# morning30.com

A mobile-first PWA for a **30-day morning-exercise** habit: a daily floor-and-mat
routine you can run with the screen in your pocket, guided entirely by voice.

- **Repository:** <https://github.com/sfzorin/morning30>
- **License:** [Apache-2.0](LICENSE)
- **Status:** active development

Each day is a single session — **warm-up once → two complete main rounds → cool-down once** —
on a 30-day cycle. New accounts run a short public **demo**. Everyone else keeps
their own plan as JSON on the account (no named picker). Progress, streaks and
preferences are stored per user on the server.

Built with the [doors](https://github.com/doors-dev/doors) Go framework and its
GoX templates (server-driven reactive UI), a pure-Go SQLite database, and a
small client-side workout engine.

---

## Features

- **One public demo, personal plans.** Guests and new accounts run a short
  no-jump demo cycle. Existing named plans (Sergey / Ilya / Vlad) were copied
  into that account's personal JSON on first load; cycle position and streaks
  stay put. Settings edit *this* account only — there is no Sergey/Vlad/Ilya
  picker.
- **In-app editor.** Reorder, add, remove and retime exercises; all three phases
  remain nonempty, with exactly two main rounds. Save as this account's version,
  or restore the demo.
- **Exercise JSON.** Download any exercise as a file, paste or upload one back.
  Custom exercises overlay the built-in library. Each exercise can use one
  catalog still (static) or two (they swap once a second). Legacy built-in
  PNGs remain the fallback until photos are assigned.
- **MCP.** Agents write exercises and the 30-day cycle straight into the user's
  SQLite row (`custom_exercises` / `program_json`). Same data the editor and
  file import use. Progress is never reset. They also upload stills into the
  shared pose catalog and pick 1–2 photos per exercise.
- **Universal difficulty.** A per-user level from **−3 to +3** scales every
  program by ±10% per step (reps and seconds only; breaths and warm-up/cool-down
  stay fixed). Prescribed programs can lock their doses with `fixedDifficulty`.
- **One rest policy.** Account settings control every pause between exercises:
  warm-up and cool-down 0–40 seconds, main 5–40. A pause belongs to the upcoming
  exercise's phase; the final exercise has no following pause. The same resolved
  values drive estimates and player timers. Program JSON cannot prescribe pauses;
  old `rest`/`sideRest` fields are ignored on import and omitted on export.
- **Daily phases.** Every program includes a warm-up, two main rounds and a
  cool-down. Empty day phases inherit the shared phase; missing shared phases
  and empty main sequences use demo defaults. Legacy identical expanded rounds
  collapse to one sequence; asymmetric old rounds are rejected visibly on import.
  `repSeconds` only estimates duration. `selfManagedSides` uses one Done tap for
  both sides, or a combined hold timer with the time per side shown.
  See [Ilya's progressive 30-day bundle](programs/ilya-tennis/README.md).
- **Two exercise types.** Reps (tap *Done*) and timed (countdown). Counts are
  shown as-is; timed per-side exercises include a shared three-second spoken
  side-change countdown between working halves. Rep sets use a single Done tap
  after both sides.
- **Voice-first player (client-side).** Progress bar, pause / skip / back, rest
  periods, screen wake-lock, and a celebratory finish (applause, confetti,
  localized encouragement). The player runs entirely in the browser from a
  server-provided JSON payload. Leaving the app pauses the workout; returning
  or reopening offers to resume the saved exercise, side and timer.
- **Reliable completion.** Results are queued on the device before sending to
  `/api/workouts/complete`. Retries retain the original local date and cannot
  advance the cycle twice. The player confirms the streak only after the server
  has saved it; the home page also retries pending results.
- **Spoken guidance** via the browser Web Speech API: a start cue, a rest
  announcement of the next exercise with its count, a 3-2-1 lead-in, a half-time
  marker on timed sets, and running technique narration (how to do it / common
  mistakes). Four levels in settings: off · minimal · normal · detailed.
- **Male / female figures.** A settings toggle picks the variant. Moving
  exercises show two catalog stills that alternate every second; holds use one.
- **Accounts.** Email + password (bcrypt). A **guest mode** lets people start
  immediately; signing up later upgrades the guest in place, keeping all progress.
- **Streaks & calendar.** Consecutive-day streaks and a monthly activity calendar.
- **7 languages.** Russian, English, Turkish, German, Spanish, French, Italian —
  UI strings, exercise names, and voice cues.
- **Installable PWA** (web manifest + service worker + icons).

---

## Tech stack

| Layer        | Choice |
|--------------|--------|
| Language     | Go (≥ 1.25) |
| UI framework | [doors](https://github.com/doors-dev/doors) + GoX templates (`.gox` → generated `.x.go`) |
| Database     | SQLite via pure-Go [`modernc.org/sqlite`](https://pkg.go.dev/modernc.org/sqlite) (no CGO) |
| Auth         | bcrypt (`golang.org/x/crypto`), cookie sessions |
| Workout player | Vanilla JavaScript + Web Speech API |

---

## Architecture

The app is **server-driven** for everything except the workout itself:

- **Pages** (`segments/pages/`) — sign-in/up, home (today + calendar + streak),
  settings, and the router — are doors components rendered and updated on the
  server. A shared reactive `Session` value drives auth-gated routing.
- **The workout player** (`assets/player.js`) is the one client-side piece. The
  server builds a self-contained JSON payload (the day's items, localized labels
  and voice cues, the user's voice mode and figure variant) and hands it to the
  player, which owns all timing, speech, pausing and rest. `assets/workout-state.js`
  keeps per-account checkpoints and a completion outbox in local storage. A
  program fingerprint and captured cycle position prevent resuming stale state.
  The authenticated completion endpoint is independent of Doors page lifetimes;
  one SQLite transaction records the receipt, activity date and conditional cycle
  advance. Scheduled cycle restarts advance to a new boundary without deleting
  the activity calendar.
- **Layered exercise catalog** (`internal/catalog/`) resolves each exercise
  through three layers, the first match winning:
  `per-user custom → global DB library → built-in code`.
- **Content** (`internal/content/`) holds the exercise library, the public demo,
  and leftover named cycles used only to snapshot existing accounts into a
  personal file. A program is a shared warm-up/cool-down plus a per-day main
  block, flattened and difficulty-scaled at run time.
- **i18n** (`internal/i18n/`) holds all 7-language UI strings, exercise names and
  spoken cues.

---

## Project layout

```
main.go              entry point: open DB, bootstrap session, routing, static files
app.gox              HTML shell, PWA <head>, service-worker registration, root router
path/                URL path model (/, /login, /register, /day/:n, /settings)
internal/
  content/           exercise library + demo 30-day cycle + legacy snapshots
  catalog/           per-user exercise catalog (custom over DB over code)
  i18n/              7-language UI strings, exercise names, voice cues
  db/                SQLite: users, sessions, activity, exercise library, pose stills
  auth/              bcrypt, session tokens, the shared Session value
  app/               process-wide deps (DB) + per-user program/workout resolution
  mcp/               stdio JSON-RPC + optional Bearer HTTP /api
cmd/morning-mcp      stdio MCP server (same SQLite file as the app)
segments/pages/      router, auth, home, workout, settings, shared header
components/           shared UI (e.g. 404)
assets/
  style.css          mobile-first dark theme
  player.js          client-side workout engine (timer, TTS, rest, finish)
  workout-state.js   per-account checkpoints and durable completion retries
  workout-home.js    resume link and pending-result status on the dashboard
  exupload.js        per-exercise JSON import helper
  static/            manifest, service worker, icons, exercise figures
cmd/genicons         PWA icon generator
deploy/              production docker-compose
Dockerfile           multi-stage container build
```

> **Note:** leftover built-in figures live under `assets/static/ex/<m|f>/<id>.png`.
> New stills go into the SQLite pose catalog (`poses` / `pose_images`) and are
> served at `/poses/{id}?sex=m|f`.

---

## Getting started

**Prerequisites**

- Go ≥ 1.25
- The `gox` CLI — a prebuilt binary pinned to the `github.com/doors-dev/gox`
  module version, kept in `.tooling/bin/` (it cannot be `go install`ed).

**Run**

```bash
./dev.sh run          # regenerate templates, then `go run .`  →  http://localhost:8080
```

`dev.sh` sets `PATH` and `GOFLAGS=-mod=mod` for you. Other commands:

```bash
./dev.sh gen          # regenerate .x.go from .gox (run after editing any .gox)
./dev.sh build        # gen + go build ./...
./dev.sh tidy         # go mod tidy
```

> Edit **`.gox`** files only — the `.x.go` files next to them are generated and
> overwritten by `gox gen`.

**Configuration**

| Env                 | Default      | Description |
|---------------------|--------------|-------------|
| `MORNING_DB`        | `morning.db` | SQLite database file path (the container points this at the persistent `/data` volume) |
| `MORNING_MCP_TOKEN` | *(unset)*    | Optional **admin** Bearer token: list/edit any account. Per-user tokens from Settings work without this. |

The server listens on `:8080`.

---

## MCP (agents)

Agents edit **this user's** overlay for exercises and the 30-day cycle, plus a
**shared pose catalog** (stills in SQLite that every account can pick from):

- `custom_exercises` — per-exercise JSON (same document the settings download/upload uses)
- `program_json` — the 30-day cycle (same shape the in-app editor saves)

stdio (local Cursor / Claude against the same DB file):

```bash
MORNING_DB=morning.db go run ./cmd/morning-mcp
```

Tools: `list_accounts`, `get_program`, `set_program`, `restore_demo`,
`list_exercises`, `list_custom_exercises`, `get_exercise`, `put_exercise`,
`list_figures`, `list_poses`, `put_pose`, `set_exercise_poses`.
`account` is a user id, email, or unique display name.
`put_exercise` may set `poses` to 1–2 ids from `list_poses` (`svg` stays the
legacy figure fallback). `put_pose` uploads a still (base64 png/jpeg/webp).
`set_exercise_poses` assigns those stills on a global library exercise.

HTTP is always on `/api/*`. Issue a token in **Settings → MCP token** (shown once;
**Reissue** invalidates the old one). That Bearer is scoped to the account:

```bash
curl -H "Authorization: Bearer m30_…" http://localhost:8080/api/accounts
curl -H "Authorization: Bearer m30_…" \
  -X PUT --data-binary @P08.json \
  http://localhost:8080/api/accounts/me/exercises/P08
curl -H "Authorization: Bearer m30_…" \
  -H "Content-Type: image/png" \
  --data-binary @plank-high.png \
  "http://localhost:8080/api/poses/plank-high?sex=m&kind=plank&label=High%20plank"
```

`MORNING_MCP_TOKEN` is an optional operator override that can still address any
account by id/email/name.

---

## Building & deployment

A multi-stage [`Dockerfile`](Dockerfile) produces a small static binary on a slim
Debian runtime; [`deploy/docker-compose.yml`](deploy/docker-compose.yml) runs it
with a persistent data volume. Production deploys are automated via GitHub
Actions (build → ship → restart); the deployment **secrets, volume and DNS are
managed outside this repository** by the operator.

---

## Internationalization

All user-facing text is centralized in `internal/i18n` as `[7]string` tables in
the order `[ru, en, tr, de, es, fr, it]`. Exercise technique content lives in
`internal/content` per language. Adding a language means adding a column to those
tables and a `Lang` entry — there are no translation files to wire up.

---

## Contributing

Issues and pull requests are welcome — see [CONTRIBUTING.md](CONTRIBUTING.md) for
the build/regenerate workflow and conventions.

## Leap Fitness How-To import

Offline tooling (in `scripts/`, Python) that builds an exercise catalogue from the
public **"How to Do …"** videos on the *Leap Fitness Official* YouTube channel.
No video is ever downloaded — only public titles/URLs are read, and all card text
is written **originally** in this project's house style (nothing is copied from a
video, description or subtitle). Every card carries its source video link and a
`review_status` for manual review before publishing.

```bash
python3 -m venv .venv && .venv/bin/pip install -r requirements.txt

.venv/bin/python scripts/fetch_leap_howto_videos.py       # → data/leap_howto_videos.raw.json
.venv/bin/python scripts/enrich_leap_dates.py             # (optional) fill publish dates
.venv/bin/python scripts/normalize_leap_exercises.py      # dedup (keep newest) + match + group
.venv/bin/python scripts/generate_leap_exercise_cards.py  # assemble house-style cards
.venv/bin/python scripts/validate_leap_exercise_cards.py  # validate
```

Source preference: the **YouTube Data API** when `YOUTUBE_API_KEY` is set
(`.env`), otherwise **yt-dlp** (metadata only). Card text comes from the authored
house-style pass; with `OPENAI_API_KEY` set the generator can fill any gaps via
OpenAI, else it uses safe generic drafts. See [.env.example](.env.example).

Outputs:

| File | What |
|------|------|
| `data/leap_howto_videos.raw.json` | every "How to Do" video found (title, URL, thumbnail) |
| `data/leap_howto_exercises.normalized.json` | unique exercises (freshest duplicate kept), grouped, matched to existing |
| `data/leap_video_map.json` | exercises that match the app's library → just attach a video |
| `content/leap_exercise_cards.en.json` / `.md` | the assembled cards (English) |
| `content/leap_exercise_cards.review.csv` | flat review sheet (`exercise_id,name,video_url,status,issues`) |
| `content/leap_exercise_cards.validation.json` | validation report |

## License

Licensed under the **Apache License, Version 2.0**. See [LICENSE](LICENSE) and
[NOTICE](NOTICE).
