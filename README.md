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
- **Ready programs.** A catalog of 50 ready 30-day programs in seven sections —
  first steps (morning workout, 50+), strength (plank, abs, push-ups, arms,
  pull-ups, legs), weight loss, stretching (seven programs, from splits to a
  bedtime stretch), back and balance (posture, desk workers, tai chi, back care
  for disc protrusions, back symmetry for scoliosis, healthy feet), kids
  (thirteen programs: general, agility, posture, tennis, football, swimming,
  wrestling and judo, karate, gymnastics, dance, hockey and skating,
  basketball and volleyball, flat feet) and sports (runners, football, tennis,
  skiing, martial arts, basketball and volleyball, swimming, wrestling and
  judo, cycling, hockey, climbing, golf, padel and badminton, hiking, rowing) —
  with one-line descriptions, a preview of the 30-day plan, and "no jumps" /
  "no equipment" filters read from the programs themselves. Picking a kids'
  program switches the figure to a child. Each course also has a public page per language
  (`/{lang}/courses/...`, listed in the sitemap).
- **About you and the fitness index.** A newcomer (guest or new account)
  first sees "About you": for me or for a child, figure, age, height, weight
  and what bothers them (back, knees, blood pressure, feet). A voice-guided
  self-test follows — push-ups, squats in one minute, a plank hold and
  pull-ups (each can be skipped). `internal/fitness` compares each result with
  a typical one for the age and sex (typical = 50 points, twice that = 100,
  children have their own norms), averages them into an index 0–100 with a
  level (getting started, fair, good, excellent) and recommends a program,
  its difficulty and rounds: back → back care (posture for children), feet →
  healthy feet, blood pressure or 50+ → workout 50+, otherwise the morning
  workout; knees, a high BMI, blood pressure and age 60+ make it gentler. The
  result screen shows the index (and the previous one after a retest) and
  starts the program. Settings → Profile → Fitness reopens it; the answers and
  tests are deleted with the account.
- **Selected programs.** The user's own shortlist. One tap on a program
  switches to it; the star adds or removes it without switching. Every
  selected program keeps its own copy (with the user's edits) and its own
  cycle day, so switching never loses anything; the streak and calendar are
  common. Removing a program that holds the user's edits asks first. The home
  screen offers quick switching between them (popular ready programs until
  there is a second one).
- **Day screen and practical library.** Tapping the day on the home screen
  lists its exercises; each one opens its card, a replacement picker (easier,
  similar, harder exercises of the same group, for today, every day or chosen
  days) or removal. The library has search, groups and a "no jumps" filter;
  "+" adds an exercise to a part of the current program on chosen days with a
  chosen amount. Edits always go to the user's version of the program.
- **Program page.** Tapping the program name on the home screen opens its
  page: rename it, and set its difficulty, number of main rounds (1–10) and the
  three pauses. These belong to each program and travel with it when switching.
  From there: today's exercises, the day-by-day editor (reorder, add, remove and
  retime exercises; all three phases stay nonempty), the library and the JSON
  download. General settings keep only what belongs to the person — voice,
  figure, language, name, account. Every setting is saved as soon as it changes.
- **AI coach.** On the program page, "AI coach" opens a chat that edits the
  current program — and only it: add, replace and remove exercises on chosen
  days, change doses, rounds, difficulty and rests, rename the program, and
  create new exercises in the person's library. Each request carries the
  program day by day, the program rules and the "About you" answers, so nobody
  has to explain where they are. The list of changes under a reply is computed
  by the server from the program before and after, and **Undo** brings the
  program back. Up to 40 messages per account per day.
- **Exercise documents.** Besides the original built-ins, exercises can be
  complete JSON documents (definition plus name, technique, mistakes,
  breathing, safety and voice cues in every language) in
  `internal/content/exercises/*.json`; they join the built-in library at
  start-up. A test checks every document in every language.
- **Exercise JSON.** Download any exercise as a file, paste or upload one back.
  Custom exercises overlay the built-in library. Each exercise can use one
  catalog still (static) or two (they swap once a second). Legacy built-in
  PNGs remain the fallback until photos are assigned.
- **Local MCP for the operator.** `cmd/morning-mcp` (stdio, on the machine with
  the database file) edits any account's exercises and program and uploads pose
  stills. There is no network API and no per-user tokens.
- **Universal difficulty.** A per-user level from **−3 to +3** scales every
  program by ±10% per step (reps and seconds only; breaths and warm-up/cool-down
  stay fixed). Prescribed programs can lock their doses with `fixedDifficulty`.
- **One rest policy.** Account settings control every pause between exercises:
  warm-up and cool-down 0–40 seconds, main 5–40. A pause belongs to the upcoming
  exercise's phase; the final exercise has no following pause. The same resolved
  values drive estimates and player timers. Program JSON cannot prescribe pauses;
  old `rest`/`sideRest` fields are ignored on import and omitted on export.
- **Daily phases.** Every program includes a warm-up, its main rounds (two by default) and a
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
  has saved it; the home page also retries pending results. Finishing the final
  timer while leaving the app records that finish immediately, before rendering
  the celebration. Completed checkpoints are recovered into the pending queue
  even if the program or cycle position has since changed. Each newly finished
  attempt gets its own completion ID, so two pages resumed from one checkpoint
  cannot mistake an earlier receipt for the new result; retries keep that ID.
- **Spoken guidance** via the browser Web Speech API: a start cue, a rest
  announcement of the next exercise with its count, a 3-2-1 lead-in, a half-time
  marker on timed sets, and running technique narration (how to do it / common
  mistakes). Four levels in settings: off · minimal · normal · detailed.
- **Four figures.** Man, woman, boy or girl, chosen in settings and in each
  exercise card. A child figure without its own picture is shown as the adult
  of the same sex. Picking a children's ready program switches to the child
  figure (and back for an adult program). Moving exercises show two catalog
  stills that alternate every second; holds use one.
- **Accounts.** Email + password (bcrypt). A **guest mode** lets people start
  immediately; signing up later upgrades the guest in place, keeping all progress.
- **Streaks & calendar.** Consecutive-day streaks and a monthly activity calendar.
- **28 languages.** Russian, English, Turkish, German, Spanish, French, Italian,
  Portuguese (Brazil), Polish, Dutch, Swedish, Danish, Norwegian, Finnish, Czech,
  Romanian, Hungarian, Greek, Japanese, Korean, Chinese (Simplified and
  Traditional), Indonesian, Vietnamese, Thai, Arabic, Hebrew and Persian — the
  interface, voice cues, every exercise card and every program. Arabic, Hebrew
  and Persian lay the screens out right to left. The language follows the
  browser (zh-TW/zh-HK open Traditional Chinese) and can be changed in settings.
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
- **i18n** (`internal/i18n/`) holds the UI strings, exercise names and spoken
  cues, one file per language in `locales/`.

---

## Project layout

```
main.go              entry point: open DB, bootstrap session, routing, static files
app.gox              HTML shell, PWA <head>, service-worker registration, root router
path/                URL path model (/, /login, /register, /day/:n, /settings)
internal/
  content/           exercise library + demo 30-day cycle + legacy snapshots
  catalog/           per-user exercise catalog (custom over DB over code)
  i18n/              UI strings, exercise names, voice cues (locales/<lang>.json)
  db/                SQLite: users, sessions, activity, exercise library, pose stills
  auth/              bcrypt, session tokens, the shared Session value
  app/               process-wide deps (DB) + per-user program/workout resolution
  mcp/               stdio JSON-RPC tools for cmd/morning-mcp
  coach/             the AI coach: model client, scoped tools, context, change list
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
| `AI_API_KEY`        | *(unset)*    | Key of an OpenAI-compatible chat API (OpenRouter). Without it the AI coach says it is unavailable. Keep it in the server's environment or a local `.env` (git-ignored), never in the repository. |
| `AI_BASE_URL`       | `https://openrouter.ai/api/v1` | The chat API |
| `AI_MODEL`          | `~deepseek/deepseek-flash-latest@deepseek` | `model@provider`. `~…-latest` follows the maker's newest model of that family by itself; `@provider` pins requests to the maker's own servers (no resellers). |
| `AI_FALLBACK`       | `qwen/qwen3.8-flash@alibaba` | Used when the main model fails; `none` turns it off |

`go run ./cmd/coachbench model@provider …` runs real program edits against
models and prints which passed, the cost per message and the time — use it
before changing `AI_MODEL`.

The server listens on `:8080`.

---

## MCP (operator, local only)

`cmd/morning-mcp` is a stdio MCP server for the operator's own tools (Cursor,
Claude Code) on the machine that holds the database file. It is not reachable
over the network; inside the app, people edit programs with the AI coach.

```bash
MORNING_DB=morning.db go run ./cmd/morning-mcp
```

Tools: `list_accounts`, `list_programs`, `get_ready_program`, `list_selected`,
`select_program`, `star_program`, `unselect_program`, `rename_program`,
`get_program`, `set_program`, `import_program`, `restore_demo`, `get_day`,
`add_exercise`, `replace_exercise`, `remove_exercise`, `list_exercises`,
`similar_exercises`, `get_settings`, `set_settings`, `list_custom_exercises`,
`get_exercise`, `put_exercise`, `list_figures`, `list_poses`, `put_pose`,
`set_exercise_poses`. `account` is a user id, email or unique display name;
`days` is `"today"`, `"all"` (default), a day number or a list. Nothing changes
the streak; `put_pose` uploads a still (base64 png/jpeg/webp) for a figure
`m`, `f`, `b` or `g`.

---

## Building & deployment

A multi-stage [`Dockerfile`](Dockerfile) produces a small static binary on a slim
Debian runtime; [`deploy/docker-compose.yml`](deploy/docker-compose.yml) runs it
with a persistent data volume. Production deploys are automated via GitHub
Actions (build → ship → restart); the deployment **secrets, volume and DNS are
managed outside this repository** by the operator.

---

## Internationalization

The list of languages lives in `internal/langs` (code, name, speech tag,
right-to-left). Each language has:

- `internal/i18n/locales/<lang>.json` — UI strings (with `%d` templates such as
  `fmt.day` so every language puts numbers where it needs them, and plural
  variants such as `fmt.exercises|few` picked by CLDR rules), exercise names,
  month and weekday names;
- `internal/content/i18n/{courses,details,exercises}.<lang>.json` — program
  titles and descriptions, the cards of the original exercises, and the cards
  of the bundled exercises (the first seven languages keep theirs in the Go
  tables and exercise documents).

Anything missing falls back to English, and tests check that every language is
complete: every UI key with the same placeholders, every exercise card with the
same number of steps, and mistakes phrased as prohibitions in that language.
Adding a language means adding it to `internal/langs` and adding those files.

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
