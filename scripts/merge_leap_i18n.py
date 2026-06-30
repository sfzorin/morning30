#!/usr/bin/env python3
"""Merge the English cards with the 6 translations into one localized catalogue.

Input:
  content/leap_exercise_cards.en.json        (English, authoritative structure)
  data/leap_i18n_out/*.json                  (ru/tr/de/es/fr/it per exercise)
Output:
  content/leap_exercise_cards.i18n.json      (all 7 languages per card)

Language order matches the app's i18n tuple: [ru, en, tr, de, es, fr, it].
"""
from __future__ import annotations

import glob
import json

import leap_common as lc

LANGS6 = ["ru", "tr", "de", "es", "fr", "it"]
ORDER = ["ru", "en", "tr", "de", "es", "fr", "it"]
EN = lc.CONTENT / "leap_exercise_cards.en.json"
OUT = lc.CONTENT / "leap_exercise_cards.i18n.json"


def main() -> None:
    en = {c["exercise_id"]: c for c in lc.read_json(EN)}
    tr: dict[str, dict] = {}
    for f in glob.glob(str(lc.DATA / "leap_i18n_out" / "*.json")):
        for c in lc.read_json(f):
            tr[c["exercise_id"]] = c

    out = []
    for eid, c in en.items():
        t = tr.get(eid, {})
        names = {"en": c["exercise_name"]}
        cards = {"en": {k: c[k] for k in ("short_description", "how_to", "correct", "mistakes")}}
        for L in LANGS6:
            d = t.get(L, {})
            names[L] = d.get("name", c["exercise_name"])
            cards[L] = {
                "short_description": d.get("short_description", ""),
                "how_to": d.get("how_to", []),
                "correct": d.get("correct", []),
                "mistakes": d.get("mistakes", []),
            }
        out.append({
            "exercise_id": eid,
            "slug": c["slug"],
            "group": c["group"],
            "group_label": c["group_label"],
            "names": {k: names[k] for k in ORDER},
            "cards": {k: cards[k] for k in ORDER},
            "source": c["source"],
            "review_status": c["review_status"],
        })

    lc.write_json(OUT, out)
    print(json.dumps({
        "cards": len(out),
        "languages": ORDER,
        "output": str(OUT.relative_to(lc.ROOT)),
    }, indent=2))


if __name__ == "__main__":
    main()
