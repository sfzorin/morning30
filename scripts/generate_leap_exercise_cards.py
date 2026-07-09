#!/usr/bin/env python3
"""Step 3 — assemble the final exercise cards.

Text layers, in priority order:
  1. Authored drafts in data/leap_gen_out/*.json (the house-style pass; on this
     project they are written by review-grade model batches).
  2. OpenAI (if OPENAI_API_KEY is set and a card has no authored draft).
  3. A safe generic template keyed by the exercise group (never fails).

Matched exercises (already in the app's library) are NOT re-carded — they only
need a video, written to data/leap_video_map.json.

Outputs:
  content/leap_exercise_cards.en.json
  content/leap_exercise_cards.en.md
  content/leap_exercise_cards.review.csv
  data/leap_video_map.json
"""
from __future__ import annotations

import csv
import glob
import json
import os

import leap_common as lc

NORMALIZED = lc.DATA / "leap_howto_exercises.normalized.json"
GEN_OUT = lc.DATA / "leap_gen_out"
OUT_JSON = lc.CONTENT / "leap_exercise_cards.en.json"
OUT_MD = lc.CONTENT / "leap_exercise_cards.en.md"
OUT_CSV = lc.CONTENT / "leap_exercise_cards.review.csv"
VIDEO_MAP = lc.DATA / "leap_video_map.json"


def load_authored() -> dict:
    text = {}
    for f in sorted(glob.glob(str(GEN_OUT / "batch_*.json"))):
        try:
            for c in json.load(open(f, encoding="utf-8")):
                text[c["exercise_id"]] = c
        except Exception:
            continue
    return text


def template_card(name: str, group_key: str) -> dict:
    """Generic, safe draft when no authored/AI text exists."""
    return {
        "short_description": f"A bodyweight exercise: {name.lower()}.",
        "how_to": [
            "Set up in a stable starting position.",
            "Move through the full range with control.",
            "Return slowly to the start.",
            "Keep a steady, even pace.",
        ],
        "correct": [
            "Keep your core engaged.",
            "Move with control.",
            "Breathe steadily.",
        ],
        "mistakes": [
            "Don't rush the movement.",
            "Don't hold your breath.",
            "Don't lose your form when tired.",
        ],
    }


def main() -> None:
    lc.load_dotenv()
    norm = lc.read_json(NORMALIZED)
    authored = load_authored()
    use_openai = bool(os.environ.get("OPENAI_API_KEY", "").strip()) and not authored

    cards, video_map = [], []
    for r in norm:
        if r.get("matched_exercise_id"):
            video_map.append({
                "exercise_id": r["matched_exercise_id"],
                "matched_from": r["exercise_name"],
                "video_url": r["source"]["video_url"],
                "video_id": r["source"]["video_id"],
                "video_title": r["source"]["video_title"],
            })
            continue

        eid = r["exercise_id"]
        if eid in authored:
            t = authored[eid]
            status = "needs_human_review"
        elif use_openai:
            t = _openai_card(r["exercise_name"], r["source"]["video_title"])  # noqa
            status = "needs_human_review"
        else:
            t = template_card(r["exercise_name"], r["group"])
            status = "needs_manual_review"

        cards.append({
            "exercise_id": eid,
            "exercise_name": r["exercise_name"],
            "slug": r["slug"],
            "group": r["group"],
            "group_label": r["group_label"],
            "short_description": t["short_description"],
            "how_to": t["how_to"],
            "correct": t["correct"],
            "mistakes": t["mistakes"],
            "source": {
                "video_title": r["source"]["video_title"],
                "video_url": r["source"]["video_url"],
                "video_id": r["source"]["video_id"],
                "thumbnail": r["source"].get("thumbnail", ""),
            },
            "review_status": status,
        })

    lc.write_json(OUT_JSON, cards)
    lc.write_json(VIDEO_MAP, video_map)
    # Embeddable copy for the Go app (assets/ is the go:embed root).
    lc.write_json(lc.ROOT / "assets" / "leap_video_map.json", video_map)
    _write_md(cards)
    _write_csv(cards)

    needs = sum(1 for c in cards if c["review_status"] != "ok")
    print(json.dumps({
        "cards_written": len(cards),
        "video_map_matched": len(video_map),
        "needs_review": needs,
        "text_source": "authored" if authored else ("openai" if use_openai else "template"),
        "outputs": [str(p.relative_to(lc.ROOT)) for p in (OUT_JSON, OUT_MD, OUT_CSV, VIDEO_MAP)],
    }, indent=2))


def _write_md(cards: list) -> None:
    by_group: dict[str, list] = {}
    for c in cards:
        by_group.setdefault(c["group_label"], []).append(c)
    lines = ["# Leap Fitness — exercise cards (English)", "",
             f"{len(cards)} cards, grouped. All drafts — review before publishing.", ""]
    for group in sorted(by_group):
        lines.append(f"# {group}\n")
        for c in by_group[group]:
            lines.append(f"## {c['exercise_name']}\n")
            lines.append(c["short_description"] + "\n")
            lines.append("**How to**")
            lines += c["how_to"]
            lines.append("\n**Correct**")
            lines += c["correct"]
            lines.append("\n**Mistakes**")
            lines += c["mistakes"]
            lines.append("\n**Source**")
            lines.append(f"Video: {c['source']['video_title']}")
            lines.append(f"URL: {c['source']['video_url']}\n")
    OUT_MD.parent.mkdir(parents=True, exist_ok=True)
    OUT_MD.write_text("\n".join(lines) + "\n", encoding="utf-8")


def _write_csv(cards: list) -> None:
    OUT_CSV.parent.mkdir(parents=True, exist_ok=True)
    with open(OUT_CSV, "w", newline="", encoding="utf-8") as fh:
        w = csv.writer(fh)
        w.writerow(["exercise_id", "exercise_name", "group", "video_url", "review_status"])
        for c in cards:
            w.writerow([c["exercise_id"], c["exercise_name"], c["group_label"],
                        c["source"]["video_url"], c["review_status"]])


def _openai_card(name: str, video_title: str) -> dict:
    """Best-effort single-card generation via OpenAI (only when no authored draft
    exists and OPENAI_API_KEY is set). Kept minimal; the authored pass is primary."""
    import requests
    guide = (lc.DATA / "leap_gen_GUIDE.md").read_text(encoding="utf-8")
    prompt = (f"{guide}\n\nWrite ONE card as a JSON object for:\n"
              f"exercise_name: {name}\nvideo_title: {video_title}\n"
              "Return only the JSON object.")
    r = requests.post(
        "https://api.openai.com/v1/chat/completions",
        headers={"Authorization": f"Bearer {os.environ['OPENAI_API_KEY']}"},
        json={"model": "gpt-4o-mini", "temperature": 0.4,
              "messages": [{"role": "user", "content": prompt}],
              "response_format": {"type": "json_object"}},
        timeout=60)
    r.raise_for_status()
    obj = json.loads(r.json()["choices"][0]["message"]["content"])
    return {k: obj[k] for k in ("short_description", "how_to", "correct", "mistakes")}


if __name__ == "__main__":
    main()
