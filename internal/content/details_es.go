package content

// detailsES is the Spanish rich content for every exercise (current library IDs).
var detailsES = map[string]Detail{
	// ---- Calentamiento ----
	"W01": {
		Desc:    "Un calentamiento simple de hombros para preparar el cuello, los hombros y la parte alta de la espalda.",
		HowTo:   []string{"Ponte de pie, brazos relajados.", "Haz círculos lentos con los hombros hacia atrás.", "Luego lentos hacia delante.", "El movimiento es suave y controlado."},
		Correct: []string{"Relaja el cuello.", "Mueve los hombros con suavidad.", "Suelta los brazos.", "Mantén el cuerpo erguido."},
		Wrong:   []string{"No encojas los hombros con fuerza.", "No te muevas demasiado rápido.", "No tenses el cuello.", "No redondees la parte alta de la espalda."},
	},
	"W02": {
		Desc:    "Un calentamiento de hombros y espalda alta que activa los deltoides y los estabilizadores del hombro.",
		HowTo:   []string{"Ponte de pie y sube los brazos a la altura de los hombros.", "Haz pequeños círculos hacia delante.", "Luego pequeños círculos hacia atrás.", "Mantén los círculos pequeños y controlados."},
		Correct: []string{"Mantén los brazos a la altura de los hombros.", "Haz los círculos pequeños.", "Mantén los hombros abajo.", "Relaja el cuello."},
		Wrong:   []string{"No hagas grandes vaivenes.", "No subas los hombros a las orejas.", "No arquees la zona lumbar.", "No aguantes la respiración."},
	},
	"W03": {
		Desc:    "Un ejercicio suave de rotación para el tronco y la parte alta de la columna.",
		HowTo:   []string{"Pies al ancho de caderas.", "Las caderas miran sobre todo al frente.", "Gira la parte superior del cuerpo a un lado.", "Vuelve al centro y gira al otro lado."},
		Correct: []string{"Mantén los pies estables.", "Deja que el movimiento salga del tronco.", "Gira con suavidad.", "No tuerzas las rodillas."},
		Wrong:   []string{"No gires las rodillas con el cuerpo.", "No gires demasiado rápido.", "No te inclines hacia atrás.", "No fuerces el rango de movimiento."},
	},
	"W04": {
		Desc:    "Un ejercicio de movilidad para las caderas, los isquios y el control lumbar.",
		HowTo:   []string{"Ponte de pie, pies al ancho de caderas.", "Suaviza ligeramente las rodillas.", "Lleva las caderas hacia atrás.", "Baja el torso con la espalda recta.", "Vuelve empujando las caderas hacia delante."},
		Correct: []string{"Mantén la espalda recta.", "Empieza el movimiento desde las caderas.", "Mantén las rodillas suaves, sin flexionarlas mucho.", "Deja que el cuello siga la línea de la columna."},
		Wrong:   []string{"No redondees la espalda.", "No lo conviertas en una sentadilla.", "No bajes demasiado.", "No subas mucho la cabeza."},
	},
	"W05": {
		Desc:    "Un calentamiento de piernas controlado en un rango seguro y poco profundo.",
		HowTo:   []string{"Pies algo más anchos que las caderas.", "Lleva las caderas un poco hacia atrás.", "Flexiona las rodillas solo hasta una profundidad cómoda.", "Vuelve a ponerte de pie lentamente."},
		Correct: []string{"Apunta las rodillas hacia los dedos.", "Mantén los talones en el suelo.", "Mantén la espalda larga.", "Baja a una profundidad moderada."},
		Wrong:   []string{"No bajes demasiado.", "No dejes que las rodillas se hundan hacia dentro.", "No levantes los talones.", "No bajes demasiado rápido."},
	},
	"W06": {
		Desc:    "Un calentamiento para los hombros, las muñecas y el core.",
		HowTo:   []string{"Empieza en plancha alta.", "Las manos bajo los hombros.", "Desplaza el cuerpo ligeramente hacia delante.", "Vuelve a la posición inicial.", "El movimiento es pequeño y controlado."},
		Correct: []string{"Mantén el cuerpo en una línea.", "Mantén los hombros estables.", "Activa el core.", "Muévete con suavidad."},
		Wrong:   []string{"No dejes caer la cadera.", "No subas los hombros a las orejas.", "No vayas demasiado hacia delante.", "No dejes que la zona lumbar se hunda."},
	},

	// ---- Abdomen / core ----
	"C01": {
		Desc:    "Un ejercicio de estabilidad del core para el abdomen, los hombros, los glúteos y el tronco.",
		HowTo:   []string{"Coloca los codos bajo los hombros.", "Lleva los pies hacia atrás.", "Forma una línea recta de los hombros a los talones.", "Aprieta el abdomen y los glúteos.", "Mantén la posición."},
		Correct: []string{"Mantén el cuerpo en una línea.", "Coloca los codos bajo los hombros.", "Activa el core.", "Respira de forma constante."},
		Wrong:   []string{"No dejes caer la zona lumbar.", "No subas mucho la cadera.", "No mires al frente.", "No aguantes la respiración."},
	},
	"C02": {
		Desc:    "Una plancha corta e intensa centrada en la tensión de todo el cuerpo.",
		HowTo:   []string{"Empieza en plancha de antebrazos.", "Aprieta fuerte el abdomen.", "Aprieta los glúteos.", "Imagina que llevas los codos hacia los pies.", "Mantén con el máximo control."},
		Correct: []string{"Tensa todo el cuerpo.", "Aprieta los glúteos.", "Mantén el core firme.", "Controla la respiración."},
		Wrong:   []string{"No la mantengas como una plancha relajada.", "No dejes que la cadera se hunda.", "No aguantes la respiración.", "No intentes aguantar demasiado."},
	},
	"C03": {
		Desc:    "Un ejercicio de abdomen bajo con movimientos cruzados controlados de las piernas.",
		HowTo:   []string{"Túmbate boca arriba.", "Presiona suavemente la zona lumbar contra el suelo.", "Eleva las piernas.", "Cruza una pierna sobre la otra.", "Continúa alternando el cruce."},
		Correct: []string{"Mantén la zona lumbar cerca del suelo.", "Controla las piernas.", "Relaja el cuello.", "Muévete con suavidad."},
		Wrong:   []string{"No arquees la zona lumbar.", "No bajes demasiado las piernas.", "No tires de la cabeza hacia delante.", "No te muevas demasiado rápido."},
	},
	"C04": {
		Desc:    "Un ejercicio de abdomen bajo con patadas cortas y alternas.",
		HowTo:   []string{"Túmbate boca arriba.", "Presiona la zona lumbar contra el suelo.", "Eleva las piernas.", "Mueve las piernas arriba y abajo en patadas cortas y alternas."},
		Correct: []string{"Haz un rango de movimiento pequeño.", "Mantén la zona lumbar estable.", "No bajes demasiado las piernas.", "Respira de forma constante."},
		Wrong:   []string{"No arquees la zona lumbar.", "No hagas patadas grandes.", "No tenses el cuello.", "No aguantes la respiración."},
	},
	"C05": {
		Desc:    "Un ejercicio abdominal centrado en elevar la pelvis con control.",
		HowTo:   []string{"Túmbate boca arriba.", "Flexiona las rodillas y eleva los pies.", "Usa el abdomen para enrollar la pelvis ligeramente hacia arriba.", "Baja lentamente."},
		Correct: []string{"Deja que el movimiento salga del abdomen bajo.", "Sube la pelvis con control.", "Relaja el cuello.", "Baja despacio."},
		Wrong:   []string{"No balancees las piernas.", "No uses impulso.", "No tires del cuello.", "No dejes caer la cadera rápido."},
	},
	"C06": {
		Desc:    "Un mantenimiento estático del core con palanca acortada para mejor control.",
		HowTo:   []string{"Túmbate boca arriba.", "Presiona la zona lumbar contra el suelo.", "Eleva ligeramente los hombros.", "Rodillas flexionadas o piernas más altas.", "Mantén la posición."},
		Correct: []string{"Mantén la zona lumbar en el suelo.", "Mantén el abdomen firme.", "Relaja el cuello.", "Sigue respirando."},
		Wrong:   []string{"No arquees la zona lumbar.", "No bajes demasiado las piernas.", "No lleves el mentón hacia delante.", "No aguantes la respiración."},
	},
	"C07": {
		Desc:    "Un ejercicio lateral del core para los oblicuos y la estabilidad de hombro y cadera.",
		HowTo:   []string{"Túmbate de lado.", "Coloca el codo bajo el hombro.", "Eleva la cadera.", "El cuerpo en una línea recta.", "Mantén."},
		Correct: []string{"Coloca el codo bajo el hombro.", "Mantén la cadera arriba.", "Mantén el cuerpo largo.", "Mantén el cuello neutro."},
		Wrong:   []string{"No dejes caer la cadera.", "No encojas el hombro hacia la oreja.", "No ruedes el pecho hacia delante.", "No aguantes la respiración."},
	},
	"C08": {
		Desc:    "Una variación más exigente de la plancha lateral para los oblicuos y la cadera lateral.",
		HowTo:   []string{"Empieza en plancha lateral.", "Baja ligeramente la cadera.", "Vuelve a subir la cadera.", "Repite con control.", "Cambia de lado."},
		Correct: []string{"Haz un movimiento pequeño y controlado.", "Mantén el codo bajo el hombro.", "Mueve la cadera en vertical.", "Mantén el core firme."},
		Wrong:   []string{"No bajes demasiado.", "No tuerzas el tronco.", "No te hundas en el hombro.", "No te muevas demasiado rápido."},
	},
	"C10": {
		Desc:    "Un ejercicio de core y tríceps alternando entre plancha de antebrazos y plancha alta.",
		HowTo:   []string{"Empieza en plancha de antebrazos.", "Apoya una mano en el suelo y empuja hacia arriba.", "Sube la otra mano a la plancha alta.", "Baja de nuevo a los antebrazos.", "Alterna la mano que lidera."},
		Correct: []string{"Mantén la cadera estable.", "Activa el core.", "Muévete con control.", "Mantén los hombros sobre las manos o los codos."},
		Wrong:   []string{"No balancees la cadera de lado a lado.", "No dejes caer la zona lumbar.", "No te muevas demasiado rápido.", "No dejes que los hombros se hundan."},
	},
	"C11": {
		Desc:    "Un ejercicio de core controlado y una alternativa más segura a las elevaciones de pierna más duras.",
		HowTo:   []string{"Túmbate boca arriba.", "Eleva las piernas a posición de mesa.", "Baja un talón hacia el suelo.", "Toca ligeramente y vuelve.", "Alterna los lados."},
		Correct: []string{"Mantén la zona lumbar estable.", "Muévete despacio.", "Toca el talón con suavidad.", "Mantén la pelvis quieta."},
		Wrong:   []string{"No dejes caer la pierna demasiado rápido.", "No arquees la zona lumbar.", "No muevas las dos piernas a la vez.", "No te apresures."},
	},
	"C12": {
		Desc:    "Un ejercicio de control del core para el abdomen profundo y la estabilidad de la columna.",
		HowTo:   []string{"Túmbate boca arriba.", "Eleva brazos y piernas a posición de mesa.", "Extiende el brazo y la pierna contrarios.", "Mantén la zona lumbar estable.", "Vuelve y cambia de lado."},
		Correct: []string{"Mantén la zona lumbar cerca del suelo.", "Muévete despacio.", "Mueve juntos el brazo y la pierna contrarios.", "Mantén el tronco quieto."},
		Wrong:   []string{"No arquees la zona lumbar.", "No te muevas demasiado rápido.", "No bajes demasiado la pierna.", "No saques las costillas."},
	},

	// ---- Flexiones / tríceps / hombros ----
	"P01": {
		Desc:    "Un ejercicio básico de tren superior para el pecho, el tríceps, los hombros y el core.",
		HowTo:   []string{"Empieza en plancha alta.", "Manos bajo los hombros o algo más anchas.", "Baja el cuerpo con control.", "Empuja hacia arriba manteniendo el cuerpo recto."},
		Correct: []string{"Mantén el cuerpo en una línea.", "Mantén el core firme.", "Mantén los codos a unos 30 a 45 grados.", "Muévete con control."},
		Wrong:   []string{"No dejes caer la cadera.", "No subas mucho la cadera.", "No abras los codos hacia fuera.", "No caigas en la posición baja."},
	},
	"P02": {
		Desc:    "Una variación de flexión con más énfasis en el tríceps.",
		HowTo:   []string{"Empieza en plancha alta.", "Coloca las manos más juntas que en una flexión normal.", "Mantén los codos cerca del cuerpo.", "Baja despacio y empuja hacia arriba."},
		Correct: []string{"Mantén los codos cerca del cuerpo.", "Mantén el cuerpo recto.", "Mantén los hombros lejos de las orejas.", "Mantén las muñecas cómodas."},
		Wrong:   []string{"No pongas las manos demasiado juntas.", "No abras los codos.", "No dejes caer la zona lumbar.", "No fuerces si te duele la muñeca o el codo."},
	},
	"P03": {
		Desc:    "Flexiones controladas con una breve pausa en la posición baja.",
		HowTo:   []string{"Empieza en posición de flexión normal.", "Baja despacio.", "Pausa un segundo cerca del fondo.", "Empuja hacia arriba sin perder la posición del cuerpo."},
		Correct: []string{"Haz una pausa controlada.", "Mantén el core firme.", "Mantén el pecho activo.", "Mantén el cuerpo recto."},
		Wrong:   []string{"No te relajes abajo.", "No dejes caer la cadera.", "No aguantes la respiración.", "No rebotes desde el fondo."},
	},
	"P04": {
		Desc:    "Una variación lenta de flexión que aumenta el control y el tiempo bajo tensión.",
		HowTo:   []string{"Empieza en plancha alta.", "Baja durante unos tres segundos.", "Pausa un segundo cerca del fondo.", "Empuja hacia arriba durante unos tres segundos."},
		Correct: []string{"Mantén un ritmo lento y uniforme.", "Mantén el cuerpo recto.", "Controla los codos.", "Respira de forma constante."},
		Wrong:   []string{"No apresures el movimiento.", "No te dejes caer abajo.", "No dejes caer la cadera.", "No pierdas la tensión del core."},
	},
	"P05": {
		Desc:    "Una variación asimétrica de flexión que reta al pecho, los hombros y el core.",
		HowTo:   []string{"Empieza en posición de flexión.", "Coloca una mano algo adelantada y la otra algo atrasada.", "Baja con control.", "Empuja hacia arriba.", "Cambia la posición de las manos en el otro lado."},
		Correct: []string{"Mantén el cuerpo recto.", "Mantén el core firme.", "Mantén ambos hombros estables.", "Muévete con control."},
		Wrong:   []string{"No tuerzas el tronco.", "No dejes caer un hombro.", "No pongas las manos demasiado anchas.", "No fuerces el rango si te molesta el hombro."},
	},
	"P06": {
		Desc:    "Una variación de flexión centrada en los hombros.",
		HowTo:   []string{"Empieza en plancha alta.", "Acerca un poco los pies y eleva la cadera.", "La cabeza entre los brazos.", "Baja la cabeza hacia el suelo.", "Empuja de nuevo hacia arriba."},
		Correct: []string{"Mantén la cadera alta.", "Flexiona los codos con control.", "Mantén los hombros estables.", "Mantén el cuello neutro."},
		Wrong:   []string{"No caigas en una flexión normal.", "No lleves la cabeza demasiado adelante.", "No abras mucho los codos.", "No fuerces si te duele el hombro."},
	},
	"P08": {
		Desc:    "Un ejercicio de tríceps y hombros desde una posición de apoyo en antebrazos.",
		HowTo:   []string{"Empieza sobre los antebrazos, el cuerpo largo.", "Codos bajo los hombros o algo por delante.", "Empuja con las manos y los antebrazos.", "Extiende ligeramente los codos.", "Vuelve con control."},
		Correct: []string{"Mantén el core firme.", "Mueve los codos con suavidad.", "Mantén los hombros abajo.", "Trabaja en un rango cómodo."},
		Wrong:   []string{"No fuerces los codos.", "No dejes caer la cadera.", "No encojas los hombros.", "No te muevas si te duele el codo o el hombro."},
	},

	// ---- Espalda / postura ----
	"B01": {
		Desc:    "Un ejercicio de espalda centrado en la postura para la espalda alta, las escápulas y los extensores de la columna.",
		HowTo:   []string{"Túmbate boca abajo.", "Los brazos a lo largo del cuerpo.", "Gira ligeramente los pulgares hacia fuera.", "Eleva un poco el pecho.", "Lleva las escápulas atrás y abajo.", "Mantén."},
		Correct: []string{"Eleva solo un poco.", "Mantén el cuello neutro.", "Lleva las escápulas atrás y abajo.", "Activa ligeramente los glúteos."},
		Wrong:   []string{"No eleves demasiado.", "No eches la cabeza atrás.", "No hiperextiendas la zona lumbar.", "No subas los hombros a las orejas."},
	},
	"B02": {
		Desc:    "Un ejercicio de espalda alta y escápulas boca abajo.",
		HowTo:   []string{"Túmbate boca abajo.", "La cabeza mira al suelo.", "Lleva los brazos en un arco amplio desde los lados hacia arriba.", "Vuelve por el mismo camino.", "Movimiento controlado."},
		Correct: []string{"Relaja el cuello.", "Mueve las escápulas con suavidad.", "Eleva el pecho solo un poco o mantenlo neutro.", "Mueve los brazos sin prisa."},
		Wrong:   []string{"No levantes la cabeza.", "No hiperextiendas mucho la zona lumbar.", "No balancees los brazos.", "No subas los hombros a las orejas."},
	},
	"B03": {
		Desc:    "Un ejercicio de espalda alta para las escápulas y los hombros posteriores.",
		HowTo:   []string{"Túmbate boca abajo.", "Flexiona los codos en forma de W.", "Eleva un poco los codos y las manos.", "Aprieta las escápulas.", "Baja con control."},
		Correct: []string{"Haz un rango de movimiento pequeño.", "Deja que el trabajo lo hagan las escápulas.", "Mantén el cuello neutro.", "Mantén los hombros lejos de las orejas."},
		Wrong:   []string{"No eleves demasiado.", "No levantes la cabeza.", "No uses impulso.", "No hiperextiendas la zona lumbar."},
	},
	"B04": {
		Desc:    "Un ejercicio de escápulas para el trapecio inferior y la espalda alta.",
		HowTo:   []string{"Túmbate boca abajo.", "Extiende los brazos en diagonal hacia arriba en forma de Y.", "Eleva un poco los brazos del suelo.", "Baja con control."},
		Correct: []string{"Sube los brazos solo un poco.", "Mantén el cuello neutro.", "Mantén los hombros abajo.", "Muévete con control."},
		Wrong:   []string{"No eleves demasiado.", "No subas los hombros a las orejas.", "No arquees la zona lumbar.", "No te muevas demasiado rápido."},
	},
	"B05": {
		Desc:    "Un ejercicio de espalda alta para los hombros posteriores y el control de las escápulas.",
		HowTo:   []string{"Túmbate boca abajo.", "Extiende los brazos a los lados en forma de T.", "Eleva un poco los brazos.", "Aprieta las escápulas.", "Baja lentamente."},
		Correct: []string{"Mantén los brazos en la línea de los hombros.", "Relaja el cuello.", "Mueve las escápulas con suavidad.", "Eleva solo un poco."},
		Wrong:   []string{"No lances los brazos hacia arriba.", "No levantes la cabeza.", "No encojas los hombros.", "No fuerces un rango amplio."},
	},
	"B06": {
		Desc:    "Un ejercicio de espalda que imita un movimiento de tracción sin equipo.",
		HowTo:   []string{"Túmbate boca abajo con los brazos por encima de la cabeza.", "Eleva un poco el pecho.", "Lleva los codos hacia abajo hacia las costillas.", "Aprieta las escápulas.", "Extiende los brazos de nuevo hacia delante."},
		Correct: []string{"Baja los codos con control.", "Aprieta las escápulas con suavidad.", "Mantén el cuello neutro.", "Mantén la zona lumbar cómoda."},
		Wrong:   []string{"No des tirones con los brazos.", "No eleves demasiado.", "No eches la cabeza atrás.", "No subas los hombros a las orejas."},
	},
	"B07": {
		Desc:    "Un ejercicio de espalda y coordinación boca abajo.",
		HowTo:   []string{"Túmbate boca abajo.", "Eleva un poco el pecho.", "Mueve el brazo y la pierna contrarios en un patrón de nado controlado.", "El movimiento es pequeño y constante."},
		Correct: []string{"Mantén el cuello neutro.", "Muévete con control.", "Mantén la zona lumbar cómoda.", "Sigue respirando."},
		Wrong:   []string{"No te muevas demasiado rápido.", "No eleves demasiado.", "No eches la cabeza atrás.", "No fuerces la zona lumbar."},
	},
	"B08": {
		Desc:    "Un ejercicio de cadena posterior para los glúteos, los isquios, los hombros y la espalda.",
		HowTo:   []string{"Siéntate en el suelo con las piernas extendidas.", "Coloca las manos detrás de la cadera.", "Empuja con las manos y los talones.", "Eleva la cadera.", "Mantén el cuerpo en una línea larga."},
		Correct: []string{"Abre el pecho.", "Mantén la cadera arriba.", "Mantén los hombros estables.", "Mantén el cuello neutro."},
		Wrong:   []string{"No dejes caer la cadera.", "No encojas los hombros.", "No eches la cabeza atrás.", "No fuerces las muñecas."},
	},
	"B09": {
		Desc:    "Un ejercicio de espalda de rango pequeño para los extensores de la columna y la postura.",
		HowTo:   []string{"Túmbate boca abajo.", "Las manos junto al cuerpo o ligeramente cerca del pecho.", "Eleva un poco el pecho.", "Haz pequeños impulsos controlados.", "Cuello neutro."},
		Correct: []string{"Eleva solo un poco.", "Haz impulsos suaves.", "Mantén el cuello largo.", "Controla la zona lumbar."},
		Wrong:   []string{"No impulses demasiado alto.", "No uses impulso.", "No mires al frente.", "No sigas si sientes dolor agudo en la zona lumbar."},
	},
	"B10": {
		Desc:    "Un ejercicio de espalda más exigente que combina una pequeña elevación de cobra con una tracción de brazos en forma de W.",
		HowTo:   []string{"Túmbate boca abajo con los brazos al frente.", "Eleva un poco el pecho.", "Lleva los codos abajo y atrás en forma de W.", "Aprieta las escápulas.", "Extiende de nuevo hacia delante y baja con control."},
		Correct: []string{"Eleva el pecho solo un poco.", "Haz la tracción en W con suavidad.", "Lleva las escápulas atrás y abajo.", "Mantén el cuello largo."},
		Wrong:   []string{"No hiperextiendas la zona lumbar.", "No tires con el cuello.", "No te muevas demasiado rápido.", "No subas los hombros a las orejas."},
	},

	// ---- Brazos / bíceps ----
	"A01": {
		Desc:    "Un ejercicio de bíceps sin equipo: una mano opone resistencia a la otra.",
		HowTo:   []string{"De pie o sentado, erguido.", "Flexiona un codo como en un curl.", "Con la otra mano presiona hacia abajo el antebrazo.", "Sube despacio contra la resistencia.", "Cambia de lado."},
		Correct: []string{"Mantén una resistencia constante.", "Muévete despacio.", "Mantén el codo cerca del cuerpo.", "Mantén el hombro abajo."},
		Wrong:   []string{"No te muevas sin resistencia real.", "No des tirones con el brazo.", "No subas el hombro.", "No aguantes la respiración."},
	},
	"A02": {
		Desc:    "Un mantenimiento estático de bíceps con autorresistencia.",
		HowTo:   []string{"Flexiona un codo a unos 90 grados.", "Con la otra mano presiona hacia abajo.", "El brazo que trabaja opone resistencia sin moverse.", "Mantén la tensión.", "Cambia de lado."},
		Correct: []string{"Mantén estable el ángulo del codo.", "Mantén una tensión constante.", "Relaja los hombros.", "Sigue respirando."},
		Wrong:   []string{"No presiones a tirones.", "No tuerzas el tronco.", "No subas el hombro.", "No aguantes la respiración."},
	},
	"A03": {
		Desc:    "Un ejercicio isométrico de brazo y antebrazo con las manos una contra otra.",
		HowTo:   []string{"Engancha los dedos o las manos.", "Codos ligeramente flexionados.", "Tira de las manos para separarlas.", "Mantén una tensión constante.", "Relaja despacio."},
		Correct: []string{"Mantén la tensión controlada.", "Mantén los hombros abajo.", "Mantén las muñecas cómodas.", "Respira con tranquilidad."},
		Wrong:   []string{"No tires de golpe.", "No encojas los hombros.", "No flexiones en exceso las muñecas.", "No aguantes la respiración."},
	},
	"A04": {
		Desc:    "Un ejercicio de bajada lenta para el bíceps con autorresistencia.",
		HowTo:   []string{"Empieza con un codo flexionado.", "Con la otra mano crea resistencia.", "Baja el antebrazo lentamente.", "Mantén la resistencia durante la bajada.", "Cambia de lado."},
		Correct: []string{"Baja despacio.", "Mantén una resistencia constante.", "Mantén el codo cerca del cuerpo.", "Relaja el hombro."},
		Wrong:   []string{"No dejes caer el brazo rápido.", "No te quedes sin resistencia.", "No tuerzas el tronco.", "No aguantes la respiración."},
	},

	// ---- Piernas / glúteos ----
	"L01": {
		Desc:    "Un ejercicio de piernas controlado para los muslos, las caderas y la postura.",
		HowTo:   []string{"Pies algo más anchos que las caderas.", "Lleva las caderas hacia atrás.", "Flexiona las rodillas a una profundidad cómoda y poco profunda.", "Vuelve a ponerte de pie lentamente."},
		Correct: []string{"Apunta las rodillas hacia los dedos.", "Mantén los talones abajo.", "Mantén la espalda larga.", "Baja a una profundidad cómoda."},
		Wrong:   []string{"No bajes demasiado.", "No dejes que las rodillas se hundan hacia dentro.", "No levantes los talones.", "No bajes rápido."},
	},
	"L02": {
		Desc:    "Una variación de sentadilla controlada con una breve pausa en un rango seguro.",
		HowTo:   []string{"De pie y estable.", "Baja a una sentadilla cómoda y poco profunda.", "Pausa breve.", "Vuelve a ponerte de pie con control."},
		Correct: []string{"Haz una pausa estable.", "Apunta las rodillas hacia los dedos.", "Abre el pecho.", "Mantén los talones en el suelo."},
		Wrong:   []string{"No bajes demasiado.", "No rebotes en la posición baja.", "No dejes que las rodillas se hundan hacia dentro.", "No aguantes la respiración."},
	},
	"L03": {
		Desc:    "Un ejercicio de movilidad para las caderas, los isquios y el control lumbar.",
		HowTo:   []string{"Ponte de pie, pies al ancho de caderas.", "Suaviza ligeramente las rodillas.", "Lleva las caderas hacia atrás.", "Baja el torso con la espalda recta.", "Vuelve empujando las caderas hacia delante."},
		Correct: []string{"Mantén la espalda recta.", "Empieza el movimiento desde las caderas.", "Mantén las rodillas suaves, sin flexionarlas mucho.", "Deja que el cuello siga la línea de la columna."},
		Wrong:   []string{"No redondees la espalda.", "No lo conviertas en una sentadilla.", "No bajes demasiado.", "No subas mucho la cabeza."},
	},
	"L04": {
		Desc:    "Un ejercicio de glúteos más exigente con una pierna cada vez.",
		HowTo:   []string{"Túmbate boca arriba.", "Flexiona una rodilla con ese pie en el suelo.", "Extiende o eleva la otra pierna.", "Empuja con el talón que trabaja.", "Eleva la cadera y baja despacio."},
		Correct: []string{"Deja que la elevación la haga el glúteo que trabaja.", "Mantén la cadera nivelada.", "Muévete con control.", "Mantén la zona lumbar neutra."},
		Wrong:   []string{"No tuerzas la pelvis.", "No empujes con los dedos del pie.", "No arquees la zona lumbar.", "No bajes demasiado rápido."},
	},
	"L05": {
		Desc:    "Un ejercicio de glúteos y estabilidad de la pelvis.",
		HowTo:   []string{"Empieza en posición de puente de glúteos.", "Mantén la cadera elevada.", "Eleva un pie ligeramente del suelo.", "Bájalo y cambia de lado.", "Mantén la pelvis estable."},
		Correct: []string{"Mantén la cadera nivelada.", "Aprieta los glúteos.", "Muévete despacio.", "Mantén la zona lumbar cómoda."},
		Wrong:   []string{"No dejes caer la cadera.", "No te balancees de lado a lado.", "No te muevas demasiado rápido.", "No arquees la zona lumbar."},
	},
	"L06": {
		Desc:    "Un ejercicio de glúteos que fortalece las caderas y la cadena posterior.",
		HowTo:   []string{"Túmbate boca arriba.", "Flexiona las rodillas, pies en el suelo.", "Eleva la cadera.", "Pausa arriba.", "Baja lentamente."},
		Correct: []string{"Empuja con los talones.", "Aprieta los glúteos arriba.", "Mantén las costillas abajo.", "No hiperextiendas la zona lumbar."},
		Wrong:   []string{"No hiperextiendas la zona lumbar arriba.", "No levantes los talones.", "No dejes caer la cadera demasiado rápido.", "No pongas los pies demasiado juntos."},
	},
	"L07": {
		Desc:    "Un ejercicio de pierna baja para las pantorrillas, los tobillos y el control del pie.",
		HowTo:   []string{"De pie, pies al ancho de caderas.", "Sube sobre las puntas de los pies.", "Pausa breve arriba.", "Baja los talones lentamente."},
		Correct: []string{"Mantén el cuerpo erguido.", "Muévete con suavidad.", "Mantén los tobillos alineados.", "Baja con control."},
		Wrong:   []string{"No dejes caer los talones rápido.", "No te balancees hacia delante.", "No vuelques los tobillos hacia fuera.", "No uses impulso."},
	},
	"L09": {
		Desc:    "Un ejercicio de cadena lateral para los glúteos laterales, los oblicuos y la estabilidad de la cadera.",
		HowTo:   []string{"Empieza en plancha lateral.", "El cuerpo en una línea recta.", "Eleva un poco la pierna de arriba.", "Bájala con control.", "Repite y cambia de lado."},
		Correct: []string{"Mantén la cadera arriba.", "Mueve la pierna de arriba despacio.", "Mantén el tronco estable.", "Mantén el hombro fuerte."},
		Wrong:   []string{"No dejes caer la cadera.", "No balancees la pierna.", "No ruedes el cuerpo hacia atrás.", "No te hundas en el hombro."},
	},

	// ---- Vuelta a la calma ----
	"CD01": {
		Desc:    "Un estiramiento suave del pecho y la parte frontal de los hombros.",
		HowTo:   []string{"Túmbate boca arriba.", "Abre los brazos a los lados.", "Relaja los hombros.", "Deja que el pecho se abra con suavidad.", "Respira despacio."},
		Correct: []string{"Deja que el pecho se abra con suavidad.", "Relaja los hombros.", "Mantén la zona lumbar cómoda.", "No fuerces si te duelen los hombros."},
		Wrong:   []string{"No fuerces los brazos hacia abajo.", "No arquees la zona lumbar.", "No tenses el cuello.", "No estires si te duele."},
	},
	"CD02": {
		Desc:    "Un estiramiento suave de la parte posterior del muslo.",
		HowTo:   []string{"Túmbate boca arriba.", "Eleva una pierna.", "Sujeta por detrás del muslo o la pantorrilla.", "Lleva la pierna suavemente hacia ti.", "Cambia de lado."},
		Correct: []string{"Estira con suavidad.", "Deja la rodilla algo flexionada.", "Mantén la zona lumbar tranquila.", "Respira despacio."},
		Wrong:   []string{"No tires demasiado fuerte.", "No fuerces la pierna recta.", "No eleves la cadera.", "No sigas si sientes dolor agudo detrás de la rodilla."},
	},
	"CD03": {
		Desc:    "Un estiramiento suave de rotación para la columna y el tronco.",
		HowTo:   []string{"Túmbate boca arriba.", "Flexiona las rodillas.", "Deja que las rodillas caigan suavemente a un lado.", "Mantén los hombros en el suelo.", "Vuelve al centro y cambia de lado."},
		Correct: []string{"Muévete despacio.", "Mantén los hombros en el suelo.", "Respira de forma relajada.", "Estira de forma cómoda."},
		Wrong:   []string{"No fuerces las rodillas al suelo.", "No gires rápido.", "No levantes el hombro.", "No estires si te duele."},
	},
	"CD04": {
		Desc:    "Un estiramiento suave de la parte frontal del cuerpo y una leve extensión de la espalda.",
		HowTo:   []string{"Túmbate boca abajo.", "Coloca los antebrazos en el suelo.", "Eleva el pecho con suavidad.", "Mantén el cuello largo.", "Respira despacio."},
		Correct: []string{"Eleva solo un poco.", "Mantén los hombros lejos de las orejas.", "Mantén la zona lumbar cómoda.", "Mantén el cuello neutro."},
		Wrong:   []string{"No hiperextiendas la zona lumbar.", "No eches la cabeza atrás.", "No encojas los hombros.", "No fuerces la posición."},
	},
	"CD05": {
		Desc:    "Un ejercicio de respiración tranquila para terminar la sesión.",
		HowTo:   []string{"Túmbate boca arriba.", "Relaja los hombros y la mandíbula.", "Si es cómodo, pon una mano en el vientre.", "Inhala suavemente.", "Exhala lentamente."},
		Correct: []string{"Respira con tranquilidad.", "Relaja los hombros.", "Relaja la cara.", "Deja que el cuerpo se calme."},
		Wrong:   []string{"No respires con demasiada fuerza.", "No aguantes la respiración.", "No tenses el cuello.", "No arquees la zona lumbar."},
	},

	// ---- Añadidos del set Vlad (calentamiento / cardio / pliometría / zancadas) ----
	"W07": {
		Desc:    "Un ejercicio de apertura tranquilo: respiración profunda con una extensión suave de todo el cuerpo para alargar la columna.",
		HowTo:   []string{"Ponte de pie y erguido, pies al ancho de caderas.", "Inhala y lleva ambos brazos por encima de la cabeza.", "Alarga la columna con suavidad.", "Exhala y baja los brazos, dejando caer los hombros."},
		Correct: []string{"Respira de forma lenta y completa.", "Relaja los hombros al exhalar.", "Mantén las costillas abajo, sin arquear de más.", "Muévete sin prisa."},
		Wrong:   []string{"No aguantes la respiración.", "No arquees con fuerza la zona lumbar.", "No subas los hombros a las orejas.", "No apresures la extensión."},
	},
	"W08": {
		Desc:    "Un ejercicio suave de rotación para calentar la columna y el tronco.",
		HowTo:   []string{"Ponte de pie, pies al ancho de caderas.", "Deja que los brazos se balanceen sueltos.", "Gira la parte superior del cuerpo a un lado.", "Fluye con suavidad hacia el otro lado."},
		Correct: []string{"Deja que el movimiento salga del tronco.", "Mantén las caderas mirando sobre todo al frente.", "Mantén los pies estables.", "Lleva un ritmo suave y uniforme."},
		Wrong:   []string{"No tuerzas con fuerza las rodillas.", "No des tirones con los brazos.", "No gires demasiado rápido.", "No aguantes la respiración."},
	},
	"W09": {
		Desc:    "Un calentamiento dinámico de hombros y pecho para preparar el tren superior para flexiones y saltos.",
		HowTo:   []string{"Ponte de pie y erguido, pies al ancho de caderas.", "Balancea los brazos hacia delante y hacia atrás de forma controlada.", "Deja que el pecho se abra con naturalidad al llevar los brazos atrás.", "Mantén el movimiento suave y relajado."},
		Correct: []string{"Mueve los brazos con libertad.", "Mantén los hombros relajados.", "Deja que el pecho se abra con suavidad.", "Mantén el cuerpo erguido."},
		Wrong:   []string{"No te balancees con demasiada fuerza.", "No subas los hombros a las orejas.", "No arquees la zona lumbar.", "No conviertas el movimiento en un impulso descontrolado."},
	},
	"C13": {
		Desc:    "Un ejercicio abdominal de rotación para los oblicuos y el abdomen frontal.",
		HowTo:   []string{"Túmbate boca arriba.", "Sube las rodillas.", "Coloca las manos suaves cerca de la cabeza.", "Gira el tronco y lleva un codo hacia la rodilla contraria.", "Cambia de lado en un ritmo controlado."},
		Correct: []string{"Deja que la rotación salga del tronco.", "Mantén la zona lumbar controlada.", "Relaja el cuello.", "Muévete con suavidad."},
		Wrong:   []string{"No tires del cuello.", "No muevas solo los codos.", "No apresures las repeticiones.", "No dejes que la zona lumbar se arquee."},
	},
	"B11": {
		Desc:    "Un ejercicio de espalda que eleva el brazo y la pierna contrarios en un patrón de nado alterno.",
		HowTo:   []string{"Túmbate boca abajo con los brazos por encima de la cabeza.", "Eleva un poco el pecho y las piernas.", "Sube un brazo y la pierna contraria.", "Baja y cambia de lado, alternando con suavidad."},
		Correct: []string{"Eleva solo un poco.", "Deja que el cuello siga la línea de la columna.", "Activa los glúteos.", "Muévete de forma constante."},
		Wrong:   []string{"No fuerces la zona lumbar.", "No levantes la cabeza.", "No te muevas a tirones.", "No aguantes la respiración."},
	},
	"L10": {
		Desc:    "Una sentadilla básica con el peso del cuerpo para las piernas, las caderas y los glúteos.",
		HowTo:   []string{"Ponte de pie con los pies más o menos al ancho de los hombros.", "Lleva las caderas un poco hacia atrás.", "Flexiona las rodillas y baja a una sentadilla.", "Mantén el pecho abierto.", "Vuelve a ponerte de pie empujando con todo el pie."},
		Correct: []string{"Apunta las rodillas hacia la línea de los pies.", "Mantén los talones en el suelo.", "Mantén la espalda larga.", "Controla la profundidad."},
		Wrong:   []string{"No dejes que las rodillas se hundan hacia dentro.", "No levantes los talones.", "No redondees la espalda.", "No bajes demasiado rápido."},
	},
	"L11": {
		Desc:    "Un ejercicio del tren inferior para las piernas, los glúteos y el equilibrio.",
		HowTo:   []string{"Ponte de pie y erguido.", "Da un paso atrás con una pierna.", "Baja a una zancada con control.", "Empuja con el pie delantero para volver de pie.", "Cambia de lado."},
		Correct: []string{"Apunta la rodilla delantera hacia la línea de los pies.", "Mantén el torso erguido.", "Da el paso atrás de forma controlada.", "Mantén el talón delantero en el suelo."},
		Wrong:   []string{"No dejes que la rodilla delantera se hunda hacia dentro.", "No te inclines demasiado hacia delante.", "No bajes rápido.", "No empujes con la pierna trasera con demasiada fuerza."},
	},
	"L12": {
		Desc:    "Una sentadilla sin peso hecha despacio para ganar control y fuerza de piernas.",
		HowTo:   []string{"Pies al ancho de los hombros.", "Baja durante unos tres segundos.", "Pausa breve en la posición baja.", "Ponte de pie con control."},
		Correct: []string{"Mantén un ritmo lento.", "Apunta las rodillas hacia los dedos.", "Mantén los talones en el suelo.", "Mantén la espalda neutra."},
		Wrong:   []string{"No bajes rápido.", "No rebotes desde el fondo.", "No dejes que las rodillas se hundan hacia dentro.", "No aguantes la respiración."},
	},
	"J01": {
		Desc:    "Un calentamiento suave de tobillos y pantorrillas con saltitos cortos y rápidos.",
		HowTo:   []string{"Ponte de pie y erguido, pies más o menos al ancho de caderas.", "Mantén las rodillas algo suaves.", "Haz saltitos cortos y rápidos usando sobre todo los tobillos.", "Aterriza en silencio y repite."},
		Correct: []string{"Mantén los saltos bajos.", "Aterriza con suavidad y sin ruido.", "Mantén las rodillas algo flexionadas.", "Mantén el cuerpo erguido."},
		Wrong:   []string{"No saltes demasiado alto.", "No aterrices con ruido.", "No bloquees las rodillas.", "No vuelques los tobillos hacia dentro ni hacia fuera."},
	},
	"J02": {
		Desc:    "Una alternativa de calentamiento de bajo impacto a las rodillas altas.",
		HowTo:   []string{"Ponte de pie y erguido.", "Eleva una rodilla hacia la altura de la cadera.", "Bájala con control.", "Cambia de lado y sigue marchando.", "Mantén los brazos moviéndose con naturalidad."},
		Correct: []string{"Mantén el torso erguido.", "Sube las rodillas con control.", "Apoya los pies con suavidad.", "Respira de forma constante."},
		Wrong:   []string{"No te inclines hacia atrás.", "No balancees las piernas sin control.", "No pises con fuerza.", "No aguantes la respiración."},
	},
	"J03": {
		Desc:    "Un ejercicio cardiovascular de todo el cuerpo que sube las pulsaciones y calienta los hombros, las caderas y las piernas.",
		HowTo:   []string{"Ponte de pie y erguido, brazos a los lados.", "Salta abriendo los pies mientras subes los brazos por encima de la cabeza.", "Salta cerrando los pies mientras bajas los brazos.", "Repite con un ritmo constante."},
		Correct: []string{"Aterriza con suavidad.", "Mantén las rodillas algo flexionadas.", "Mueve los brazos con suavidad.", "Lleva una respiración rítmica."},
		Wrong:   []string{"No aterrices fuerte.", "No bloquees las rodillas.", "No encojas los hombros.", "No te muevas demasiado rápido ni pierdas el control."},
	},
	"J04": {
		Desc:    "Un ejercicio dinámico de core y cardio realizado desde plancha alta.",
		HowTo:   []string{"Empieza en plancha alta.", "Mantén las manos bajo los hombros.", "Lleva una rodilla hacia el pecho.", "Cambia de pierna con ritmo.", "Mantén la cadera estable."},
		Correct: []string{"Mantén el cuerpo en una plancha firme.", "Mantén el core apretado.", "Mantén los hombros sobre las manos.", "Mueve las rodillas con control."},
		Wrong:   []string{"No dejes que la cadera rebote hacia arriba.", "No dejes caer la zona lumbar.", "No muevas los pies de forma descontrolada.", "No dejes que los hombros se hundan."},
	},
	"J05": {
		Desc:    "Un ejercicio pliométrico de piernas para la potencia, el acondicionamiento y la coordinación.",
		HowTo:   []string{"Empieza con los pies más o menos al ancho de los hombros.", "Baja a una sentadilla controlada.", "Salta hacia arriba.", "Aterriza con suavidad con las rodillas algo flexionadas.", "Recoloca tu posición antes del siguiente salto."},
		Correct: []string{"Aterriza en silencio.", "Apunta las rodillas hacia la línea de los pies.", "Mantén el pecho abierto.", "Controla cada salto."},
		Wrong:   []string{"No aterrices fuerte.", "No dejes que las rodillas se hundan hacia dentro.", "No saltes antes de que la sentadilla sea estable.", "No te apresures sin recolocarte."},
	},
	"J06": {
		Desc:    "Un ejercicio de salto lateral para las piernas, las caderas, el equilibrio y la coordinación.",
		HowTo:   []string{"Apóyate sobre una pierna.", "Salta de lado hacia la otra pierna.", "Aterriza con suavidad y estabilízate.", "Repite de un lado a otro.", "Usa los brazos para el equilibrio."},
		Correct: []string{"Aterriza en silencio.", "Apunta la rodilla hacia la línea de los pies.", "Mantén el torso controlado.", "Muévete de lado a lado con suavidad."},
		Wrong:   []string{"No dejes que la rodilla se hunda hacia dentro al aterrizar.", "No saltes demasiado lejos demasiado pronto.", "No te dejes caer en el aterrizaje.", "No tuerzas la rodilla al contacto."},
	},
	"J07": {
		Desc:    "Un ejercicio cardiovascular que sube las pulsaciones y activa las caderas, el core y las piernas.",
		HowTo:   []string{"Ponte de pie y erguido.", "Corre en el sitio elevando bien las rodillas.", "Mueve los brazos con naturalidad.", "Mantén el ritmo rápido pero controlado."},
		Correct: []string{"Sube las rodillas hacia la altura de la cadera.", "Aterriza ligero.", "Mantén el core activo.", "Mantén el torso erguido."},
		Wrong:   []string{"No te inclines hacia atrás.", "No aterrices con fuerza.", "No dejes que los pies golpeen el suelo.", "No pierdas la postura al aumentar el ritmo."},
	},
	"J08": {
		Desc:    "Un ejercicio de acondicionamiento de todo el cuerpo que combina una sentadilla, una plancha y un salto. En este programa, usa la versión sin flexión salvo que se indique una flexión específicamente.",
		HowTo:   []string{"Ponte de pie y erguido.", "Baja a sentadilla y apoya las manos en el suelo.", "Salta o lleva los pies atrás a una plancha.", "Salta o lleva los pies hacia delante.", "Ponte de pie y termina con un pequeño salto."},
		Correct: []string{"Apoya las manos bajo los hombros.", "Mantén la plancha firme.", "Apoya los pies con suavidad.", "Controla el salto."},
		Wrong:   []string{"No dejes caer la zona lumbar en la plancha.", "No aterrices con fuerza.", "No apresures las repeticiones con técnica descuidada.", "No dejes que las rodillas se hundan hacia dentro al saltar hacia delante."},
	},
	"J09": {
		Desc:    "Un ejercicio pliométrico avanzado de piernas para la potencia, el acondicionamiento y la coordinación.",
		HowTo:   []string{"Empieza en posición de zancada.", "Salta hacia arriba.", "Cambia de pierna en el aire.", "Aterriza con suavidad en la zancada contraria.", "Recupera el equilibrio antes de la siguiente repetición."},
		Correct: []string{"Aterriza con suavidad y control.", "Apunta la rodilla delantera hacia la línea de los pies.", "Mantén el torso erguido.", "Mantén el rango de movimiento seguro."},
		Wrong:   []string{"No aterrices fuerte.", "No dejes que la rodilla delantera se hunda hacia dentro.", "No bajes demasiado.", "No te muevas más rápido de lo que puedes controlar."},
	},
	"CD07": {
		Desc:    "Un estiramiento de descanso para la espalda, las caderas y los hombros para relajarse.",
		HowTo:   []string{"Arrodíllate y lleva las caderas hacia atrás, hacia los talones.", "Extiende los brazos hacia delante en el suelo.", "Deja que la frente descanse en el suelo.", "Respira despacio y relájate."},
		Correct: []string{"Deja que las caderas se acomoden hacia los talones.", "Alarga la espalda con suavidad.", "Relaja los hombros.", "Respira despacio."},
		Wrong:   []string{"No fuerces las caderas hacia abajo.", "No tenses los hombros.", "No aguantes la respiración.", "No fuerces las rodillas."},
	},
}
