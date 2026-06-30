#!/usr/bin/env python3
"""Step 4 — validate the assembled cards.

Checks each card for required fields, line-count bounds, line emptiness, the
YouTube URL shape, id/slug uniqueness, banned phrases and review_status.

Outputs:
  content/leap_exercise_cards.validation.json
  content/leap_exercise_cards.review.csv   (exercise_id,exercise_name,video_url,status,issues)
"""
from __future__ import annotations

import csv
import json

import leap_common as lc

CARDS = lc.CONTENT / "leap_exercise_cards.en.json"
VALID_OUT = lc.CONTENT / "leap_exercise_cards.validation.json"
CSV_OUT = lc.CONTENT / "leap_exercise_cards.review.csv"
URL_PREFIX = "https://www.youtube.com/watch?v="


def check(card: dict, seen_ids: set, seen_slugs: set) -> list[str]:
    issues: list[str] = []

    def need(field):
        if not card.get(field):
            issues.append(f"missing:{field}")

    for f in ("exercise_id", "exercise_name", "slug", "short_description", "review_status"):
        need(f)

    def bounded(field, lo, hi):
        arr = card.get(field) or []
        if not (lo <= len(arr) <= hi):
            issues.append(f"{field}:count={len(arr)}(want {lo}-{hi})")
        if any(not str(s).strip() for s in arr):
            issues.append(f"{field}:empty_line")

    bounded("how_to", 3, 6)
    bounded("correct", 3, 5)
    bounded("mistakes", 3, 5)

    url = (card.get("source") or {}).get("video_url", "")
    if not url:
        issues.append("missing:source.video_url")
    elif not url.startswith(URL_PREFIX):
        issues.append("bad:video_url")

    eid, slug = card.get("exercise_id"), card.get("slug")
    if eid in seen_ids:
        issues.append("dup:exercise_id")
    if slug in seen_slugs:
        issues.append("dup:slug")

    blob = json.dumps(card, ensure_ascii=False).lower()
    for b in lc.BANNED:
        if b in blob:
            issues.append(f"banned:{b}")
    return issues


def main() -> None:
    cards = lc.read_json(CARDS)
    seen_ids: set = set()
    seen_slugs: set = set()
    rows, results = [], []
    ok = 0
    for c in cards:
        issues = check(c, seen_ids, seen_slugs)
        seen_ids.add(c.get("exercise_id"))
        seen_slugs.add(c.get("slug"))
        status = "ok" if not issues else "issues"
        if status == "ok":
            ok += 1
        results.append({"exercise_id": c.get("exercise_id"), "status": status, "issues": issues})
        rows.append([c.get("exercise_id"), c.get("exercise_name"),
                     (c.get("source") or {}).get("video_url", ""), status, "; ".join(issues)])

    summary = {
        "total": len(cards),
        "ok": ok,
        "with_issues": len(cards) - ok,
        "results": results,
    }
    lc.write_json(VALID_OUT, summary)
    CSV_OUT.parent.mkdir(parents=True, exist_ok=True)
    with open(CSV_OUT, "w", newline="", encoding="utf-8") as fh:
        w = csv.writer(fh)
        w.writerow(["exercise_id", "exercise_name", "video_url", "status", "issues"])
        w.writerows(rows)

    print(json.dumps({k: summary[k] for k in ("total", "ok", "with_issues")}, indent=2))
    print(f"wrote {VALID_OUT.relative_to(lc.ROOT)} and {CSV_OUT.relative_to(lc.ROOT)}")


if __name__ == "__main__":
    main()
