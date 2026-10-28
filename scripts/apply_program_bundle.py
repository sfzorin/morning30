#!/usr/bin/env python3
"""Validate/apply a personal program bundle without seeding other accounts.

Dry run is the default. Apply only to an explicitly identified existing account,
back up its program fields, and change those fields atomically. Progress, tokens,
preferences, and other accounts are never written. With --images-only, preserve
the saved program and all exercise fields except supplied illustration metadata.
"""
import argparse
import datetime
import json
import os
import re
from pathlib import Path
import sqlite3
import struct


LANGUAGES = {"ru", "en", "tr", "de", "es", "fr", "it"}


def validate_reaction(doc):
    """Reaction cues use player timing; exercise authors supply directions only."""
    reaction = doc.get("reaction")
    if reaction is None:
        return
    if not isinstance(reaction, dict) or set(reaction) != {"directions"}:
        raise ValueError("Invalid reaction")
    directions = reaction["directions"]
    if (doc.get("unit") != "seconds" or doc.get("perSide") is not False
            or not isinstance(directions, list) or len(directions) != 2
            or any(not isinstance(direction, str) for direction in directions)
            or set(directions) != {"left", "right"}):
        raise ValueError("Invalid reaction")


def animation_poses(doc):
    animation = doc.get("animation")
    if animation is None:
        return []
    if not isinstance(animation, dict):
        raise ValueError("Invalid animation")
    if "otherPose" in animation and not isinstance(animation["otherPose"], str):
        raise ValueError("Invalid animation")
    poses = [animation.get("pose")]
    if animation.get("otherPose"):
        poses.append(animation["otherPose"])
    frames, kind = animation.get("frames"), animation.get("kind")
    if (type(frames) is not int or frames not in {2, 3}
            or not isinstance(kind, str) or kind not in {"movement", "alternating", "sequence"}
            or (kind == "alternating" and frames != 3)
            or any(not isinstance(pose, str) or not re.fullmatch(r"[a-z0-9][a-z0-9_-]{0,63}", pose) for pose in poses)):
        raise ValueError("Invalid animation")
    return poses


def validate_strip(data, frames):
    if len(data) < 24 or data[:8] != b"\x89PNG\r\n\x1a\n" or data[12:16] != b"IHDR":
        raise ValueError("Invalid animation PNG")
    width, height = struct.unpack(">II", data[16:24])
    if height < 64 or width != frames * height:
        raise ValueError("Animation must contain equal square frames")


def load_bundle(folder):
    folder = Path(folder)
    program = json.loads((folder / "program.json").read_text())
    docs = json.loads((folder / "exercises.json").read_text())
    exercises = {d["id"]: d for d in docs}
    if len(exercises) != len(docs):
        raise ValueError("Duplicate exercise IDs")
    for doc in docs:
        animation_poses(doc)
        validate_reaction(doc)
        if doc["unit"] not in {"seconds", "reps", "breaths"}:
            raise ValueError(f"Invalid unit: {doc['id']}")
        if set(doc["content"]) != LANGUAGES:
            raise ValueError(f"Missing translations: {doc['id']}")
        for lang, loc in doc["content"].items():
            for field in ("name", "hint", "warning", "desc", "howTo", "correct", "wrong", "breathing", "safety"):
                if not loc.get(field):
                    raise ValueError(f"Missing {doc['id']}.{lang}.{field}")
            for field in ("start", "mid", "finish"):
                if not loc.get("voice", {}).get(field):
                    raise ValueError(f"Missing {doc['id']}.{lang}.voice.{field}")
    if program.get("totalDays") != 30 or len(program["days"]) != 30:
        raise ValueError("Expected a 30-day cycle")
    if program.get("warmupRounds", 1) != 1 or not program.get("warmup") or not program.get("cooldown"):
        raise ValueError("Every program needs one warm-up and one cool-down")
    for number, day in enumerate(program["days"], 1):
        if day["day"] != number or day["rounds"] != 2 or not day["items"]:
            raise ValueError("Invalid day numbering/rounds")
        if not day.get("warmup", program["warmup"]) or not day.get("cooldown", program["cooldown"]):
            raise ValueError("Every day needs a warm-up and cool-down")
        items = day.get("warmup", program["warmup"]) + day["items"] + day.get("cooldown", program["cooldown"])
        for item in items:
            doc = exercises[item["id"]]
            if item["unit"] != doc["unit"] or item["perSide"] != doc["perSide"]:
                raise ValueError(f"Definition mismatch: {item['id']}")
            if not isinstance(item["value"], int) or item["value"] <= 0:
                raise ValueError("Invalid exercise count")
            for field in ("rest", "sideRest"):
                if field in item:
                    raise ValueError("Exercise pauses belong to user settings, not the program")
    manifest_path = folder / "images.json"
    images = json.loads(manifest_path.read_text()) if manifest_path.exists() else []
    for image in images:
        path = (folder / image["file"]).resolve()
        if not path.is_relative_to(folder.resolve()):
            raise ValueError("Image must be inside bundle")
        data = path.read_bytes()
        if not data.startswith(b"\x89PNG\r\n\x1a\n") or len(data) > 4 * 1024 * 1024:
            raise ValueError(f"Invalid PNG: {path}")
        image["data"] = data
    by_id = {image["id"]: image["data"] for image in images}
    for doc in docs:
        for pose in animation_poses(doc):
            if pose in by_id:
                validate_strip(by_id[pose], doc["animation"]["frames"])
    return program, exercises, images


def apply_bundle(connection, account, expected_name, program, exercises, images, backup_path, *, images_only=False):
    connection.execute("BEGIN IMMEDIATE")
    try:
        row = connection.execute(
            "SELECT id,name,program_json,custom_exercises,active_program,position FROM users WHERE id=?",
            (account,),
        ).fetchone()
        if row is None or row["name"] != expected_name:
            raise ValueError("Target account does not match the expected name")
        current = json.loads(row["custom_exercises"] or "{}")
        if images_only:
            # Validate every assignment before creating a backup or inserting any
            # pictures. Stage two must not create or replace exercise content.
            for exercise_id, doc in exercises.items():
                if exercise_id not in current or not isinstance(current[exercise_id], dict):
                    raise ValueError(f"Existing exercise not found: {exercise_id}")
                poses = doc.get("poses")
                if not isinstance(poses, list) or not poses or not all(
                    isinstance(pose, str) and pose.strip() for pose in poses
                ):
                    raise ValueError(f"Nonempty pose assignment required: {exercise_id}")
            for exercise_id, doc in exercises.items():
                current[exercise_id]["poses"] = list(doc["poses"])
                if "animation" in doc:
                    current[exercise_id]["animation"] = doc["animation"]
        else:
            current.update(exercises)
        supplied = {image["id"] for image in images}
        supplied_data = {image["id"]: image["data"] for image in images}
        for doc in exercises.values():
            if not images_only:
                validate_reaction(doc)
            strips = animation_poses(doc)
            for pose in doc.get("poses", []) + strips:
                if pose not in supplied and not connection.execute(
                    "SELECT 1 FROM pose_images WHERE pose_id=? LIMIT 1", (pose,)
                ).fetchone():
                    raise ValueError(f"Pose image not found: {pose}")
            for pose in strips:
                data = supplied_data.get(pose)
                if data is None:
                    data = connection.execute("SELECT data FROM pose_images WHERE pose_id=? LIMIT 1", (pose,)).fetchone()[0]
                validate_strip(data, doc["animation"]["frames"])
        backup = dict(row)
        backup["saved_at"] = datetime.datetime.now(datetime.timezone.utc).isoformat()
        backup_path = Path(backup_path)
        fd = os.open(backup_path, os.O_WRONLY | os.O_CREAT | os.O_EXCL, 0o600)
        with os.fdopen(fd, "w") as file:
            json.dump(backup, file, ensure_ascii=False, indent=2)
            file.flush()
            os.fsync(file.fileno())
        for image in images:
            # Bundle pose IDs are immutable. Refuse collisions instead of
            # replacing pictures another account may already be using.
            old = connection.execute(
                "SELECT data FROM pose_images WHERE pose_id=? AND sex=?",
                (image["id"], image.get("sex", "m")),
            ).fetchone()
            if old is not None and old[0] != image["data"]:
                raise ValueError(f"Existing pose differs: {image['id']}")
            connection.execute(
                "INSERT OR IGNORE INTO poses(id,kind,label) VALUES(?,?,?)",
                (image["id"], image.get("kind", "tennis"), image.get("label", image["id"])),
            )
            connection.execute(
                "INSERT OR IGNORE INTO pose_images(pose_id,sex,mime,data) VALUES(?,?,?,?)",
                (image["id"], image.get("sex", "m"), "image/png", image["data"]),
            )
        if images_only:
            connection.execute(
                "UPDATE users SET custom_exercises=? WHERE id=?",
                (json.dumps(current, ensure_ascii=False), account),
            )
        else:
            connection.execute(
                "UPDATE users SET program_json=?,custom_exercises=?,active_program='custom' WHERE id=?",
                (json.dumps(program, ensure_ascii=False), json.dumps(current, ensure_ascii=False), account),
            )
        connection.commit()
        return row["position"]
    except Exception:
        connection.rollback()
        raise


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("bundle", type=Path)
    parser.add_argument("--db", type=Path)
    parser.add_argument("--account", type=int)
    parser.add_argument("--expected-name")
    parser.add_argument("--apply", action="store_true")
    parser.add_argument(
        "--images-only", action="store_true",
        help="Only assign illustrations to existing exercises; preserve the saved program and exercise content",
    )
    parser.add_argument("--backup", type=Path)
    args = parser.parse_args()
    program, exercises, images = load_bundle(args.bundle)
    print(f"Validated: {len(exercises)} exercises, 7 languages, 30 days, {len(images)} images")
    if not args.apply:
        return
    if not all((args.db, args.account, args.expected_name, args.backup)):
        parser.error("--apply requires --db, --account, --expected-name and --backup")
    connection = sqlite3.connect(f"{args.db.resolve().as_uri()}?mode=rw", uri=True, timeout=10)
    connection.row_factory = sqlite3.Row
    try:
        position = apply_bundle(
            connection, args.account, args.expected_name, program, exercises, images, args.backup,
            images_only=args.images_only,
        )
        print(f"Applied to account {args.account}; position {position} preserved; backup: {args.backup}")
    finally:
        connection.close()


if __name__ == "__main__":
    main()
