package i18n

// exerciseNames maps exercise ID to names in Order = [ru, en, tr, de, es, fr, it].
var exerciseNames = map[string][7]string{
	// ---- Warm-up ----
	"W01": {"Круги плечами", "Shoulder circles", "Omuz çevirme", "Schulterkreisen", "Círculos de hombros", "Cercles d'épaules", "Cerchi con le spalle"},
	"W02": {"Малые круги руками", "Small arm circles", "Küçük kol çevirme", "Kleine Armkreise", "Círculos pequeños de brazos", "Petits cercles des bras", "Piccoli cerchi con le braccia"},
	"W03": {"Повороты корпуса стоя", "Standing twists", "Ayakta gövde dönüşü", "Rumpfdrehung im Stehen", "Giros de torso de pie", "Rotations du buste debout", "Torsioni del busto in piedi"},
	"W04": {"Наклон с прямой спиной", "Hip hinge", "Kalça menteşesi", "Hüftbeuge", "Bisagra de cadera", "Charnière de hanche", "Hip hinge"},
	"W05": {"Неглубокий присед", "Shallow squat", "Sığ çömelme", "Flache Kniebeuge", "Sentadilla poco profunda", "Squat peu profond", "Squat poco profondo"},
	"W06": {"Высокая планка с переносом веса", "High plank weight shift", "Yüksek plankta ağırlık aktarma", "Hoher Stütz mit Gewichtsverlagerung", "Plancha alta con cambio de peso", "Planche haute, transfert de poids", "Plank alto con spostamento del peso"},

	// ---- Abs / core ----
	"C01": {"Планка на локтях", "Forearm plank", "Dirsek plank", "Unterarmstütz", "Plancha de antebrazos", "Planche sur les avant-bras", "Plank sugli avambracci"},
	"C02": {"RKC-планка", "RKC plank", "RKC plank", "RKC-Plank", "Plancha RKC", "Planche RKC", "Plank RKC"},
	"C03": {"Ножницы", "Scissors", "Makas", "Schere", "Tijeras", "Ciseaux", "Forbici"},
	"C04": {"Флаттер", "Flutter kicks", "Çırpma hareketi", "Flutter Kicks", "Patada de aleteo", "Battements de jambes", "Flutter kicks"},
	"C05": {"Обратные скручивания", "Reverse crunch", "Ters mekik", "Reverse Crunch", "Encogimiento inverso", "Crunch inversé", "Crunch inverso"},
	"C06": {"Hollow hold (лёгкий)", "Hollow hold (easy)", "Hollow hold (kolay)", "Hollow Hold (leicht)", "Hollow hold (fácil)", "Hollow hold (facile)", "Hollow hold (facile)"},
	"C07": {"Боковая планка", "Side plank", "Yan plank", "Seitstütz", "Plancha lateral", "Planche latérale", "Plank laterale"},
	"C08": {"Боковая планка с подъёмом таза", "Side plank hip lift", "Yan plankta kalça kaldırma", "Seitstütz mit Hüftheben", "Plancha lateral con elevación de cadera", "Planche latérale, élévation de hanche", "Plank laterale con sollevamento dell'anca"},
	"C09": {"Планка с касанием плеч", "Plank shoulder taps", "Plankta omuz dokunuşu", "Plank mit Schultertippen", "Plancha con toques de hombro", "Planche, touches d'épaule", "Plank con tocco delle spalle"},
	"C10": {"Планка вверх-вниз", "Up-down plank", "Aşağı-yukarı plank", "Hoch-tief-Plank", "Plancha arriba-abajo", "Planche haute-basse", "Plank su e giù"},
	"C11": {"Касания пяток лёжа", "Heel taps", "Topuk dokunuşu", "Fersen-Tippen", "Toques de talón", "Touches de talon", "Tocchi del tallone"},
	"C12": {"Дэд-баг (усложнённый)", "Dead bug (advanced)", "Ölü böcek (ileri)", "Dead Bug (fortgeschritten)", "Bicho muerto (avanzado)", "Dead bug (avancé)", "Dead bug (avanzato)"},

	// ---- Push-ups / triceps / shoulders ----
	"P01": {"Отжимания", "Push-ups", "Şınav", "Liegestütze", "Flexiones", "Pompes", "Piegamenti"},
	"P02": {"Узкие отжимания", "Narrow push-ups", "Dar şınav", "Enge Liegestütze", "Flexiones cerradas", "Pompes serrées", "Piegamenti stretti"},
	"P03": {"Отжимания с паузой", "Paused push-ups", "Duraklamalı şınav", "Liegestütze mit Pause", "Flexiones con pausa", "Pompes avec pause", "Piegamenti con pausa"},
	"P04": {"Медленные отжимания 3-1-3", "Slow push-ups 3-1-3", "Yavaş şınav 3-1-3", "Langsame Liegestütze 3-1-3", "Flexiones lentas 3-1-3", "Pompes lentes 3-1-3", "Piegamenti lenti 3-1-3"},
	"P05": {"Отжимания со смещением рук", "Staggered push-ups", "Kaydırmalı şınav", "Versetzte Liegestütze", "Flexiones escalonadas", "Pompes décalées", "Piegamenti sfalsati"},
	"P06": {"Отжимания «домиком»", "Pike push-ups", "Pike şınav", "Pike-Liegestütze", "Flexiones pica", "Pompes piquées", "Piegamenti pike"},
	"P08": {"Отжимания сфинкса", "Sphinx push-ups", "Sfenks şınav", "Sphinx-Liegestütze", "Flexiones esfinge", "Pompes sphinx", "Piegamenti sfinge"},

	// ---- Back / scapula ----
	"B01": {"Кобра (удержание)", "Cobra hold", "Kobra tutuş", "Kobra halten", "Cobra isométrica", "Cobra tenue", "Cobra tenuta"},
	"B02": {"Обратные снежные ангелы", "Reverse snow angels", "Ters kar meleği", "Reverse Snow Angels", "Ángeles de nieve inversos", "Anges de neige inversés", "Angeli di neve inversi"},
	"B03": {"W-подъём лёжа", "Prone W-raise", "W kaldırma", "W-Heben", "Elevación en W", "Levée en W", "Sollevamento a W"},
	"B04": {"Y-подъём лёжа", "Prone Y-raise", "Y kaldırma", "Y-Heben", "Elevación en Y", "Levée en Y", "Sollevamento a Y"},
	"B05": {"T-подъём лёжа", "Prone T-raise", "T kaldırma", "T-Heben", "Elevación en T", "Levée en T", "Sollevamento a T"},
	"B06": {"Супермен с притягиванием", "Superman pull-down", "Superman çekiş", "Superman-Zug", "Superman con jalón", "Superman tirage", "Superman pull-down"},
	"B07": {"Пловец лёжа", "Swimmers", "Yüzücü", "Schwimmer", "Nadadores", "Nageur", "Nuotatore"},
	"B08": {"Обратная планка", "Reverse plank", "Ters plank", "Umgekehrter Stütz", "Plancha inversa", "Planche inversée", "Plank inverso"},
	"B09": {"Разгибания спины (пульс)", "Back extension pulses", "Sırt esnetme (pulse)", "Rückenstrecker-Pulse", "Extensiones de espalda (pulsos)", "Extensions du dos (pulsations)", "Estensioni schiena (pulse)"},
	"B10": {"Кобра с W-тягой", "Cobra to W-pull", "Kobra + W çekiş", "Kobra mit W-Zug", "Cobra con tirón en W", "Cobra avec tirage en W", "Cobra con tirata a W"},

	// ---- Arms / biceps ----
	"A01": {"Бицепс с самосопротивлением", "Self-resist curl", "Pazı (kendine direnç)", "Bizeps mit Eigenwiderstand", "Curl con autorresistencia", "Curl en auto-résistance", "Curl con autoresistenza"},
	"A02": {"Изометрия бицепса 90°", "Biceps isometric 90°", "Pazı izometrik 90°", "Bizeps isometrisch 90°", "Bíceps isométrico 90°", "Biceps isométrique 90°", "Bicipiti isometrico 90°"},
	"A03": {"«Крюк» ладонями", "Palm-hook pull", "Avuç kanca çekişi", "Handflächen-Haken", "Gancho de palmas", "Crochet des paumes", "Aggancio dei palmi"},
	"A04": {"Негатив с самосопротивлением", "Slow negative curl", "Yavaş negatif (kendine direnç)", "Langsame Negativ-Curls", "Negativo lento con autorresistencia", "Négatif lent en auto-résistance", "Negativo lento con autoresistenza"},

	// ---- Legs / glutes ----
	"L01": {"Неглубокий присед", "Shallow squat", "Sığ çömelme", "Flache Kniebeuge", "Sentadilla poco profunda", "Squat peu profond", "Squat poco profondo"},
	"L02": {"Присед с паузой", "Paused squat", "Duraklamalı çömelme", "Kniebeuge mit Pause", "Sentadilla con pausa", "Squat avec pause", "Squat con pausa"},
	"L03": {"Наклон с прямой спиной", "Hip hinge", "Kalça menteşesi", "Hüftbeuge", "Bisagra de cadera", "Charnière de hanche", "Hip hinge"},
	"L04": {"Мост на одной ноге", "Single-leg glute bridge", "Tek bacak köprü", "Einbeinige Glute Bridge", "Puente a una pierna", "Pont fessier une jambe", "Ponte glutei a una gamba"},
	"L05": {"Мост с шагами", "Glute bridge march", "Köprüde yürüyüş", "Glute Bridge March", "Puente con marcha", "Pont fessier en marche", "Ponte glutei con marcia"},
	"L06": {"Ягодичный мост с паузой", "Glute bridge hold", "Duraklamalı köprü", "Glute Bridge mit Pause", "Puente con pausa", "Pont fessier avec pause", "Ponte glutei con pausa"},
	"L07": {"Подъём на носки", "Calf raises", "Topuk kaldırma", "Wadenheben", "Elevación de pantorrillas", "Mollets debout", "Sollevamento polpacci"},
	"L09": {"Боковая планка с подъёмом ноги", "Side plank leg lift", "Yan plankta bacak kaldırma", "Seitstütz mit Beinheben", "Plancha lateral con elevación de pierna", "Planche latérale, élévation de jambe", "Plank laterale con sollevamento gamba"},

	// ---- Cool-down ----
	"CD01": {"Растяжка груди", "Chest stretch", "Göğüs esnetme", "Brustdehnung", "Estiramiento de pecho", "Étirement pectoral", "Allungamento pettorali"},
	"CD02": {"Растяжка задней поверхности бедра", "Hamstring stretch", "Arka bacak esnetme", "Beinrückseite dehnen", "Estiramiento de isquios", "Étirement des ischio-jambiers", "Allungamento femorali"},
	"CD03": {"Скручивание лёжа", "Supine twist", "Sırtüstü dönüş", "Liegende Drehung", "Torsión tumbado", "Torsion allongée", "Torsione sdraiato"},
	"CD04": {"Поза сфинкса", "Sphinx pose", "Sfenks pozu", "Sphinx-Haltung", "Postura de esfinge", "Posture du sphinx", "Posizione della sfinge"},
	"CD05": {"Спокойное дыхание лёжа", "Lying breathing", "Sırtüstü sakin nefes", "Ruhige Atmung im Liegen", "Respiración tumbada", "Respiration allongée", "Respirazione sdraiato"},
}
