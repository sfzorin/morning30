package content

// The "… in 30 days" series: goal-named 30-day courses built from the same
// exercise library as the demo. They live in their own registry on purpose.
// StandardPrograms (what accounts can be on) and LegacyNamed (old snapshots)
// do not list them, so no existing account, program or cycle position is
// affected until a course is explicitly offered somewhere.

// Course is one goal-named 30-day program of the series.
type Course struct {
	Key      string            // stable id, e.g. "plank30"
	Slug     string            // public URL name, e.g. "plank-in-30-days"
	Category string            // catalog section ("" = from courseCategory)
	Kids     bool              // a children's program
	Title    map[string]string // language code → display title
	About    map[string]string // language code → one-paragraph description
	plan     coursePlan
}

// TitleIn returns the title in language l, falling back to English.
func (c Course) TitleIn(l string) string {
	if t := c.Title[l]; t != "" {
		return t
	}
	return c.Title["en"]
}

// AboutIn returns the description in language l, falling back to English.
func (c Course) AboutIn(l string) string {
	if t := c.About[l]; t != "" {
		return t
	}
	return c.About["en"]
}

// IsRecoveryDay reports whether a day of the series is a lighter session.
func IsRecoveryDay(day int) bool { return isRecoveryDay(day) }

// Resolve builds the course as an ordinary editable program: the same shape
// the demo, the editor, JSON import and the player already use.
func (c Course) Resolve() Resolved {
	days := make(map[int][]setDef, TotalDays)
	for day := 1; day <= TotalDays; day++ {
		days[day] = c.plan.block(day)
	}
	return programSpec{name: c.Title["en"], warmup: c.plan.warmup, cooldown: c.plan.cooldown, days: days}.resolve()
}

// Series30 returns the courses in display order.
func Series30() []Course {
	out := make([]Course, len(series30))
	copy(out, series30)
	return out
}

// CourseByKey returns a course of the series by its key.
func CourseByKey(key string) (Course, bool) {
	for _, c := range series30 {
		if c.Key == key {
			return c, true
		}
	}
	return Course{}, false
}

// CourseBySlug returns a course of the series by its public URL name.
func CourseBySlug(slug string) (Course, bool) {
	for _, c := range series30 {
		if c.Slug == slug {
			return c, true
		}
	}
	return Course{}, false
}

// dose is one main station whose level-0 dose per round grows linearly from
// From (its first day) to To (day 29, the peak). Since delays the station's
// first appearance, which is how sessions gain stations over the month.
type dose struct {
	ID       string
	From, To int
	Since    int
}

// coursePlan is a course in compact form. Training days rotate through
// sessions a → b → c; the demo's recovery days (4, 7, 11, 14, 18, 21, 25, 28)
// run the lighter recovery session, which is shortest in the first week;
// day 30 is a fixed final session.
type coursePlan struct {
	warmup, cooldown []setDef
	a, b, c          []dose
	recovery         []dose
	final            []setDef
}

const peakDay = TotalDays - 1

// block returns the main block (one round) for a day.
func (p coursePlan) block(day int) []setDef {
	day = clampDay(day)
	if day == TotalDays {
		return p.final
	}
	if isRecoveryDay(day) {
		return dosesOn(p.recovery, day)
	}
	training := 0 // zero-based index among the training days so far
	for d := 1; d < day; d++ {
		if !isRecoveryDay(d) {
			training++
		}
	}
	return dosesOn([][]dose{p.a, p.b, p.c}[training%3], day)
}

// dosesOn interpolates a session's stations for a day. Timed work is rounded
// to five seconds so on-screen and spoken counts stay simple.
func dosesOn(session []dose, day int) []setDef {
	out := make([]setDef, 0, len(session))
	for _, s := range session {
		start := s.Since
		if start < 1 {
			start = 1
		}
		if day < start {
			continue
		}
		val := s.From
		if span := peakDay - start; span > 0 {
			val = s.From + ((s.To-s.From)*(day-start)*2+span)/(span*2) // round half up
		}
		if Pool[s.ID].Unit == Seconds {
			val = (val + 2) / 5 * 5
		}
		out = append(out, setDef{s.ID, val})
	}
	return out
}

var series30 = []Course{
	{
		Key:  "plank30",
		Slug: "plank-in-30-days",
		About: map[string]string{
			"ru": "Планка каждый день: от 20 секунд до 70 за подход. Классическая, боковая, динамическая и обратная планки чередуются, а упражнения для спины держат баланс.",
			"en": "A plank every day, from 20 seconds up to 70 per set. Front, side, moving and reverse planks rotate, with back work to keep things balanced.",
			"tr": "Her gün plank: set başına 20 saniyeden 70 saniyeye. Klasik, yan, hareketli ve ters plank dönüşümlü gelir; sırt egzersizleri dengeyi korur.",
			"de": "Jeden Tag Plank: von 20 bis 70 Sekunden pro Satz. Klassischer, seitlicher, dynamischer und umgekehrter Plank wechseln sich ab, Rückenübungen sorgen für Ausgleich.",
			"es": "Plancha cada día: de 20 a 70 segundos por serie. Se alternan la plancha frontal, lateral, dinámica e invertida, con ejercicios de espalda para equilibrar.",
			"fr": "Du gainage chaque jour : de 20 à 70 secondes par série. Planche classique, latérale, dynamique et inversée alternent, avec du travail du dos pour l'équilibre.",
			"it": "Plank ogni giorno: da 20 a 70 secondi per serie. Plank frontale, laterale, dinamico e inverso si alternano, con esercizi per la schiena per bilanciare.",
		},
		Title: map[string]string{
			"ru": "Планка за 30 дней", "en": "Plank in 30 Days", "tr": "30 Günde Plank",
			"de": "Plank in 30 Tagen", "es": "Plancha en 30 días", "fr": "Gainage en 30 jours",
			"it": "Plank in 30 giorni",
		},
		// Front plank grows 20 s → 70 s per round; side, dynamic and reverse
		// planks rotate around it, with back extension as the counterweight.
		plan: coursePlan{
			warmup:   []setDef{{"W01", 15}, {"W09", 20}, {"W03", 6}, {"W04", 8}, {"W05", 8}, {"W06", 20}},
			cooldown: []setDef{{"CD04", 25}, {"CD07", 25}, {"CD03", 25}, {"CD01", 25}},
			a: []dose{
				{ID: "C01", From: 20, To: 60}, {ID: "C09", From: 6, To: 14}, {ID: "B01", From: 20, To: 35},
				{ID: "C07", From: 15, To: 30}, {ID: "C02", From: 10, To: 25, Since: 13},
			},
			b: []dose{
				{ID: "C07", From: 15, To: 35}, {ID: "C10", From: 4, To: 10}, {ID: "B08", From: 15, To: 35},
				{ID: "C08", From: 5, To: 12}, {ID: "C01", From: 20, To: 55},
			},
			c: []dose{
				{ID: "C01", From: 25, To: 70}, {ID: "C06", From: 15, To: 30}, {ID: "C09", From: 6, To: 16},
				{ID: "B06", From: 8, To: 14}, {ID: "L09", From: 4, To: 10, Since: 9},
			},
			recovery: []dose{
				{ID: "C01", From: 20, To: 40}, {ID: "C12", From: 6, To: 10}, {ID: "B01", From: 20, To: 30},
				{ID: "L06", From: 8, To: 12, Since: 9},
			},
			final: []setDef{{"C02", 25}, {"C07", 35}, {"C10", 10}, {"B08", 35}, {"C09", 16}, {"C01", 90}},
		},
	},
	{
		Key:  "abs30",
		Slug: "abs-in-30-days",
		About: map[string]string{
			"ru": "Пресс дома без инвентаря: нижняя часть, косые мышцы и весь корпус по очереди. Скручивания, ножницы, велосипед и планки с плавным ростом нагрузки.",
			"en": "Abs at home with no equipment: lower abs, obliques and the whole core in turn. Crunches, scissors, bicycles and planks with a steady build-up.",
			"tr": "Evde ekipmansız karın çalışması: alt karın, yan karın ve tüm gövde sırayla. Mekik, makas, bisiklet ve plank; yük kademeli artar.",
			"de": "Bauchtraining zu Hause ohne Geräte: unterer Bauch, schräge Bauchmuskeln und der ganze Rumpf im Wechsel. Crunches, Scheren, Radfahren und Planks mit stetiger Steigerung.",
			"es": "Abdominales en casa sin material: zona baja, oblicuos y todo el core por turnos. Crunches, tijeras, bicicleta y planchas con progresión gradual.",
			"fr": "Des abdos à la maison sans matériel : bas du ventre, obliques et tout le tronc à tour de rôle. Crunchs, ciseaux, vélo et planches avec une progression régulière.",
			"it": "Addominali a casa senza attrezzi: parte bassa, obliqui e tutto il core a turno. Crunch, forbici, bicicletta e plank con una progressione graduale.",
		},
		Title: map[string]string{
			"ru": "Пресс за 30 дней", "en": "Abs in 30 Days", "tr": "30 Günde Karın Kası",
			"de": "Bauchmuskeln in 30 Tagen", "es": "Abdominales en 30 días", "fr": "Abdos en 30 jours",
			"it": "Addominali in 30 giorni",
		},
		// a = lower abs, b = obliques and anti-rotation, c = whole trunk.
		plan: coursePlan{
			warmup:   []setDef{{"W01", 15}, {"W09", 20}, {"W03", 8}, {"W08", 20}, {"W04", 8}, {"W06", 20}},
			cooldown: []setDef{{"CD04", 30}, {"CD03", 25}, {"CD07", 25}, {"CD05", 25}},
			a: []dose{
				{ID: "C05", From: 8, To: 16}, {ID: "C03", From: 15, To: 35}, {ID: "C11", From: 8, To: 16},
				{ID: "C04", From: 15, To: 35}, {ID: "C01", From: 20, To: 50},
			},
			b: []dose{
				{ID: "C13", From: 15, To: 40}, {ID: "C07", From: 15, To: 30}, {ID: "C12", From: 6, To: 12},
				{ID: "C09", From: 6, To: 14}, {ID: "C08", From: 5, To: 12, Since: 8},
			},
			c: []dose{
				{ID: "C06", From: 15, To: 30}, {ID: "C05", From: 8, To: 16}, {ID: "J04", From: 15, To: 30},
				{ID: "B02", From: 8, To: 14}, {ID: "C10", From: 4, To: 10, Since: 10}, {ID: "C02", From: 10, To: 25},
			},
			recovery: []dose{
				{ID: "C12", From: 6, To: 10}, {ID: "C11", From: 8, To: 12}, {ID: "C01", From: 20, To: 40},
				{ID: "B01", From: 20, To: 30, Since: 9},
			},
			final: []setDef{{"C05", 16}, {"C13", 40}, {"C06", 30}, {"C08", 12}, {"J04", 30}, {"C04", 35}, {"C01", 60}},
		},
	},
	{
		Key:  "slim30",
		Slug: "lose-weight-in-30-days",
		About: map[string]string{
			"ru": "Кардио-круги и силовые упражнения на всё тело, чтобы тратить больше энергии. В курсе есть прыжки; приседания с выпрыгиванием и бёрпи появляются только во второй половине.",
			"en": "Cardio circuits and whole-body strength work to burn more energy. This course includes jumps; squat jumps and burpees appear only in the second half.",
			"tr": "Daha fazla enerji harcamak için kardiyo devreleri ve tüm vücut kuvvet çalışması. Kursta zıplama var; sıçramalı squat ve burpee yalnızca ikinci yarıda gelir.",
			"de": "Cardio-Zirkel und Kraftübungen für den ganzen Körper, um mehr Energie zu verbrauchen. Der Kurs enthält Sprünge; Squat Jumps und Burpees kommen erst in der zweiten Hälfte.",
			"es": "Circuitos de cardio y fuerza para todo el cuerpo para gastar más energía. El curso incluye saltos; las sentadillas con salto y los burpees llegan solo en la segunda mitad.",
			"fr": "Des circuits cardio et du renforcement de tout le corps pour dépenser plus d'énergie. Le cours comporte des sauts ; squats sautés et burpees n'arrivent qu'en seconde moitié.",
			"it": "Circuiti cardio e forza per tutto il corpo per consumare più energia. Il corso include salti; squat con salto e burpee arrivano solo nella seconda metà.",
		},
		Title: map[string]string{
			"ru": "Похудеть за 30 дней", "en": "Lose Weight in 30 Days", "tr": "30 Günde Kilo Ver",
			"de": "Abnehmen in 30 Tagen", "es": "Adelgaza en 30 días", "fr": "Maigrir en 30 jours",
			"it": "Dimagrire in 30 giorni",
		},
		// The one course with jumps: cardio circuits mixed with whole-body
		// strength. Squat jumps and burpees arrive only in the second half.
		plan: coursePlan{
			warmup:   []setDef{{"W01", 15}, {"W09", 20}, {"W08", 20}, {"J02", 30}, {"W05", 8}, {"W04", 8}},
			cooldown: []setDef{{"CD02", 25}, {"CD03", 25}, {"CD01", 25}, {"CD07", 25}},
			a: []dose{
				{ID: "J03", From: 20, To: 40}, {ID: "L10", From: 10, To: 18}, {ID: "J04", From: 15, To: 35},
				{ID: "P08", From: 6, To: 12}, {ID: "C13", From: 15, To: 35}, {ID: "J07", From: 15, To: 30, Since: 6},
			},
			b: []dose{
				{ID: "L11", From: 6, To: 12}, {ID: "J06", From: 15, To: 30}, {ID: "C01", From: 20, To: 45},
				{ID: "L05", From: 6, To: 12}, {ID: "J03", From: 20, To: 40}, {ID: "C05", From: 8, To: 14},
			},
			c: []dose{
				{ID: "J07", From: 15, To: 35}, {ID: "L10", From: 12, To: 20}, {ID: "C09", From: 6, To: 14},
				{ID: "J04", From: 20, To: 40}, {ID: "P01", From: 4, To: 10, Since: 10},
				{ID: "J05", From: 6, To: 12, Since: 13}, {ID: "J08", From: 4, To: 10, Since: 17},
			},
			recovery: []dose{
				{ID: "L12", From: 6, To: 10}, {ID: "L06", From: 10, To: 14}, {ID: "C12", From: 6, To: 10},
				{ID: "B06", From: 8, To: 12, Since: 9},
			},
			final: []setDef{{"J03", 45}, {"L10", 20}, {"J08", 10}, {"J04", 40}, {"P01", 10}, {"C13", 40}, {"C01", 50}},
		},
	},
	{
		Key:  "pushups30",
		Slug: "push-ups-in-30-days",
		About: map[string]string{
			"ru": "От 4 до 15 отжиманий за подход. Усложнённые варианты — узкие, медленные, с паузой — добавляются по одному, а упражнения для спины уравновешивают нагрузку.",
			"en": "From 4 to 15 push-ups per set. Harder variations — narrow, slow, paused — join one at a time, and back work balances the pressing.",
			"tr": "Set başına 4 şınavdan 15'e. Dar, yavaş ve duraklamalı gibi zor varyasyonlar birer birer eklenir; sırt çalışması yükü dengeler.",
			"de": "Von 4 auf 15 Liegestütze pro Satz. Schwerere Varianten – eng, langsam, mit Pause – kommen einzeln dazu, Rückenübungen gleichen die Belastung aus.",
			"es": "De 4 a 15 flexiones por serie. Las variantes más duras —cerradas, lentas, con pausa— se suman una a una, y el trabajo de espalda equilibra la carga.",
			"fr": "De 4 à 15 pompes par série. Les variantes plus dures — serrées, lentes, avec pause — s'ajoutent une à une, et le travail du dos équilibre l'effort.",
			"it": "Da 4 a 15 flessioni per serie. Le varianti più dure — strette, lente, con pausa — si aggiungono una alla volta, e il lavoro per la schiena bilancia lo sforzo.",
		},
		Title: map[string]string{
			"ru": "Отжимания за 30 дней", "en": "Push-Ups in 30 Days", "tr": "30 Günde Şınav",
			"de": "Liegestütze in 30 Tagen", "es": "Flexiones en 30 días", "fr": "Pompes en 30 jours",
			"it": "Flessioni in 30 giorni",
		},
		// Classic push-ups grow 4 → 15 per round. Harder variations (paused,
		// narrow, slow, pike, staggered) join one at a time; every session
		// pairs pressing with pulling or back work.
		plan: coursePlan{
			warmup:   []setDef{{"W01", 20}, {"W02", 30}, {"W09", 20}, {"W03", 6}, {"W04", 8}, {"W06", 20}},
			cooldown: []setDef{{"CD01", 30}, {"CD04", 25}, {"CD07", 25}, {"CD03", 25}},
			a: []dose{
				{ID: "P01", From: 4, To: 14}, {ID: "B06", From: 8, To: 14}, {ID: "P08", From: 6, To: 12},
				{ID: "C01", From: 20, To: 50}, {ID: "P03", From: 3, To: 8, Since: 8},
			},
			b: []dose{
				{ID: "P08", From: 6, To: 12}, {ID: "B03", From: 8, To: 14}, {ID: "P01", From: 4, To: 12},
				{ID: "C09", From: 6, To: 14}, {ID: "P02", From: 3, To: 8, Since: 10},
				{ID: "P06", From: 3, To: 7, Since: 16},
			},
			c: []dose{
				{ID: "P01", From: 5, To: 15}, {ID: "B02", From: 8, To: 14}, {ID: "C10", From: 4, To: 10},
				{ID: "A03", From: 15, To: 25}, {ID: "P04", From: 3, To: 7, Since: 12},
				{ID: "P05", From: 3, To: 6, Since: 18},
			},
			recovery: []dose{
				{ID: "P08", From: 5, To: 8}, {ID: "B01", From: 20, To: 30}, {ID: "C12", From: 6, To: 10},
				{ID: "L06", From: 8, To: 12, Since: 9},
			},
			final: []setDef{{"P01", 16}, {"B06", 14}, {"P02", 8}, {"C10", 10}, {"P06", 7}, {"P03", 8}, {"C01", 50}},
		},
	},
	{
		Key:  "posture30",
		Slug: "posture-in-30-days",
		About: map[string]string{
			"ru": "Укрепление спины для ровной осанки: лопатки, разгибатели спины и вся задняя поверхность тела. Без прыжков и инвентаря, всё на коврике.",
			"en": "Back strengthening for better posture: shoulder blades, back extensors and the whole back of the body. No jumps, no equipment, all on a mat.",
			"tr": "Daha iyi duruş için sırt güçlendirme: kürek kemikleri, sırt kasları ve vücudun tüm arka hattı. Zıplama ve ekipman yok, hepsi matta.",
			"de": "Rückenkräftigung für eine bessere Haltung: Schulterblätter, Rückenstrecker und die gesamte Körperrückseite. Ohne Sprünge, ohne Geräte, alles auf der Matte.",
			"es": "Fortalecimiento de espalda para una mejor postura: escápulas, extensores y toda la cadena posterior. Sin saltos ni material, todo en la esterilla.",
			"fr": "Renforcement du dos pour une meilleure posture : omoplates, extenseurs du dos et toute la chaîne postérieure. Sans sauts ni matériel, tout sur le tapis.",
			"it": "Rinforzo della schiena per una postura migliore: scapole, estensori e tutta la catena posteriore. Niente salti né attrezzi, tutto sul tappetino.",
		},
		Title: map[string]string{
			"ru": "Осанка за 30 дней", "en": "Better Posture in 30 Days", "tr": "30 Günde Dik Duruş",
			"de": "Bessere Haltung in 30 Tagen", "es": "Mejor postura en 30 días",
			"fr": "Une meilleure posture en 30 jours", "it": "Postura migliore in 30 giorni",
		},
		// a = shoulder blades, b = back extensors, c = the whole posterior
		// chain. Strengthening only; no jumps.
		plan: coursePlan{
			warmup:   []setDef{{"W01", 20}, {"W02", 30}, {"W03", 8}, {"W08", 20}, {"W04", 8}, {"W07", 20}},
			cooldown: []setDef{{"CD01", 30}, {"CD07", 30}, {"CD03", 25}, {"CD02", 25}},
			a: []dose{
				{ID: "B03", From: 8, To: 14}, {ID: "B04", From: 6, To: 12}, {ID: "B01", From: 20, To: 40},
				{ID: "C12", From: 6, To: 12}, {ID: "L06", From: 8, To: 14}, {ID: "B05", From: 6, To: 12, Since: 8},
			},
			b: []dose{
				{ID: "B06", From: 8, To: 16}, {ID: "B07", From: 15, To: 35}, {ID: "B02", From: 8, To: 16},
				{ID: "C01", From: 20, To: 45}, {ID: "L03", From: 8, To: 14}, {ID: "B09", From: 8, To: 16, Since: 6},
			},
			c: []dose{
				{ID: "B10", From: 6, To: 12}, {ID: "B08", From: 15, To: 30}, {ID: "B11", From: 6, To: 12},
				{ID: "C07", From: 15, To: 25}, {ID: "B01", From: 25, To: 45}, {ID: "L04", From: 5, To: 10, Since: 10},
			},
			recovery: []dose{
				{ID: "B01", From: 20, To: 30}, {ID: "C12", From: 6, To: 10}, {ID: "L06", From: 8, To: 12},
				{ID: "B02", From: 8, To: 10},
			},
			final: []setDef{{"B06", 16}, {"B10", 12}, {"B07", 35}, {"B08", 30}, {"B04", 12}, {"C12", 12}, {"C01", 50}},
		},
	},
	{
		Key:  "legs30",
		Slug: "legs-and-glutes-in-30-days",
		About: map[string]string{
			"ru": "Приседания, ягодичные мостики и выпады без инвентаря. Нагрузка растёт от 10 до 24 приседаний за подход; прыжковые упражнения появляются с 12-го дня.",
			"en": "Squats, glute bridges and lunges with no equipment. Squats grow from 10 to 24 per set; jumping exercises come in from day 12.",
			"tr": "Ekipmansız squat, kalça köprüsü ve lunge. Squat sayısı set başına 10'dan 24'e çıkar; zıplamalı egzersizler 12. günden itibaren gelir.",
			"de": "Kniebeugen, Glute Bridges und Ausfallschritte ohne Geräte. Die Kniebeugen steigen von 10 auf 24 pro Satz; Sprungübungen kommen ab Tag 12 dazu.",
			"es": "Sentadillas, puentes de glúteos y zancadas sin material. Las sentadillas suben de 10 a 24 por serie; los ejercicios con salto llegan a partir del día 12.",
			"fr": "Squats, ponts fessiers et fentes sans matériel. Les squats passent de 10 à 24 par série ; les exercices sautés arrivent à partir du jour 12.",
			"it": "Squat, ponti per i glutei e affondi senza attrezzi. Gli squat salgono da 10 a 24 per serie; gli esercizi con salto arrivano dal giorno 12.",
		},
		Title: map[string]string{
			"ru": "Ноги и ягодицы за 30 дней", "en": "Legs & Glutes in 30 Days", "tr": "30 Günde Bacak ve Kalça",
			"de": "Beine & Po in 30 Tagen", "es": "Piernas y glúteos en 30 días",
			"fr": "Jambes et fessiers en 30 jours", "it": "Gambe e glutei in 30 giorni",
		},
		// a = squats, b = glutes and hinge, c = lunges and power. Squat jumps
		// come in from day 12, jump lunges only in the last week.
		plan: coursePlan{
			warmup:   []setDef{{"W09", 20}, {"W08", 20}, {"W04", 10}, {"W05", 8}, {"J02", 30}, {"W03", 6}},
			cooldown: []setDef{{"CD02", 30}, {"CD03", 25}, {"CD07", 25}, {"CD05", 25}},
			a: []dose{
				{ID: "L10", From: 10, To: 22}, {ID: "L06", From: 10, To: 18}, {ID: "L07", From: 12, To: 24},
				{ID: "L01", From: 8, To: 14}, {ID: "C01", From: 20, To: 40}, {ID: "L02", From: 5, To: 10, Since: 8},
			},
			b: []dose{
				{ID: "L05", From: 6, To: 12}, {ID: "L03", From: 10, To: 16}, {ID: "L12", From: 6, To: 12},
				{ID: "B06", From: 8, To: 12}, {ID: "L04", From: 5, To: 10, Since: 6},
				{ID: "L09", From: 4, To: 8, Since: 14},
			},
			c: []dose{
				{ID: "L11", From: 6, To: 12}, {ID: "L10", From: 12, To: 24}, {ID: "L06", From: 10, To: 18},
				{ID: "C07", From: 15, To: 25}, {ID: "J05", From: 5, To: 10, Since: 12},
				{ID: "J09", From: 3, To: 6, Since: 22},
			},
			recovery: []dose{
				{ID: "L06", From: 10, To: 14}, {ID: "L03", From: 8, To: 12}, {ID: "L07", From: 12, To: 16},
				{ID: "C12", From: 6, To: 8},
			},
			final: []setDef{{"L10", 25}, {"L11", 12}, {"L04", 10}, {"J05", 10}, {"L07", 25}, {"C01", 45}},
		},
	},
}
