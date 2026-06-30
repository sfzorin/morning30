#!/usr/bin/env python3
"""Step 1 — collect every "How to Do" video on the Leap Fitness Official channel.

Metadata only: titles + URLs + thumbnails. No video is ever downloaded.

Source preference:
  1. YouTube Data API v3 when YOUTUBE_API_KEY is set (clean, ToS-friendly).
  2. yt-dlp flat-playlist fallback otherwise (also metadata only).

Output: data/leap_howto_videos.raw.json
"""
from __future__ import annotations

import json
import os
import subprocess
import sys

import leap_common as lc


def fetch_via_ytdlp() -> list[dict]:
    """Flat-list the channel (one network pass, no per-video download)."""
    ytdlp = os.environ.get("YTDLP_BIN", "yt-dlp")
    cmd = [
        ytdlp, "--flat-playlist", "--no-warnings", "--ignore-errors",
        "--print", "%(id)s\t%(title)s\t%(duration)s\t%(upload_date)s",
        lc.CHANNEL_URL,
    ]
    print(f"[fetch] yt-dlp flat-playlist: {lc.CHANNEL_URL}", file=sys.stderr)
    proc = subprocess.run(cmd, capture_output=True, text=True)
    if proc.returncode != 0 and not proc.stdout.strip():
        sys.exit(f"[fetch] yt-dlp failed:\n{proc.stderr[-800:]}")
    rows = []
    for line in proc.stdout.splitlines():
        parts = line.split("\t")
        if len(parts) < 2 or not parts[0]:
            continue
        vid, title = parts[0], parts[1]
        duration = parts[2] if len(parts) > 2 and parts[2] != "NA" else ""
        upload = parts[3] if len(parts) > 3 and parts[3] != "NA" else ""
        rows.append({
            "video_id": vid,
            "title": title,
            "duration": duration,
            "published_at": upload,
            "source": "yt-dlp",
        })
    return rows


def fetch_via_api(api_key: str) -> list[dict]:
    import requests  # only needed on the API path

    def get(url, **params):
        params["key"] = api_key
        r = requests.get(url, params=params, timeout=30)
        r.raise_for_status()
        return r.json()

    print("[fetch] YouTube Data API", file=sys.stderr)
    # Resolve the channel's uploads playlist.
    ch = get("https://www.googleapis.com/youtube/v3/channels",
             part="contentDetails", forHandle="LeapFitnessOfficial")
    items = ch.get("items") or []
    if not items:
        sys.exit("[fetch] channel not found via API")
    uploads = items[0]["contentDetails"]["relatedPlaylists"]["uploads"]
    rows, token = [], None
    while True:
        page = get("https://www.googleapis.com/youtube/v3/playlistItems",
                   part="snippet,contentDetails", playlistId=uploads,
                   maxResults=50, pageToken=token or "")
        for it in page.get("items", []):
            sn = it["snippet"]
            rows.append({
                "video_id": it["contentDetails"]["videoId"],
                "title": sn["title"],
                "description": sn.get("description", ""),
                "published_at": sn.get("publishedAt", ""),
                "duration": "",
                "source": "youtube",
            })
        token = page.get("nextPageToken")
        if not token:
            break
    return rows


def main() -> None:
    lc.load_dotenv()
    api_key = os.environ.get("YOUTUBE_API_KEY", "").strip()
    raw = fetch_via_api(api_key) if api_key else fetch_via_ytdlp()

    total = len(raw)
    howto = [r for r in raw if lc.is_howto(r["title"])]
    out = []
    for r in howto:
        vid = r["video_id"]
        out.append({
            "video_id": vid,
            "title": r["title"],
            "url": lc.watch_url(vid),
            "description": r.get("description", ""),
            "duration": r.get("duration", ""),
            "published_at": r.get("published_at", ""),
            "channel_title": lc.CHANNEL_TITLE,
            "thumbnail": lc.thumbnail_for(vid),
            "source": "youtube" if api_key else "yt-dlp",
        })

    lc.write_json(lc.DATA / "leap_howto_videos.raw.json", out)
    print(json.dumps({
        "total_videos_on_channel": total,
        "howto_videos": len(out),
        "output": str((lc.DATA / "leap_howto_videos.raw.json").relative_to(lc.ROOT)),
    }, indent=2))


if __name__ == "__main__":
    main()
