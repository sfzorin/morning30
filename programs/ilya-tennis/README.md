# Ilya: tennis and strength

Personal 30-day bundle: six days repeated five times, independent of weekdays.
Exercises have complete Russian, English, Turkish, German, Spanish, French and
Italian instructions and voice cues. No automatic progression or level scaling.

| Cycle day | Session | Estimated duration including rests |
| --- | --- | --- |
| 1 | Strength A and controlled stops | 27:24 |
| 2 | Footwork and coordination | 18:52 |
| 3 | Recovery | 12:11 |
| 4 | Strength B and controlled landings | 26:52 |
| 5 | Footwork and balance | 18:52 |
| 6 | Same recovery session as day 3 | 12:11 |

`program.json` stores sets consecutively, with `rounds: 1`. Recovery days disable
shared warm-up/cool-down through explicit empty arrays. Per-side exercises have
a 10-second switch; the short cool-down calf stretch switches immediately.
The initial push-up dose is six. Repetition tempo is an estimate, not a timer.
New movements need instruction and supervised technique checks; fatigue calls
for the recovery session, with no make-up workload.

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
requires the exercises to exist and only changes their `poses` fields, preserving
the live program and all exercise text as well as progress.

## Character and illustrations

All 33 exercises use the same fictional eight-year-old boy from
`references/boy-turnaround.png` and `references/boy-ready.png`. The 31 selected
PNG illustrations live in `images/`; IT02/IT03 share the ankle-rock pose and
IT01/IT31 share the walking pose. Each shows a key position; sequence, timing and
side changes are explained by the localized instructions and player.

Images were generated with the built-in ImageGen tool using both character
references. Exact prompts and any corrections are saved in `prompts-*.json`;
reference prompts are in `references/prompts.json`. Only inspected child images
are published. Adult drafts and rejected variants are excluded.
