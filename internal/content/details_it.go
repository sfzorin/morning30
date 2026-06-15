package content

// detailsIT is the Italian rich content for every exercise (current library IDs).
var detailsIT = map[string]Detail{
	// ---- Riscaldamento ----
	"W01": {
		Desc:    "Un semplice riscaldamento delle spalle per preparare collo, spalle e parte alta della schiena.",
		HowTo:   []string{"Stai dritto, braccia rilassate.", "Fai lenti cerchi con le spalle all'indietro.", "Poi lentamente in avanti.", "Il movimento resta morbido e controllato."},
		Correct: []string{"Tieni il collo rilassato.", "Muovi le spalle in modo fluido.", "Lascia le braccia sciolte.", "Tieni il corpo eretto."},
		Wrong:   []string{"Non sollevare con forza le spalle.", "Non muoverti troppo veloce.", "Non tendere il collo.", "Non incurvare la parte alta della schiena."},
	},
	"W02": {
		Desc:    "Un riscaldamento di spalle e parte alta della schiena che attiva i deltoidi e gli stabilizzatori della spalla.",
		HowTo:   []string{"Stai dritto e alza le braccia all'altezza delle spalle.", "Fai piccoli cerchi in avanti.", "Poi piccoli cerchi all'indietro.", "Mantieni i cerchi piccoli e controllati."},
		Correct: []string{"Tieni le braccia all'altezza delle spalle.", "Fai cerchi piccoli.", "Tieni le spalle basse.", "Tieni il collo rilassato."},
		Wrong:   []string{"Non fare grandi oscillazioni.", "Non alzare le spalle alle orecchie.", "Non inarcare la zona lombare.", "Non trattenere il respiro."},
	},
	"W03": {
		Desc:    "Un esercizio di rotazione dolce per il tronco e la parte alta della colonna.",
		HowTo:   []string{"Piedi alla larghezza dei fianchi.", "I fianchi restano soprattutto rivolti in avanti.", "Ruota la parte alta del corpo da un lato.", "Torna al centro e ruota dall'altro lato."},
		Correct: []string{"Tieni i piedi stabili.", "Fai partire il movimento dal tronco.", "Ruota in modo fluido.", "Non torcere le ginocchia."},
		Wrong:   []string{"Non ruotare le ginocchia insieme al corpo.", "Non ruotare troppo veloce.", "Non inclinarti all'indietro.", "Non forzare l'ampiezza."},
	},
	"W04": {
		Desc:    "Un esercizio di mobilità per fianchi, femorali e controllo della zona lombare.",
		HowTo:   []string{"Stai dritto, piedi alla larghezza dei fianchi.", "Ammorbidisci leggermente le ginocchia.", "Spingi i fianchi all'indietro.", "Abbassa il busto con la schiena dritta.", "Torna su spingendo i fianchi in avanti."},
		Correct: []string{"Tieni la schiena dritta.", "Fai partire il movimento dai fianchi.", "Tieni le ginocchia morbide e poco piegate.", "Lascia che il collo segua la linea della colonna."},
		Wrong:   []string{"Non incurvare la schiena.", "Non trasformarlo in uno squat.", "Non scendere troppo.", "Non alzare troppo la testa."},
	},
	"W05": {
		Desc:    "Un riscaldamento delle gambe controllato in un'ampiezza sicura e poco profonda.",
		HowTo:   []string{"Piedi poco più larghi dei fianchi.", "Spingi leggermente i fianchi all'indietro.", "Piega le ginocchia solo fino a una profondità comoda.", "Rialzati lentamente."},
		Correct: []string{"Tieni le ginocchia in direzione delle dita.", "Tieni i talloni a terra.", "Mantieni la schiena lunga.", "Scendi solo a una profondità moderata."},
		Wrong:   []string{"Non scendere troppo.", "Non far cedere le ginocchia verso l'interno.", "Non alzare i talloni.", "Non scendere troppo veloce."},
	},
	"W06": {
		Desc:    "Un riscaldamento per spalle, polsi e core.",
		HowTo:   []string{"Inizia in plank alto.", "Mani sotto le spalle.", "Sposta il corpo leggermente in avanti.", "Torna alla posizione iniziale.", "Il movimento è piccolo e controllato."},
		Correct: []string{"Tieni il corpo in una linea.", "Tieni le spalle stabili.", "Tieni il core attivo.", "Muoviti in modo fluido."},
		Wrong:   []string{"Non far cadere i fianchi.", "Non spingere le spalle verso le orecchie.", "Non andare troppo in avanti.", "Non far cedere la zona lombare."},
	},

	// ---- Addome / core ----
	"C01": {
		Desc:    "Un esercizio di stabilità del core per addome, spalle, glutei e tronco.",
		HowTo:   []string{"Metti i gomiti sotto le spalle.", "Porta i piedi indietro.", "Forma una linea retta dalle spalle ai talloni.", "Contrai addome e glutei.", "Mantieni la posizione."},
		Correct: []string{"Tieni il corpo in una linea.", "Tieni i gomiti sotto le spalle.", "Tieni il core attivo.", "Respira in modo regolare."},
		Wrong:   []string{"Non far cedere la zona lombare.", "Non alzare troppo i fianchi.", "Non guardare avanti.", "Non trattenere il respiro."},
	},
	"C02": {
		Desc:    "Un plank breve e intenso incentrato sulla tensione di tutto il corpo.",
		HowTo:   []string{"Inizia in plank sugli avambracci.", "Contrai forte l'addome.", "Stringi i glutei.", "Immagina di tirare i gomiti verso le dita dei piedi.", "Mantieni con il massimo controllo."},
		Correct: []string{"Metti in tensione tutto il corpo.", "Tieni i glutei attivi.", "Tieni il core saldo.", "Respira in modo controllato."},
		Wrong:   []string{"Non tenerlo come un plank rilassato.", "Non far cedere i fianchi.", "Non trattenere il respiro.", "Non provare a tenere troppo a lungo."},
	},
	"C03": {
		Desc:    "Un esercizio per il basso addome con movimenti incrociati controllati delle gambe.",
		HowTo:   []string{"Sdraiati sulla schiena.", "Premi delicatamente la zona lombare a terra.", "Solleva le gambe.", "Incrocia una gamba sopra l'altra.", "Continua alternando l'incrocio."},
		Correct: []string{"Tieni la zona lombare vicina al pavimento.", "Muovi le gambe in modo controllato.", "Tieni il collo rilassato.", "Muoviti in modo fluido."},
		Wrong:   []string{"Non inarcare la zona lombare.", "Non abbassare troppo le gambe.", "Non tirare la testa in avanti.", "Non muoverti troppo veloce."},
	},
	"C04": {
		Desc:    "Un esercizio per il basso addome con brevi colpi alternati delle gambe.",
		HowTo:   []string{"Sdraiati sulla schiena.", "Premi la zona lombare a terra.", "Solleva le gambe.", "Muovi le gambe su e giù con brevi colpi alternati."},
		Correct: []string{"Mantieni l'ampiezza ridotta.", "Tieni la zona lombare stabile.", "Non far scendere troppo le gambe.", "Respira in modo regolare."},
		Wrong:   []string{"Non inarcare la zona lombare.", "Non fare colpi ampi.", "Non tendere il collo.", "Non trattenere il respiro."},
	},
	"C05": {
		Desc:    "Un esercizio addominale incentrato sul sollevamento del bacino con controllo.",
		HowTo:   []string{"Sdraiati sulla schiena.", "Piega le ginocchia e solleva i piedi.", "Usa l'addome per arrotolare leggermente il bacino verso l'alto.", "Riabbassa lentamente."},
		Correct: []string{"Fai partire il movimento dal basso addome.", "Solleva il bacino con controllo.", "Tieni il collo rilassato.", "Scendi lentamente."},
		Wrong:   []string{"Non oscillare le gambe.", "Non usare lo slancio.", "Non tirare il collo.", "Non far cadere il bacino veloce."},
	},
	"C06": {
		Desc:    "Una tenuta statica del core con leva accorciata per un controllo migliore.",
		HowTo:   []string{"Sdraiati sulla schiena.", "Premi la zona lombare a terra.", "Solleva leggermente le spalle.", "Ginocchia piegate o gambe più in alto.", "Mantieni la posizione."},
		Correct: []string{"Tieni la zona lombare a terra.", "Tieni l'addome saldo.", "Tieni il collo rilassato.", "Continua a respirare."},
		Wrong:   []string{"Non inarcare la zona lombare.", "Non abbassare troppo le gambe.", "Non portare il mento in avanti.", "Non trattenere il respiro."},
	},
	"C07": {
		Desc:    "Un esercizio laterale per il core, per gli obliqui e la stabilità di spalla e anca.",
		HowTo:   []string{"Sdraiati su un fianco.", "Metti il gomito sotto la spalla.", "Solleva i fianchi.", "Il corpo in una linea retta.", "Mantieni."},
		Correct: []string{"Tieni il gomito sotto la spalla.", "Tieni i fianchi sollevati.", "Mantieni il corpo lungo.", "Tieni il collo neutro."},
		Wrong:   []string{"Non far cadere i fianchi.", "Non avvicinare la spalla all'orecchio.", "Non ruotare il petto in avanti.", "Non trattenere il respiro."},
	},
	"C08": {
		Desc:    "Una variante più impegnativa del plank laterale per gli obliqui e l'anca laterale.",
		HowTo:   []string{"Inizia in plank laterale.", "Abbassa leggermente i fianchi.", "Rialza i fianchi.", "Ripeti con controllo.", "Cambia lato."},
		Correct: []string{"Fai un movimento piccolo e controllato.", "Tieni il gomito sotto la spalla.", "Fai salire i fianchi in verticale.", "Tieni il core saldo."},
		Wrong:   []string{"Non scendere troppo.", "Non torcere il tronco.", "Non cedere nella spalla.", "Non muoverti troppo veloce."},
	},
	"C10": {
		Desc:    "Un esercizio di core e tricipiti che alterna plank sugli avambracci e plank alto.",
		HowTo:   []string{"Inizia in plank sugli avambracci.", "Appoggia una mano a terra e spingi su.", "Porta l'altra mano in plank alto.", "Riscendi sugli avambracci.", "Alterna la mano che guida."},
		Correct: []string{"Tieni i fianchi stabili.", "Tieni il core attivo.", "Muoviti in modo controllato.", "Tieni le spalle sopra mani o gomiti."},
		Wrong:   []string{"Non oscillare i fianchi da lato a lato.", "Non far cedere la zona lombare.", "Non muoverti troppo veloce.", "Non far cedere le spalle."},
	},
	"C11": {
		Desc:    "Un esercizio per il core controllato e un'alternativa più sicura ai sollevamenti di gambe più impegnativi.",
		HowTo:   []string{"Sdraiati sulla schiena.", "Solleva le gambe in posizione a tavolino.", "Abbassa un tallone verso il pavimento.", "Tocca leggermente e torna.", "Alterna i lati."},
		Correct: []string{"Tieni la zona lombare stabile.", "Muoviti lentamente.", "Fai toccare il tallone delicatamente.", "Tieni il bacino fermo."},
		Wrong:   []string{"Non far cadere la gamba troppo veloce.", "Non inarcare la zona lombare.", "Non muovere entrambe le gambe insieme.", "Non affrettarti."},
	},
	"C12": {
		Desc:    "Un esercizio di controllo del core per l'addome profondo e la stabilità della colonna.",
		HowTo:   []string{"Sdraiati sulla schiena.", "Solleva braccia e gambe in posizione a tavolino.", "Estendi il braccio e la gamba opposti.", "Mantieni stabile la zona lombare.", "Torna e cambia lato."},
		Correct: []string{"Tieni la zona lombare vicina al pavimento.", "Muoviti lentamente.", "Muovi insieme il braccio e la gamba opposti.", "Tieni il tronco fermo."},
		Wrong:   []string{"Non inarcare la zona lombare.", "Non muoverti troppo veloce.", "Non abbassare troppo la gamba.", "Non far sporgere le costole."},
	},

	// ---- Piegamenti / tricipiti / spalle ----
	"P01": {
		Desc:    "Un esercizio base per la parte alta del corpo: petto, tricipiti, spalle e core.",
		HowTo:   []string{"Inizia in plank alto.", "Mani sotto le spalle o poco più larghe.", "Abbassa il corpo con controllo.", "Spingi su mantenendo il corpo dritto."},
		Correct: []string{"Tieni il corpo in una linea.", "Tieni il core saldo.", "Tieni i gomiti a circa 30–45 gradi.", "Muoviti in modo controllato."},
		Wrong:   []string{"Non far cadere i fianchi.", "Non alzare troppo i fianchi.", "Non aprire i gomiti dritti ai lati.", "Non cadere nella posizione bassa."},
	},
	"P02": {
		Desc:    "Una variante di piegamento con più lavoro sui tricipiti.",
		HowTo:   []string{"Inizia in plank alto.", "Metti le mani più vicine rispetto a un piegamento normale.", "Tieni i gomiti vicini al corpo.", "Scendi lentamente e spingi su."},
		Correct: []string{"Tieni i gomiti vicini.", "Tieni il corpo dritto.", "Tieni le spalle lontane dalle orecchie.", "Tieni i polsi comodi."},
		Wrong:   []string{"Non mettere le mani troppo vicine.", "Non aprire i gomiti.", "Non far cedere la zona lombare.", "Non forzare se senti dolore a polso o gomito."},
	},
	"P03": {
		Desc:    "Piegamenti controllati con una breve pausa nella posizione bassa.",
		HowTo:   []string{"Inizia in posizione di piegamento normale.", "Scendi lentamente.", "Pausa un secondo vicino al fondo.", "Spingi su senza perdere la posizione del corpo."},
		Correct: []string{"Fai una pausa controllata.", "Tieni il core saldo.", "Tieni il petto attivo.", "Tieni il corpo dritto."},
		Wrong:   []string{"Non rilassarti in basso.", "Non far cadere i fianchi.", "Non trattenere il respiro.", "Non rimbalzare dal fondo."},
	},
	"P04": {
		Desc:    "Una variante lenta di piegamento che aumenta controllo e tempo sotto tensione.",
		HowTo:   []string{"Inizia in plank alto.", "Scendi per circa tre secondi.", "Pausa un secondo vicino al fondo.", "Spingi su per circa tre secondi."},
		Correct: []string{"Mantieni un ritmo lento e uniforme.", "Tieni il corpo dritto.", "Tieni i gomiti controllati.", "Respira in modo regolare."},
		Wrong:   []string{"Non affrettare il movimento.", "Non crollare in basso.", "Non far cadere i fianchi.", "Non perdere la tensione del core."},
	},
	"P05": {
		Desc:    "Una variante asimmetrica di piegamento che impegna petto, spalle e core.",
		HowTo:   []string{"Inizia in posizione di piegamento.", "Metti una mano un po' avanti e l'altra un po' indietro.", "Scendi con controllo.", "Spingi su.", "Cambia la posizione delle mani dall'altro lato."},
		Correct: []string{"Tieni il corpo dritto.", "Tieni il core saldo.", "Tieni entrambe le spalle stabili.", "Muoviti in modo controllato."},
		Wrong:   []string{"Non torcere il tronco.", "Non far cadere una spalla.", "Non allargare troppo le mani.", "Non forzare l'ampiezza se la spalla dà fastidio."},
	},
	"P06": {
		Desc:    "Una variante di piegamento incentrata sulle spalle.",
		HowTo:   []string{"Inizia in plank alto.", "Avvicina un po' i piedi e solleva i fianchi.", "La testa tra le braccia.", "Abbassa la testa verso il pavimento.", "Spingi di nuovo su."},
		Correct: []string{"Tieni i fianchi alti.", "Piega i gomiti con controllo.", "Tieni le spalle stabili.", "Tieni il collo neutro."},
		Wrong:   []string{"Non cadere in un piegamento normale.", "Non portare la testa troppo in avanti.", "Non aprire troppo i gomiti.", "Non forzare se senti dolore alla spalla."},
	},
	"P08": {
		Desc:    "Un esercizio per tricipiti e spalle da una posizione di appoggio sugli avambracci.",
		HowTo:   []string{"Inizia sugli avambracci, il corpo lungo.", "Gomiti sotto o un po' davanti alle spalle.", "Spingi con mani e avambracci.", "Estendi leggermente i gomiti.", "Torna con controllo."},
		Correct: []string{"Tieni il core saldo.", "Muovi i gomiti in modo fluido.", "Tieni le spalle basse.", "Mantieni un'ampiezza comoda."},
		Wrong:   []string{"Non forzare i gomiti.", "Non far cadere i fianchi.", "Non alzare le spalle.", "Non muoverti se senti dolore a gomito o spalla."},
	},

	// ---- Schiena / postura ----
	"B01": {
		Desc:    "Un esercizio per la schiena incentrato sulla postura, per la parte alta della schiena, le scapole e gli estensori della colonna.",
		HowTo:   []string{"Sdraiati a pancia in giù.", "Le braccia lungo il corpo.", "Ruota leggermente i pollici verso l'esterno.", "Solleva un po' il petto.", "Porta le scapole indietro e in basso.", "Mantieni."},
		Correct: []string{"Solleva solo di poco.", "Tieni il collo neutro.", "Porta le scapole indietro e in basso.", "Attiva leggermente i glutei."},
		Wrong:   []string{"Non sollevare troppo.", "Non buttare indietro la testa.", "Non iperestendere la zona lombare.", "Non alzare le spalle alle orecchie."},
	},
	"B02": {
		Desc:    "Un esercizio per la parte alta della schiena e le scapole a pancia in giù.",
		HowTo:   []string{"Sdraiati sulla pancia.", "La testa guarda in basso.", "Muovi le braccia in un ampio arco dai lati verso l'alto.", "Torna per lo stesso percorso.", "Movimento controllato."},
		Correct: []string{"Tieni il collo rilassato.", "Muovi le scapole in modo fluido.", "Solleva il petto solo di poco o tienilo neutro.", "Muovi le braccia senza fretta."},
		Wrong:   []string{"Non sollevare la testa.", "Non iperestendere troppo la zona lombare.", "Non oscillare le braccia.", "Non alzare le spalle alle orecchie."},
	},
	"B03": {
		Desc:    "Un esercizio per la parte alta della schiena, per le scapole e le spalle posteriori.",
		HowTo:   []string{"Sdraiati a pancia in giù.", "Piega i gomiti a forma di W.", "Solleva un po' gomiti e mani.", "Stringi le scapole.", "Riabbassa con controllo."},
		Correct: []string{"Mantieni l'ampiezza ridotta.", "Lascia che siano le scapole a lavorare.", "Tieni il collo neutro.", "Tieni le spalle lontane dalle orecchie."},
		Wrong:   []string{"Non sollevare troppo.", "Non sollevare la testa.", "Non usare lo slancio.", "Non iperestendere la zona lombare."},
	},
	"B04": {
		Desc:    "Un esercizio per le scapole, per il trapezio inferiore e la parte alta della schiena.",
		HowTo:   []string{"Sdraiati a pancia in giù.", "Estendi le braccia in diagonale verso l'alto a forma di Y.", "Solleva un po' le braccia dal pavimento.", "Riabbassa con controllo."},
		Correct: []string{"Solleva le braccia solo un po'.", "Tieni il collo neutro.", "Tieni le spalle basse.", "Muoviti in modo controllato."},
		Wrong:   []string{"Non sollevare troppo.", "Non alzare le spalle alle orecchie.", "Non inarcare la zona lombare.", "Non muoverti troppo veloce."},
	},
	"B05": {
		Desc:    "Un esercizio per la parte alta della schiena, per le spalle posteriori e il controllo delle scapole.",
		HowTo:   []string{"Sdraiati a pancia in giù.", "Estendi le braccia ai lati a forma di T.", "Solleva un po' le braccia.", "Stringi le scapole.", "Riabbassa lentamente."},
		Correct: []string{"Tieni le braccia sulla linea delle spalle.", "Tieni il collo rilassato.", "Muovi le scapole dolcemente.", "Solleva solo di poco."},
		Wrong:   []string{"Non lanciare le braccia verso l'alto.", "Non sollevare la testa.", "Non alzare le spalle.", "Non forzare un'ampiezza ampia."},
	},
	"B06": {
		Desc:    "Un esercizio per la schiena che imita un movimento di tirata senza attrezzi.",
		HowTo:   []string{"Sdraiati a pancia in giù con le braccia sopra la testa.", "Solleva un po' il petto.", "Tira i gomiti in basso verso le costole.", "Stringi le scapole.", "Estendi di nuovo le braccia in avanti."},
		Correct: []string{"Fai scendere i gomiti con controllo.", "Stringi dolcemente le scapole.", "Tieni il collo neutro.", "Tieni la zona lombare comoda."},
		Wrong:   []string{"Non dare strattoni con le braccia.", "Non sollevare troppo.", "Non buttare indietro la testa.", "Non alzare le spalle alle orecchie."},
	},
	"B07": {
		Desc:    "Un esercizio per la schiena e la coordinazione a pancia in giù.",
		HowTo:   []string{"Sdraiati sulla pancia.", "Solleva un po' il petto.", "Muovi braccio e gamba opposti in uno schema di nuoto controllato.", "Il movimento è piccolo e costante."},
		Correct: []string{"Tieni il collo neutro.", "Muoviti in modo controllato.", "Tieni la zona lombare comoda.", "Continua a respirare."},
		Wrong:   []string{"Non muoverti troppo veloce.", "Non sollevare troppo.", "Non buttare indietro la testa.", "Non forzare la zona lombare."},
	},
	"B08": {
		Desc:    "Un esercizio per la catena posteriore: glutei, femorali, spalle e schiena.",
		HowTo:   []string{"Siediti a terra con le gambe distese.", "Metti le mani dietro i fianchi.", "Spingi con mani e talloni.", "Solleva i fianchi.", "Tieni il corpo in una linea lunga."},
		Correct: []string{"Tieni il petto aperto.", "Tieni i fianchi sollevati.", "Tieni le spalle stabili.", "Tieni il collo neutro."},
		Wrong:   []string{"Non far cadere i fianchi.", "Non alzare le spalle.", "Non buttare indietro la testa.", "Non forzare i polsi."},
	},
	"B09": {
		Desc:    "Un esercizio per la schiena di piccola ampiezza per gli estensori della colonna e la postura.",
		HowTo:   []string{"Sdraiati a pancia in giù.", "Le mani lungo il corpo o leggermente vicino al petto.", "Solleva un po' il petto.", "Fai piccole pulsazioni controllate.", "Collo neutro."},
		Correct: []string{"Solleva solo di poco.", "Fai pulsazioni dolci.", "Mantieni il collo lungo.", "Tieni la zona lombare controllata."},
		Wrong:   []string{"Non pulsare troppo in alto.", "Non usare lo slancio.", "Non guardare avanti.", "Non continuare se senti dolore acuto nella zona lombare."},
	},
	"B10": {
		Desc:    "Un esercizio per la schiena più impegnativo che unisce un piccolo sollevamento a cobra con una tirata delle braccia a forma di W.",
		HowTo:   []string{"Sdraiati a pancia in giù con le braccia in avanti.", "Solleva un po' il petto.", "Tira i gomiti in basso e indietro a forma di W.", "Stringi le scapole.", "Estendi di nuovo in avanti e riabbassa con controllo."},
		Correct: []string{"Solleva il petto solo di poco.", "Esegui la tirata a W in modo fluido.", "Porta le scapole indietro e in basso.", "Mantieni il collo lungo."},
		Wrong:   []string{"Non iperestendere la zona lombare.", "Non tirare con il collo.", "Non muoverti troppo veloce.", "Non alzare le spalle."},
	},

	// ---- Braccia / bicipiti ----
	"A01": {
		Desc:    "Un esercizio per i bicipiti senza attrezzi: una mano oppone resistenza all'altra.",
		HowTo:   []string{"In piedi o seduto, ben dritto.", "Piega un gomito come in un curl.", "Con l'altra mano premi verso il basso sull'avambraccio.", "Sali lentamente contro la resistenza.", "Cambia lato."},
		Correct: []string{"Mantieni una resistenza costante.", "Muoviti lentamente.", "Tieni il gomito vicino al corpo.", "Tieni la spalla bassa."},
		Wrong:   []string{"Non muoverti senza vera resistenza.", "Non dare strattoni col braccio.", "Non alzare la spalla.", "Non trattenere il respiro."},
	},
	"A02": {
		Desc:    "Una tenuta statica per i bicipiti con autoresistenza.",
		HowTo:   []string{"Piega un gomito a circa 90 gradi.", "Con l'altra mano premi verso il basso.", "Il braccio che lavora resiste senza muoversi.", "Mantieni la tensione.", "Cambia lato."},
		Correct: []string{"Tieni stabile l'angolo del gomito.", "Mantieni una tensione costante.", "Tieni le spalle rilassate.", "Continua a respirare."},
		Wrong:   []string{"Non premere a strattoni.", "Non torcere il tronco.", "Non alzare la spalla.", "Non trattenere il respiro."},
	},
	"A03": {
		Desc:    "Un esercizio isometrico per braccia e avambracci, le mani una contro l'altra.",
		HowTo:   []string{"Aggancia le dita o le mani.", "I gomiti leggermente piegati.", "Tira le mani per allontanarle.", "Mantieni una tensione costante.", "Rilascia lentamente."},
		Correct: []string{"Mantieni una tensione controllata.", "Tieni le spalle basse.", "Tieni i polsi comodi.", "Respira con calma."},
		Wrong:   []string{"Non tirare di scatto.", "Non alzare le spalle.", "Non piegare troppo i polsi.", "Non trattenere il respiro."},
	},
	"A04": {
		Desc:    "Un esercizio di discesa lenta per i bicipiti con autoresistenza.",
		HowTo:   []string{"Inizia con un gomito piegato.", "Con l'altra mano crea resistenza.", "Abbassa lentamente l'avambraccio.", "Mantieni la resistenza durante la discesa.", "Cambia lato."},
		Correct: []string{"Scendi lentamente.", "Mantieni una resistenza costante.", "Tieni il gomito vicino al corpo.", "Tieni la spalla rilassata."},
		Wrong:   []string{"Non far cadere il braccio veloce.", "Non muoverti senza resistenza.", "Non torcere il tronco.", "Non trattenere il respiro."},
	},

	// ---- Gambe / glutei ----
	"L01": {
		Desc:    "Un esercizio per le gambe controllato per cosce, fianchi e postura.",
		HowTo:   []string{"Piedi poco più larghi dei fianchi.", "Spingi i fianchi all'indietro.", "Piega le ginocchia fino a una profondità comoda e poco profonda.", "Rialzati lentamente."},
		Correct: []string{"Tieni le ginocchia in direzione delle dita.", "Tieni i talloni a terra.", "Mantieni la schiena lunga.", "Scendi solo a una profondità comoda."},
		Wrong:   []string{"Non scendere troppo.", "Non far cedere le ginocchia verso l'interno.", "Non alzare i talloni.", "Non scendere veloce."},
	},
	"L02": {
		Desc:    "Una variante di squat controllata con una breve pausa in un'ampiezza sicura.",
		HowTo:   []string{"In piedi e stabile.", "Scendi in uno squat comodo e poco profondo.", "Pausa breve.", "Rialzati con controllo."},
		Correct: []string{"Fai una pausa stabile.", "Tieni le ginocchia in direzione delle dita.", "Tieni il petto aperto.", "Tieni i talloni a terra."},
		Wrong:   []string{"Non scendere troppo.", "Non rimbalzare nella posizione bassa.", "Non far cedere le ginocchia verso l'interno.", "Non trattenere il respiro."},
	},
	"L03": {
		Desc:    "Un esercizio di mobilità per fianchi, femorali e controllo della zona lombare.",
		HowTo:   []string{"Stai dritto, piedi alla larghezza dei fianchi.", "Ammorbidisci leggermente le ginocchia.", "Spingi i fianchi all'indietro.", "Abbassa il busto con la schiena dritta.", "Torna su spingendo i fianchi in avanti."},
		Correct: []string{"Tieni la schiena dritta.", "Fai partire il movimento dai fianchi.", "Tieni le ginocchia morbide e poco piegate.", "Lascia che il collo segua la linea della colonna."},
		Wrong:   []string{"Non incurvare la schiena.", "Non trasformarlo in uno squat.", "Non scendere troppo.", "Non alzare troppo la testa."},
	},
	"L04": {
		Desc:    "Un esercizio per i glutei più impegnativo, una gamba alla volta.",
		HowTo:   []string{"Sdraiati sulla schiena.", "Piega un ginocchio con quel piede a terra.", "Estendi o solleva l'altra gamba.", "Spingi con il tallone che lavora.", "Solleva i fianchi e riabbassa lentamente."},
		Correct: []string{"Fai partire il sollevamento dal gluteo che lavora.", "Tieni i fianchi a livello.", "Muoviti in modo controllato.", "Tieni la zona lombare neutra."},
		Wrong:   []string{"Non torcere il bacino.", "Non spingere sulle dita dei piedi.", "Non inarcare la zona lombare.", "Non scendere troppo veloce."},
	},
	"L05": {
		Desc:    "Un esercizio per i glutei e la stabilità del bacino.",
		HowTo:   []string{"Inizia in posizione di ponte glutei.", "Tieni i fianchi sollevati.", "Solleva un piede leggermente da terra.", "Riabbassalo e cambia lato.", "Tieni il bacino stabile."},
		Correct: []string{"Tieni i fianchi a livello.", "Tieni i glutei attivi.", "Muoviti lentamente.", "Tieni la zona lombare comoda."},
		Wrong:   []string{"Non far cadere i fianchi.", "Non oscillare da lato a lato.", "Non muoverti troppo veloce.", "Non inarcare la zona lombare."},
	},
	"L06": {
		Desc:    "Un esercizio per i glutei che rafforza fianchi e catena posteriore.",
		HowTo:   []string{"Sdraiati sulla schiena.", "Piega le ginocchia, piedi a terra.", "Solleva i fianchi.", "Pausa in alto.", "Riabbassa lentamente."},
		Correct: []string{"Spingi con i talloni.", "Stringi i glutei in alto.", "Tieni le costole in basso.", "Non iperestendere la zona lombare."},
		Wrong:   []string{"Non iperestendere la zona lombare in alto.", "Non alzare i talloni.", "Non far cadere i fianchi troppo veloce.", "Non mettere i piedi troppo vicini."},
	},
	"L07": {
		Desc:    "Un esercizio per la parte bassa della gamba per polpacci, caviglie e controllo del piede.",
		HowTo:   []string{"In piedi, piedi alla larghezza dei fianchi.", "Sali sulle punte dei piedi.", "Pausa breve in alto.", "Abbassa i talloni lentamente."},
		Correct: []string{"Tieni il corpo eretto.", "Muoviti in modo fluido.", "Tieni le caviglie allineate.", "Scendi in modo controllato."},
		Wrong:   []string{"Non far cadere i talloni veloce.", "Non oscillare in avanti.", "Non far cadere le caviglie verso l'esterno.", "Non usare lo slancio."},
	},
	"L09": {
		Desc:    "Un esercizio per la catena laterale per glutei laterali, obliqui e stabilità dell'anca.",
		HowTo:   []string{"Inizia in plank laterale.", "Il corpo in una linea retta.", "Solleva un po' la gamba in alto.", "Riabbassala con controllo.", "Ripeti e cambia lato."},
		Correct: []string{"Tieni i fianchi sollevati.", "Muovi lentamente la gamba in alto.", "Tieni il tronco stabile.", "Tieni la spalla salda."},
		Wrong:   []string{"Non far cadere i fianchi.", "Non oscillare la gamba.", "Non ruotare il corpo all'indietro.", "Non cedere nella spalla."},
	},

	// ---- Defaticamento ----
	"CD01": {
		Desc:    "Un allungamento dolce per il petto e la parte anteriore delle spalle.",
		HowTo:   []string{"Sdraiati sulla schiena.", "Apri le braccia ai lati.", "Rilassa le spalle.", "Lascia che il petto si apra dolcemente.", "Respira lentamente."},
		Correct: []string{"Lascia che il petto si apra dolcemente.", "Tieni le spalle rilassate.", "Tieni la zona lombare comoda.", "Fermati se senti dolore alle spalle."},
		Wrong:   []string{"Non forzare le braccia verso il basso.", "Non inarcare la zona lombare.", "Non tendere il collo.", "Non allungare se senti dolore."},
	},
	"CD02": {
		Desc:    "Un allungamento dolce della parte posteriore della coscia.",
		HowTo:   []string{"Sdraiati sulla schiena.", "Solleva una gamba.", "Tieni dietro la coscia o il polpaccio.", "Porta la gamba dolcemente verso di te.", "Cambia lato."},
		Correct: []string{"Mantieni un allungamento leggero.", "Lascia il ginocchio un po' piegato.", "Tieni la zona lombare tranquilla.", "Respira lentamente."},
		Wrong:   []string{"Non tirare troppo forte.", "Non forzare la gamba dritta.", "Non sollevare i fianchi.", "Non continuare se senti dolore acuto dietro il ginocchio."},
	},
	"CD03": {
		Desc:    "Un allungamento dolce di rotazione per la colonna e il tronco.",
		HowTo:   []string{"Sdraiati sulla schiena.", "Piega le ginocchia.", "Lascia cadere le ginocchia dolcemente da un lato.", "Tieni le spalle a terra.", "Torna al centro e cambia lato."},
		Correct: []string{"Muoviti lentamente.", "Tieni le spalle a terra.", "Respira in modo rilassato.", "Mantieni un allungamento comodo."},
		Wrong:   []string{"Non forzare le ginocchia a terra.", "Non ruotare veloce.", "Non sollevare la spalla.", "Non allungare se senti dolore."},
	},
	"CD04": {
		Desc:    "Un allungamento dolce della parte anteriore del corpo e una lieve estensione della schiena.",
		HowTo:   []string{"Sdraiati a pancia in giù.", "Metti gli avambracci a terra.", "Solleva il petto dolcemente.", "Tieni il collo lungo.", "Respira lentamente."},
		Correct: []string{"Solleva il petto solo lievemente.", "Tieni le spalle lontane dalle orecchie.", "Tieni la zona lombare comoda.", "Tieni il collo neutro."},
		Wrong:   []string{"Non iperestendere la zona lombare.", "Non buttare indietro la testa.", "Non alzare le spalle.", "Non forzare la posizione."},
	},
	"CD05": {
		Desc:    "Un esercizio di respirazione calma per concludere la sessione.",
		HowTo:   []string{"Sdraiati sulla schiena.", "Rilassa spalle e mascella.", "Se comodo, metti una mano sulla pancia.", "Inspira dolcemente.", "Espira lentamente."},
		Correct: []string{"Respira con calma.", "Tieni le spalle rilassate.", "Tieni il viso disteso.", "Lascia che il corpo si calmi."},
		Wrong:   []string{"Non respirare troppo forte.", "Non trattenere il respiro.", "Non tendere il collo.", "Non inarcare la zona lombare."},
	},

	// ---- Aggiunte del set Vlad (riscaldamento / cardio / pliometria / affondi) ----
	"W07": {
		Desc:    "Un esercizio di apertura calmo: respirazione profonda con un dolce allungamento di tutto il corpo per allungare la colonna.",
		HowTo:   []string{"Stai dritto, piedi alla larghezza dei fianchi.", "Inspira e porta entrambe le braccia sopra la testa.", "Allungati dolcemente lungo la colonna.", "Espira e abbassa le braccia lasciando cadere le spalle."},
		Correct: []string{"Respira in modo lento e pieno.", "Rilassa le spalle sull'espirazione.", "Tieni le costole basse senza iperestendere.", "Muoviti senza fretta."},
		Wrong:   []string{"Non trattenere il respiro.", "Non inarcare con forza la zona lombare.", "Non alzare le spalle alle orecchie.", "Non affrettare l'allungamento."},
	},
	"W08": {
		Desc:    "Un esercizio di rotazione dolce per riscaldare la colonna e il tronco.",
		HowTo:   []string{"Piedi alla larghezza dei fianchi.", "Lascia oscillare le braccia in modo sciolto.", "Ruota la parte alta del corpo da un lato.", "Passa in modo fluido all'altro lato."},
		Correct: []string{"Fai partire il movimento dal tronco.", "Tieni i fianchi soprattutto rivolti in avanti.", "Tieni i piedi stabili.", "Mantieni un ritmo fluido e uniforme."},
		Wrong:   []string{"Non torcere con forza le ginocchia.", "Non dare strattoni con le braccia.", "Non ruotare troppo veloce.", "Non trattenere il respiro."},
	},
	"W09": {
		Desc:    "Un riscaldamento dinamico per spalle e petto, per preparare la parte alta del corpo a piegamenti e salti.",
		HowTo:   []string{"Stai dritto, piedi alla larghezza dei fianchi.", "Oscilla le braccia avanti e indietro in modo controllato.", "Lascia che il petto si apra naturalmente quando le braccia vanno indietro.", "Mantieni il movimento fluido e rilassato."},
		Correct: []string{"Muovi le braccia liberamente.", "Tieni le spalle rilassate.", "Lascia che il petto si apra dolcemente.", "Tieni il corpo eretto."},
		Wrong:   []string{"Non oscillare in modo troppo aggressivo.", "Non alzare le spalle alle orecchie.", "Non inarcare la zona lombare.", "Non trasformare il movimento in uno slancio incontrollato."},
	},
	"C13": {
		Desc:    "Un esercizio addominale di rotazione per gli obliqui e l'addome anteriore.",
		HowTo:   []string{"Sdraiati sulla schiena.", "Porta le ginocchia in alto.", "Appoggia le mani leggere vicino alla testa.", "Ruota il tronco e avvicina un gomito al ginocchio opposto.", "Cambia lato con un ritmo controllato."},
		Correct: []string{"Fai partire la rotazione dal tronco.", "Tieni la zona lombare controllata.", "Tieni il collo rilassato.", "Muoviti in modo fluido."},
		Wrong:   []string{"Non tirare il collo.", "Non muovere solo i gomiti.", "Non affrettare le ripetizioni.", "Non far inarcare la zona lombare."},
	},
	"B11": {
		Desc:    "Un esercizio per la schiena che solleva braccio e gamba opposti in un movimento alternato simile al nuoto.",
		HowTo:   []string{"Sdraiati a pancia in giù con le braccia sopra la testa.", "Solleva un po' il petto e le gambe.", "Alza un braccio e la gamba opposta.", "Riabbassa e cambia lato, alternando in modo fluido."},
		Correct: []string{"Solleva solo di poco.", "Lascia che il collo segua la linea della colonna.", "Tieni i glutei attivi.", "Muoviti in modo costante."},
		Wrong:   []string{"Non forzare la zona lombare.", "Non sollevare la testa.", "Non muoverti a scatti.", "Non trattenere il respiro."},
	},
	"L10": {
		Desc:    "Uno squat base a corpo libero per gambe, anche e glutei.",
		HowTo:   []string{"Stai in piedi, piedi circa alla larghezza delle spalle.", "Spingi leggermente le anche all'indietro.", "Piega le ginocchia e scendi in squat.", "Tieni il petto aperto.", "Rialzati spingendo con tutta la pianta del piede."},
		Correct: []string{"Lascia che le ginocchia seguano le punte.", "Tieni i talloni a terra.", "Mantieni la schiena lunga.", "Tieni la profondità controllata."},
		Wrong:   []string{"Non far cedere le ginocchia verso l'interno.", "Non alzare i talloni.", "Non incurvare la schiena.", "Non scendere troppo in fretta."},
	},
	"L11": {
		Desc:    "Un esercizio per la parte bassa del corpo, per gambe, glutei ed equilibrio.",
		HowTo:   []string{"Stai dritto.", "Porta una gamba indietro.", "Scendi in affondo con controllo.", "Spingi con il piede anteriore per tornare in piedi.", "Cambia lato."},
		Correct: []string{"Lascia che il ginocchio anteriore segua le punte.", "Tieni il busto eretto.", "Fai un passo indietro controllato.", "Tieni il tallone anteriore a terra."},
		Wrong:   []string{"Non far cedere il ginocchio anteriore verso l'interno.", "Non inclinarti troppo in avanti.", "Non scendere in fretta.", "Non spingere troppo con la gamba posteriore."},
	},
	"L12": {
		Desc:    "Uno squat a corpo libero eseguito lentamente per costruire controllo e forza nelle gambe.",
		HowTo:   []string{"Piedi alla larghezza delle spalle.", "Scendi in circa tre secondi.", "Pausa breve nella posizione bassa.", "Rialzati con controllo."},
		Correct: []string{"Mantieni il ritmo lento.", "Tieni le ginocchia in direzione delle dita.", "Tieni i talloni a terra.", "Tieni la schiena neutra."},
		Wrong:   []string{"Non scendere veloce.", "Non rimbalzare dalla posizione bassa.", "Non far cedere le ginocchia verso l'interno.", "Non trattenere il respiro."},
	},
	"J01": {
		Desc:    "Un leggero riscaldamento per caviglie e polpacci con piccoli salti rapidi.",
		HowTo:   []string{"Stai dritto, piedi circa alla larghezza dei fianchi.", "Tieni le ginocchia leggermente morbide.", "Fai piccoli salti rapidi usando soprattutto le caviglie.", "Atterra in modo silenzioso e ripeti."},
		Correct: []string{"Tieni i salti bassi.", "Atterra in modo morbido e silenzioso.", "Tieni le ginocchia leggermente piegate.", "Tieni il corpo eretto."},
		Wrong:   []string{"Non saltare troppo in alto.", "Non atterrare in modo rumoroso.", "Non bloccare le ginocchia.", "Non far cedere le caviglie verso l'interno o l'esterno."},
	},
	"J02": {
		Desc:    "Un'alternativa a basso impatto alle ginocchia alte.",
		HowTo:   []string{"Stai dritto.", "Solleva un ginocchio verso l'altezza dell'anca.", "Riabbassalo con controllo.", "Cambia lato e continua a marciare.", "Tieni le braccia in movimento naturale."},
		Correct: []string{"Tieni il busto eretto.", "Solleva le ginocchia con controllo.", "Fai atterrare i piedi in modo morbido.", "Respira in modo regolare."},
		Wrong:   []string{"Non inclinarti all'indietro.", "Non muovere le gambe senza controllo.", "Non sbattere i piedi a terra.", "Non trattenere il respiro."},
	},
	"J03": {
		Desc:    "Un esercizio cardio per tutto il corpo che alza la frequenza cardiaca e riscalda spalle, anche e gambe.",
		HowTo:   []string{"Stai dritto, braccia lungo i fianchi.", "Salta allargando i piedi mentre alzi le braccia sopra la testa.", "Salta riportando i piedi insieme mentre abbassi le braccia.", "Ripeti con un ritmo costante."},
		Correct: []string{"Atterra in modo morbido.", "Tieni le ginocchia leggermente piegate.", "Muovi le braccia in modo fluido.", "Respira in modo ritmico."},
		Wrong:   []string{"Non atterrare in modo duro.", "Non bloccare le ginocchia.", "Non alzare le spalle alle orecchie.", "Non muoverti troppo veloce perdendo il controllo."},
	},
	"J04": {
		Desc:    "Un esercizio dinamico di core e cardio eseguito dal plank alto.",
		HowTo:   []string{"Parti in plank alto.", "Tieni le mani sotto le spalle.", "Porta un ginocchio verso il petto.", "Cambia gamba a ritmo.", "Tieni le anche stabili."},
		Correct: []string{"Tieni il corpo in un plank solido.", "Tieni il core saldo.", "Tieni le spalle sopra le mani.", "Muovi le ginocchia con controllo."},
		Wrong:   []string{"Non far rimbalzare le anche verso l'alto.", "Non far cedere la zona lombare.", "Non muovere i piedi in modo troppo scomposto.", "Non far cedere le spalle."},
	},
	"J05": {
		Desc:    "Un esercizio pliometrico per le gambe, per potenza, condizionamento e coordinazione.",
		HowTo:   []string{"Parti con i piedi circa alla larghezza delle spalle.", "Scendi in uno squat controllato.", "Salta verso l'alto.", "Atterra morbido con le ginocchia leggermente piegate.", "Ritrova la posizione prima del salto successivo."},
		Correct: []string{"Atterra in modo silenzioso.", "Lascia che le ginocchia seguano le punte.", "Tieni il petto aperto.", "Esegui ogni salto in modo controllato."},
		Wrong:   []string{"Non atterrare in modo duro.", "Non far cedere le ginocchia verso l'interno.", "Non saltare prima che lo squat sia stabile.", "Non affrettarti senza ritrovare la posizione."},
	},
	"J06": {
		Desc:    "Un salto laterale per gambe, anche, equilibrio e coordinazione.",
		HowTo:   []string{"Stai su una gamba sola.", "Salta di lato sull'altra gamba.", "Atterra morbido e stabilizzati.", "Ripeti da un lato all'altro.", "Usa le braccia per l'equilibrio."},
		Correct: []string{"Atterra in modo silenzioso.", "Lascia che il ginocchio segua le punte.", "Tieni il busto controllato.", "Esegui il movimento laterale in modo fluido."},
		Wrong:   []string{"Non atterrare con il ginocchio che cede verso l'interno.", "Non saltare troppo lontano troppo presto.", "Non lasciarti cadere nell'atterraggio.", "Non torcere il ginocchio al contatto."},
	},
	"J07": {
		Desc:    "Un esercizio cardio che alza la frequenza cardiaca e attiva anche, core e gambe.",
		HowTo:   []string{"Stai dritto.", "Corri sul posto sollevando le ginocchia in alto.", "Muovi le braccia in modo naturale.", "Mantieni un ritmo rapido ma controllato."},
		Correct: []string{"Solleva le ginocchia verso l'altezza dell'anca.", "Atterra in modo leggero.", "Tieni il core attivo.", "Tieni il busto eretto."},
		Wrong:   []string{"Non inclinarti all'indietro.", "Non atterrare in modo pesante.", "Non sbattere i piedi a terra.", "Non perdere la postura quando il ritmo aumenta."},
	},
	"J08": {
		Desc:    "Un esercizio di condizionamento per tutto il corpo che unisce squat, plank e salto. In questo programma, usa la versione senza piegamento, salvo quando è specificamente previsto un piegamento.",
		HowTo:   []string{"Stai dritto.", "Scendi in squat e appoggia le mani a terra.", "Salta o porta i piedi indietro in plank.", "Salta o riporta i piedi in avanti.", "Rialzati e concludi con un piccolo salto."},
		Correct: []string{"Appoggia le mani sotto le spalle.", "Tieni la posizione di plank solida.", "Fai atterrare i piedi in modo morbido.", "Esegui il salto in modo controllato."},
		Wrong:   []string{"Non far cedere la zona lombare nel plank.", "Non atterrare in modo pesante.", "Non affrettare ripetizioni trascurate.", "Non far cedere le ginocchia verso l'interno saltando in avanti."},
	},
	"J09": {
		Desc:    "Un esercizio pliometrico avanzato per le gambe, per potenza, condizionamento e coordinazione.",
		HowTo:   []string{"Parti in posizione di affondo.", "Salta verso l'alto.", "Cambia gamba in aria.", "Atterra morbido nell'affondo opposto.", "Ritrova l'equilibrio prima della ripetizione successiva."},
		Correct: []string{"Atterra in modo morbido e controllato.", "Lascia che il ginocchio anteriore segua le punte.", "Tieni il busto eretto.", "Mantieni l'ampiezza del movimento sicura."},
		Wrong:   []string{"Non atterrare in modo duro.", "Non far cedere il ginocchio anteriore verso l'interno.", "Non scendere troppo in profondità.", "Non muoverti più in fretta di quanto riesci a controllare."},
	},
	"CD07": {
		Desc:    "Un allungamento riposante per schiena, fianchi e spalle per defaticarsi.",
		HowTo:   []string{"Inginocchiati e porta i fianchi indietro verso i talloni.", "Allunga le braccia in avanti sul pavimento.", "Lascia riposare la fronte a terra.", "Respira lentamente e rilassati."},
		Correct: []string{"Abbassa i fianchi verso i talloni.", "Allunga la schiena dolcemente.", "Rilassa le spalle.", "Respira lentamente."},
		Wrong:   []string{"Non forzare i fianchi verso il basso.", "Non tendere le spalle.", "Non trattenere il respiro.", "Non sforzare le ginocchia."},
	},
}
