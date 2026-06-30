#!/usr/bin/env python3
"""Step 2 — normalize the raw videos into deduplicated exercise records.

For each video: clean the title into an exercise name, derive slug + id, assign a
browse group, and — the key step for this project — try to MATCH it against the
app's existing 70 exercises so matched ones only need a video (no new card).

Input:  data/leap_howto_videos.raw.json
        data/existing_exercises.json   (id -> {ru, en})
Output: data/leap_howto_exercises.normalized.json
"""
from __future__ import annotations

import json
import re
import sys

import leap_common as lc

# Known phrasing equivalences so obvious matches aren't missed by token compare.
SYNONYMS = {
    "pushups": "push up", "pushup": "push up", "push ups": "push up",
    "pressup": "push up", "press up": "push up", "press ups": "push up",
    "situps": "sit up", "situp": "sit up",
    "jumping jack": "jumping jacks",
    "mountain climber": "mountain climbers",
    "flutter kick": "flutter kicks",
    "bicycle crunch": "bicycle crunches", "bicycle": "bicycle crunches",
    "glute bridge": "bridge", "hip bridge": "bridge",
    "air squat": "squat", "bodyweight squat": "squat",
}

_STOP = {"the", "a", "an", "with", "and"}


def _newer(a: dict, b: dict) -> bool:
    """True if video a is fresher than b. Prefer a real publish date; when dates
    are missing fall back to channel order (raw list is newest-first, so the
    already-stored b came first = newer, and a should NOT replace it)."""
    da, db = a.get("published_at", ""), b.get("published_at", "")
    if da and db:
        return da > db
    return bool(da) and not db


def canon_tokens(name: str) -> frozenset[str]:
    s = name.lower()
    s = SYNONYMS.get(s.strip(), s)
    s = re.sub(r"[^a-z0-9 ]+", " ", s)
    toks = []
    for w in s.split():
        if w in _STOP:
            continue
        w = SYNONYMS.get(w, w)
        if len(w) > 3 and w.endswith("s"):   # light de-pluralization
            w = w[:-1]
        toks.append(w)
    return frozenset(toks)


def build_existing_index(existing: dict) -> dict:
    idx = {}
    for eid, names in existing.items():
        idx[canon_tokens(names["en"])] = (eid, names["en"])
    return idx


def match_existing(name: str, idx: dict):
    """Exact token-set match only. Variants/compounds ("Jumping Push-Ups",
    "Dumbbell Hip Hinge") are deliberately NOT matched — they get their own
    card; only the genuinely same movement reuses an existing exercise."""
    toks = canon_tokens(name)
    if not toks:
        return None
    return idx.get(toks)


def main() -> None:
    raw = lc.read_json(lc.DATA / "leap_howto_videos.raw.json")
    existing = lc.read_json(lc.DATA / "existing_exercises.json")
    idx = build_existing_index(existing)

    by_slug: dict[str, dict] = {}
    order: list[str] = []
    matched_ct = 0
    for v in raw:
        name = lc.clean_name(v["title"])
        slug = lc.slugify(name)
        if not slug:
            continue
        src = {
            "video_id": v["video_id"],
            "video_title": v["title"],
            "video_url": v["url"],
            "thumbnail": v.get("thumbnail", ""),
            "published_at": v.get("published_at", ""),
        }
        if slug in by_slug:                       # dedup: keep the FRESHEST video
            rec = by_slug[slug]
            if _newer(src, rec["source"]):
                rec["duplicate_sources"].append(rec["source"])
                rec["source"] = src
            else:
                rec["duplicate_sources"].append(src)
            continue
        gkey, glabel = lc.classify_group(name)
        m = match_existing(name, idx)
        rec = {
            "exercise_id": f"leap_{slug.replace('-', '_')}",
            "exercise_name": name,
            "slug": slug,
            "group": gkey,
            "group_label": glabel,
            "matched_exercise_id": m[0] if m else None,
            "matched_exercise_name": m[1] if m else None,
            "source": src,
            "duplicate_sources": [],
        }
        if m:
            matched_ct += 1
        by_slug[slug] = rec
        order.append(slug)

    records = [by_slug[s] for s in order]
    lc.write_json(lc.DATA / "leap_howto_exercises.normalized.json", records)

    dupes = sum(len(r["duplicate_sources"]) for r in records)
    groups: dict[str, int] = {}
    for r in records:
        groups[r["group_label"]] = groups.get(r["group_label"], 0) + 1
    print(json.dumps({
        "videos_in": len(raw),
        "unique_exercises": len(records),
        "duplicate_videos_folded": dupes,
        "matched_to_existing": matched_ct,
        "needs_new_card": len(records) - matched_ct,
        "groups": dict(sorted(groups.items(), key=lambda kv: -kv[1])),
        "output": str((lc.DATA / "leap_howto_exercises.normalized.json").relative_to(lc.ROOT)),
    }, indent=2, ensure_ascii=False))


if __name__ == "__main__":
    main()
