# Translation guide — exercise cards into 6 languages

You translate English exercise cards into **Russian, Turkish, German, Spanish,
French, Italian**. The English source is already in the project's house style;
keep that style in every language.

## Input

A JSON array of cards, each:
`{exercise_id, exercise_name, short_description, how_to[], correct[], mistakes[]}`

## Output

A JSON array, one object per input card, SAME ORDER, shape:

```json
{
  "exercise_id": "<copy exactly>",
  "ru": {"name": "...", "short_description": "...", "how_to": [...], "correct": [...], "mistakes": [...]},
  "tr": { ... },
  "de": { ... },
  "es": { ... },
  "fr": { ... },
  "it": { ... }
}
```

Write RAW JSON only (no markdown fences, no prose).

## Rules

- Translate ALL fields: `name` (from exercise_name), `short_description`, and
  every line of `how_to`, `correct`, `mistakes`.
- Keep the SAME number of lines per array as the English source.
- Natural, idiomatic fitness language — translate the meaning, not word-for-word.
- Keep the house voice per language:
  - `how_to`: short imperative steps (the language's natural command form).
  - `correct`: short cues, second person ("Keep your…" → the natural equivalent,
    e.g. RU "Держи…/Держите…" — use the informal **ты** form to match the app).
  - `mistakes`: short negative cues — the English starts each with "Don't …";
    render the natural negative imperative (RU "Не …", DE "Nicht …", ES "No …",
    FR "Ne …", IT "Non …", TR "… -ma/-me").
- Exercise names: use the established term in each language; if a movement has no
  common local name (many yoga poses), keep a recognizable form (transliterate or
  keep the English/Sanskrit name) rather than inventing one.
- No added hype, no emojis, no brand names, no medical claims.
- Use the informal singular ("ты"/"du"/"tú"/"tu") to match the app's tone.

## Example (one card → output object)

Input card:
`{"exercise_id":"leap_squat_jacks","exercise_name":"Squat Jacks","short_description":"A cardio squat that mixes a jumping jack with a squat to work the legs, glutes and heart.","how_to":["Stand with your feet together and your hands in front.","Jump your feet out and drop into a squat.","Jump your feet back together as you stand up.","Repeat with a steady rhythm."],"correct":["Land softly.","Keep your knees over your toes.","Keep your chest up.","Keep your heels down in the squat."],"mistakes":["Don't land hard.","Don't let your knees cave inward.","Don't round your back.","Don't rush and lose control."]}`

Output object:
```json
{
  "exercise_id": "leap_squat_jacks",
  "ru": {"name": "Сквот-джеки", "short_description": "Кардио-присед, сочетающий прыжковый джек с приседом — работают ноги, ягодицы и сердце.", "how_to": ["Встань, ноги вместе, руки перед собой.", "Прыжком расставь ноги и опустись в присед.", "Прыжком верни ноги вместе и выпрямись.", "Повторяй в ровном ритме."], "correct": ["Приземляйся мягко.", "Держи колени над носками.", "Держи грудь раскрытой.", "Держи пятки на полу в приседе."], "mistakes": ["Не приземляйся жёстко.", "Не своди колени внутрь.", "Не округляй спину.", "Не торопись и не теряй контроль."]},
  "tr": {"name": "Squat Jack", "short_description": "Jumping jack ile squat'ı birleştiren, bacakları, kalçaları ve kalbi çalıştıran kardiyo squat.", "how_to": ["Ayakların bitişik, ellerin önde dur.", "Sıçrayarak ayaklarını aç ve squat'a in.", "Sıçrayarak ayaklarını birleştir ve doğrul.", "Sabit bir ritimde tekrarla."], "correct": ["Yumuşak in.", "Dizlerini ayak parmaklarının üstünde tut.", "Göğsünü dik tut.", "Squat'ta topuklarını yerde tut."], "mistakes": ["Sert inme.", "Dizlerini içe bükme.", "Sırtını yuvarlama.", "Acele edip kontrolü kaybetme."]},
  "de": {"name": "Squat Jacks", "short_description": "Eine Cardio-Kniebeuge, die Jumping Jack und Kniebeuge verbindet und Beine, Gesäß und Herz fordert.", "how_to": ["Stell dich hin, Füße zusammen, Hände vorn.", "Spring die Füße auseinander und geh in die Hocke.", "Spring die Füße zusammen und richte dich auf.", "Wiederhole in gleichmäßigem Rhythmus."], "correct": ["Lande weich.", "Halte die Knie über den Zehen.", "Halte die Brust aufrecht.", "Halte die Fersen in der Hocke am Boden."], "mistakes": ["Lande nicht hart.", "Lass die Knie nicht nach innen kippen.", "Mach keinen runden Rücken.", "Überhaste dich nicht und verlier nicht die Kontrolle."]},
  "es": {"name": "Squat Jacks", "short_description": "Una sentadilla cardio que combina el jumping jack con la sentadilla para trabajar piernas, glúteos y corazón.", "how_to": ["Ponte de pie con los pies juntos y las manos delante.", "Salta abriendo los pies y baja a una sentadilla.", "Salta juntando los pies y ponte de pie.", "Repite con un ritmo constante."], "correct": ["Cae suave.", "Mantén las rodillas sobre los dedos.", "Mantén el pecho arriba.", "Mantén los talones en el suelo en la sentadilla."], "mistakes": ["No caigas con fuerza.", "No dejes que las rodillas se hundan hacia dentro.", "No redondees la espalda.", "No te apresures ni pierdas el control."]},
  "fr": {"name": "Squat Jacks", "short_description": "Un squat cardio qui mêle le jumping jack au squat pour travailler les jambes, les fessiers et le cœur.", "how_to": ["Tiens-toi pieds joints, mains devant.", "Saute en écartant les pieds et descends en squat.", "Saute en rejoignant les pieds et redresse-toi.", "Répète à un rythme régulier."], "correct": ["Atterris en douceur.", "Garde les genoux au-dessus des orteils.", "Garde la poitrine ouverte.", "Garde les talons au sol dans le squat."], "mistakes": ["N'atterris pas brutalement.", "Ne laisse pas les genoux rentrer vers l'intérieur.", "N'arrondis pas le dos.", "Ne te précipite pas au point de perdre le contrôle."]},
  "it": {"name": "Squat Jacks", "short_description": "Uno squat cardio che unisce il jumping jack allo squat per lavorare gambe, glutei e cuore.", "how_to": ["Mettiti in piedi con i piedi uniti e le mani davanti.", "Salta allargando i piedi e scendi in squat.", "Salta riunendo i piedi e raddrizzati.", "Ripeti con un ritmo costante."], "correct": ["Atterra in modo morbido.", "Tieni le ginocchia sopra le punte dei piedi.", "Tieni il petto alto.", "Tieni i talloni a terra nello squat."], "mistakes": ["Non atterrare con forza.", "Non far cedere le ginocchia verso l'interno.", "Non incurvare la schiena.", "Non avere fretta perdendo il controllo."]}
}
```
