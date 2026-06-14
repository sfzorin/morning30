package content

// detailsIT is the Italian rich content for every exercise (current library IDs).
var detailsIT = map[string]Detail{
	// ---- Riscaldamento ----
	"W01": {
		Desc:    "Un semplice riscaldamento delle spalle per preparare collo, spalle e parte alta della schiena.",
		HowTo:   []string{"Stai dritto, braccia rilassate.", "Fai lenti cerchi con le spalle all'indietro.", "Poi lentamente in avanti.", "Il movimento resta morbido e controllato."},
		Correct: []string{"Collo rilassato.", "Le spalle si muovono in modo fluido.", "Braccia sciolte.", "Corpo eretto."},
		Wrong:   []string{"Sollevare con forza le spalle.", "Muoversi troppo veloce.", "Tendere il collo.", "Incurvare la parte alta della schiena."},
	},
	"W02": {
		Desc:    "Un riscaldamento di spalle e parte alta della schiena che attiva i deltoidi e gli stabilizzatori della spalla.",
		HowTo:   []string{"Stai dritto e alza le braccia all'altezza delle spalle.", "Fai piccoli cerchi in avanti.", "Poi piccoli cerchi all'indietro.", "Mantieni i cerchi piccoli e controllati."},
		Correct: []string{"Braccia all'altezza delle spalle.", "I cerchi sono piccoli.", "Spalle basse.", "Collo rilassato."},
		Wrong:   []string{"Fare grandi oscillazioni.", "Alzare le spalle alle orecchie.", "Inarcare la zona lombare.", "Trattenere il respiro."},
	},
	"W03": {
		Desc:    "Un esercizio di rotazione dolce per il tronco e la parte alta della colonna.",
		HowTo:   []string{"Piedi alla larghezza dei fianchi.", "I fianchi restano soprattutto rivolti in avanti.", "Ruota la parte alta del corpo da un lato.", "Torna al centro e ruota dall'altro lato."},
		Correct: []string{"I piedi restano stabili.", "Il movimento viene dal tronco.", "La rotazione è fluida.", "Senza torcere le ginocchia."},
		Wrong:   []string{"Ruotare le ginocchia con il corpo.", "Ruotare troppo veloce.", "Inclinarsi all'indietro.", "Forzare l'ampiezza."},
	},
	"W04": {
		Desc:    "Un esercizio di mobilità per fianchi, femorali e controllo della zona lombare.",
		HowTo:   []string{"Stai dritto, piedi alla larghezza dei fianchi.", "Ammorbidisci leggermente le ginocchia.", "Spingi i fianchi all'indietro.", "Abbassa il busto con la schiena dritta.", "Torna su spingendo i fianchi in avanti."},
		Correct: []string{"La schiena resta dritta.", "Il movimento parte dai fianchi.", "Ginocchia morbide, poco piegate.", "Il collo segue la linea della colonna."},
		Wrong:   []string{"Incurvare la schiena.", "Trasformarlo in uno squat.", "Scendere troppo.", "Alzare troppo la testa."},
	},
	"W05": {
		Desc:    "Un riscaldamento delle gambe controllato in un'ampiezza sicura e poco profonda.",
		HowTo:   []string{"Piedi poco più larghi dei fianchi.", "Spingi leggermente i fianchi all'indietro.", "Piega le ginocchia solo fino a una profondità comoda.", "Rialzati lentamente."},
		Correct: []string{"Ginocchia in direzione delle dita.", "Talloni a terra.", "Schiena lunga.", "Profondità moderata."},
		Wrong:   []string{"Scendere troppo.", "Lasciare cedere le ginocchia verso l'interno.", "Alzare i talloni.", "Scendere troppo veloce."},
	},
	"W06": {
		Desc:    "Un riscaldamento per spalle, polsi e core.",
		HowTo:   []string{"Inizia in plank alto.", "Mani sotto le spalle.", "Sposta il corpo leggermente in avanti.", "Torna alla posizione iniziale.", "Il movimento è piccolo e controllato."},
		Correct: []string{"Corpo in una linea.", "Spalle stabili.", "Core attivo.", "Movimento fluido."},
		Wrong:   []string{"Lasciar cadere i fianchi.", "Spingere le spalle verso le orecchie.", "Andare troppo in avanti.", "Lasciare cedere la zona lombare."},
	},

	// ---- Addome / core ----
	"C01": {
		Desc:    "Un esercizio di stabilità del core per addome, spalle, glutei e tronco.",
		HowTo:   []string{"Metti i gomiti sotto le spalle.", "Porta i piedi indietro.", "Forma una linea retta dalle spalle ai talloni.", "Contrai addome e glutei.", "Mantieni la posizione."},
		Correct: []string{"Corpo in una linea.", "Gomiti sotto le spalle.", "Core attivo.", "Respiro regolare."},
		Wrong:   []string{"Lasciar cedere la zona lombare.", "Alzare troppo i fianchi.", "Guardare avanti.", "Trattenere il respiro."},
	},
	"C02": {
		Desc:    "Un plank breve e intenso incentrato sulla tensione di tutto il corpo.",
		HowTo:   []string{"Inizia in plank sugli avambracci.", "Contrai forte l'addome.", "Stringi i glutei.", "Immagina di tirare i gomiti verso le dita dei piedi.", "Mantieni con il massimo controllo."},
		Correct: []string{"Tensione di tutto il corpo.", "Glutei attivi.", "Core saldo.", "Respiro controllato."},
		Wrong:   []string{"Tenerlo come un plank rilassato.", "Lasciar cedere i fianchi.", "Trattenere il respiro.", "Provare a tenere troppo a lungo."},
	},
	"C03": {
		Desc:    "Un esercizio per il basso addome con movimenti incrociati controllati delle gambe.",
		HowTo:   []string{"Sdraiati sulla schiena.", "Premi delicatamente la zona lombare a terra.", "Solleva le gambe.", "Incrocia una gamba sopra l'altra.", "Continua alternando l'incrocio."},
		Correct: []string{"Zona lombare vicina al pavimento.", "Gambe controllate.", "Collo rilassato.", "Movimento fluido."},
		Wrong:   []string{"Inarcare la zona lombare.", "Abbassare troppo le gambe.", "Tirare la testa in avanti.", "Muoversi troppo veloce."},
	},
	"C04": {
		Desc:    "Un esercizio per il basso addome con brevi colpi alternati delle gambe.",
		HowTo:   []string{"Sdraiati sulla schiena.", "Premi la zona lombare a terra.", "Solleva le gambe.", "Muovi le gambe su e giù con brevi colpi alternati."},
		Correct: []string{"Ampiezza ridotta.", "Zona lombare stabile.", "Le gambe non scendono troppo.", "Respiro regolare."},
		Wrong:   []string{"Inarcare la zona lombare.", "Fare colpi ampi.", "Tendere il collo.", "Trattenere il respiro."},
	},
	"C05": {
		Desc:    "Un esercizio addominale incentrato sul sollevamento del bacino con controllo.",
		HowTo:   []string{"Sdraiati sulla schiena.", "Piega le ginocchia e solleva i piedi.", "Usa l'addome per arrotolare leggermente il bacino verso l'alto.", "Riabbassa lentamente."},
		Correct: []string{"Il movimento viene dal basso addome.", "Il bacino sale con controllo.", "Collo rilassato.", "La discesa è lenta."},
		Wrong:   []string{"Oscillare le gambe.", "Usare lo slancio.", "Tirare il collo.", "Lasciar cadere il bacino veloce."},
	},
	"C06": {
		Desc:    "Una tenuta statica del core con leva accorciata per un controllo migliore.",
		HowTo:   []string{"Sdraiati sulla schiena.", "Premi la zona lombare a terra.", "Solleva leggermente le spalle.", "Ginocchia piegate o gambe più in alto.", "Mantieni la posizione."},
		Correct: []string{"Zona lombare a terra.", "Addome saldo.", "Collo rilassato.", "Il respiro continua."},
		Wrong:   []string{"Inarcare la zona lombare.", "Abbassare troppo le gambe.", "Portare il mento in avanti.", "Trattenere il respiro."},
	},
	"C07": {
		Desc:    "Un esercizio laterale per il core, per gli obliqui e la stabilità di spalla e anca.",
		HowTo:   []string{"Sdraiati su un fianco.", "Metti il gomito sotto la spalla.", "Solleva i fianchi.", "Il corpo in una linea retta.", "Mantieni."},
		Correct: []string{"Gomito sotto la spalla.", "Fianchi sollevati.", "Corpo lungo.", "Collo neutro."},
		Wrong:   []string{"Lasciar cadere i fianchi.", "Avvicinare la spalla all'orecchio.", "Ruotare il petto in avanti.", "Trattenere il respiro."},
	},
	"C08": {
		Desc:    "Una variante più impegnativa del plank laterale per gli obliqui e l'anca laterale.",
		HowTo:   []string{"Inizia in plank laterale.", "Abbassa leggermente i fianchi.", "Rialza i fianchi.", "Ripeti con controllo.", "Cambia lato."},
		Correct: []string{"Movimento piccolo e controllato.", "Gomito sotto la spalla.", "I fianchi salgono in verticale.", "Core saldo."},
		Wrong:   []string{"Scendere troppo.", "Torcere il tronco.", "Cedere nella spalla.", "Muoversi troppo veloce."},
	},
	"C10": {
		Desc:    "Un esercizio di core e tricipiti che alterna plank sugli avambracci e plank alto.",
		HowTo:   []string{"Inizia in plank sugli avambracci.", "Appoggia una mano a terra e spingi su.", "Porta l'altra mano in plank alto.", "Riscendi sugli avambracci.", "Alterna la mano che guida."},
		Correct: []string{"Fianchi stabili.", "Core attivo.", "Movimento controllato.", "Spalle sopra mani o gomiti."},
		Wrong:   []string{"Oscillare i fianchi da lato a lato.", "Lasciar cedere la zona lombare.", "Muoversi troppo veloce.", "Lasciar cedere le spalle."},
	},
	"C11": {
		Desc:    "Un esercizio per il core controllato e un'alternativa più sicura ai sollevamenti di gambe più impegnativi.",
		HowTo:   []string{"Sdraiati sulla schiena.", "Solleva le gambe in posizione a tavolino.", "Abbassa un tallone verso il pavimento.", "Tocca leggermente e torna.", "Alterna i lati."},
		Correct: []string{"Zona lombare stabile.", "Movimento lento.", "Il tallone tocca delicatamente.", "Bacino fermo."},
		Wrong:   []string{"Lasciar cadere la gamba troppo veloce.", "Inarcare la zona lombare.", "Muovere entrambe le gambe insieme.", "Affrettarsi."},
	},
	"C12": {
		Desc:    "Un esercizio di controllo del core per l'addome profondo e la stabilità della colonna.",
		HowTo:   []string{"Sdraiati sulla schiena.", "Solleva braccia e gambe in posizione a tavolino.", "Estendi il braccio e la gamba opposti.", "Mantieni stabile la zona lombare.", "Torna e cambia lato."},
		Correct: []string{"Zona lombare vicina al pavimento.", "Movimento lento.", "Il braccio e la gamba opposti si muovono insieme.", "Il tronco resta fermo."},
		Wrong:   []string{"Inarcare la zona lombare.", "Muoversi troppo veloce.", "Abbassare troppo la gamba.", "Far sporgere le costole."},
	},

	// ---- Piegamenti / tricipiti / spalle ----
	"P01": {
		Desc:    "Un esercizio base per la parte alta del corpo: petto, tricipiti, spalle e core.",
		HowTo:   []string{"Inizia in plank alto.", "Mani sotto le spalle o poco più larghe.", "Abbassa il corpo con controllo.", "Spingi su mantenendo il corpo dritto."},
		Correct: []string{"Corpo in una linea.", "Core saldo.", "Gomiti a circa 30–45 gradi.", "Movimento controllato."},
		Wrong:   []string{"Lasciar cadere i fianchi.", "Alzare troppo i fianchi.", "Aprire i gomiti dritti ai lati.", "Cadere nella posizione bassa."},
	},
	"P02": {
		Desc:    "Una variante di piegamento con più lavoro sui tricipiti.",
		HowTo:   []string{"Inizia in plank alto.", "Metti le mani più vicine rispetto a un piegamento normale.", "Tieni i gomiti vicini al corpo.", "Scendi lentamente e spingi su."},
		Correct: []string{"Gomiti vicini.", "Corpo dritto.", "Spalle lontane dalle orecchie.", "Polsi comodi."},
		Wrong:   []string{"Mettere le mani troppo vicine.", "Aprire i gomiti.", "Lasciar cedere la zona lombare.", "Forzare nonostante dolore a polso o gomito."},
	},
	"P03": {
		Desc:    "Piegamenti controllati con una breve pausa nella posizione bassa.",
		HowTo:   []string{"Inizia in posizione di piegamento normale.", "Scendi lentamente.", "Pausa un secondo vicino al fondo.", "Spingi su senza perdere la posizione del corpo."},
		Correct: []string{"Pausa controllata.", "Core saldo.", "Petto attivo.", "Il corpo resta dritto."},
		Wrong:   []string{"Rilassarsi in basso.", "Lasciar cadere i fianchi.", "Trattenere il respiro.", "Rimbalzare dal fondo."},
	},
	"P04": {
		Desc:    "Una variante lenta di piegamento che aumenta controllo e tempo sotto tensione.",
		HowTo:   []string{"Inizia in plank alto.", "Scendi per circa tre secondi.", "Pausa un secondo vicino al fondo.", "Spingi su per circa tre secondi."},
		Correct: []string{"Ritmo lento e uniforme.", "Corpo dritto.", "Gomiti controllati.", "Respiro regolare."},
		Wrong:   []string{"Affrettare il movimento.", "Crollare in basso.", "Lasciar cadere i fianchi.", "Perdere la tensione del core."},
	},
	"P05": {
		Desc:    "Una variante asimmetrica di piegamento che impegna petto, spalle e core.",
		HowTo:   []string{"Inizia in posizione di piegamento.", "Metti una mano un po' avanti e l'altra un po' indietro.", "Scendi con controllo.", "Spingi su.", "Cambia la posizione delle mani dall'altro lato."},
		Correct: []string{"Il corpo resta dritto.", "Core saldo.", "Entrambe le spalle stabili.", "Movimento controllato."},
		Wrong:   []string{"Torcere il tronco.", "Lasciar cadere una spalla.", "Allargare troppo le mani.", "Forzare l'ampiezza se la spalla dà fastidio."},
	},
	"P06": {
		Desc:    "Una variante di piegamento incentrata sulle spalle.",
		HowTo:   []string{"Inizia in plank alto.", "Avvicina un po' i piedi e solleva i fianchi.", "La testa tra le braccia.", "Abbassa la testa verso il pavimento.", "Spingi di nuovo su."},
		Correct: []string{"Fianchi alti.", "I gomiti si piegano con controllo.", "Spalle stabili.", "Collo neutro."},
		Wrong:   []string{"Cadere in un piegamento normale.", "Portare la testa troppo in avanti.", "Aprire troppo i gomiti.", "Forzare nonostante dolore alla spalla."},
	},
	"P08": {
		Desc:    "Un esercizio per tricipiti e spalle da una posizione di appoggio sugli avambracci.",
		HowTo:   []string{"Inizia sugli avambracci, il corpo lungo.", "Gomiti sotto o un po' davanti alle spalle.", "Spingi con mani e avambracci.", "Estendi leggermente i gomiti.", "Torna con controllo."},
		Correct: []string{"Core saldo.", "I gomiti si muovono in modo fluido.", "Spalle basse.", "Ampiezza comoda."},
		Wrong:   []string{"Forzare i gomiti.", "Lasciar cadere i fianchi.", "Alzare le spalle.", "Muoversi nonostante dolore a gomito o spalla."},
	},

	// ---- Schiena / postura ----
	"B01": {
		Desc:    "Un esercizio per la schiena incentrato sulla postura, per la parte alta della schiena, le scapole e gli estensori della colonna.",
		HowTo:   []string{"Sdraiati a pancia in giù.", "Le braccia lungo il corpo.", "Ruota leggermente i pollici verso l'esterno.", "Solleva un po' il petto.", "Porta le scapole indietro e in basso.", "Mantieni."},
		Correct: []string{"Il sollevamento è piccolo.", "Collo neutro.", "Scapole indietro e in basso.", "Glutei leggermente attivi."},
		Wrong:   []string{"Sollevare troppo.", "Buttare indietro la testa.", "Iperestendere la zona lombare.", "Alzare le spalle alle orecchie."},
	},
	"B02": {
		Desc:    "Un esercizio per la parte alta della schiena e le scapole a pancia in giù.",
		HowTo:   []string{"Sdraiati sulla pancia.", "La testa guarda in basso.", "Muovi le braccia in un ampio arco dai lati verso l'alto.", "Torna per lo stesso percorso.", "Movimento controllato."},
		Correct: []string{"Collo rilassato.", "Le scapole si muovono in modo fluido.", "Il sollevamento del petto è piccolo o neutro.", "Le braccia senza fretta."},
		Wrong:   []string{"Sollevare la testa.", "Iperestendere troppo la zona lombare.", "Oscillare le braccia.", "Alzare le spalle alle orecchie."},
	},
	"B03": {
		Desc:    "Un esercizio per la parte alta della schiena, per le scapole e le spalle posteriori.",
		HowTo:   []string{"Sdraiati a pancia in giù.", "Piega i gomiti a forma di W.", "Solleva un po' gomiti e mani.", "Stringi le scapole.", "Riabbassa con controllo."},
		Correct: []string{"Ampiezza ridotta.", "Il lavoro lo fanno le scapole.", "Collo neutro.", "Spalle lontane dalle orecchie."},
		Wrong:   []string{"Sollevare troppo.", "Sollevare la testa.", "Usare lo slancio.", "Iperestendere la zona lombare."},
	},
	"B04": {
		Desc:    "Un esercizio per le scapole, per il trapezio inferiore e la parte alta della schiena.",
		HowTo:   []string{"Sdraiati a pancia in giù.", "Estendi le braccia in diagonale verso l'alto a forma di Y.", "Solleva un po' le braccia dal pavimento.", "Riabbassa con controllo."},
		Correct: []string{"Le braccia salgono solo un po'.", "Collo neutro.", "Spalle basse.", "Movimento controllato."},
		Wrong:   []string{"Sollevare troppo.", "Alzare le spalle alle orecchie.", "Inarcare la zona lombare.", "Muoversi troppo veloce."},
	},
	"B05": {
		Desc:    "Un esercizio per la parte alta della schiena, per le spalle posteriori e il controllo delle scapole.",
		HowTo:   []string{"Sdraiati a pancia in giù.", "Estendi le braccia ai lati a forma di T.", "Solleva un po' le braccia.", "Stringi le scapole.", "Riabbassa lentamente."},
		Correct: []string{"Le braccia sulla linea delle spalle.", "Collo rilassato.", "Le scapole si muovono dolcemente.", "Il sollevamento è piccolo."},
		Wrong:   []string{"Lanciare le braccia verso l'alto.", "Sollevare la testa.", "Alzare le spalle.", "Forzare un'ampiezza ampia."},
	},
	"B06": {
		Desc:    "Un esercizio per la schiena che imita un movimento di tirata senza attrezzi.",
		HowTo:   []string{"Sdraiati a pancia in giù con le braccia sopra la testa.", "Solleva un po' il petto.", "Tira i gomiti in basso verso le costole.", "Stringi le scapole.", "Estendi di nuovo le braccia in avanti."},
		Correct: []string{"I gomiti scendono con controllo.", "Le scapole si stringono dolcemente.", "Collo neutro.", "Zona lombare comoda."},
		Wrong:   []string{"Dare strattoni con le braccia.", "Sollevare troppo.", "Buttare indietro la testa.", "Alzare le spalle alle orecchie."},
	},
	"B07": {
		Desc:    "Un esercizio per la schiena e la coordinazione a pancia in giù.",
		HowTo:   []string{"Sdraiati sulla pancia.", "Solleva un po' il petto.", "Muovi braccio e gamba opposti in uno schema di nuoto controllato.", "Il movimento è piccolo e costante."},
		Correct: []string{"Collo neutro.", "Movimento controllato.", "Zona lombare comoda.", "Il respiro continua."},
		Wrong:   []string{"Muoversi troppo veloce.", "Sollevare troppo.", "Buttare indietro la testa.", "Forzare la zona lombare."},
	},
	"B08": {
		Desc:    "Un esercizio per la catena posteriore: glutei, femorali, spalle e schiena.",
		HowTo:   []string{"Siediti a terra con le gambe distese.", "Metti le mani dietro i fianchi.", "Spingi con mani e talloni.", "Solleva i fianchi.", "Tieni il corpo in una linea lunga."},
		Correct: []string{"Petto aperto.", "Fianchi sollevati.", "Spalle stabili.", "Collo neutro."},
		Wrong:   []string{"Lasciar cadere i fianchi.", "Alzare le spalle.", "Buttare indietro la testa.", "Forzare i polsi."},
	},
	"B09": {
		Desc:    "Un esercizio per la schiena di piccola ampiezza per gli estensori della colonna e la postura.",
		HowTo:   []string{"Sdraiati a pancia in giù.", "Le mani lungo il corpo o leggermente vicino al petto.", "Solleva un po' il petto.", "Fai piccole pulsazioni controllate.", "Collo neutro."},
		Correct: []string{"Sollevamento piccolo.", "Pulsazioni dolci.", "Collo lungo.", "Zona lombare controllata."},
		Wrong:   []string{"Pulsare troppo in alto.", "Usare lo slancio.", "Guardare avanti.", "Sentire dolore acuto nella zona lombare."},
	},
	"B10": {
		Desc:    "Un esercizio per la schiena più impegnativo che unisce un piccolo sollevamento a cobra con una tirata delle braccia a forma di W.",
		HowTo:   []string{"Sdraiati a pancia in giù con le braccia in avanti.", "Solleva un po' il petto.", "Tira i gomiti in basso e indietro a forma di W.", "Stringi le scapole.", "Estendi di nuovo in avanti e riabbassa con controllo."},
		Correct: []string{"Piccolo sollevamento del petto.", "Tirata a W fluida.", "Scapole indietro e in basso.", "Collo lungo."},
		Wrong:   []string{"Iperestendere la zona lombare.", "Tirare con il collo.", "Muoversi troppo veloce.", "Alzare le spalle."},
	},

	// ---- Braccia / bicipiti ----
	"A01": {
		Desc:    "Un esercizio per i bicipiti senza attrezzi: una mano oppone resistenza all'altra.",
		HowTo:   []string{"In piedi o seduto, ben dritto.", "Piega un gomito come in un curl.", "Con l'altra mano premi verso il basso sull'avambraccio.", "Sali lentamente contro la resistenza.", "Cambia lato."},
		Correct: []string{"Resistenza costante.", "Movimento lento.", "Il gomito vicino al corpo.", "La spalla bassa."},
		Wrong:   []string{"Muoversi senza vera resistenza.", "Dare strattoni col braccio.", "Alzare la spalla.", "Trattenere il respiro."},
	},
	"A02": {
		Desc:    "Una tenuta statica per i bicipiti con autoresistenza.",
		HowTo:   []string{"Piega un gomito a circa 90 gradi.", "Con l'altra mano premi verso il basso.", "Il braccio che lavora resiste senza muoversi.", "Mantieni la tensione.", "Cambia lato."},
		Correct: []string{"L'angolo del gomito resta stabile.", "Tensione costante.", "Spalle rilassate.", "Il respiro continua."},
		Wrong:   []string{"Premere a strattoni.", "Torcere il tronco.", "Alzare la spalla.", "Trattenere il respiro."},
	},
	"A03": {
		Desc:    "Un esercizio isometrico per braccia e avambracci, le mani una contro l'altra.",
		HowTo:   []string{"Aggancia le dita o le mani.", "I gomiti leggermente piegati.", "Tira le mani per allontanarle.", "Mantieni una tensione costante.", "Rilascia lentamente."},
		Correct: []string{"Tensione controllata.", "Spalle basse.", "Polsi comodi.", "Respiro calmo."},
		Wrong:   []string{"Tirare di scatto.", "Alzare le spalle.", "Piegare troppo i polsi.", "Trattenere il respiro."},
	},
	"A04": {
		Desc:    "Un esercizio di discesa lenta per i bicipiti con autoresistenza.",
		HowTo:   []string{"Inizia con un gomito piegato.", "Con l'altra mano crea resistenza.", "Abbassa lentamente l'avambraccio.", "Mantieni la resistenza durante la discesa.", "Cambia lato."},
		Correct: []string{"La discesa è lenta.", "Resistenza costante.", "Il gomito vicino al corpo.", "La spalla rilassata."},
		Wrong:   []string{"Lasciar cadere il braccio veloce.", "Non usare resistenza.", "Torcere il tronco.", "Trattenere il respiro."},
	},

	// ---- Gambe / glutei ----
	"L01": {
		Desc:    "Un esercizio per le gambe controllato per cosce, fianchi e postura.",
		HowTo:   []string{"Piedi poco più larghi dei fianchi.", "Spingi i fianchi all'indietro.", "Piega le ginocchia fino a una profondità comoda e poco profonda.", "Rialzati lentamente."},
		Correct: []string{"Ginocchia in direzione delle dita.", "Talloni a terra.", "Schiena lunga.", "Profondità comoda."},
		Wrong:   []string{"Scendere troppo.", "Lasciare cedere le ginocchia verso l'interno.", "Alzare i talloni.", "Scendere veloce."},
	},
	"L02": {
		Desc:    "Una variante di squat controllata con una breve pausa in un'ampiezza sicura.",
		HowTo:   []string{"In piedi e stabile.", "Scendi in uno squat comodo e poco profondo.", "Pausa breve.", "Rialzati con controllo."},
		Correct: []string{"Pausa stabile.", "Ginocchia in direzione delle dita.", "Petto aperto.", "Talloni a terra."},
		Wrong:   []string{"Scendere troppo.", "Rimbalzare nella posizione bassa.", "Lasciare cedere le ginocchia verso l'interno.", "Trattenere il respiro."},
	},
	"L03": {
		Desc:    "Un esercizio di mobilità per fianchi, femorali e controllo della zona lombare.",
		HowTo:   []string{"Stai dritto, piedi alla larghezza dei fianchi.", "Ammorbidisci leggermente le ginocchia.", "Spingi i fianchi all'indietro.", "Abbassa il busto con la schiena dritta.", "Torna su spingendo i fianchi in avanti."},
		Correct: []string{"La schiena resta dritta.", "Il movimento parte dai fianchi.", "Ginocchia morbide, poco piegate.", "Il collo segue la linea della colonna."},
		Wrong:   []string{"Incurvare la schiena.", "Trasformarlo in uno squat.", "Scendere troppo.", "Alzare troppo la testa."},
	},
	"L04": {
		Desc:    "Un esercizio per i glutei più impegnativo, una gamba alla volta.",
		HowTo:   []string{"Sdraiati sulla schiena.", "Piega un ginocchio con quel piede a terra.", "Estendi o solleva l'altra gamba.", "Spingi con il tallone che lavora.", "Solleva i fianchi e riabbassa lentamente."},
		Correct: []string{"Il sollevamento viene dal gluteo che lavora.", "I fianchi restano a livello.", "Movimento controllato.", "Zona lombare neutra."},
		Wrong:   []string{"Torcere il bacino.", "Spingere sulle dita dei piedi.", "Inarcare la zona lombare.", "Scendere troppo veloce."},
	},
	"L05": {
		Desc:    "Un esercizio per i glutei e la stabilità del bacino.",
		HowTo:   []string{"Inizia in posizione di ponte glutei.", "Tieni i fianchi sollevati.", "Solleva un piede leggermente da terra.", "Riabbassalo e cambia lato.", "Tieni il bacino stabile."},
		Correct: []string{"I fianchi restano a livello.", "Glutei attivi.", "Movimento lento.", "Zona lombare comoda."},
		Wrong:   []string{"Lasciar cadere i fianchi.", "Oscillare da lato a lato.", "Muoversi troppo veloce.", "Inarcare la zona lombare."},
	},
	"L06": {
		Desc:    "Un esercizio per i glutei che rafforza fianchi e catena posteriore.",
		HowTo:   []string{"Sdraiati sulla schiena.", "Piega le ginocchia, piedi a terra.", "Solleva i fianchi.", "Pausa in alto.", "Riabbassa lentamente."},
		Correct: []string{"Spingi con i talloni.", "I glutei si stringono in alto.", "Le costole in basso.", "La zona lombare non si iperestende."},
		Wrong:   []string{"Iperestendere la zona lombare in alto.", "Alzare i talloni.", "Lasciar cadere i fianchi troppo veloce.", "Mettere i piedi troppo vicini."},
	},
	"L07": {
		Desc:    "Un esercizio per la parte bassa della gamba per polpacci, caviglie e controllo del piede.",
		HowTo:   []string{"In piedi, piedi alla larghezza dei fianchi.", "Sali sulle punte dei piedi.", "Pausa breve in alto.", "Abbassa i talloni lentamente."},
		Correct: []string{"Corpo eretto.", "Movimento fluido.", "Caviglie allineate.", "Discesa controllata."},
		Wrong:   []string{"Lasciar cadere i talloni veloce.", "Oscillare in avanti.", "Far cadere le caviglie verso l'esterno.", "Usare lo slancio."},
	},
	"L09": {
		Desc:    "Un esercizio per la catena laterale per glutei laterali, obliqui e stabilità dell'anca.",
		HowTo:   []string{"Inizia in plank laterale.", "Il corpo in una linea retta.", "Solleva un po' la gamba in alto.", "Riabbassala con controllo.", "Ripeti e cambia lato."},
		Correct: []string{"Fianchi sollevati.", "La gamba in alto si muove lentamente.", "Tronco stabile.", "Spalla salda."},
		Wrong:   []string{"Lasciar cadere i fianchi.", "Oscillare la gamba.", "Ruotare il corpo all'indietro.", "Cedere nella spalla."},
	},

	// ---- Defaticamento ----
	"CD01": {
		Desc:    "Un allungamento dolce per il petto e la parte anteriore delle spalle.",
		HowTo:   []string{"Sdraiati sulla schiena.", "Apri le braccia ai lati.", "Rilassa le spalle.", "Lascia che il petto si apra dolcemente.", "Respira lentamente."},
		Correct: []string{"Il petto si apre dolcemente.", "Spalle rilassate.", "Zona lombare comoda.", "Nessun dolore alle spalle."},
		Wrong:   []string{"Forzare le braccia verso il basso.", "Inarcare la zona lombare.", "Tendere il collo.", "Allungare nonostante il dolore."},
	},
	"CD02": {
		Desc:    "Un allungamento dolce della parte posteriore della coscia.",
		HowTo:   []string{"Sdraiati sulla schiena.", "Solleva una gamba.", "Tieni dietro la coscia o il polpaccio.", "Porta la gamba dolcemente verso di te.", "Cambia lato."},
		Correct: []string{"L'allungamento è leggero.", "Il ginocchio può restare un po' piegato.", "Zona lombare tranquilla.", "Respiro lento."},
		Wrong:   []string{"Tirare troppo forte.", "Forzare la gamba dritta.", "Sollevare i fianchi.", "Sentire dolore acuto dietro il ginocchio."},
	},
	"CD03": {
		Desc:    "Un allungamento dolce di rotazione per la colonna e il tronco.",
		HowTo:   []string{"Sdraiati sulla schiena.", "Piega le ginocchia.", "Lascia cadere le ginocchia dolcemente da un lato.", "Tieni le spalle a terra.", "Torna al centro e cambia lato."},
		Correct: []string{"Movimento lento.", "Spalle a terra.", "Respiro rilassato.", "Allungamento comodo."},
		Wrong:   []string{"Forzare le ginocchia a terra.", "Ruotare veloce.", "Sollevare la spalla.", "Allungare nonostante il dolore."},
	},
	"CD04": {
		Desc:    "Un allungamento dolce della parte anteriore del corpo e una lieve estensione della schiena.",
		HowTo:   []string{"Sdraiati a pancia in giù.", "Metti gli avambracci a terra.", "Solleva il petto dolcemente.", "Tieni il collo lungo.", "Respira lentamente."},
		Correct: []string{"Il sollevamento è lieve.", "Spalle lontane dalle orecchie.", "Zona lombare comoda.", "Collo neutro."},
		Wrong:   []string{"Iperestendere la zona lombare.", "Buttare indietro la testa.", "Alzare le spalle.", "Forzare la posizione."},
	},
	"CD05": {
		Desc:    "Un esercizio di respirazione calma per concludere la sessione.",
		HowTo:   []string{"Sdraiati sulla schiena.", "Rilassa spalle e mascella.", "Se comodo, metti una mano sulla pancia.", "Inspira dolcemente.", "Espira lentamente."},
		Correct: []string{"Respiro calmo.", "Spalle rilassate.", "Viso disteso.", "Il corpo si calma."},
		Wrong:   []string{"Respirare troppo forte.", "Trattenere il respiro.", "Tendere il collo.", "Inarcare la zona lombare."},
	},

	// ---- Aggiunte del set Vlad (riscaldamento / cardio / pliometria / affondi) ----
	"W07": {
		Desc:    "Un esercizio di apertura calmo: respirazione profonda con un dolce allungamento di tutto il corpo per allungare la colonna.",
		HowTo:   []string{"Stai dritto, piedi alla larghezza dei fianchi.", "Inspira e porta entrambe le braccia sopra la testa.", "Allungati dolcemente lungo la colonna.", "Espira e abbassa le braccia lasciando cadere le spalle."},
		Correct: []string{"Il respiro è lento e pieno.", "Le spalle si rilassano sull'espirazione.", "Le costole restano basse, senza iperestendere.", "Il movimento è senza fretta."},
		Wrong:   []string{"Trattenere il respiro.", "Inarcare con forza la zona lombare.", "Alzare le spalle alle orecchie.", "Affrettare l'allungamento."},
	},
	"W08": {
		Desc:    "Un esercizio di rotazione dolce per riscaldare la colonna e il tronco.",
		HowTo:   []string{"Piedi alla larghezza dei fianchi.", "Lascia oscillare le braccia in modo sciolto.", "Ruota la parte alta del corpo da un lato.", "Passa in modo fluido all'altro lato."},
		Correct: []string{"Il movimento viene dal tronco.", "I fianchi restano soprattutto rivolti in avanti.", "I piedi restano stabili.", "Il ritmo è fluido e uniforme."},
		Wrong:   []string{"Torcere con forza le ginocchia.", "Dare strattoni con le braccia.", "Ruotare troppo veloce.", "Trattenere il respiro."},
	},
	"W09": {
		Desc:    "Un riscaldamento dinamico per spalle e petto, per preparare la parte alta del corpo a piegamenti e salti.",
		HowTo:   []string{"Stai dritto, piedi alla larghezza dei fianchi.", "Oscilla le braccia avanti e indietro in modo controllato.", "Lascia che il petto si apra naturalmente quando le braccia vanno indietro.", "Mantieni il movimento fluido e rilassato."},
		Correct: []string{"Le braccia si muovono liberamente.", "Le spalle restano rilassate.", "Il petto si apre dolcemente.", "Il corpo resta eretto."},
		Wrong:   []string{"Oscillare in modo troppo aggressivo.", "Alzare le spalle alle orecchie.", "Inarcare la zona lombare.", "Trasformare il movimento in uno slancio incontrollato."},
	},
	"C13": {
		Desc:    "Un esercizio addominale di rotazione per gli obliqui e l'addome anteriore.",
		HowTo:   []string{"Sdraiati sulla schiena.", "Porta le ginocchia in alto.", "Appoggia le mani leggere vicino alla testa.", "Ruota il tronco e avvicina un gomito al ginocchio opposto.", "Cambia lato con un ritmo controllato."},
		Correct: []string{"La rotazione viene dal tronco.", "La zona lombare resta controllata.", "Il collo resta rilassato.", "Il movimento è fluido."},
		Wrong:   []string{"Tirare il collo.", "Muovere solo i gomiti.", "Affrettare le ripetizioni.", "Lasciare inarcare la zona lombare."},
	},
	"B11": {
		Desc:    "Un esercizio per la schiena che solleva braccio e gamba opposti in un movimento alternato simile al nuoto.",
		HowTo:   []string{"Sdraiati a pancia in giù con le braccia sopra la testa.", "Solleva un po' il petto e le gambe.", "Alza un braccio e la gamba opposta.", "Riabbassa e cambia lato, alternando in modo fluido."},
		Correct: []string{"Il sollevamento resta piccolo.", "Il collo segue la linea della colonna.", "I glutei restano attivi.", "Il movimento è costante."},
		Wrong:   []string{"Forzare la zona lombare.", "Sollevare la testa.", "Muoversi a scatti.", "Trattenere il respiro."},
	},
	"L10": {
		Desc:    "Uno squat base a corpo libero per gambe, anche e glutei.",
		HowTo:   []string{"Stai in piedi, piedi circa alla larghezza delle spalle.", "Spingi leggermente le anche all'indietro.", "Piega le ginocchia e scendi in squat.", "Tieni il petto aperto.", "Rialzati spingendo con tutta la pianta del piede."},
		Correct: []string{"Le ginocchia seguono le punte.", "I talloni restano a terra.", "La schiena resta lunga.", "La profondità resta controllata."},
		Wrong:   []string{"Lasciar cedere le ginocchia verso l'interno.", "Alzare i talloni.", "Incurvare la schiena.", "Scendere troppo in fretta."},
	},
	"L11": {
		Desc:    "Un esercizio per la parte bassa del corpo, per gambe, glutei ed equilibrio.",
		HowTo:   []string{"Stai dritto.", "Porta una gamba indietro.", "Scendi in affondo con controllo.", "Spingi con il piede anteriore per tornare in piedi.", "Cambia lato."},
		Correct: []string{"Il ginocchio anteriore segue le punte.", "Il busto resta eretto.", "Il passo indietro è controllato.", "Il tallone anteriore resta a terra."},
		Wrong:   []string{"Lasciar cedere il ginocchio anteriore verso l'interno.", "Inclinarsi troppo in avanti.", "Scendere in fretta.", "Spingere troppo con la gamba posteriore."},
	},
	"L12": {
		Desc:    "Uno squat a corpo libero eseguito lentamente per costruire controllo e forza nelle gambe.",
		HowTo:   []string{"Piedi alla larghezza delle spalle.", "Scendi in circa tre secondi.", "Pausa breve nella posizione bassa.", "Rialzati con controllo."},
		Correct: []string{"Il ritmo resta lento.", "Ginocchia in direzione delle dita.", "Talloni a terra.", "Schiena neutra."},
		Wrong:   []string{"Scendere veloce.", "Rimbalzare dalla posizione bassa.", "Lasciare cedere le ginocchia verso l'interno.", "Trattenere il respiro."},
	},
	"J01": {
		Desc:    "Un leggero riscaldamento per caviglie e polpacci con piccoli salti rapidi.",
		HowTo:   []string{"Stai dritto, piedi circa alla larghezza dei fianchi.", "Tieni le ginocchia leggermente morbide.", "Fai piccoli salti rapidi usando soprattutto le caviglie.", "Atterra in modo silenzioso e ripeti."},
		Correct: []string{"I salti restano bassi.", "Gli atterraggi sono morbidi e silenziosi.", "Le ginocchia restano leggermente piegate.", "Il corpo resta eretto."},
		Wrong:   []string{"Saltare troppo in alto.", "Atterrare in modo rumoroso.", "Bloccare le ginocchia.", "Far cedere le caviglie verso l'interno o l'esterno."},
	},
	"J02": {
		Desc:    "Un'alternativa a basso impatto alle ginocchia alte.",
		HowTo:   []string{"Stai dritto.", "Solleva un ginocchio verso l'altezza dell'anca.", "Riabbassalo con controllo.", "Cambia lato e continua a marciare.", "Tieni le braccia in movimento naturale."},
		Correct: []string{"Il busto resta eretto.", "Le ginocchia salgono con controllo.", "I piedi atterrano in modo morbido.", "Il respiro resta regolare."},
		Wrong:   []string{"Inclinarsi all'indietro.", "Muovere le gambe senza controllo.", "Sbattere i piedi a terra.", "Trattenere il respiro."},
	},
	"J03": {
		Desc:    "Un esercizio cardio per tutto il corpo che alza la frequenza cardiaca e riscalda spalle, anche e gambe.",
		HowTo:   []string{"Stai dritto, braccia lungo i fianchi.", "Salta allargando i piedi mentre alzi le braccia sopra la testa.", "Salta riportando i piedi insieme mentre abbassi le braccia.", "Ripeti con un ritmo costante."},
		Correct: []string{"Atterra in modo morbido.", "Tieni le ginocchia leggermente piegate.", "Le braccia si muovono in modo fluido.", "Il respiro resta ritmico."},
		Wrong:   []string{"Atterrare in modo duro.", "Bloccare le ginocchia.", "Alzare le spalle alle orecchie.", "Muoversi troppo veloce perdendo il controllo."},
	},
	"J04": {
		Desc:    "Un esercizio dinamico di core e cardio eseguito dal plank alto.",
		HowTo:   []string{"Parti in plank alto.", "Tieni le mani sotto le spalle.", "Porta un ginocchio verso il petto.", "Cambia gamba a ritmo.", "Tieni le anche stabili."},
		Correct: []string{"Il corpo resta in un plank solido.", "Il core resta saldo.", "Le spalle restano sopra le mani.", "Le ginocchia si muovono con controllo."},
		Wrong:   []string{"Lasciar rimbalzare le anche verso l'alto.", "Lasciar cedere la zona lombare.", "Muovere i piedi in modo troppo scomposto.", "Lasciar cedere le spalle."},
	},
	"J05": {
		Desc:    "Un esercizio pliometrico per le gambe, per potenza, condizionamento e coordinazione.",
		HowTo:   []string{"Parti con i piedi circa alla larghezza delle spalle.", "Scendi in uno squat controllato.", "Salta verso l'alto.", "Atterra morbido con le ginocchia leggermente piegate.", "Ritrova la posizione prima del salto successivo."},
		Correct: []string{"Atterra in modo silenzioso.", "Le ginocchia seguono le punte.", "Il petto resta aperto.", "Ogni salto è controllato."},
		Wrong:   []string{"Atterrare in modo duro.", "Lasciar cedere le ginocchia verso l'interno.", "Saltare prima che lo squat sia stabile.", "Affrettarsi senza ritrovare la posizione."},
	},
	"J06": {
		Desc:    "Un salto laterale per gambe, anche, equilibrio e coordinazione.",
		HowTo:   []string{"Stai su una gamba sola.", "Salta di lato sull'altra gamba.", "Atterra morbido e stabilizzati.", "Ripeti da un lato all'altro.", "Usa le braccia per l'equilibrio."},
		Correct: []string{"Gli atterraggi sono silenziosi.", "Il ginocchio segue le punte.", "Il busto resta controllato.", "Il movimento laterale è fluido."},
		Wrong:   []string{"Atterrare con il ginocchio che cede verso l'interno.", "Saltare troppo lontano troppo presto.", "Lasciarsi cadere nell'atterraggio.", "Torcere il ginocchio al contatto."},
	},
	"J07": {
		Desc:    "Un esercizio cardio che alza la frequenza cardiaca e attiva anche, core e gambe.",
		HowTo:   []string{"Stai dritto.", "Corri sul posto sollevando le ginocchia in alto.", "Muovi le braccia in modo naturale.", "Mantieni un ritmo rapido ma controllato."},
		Correct: []string{"Le ginocchia salgono verso l'altezza dell'anca.", "Gli atterraggi restano leggeri.", "Il core resta attivo.", "Il busto resta eretto."},
		Wrong:   []string{"Inclinarsi all'indietro.", "Atterrare in modo pesante.", "Sbattere i piedi a terra.", "Perdere la postura quando il ritmo aumenta."},
	},
	"J08": {
		Desc:    "Un esercizio di condizionamento per tutto il corpo che unisce squat, plank e salto. In questo programma, usa la versione senza piegamento, salvo quando è specificamente previsto un piegamento.",
		HowTo:   []string{"Stai dritto.", "Scendi in squat e appoggia le mani a terra.", "Salta o porta i piedi indietro in plank.", "Salta o riporta i piedi in avanti.", "Rialzati e concludi con un piccolo salto."},
		Correct: []string{"Le mani atterrano sotto le spalle.", "La posizione di plank è solida.", "I piedi atterrano in modo morbido.", "Il salto è controllato."},
		Wrong:   []string{"Lasciar cedere la zona lombare nel plank.", "Atterrare in modo pesante.", "Affrettare ripetizioni trascurate.", "Lasciar cedere le ginocchia verso l'interno saltando in avanti."},
	},
	"J09": {
		Desc:    "Un esercizio pliometrico avanzato per le gambe, per potenza, condizionamento e coordinazione.",
		HowTo:   []string{"Parti in posizione di affondo.", "Salta verso l'alto.", "Cambia gamba in aria.", "Atterra morbido nell'affondo opposto.", "Ritrova l'equilibrio prima della ripetizione successiva."},
		Correct: []string{"L'atterraggio è morbido e controllato.", "Il ginocchio anteriore segue le punte.", "Il busto resta eretto.", "L'ampiezza del movimento resta sicura."},
		Wrong:   []string{"Atterrare in modo duro.", "Far cedere il ginocchio anteriore verso l'interno.", "Scendere troppo in profondità.", "Muoversi più in fretta di quanto riesci a controllare."},
	},
	"CD07": {
		Desc:    "Un allungamento riposante per schiena, fianchi e spalle per defaticarsi.",
		HowTo:   []string{"Inginocchiati e porta i fianchi indietro verso i talloni.", "Allunga le braccia in avanti sul pavimento.", "Lascia riposare la fronte a terra.", "Respira lentamente e rilassati."},
		Correct: []string{"I fianchi si abbassano verso i talloni.", "La schiena si allunga dolcemente.", "Le spalle si rilassano.", "Il respiro resta lento."},
		Wrong:   []string{"Forzare i fianchi verso il basso.", "Tendere le spalle.", "Trattenere il respiro.", "Sforzare le ginocchia."},
	},
}
