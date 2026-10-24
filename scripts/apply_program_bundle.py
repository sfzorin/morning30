#!/usr/bin/env python3
"""Validate/apply a personal program bundle without seeding other accounts.

Dry run is the default. Apply only to an explicitly identified existing account,
back up its program fields, and change those fields atomically. Progress, tokens,
preferences, and other accounts are never written.
"""
import argparse
import datetime
import json
import os
from pathlib import Path
import sqlite3


LANGUAGES = {"ru", "en", "tr", "de", "es", "fr", "it"}


def load_bundle(folder):
    folder = Path(folder)
    program = json.loads((folder / "program.json").read_text())
    docs = json.loads((folder / "exercises.json").read_text())
    exercises = {d["id"]: d for d in docs}
    if len(exercises) != len(docs):
        raise ValueError("Duplicate exercise IDs")
    for doc in docs:
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
    for number, day in enumerate(program["days"], 1):
        if day["day"] != number or day["rounds"] < 1:
            raise ValueError("Invalid day numbering/rounds")
        items = day.get("warmup", program["warmup"]) + day["items"] + day.get("cooldown", program["cooldown"])
        for item in items:
            doc = exercises[item["id"]]
            if item["unit"] != doc["unit"] or item["perSide"] != doc["perSide"]:
                raise ValueError(f"Definition mismatch: {item['id']}")
            if not isinstance(item["value"], int) or item["value"] <= 0:
                raise ValueError("Invalid exercise count")
            for field in ("rest", "sideRest"):
                if field in item and not 0 <= item[field] <= 90:
                    raise ValueError(f"Invalid {field}")
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
    return program, exercises, images


def apply_bundle(connection, account, expected_name, program, exercises, images, backup_path):
    connection.execute("BEGIN IMMEDIATE")
    try:
        row = connection.execute(
            "SELECT id,name,program_json,custom_exercises,active_program,position FROM users WHERE id=?",
            (account,),
        ).fetchone()
        if row is None or row["name"] != expected_name:
            raise ValueError("Target account does not match the expected name")
        current = json.loads(row["custom_exercises"] or "{}")
        current.update(exercises)
        supplied = {image["id"] for image in images}
        for doc in exercises.values():
            for pose in doc.get("poses", []):
                if pose not in supplied and not connection.execute(
                    "SELECT 1 FROM pose_images WHERE pose_id=? LIMIT 1", (pose,)
                ).fetchone():
                    raise ValueError(f"Pose image not found: {pose}")
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
        position = apply_bundle(connection, args.account, args.expected_name, program, exercises, images, args.backup)
        print(f"Applied to account {args.account}; position {position} preserved; backup: {args.backup}")
    finally:
        connection.close()


if __name__ == "__main__":
    main()
