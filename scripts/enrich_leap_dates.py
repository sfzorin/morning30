#!/usr/bin/env python3
"""Optional enrichment — fill published_at + duration on the raw videos.

The flat-playlist crawl is fast but omits dates, so the dedup step falls back to
channel order (newest first). This pass fetches the real upload date + duration
per video (metadata only, NO download) so dedup can keep the genuinely freshest
duplicate, and reviewers see real dates. Progress is written incrementally so the
run is resumable.

Usage: python scripts/enrich_leap_dates.py
Output: rewrites data/leap_howto_videos.raw.json (published_at/duration filled)
"""
from __future__ import annotations

import json
import os
import subprocess
import sys
from pathlib import Path

import leap_common as lc

RAW = lc.DATA / "leap_howto_videos.raw.json"


def ytdlp_bin() -> str:
    cand = lc.ROOT / ".venv" / "bin" / "yt-dlp"
    return str(cand) if cand.exists() else os.environ.get("YTDLP_BIN", "yt-dlp")


def main() -> None:
    raw = lc.read_json(RAW)
    todo = [v for v in raw if not v.get("published_at")]
    print(f"[enrich] {len(todo)}/{len(raw)} videos need dates", file=sys.stderr)
    if not todo:
        return
    by_id = {v["video_id"]: v for v in raw}
    urls = [v["url"] for v in todo]
    proc = subprocess.Popen(
        [ytdlp_bin(), "--skip-download", "--no-warnings", "--ignore-errors",
         "--sleep-requests", "1.5", "--retries", "3",  # avoid YouTube rate-limit
         "--print", "%(id)s|%(upload_date)s|%(duration)s", *urls],
        stdout=subprocess.PIPE, text=True,
    )
    done = 0
    for line in proc.stdout:                       # type: ignore[union-attr]
        parts = line.strip().split("|")
        if len(parts) < 2 or parts[0] not in by_id:
            continue
        vid, up = parts[0], parts[1]
        dur = parts[2] if len(parts) > 2 and parts[2] != "NA" else ""
        if up and up != "NA":
            # YYYYMMDD -> YYYY-MM-DD
            by_id[vid]["published_at"] = f"{up[:4]}-{up[4:6]}-{up[6:8]}" if len(up) == 8 else up
        by_id[vid]["duration"] = dur
        done += 1
        if done % 25 == 0:
            lc.write_json(RAW, raw)
            print(f"[enrich] {done}/{len(todo)}", file=sys.stderr)
    proc.wait()
    lc.write_json(RAW, raw)
    print(f"[enrich] done: {done} videos updated", file=sys.stderr)


if __name__ == "__main__":
    main()
