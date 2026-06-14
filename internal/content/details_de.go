package content

// detailsDE is the German rich content for every exercise (current library IDs).
var detailsDE = map[string]Detail{
	// ---- Aufwärmen ----
	"W01": {
		Desc:    "Ein einfaches Schulter-Aufwärmen für Nacken, Schultern und oberen Rücken.",
		HowTo:   []string{"Aufrecht stehen, Arme locker.", "Die Schultern langsam rückwärts kreisen.", "Dann langsam vorwärts.", "Die Bewegung bleibt weich und kontrolliert."},
		Correct: []string{"Nacken entspannt.", "Schultern bewegen sich weich.", "Arme locker.", "Körper aufrecht."},
		Wrong:   []string{"Stark mit den Schultern zucken.", "Zu schnell bewegen.", "Den Nacken anspannen.", "Den oberen Rücken runden."},
	},
	"W02": {
		Desc:    "Ein Schulter- und Oberrücken-Aufwärmen, das Deltas und Schulterstabilisatoren aktiviert.",
		HowTo:   []string{"Aufrecht stehen, Arme auf Schulterhöhe heben.", "Kleine Kreise vorwärts.", "Dann kleine Kreise rückwärts.", "Die Kreise klein und kontrolliert halten."},
		Correct: []string{"Arme auf Schulterhöhe.", "Kreise sind klein.", "Schultern bleiben unten.", "Nacken entspannt."},
		Wrong:   []string{"Große Schwünge machen.", "Schultern zu den Ohren heben.", "Den unteren Rücken überstrecken.", "Die Luft anhalten."},
	},
	"W03": {
		Desc:    "Eine sanfte Rotationsübung für Rumpf und obere Wirbelsäule.",
		HowTo:   []string{"Füße hüftbreit.", "Das Becken zeigt überwiegend nach vorn.", "Den Oberkörper zu einer Seite drehen.", "Zur Mitte zurück und zur anderen Seite drehen."},
		Correct: []string{"Füße bleiben stabil.", "Bewegung kommt aus dem Rumpf.", "Drehung ist weich.", "Keine Drehung in den Knien."},
		Wrong:   []string{"Die Knie mitdrehen.", "Zu schnell drehen.", "Nach hinten lehnen.", "Den Bewegungsumfang erzwingen."},
	},
	"W04": {
		Desc:    "Eine Bewegungsübung für Hüften, Beinrückseite und Kontrolle des unteren Rückens.",
		HowTo:   []string{"Aufrecht stehen, Füße hüftbreit.", "Die Knie leicht weich machen.", "Das Becken nach hinten schieben.", "Den Oberkörper mit geradem Rücken senken.", "Über die Hüften zurück nach oben."},
		Correct: []string{"Rücken bleibt gerade.", "Bewegung startet aus der Hüfte.", "Knie sind weich, nicht tief gebeugt.", "Nacken folgt der Linie der Wirbelsäule."},
		Wrong:   []string{"Den Rücken runden.", "Daraus eine Kniebeuge machen.", "Zu tief greifen.", "Den Kopf zu hoch heben."},
	},
	"W05": {
		Desc:    "Ein kontrolliertes Bein-Aufwärmen in sicherem, flachem Bereich.",
		HowTo:   []string{"Füße etwas breiter als hüftbreit.", "Das Becken leicht nach hinten schieben.", "Die Knie nur bis zu einer angenehmen Tiefe beugen.", "Langsam aufrichten."},
		Correct: []string{"Knie in Richtung der Zehen.", "Fersen am Boden.", "Rücken lang.", "Tiefe moderat."},
		Wrong:   []string{"Zu tief gehen.", "Die Knie nach innen fallen lassen.", "Die Fersen heben.", "Schnell nach unten fallen."},
	},
	"W06": {
		Desc:    "Ein Aufwärmen für Schultern, Handgelenke und Rumpf.",
		HowTo:   []string{"In einer hohen Plank beginnen.", "Hände unter den Schultern.", "Den Körper leicht nach vorn verlagern.", "Zurück in die Ausgangsposition.", "Die Bewegung klein und kontrolliert."},
		Correct: []string{"Körper in einer Linie.", "Schultern stabil.", "Rumpf aktiv.", "Bewegung weich."},
		Wrong:   []string{"Das Becken absenken.", "Die Schultern zu den Ohren schieben.", "Zu weit nach vorn gehen.", "Den unteren Rücken durchhängen lassen."},
	},

	// ---- Bauch / Rumpf ----
	"C01": {
		Desc:    "Eine Rumpfstabilitätsübung für Bauch, Schultern, Gesäß und Rumpf.",
		HowTo:   []string{"Die Ellbogen unter die Schultern.", "Die Füße nach hinten setzen.", "Eine gerade Linie von Schultern zu Fersen.", "Bauch und Gesäß anspannen.", "Die Position halten."},
		Correct: []string{"Körper in einer Linie.", "Ellbogen unter den Schultern.", "Rumpf aktiv.", "Atmung gleichmäßig."},
		Wrong:   []string{"Den unteren Rücken durchhängen lassen.", "Das Becken zu hoch heben.", "Nach vorn schauen.", "Die Luft anhalten."},
	},
	"C02": {
		Desc:    "Eine kurze, intensive Plank mit Fokus auf Ganzkörperspannung.",
		HowTo:   []string{"In einer Unterarm-Plank beginnen.", "Den Bauch stark anspannen.", "Das Gesäß anspannen.", "Die Ellbogen gedanklich zu den Zehen ziehen.", "Mit maximaler Kontrolle halten."},
		Correct: []string{"Ganzkörperspannung.", "Gesäß aktiv.", "Rumpf fest.", "Atmung kontrolliert."},
		Wrong:   []string{"Wie eine lockere Plank halten.", "Das Becken durchhängen lassen.", "Die Luft anhalten.", "Zu lange halten wollen."},
	},
	"C03": {
		Desc:    "Eine Unterbauch-Übung mit kontrollierten kreuzenden Beinbewegungen.",
		HowTo:   []string{"Auf den Rücken legen.", "Den unteren Rücken sanft zum Boden drücken.", "Die Beine heben.", "Ein Bein über das andere kreuzen.", "Das Kreuzen abwechselnd fortsetzen."},
		Correct: []string{"Unterer Rücken nah am Boden.", "Beine kontrolliert.", "Nacken entspannt.", "Bewegung weich."},
		Wrong:   []string{"Den unteren Rücken überstrecken.", "Die Beine zu tief senken.", "Den Kopf nach vorn ziehen.", "Zu schnell bewegen."},
	},
	"C04": {
		Desc:    "Eine Unterbauch-Übung mit kurzen, wechselnden Beinschlägen.",
		HowTo:   []string{"Auf den Rücken legen.", "Den unteren Rücken zum Boden drücken.", "Die Beine heben.", "Die Beine in kurzen, wechselnden Schlägen auf und ab bewegen."},
		Correct: []string{"Kleiner Bewegungsumfang.", "Unterer Rücken stabil.", "Beine sinken nicht zu tief.", "Atmung gleichmäßig."},
		Wrong:   []string{"Den unteren Rücken überstrecken.", "Große Schläge machen.", "Den Nacken anspannen.", "Die Luft anhalten."},
	},
	"C05": {
		Desc:    "Eine Bauchübung mit Fokus auf kontrolliertes Heben des Beckens.",
		HowTo:   []string{"Auf den Rücken legen.", "Knie beugen, Füße heben.", "Mit dem Bauch das Becken leicht nach oben rollen.", "Langsam absenken."},
		Correct: []string{"Bewegung kommt aus dem Unterbauch.", "Becken hebt kontrolliert.", "Nacken entspannt.", "Absenken langsam."},
		Wrong:   []string{"Die Beine schwingen.", "Schwung benutzen.", "Am Nacken ziehen.", "Das Becken schnell fallen lassen."},
	},
	"C06": {
		Desc:    "Ein statisches Rumpf-Halten mit verkürztem Hebel für bessere Kontrolle.",
		HowTo:   []string{"Auf den Rücken legen.", "Den unteren Rücken zum Boden drücken.", "Die Schultern leicht heben.", "Knie gebeugt oder Beine höher.", "Die Position halten."},
		Correct: []string{"Unterer Rücken am Boden.", "Bauch fest.", "Nacken entspannt.", "Atmung läuft weiter."},
		Wrong:   []string{"Den unteren Rücken überstrecken.", "Die Beine zu tief senken.", "Das Kinn nach vorn ziehen.", "Die Luft anhalten."},
	},
	"C07": {
		Desc:    "Eine seitliche Rumpfübung für die schrägen Bauchmuskeln, Schulter- und Hüftstabilität.",
		HowTo:   []string{"Auf eine Seite legen.", "Den Ellbogen unter die Schulter.", "Das Becken heben.", "Der Körper in einer geraden Linie.", "Halten."},
		Correct: []string{"Ellbogen unter der Schulter.", "Becken oben.", "Körper lang.", "Nacken neutral."},
		Wrong:   []string{"Das Becken absenken.", "Die Schulter zum Ohr ziehen.", "Die Brust nach vorn rollen.", "Die Luft anhalten."},
	},
	"C08": {
		Desc:    "Eine stärkere Seitstütz-Variante für die schrägen Bauchmuskeln und die seitliche Hüfte.",
		HowTo:   []string{"In einem Seitstütz beginnen.", "Das Becken leicht senken.", "Das Becken wieder heben.", "Mit Kontrolle wiederholen.", "Seite wechseln."},
		Correct: []string{"Bewegung klein und kontrolliert.", "Ellbogen unter der Schulter.", "Becken bewegt sich vertikal.", "Rumpf fest."},
		Wrong:   []string{"Zu tief absinken.", "Den Rumpf verdrehen.", "In die Schulter einsinken.", "Zu schnell bewegen."},
	},
	"C10": {
		Desc:    "Eine Rumpf- und Trizeps-Übung im Wechsel zwischen Unterarm- und hoher Plank.",
		HowTo:   []string{"In einer Unterarm-Plank beginnen.", "Eine Hand auf den Boden setzen und hochdrücken.", "Die andere Hand in die hohe Plank bringen.", "Zurück auf die Unterarme.", "Die führende Hand abwechseln."},
		Correct: []string{"Becken stabil.", "Rumpf aktiv.", "Bewegung kontrolliert.", "Schultern über Händen oder Ellbogen."},
		Wrong:   []string{"Das Becken seitlich schaukeln.", "Den unteren Rücken durchhängen lassen.", "Zu schnell bewegen.", "Die Schultern einsinken lassen."},
	},
	"C11": {
		Desc:    "Eine kontrollierte Bauchübung und eine sicherere Alternative zu schwereren Beinheben.",
		HowTo:   []string{"Auf den Rücken legen.", "Die Beine in die Tabletop-Position heben.", "Eine Ferse zum Boden senken.", "Leicht antippen und zurück.", "Seiten wechseln."},
		Correct: []string{"Unterer Rücken stabil.", "Bewegung langsam.", "Ferse berührt sanft.", "Becken ruhig."},
		Wrong:   []string{"Das Bein zu schnell fallen lassen.", "Den unteren Rücken überstrecken.", "Beide Beine gleichzeitig bewegen.", "Hetzen."},
	},
	"C12": {
		Desc:    "Eine Rumpfkontroll-Übung für die tiefe Bauchmuskulatur und Wirbelsäulenstabilität.",
		HowTo:   []string{"Auf den Rücken legen.", "Arme und Beine in die Tabletop-Position heben.", "Den gegenüberliegenden Arm und das Bein ausstrecken.", "Den unteren Rücken stabil halten.", "Zurück und Seite wechseln."},
		Correct: []string{"Unterer Rücken nah am Boden.", "Bewegung langsam.", "Gegenüberliegender Arm und Bein bewegen sich zusammen.", "Rumpf bleibt ruhig."},
		Wrong:   []string{"Den unteren Rücken überstrecken.", "Zu schnell bewegen.", "Das Bein zu tief senken.", "Die Rippen herausdrücken."},
	},

	// ---- Liegestütze / Trizeps / Schultern ----
	"P01": {
		Desc:    "Eine grundlegende Oberkörperübung für Brust, Trizeps, Schultern und Rumpf.",
		HowTo:   []string{"In einer hohen Plank beginnen.", "Hände unter den Schultern oder etwas breiter.", "Den Körper kontrolliert senken.", "Hochdrücken und den Körper gerade halten."},
		Correct: []string{"Körper in einer Linie.", "Rumpf fest.", "Ellbogen etwa 30–45 Grad.", "Bewegung kontrolliert."},
		Wrong:   []string{"Das Becken durchhängen lassen.", "Das Becken zu hoch heben.", "Die Ellbogen gerade abspreizen.", "In die untere Position fallen."},
	},
	"P02": {
		Desc:    "Eine Liegestütz-Variante mit mehr Fokus auf den Trizeps.",
		HowTo:   []string{"In einer hohen Plank beginnen.", "Die Hände enger als bei normalen Liegestützen setzen.", "Die Ellbogen nah am Körper halten.", "Langsam senken und hochdrücken."},
		Correct: []string{"Ellbogen nah.", "Körper gerade.", "Schultern weg von den Ohren.", "Handgelenke angenehm."},
		Wrong:   []string{"Die Hände zu eng setzen.", "Die Ellbogen abspreizen.", "Den unteren Rücken durchhängen lassen.", "Bei Handgelenk- oder Ellbogenschmerz weitermachen."},
	},
	"P03": {
		Desc:    "Kontrollierte Liegestütze mit einer kurzen Pause in der unteren Position.",
		HowTo:   []string{"In normaler Liegestützposition beginnen.", "Langsam senken.", "Unten eine Sekunde pausieren.", "Hochdrücken, ohne die Körperposition zu verlieren."},
		Correct: []string{"Pause kontrolliert.", "Rumpf fest.", "Brust aktiv.", "Körper bleibt gerade."},
		Wrong:   []string{"Unten entspannen.", "Das Becken durchhängen lassen.", "Die Luft anhalten.", "Aus der unteren Position springen."},
	},
	"P04": {
		Desc:    "Eine langsame Liegestütz-Variante, die Kontrolle und Spannungsdauer erhöht.",
		HowTo:   []string{"In einer hohen Plank beginnen.", "Etwa drei Sekunden senken.", "Unten eine Sekunde pausieren.", "Etwa drei Sekunden hochdrücken."},
		Correct: []string{"Langsames, gleichmäßiges Tempo.", "Körper gerade.", "Ellbogen kontrolliert.", "Atmung gleichmäßig."},
		Wrong:   []string{"Die Bewegung hetzen.", "Unten zusammensacken.", "Das Becken durchhängen lassen.", "Die Rumpfspannung verlieren."},
	},
	"P05": {
		Desc:    "Eine asymmetrische Liegestütz-Variante, die Brust, Schultern und Rumpf fordert.",
		HowTo:   []string{"In Liegestützposition beginnen.", "Eine Hand etwas vor, die andere etwas zurück setzen.", "Kontrolliert senken.", "Hochdrücken.", "Auf der anderen Seite die Handposition wechseln."},
		Correct: []string{"Körper bleibt gerade.", "Rumpf fest.", "Beide Schultern stabil.", "Bewegung kontrolliert."},
		Wrong:   []string{"Den Rumpf verdrehen.", "Eine Schulter absenken.", "Die Hände zu breit setzen.", "Den Umfang bei Schulterbeschwerden erzwingen."},
	},
	"P06": {
		Desc:    "Eine schulterbetonte Liegestütz-Variante.",
		HowTo:   []string{"In einer hohen Plank beginnen.", "Die Füße etwas näher setzen und das Becken heben.", "Den Kopf zwischen den Armen.", "Den Kopf Richtung Boden senken.", "Wieder hochdrücken."},
		Correct: []string{"Becken hoch.", "Ellbogen kontrolliert beugen.", "Schultern stabil.", "Nacken neutral."},
		Wrong:   []string{"In eine normale Liegestütze sinken.", "Den Kopf zu weit nach vorn führen.", "Die Ellbogen zu stark abspreizen.", "Bei Schulterschmerz weitermachen."},
	},
	"P08": {
		Desc:    "Eine Trizeps- und Schulterübung aus dem Unterarmstütz.",
		HowTo:   []string{"Auf den Unterarmen beginnen, der Körper lang.", "Ellbogen unter oder leicht vor den Schultern.", "Über Hände und Unterarme drücken.", "Die Ellbogen leicht strecken.", "Kontrolliert zurück."},
		Correct: []string{"Rumpf fest.", "Ellbogen bewegen sich weich.", "Schultern unten.", "Bewegungsumfang angenehm."},
		Wrong:   []string{"Die Ellbogen erzwingen.", "Das Becken durchhängen lassen.", "Die Schultern hochziehen.", "Bei Ellbogen- oder Schulterschmerz bewegen."},
	},

	// ---- Rücken / Haltung ----
	"B01": {
		Desc:    "Eine haltungsorientierte Rückenübung für oberen Rücken, Schulterblätter und Rückenstrecker.",
		HowTo:   []string{"Auf den Bauch legen.", "Die Arme am Körper.", "Die Daumen leicht nach außen drehen.", "Die Brust leicht heben.", "Die Schulterblätter nach hinten und unten ziehen.", "Halten."},
		Correct: []string{"Heben ist klein.", "Nacken neutral.", "Schulterblätter nach hinten und unten.", "Gesäß leicht aktiv."},
		Wrong:   []string{"Zu hoch heben.", "Den Kopf nach hinten werfen.", "Den unteren Rücken überstrecken.", "Die Schultern zu den Ohren ziehen."},
	},
	"B02": {
		Desc:    "Eine Übung für oberen Rücken und Schulterblätter in Bauchlage.",
		HowTo:   []string{"Auf den Bauch legen.", "Der Kopf schaut nach unten.", "Die Arme in einem weiten Bogen von den Seiten nach oben führen.", "Auf demselben Weg zurück.", "Bewegung kontrolliert."},
		Correct: []string{"Nacken entspannt.", "Schulterblätter bewegen sich weich.", "Brustheben klein oder neutral.", "Arme ohne Hast."},
		Wrong:   []string{"Den Kopf heben.", "Den unteren Rücken zu stark überstrecken.", "Die Arme schwingen.", "Die Schultern zu den Ohren ziehen."},
	},
	"B03": {
		Desc:    "Eine Oberrücken-Übung für Schulterblätter und hintere Schultern.",
		HowTo:   []string{"Auf den Bauch legen.", "Die Ellbogen in eine W-Form beugen.", "Ellbogen und Hände leicht heben.", "Die Schulterblätter zusammenziehen.", "Kontrolliert absenken."},
		Correct: []string{"Kleiner Bewegungsumfang.", "Die Arbeit machen die Schulterblätter.", "Nacken neutral.", "Schultern weg von den Ohren."},
		Wrong:   []string{"Zu hoch heben.", "Den Kopf heben.", "Schwung benutzen.", "Den unteren Rücken überstrecken."},
	},
	"B04": {
		Desc:    "Eine Schulterblatt-Übung für den unteren Trapez und oberen Rücken.",
		HowTo:   []string{"Auf den Bauch legen.", "Die Arme diagonal über Kopf in eine Y-Form strecken.", "Die Arme leicht vom Boden heben.", "Kontrolliert absenken."},
		Correct: []string{"Arme heben nur wenig.", "Nacken neutral.", "Schultern unten.", "Bewegung kontrolliert."},
		Wrong:   []string{"Zu hoch heben.", "Die Schultern zu den Ohren ziehen.", "Den unteren Rücken überstrecken.", "Zu schnell bewegen."},
	},
	"B05": {
		Desc:    "Eine Oberrücken-Übung für hintere Schultern und Schulterblattkontrolle.",
		HowTo:   []string{"Auf den Bauch legen.", "Die Arme in einer T-Form zur Seite strecken.", "Die Arme leicht heben.", "Die Schulterblätter zusammenziehen.", "Langsam absenken."},
		Correct: []string{"Arme auf Schulterlinie.", "Nacken entspannt.", "Schulterblätter bewegen sich sanft.", "Heben ist klein."},
		Wrong:   []string{"Die Arme nach oben werfen.", "Den Kopf heben.", "Die Schultern hochziehen.", "Einen großen Umfang erzwingen."},
	},
	"B06": {
		Desc:    "Eine Rückenübung, die eine Zugbewegung ohne Geräte imitiert.",
		HowTo:   []string{"Auf den Bauch legen, Arme über Kopf.", "Die Brust leicht heben.", "Die Ellbogen Richtung Rippen ziehen.", "Die Schulterblätter zusammenziehen.", "Die Arme wieder nach vorn strecken."},
		Correct: []string{"Ellbogen ziehen kontrolliert nach unten.", "Schulterblätter ziehen sanft zusammen.", "Nacken neutral.", "Unterer Rücken angenehm."},
		Wrong:   []string{"Mit den Armen reißen.", "Zu hoch heben.", "Den Kopf nach hinten werfen.", "Die Schultern zu den Ohren heben."},
	},
	"B07": {
		Desc:    "Eine Rücken- und Koordinationsübung in Bauchlage.",
		HowTo:   []string{"Auf den Bauch legen.", "Die Brust leicht heben.", "Gegenüberliegenden Arm und Bein in einem kontrollierten Schwimmmuster bewegen.", "Die Bewegung klein und gleichmäßig."},
		Correct: []string{"Nacken neutral.", "Bewegung kontrolliert.", "Unterer Rücken angenehm.", "Atmung läuft weiter."},
		Wrong:   []string{"Zu schnell bewegen.", "Zu hoch heben.", "Den Kopf nach hinten werfen.", "Den unteren Rücken überlasten."},
	},
	"B08": {
		Desc:    "Eine Übung für die hintere Kette: Gesäß, Beinrückseite, Schultern und Rücken.",
		HowTo:   []string{"Mit gestreckten Beinen auf den Boden setzen.", "Die Hände hinter dem Becken.", "Über Hände und Fersen drücken.", "Das Becken heben.", "Den Körper in einer langen Linie halten."},
		Correct: []string{"Brust offen.", "Becken oben.", "Schultern stabil.", "Nacken neutral."},
		Wrong:   []string{"Das Becken absenken.", "Die Schultern hochziehen.", "Den Kopf nach hinten werfen.", "Die Handgelenke überlasten."},
	},
	"B09": {
		Desc:    "Eine Rückenübung mit kleinem Umfang für Rückenstrecker und Haltung.",
		HowTo:   []string{"Auf den Bauch legen.", "Die Hände neben den Körper oder leicht zur Brust.", "Die Brust leicht heben.", "Kleine, kontrollierte Impulse machen.", "Nacken neutral."},
		Correct: []string{"Kleines Heben.", "Weiche Impulse.", "Nacken lang.", "Unterer Rücken kontrolliert."},
		Wrong:   []string{"Zu hoch pulsieren.", "Schwung benutzen.", "Nach vorn schauen.", "Stechenden Schmerz im unteren Rücken spüren."},
	},
	"B10": {
		Desc:    "Eine stärkere Rückenübung, die ein kleines Kobra-Heben mit einem W-förmigen Armzug verbindet.",
		HowTo:   []string{"Auf den Bauch legen, Arme nach vorn.", "Die Brust leicht heben.", "Die Ellbogen nach unten und hinten in eine W-Form ziehen.", "Die Schulterblätter zusammenziehen.", "Wieder nach vorn strecken und kontrolliert absenken."},
		Correct: []string{"Kleines Brustheben.", "Weicher W-Zug.", "Schulterblätter nach hinten und unten.", "Nacken lang."},
		Wrong:   []string{"Den unteren Rücken überstrecken.", "Mit dem Nacken ziehen.", "Zu schnell bewegen.", "Die Schultern hochziehen."},
	},

	// ---- Arme / Bizeps ----
	"A01": {
		Desc:    "Eine Bizeps-Übung ohne Geräte: eine Hand leistet der anderen Widerstand.",
		HowTo:   []string{"Aufrecht stehen oder sitzen.", "Einen Ellbogen wie bei einem Curl beugen.", "Mit der anderen Hand auf den Unterarm drücken.", "Langsam gegen den Widerstand nach oben.", "Seite wechseln."},
		Correct: []string{"Widerstand gleichmäßig.", "Bewegung langsam.", "Ellbogen nah am Körper.", "Schulter unten."},
		Wrong:   []string{"Ohne echten Widerstand bewegen.", "Mit dem Arm reißen.", "Die Schulter heben.", "Die Luft anhalten."},
	},
	"A02": {
		Desc:    "Ein statisches Bizeps-Halten mit Eigenwiderstand.",
		HowTo:   []string{"Den Ellbogen auf etwa 90 Grad beugen.", "Mit der anderen Hand nach unten drücken.", "Der arbeitende Arm leistet ohne Bewegung Widerstand.", "Die Spannung halten.", "Seite wechseln."},
		Correct: []string{"Ellbogenwinkel stabil.", "Spannung gleichmäßig.", "Schultern entspannt.", "Atmung läuft weiter."},
		Wrong:   []string{"In kurzen Stößen drücken.", "Den Rumpf verdrehen.", "Die Schulter heben.", "Die Luft anhalten."},
	},
	"A03": {
		Desc:    "Eine isometrische Arm- und Unterarmübung mit den Händen gegeneinander.",
		HowTo:   []string{"Die Finger oder Hände einhaken.", "Die Ellbogen leicht gebeugt.", "Die Hände voneinander wegziehen.", "Gleichmäßige Spannung halten.", "Langsam lösen."},
		Correct: []string{"Spannung kontrolliert.", "Schultern unten.", "Handgelenke angenehm.", "Atmung ruhig."},
		Wrong:   []string{"Mit einem Ruck ziehen.", "Die Schultern hochziehen.", "Die Handgelenke überbeugen.", "Die Luft anhalten."},
	},
	"A04": {
		Desc:    "Eine langsame Absenkübung für den Bizeps mit Eigenwiderstand.",
		HowTo:   []string{"Mit gebeugtem Ellbogen beginnen.", "Mit der anderen Hand Widerstand erzeugen.", "Den Unterarm langsam absenken.", "Während des Absenkens Widerstand halten.", "Seite wechseln."},
		Correct: []string{"Absenken langsam.", "Widerstand konstant.", "Ellbogen nah am Körper.", "Schulter entspannt."},
		Wrong:   []string{"Den Arm schnell fallen lassen.", "Keinen Widerstand benutzen.", "Den Rumpf verdrehen.", "Die Luft anhalten."},
	},

	// ---- Beine / Gesäß ----
	"L01": {
		Desc:    "Eine kontrollierte Beinübung für Oberschenkel, Hüften und Haltung.",
		HowTo:   []string{"Füße etwas breiter als hüftbreit.", "Das Becken nach hinten schieben.", "Die Knie bis zu einer angenehmen, flachen Tiefe beugen.", "Langsam aufrichten."},
		Correct: []string{"Knie in Richtung der Zehen.", "Fersen unten.", "Rücken lang.", "Tiefe angenehm."},
		Wrong:   []string{"Zu tief gehen.", "Die Knie nach innen fallen lassen.", "Die Fersen heben.", "Schnell fallen."},
	},
	"L02": {
		Desc:    "Eine kontrollierte Kniebeuge-Variante mit kurzem Halten im sicheren Bereich.",
		HowTo:   []string{"Stabil stehen.", "In eine angenehme, flache Kniebeuge senken.", "Kurz pausieren.", "Kontrolliert aufrichten."},
		Correct: []string{"Pause stabil.", "Knie in Richtung der Zehen.", "Brust offen.", "Fersen am Boden."},
		Wrong:   []string{"Zu tief gehen.", "In der unteren Position federn.", "Die Knie nach innen fallen lassen.", "Die Luft anhalten."},
	},
	"L03": {
		Desc:    "Eine Bewegungsübung für Hüften, Beinrückseite und Kontrolle des unteren Rückens.",
		HowTo:   []string{"Aufrecht stehen, Füße hüftbreit.", "Die Knie leicht weich machen.", "Das Becken nach hinten schieben.", "Den Oberkörper mit geradem Rücken senken.", "Über die Hüften zurück nach oben."},
		Correct: []string{"Rücken bleibt gerade.", "Bewegung startet aus der Hüfte.", "Knie sind weich, nicht tief gebeugt.", "Nacken folgt der Linie der Wirbelsäule."},
		Wrong:   []string{"Den Rücken runden.", "Daraus eine Kniebeuge machen.", "Zu tief greifen.", "Den Kopf zu hoch heben."},
	},
	"L04": {
		Desc:    "Eine stärkere Gesäßübung mit jeweils einem Bein.",
		HowTo:   []string{"Auf den Rücken legen.", "Ein Knie beugen, der Fuß bleibt am Boden.", "Das andere Bein strecken oder heben.", "Über die arbeitende Ferse drücken.", "Das Becken heben und langsam senken."},
		Correct: []string{"Heben aus dem arbeitenden Gesäß.", "Becken bleibt gerade.", "Bewegung kontrolliert.", "Unterer Rücken neutral."},
		Wrong:   []string{"Das Becken verdrehen.", "Über die Zehen drücken.", "Den unteren Rücken überstrecken.", "Zu schnell fallen."},
	},
	"L05": {
		Desc:    "Eine Gesäß- und Beckenstabilitätsübung.",
		HowTo:   []string{"In einer Gesäßbrücke beginnen.", "Das Becken oben halten.", "Einen Fuß leicht vom Boden heben.", "Absenken und Seite wechseln.", "Das Becken stabil halten."},
		Correct: []string{"Becken bleibt gerade.", "Gesäß aktiv.", "Bewegung langsam.", "Unterer Rücken angenehm."},
		Wrong:   []string{"Das Becken absinken lassen.", "Seitlich schaukeln.", "Zu schnell bewegen.", "Den unteren Rücken überstrecken."},
	},
	"L06": {
		Desc:    "Eine Gesäßübung, die Hüften und hintere Kette stärkt.",
		HowTo:   []string{"Auf den Rücken legen.", "Knie beugen, Füße am Boden.", "Das Becken heben.", "Oben pausieren.", "Langsam absenken."},
		Correct: []string{"Über die Fersen drücken.", "Gesäß oben anspannen.", "Rippen unten.", "Unterer Rücken überstreckt nicht."},
		Wrong:   []string{"Oben den unteren Rücken überstrecken.", "Die Fersen heben.", "Das Becken zu schnell fallen lassen.", "Die Füße zu nah setzen."},
	},
	"L07": {
		Desc:    "Eine Unterschenkelübung für Waden, Sprunggelenke und Fußkontrolle.",
		HowTo:   []string{"Aufrecht stehen, Füße hüftbreit.", "Auf die Fußballen heben.", "Oben kurz pausieren.", "Die Fersen langsam senken."},
		Correct: []string{"Körper aufrecht.", "Bewegung weich.", "Sprunggelenke ausgerichtet.", "Absenken kontrolliert."},
		Wrong:   []string{"Die Fersen schnell fallen lassen.", "Nach vorn schaukeln.", "Die Sprunggelenke nach außen kippen.", "Schwung benutzen."},
	},
	"L09": {
		Desc:    "Eine Übung für die seitliche Kette: seitliches Gesäß, schräge Bauchmuskeln und Hüftstabilität.",
		HowTo:   []string{"In einem Seitstütz beginnen.", "Der Körper in einer geraden Linie.", "Das obere Bein leicht heben.", "Kontrolliert absenken.", "Wiederholen und Seite wechseln."},
		Correct: []string{"Becken oben.", "Oberes Bein bewegt sich langsam.", "Rumpf stabil.", "Schulter stark."},
		Wrong:   []string{"Das Becken absenken.", "Das Bein schwingen.", "Den Körper nach hinten rollen.", "In die Schulter einsinken."},
	},

	// ---- Abwärmen ----
	"CD01": {
		Desc:    "Eine sanfte Dehnung für Brust und vordere Schultern.",
		HowTo:   []string{"Auf den Rücken legen.", "Die Arme zur Seite öffnen.", "Die Schultern entspannen.", "Die Brust sanft öffnen lassen.", "Langsam atmen."},
		Correct: []string{"Brust öffnet sanft.", "Schultern entspannt.", "Unterer Rücken angenehm.", "Kein Schmerz in den Schultern."},
		Wrong:   []string{"Die Arme nach unten zwingen.", "Den unteren Rücken überstrecken.", "Den Nacken anspannen.", "Durch Schmerz dehnen."},
	},
	"CD02": {
		Desc:    "Eine sanfte Dehnung für die Rückseite des Oberschenkels.",
		HowTo:   []string{"Auf den Rücken legen.", "Ein Bein heben.", "Hinter Oberschenkel oder Wade fassen.", "Das Bein sanft zu sich ziehen.", "Seite wechseln."},
		Correct: []string{"Dehnung mild.", "Knie darf leicht gebeugt bleiben.", "Unterer Rücken ruhig.", "Atmung langsam."},
		Wrong:   []string{"Zu stark ziehen.", "Das Bein gewaltsam strecken.", "Das Becken heben.", "Stechenden Schmerz hinter dem Knie spüren."},
	},
	"CD03": {
		Desc:    "Eine sanfte Rotationsdehnung für Wirbelsäule und Rumpf.",
		HowTo:   []string{"Auf den Rücken legen.", "Die Knie beugen.", "Die Knie sanft zu einer Seite sinken lassen.", "Die Schultern am Boden halten.", "Zur Mitte zurück und Seite wechseln."},
		Correct: []string{"Bewegung langsam.", "Schultern am Boden.", "Atmung entspannt.", "Dehnung angenehm."},
		Wrong:   []string{"Die Knie zum Boden zwingen.", "Schnell drehen.", "Die Schulter heben.", "Durch Schmerz dehnen."},
	},
	"CD04": {
		Desc:    "Eine sanfte Dehnung der Körpervorderseite und eine milde Rückenstreckung.",
		HowTo:   []string{"Auf den Bauch legen.", "Die Unterarme auf den Boden.", "Die Brust sanft heben.", "Den Nacken lang halten.", "Langsam atmen."},
		Correct: []string{"Heben ist mild.", "Schultern weg von den Ohren.", "Unterer Rücken angenehm.", "Nacken neutral."},
		Wrong:   []string{"Den unteren Rücken überstrecken.", "Den Kopf nach hinten werfen.", "Die Schultern hochziehen.", "Die Position erzwingen."},
	},
	"CD05": {
		Desc:    "Eine ruhige Atemübung zum Abschluss der Einheit.",
		HowTo:   []string{"Auf den Rücken legen.", "Schultern und Kiefer entspannen.", "Wenn angenehm, eine Hand auf den Bauch legen.", "Sanft einatmen.", "Langsam ausatmen."},
		Correct: []string{"Atmung ruhig.", "Schultern entspannt.", "Gesicht weich.", "Der Körper kommt zur Ruhe."},
		Wrong:   []string{"Zu kräftig atmen.", "Die Luft anhalten.", "Den Nacken anspannen.", "Den unteren Rücken überstrecken."},
	},

	// ---- Vlad-Set Ergänzungen (Aufwärmen / Cardio / Plyometrie / Ausfallschritte) ----
	"W07": {
		Desc:    "Eine ruhige Eröffnungsübung: tiefe Atmung mit einem sanften Ganzkörper-Strecken, das die Wirbelsäule länger macht.",
		HowTo:   []string{"Aufrecht stehen, Füße hüftbreit.", "Einatmen und beide Arme über den Kopf strecken.", "Die Wirbelsäule sanft in die Länge ziehen.", "Ausatmen und die Arme senken, die Schultern fallen lassen."},
		Correct: []string{"Atmung langsam und voll.", "Schultern entspannen beim Ausatmen.", "Rippen bleiben unten, kein Überstrecken.", "Bewegung ohne Hast."},
		Wrong:   []string{"Die Luft anhalten.", "Den unteren Rücken stark überstrecken.", "Die Schultern zu den Ohren ziehen.", "Das Strecken hetzen."},
	},
	"W08": {
		Desc:    "Eine sanfte Rotationsübung, um Wirbelsäule und Rumpf aufzuwärmen.",
		HowTo:   []string{"Füße hüftbreit.", "Die Arme locker schwingen lassen.", "Den Oberkörper zu einer Seite drehen.", "Weich zur anderen Seite fließen."},
		Correct: []string{"Bewegung kommt aus dem Rumpf.", "Becken zeigt überwiegend nach vorn.", "Füße bleiben stabil.", "Tempo weich und gleichmäßig."},
		Wrong:   []string{"Stark in den Knien drehen.", "Mit den Armen reißen.", "Zu schnell drehen.", "Die Luft anhalten."},
	},
	"C13": {
		Desc:    "Eine dynamische Rumpfübung für Bauch und schräge Bauchmuskeln.",
		HowTo:   []string{"Auf den Rücken legen, Hände locker hinter dem Kopf.", "Die Schultern heben und die Knie anziehen.", "Einen Ellbogen Richtung gegenüberliegendes Knie führen.", "In einer gleichmäßigen Trittbewegung die Seiten wechseln."},
		Correct: []string{"Unterer Rücken bleibt am Boden.", "Bewegung kontrolliert, nicht gerissen.", "Ellbogen bleiben weit.", "Atmung gleichmäßig."},
		Wrong:   []string{"Am Nacken ziehen.", "Den unteren Rücken vom Boden heben.", "Den Rhythmus hetzen.", "Die Luft anhalten."},
	},
	"B11": {
		Desc:    "Eine Rückenübung, die gegenüberliegenden Arm und Bein in einer wechselnden, schwimmähnlichen Bewegung hebt.",
		HowTo:   []string{"Auf den Bauch legen, Arme über Kopf.", "Brust und Beine leicht heben.", "Einen Arm und das gegenüberliegende Bein anheben.", "Absenken und Seite wechseln, weich abwechselnd."},
		Correct: []string{"Heben bleibt klein.", "Nacken folgt der Wirbelsäule.", "Gesäß bleibt aktiv.", "Bewegung gleichmäßig."},
		Wrong:   []string{"Den unteren Rücken überlasten.", "Den Kopf heben.", "In Rucken bewegen.", "Die Luft anhalten."},
	},
	"L10": {
		Desc:    "Eine grundlegende Beinübung für Oberschenkel, Gesäß und Hüften.",
		HowTo:   []string{"Füße schulterbreit.", "Das Becken nach hinten schieben und die Knie beugen.", "Senken, bis die Oberschenkel fast parallel sind.", "Über die Fersen nach oben drücken."},
		Correct: []string{"Knie in Richtung der Zehen.", "Fersen bleiben am Boden.", "Brust bleibt offen.", "Rücken bleibt neutral."},
		Wrong:   []string{"Die Knie nach innen fallen lassen.", "Die Fersen heben.", "Den Rücken runden.", "Ohne Kontrolle nach unten fallen."},
	},
	"L11": {
		Desc:    "Eine einbeinige Übung für Beine und Gesäß, die schonend für die Knie ist.",
		HowTo:   []string{"Aufrecht stehen, Füße hüftbreit.", "Mit einem Fuß einen Schritt zurück machen.", "Senken, bis beide Knie fast 90 Grad erreichen.", "Über die vordere Ferse zurückdrücken."},
		Correct: []string{"Vorderes Knie bleibt über der Ferse.", "Oberkörper bleibt aufrecht.", "Schritt kontrolliert.", "Hinteres Knie senkt sich sanft."},
		Wrong:   []string{"Das vordere Knie nach innen fallen lassen.", "Den Oberkörper weit nach vorn lehnen.", "Das hintere Knie auf den Boden schlagen.", "Die Schritte hetzen."},
	},
	"L12": {
		Desc:    "Eine langsam ausgeführte Kniebeuge ohne Gewicht, um Kontrolle und Beinkraft aufzubauen.",
		HowTo:   []string{"Füße schulterbreit.", "Über etwa drei Sekunden senken.", "Unten kurz pausieren.", "Mit Kontrolle aufrichten."},
		Correct: []string{"Tempo bleibt langsam.", "Knie in Richtung der Zehen.", "Fersen bleiben am Boden.", "Rücken bleibt neutral."},
		Wrong:   []string{"Schnell nach unten fallen.", "Aus der unteren Position federn.", "Die Knie nach innen fallen lassen.", "Die Luft anhalten."},
	},
	"J01": {
		Desc:    "Kleine, federnde Sprünge von den Fußballen, um Waden und Sprunggelenke vorzubereiten.",
		HowTo:   []string{"Aufrecht stehen, Füße zusammen.", "Leicht von den Fußballen abfedern.", "Die Knie leicht weich halten.", "Weich und leise landen."},
		Correct: []string{"Sprünge bleiben niedrig.", "Landungen sind leise.", "Knie bleiben weich.", "Körper bleibt aufrecht."},
		Wrong:   []string{"Hart und laut landen.", "Die Knie durchstrecken.", "Zu hoch springen.", "Die Sprunggelenke nach innen kippen."},
	},
	"J02": {
		Desc:    "Ein gelenkschonender Marsch mit Kniehub, ein sanftes Cardio-Aufwärmen.",
		HowTo:   []string{"Aufrecht stehen.", "Auf der Stelle marschieren und ein Knie bis auf Hüfthöhe heben.", "Den gegenüberliegenden Arm natürlich mitschwingen.", "Ein gleichmäßiges, angenehmes Tempo halten."},
		Correct: []string{"Oberkörper bleibt aufrecht.", "Knie heben bis auf angenehme Höhe.", "Schritte bleiben leicht.", "Atmung bleibt locker."},
		Wrong:   []string{"Nach hinten lehnen.", "Mit den Füßen stampfen.", "Die Luft anhalten.", "Die Kontrolle hetzen."},
	},
	"J03": {
		Desc:    "Eine klassische Ganzkörper-Cardio-Bewegung, die den ganzen Körper aufwärmt.",
		HowTo:   []string{"Füße zusammen, Arme an den Seiten.", "Die Füße weit springen, während die Arme über den Kopf gehen.", "Zurück in die Ausgangsposition springen.", "Einen gleichmäßigen Rhythmus halten."},
		Correct: []string{"Landungen bleiben weich.", "Knie in Richtung der Zehen.", "Bewegung rhythmisch.", "Atmung gleichmäßig."},
		Wrong:   []string{"Hart landen.", "Die Knie nach innen fallen lassen.", "Die Schultern hochziehen.", "Die Luft anhalten."},
	},
	"J04": {
		Desc:    "Eine dynamische Rumpf- und Cardio-Bewegung, die aus der Plank die Knie zur Brust treibt.",
		HowTo:   []string{"In einer hohen Plank beginnen, Hände unter den Schultern.", "Ein Knie Richtung Brust treiben.", "Schnell die Beine wechseln.", "Das Becken tief und stabil halten."},
		Correct: []string{"Becken bleibt gerade.", "Schultern bleiben über den Händen.", "Rumpf bleibt fest.", "Tempo bleibt kontrolliert."},
		Wrong:   []string{"Das Becken nach oben springen lassen.", "Den unteren Rücken durchhängen lassen.", "Die Schultern zu weit nach vorn schieben.", "Die Luft anhalten."},
	},
	"J05": {
		Desc:    "Eine explosive Unterkörper-Bewegung: in die Kniebeuge gehen, dann kraftvoll nach oben springen.",
		HowTo:   []string{"Füße schulterbreit.", "In eine Kniebeuge senken.", "Explosiv nach oben springen.", "Weich landen und in die nächste Kniebeuge abfedern."},
		Correct: []string{"Landungen weich und leise.", "Knie in Richtung der Zehen.", "Brust bleibt offen.", "Hüften federn die Landung ab."},
		Wrong:   []string{"Steif und laut landen.", "Die Knie nach innen fallen lassen.", "Den Rücken runden.", "Auf gestreckten Beinen landen."},
	},
	"J06": {
		Desc:    "Ein seitlicher Hüpfer von Seite zu Seite für Kraft, Gleichgewicht und Koordination.",
		HowTo:   []string{"Auf einem Bein stehen.", "Seitlich auf das andere Bein hüpfen.", "Das nachziehende Bein hinten ausschwingen lassen.", "Weich landen und umkehren."},
		Correct: []string{"Landungen bleiben weich.", "Knie in Richtung des Fußes.", "Hüften bleiben im Gleichgewicht.", "Bewegung bleibt weich."},
		Wrong:   []string{"Das Knie nach innen fallen lassen.", "Hart landen.", "Das Gleichgewicht verlieren.", "Ohne Kontrolle hetzen."},
	},
	"J07": {
		Desc:    "Eine Cardio-Bewegung mit Laufen auf der Stelle und hohem Kniehub.",
		HowTo:   []string{"Aufrecht stehen.", "Auf der Stelle laufen und die Knie bis auf Hüfthöhe heben.", "Auf den Fußballen bleiben.", "Die Arme im Rhythmus mitpumpen."},
		Correct: []string{"Oberkörper bleibt aufrecht.", "Knie heben hoch.", "Landungen bleiben leicht.", "Tempo bleibt gleichmäßig."},
		Wrong:   []string{"Nach hinten lehnen.", "Mit den Füßen stampfen.", "Zusammensacken.", "Die Luft anhalten."},
	},
	"J08": {
		Desc:    "Eine Ganzkörper-Konditionsübung, die Kniebeuge, Plank und Sprung verbindet.",
		HowTo:   []string{"Aus dem Stand in die Hocke gehen und die Hände aufsetzen.", "Die Füße in die Plank springen oder zurücksetzen.", "Die Füße wieder unter die Hüften bringen.", "Aufstehen und mit einem Strecken nach oben springen."},
		Correct: []string{"Wirbelsäule bleibt neutral in der Plank.", "Landungen bleiben weich.", "Tempo bleibt gleichmäßig.", "Rumpf bleibt aktiv."},
		Wrong:   []string{"Das Becken in der Plank durchhängen lassen.", "Hart landen.", "Die Form hetzen.", "Die Luft anhalten."},
	},
	"J09": {
		Desc:    "Ein explosiver Ausfallschritt, der die Beine in der Luft wechselt, für Kraft und Koordination.",
		HowTo:   []string{"In einer Ausfallschritt-Position beginnen.", "Nach oben springen und die Beine in der Luft wechseln.", "Weich in einem Ausfallschritt auf der anderen Seite landen.", "Im Wechsel weitermachen."},
		Correct: []string{"Landungen bleiben weich.", "Vorderes Knie in Richtung des Fußes.", "Oberkörper bleibt aufrecht.", "Hüften federn jede Landung ab."},
		Wrong:   []string{"Hart landen.", "Das vordere Knie nach innen fallen lassen.", "Weit nach vorn lehnen.", "Die Kontrolle beim Wechsel verlieren."},
	},
	"CD07": {
		Desc:    "Eine ruhige Dehnung für Rücken, Hüften und Schultern zum Herunterkommen.",
		HowTo:   []string{"Knien und das Becken Richtung Fersen sinken lassen.", "Die Arme nach vorn auf den Boden strecken.", "Die Stirn ablegen.", "Langsam atmen und entspannen."},
		Correct: []string{"Becken sinkt Richtung Fersen.", "Rücken wird sanft länger.", "Schultern entspannen.", "Atmung bleibt langsam."},
		Wrong:   []string{"Das Becken nach unten zwingen.", "Die Schultern anspannen.", "Die Luft anhalten.", "Die Knie überlasten."},
	},
}
