# Ilya: tennis and strength

Personal 30-day bundle: six days repeated five times, independent of weekdays.
Exercises have complete Russian, English, Turkish, German, Spanish, French and
Italian instructions and voice cues. No automatic progression or level scaling.

| Cycle day | Session | Estimated duration including rests |
| --- | --- | --- |
| 1 | Strength A and controlled stops | 26:20 |
| 2 | Footwork and coordination | 20:38 |
| 3 | Recovery | 19:52 |
| 4 | Strength B and controlled landings | 27:08 |
| 5 | Footwork and balance | 20:38 |
| 6 | Same recovery session as day 3 | 19:52 |

Every day has **two complete main rounds**: A → B → C → …, then the same
sequence again. `program.json` stores each main exercise once with `rounds: 2`.
The shared warm-up and cool-down run once every day, including recovery days.
Recovery days keep two lighter main rounds between these shared phases.

`selfManagedSides: true` means the child handles both sides on one exercise
screen. Repetition counts are **per side** where marked; tap Done after both.
Static unilateral holds use one combined timer lasting twice the per-side dose;
switch sides independently halfway through. No occurrence has `sideRest`.
The compatible player must be deployed before installing this program.

The earlier volume is preserved when converting consecutive sets to circuits:
IT07 changes from 3 × 4 to 2 × 6 total repetitions; IT15 and IT31 change from
3 × 20 seconds to 2 × 30 seconds. Single recovery sets are divided between the
two rounds; repeated stretches keep their earlier per-round dose. Push-ups
remain six per round. There is no new progression or maximum-effort work.

### Doses in each main round

Counts followed by “each” apply to each side. Times followed by “each” have one
combined timer; for example, 20 seconds each displays 40 seconds total.

| Day | Main sequence, repeated twice | Pause after each exercise |
| --- | --- | --- |
| 1 | IT07 ×6 total; IT08 ×6 each; IT09 ×6; IT10 ×8; IT11 ×8 each; IT12 ×5 each | 60 s |
| 2 | IT15 30 s; IT16 ×4 full compass sequences; IT17 20 s each; IT18 ×20 total touches; IT19 ×6 each | 45 s |
| 3, 6 | IT01 60 s; IT03 ×4 each; IT20 ×4; IT21 ×3 each; IT22 ×4; IT13 20 s each; IT23 20 s each; IT14 30 s | 15 s |
| 4 | IT24 ×6 total; IT25 ×10; IT09 ×6; IT26 ×10; IT27 ×6 each; IT28 15 s each | 90 s after IT24, otherwise 60 s |
| 5 | IT29 ×6 full crossing-and-return sequences; IT30 ×10 total diagonal steps with return; IT31 30 s; IT32 ×3 three-tap sequences each; IT33 ×6 each | 45 s |

The pause also separates the two main rounds. Every day the same pause
follows the final main exercise before cool-down. No rest follows the final
exercise of the whole session. Durations estimate repetition tempo, include
prescribed rests, and exclude the player's initial three-second ready count.
They are estimates, not targets to rush toward.

### Counted warm-up and coordination work

The shared warm-up is IT01 45 seconds, IT02 ×12 alternating ankle rocks total,
IT04 ×12 movements (4 shoulder circles backward, 4 forward, 4 complete arm
swings), IT05 ×10 hip hinges and IT06 ×10 shallow squats, with 15 seconds after
each exercise. The four counted warm-up exercises previously used 45-second
timers. Main drills IT16, IT18, IT29 and IT30 now use the counts in the table.

Walking, continuous side steps (IT15), continuous rhythm steps (IT31), quiet
breathing and static holds remain timed. All seven instruction languages define
what counts as one repetition and explain independent side changes. The
cool-down remains IT01 60 seconds, IT13 20 seconds per side and IT14 20 seconds,
without added pauses.

New movements need instruction and supervised technique checks; fatigue calls
for the recovery session, with no make-up workload.

The `wrong` cues are spoken directly, without a heading. Author each as a negative
instruction in all seven languages (for example, “Не своди колени и не отрывай
пятки”), keeping every prohibited action in a compound sentence negated.

## Validate and install

```sh
python3 scripts/apply_program_bundle.py programs/ilya-tennis
go test ./programs/ilya-tennis
```

Deploy the compatible player before installing the bundle. On the database host,
use the verified numeric account ID and exact display name:

```sh
python3 scripts/apply_program_bundle.py programs/ilya-tennis \
  --db /path/to/morning.db --account ID --expected-name 'EXACT NAME' \
  --apply --backup /private/backup/account-before.json
```

The installer requires an existing database, makes an exclusive mode-0600 backup
of the account's program fields, merges exercise IDs, and updates only the target
account's program/library/active-program fields in one transaction. Position,
activity, preferences, credentials and other accounts are preserved. It does not
run global seeding or legacy account migration.

Pictures are published separately after the program. `images.json`, when present,
lists bundle-local PNG files and immutable pose IDs; existing different images
cause the transaction to fail. Exercises without pictures display text and voice
until their pose references are installed. Keep account backups outside Git.

For the second publication, add `--images-only` to the install command. This
requires the exercises to exist and only changes their `poses` and supplied
`animation` fields, preserving the live program, exercise text and progress.

## Character and illustrations

All 33 exercises use the same fictional eight-year-old boy from
`references/boy-turnaround.png` and `references/boy-ready.png`. The 31 selected
PNG illustrations live in `images/`; IT02/IT03 share the ankle-rock pose and
IT01/IT31 share the walking pose. These stills remain library thumbnails and
fallbacks if a demonstration cannot load.

`animations/*.png` preserves the original generated strips; `animations/aligned/`
contains the published strips of three equal 640×640 square frames. The
player shows one frame at a time, with a one-second interval. A `movement` strip
plays start–middle–end–middle; an `alternating` strip plays neutral–one side–neutral–
other side. Complex exercises can use `otherPose` for a second three-frame strip:
`movement` finishes and returns on each side; `sequence` walks through all six
frames in order; `alternating` returns to neutral between all four directions
of the compass exercise. This is a technique demonstration, not a rep counter or a pace
to match. The prescribed dose and self-managed sides remain unchanged.

Tapping the figure opens the technique card and pauses the workout clock. The
demonstration can be paused, played or inspected frame by frame. Leaving the app
stops playback; the workout still requires Continue on return. Reduced-motion
preferences disable autoplay, with manual playback available in the card.

Only the current exercise's strips are loaded. Animation URLs are immutable pose
IDs; missing images fall back to the original still. The bundle installer validates
frame count, references and square-cell dimensions before changing the account.

The final frames are cropped and registered algorithmically, without redrawing:

```sh
python3 scripts/align_exercise_frames.py programs/ilya-tennis/animations \
  programs/ilya-tennis/animations/aligned
python3 -m unittest discover -s scripts -p 'test_*.py'
```

This offline authoring tool requires Pillow and numpy. It detects actual gaps
between the source figures (which may not be equal thirds), aligns camera scale
and anchors, then applies a shared crop to every pose of an exercise. It preserves
posture changes instead of enlarging a lower pose to fill its frame. Hinges use
planted feet; jumps retain the landing floor; line steps use the blue line;
stationary-head drills keep the head fixed; floor exercises use the whole visible
mat, including disconnected regions. `aligned/alignment.json` records every
source rectangle, anchor and uniform affine transform. The original images stay
unchanged and allow the process to be rerun exactly.

Images were generated with the built-in ImageGen tool using both character
references. Exact prompts and any corrections are saved in `prompts-*.json`;
reference prompts are in `references/prompts.json`. Only inspected child images
are published. Adult drafts and rejected variants are excluded.
