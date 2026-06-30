# Exercise card generation — house style guide

You write short, original English instruction cards for bodyweight/fitness
exercises. The text must be ORIGINAL (never copied from any video, subtitle or
description). You only get an exercise name + its category; infer safe, sensible
instructions from the name. These are drafts for human review.

## Output schema (per exercise)

```json
{
  "exercise_id": "<copy EXACTLY from input>",
  "short_description": "one short sentence",
  "how_to": ["line", "line", "line", "line"],
  "correct": ["line", "line", "line", "line"],
  "mistakes": ["line", "line", "line", "line"]
}
```

- `short_description`: exactly ONE short sentence naming the movement and what it
  works (muscles / purpose). No hype.
- `how_to`: 3–6 short imperative steps ("Stand tall…", "Lower your hips…").
- `correct`: 3–5 short cues, written in the SECOND PERSON, most starting with
  "Keep your…" / "Relax your…" / a short imperative ("Land softly.").
- `mistakes`: 3–5 short lines, EACH starting with "Don't …".
- Every line is short (≈3–9 words). No paragraphs. No empty strings.

## Voice rules (strict)

- English only. Second person ("you/your"). NEVER first person ("I/we").
- No motivational hype. No emojis.
- Never mention Leap Fitness, YouTube, or any brand/channel.
- No medical or results claims.
- BANNED phrases (must never appear, any case): "no pain no gain",
  "burn fat fast", "guaranteed", "cure", "medical treatment",
  "six pack in 30 days".

## Per-type safety points — fold these in, phrased in the house voice

If the name involves **jump / jumping / hops / burpee / high knees / jacks / skater**:
- correct must include: "Land softly." and "Keep your knees over your toes."
- mistakes must include: "Don't land hard." and "Don't let your knees cave inward."

If the name involves **plank**:
- correct must include: "Keep your body in one line." and "Keep your core engaged."
- mistakes must include: "Don't let your lower back sag." and "Don't lift your hips too high."

If the name involves **push-up / push up / pushup / press-up / dip**:
- correct must include: "Keep your body straight." and "Keep your elbows controlled."
- mistakes must include: "Don't let your hips drop." and "Don't flare your elbows out wide."

If the name involves **crunch / kicks / scissors / leg raise / leg lift / bicycle / sit-up**:
- correct must include: "Keep your lower back controlled."
- mistakes must include: "Don't arch your lower back." and "Don't pull on your neck."

If the name involves **cobra / superman / snow angel / swimmers / prone / Y/T/W raise / locust**:
- correct must include: "Keep your neck neutral." and "Move with control."
- mistakes must include: "Don't throw your head back." and "Don't over-arch your lower back."

If the name involves **squat / lunge / split squat**:
- correct must include: "Keep your knees over your toes." (and "Keep your heels down." where it applies)
- mistakes must include: "Don't let your knees cave inward." and "Don't drop too fast."

(For yoga poses, stretches, dumbbell/equipment, face/neck, wall, etc. — no
mandatory lines; just follow the voice rules and write sensible cues. For a hold
or stretch, "Breathe steadily." / "Don't hold your breath." / "Don't force the
stretch." fit well.)

## Approved examples (match this exactly)

```json
[
  {
    "exercise_id": "leap_push_up_hold",
    "short_description": "An isometric push-up that builds chest, shoulder and core strength by holding the lowered position.",
    "how_to": ["Start in a push-up position with your hands under your shoulders.", "Lower halfway down until your elbows are bent.", "Hold there with your body in a straight line.", "Breathe steadily and keep your core tight."],
    "correct": ["Keep your body straight from head to heels.", "Keep your core engaged.", "Keep your elbows close to your body.", "Breathe steadily."],
    "mistakes": ["Don't let your hips drop.", "Don't raise your hips too high.", "Don't flare your elbows out wide.", "Don't hold your breath."]
  },
  {
    "exercise_id": "leap_squat_jacks",
    "short_description": "A cardio squat that mixes a jumping jack with a squat to work the legs, glutes and heart.",
    "how_to": ["Stand with your feet together and your hands in front.", "Jump your feet out and drop into a squat.", "Jump your feet back together as you stand up.", "Repeat with a steady rhythm."],
    "correct": ["Land softly.", "Keep your knees over your toes.", "Keep your chest up.", "Keep your heels down in the squat."],
    "mistakes": ["Don't land hard.", "Don't let your knees cave inward.", "Don't round your back.", "Don't rush and lose control."]
  },
  {
    "exercise_id": "leap_wide_leg_hamstring_stretch",
    "short_description": "A wide-leg forward fold that lengthens the hamstrings and inner thighs.",
    "how_to": ["Set your feet wide apart.", "Keep your back straight and hinge from your hips.", "Reach your hands toward the floor.", "Hold where you feel a gentle stretch."],
    "correct": ["Keep your back straight.", "Hinge from your hips.", "Keep a soft bend in your knees.", "Breathe slowly."],
    "mistakes": ["Don't round your back.", "Don't lock your knees.", "Don't bounce.", "Don't force the stretch."]
  }
]
```
