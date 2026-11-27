# Ilya: progressive strength and movement

This 30-day program varies both the number of exercises and their selection.
The first day has 22 exercise screens; the largest sessions have 32. Recovery
days return to 20 or 22 screens between fuller sessions. The count includes
six warm-up exercises, both complete main rounds, and six cool-down stretches.
The main sequence therefore contains four to ten stations:

`total exercise screens = 6 + 2 × main stations + 6`

A repeated station counts once in each round. Both sides share one screen;
pauses are not exercise screens. Every day includes push-ups and one plank
variation. The complete main sequence runs in order twice. The program uses
existing exercise definitions, translations and illustrations throughout.

M means a moderate session, H a heavier session, and R active recovery. These
are authoring labels in this document, not additional player modes. Recovery
sessions retain the daily routine with fewer stations, reduced push-up/plank
doses and gentle coordination or posture work. They omit the fast reaction
exercise. Heavier sessions are separated by moderate or recovery days. When
hops are included, they are the first main station.

The ordinary difficulty control remains available; all doses below are level 0.
Pauses come only from the account's warm-up/main/cool-down settings. A transition
uses the next exercise's phase setting, including transitions between rounds
and phases; the final exercise has no following pause. Changing those settings
changes the countdown and estimated duration.

`selfManagedSides: true` keeps both sides together. Counted exercises use
repetitions and Done; the user finishes both sides before pressing Done when
`/side` is specified. Timed one-sided holds use twice the per-side dose, with the
existing side-change cue and three-second countdown halfway through. Repetition
tempo fields estimate duration; they do not turn repetitions into timer work.

Installing the revised program preserves the current program position, completed
workouts and streak. It does not restart the cycle: the existing position selects
the corresponding day of the revised table. The cycle remains independent of
weekdays and tennis sessions.

General authoring references include gradual progression, varied exercises,
nonconsecutive strength-training sessions, dynamic preparation and stretching
after exercise in the [AAP report, *Resistance Training for Children and
Adolescents*](https://publications.aap.org/pediatrics/article/145/6/e20201011/76942/Resistance-Training-for-Children-and-Adolescents).
The exact daily counts and doses below are this program's authored schedule.

## Count progression and main doses

| Days | Six consecutive day totals |
| --- | --- |
| 1–6 | 22, 24, 20, 24, 22, 20 |
| 7–12 | 24, 26, 20, 26, 24, 22 |
| 13–18 | 26, 28, 22, 28, 26, 20 |
| 19–24 | 28, 30, 22, 30, 28, 22 |
| 25–30 | 30, 32, 22, 30, 22, 32 |

These targets apply **per main round**, when the exercise appears. Values for
counted exercises are repetitions; continuous movement values are seconds.

| Exercise / session | Days 1–6 | 7–12 | 13–18 | 19–24 | 25–30 |
| --- | ---: | ---: | ---: | ---: | ---: |
| IT09 push-ups, M | 6 | 7 | 8 | 9 | 10 |
| IT09 push-ups, H | 6 | 8 | 10 | 12 | 14 |
| IT09 push-ups, R | 4 | 4 | 5 | 5 | 6 |
| IT41 side plank, R, seconds each | 12 | 12 | 15 | 15 | 15 |
| IT39 shoulder taps, M, total | 8 | 8 | 10 | 10 | 12 |
| IT39 shoulder taps, H, total | 10 | 10 | 12 | 12 | 14 |
| IT40 foot lifts / IT48 side steps, M, total | 6 | 6 | 8 | 8 | 10 |
| IT40 foot lifts / IT48 side steps, H, total | 8 | 8 | 10 | 10 | 12 |
| IT34/35/36/43/49 continuous movement, M, seconds | 35 | 35 | 40 | 40 | 40 |
| IT34/35/36/43/49 continuous movement, H, seconds | 40 | 40 | 45 | 45 | 45 |

IT09 retains the controlled two-second lowering phase. Moving planks count
individual alternating actions: one tap, one foot lift, or one foot out and back.
IT46 counts individual Y, T and W raises, so 9 or 12 repetitions make three or
four complete Y–T–W sequences. Hops remain six per round. The targets retain
controlled form and a small reserve; the existing difficulty control adjusts
main-work doses when needed.

IT36 ends each M and H main round. It uses the existing left/right reaction
signals with `cadence: {"startSeconds": 3, "endSeconds": 1}`. Intervals shorten
across the exercise while its clock continues. Directions mean the child's own
left and right; make short side steps and return to the centre. Visible cues
remain available with voice disabled. Recovery days do not contain IT36.

## All 30 sessions

Each sequence below is performed **twice**. `×N` means counted repetitions;
`Ns` means seconds; `/side` means that dose on each side. Other counts are totals,
including alternating actions. The exercise ID identifies the illustrated,
translated instruction in `exercises.json`.

Times are the application's estimates at level 0 with rests of **15 / 20 / 10
seconds** for warm-up/main/cool-down. They include the actual warm-up override,
both main rounds, all cool-down stretches and timed side changes. Early sessions
and recovery days are deliberately shorter; they are not padded to twenty
minutes. The longest estimate is 30:08. Actual time depends on repetition tempo,
rest preferences and difficulty; the initial three-second ready count is omitted.

| Day | Type | Total screens | Main sequence, twice | Estimated time |
| ---: | :---: | ---: | --- | ---: |
| 1 | M | 22 | IT34 35s; IT09 ×6; IT08 ×5/side; IT39 ×8; IT36 35s | 17:18 |
| 2 | H | 24 | IT35 40s; IT09 ×6; IT38 ×10; IT40 ×8; IT27 ×6/side; IT36 40s | 19:38 |
| 3 | R | 20 | IT18 ×16; IT09 ×4; IT10 ×6; IT41 12s/side | 14:18 |
| 4 | H | 24 | IT07 ×6; IT09 ×6; IT47 ×6/side; IT39 ×10; IT34 40s; IT36 40s | 19:30 |
| 5 | M | 22 | IT49 35s; IT09 ×6; IT44 ×8; IT48 ×6; IT36 35s | 16:50 |
| 6 | R | 20 | IT16 ×3; IT09 ×4; IT20 ×6; IT41 12s/side | 14:22 |
| 7 | M | 24 | IT43 35s; IT09 ×7; IT42 ×5/side; IT40 ×6; IT46 ×9; IT36 35s | 19:02 |
| 8 | H | 26 | IT43 40s; IT09 ×8; IT42 ×6/side; IT48 ×8; IT38 ×10; IT46 ×9; IT36 40s | 21:24 |
| 9 | R | 20 | IT29 ×10; IT09 ×4; IT21 ×4/side; IT41 12s/side | 14:50 |
| 10 | H | 26 | IT24 ×6; IT09 ×8; IT08 ×6/side; IT41 15s/side; IT35 40s; IT26 ×8; IT36 40s | 21:36 |
| 11 | M | 24 | IT34 35s; IT09 ×7; IT45 ×5/side; IT48 ×6; IT10 ×9; IT36 35s | 19:02 |
| 12 | R | 22 | IT32 ×3/side; IT09 ×4; IT22 ×8; IT41 12s/side; IT10 ×6 | 16:02 |
| 13 | M | 26 | IT49 40s; IT09 ×8; IT44 ×10; IT40 ×8; IT26 ×10; IT32 ×3/side; IT36 40s | 21:18 |
| 14 | H | 28 | IT34 45s; IT09 ×10; IT47 ×7/side; IT39 ×12; IT27 ×7/side; IT46 ×12; IT43 45s; IT36 45s | 25:32 |
| 15 | R | 22 | IT18 ×16; IT09 ×5; IT12 ×4/side; IT41 15s/side; IT10 ×6 | 16:20 |
| 16 | H | 28 | IT24 ×6; IT09 ×10; IT38 ×12; IT48 ×10; IT45 ×5/side; IT26 ×10; IT35 45s; IT36 45s | 23:50 |
| 17 | M | 26 | IT43 40s; IT09 ×8; IT42 ×6/side; IT40 ×8; IT27 ×6/side; IT16 ×3; IT36 40s | 22:06 |
| 18 | R | 20 | IT30 ×8; IT09 ×5; IT20 ×6; IT41 15s/side | 14:40 |
| 19 | M | 28 | IT07 ×6; IT09 ×9; IT47 ×6/side; IT48 ×8; IT27 ×6/side; IT46 ×12; IT35 40s; IT36 40s | 24:16 |
| 20 | H | 30 | IT49 45s; IT09 ×12; IT08 ×7/side; IT40 ×10; IT38 ×12; IT46 ×12; IT32 ×3/side; IT34 45s; IT36 45s | 27:08 |
| 21 | R | 22 | IT29 ×10; IT09 ×5; IT21 ×4/side; IT41 15s/side; IT10 ×6 | 16:24 |
| 22 | H | 30 | IT24 ×6; IT09 ×12; IT47 ×7/side; IT41 20s/side; IT27 ×7/side; IT26 ×10; IT43 45s; IT11 ×8/side; IT36 45s | 27:56 |
| 23 | M | 28 | IT49 40s; IT09 ×9; IT38 ×10; IT39 ×10; IT45 ×5/side; IT26 ×10; IT35 40s; IT36 40s | 23:24 |
| 24 | R | 22 | IT16 ×3; IT09 ×5; IT22 ×8; IT41 15s/side; IT18 ×16 | 16:16 |
| 25 | M | 30 | IT34 40s; IT09 ×10; IT08 ×6/side; IT39 ×12; IT45 ×5/side; IT26 ×10; IT43 40s; IT11 ×8/side; IT36 40s | 26:30 |
| 26 | H | 32 | IT07 ×6; IT09 ×14; IT42 ×8/side; IT39 ×14; IT38 ×12; IT46 ×12; IT43 45s; IT27 ×8/side; IT35 45s; IT36 45s | 30:08 |
| 27 | R | 22 | IT32 ×3/side; IT09 ×6; IT20 ×6; IT41 15s/side; IT10 ×6 | 16:26 |
| 28 | H | 30 | IT34 45s; IT09 ×14; IT08 ×8/side; IT48 ×12; IT44 ×12; IT26 ×10; IT49 45s; IT32 ×3/side; IT36 45s | 27:36 |
| 29 | R | 22 | IT30 ×8; IT09 ×6; IT21 ×4/side; IT41 15s/side; IT18 ×16 | 16:38 |
| 30 | H | 32 | IT07 ×6; IT09 ×14; IT47 ×8/side; IT40 ×12; IT46 ×12; IT35 45s; IT38 ×12; IT26 ×10; IT49 45s; IT36 45s | 29:16 |

Duration is computed from each expanded item: seconds use the displayed dose;
repetitions use `value × repSeconds` (default three seconds); per-side work is
doubled, with three seconds added for a timed side change. Each item except the
last then adds the next item's phase rest. This includes per-day warm-ups rather
than assuming every day's preparation has the same duration.

## Six warm-up exercises and six stretches every day

Preparation A is the shared warm-up: short-step jog 40 seconds, ankle rocks
12 total, shoulder circles/arm swings 12 total, reverse lunge 4 each, preparation
turn 6 each, shallow squat 8. Preparation B keeps six exercises and replaces the
lunges/turns with hip hinges 8 and cross-body march 16 total. Recovery preparation
uses B with a 30-second jog. On days 3, 15, 24 and 29, its march is replaced by
preparation turns 4 each so the same march is not repeated in warm-up and main.

| Preparation | Days | Sequence, once |
| --- | --- | --- |
| A | 1, 5, 7, 11, 13, 17, 19, 23, 25 | IT37 40s; IT02 ×12; IT04 ×12; IT08 ×4/side; IT19 ×6/side; IT06 ×8 |
| B | 2, 4, 8, 10, 14, 16, 20, 22, 26, 28, 30 | IT37 40s; IT02 ×12; IT04 ×12; IT05 ×8; IT18 ×16; IT06 ×8 |
| Recovery B with turns | 3, 15, 24, 29 | IT37 30s; IT02 ×12; IT04 ×12; IT05 ×8; IT19 ×4/side; IT06 ×8 |
| Recovery B | 6, 9, 12, 18, 21, 27 | IT37 30s; IT02 ×12; IT04 ×12; IT05 ×8; IT18 ×16; IT06 ×8 |

The cool-down always contains the following six stretches, once after the second
main round. Its pauses use the account's cool-down setting.

| Exercise | Dose |
| --- | --- |
| IT13 calf stretch | 20 seconds each side |
| IT51 standing quadriceps stretch by a wall | 20 seconds each side |
| IT50 half-kneeling hip flexor stretch | 20 seconds each side |
| IT52 lying figure-four glute stretch | 20 seconds each side |
| IT54 chest stretch with the forearm on a wall | 15 seconds each side |
| IT53 child's pose for the back and lats | 30 seconds |

All 54 definitions remain in the library. Instructions and voice cues cover
Russian, English, Turkish, German, Spanish, French and Italian. Mistake cues
continue to be spoken directly as negative instructions, without an introductory
heading. Keep every prohibited action in a compound sentence negated.

## Validate and install

```sh
python3 scripts/apply_program_bundle.py programs/ilya-tennis
go test ./programs/ilya-tennis
```

This update uses the existing player. On the database host, use the verified
numeric account ID and exact display name:

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

Keep account backups outside Git.

## Character and illustrations

The app resolves illustrations for this program by exercise ID from the shared
library. Each exercise has a man, woman, boy and girl. Set the figure in the
account settings. The [boy's character sheet](../character-references/turnarounds/boy.png)
shows the approved front, back and three-quarter views.

The player shows one frame for holds and uses two or three phases for movement.
The technique card pauses the workout clock and lets the user play or step
through the phases. During active reaction work, the figure stays neutral while
the signal changes. The [exercise artwork guide](../library-illustrations/README.md)
describes where the images and phase definitions are stored.
