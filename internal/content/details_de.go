package content

// detailsDE is the German rich content for every exercise (current library IDs).
var detailsDE = map[string]Detail{
	// ---- Aufwärmen ----
	"W01": {
		Desc:    "Ein einfaches Schulter-Aufwärmen für Nacken, Schultern und oberen Rücken.",
		HowTo:   []string{"Aufrecht stehen, Arme locker.", "Die Schultern langsam rückwärts kreisen.", "Dann langsam vorwärts.", "Die Bewegung bleibt weich und kontrolliert."},
		Correct: []string{"Halte den Nacken entspannt.", "Bewege die Schultern weich.", "Lass die Arme locker.", "Bleib aufrecht."},
		Wrong:   []string{"Zuck nicht stark mit den Schultern.", "Bewege dich nicht zu schnell.", "Spann den Nacken nicht an.", "Mach den oberen Rücken nicht rund."},
	},
	"W02": {
		Desc:    "Ein Schulter- und Oberrücken-Aufwärmen, das Deltas und Schulterstabilisatoren aktiviert.",
		HowTo:   []string{"Aufrecht stehen, Arme auf Schulterhöhe heben.", "Kleine Kreise vorwärts.", "Dann kleine Kreise rückwärts.", "Die Kreise klein und kontrolliert halten."},
		Correct: []string{"Halte die Arme auf Schulterhöhe.", "Mach die Kreise klein.", "Halte die Schultern unten.", "Halte den Nacken entspannt."},
		Wrong:   []string{"Mach keine großen Schwünge.", "Zieh die Schultern nicht zu den Ohren.", "Überstreck nicht den unteren Rücken.", "Halte nicht die Luft an."},
	},
	"W03": {
		Desc:    "Eine sanfte Rotationsübung für Rumpf und obere Wirbelsäule.",
		HowTo:   []string{"Füße hüftbreit.", "Das Becken zeigt überwiegend nach vorn.", "Den Oberkörper zu einer Seite drehen.", "Zur Mitte zurück und zur anderen Seite drehen."},
		Correct: []string{"Halte die Füße stabil.", "Dreh dich aus dem Rumpf.", "Dreh dich weich.", "Halte die Knie ruhig."},
		Wrong:   []string{"Dreh die Knie nicht mit.", "Dreh dich nicht zu schnell.", "Lehn dich nicht nach hinten.", "Erzwing den Bewegungsumfang nicht."},
	},
	"W04": {
		Desc:    "Eine Bewegungsübung für Hüften, Beinrückseite und Kontrolle des unteren Rückens.",
		HowTo:   []string{"Aufrecht stehen, Füße hüftbreit.", "Die Knie leicht weich machen.", "Das Becken nach hinten schieben.", "Den Oberkörper mit geradem Rücken senken.", "Über die Hüften zurück nach oben."},
		Correct: []string{"Halte den Rücken gerade.", "Starte die Bewegung aus der Hüfte.", "Halte die Knie weich, nicht tief gebeugt.", "Lass den Nacken der Wirbelsäule folgen."},
		Wrong:   []string{"Mach den Rücken nicht rund.", "Mach daraus keine Kniebeuge.", "Greif nicht zu tief.", "Heb den Kopf nicht zu hoch."},
	},
	"W05": {
		Desc:    "Ein kontrolliertes Bein-Aufwärmen in sicherem, flachem Bereich.",
		HowTo:   []string{"Füße etwas breiter als hüftbreit.", "Das Becken leicht nach hinten schieben.", "Die Knie nur bis zu einer angenehmen Tiefe beugen.", "Langsam aufrichten."},
		Correct: []string{"Halte die Knie über den Zehen.", "Halte die Fersen am Boden.", "Mach den Rücken lang.", "Halte die Tiefe moderat."},
		Wrong:   []string{"Geh nicht zu tief.", "Lass die Knie nicht nach innen fallen.", "Heb die Fersen nicht.", "Fall nicht schnell nach unten."},
	},
	"W06": {
		Desc:    "Ein Aufwärmen für Schultern, Handgelenke und Rumpf.",
		HowTo:   []string{"In einer hohen Plank beginnen.", "Hände unter den Schultern.", "Den Körper leicht nach vorn verlagern.", "Zurück in die Ausgangsposition.", "Die Bewegung klein und kontrolliert."},
		Correct: []string{"Halte den Körper in einer Linie.", "Halte die Schultern stabil.", "Spann den Rumpf an.", "Bewege dich weich."},
		Wrong:   []string{"Senk das Becken nicht ab.", "Schieb die Schultern nicht zu den Ohren.", "Geh nicht zu weit nach vorn.", "Lass den unteren Rücken nicht durchhängen."},
	},

	// ---- Bauch / Rumpf ----
	"C01": {
		Desc:    "Eine Rumpfstabilitätsübung für Bauch, Schultern, Gesäß und Rumpf.",
		HowTo:   []string{"Die Ellbogen unter die Schultern.", "Die Füße nach hinten setzen.", "Eine gerade Linie von Schultern zu Fersen.", "Bauch und Gesäß anspannen.", "Die Position halten."},
		Correct: []string{"Halte den Körper in einer Linie.", "Halte die Ellbogen unter den Schultern.", "Spann den Rumpf an.", "Atme gleichmäßig weiter."},
		Wrong:   []string{"Lass den unteren Rücken nicht durchhängen.", "Heb das Becken nicht zu hoch.", "Schau nicht nach vorn.", "Halte nicht die Luft an."},
	},
	"C02": {
		Desc:    "Eine kurze, intensive Plank mit Fokus auf Ganzkörperspannung.",
		HowTo:   []string{"In einer Unterarm-Plank beginnen.", "Den Bauch stark anspannen.", "Das Gesäß anspannen.", "Die Ellbogen gedanklich zu den Zehen ziehen.", "Mit maximaler Kontrolle halten."},
		Correct: []string{"Spann den ganzen Körper an.", "Spann das Gesäß an.", "Halte den Rumpf fest.", "Atme kontrolliert weiter."},
		Wrong:   []string{"Halte nicht wie eine lockere Plank.", "Lass das Becken nicht durchhängen.", "Halte nicht die Luft an.", "Halte nicht zu lange."},
	},
	"C03": {
		Desc:    "Eine Unterbauch-Übung mit kontrollierten kreuzenden Beinbewegungen.",
		HowTo:   []string{"Auf den Rücken legen.", "Den unteren Rücken sanft zum Boden drücken.", "Die Beine heben.", "Ein Bein über das andere kreuzen.", "Das Kreuzen abwechselnd fortsetzen."},
		Correct: []string{"Halte den unteren Rücken nah am Boden.", "Bewege die Beine kontrolliert.", "Halte den Nacken entspannt.", "Bewege dich weich."},
		Wrong:   []string{"Überstreck nicht den unteren Rücken.", "Senk die Beine nicht zu tief.", "Zieh den Kopf nicht nach vorn.", "Bewege dich nicht zu schnell."},
	},
	"C04": {
		Desc:    "Eine Unterbauch-Übung mit kurzen, wechselnden Beinschlägen.",
		HowTo:   []string{"Auf den Rücken legen.", "Den unteren Rücken zum Boden drücken.", "Die Beine heben.", "Die Beine in kurzen, wechselnden Schlägen auf und ab bewegen."},
		Correct: []string{"Halte den Bewegungsumfang klein.", "Halte den unteren Rücken stabil.", "Lass die Beine nicht zu tief sinken.", "Atme gleichmäßig weiter."},
		Wrong:   []string{"Überstreck nicht den unteren Rücken.", "Mach keine großen Schläge.", "Spann den Nacken nicht an.", "Halte nicht die Luft an."},
	},
	"C05": {
		Desc:    "Eine Bauchübung mit Fokus auf kontrolliertes Heben des Beckens.",
		HowTo:   []string{"Auf den Rücken legen.", "Knie beugen, Füße heben.", "Mit dem Bauch das Becken leicht nach oben rollen.", "Langsam absenken."},
		Correct: []string{"Hol die Bewegung aus dem Unterbauch.", "Heb das Becken kontrolliert.", "Halte den Nacken entspannt.", "Senk dich langsam ab."},
		Wrong:   []string{"Schwing die Beine nicht.", "Nimm keinen Schwung.", "Zieh nicht am Nacken.", "Lass das Becken nicht schnell fallen."},
	},
	"C06": {
		Desc:    "Ein statisches Rumpf-Halten mit verkürztem Hebel für bessere Kontrolle.",
		HowTo:   []string{"Auf den Rücken legen.", "Den unteren Rücken zum Boden drücken.", "Die Schultern leicht heben.", "Knie gebeugt oder Beine höher.", "Die Position halten."},
		Correct: []string{"Halte den unteren Rücken am Boden.", "Halte den Bauch fest.", "Halte den Nacken entspannt.", "Atme weiter."},
		Wrong:   []string{"Überstreck nicht den unteren Rücken.", "Senk die Beine nicht zu tief.", "Zieh das Kinn nicht nach vorn.", "Halte nicht die Luft an."},
	},
	"C07": {
		Desc:    "Eine seitliche Rumpfübung für die schrägen Bauchmuskeln, Schulter- und Hüftstabilität.",
		HowTo:   []string{"Auf eine Seite legen.", "Den Ellbogen unter die Schulter.", "Das Becken heben.", "Der Körper in einer geraden Linie.", "Halten."},
		Correct: []string{"Halte den Ellbogen unter der Schulter.", "Halte das Becken oben.", "Mach den Körper lang.", "Halte den Nacken neutral."},
		Wrong:   []string{"Senk das Becken nicht ab.", "Zieh die Schulter nicht zum Ohr.", "Roll die Brust nicht nach vorn.", "Halte nicht die Luft an."},
	},
	"C08": {
		Desc:    "Eine stärkere Seitstütz-Variante für die schrägen Bauchmuskeln und die seitliche Hüfte.",
		HowTo:   []string{"In einem Seitstütz beginnen.", "Das Becken leicht senken.", "Das Becken wieder heben.", "Mit Kontrolle wiederholen.", "Seite wechseln."},
		Correct: []string{"Bewege dich klein und kontrolliert.", "Halte den Ellbogen unter der Schulter.", "Bewege das Becken nach oben und unten.", "Halte den Rumpf fest."},
		Wrong:   []string{"Sink nicht zu tief ab.", "Verdreh den Rumpf nicht.", "Sink nicht in die Schulter ein.", "Bewege dich nicht zu schnell."},
	},
	"C10": {
		Desc:    "Eine Rumpf- und Trizeps-Übung im Wechsel zwischen Unterarm- und hoher Plank.",
		HowTo:   []string{"In einer Unterarm-Plank beginnen.", "Eine Hand auf den Boden setzen und hochdrücken.", "Die andere Hand in die hohe Plank bringen.", "Zurück auf die Unterarme.", "Die führende Hand abwechseln."},
		Correct: []string{"Halte das Becken stabil.", "Spann den Rumpf an.", "Bewege dich kontrolliert.", "Halte die Schultern über Händen oder Ellbogen."},
		Wrong:   []string{"Lass das Becken nicht seitlich schaukeln.", "Lass den unteren Rücken nicht durchhängen.", "Bewege dich nicht zu schnell.", "Lass die Schultern nicht einsinken."},
	},
	"C11": {
		Desc:    "Eine kontrollierte Bauchübung und eine sicherere Alternative zu schwereren Beinheben.",
		HowTo:   []string{"Auf den Rücken legen.", "Die Beine in die Tabletop-Position heben.", "Eine Ferse zum Boden senken.", "Leicht antippen und zurück.", "Seiten wechseln."},
		Correct: []string{"Halte den unteren Rücken stabil.", "Bewege dich langsam.", "Setz die Ferse sanft auf.", "Halte das Becken ruhig."},
		Wrong:   []string{"Lass das Bein nicht zu schnell fallen.", "Überstreck nicht den unteren Rücken.", "Bewege nicht beide Beine gleichzeitig.", "Hetz nicht."},
	},
	"C12": {
		Desc:    "Eine Rumpfkontroll-Übung für die tiefe Bauchmuskulatur und Wirbelsäulenstabilität.",
		HowTo:   []string{"Auf den Rücken legen.", "Arme und Beine in die Tabletop-Position heben.", "Den gegenüberliegenden Arm und das Bein ausstrecken.", "Den unteren Rücken stabil halten.", "Zurück und Seite wechseln."},
		Correct: []string{"Halte den unteren Rücken nah am Boden.", "Bewege dich langsam.", "Bewege gegenüberliegenden Arm und Bein zusammen.", "Halte den Rumpf ruhig."},
		Wrong:   []string{"Überstreck nicht den unteren Rücken.", "Bewege dich nicht zu schnell.", "Senk das Bein nicht zu tief.", "Drück die Rippen nicht heraus."},
	},

	// ---- Liegestütze / Trizeps / Schultern ----
	"P01": {
		Desc:    "Eine grundlegende Oberkörperübung für Brust, Trizeps, Schultern und Rumpf.",
		HowTo:   []string{"In einer hohen Plank beginnen.", "Hände unter den Schultern oder etwas breiter.", "Den Körper kontrolliert senken.", "Hochdrücken und den Körper gerade halten."},
		Correct: []string{"Halte den Körper in einer Linie.", "Halte den Rumpf fest.", "Halte die Ellbogen etwa 30 bis 45 Grad.", "Bewege dich kontrolliert."},
		Wrong:   []string{"Lass das Becken nicht durchhängen.", "Heb das Becken nicht zu hoch.", "Spreiz die Ellbogen nicht gerade ab.", "Fall nicht in die untere Position."},
	},
	"P02": {
		Desc:    "Eine Liegestütz-Variante mit mehr Fokus auf den Trizeps.",
		HowTo:   []string{"In einer hohen Plank beginnen.", "Die Hände enger als bei normalen Liegestützen setzen.", "Die Ellbogen nah am Körper halten.", "Langsam senken und hochdrücken."},
		Correct: []string{"Halte die Ellbogen nah am Körper.", "Halte den Körper gerade.", "Halte die Schultern weg von den Ohren.", "Halte die Handgelenke angenehm."},
		Wrong:   []string{"Setz die Hände nicht zu eng.", "Spreiz die Ellbogen nicht ab.", "Lass den unteren Rücken nicht durchhängen.", "Mach bei Handgelenk- oder Ellbogenschmerz nicht weiter."},
	},
	"P03": {
		Desc:    "Kontrollierte Liegestütze mit einer kurzen Pause in der unteren Position.",
		HowTo:   []string{"In normaler Liegestützposition beginnen.", "Langsam senken.", "Unten eine Sekunde pausieren.", "Hochdrücken, ohne die Körperposition zu verlieren."},
		Correct: []string{"Pausiere kontrolliert.", "Halte den Rumpf fest.", "Halte die Brust aktiv.", "Halte den Körper gerade."},
		Wrong:   []string{"Entspann unten nicht.", "Lass das Becken nicht durchhängen.", "Halte nicht die Luft an.", "Spring nicht aus der unteren Position."},
	},
	"P04": {
		Desc:    "Eine langsame Liegestütz-Variante, die Kontrolle und Spannungsdauer erhöht.",
		HowTo:   []string{"In einer hohen Plank beginnen.", "Etwa drei Sekunden senken.", "Unten eine Sekunde pausieren.", "Etwa drei Sekunden hochdrücken."},
		Correct: []string{"Halte ein langsames, gleichmäßiges Tempo.", "Halte den Körper gerade.", "Beuge die Ellbogen kontrolliert.", "Atme gleichmäßig weiter."},
		Wrong:   []string{"Hetz die Bewegung nicht.", "Sack unten nicht zusammen.", "Lass das Becken nicht durchhängen.", "Verlier die Rumpfspannung nicht."},
	},
	"P05": {
		Desc:    "Eine asymmetrische Liegestütz-Variante, die Brust, Schultern und Rumpf fordert.",
		HowTo:   []string{"In Liegestützposition beginnen.", "Eine Hand etwas vor, die andere etwas zurück setzen.", "Kontrolliert senken.", "Hochdrücken.", "Auf der anderen Seite die Handposition wechseln."},
		Correct: []string{"Halte den Körper gerade.", "Halte den Rumpf fest.", "Halte beide Schultern stabil.", "Bewege dich kontrolliert."},
		Wrong:   []string{"Verdreh den Rumpf nicht.", "Senk keine Schulter ab.", "Setz die Hände nicht zu breit.", "Erzwing den Umfang bei Schulterbeschwerden nicht."},
	},
	"P06": {
		Desc:    "Eine schulterbetonte Liegestütz-Variante.",
		HowTo:   []string{"In einer hohen Plank beginnen.", "Die Füße etwas näher setzen und das Becken heben.", "Den Kopf zwischen den Armen.", "Den Kopf Richtung Boden senken.", "Wieder hochdrücken."},
		Correct: []string{"Halte das Becken hoch.", "Beuge die Ellbogen kontrolliert.", "Halte die Schultern stabil.", "Halte den Nacken neutral."},
		Wrong:   []string{"Sink nicht in eine normale Liegestütze.", "Führ den Kopf nicht zu weit nach vorn.", "Spreiz die Ellbogen nicht zu stark ab.", "Mach bei Schulterschmerz nicht weiter."},
	},
	"P08": {
		Desc:    "Eine Trizeps- und Schulterübung aus dem Unterarmstütz.",
		HowTo:   []string{"Auf den Unterarmen beginnen, der Körper lang.", "Ellbogen unter oder leicht vor den Schultern.", "Über Hände und Unterarme drücken.", "Die Ellbogen leicht strecken.", "Kontrolliert zurück."},
		Correct: []string{"Halte den Rumpf fest.", "Bewege die Ellbogen weich.", "Halte die Schultern unten.", "Halte den Bewegungsumfang angenehm."},
		Wrong:   []string{"Erzwing die Ellbogen nicht.", "Lass das Becken nicht durchhängen.", "Zieh die Schultern nicht hoch.", "Bewege dich nicht bei Ellbogen- oder Schulterschmerz."},
	},

	// ---- Rücken / Haltung ----
	"B01": {
		Desc:    "Eine haltungsorientierte Rückenübung für oberen Rücken, Schulterblätter und Rückenstrecker.",
		HowTo:   []string{"Auf den Bauch legen.", "Die Arme am Körper.", "Die Daumen leicht nach außen drehen.", "Die Brust leicht heben.", "Die Schulterblätter nach hinten und unten ziehen.", "Halten."},
		Correct: []string{"Heb nur ein kleines Stück.", "Halte den Nacken neutral.", "Zieh die Schulterblätter nach hinten und unten.", "Spann das Gesäß leicht an."},
		Wrong:   []string{"Heb nicht zu hoch.", "Wirf den Kopf nicht nach hinten.", "Überstreck nicht den unteren Rücken.", "Zieh die Schultern nicht zu den Ohren."},
	},
	"B02": {
		Desc:    "Eine Übung für oberen Rücken und Schulterblätter in Bauchlage.",
		HowTo:   []string{"Auf den Bauch legen.", "Der Kopf schaut nach unten.", "Die Arme in einem weiten Bogen von den Seiten nach oben führen.", "Auf demselben Weg zurück.", "Bewegung kontrolliert."},
		Correct: []string{"Halte den Nacken entspannt.", "Bewege die Schulterblätter weich.", "Heb die Brust nur leicht.", "Bewege die Arme ohne Hast."},
		Wrong:   []string{"Heb den Kopf nicht.", "Überstreck den unteren Rücken nicht zu stark.", "Schwing die Arme nicht.", "Zieh die Schultern nicht zu den Ohren."},
	},
	"B03": {
		Desc:    "Eine Oberrücken-Übung für Schulterblätter und hintere Schultern.",
		HowTo:   []string{"Auf den Bauch legen.", "Die Ellbogen in eine W-Form beugen.", "Ellbogen und Hände leicht heben.", "Die Schulterblätter zusammenziehen.", "Kontrolliert absenken."},
		Correct: []string{"Halte den Bewegungsumfang klein.", "Arbeite aus den Schulterblättern.", "Halte den Nacken neutral.", "Halte die Schultern weg von den Ohren."},
		Wrong:   []string{"Heb nicht zu hoch.", "Heb den Kopf nicht.", "Nimm keinen Schwung.", "Überstreck nicht den unteren Rücken."},
	},
	"B04": {
		Desc:    "Eine Schulterblatt-Übung für den unteren Trapez und oberen Rücken.",
		HowTo:   []string{"Auf den Bauch legen.", "Die Arme diagonal über Kopf in eine Y-Form strecken.", "Die Arme leicht vom Boden heben.", "Kontrolliert absenken."},
		Correct: []string{"Heb die Arme nur wenig.", "Halte den Nacken neutral.", "Halte die Schultern unten.", "Bewege dich kontrolliert."},
		Wrong:   []string{"Heb nicht zu hoch.", "Zieh die Schultern nicht zu den Ohren.", "Überstreck nicht den unteren Rücken.", "Bewege dich nicht zu schnell."},
	},
	"B05": {
		Desc:    "Eine Oberrücken-Übung für hintere Schultern und Schulterblattkontrolle.",
		HowTo:   []string{"Auf den Bauch legen.", "Die Arme in einer T-Form zur Seite strecken.", "Die Arme leicht heben.", "Die Schulterblätter zusammenziehen.", "Langsam absenken."},
		Correct: []string{"Halte die Arme auf Schulterlinie.", "Halte den Nacken entspannt.", "Bewege die Schulterblätter sanft.", "Heb nur ein kleines Stück."},
		Wrong:   []string{"Wirf die Arme nicht nach oben.", "Heb den Kopf nicht.", "Zieh die Schultern nicht hoch.", "Erzwing keinen großen Umfang."},
	},
	"B06": {
		Desc:    "Eine Rückenübung, die eine Zugbewegung ohne Geräte imitiert.",
		HowTo:   []string{"Auf den Bauch legen, Arme über Kopf.", "Die Brust leicht heben.", "Die Ellbogen Richtung Rippen ziehen.", "Die Schulterblätter zusammenziehen.", "Die Arme wieder nach vorn strecken."},
		Correct: []string{"Zieh die Ellbogen kontrolliert nach unten.", "Zieh die Schulterblätter sanft zusammen.", "Halte den Nacken neutral.", "Halte den unteren Rücken angenehm."},
		Wrong:   []string{"Reiß nicht mit den Armen.", "Heb nicht zu hoch.", "Wirf den Kopf nicht nach hinten.", "Heb die Schultern nicht zu den Ohren."},
	},
	"B07": {
		Desc:    "Eine Rücken- und Koordinationsübung in Bauchlage.",
		HowTo:   []string{"Auf den Bauch legen.", "Die Brust leicht heben.", "Gegenüberliegenden Arm und Bein in einem kontrollierten Schwimmmuster bewegen.", "Die Bewegung klein und gleichmäßig."},
		Correct: []string{"Halte den Nacken neutral.", "Bewege dich kontrolliert.", "Halte den unteren Rücken angenehm.", "Atme weiter."},
		Wrong:   []string{"Bewege dich nicht zu schnell.", "Heb nicht zu hoch.", "Wirf den Kopf nicht nach hinten.", "Überlaste den unteren Rücken nicht."},
	},
	"B08": {
		Desc:    "Eine Übung für die hintere Kette: Gesäß, Beinrückseite, Schultern und Rücken.",
		HowTo:   []string{"Mit gestreckten Beinen auf den Boden setzen.", "Die Hände hinter dem Becken.", "Über Hände und Fersen drücken.", "Das Becken heben.", "Den Körper in einer langen Linie halten."},
		Correct: []string{"Halte die Brust offen.", "Halte das Becken oben.", "Halte die Schultern stabil.", "Halte den Nacken neutral."},
		Wrong:   []string{"Senk das Becken nicht ab.", "Zieh die Schultern nicht hoch.", "Wirf den Kopf nicht nach hinten.", "Überlaste die Handgelenke nicht."},
	},
	"B09": {
		Desc:    "Eine Rückenübung mit kleinem Umfang für Rückenstrecker und Haltung.",
		HowTo:   []string{"Auf den Bauch legen.", "Die Hände neben den Körper oder leicht zur Brust.", "Die Brust leicht heben.", "Kleine, kontrollierte Impulse machen.", "Nacken neutral."},
		Correct: []string{"Heb nur ein kleines Stück.", "Mach weiche Impulse.", "Mach den Nacken lang.", "Halte den unteren Rücken kontrolliert."},
		Wrong:   []string{"Pulsier nicht zu hoch.", "Nimm keinen Schwung.", "Schau nicht nach vorn.", "Ignorier keinen stechenden Schmerz im unteren Rücken."},
	},
	"B10": {
		Desc:    "Eine stärkere Rückenübung, die ein kleines Kobra-Heben mit einem W-förmigen Armzug verbindet.",
		HowTo:   []string{"Auf den Bauch legen, Arme nach vorn.", "Die Brust leicht heben.", "Die Ellbogen nach unten und hinten in eine W-Form ziehen.", "Die Schulterblätter zusammenziehen.", "Wieder nach vorn strecken und kontrolliert absenken."},
		Correct: []string{"Heb die Brust nur leicht.", "Zieh weich in die W-Form.", "Zieh die Schulterblätter nach hinten und unten.", "Mach den Nacken lang."},
		Wrong:   []string{"Überstreck nicht den unteren Rücken.", "Zieh nicht mit dem Nacken.", "Bewege dich nicht zu schnell.", "Zieh die Schultern nicht hoch."},
	},

	// ---- Arme / Bizeps ----
	"A01": {
		Desc:    "Eine Bizeps-Übung ohne Geräte: eine Hand leistet der anderen Widerstand.",
		HowTo:   []string{"Aufrecht stehen oder sitzen.", "Einen Ellbogen wie bei einem Curl beugen.", "Mit der anderen Hand auf den Unterarm drücken.", "Langsam gegen den Widerstand nach oben.", "Seite wechseln."},
		Correct: []string{"Halte den Widerstand gleichmäßig.", "Bewege dich langsam.", "Halte den Ellbogen nah am Körper.", "Halte die Schulter unten."},
		Wrong:   []string{"Bewege dich nicht ohne echten Widerstand.", "Reiß nicht mit dem Arm.", "Heb die Schulter nicht.", "Halte nicht die Luft an."},
	},
	"A02": {
		Desc:    "Ein statisches Bizeps-Halten mit Eigenwiderstand.",
		HowTo:   []string{"Den Ellbogen auf etwa 90 Grad beugen.", "Mit der anderen Hand nach unten drücken.", "Der arbeitende Arm leistet ohne Bewegung Widerstand.", "Die Spannung halten.", "Seite wechseln."},
		Correct: []string{"Halte den Ellbogenwinkel stabil.", "Halte die Spannung gleichmäßig.", "Halte die Schultern entspannt.", "Atme weiter."},
		Wrong:   []string{"Drück nicht in kurzen Stößen.", "Verdreh den Rumpf nicht.", "Heb die Schulter nicht.", "Halte nicht die Luft an."},
	},
	"A03": {
		Desc:    "Eine isometrische Arm- und Unterarmübung mit den Händen gegeneinander.",
		HowTo:   []string{"Die Finger oder Hände einhaken.", "Die Ellbogen leicht gebeugt.", "Die Hände voneinander wegziehen.", "Gleichmäßige Spannung halten.", "Langsam lösen."},
		Correct: []string{"Halte die Spannung kontrolliert.", "Halte die Schultern unten.", "Halte die Handgelenke angenehm.", "Atme ruhig weiter."},
		Wrong:   []string{"Zieh nicht mit einem Ruck.", "Zieh die Schultern nicht hoch.", "Überbeug die Handgelenke nicht.", "Halte nicht die Luft an."},
	},
	"A04": {
		Desc:    "Eine langsame Absenkübung für den Bizeps mit Eigenwiderstand.",
		HowTo:   []string{"Mit gebeugtem Ellbogen beginnen.", "Mit der anderen Hand Widerstand erzeugen.", "Den Unterarm langsam absenken.", "Während des Absenkens Widerstand halten.", "Seite wechseln."},
		Correct: []string{"Senk langsam ab.", "Halte den Widerstand konstant.", "Halte den Ellbogen nah am Körper.", "Halte die Schulter entspannt."},
		Wrong:   []string{"Lass den Arm nicht schnell fallen.", "Bewege dich nicht ohne Widerstand.", "Verdreh den Rumpf nicht.", "Halte nicht die Luft an."},
	},

	// ---- Beine / Gesäß ----
	"L01": {
		Desc:    "Eine kontrollierte Beinübung für Oberschenkel, Hüften und Haltung.",
		HowTo:   []string{"Füße etwas breiter als hüftbreit.", "Das Becken nach hinten schieben.", "Die Knie bis zu einer angenehmen, flachen Tiefe beugen.", "Langsam aufrichten."},
		Correct: []string{"Halte die Knie über den Zehen.", "Halte die Fersen unten.", "Mach den Rücken lang.", "Halte die Tiefe angenehm."},
		Wrong:   []string{"Geh nicht zu tief.", "Lass die Knie nicht nach innen fallen.", "Heb die Fersen nicht.", "Fall nicht schnell."},
	},
	"L02": {
		Desc:    "Eine kontrollierte Kniebeuge-Variante mit kurzem Halten im sicheren Bereich.",
		HowTo:   []string{"Stabil stehen.", "In eine angenehme, flache Kniebeuge senken.", "Kurz pausieren.", "Kontrolliert aufrichten."},
		Correct: []string{"Halte die Pause stabil.", "Halte die Knie über den Zehen.", "Halte die Brust offen.", "Halte die Fersen am Boden."},
		Wrong:   []string{"Geh nicht zu tief.", "Feder nicht in der unteren Position.", "Lass die Knie nicht nach innen fallen.", "Halte nicht die Luft an."},
	},
	"L03": {
		Desc:    "Eine Bewegungsübung für Hüften, Beinrückseite und Kontrolle des unteren Rückens.",
		HowTo:   []string{"Aufrecht stehen, Füße hüftbreit.", "Die Knie leicht weich machen.", "Das Becken nach hinten schieben.", "Den Oberkörper mit geradem Rücken senken.", "Über die Hüften zurück nach oben."},
		Correct: []string{"Halte den Rücken gerade.", "Starte die Bewegung aus der Hüfte.", "Halte die Knie weich, nicht tief gebeugt.", "Lass den Nacken der Wirbelsäule folgen."},
		Wrong:   []string{"Mach den Rücken nicht rund.", "Mach daraus keine Kniebeuge.", "Greif nicht zu tief.", "Heb den Kopf nicht zu hoch."},
	},
	"L04": {
		Desc:    "Eine stärkere Gesäßübung mit jeweils einem Bein.",
		HowTo:   []string{"Auf den Rücken legen.", "Ein Knie beugen, der Fuß bleibt am Boden.", "Das andere Bein strecken oder heben.", "Über die arbeitende Ferse drücken.", "Das Becken heben und langsam senken."},
		Correct: []string{"Heb aus dem arbeitenden Gesäß.", "Halte das Becken gerade.", "Bewege dich kontrolliert.", "Halte den unteren Rücken neutral."},
		Wrong:   []string{"Verdreh das Becken nicht.", "Drück nicht über die Zehen.", "Überstreck nicht den unteren Rücken.", "Fall nicht zu schnell."},
	},
	"L05": {
		Desc:    "Eine Gesäß- und Beckenstabilitätsübung.",
		HowTo:   []string{"In einer Gesäßbrücke beginnen.", "Das Becken oben halten.", "Einen Fuß leicht vom Boden heben.", "Absenken und Seite wechseln.", "Das Becken stabil halten."},
		Correct: []string{"Halte das Becken gerade.", "Spann das Gesäß an.", "Bewege dich langsam.", "Halte den unteren Rücken angenehm."},
		Wrong:   []string{"Lass das Becken nicht absinken.", "Schaukel nicht seitlich.", "Bewege dich nicht zu schnell.", "Überstreck nicht den unteren Rücken."},
	},
	"L06": {
		Desc:    "Eine Gesäßübung, die Hüften und hintere Kette stärkt.",
		HowTo:   []string{"Auf den Rücken legen.", "Knie beugen, Füße am Boden.", "Das Becken heben.", "Oben pausieren.", "Langsam absenken."},
		Correct: []string{"Drück über die Fersen.", "Spann oben das Gesäß an.", "Halte die Rippen unten.", "Halte den unteren Rücken neutral."},
		Wrong:   []string{"Überstreck oben nicht den unteren Rücken.", "Heb die Fersen nicht.", "Lass das Becken nicht zu schnell fallen.", "Setz die Füße nicht zu nah."},
	},
	"L07": {
		Desc:    "Eine Unterschenkelübung für Waden, Sprunggelenke und Fußkontrolle.",
		HowTo:   []string{"Aufrecht stehen, Füße hüftbreit.", "Auf die Fußballen heben.", "Oben kurz pausieren.", "Die Fersen langsam senken."},
		Correct: []string{"Bleib aufrecht.", "Bewege dich weich.", "Halte die Sprunggelenke ausgerichtet.", "Senk kontrolliert ab."},
		Wrong:   []string{"Lass die Fersen nicht schnell fallen.", "Schaukel nicht nach vorn.", "Kipp die Sprunggelenke nicht nach außen.", "Nimm keinen Schwung."},
	},
	"L09": {
		Desc:    "Eine Übung für die seitliche Kette: seitliches Gesäß, schräge Bauchmuskeln und Hüftstabilität.",
		HowTo:   []string{"In einem Seitstütz beginnen.", "Der Körper in einer geraden Linie.", "Das obere Bein leicht heben.", "Kontrolliert absenken.", "Wiederholen und Seite wechseln."},
		Correct: []string{"Halte das Becken oben.", "Bewege das obere Bein langsam.", "Halte den Rumpf stabil.", "Halte die Schulter stark."},
		Wrong:   []string{"Senk das Becken nicht ab.", "Schwing das Bein nicht.", "Roll den Körper nicht nach hinten.", "Sink nicht in die Schulter ein."},
	},

	// ---- Abwärmen ----
	"CD01": {
		Desc:    "Eine sanfte Dehnung für Brust und vordere Schultern.",
		HowTo:   []string{"Auf den Rücken legen.", "Die Arme zur Seite öffnen.", "Die Schultern entspannen.", "Die Brust sanft öffnen lassen.", "Langsam atmen."},
		Correct: []string{"Lass die Brust sanft öffnen.", "Halte die Schultern entspannt.", "Halte den unteren Rücken angenehm.", "Halte die Schultern schmerzfrei."},
		Wrong:   []string{"Zwing die Arme nicht nach unten.", "Überstreck nicht den unteren Rücken.", "Spann den Nacken nicht an.", "Dehn nicht durch Schmerz."},
	},
	"CD02": {
		Desc:    "Eine sanfte Dehnung für die Rückseite des Oberschenkels.",
		HowTo:   []string{"Auf den Rücken legen.", "Ein Bein heben.", "Hinter Oberschenkel oder Wade fassen.", "Das Bein sanft zu sich ziehen.", "Seite wechseln."},
		Correct: []string{"Halte die Dehnung mild.", "Lass das Knie leicht gebeugt.", "Halte den unteren Rücken ruhig.", "Atme langsam."},
		Wrong:   []string{"Zieh nicht zu stark.", "Streck das Bein nicht gewaltsam.", "Heb das Becken nicht.", "Ignorier keinen stechenden Schmerz hinter dem Knie."},
	},
	"CD03": {
		Desc:    "Eine sanfte Rotationsdehnung für Wirbelsäule und Rumpf.",
		HowTo:   []string{"Auf den Rücken legen.", "Die Knie beugen.", "Die Knie sanft zu einer Seite sinken lassen.", "Die Schultern am Boden halten.", "Zur Mitte zurück und Seite wechseln."},
		Correct: []string{"Bewege dich langsam.", "Halte die Schultern am Boden.", "Atme entspannt.", "Halte die Dehnung angenehm."},
		Wrong:   []string{"Zwing die Knie nicht zum Boden.", "Dreh dich nicht schnell.", "Heb die Schulter nicht.", "Dehn nicht durch Schmerz."},
	},
	"CD04": {
		Desc:    "Eine sanfte Dehnung der Körpervorderseite und eine milde Rückenstreckung.",
		HowTo:   []string{"Auf den Bauch legen.", "Die Unterarme auf den Boden.", "Die Brust sanft heben.", "Den Nacken lang halten.", "Langsam atmen."},
		Correct: []string{"Heb nur sanft.", "Halte die Schultern weg von den Ohren.", "Halte den unteren Rücken angenehm.", "Halte den Nacken neutral."},
		Wrong:   []string{"Überstreck nicht den unteren Rücken.", "Wirf den Kopf nicht nach hinten.", "Zieh die Schultern nicht hoch.", "Erzwing die Position nicht."},
	},
	"CD05": {
		Desc:    "Eine ruhige Atemübung zum Abschluss der Einheit.",
		HowTo:   []string{"Auf den Rücken legen.", "Schultern und Kiefer entspannen.", "Wenn angenehm, eine Hand auf den Bauch legen.", "Sanft einatmen.", "Langsam ausatmen."},
		Correct: []string{"Atme ruhig.", "Halte die Schultern entspannt.", "Halte das Gesicht weich.", "Lass den Körper zur Ruhe kommen."},
		Wrong:   []string{"Atme nicht zu kräftig.", "Halte nicht die Luft an.", "Spann den Nacken nicht an.", "Überstreck nicht den unteren Rücken."},
	},

	// ---- Vlad-Set Ergänzungen (Aufwärmen / Cardio / Plyometrie / Ausfallschritte) ----
	"W07": {
		Desc:    "Eine ruhige Eröffnungsübung: tiefe Atmung mit einem sanften Ganzkörper-Strecken, das die Wirbelsäule länger macht.",
		HowTo:   []string{"Aufrecht stehen, Füße hüftbreit.", "Einatmen und beide Arme über den Kopf strecken.", "Die Wirbelsäule sanft in die Länge ziehen.", "Ausatmen und die Arme senken, die Schultern fallen lassen."},
		Correct: []string{"Atme langsam und voll.", "Lass die Schultern beim Ausatmen sinken.", "Halte die Rippen unten, ohne zu überstrecken.", "Bewege dich ohne Hast."},
		Wrong:   []string{"Halte nicht die Luft an.", "Überstreck den unteren Rücken nicht stark.", "Zieh die Schultern nicht zu den Ohren.", "Hetz das Strecken nicht."},
	},
	"W08": {
		Desc:    "Eine sanfte Rotationsübung, um Wirbelsäule und Rumpf aufzuwärmen.",
		HowTo:   []string{"Füße hüftbreit.", "Die Arme locker schwingen lassen.", "Den Oberkörper zu einer Seite drehen.", "Weich zur anderen Seite fließen."},
		Correct: []string{"Dreh dich aus dem Rumpf.", "Lass das Becken überwiegend nach vorn zeigen.", "Halte die Füße stabil.", "Bewege dich weich und gleichmäßig."},
		Wrong:   []string{"Dreh nicht stark in den Knien.", "Reiß nicht mit den Armen.", "Dreh dich nicht zu schnell.", "Halte nicht die Luft an."},
	},
	"W09": {
		Desc:    "Ein dynamisches Schulter- und Brust-Aufwärmen, das den Oberkörper auf Liegestütze und Sprungübungen vorbereitet.",
		HowTo:   []string{"Aufrecht stehen, Füße hüftbreit.", "Die Arme kontrolliert nach vorn und hinten schwingen.", "Die Brust beim Zurückschwingen natürlich öffnen lassen.", "Die Bewegung weich und entspannt halten."},
		Correct: []string{"Bewege die Arme frei.", "Halte die Schultern entspannt.", "Lass die Brust sanft öffnen.", "Bleib aufrecht."},
		Wrong:   []string{"Schwing nicht zu kräftig.", "Heb die Schultern nicht zu den Ohren.", "Überstreck nicht den unteren Rücken.", "Lass die Bewegung nicht in unkontrollierten Schwung übergehen."},
	},
	"C13": {
		Desc:    "Eine Rotationsübung für die schrägen und vorderen Bauchmuskeln.",
		HowTo:   []string{"Auf den Rücken legen.", "Die Knie anziehen.", "Die Hände locker neben den Kopf legen.", "Den Oberkörper drehen und einen Ellbogen Richtung gegenüberliegendes Knie führen.", "In einem kontrollierten Rhythmus die Seiten wechseln."},
		Correct: []string{"Dreh dich aus dem Rumpf.", "Halte den unteren Rücken kontrolliert.", "Halte den Nacken entspannt.", "Bewege dich weich."},
		Wrong:   []string{"Zieh nicht am Nacken.", "Bewege nicht nur die Ellbogen.", "Hetz die Wiederholungen nicht.", "Überstreck nicht den unteren Rücken."},
	},
	"B11": {
		Desc:    "Eine Rückenübung, die gegenüberliegenden Arm und Bein in einer wechselnden, schwimmähnlichen Bewegung hebt.",
		HowTo:   []string{"Auf den Bauch legen, Arme über Kopf.", "Brust und Beine leicht heben.", "Einen Arm und das gegenüberliegende Bein anheben.", "Absenken und Seite wechseln, weich abwechselnd."},
		Correct: []string{"Heb nur ein kleines Stück.", "Lass den Nacken der Wirbelsäule folgen.", "Halte das Gesäß aktiv.", "Bewege dich gleichmäßig."},
		Wrong:   []string{"Überlaste den unteren Rücken nicht.", "Heb den Kopf nicht.", "Bewege dich nicht ruckartig.", "Halte nicht die Luft an."},
	},
	"L10": {
		Desc:    "Eine grundlegende Kniebeuge mit dem eigenen Körpergewicht für Beine, Hüften und Gesäß.",
		HowTo:   []string{"Etwa schulterbreit stehen.", "Das Becken leicht nach hinten schieben.", "Die Knie beugen und in die Kniebeuge senken.", "Die Brust offen halten.", "Über den ganzen Fuß wieder aufrichten."},
		Correct: []string{"Halte die Knie über den Zehen.", "Halte die Fersen am Boden.", "Mach den Rücken lang.", "Halte die Tiefe kontrolliert."},
		Wrong:   []string{"Lass die Knie nicht nach innen fallen.", "Heb die Fersen nicht.", "Mach den Rücken nicht rund.", "Fall nicht zu schnell nach unten."},
	},
	"L11": {
		Desc:    "Eine Unterkörperübung für Beine, Gesäß und Gleichgewicht.",
		HowTo:   []string{"Aufrecht stehen.", "Mit einem Bein einen Schritt zurück machen.", "Kontrolliert in einen Ausfallschritt senken.", "Über den vorderen Fuß zurück in den Stand drücken.", "Seite wechseln."},
		Correct: []string{"Halte das vordere Knie über den Zehen.", "Bleib aufrecht.", "Mach den Schritt zurück kontrolliert.", "Halte die vordere Ferse am Boden."},
		Wrong:   []string{"Lass das vordere Knie nicht nach innen fallen.", "Lehn dich nicht zu weit nach vorn.", "Senk dich nicht zu schnell ab.", "Stoß dich nicht zu stark vom hinteren Bein ab."},
	},
	"L12": {
		Desc:    "Eine langsam ausgeführte Kniebeuge ohne Gewicht, um Kontrolle und Beinkraft aufzubauen.",
		HowTo:   []string{"Füße schulterbreit.", "Über etwa drei Sekunden senken.", "Unten kurz pausieren.", "Mit Kontrolle aufrichten."},
		Correct: []string{"Halte das Tempo langsam.", "Halte die Knie über den Zehen.", "Halte die Fersen am Boden.", "Halte den Rücken neutral."},
		Wrong:   []string{"Fall nicht schnell nach unten.", "Feder nicht aus der unteren Position.", "Lass die Knie nicht nach innen fallen.", "Halte nicht die Luft an."},
	},
	"J01": {
		Desc:    "Ein leichtes Sprunggelenk- und Waden-Aufwärmen mit kleinen, schnellen Sprüngen.",
		HowTo:   []string{"Aufrecht stehen, Füße etwa hüftbreit.", "Die Knie leicht weich halten.", "Kleine, schnelle Sprünge überwiegend aus den Sprunggelenken machen.", "Leise landen und wiederholen."},
		Correct: []string{"Halte die Sprünge niedrig.", "Lande weich und leise.", "Halte die Knie leicht gebeugt.", "Bleib aufrecht."},
		Wrong:   []string{"Spring nicht zu hoch.", "Lande nicht laut.", "Streck die Knie nicht durch.", "Lass die Sprunggelenke nicht nach innen oder außen kippen."},
	},
	"J02": {
		Desc:    "Eine gelenkschonende Aufwärm-Alternative zum hohen Kniehub.",
		HowTo:   []string{"Aufrecht stehen.", "Ein Knie bis auf Hüfthöhe heben.", "Kontrolliert wieder absenken.", "Seite wechseln und weiter marschieren.", "Die Arme natürlich mitschwingen."},
		Correct: []string{"Bleib aufrecht.", "Heb die Knie kontrolliert.", "Setz die Füße weich auf.", "Atme gleichmäßig weiter."},
		Wrong:   []string{"Lehn dich nicht nach hinten.", "Schwing die Beine nicht ohne Kontrolle.", "Stampf nicht mit den Füßen auf.", "Halte nicht die Luft an."},
	},
	"J03": {
		Desc:    "Eine Ganzkörper-Cardio-Übung, die den Puls erhöht und Schultern, Hüften und Beine aufwärmt.",
		HowTo:   []string{"Aufrecht stehen, Arme an den Seiten.", "Die Füße auseinander springen und die Arme über den Kopf heben.", "Die Füße wieder zusammen springen und die Arme senken.", "In einem gleichmäßigen Rhythmus wiederholen."},
		Correct: []string{"Lande weich.", "Halte die Knie leicht gebeugt.", "Bewege die Arme weich.", "Atme rhythmisch weiter."},
		Wrong:   []string{"Lande nicht hart.", "Streck die Knie nicht durch.", "Zieh die Schultern nicht hoch.", "Bewege dich nicht so schnell, dass du die Kontrolle verlierst."},
	},
	"J04": {
		Desc:    "Eine dynamische Rumpf- und Cardio-Übung aus der hohen Plank.",
		HowTo:   []string{"In einer hohen Plank beginnen.", "Die Hände unter den Schultern halten.", "Ein Knie Richtung Brust treiben.", "Die Beine rhythmisch wechseln.", "Das Becken stabil halten."},
		Correct: []string{"Halte den Körper in einer festen Plank.", "Halte den Rumpf fest.", "Halte die Schultern über den Händen.", "Bewege die Knie kontrolliert."},
		Wrong:   []string{"Lass das Becken nicht nach oben springen.", "Lass den unteren Rücken nicht durchhängen.", "Bewege die Füße nicht zu wild.", "Lass die Schultern nicht einsinken."},
	},
	"J05": {
		Desc:    "Eine plyometrische Beinübung für Kraft, Kondition und Koordination.",
		HowTo:   []string{"Mit etwa schulterbreitem Stand beginnen.", "In eine kontrollierte Kniebeuge senken.", "Nach oben springen.", "Weich mit leicht gebeugten Knien landen.", "Die Position vor dem nächsten Sprung neu einnehmen."},
		Correct: []string{"Lande leise.", "Halte die Knie über den Zehen.", "Halte die Brust offen.", "Mach jeden Sprung kontrolliert."},
		Wrong:   []string{"Lande nicht hart.", "Lass die Knie nicht nach innen fallen.", "Spring nicht, bevor die Kniebeuge stabil ist.", "Hetz nicht ohne neues Einnehmen der Position."},
	},
	"J06": {
		Desc:    "Eine seitliche Sprungübung für Beine, Hüften, Gleichgewicht und Koordination.",
		HowTo:   []string{"Auf einem Bein stehen.", "Seitlich auf das andere Bein hüpfen.", "Weich landen und stabilisieren.", "Von Seite zu Seite wiederholen.", "Die Arme zum Gleichgewicht nutzen."},
		Correct: []string{"Lande leise.", "Halte das Knie über den Zehen.", "Halte den Oberkörper kontrolliert.", "Bewege dich weich von Seite zu Seite."},
		Wrong:   []string{"Lande nicht mit nach innen fallendem Knie.", "Spring nicht zu früh zu weit.", "Fall nicht in die Landung.", "Verdreh das Knie nicht beim Aufsetzen."},
	},
	"J07": {
		Desc:    "Eine Cardio-Übung, die den Puls erhöht und Hüften, Rumpf und Beine aktiviert.",
		HowTo:   []string{"Aufrecht stehen.", "Auf der Stelle laufen und die Knie hoch heben.", "Die Arme natürlich mitbewegen.", "Den Rhythmus schnell, aber kontrolliert halten."},
		Correct: []string{"Heb die Knie bis auf Hüfthöhe.", "Lande leicht.", "Halte den Rumpf aktiv.", "Bleib aufrecht."},
		Wrong:   []string{"Lehn dich nicht nach hinten.", "Lande nicht hart.", "Lass die Füße nicht auf den Boden klatschen.", "Verlier mit steigendem Tempo nicht die Haltung."},
	},
	"J08": {
		Desc:    "Eine Ganzkörper-Konditionsübung, die Kniebeuge, Plank und Sprung verbindet. In diesem Programm die Variante ohne Liegestütz nutzen, sofern kein Liegestütz ausdrücklich vorgesehen ist.",
		HowTo:   []string{"Aufrecht stehen.", "In die Hocke gehen und die Hände auf den Boden setzen.", "Die Füße in die Plank springen oder zurücksetzen.", "Die Füße wieder nach vorn springen oder setzen.", "Aufstehen und mit einem kleinen Sprung abschließen."},
		Correct: []string{"Setz die Hände unter den Schultern auf.", "Halte die Plank fest.", "Setz die Füße weich auf.", "Mach den Sprung kontrolliert."},
		Wrong:   []string{"Lass den unteren Rücken in der Plank nicht durchhängen.", "Lande nicht hart.", "Hetz nicht mit schlampigen Wiederholungen.", "Lass beim Vorspringen die Knie nicht nach innen fallen."},
	},
	"J09": {
		Desc:    "Eine fortgeschrittene plyometrische Beinübung für Kraft, Kondition und Koordination.",
		HowTo:   []string{"In einer Ausfallschritt-Position beginnen.", "Nach oben springen.", "Die Beine in der Luft wechseln.", "Weich im Ausfallschritt auf der anderen Seite landen.", "Das Gleichgewicht vor der nächsten Wiederholung neu finden."},
		Correct: []string{"Lande weich und kontrolliert.", "Halte das vordere Knie über den Zehen.", "Bleib aufrecht.", "Halte den Bewegungsumfang sicher."},
		Wrong:   []string{"Lande nicht hart.", "Lass das vordere Knie nicht nach innen fallen.", "Senk dich nicht zu tief ab.", "Bewege dich nicht schneller, als du kontrollieren kannst."},
	},
	"CD07": {
		Desc:    "Eine ruhige Dehnung für Rücken, Hüften und Schultern zum Herunterkommen.",
		HowTo:   []string{"Knien und das Becken Richtung Fersen sinken lassen.", "Die Arme nach vorn auf den Boden strecken.", "Die Stirn ablegen.", "Langsam atmen und entspannen."},
		Correct: []string{"Lass das Becken Richtung Fersen sinken.", "Mach den Rücken sanft länger.", "Lass die Schultern entspannen.", "Atme langsam weiter."},
		Wrong:   []string{"Zwing das Becken nicht nach unten.", "Spann die Schultern nicht an.", "Halte nicht die Luft an.", "Überlaste die Knie nicht."},
	},
}
