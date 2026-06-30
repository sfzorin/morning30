"""Shared helpers for the Leap Fitness How-To import pipeline.

No third-party imports here so the deterministic steps run on a bare Python.
"""
from __future__ import annotations

import json
import os
import re
import unicodedata
from pathlib import Path

ROOT = Path(__file__).resolve().parent.parent
DATA = ROOT / "data"
CONTENT = ROOT / "content"
CHANNEL_URL = "https://www.youtube.com/@LeapFitnessOfficial/videos"
CHANNEL_TITLE = "Leap Fitness Official"

# Title prefixes that mark a "How to Do" tutorial. Matched case-insensitively,
# tolerating both the ASCII ":" and the full-width "：" the channel sometimes uses.
HOWTO_RE = re.compile(r"how\s*to\s*do", re.IGNORECASE)
PREFIX_RE = re.compile(r"^\s*how\s*to\s*do\s*[:：\-–—]?\s*", re.IGNORECASE)


def load_dotenv() -> None:
    """Minimal .env loader (avoids a hard python-dotenv dependency)."""
    env = ROOT / ".env"
    if not env.exists():
        return
    for line in env.read_text(encoding="utf-8").splitlines():
        line = line.strip()
        if not line or line.startswith("#") or "=" not in line:
            continue
        key, _, val = line.partition("=")
        os.environ.setdefault(key.strip(), val.strip())


def is_howto(title: str) -> bool:
    return bool(HOWTO_RE.search(title or ""))


def strip_emoji(text: str) -> str:
    out = []
    for ch in text:
        cat = unicodedata.category(ch)
        # Drop symbols/pictographs (So) and unassigned/private-use, keep letters,
        # marks, numbers, punctuation and separators.
        if cat in ("So", "Cs", "Co", "Cn"):
            continue
        out.append(ch)
    return "".join(out)


def clean_name(title: str) -> str:
    """Turn a raw video title into a tidy Title-Cased exercise name."""
    name = strip_emoji(title or "")
    name = PREFIX_RE.sub("", name)            # drop the "How to Do:" lead-in
    name = name.replace("：", " ").replace("|", " ")
    name = re.sub(r"\s+", " ", name).strip(" -–—:|")
    name = re.sub(r"\s+", " ", name).strip()
    return title_case(name)


_SMALL = {"and", "or", "the", "a", "an", "to", "of", "with", "on", "in", "for"}
# Words to keep upper-cased when they appear (acronyms / named holds).
_KEEP_UPPER = {"rkc", "iso", "v"}


def title_case(name: str) -> str:
    words = name.split(" ")
    out = []
    for i, w in enumerate(words):
        low = w.lower()
        bare = re.sub(r"[^a-z0-9]", "", low)
        if bare in _KEEP_UPPER:
            out.append(w.upper())
        elif i != 0 and bare in _SMALL:
            out.append(low)
        elif "-" in w:
            out.append("-".join(p[:1].upper() + p[1:].lower() if p else p for p in w.split("-")))
        else:
            out.append(w[:1].upper() + w[1:].lower() if w else w)
    return " ".join(out)


def slugify(name: str) -> str:
    s = unicodedata.normalize("NFKD", name).encode("ascii", "ignore").decode("ascii")
    s = s.lower()
    s = re.sub(r"[^a-z0-9]+", "-", s).strip("-")
    s = re.sub(r"-{2,}", "-", s)
    return s


def thumbnail_for(video_id: str) -> str:
    return f"https://i.ytimg.com/vi/{video_id}/hqdefault.jpg"


def watch_url(video_id: str) -> str:
    return f"https://www.youtube.com/watch?v={video_id}"


# ---- Grouping taxonomy ----------------------------------------------------
# Ordered: the first rule whose keyword is found wins. Broad buckets so the
# library can be browsed by category instead of scrolling 700+ items.
GROUP_RULES = [
    ("dumbbell", "Dumbbell", ["dumbbell", "db "]),
    ("equipment", "Equipment", ["kettlebell", "barbell", "resistance band", " band", "ball", "bench", "chair", "step ", "roller", "weight"]),
    ("wall", "Wall", ["wall"]),
    ("plank", "Planks", ["plank"]),
    ("push", "Push-ups & Pressing", ["push-up", "push up", "pushup", "press-up", "dip", "press"]),
    ("squat", "Squats", ["squat"]),
    ("lunge", "Lunges", ["lunge", "split squat"]),
    ("glute", "Glutes & Hips", ["glute", "bridge", "hip thrust", "hip raise", "donkey", "fire hydrant", "clamshell", "kickback"]),
    ("face", "Face & Neck", ["cheek", "jaw", "chin", "eye lift", "eye ", "lips", "lip ", "face", "marilyn", "neck firm", "double chin", "smile", "forehead", "brow"]),
    ("core", "Core & Abs", ["crunch", "sit-up", "sit up", "situp", "leg raise", "leg lift", "leg drop", "flutter", "scissor", "bicycle", "hollow", "dead bug", "v-up", "toe touch", "toe tap", "heel tap", "russian twist", "knee tuck", "knees to chest", "knee to chest", "knee hug", "windshield", "body saw", "boat", "ab ", "abs"]),
    ("cardio", "Cardio & Jumps", ["jumping jack", "burpee", "mountain climber", "high knee", "jump", "hop", "skater", "jog", "march", "shuffle", "skip", "star "]),
    ("back", "Back & Posture", ["superman", "cobra", "swimmer", "snow angel", "y-raise", "t-raise", "w-raise", "prone", "reverse fly", "good morning", "bird dog"]),
    ("legs", "Legs & Calves", ["calf", "calve", "leg curl", "leg extension", "wall sit", "step-up", "step up"]),
    ("arms", "Arms & Shoulders", ["curl", "tricep", "bicep", "shoulder", "arm circle", "lateral raise", "front raise", "pike"]),
    ("yoga", "Yoga & Poses", ["pose", "warrior", "down dog", "downward", "upward", "pigeon", "lotus", "namaskar", "mudra", "asana", "crescent", "garland", "locust", "triangle", "thunderbolt", "sphinx", "happy baby", "eagle", "gate", "camel", "child", "savasana", "salutation", "cat cow", "cat-cow", "thread the needle", "half moon", "cow face", "lizard", "frog", "seated forward", "standing forward", "side angle", "forward bend", "spinal roll", "boat twist", "bow", "bridge pose"]),
    ("mobility", "Mobility & Stretch", ["stretch", "twist", "rotation", "circle", "swing", "reach", "mobility", "warm-up", "warm up", "breathing", "neck", "open"]),
]
DEFAULT_GROUP = ("other", "Other")


def classify_group(name: str):
    low = name.lower()
    for key, label, kws in GROUP_RULES:
        for kw in kws:
            if kw in low:
                return key, label
    return DEFAULT_GROUP


# ---- Banned marketing/medical phrases (validation) ------------------------
BANNED = [
    "no pain no gain",
    "burn fat fast",
    "guaranteed",
    "cure",
    "medical treatment",
    "six pack in 30 days",
]


def read_json(path: Path):
    return json.loads(Path(path).read_text(encoding="utf-8"))


def write_json(path: Path, obj) -> None:
    Path(path).parent.mkdir(parents=True, exist_ok=True)
    Path(path).write_text(json.dumps(obj, ensure_ascii=False, indent=2) + "\n", encoding="utf-8")
