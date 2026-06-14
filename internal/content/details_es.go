package content

// detailsES is the Spanish rich content for every exercise (current library IDs).
var detailsES = map[string]Detail{
	// ---- Calentamiento ----
	"W01": {
		Desc:    "Un calentamiento simple de hombros para preparar el cuello, los hombros y la parte alta de la espalda.",
		HowTo:   []string{"Ponte de pie, brazos relajados.", "Haz círculos lentos con los hombros hacia atrás.", "Luego lentos hacia delante.", "El movimiento es suave y controlado."},
		Correct: []string{"Cuello relajado.", "Los hombros se mueven con suavidad.", "Brazos sueltos.", "Cuerpo erguido."},
		Wrong:   []string{"Encoger los hombros con fuerza.", "Moverse demasiado rápido.", "Tensar el cuello.", "Redondear la parte alta de la espalda."},
	},
	"W02": {
		Desc:    "Un calentamiento de hombros y espalda alta que activa los deltoides y los estabilizadores del hombro.",
		HowTo:   []string{"Ponte de pie y sube los brazos a la altura de los hombros.", "Haz pequeños círculos hacia delante.", "Luego pequeños círculos hacia atrás.", "Mantén los círculos pequeños y controlados."},
		Correct: []string{"Brazos a la altura de los hombros.", "Los círculos son pequeños.", "Hombros abajo.", "Cuello relajado."},
		Wrong:   []string{"Hacer grandes vaivenes.", "Subir los hombros a las orejas.", "Arquear la zona lumbar.", "Aguantar la respiración."},
	},
	"W03": {
		Desc:    "Un ejercicio suave de rotación para el tronco y la parte alta de la columna.",
		HowTo:   []string{"Pies al ancho de caderas.", "Las caderas miran sobre todo al frente.", "Gira la parte superior del cuerpo a un lado.", "Vuelve al centro y gira al otro lado."},
		Correct: []string{"Los pies estables.", "El movimiento sale del tronco.", "La rotación es suave.", "Sin torcer las rodillas."},
		Wrong:   []string{"Girar las rodillas con el cuerpo.", "Girar demasiado rápido.", "Inclinarse hacia atrás.", "Forzar el rango de movimiento."},
	},
	"W04": {
		Desc:    "Un ejercicio de movilidad para las caderas, los isquios y el control lumbar.",
		HowTo:   []string{"Ponte de pie, pies al ancho de caderas.", "Suaviza ligeramente las rodillas.", "Lleva las caderas hacia atrás.", "Baja el torso con la espalda recta.", "Vuelve empujando las caderas hacia delante."},
		Correct: []string{"La espalda se mantiene recta.", "El movimiento empieza en las caderas.", "Rodillas suaves, no muy flexionadas.", "El cuello sigue la línea de la columna."},
		Wrong:   []string{"Redondear la espalda.", "Convertirlo en sentadilla.", "Bajar demasiado.", "Subir mucho la cabeza."},
	},
	"W05": {
		Desc:    "Un calentamiento de piernas controlado en un rango seguro y poco profundo.",
		HowTo:   []string{"Pies algo más anchos que las caderas.", "Lleva las caderas un poco hacia atrás.", "Flexiona las rodillas solo hasta una profundidad cómoda.", "Vuelve a ponerte de pie lentamente."},
		Correct: []string{"Rodillas en dirección a los dedos.", "Talones en el suelo.", "Espalda larga.", "Profundidad moderada."},
		Wrong:   []string{"Bajar demasiado.", "Dejar caer las rodillas hacia dentro.", "Levantar los talones.", "Bajar demasiado rápido."},
	},
	"W06": {
		Desc:    "Un calentamiento para los hombros, las muñecas y el core.",
		HowTo:   []string{"Empieza en plancha alta.", "Las manos bajo los hombros.", "Desplaza el cuerpo ligeramente hacia delante.", "Vuelve a la posición inicial.", "El movimiento es pequeño y controlado."},
		Correct: []string{"Cuerpo en una línea.", "Hombros estables.", "Core activo.", "Movimiento suave."},
		Wrong:   []string{"Dejar caer la cadera.", "Subir los hombros a las orejas.", "Ir demasiado hacia delante.", "Dejar que la zona lumbar se hunda."},
	},

	// ---- Abdomen / core ----
	"C01": {
		Desc:    "Un ejercicio de estabilidad del core para el abdomen, los hombros, los glúteos y el tronco.",
		HowTo:   []string{"Coloca los codos bajo los hombros.", "Lleva los pies hacia atrás.", "Forma una línea recta de los hombros a los talones.", "Aprieta el abdomen y los glúteos.", "Mantén la posición."},
		Correct: []string{"Cuerpo en una línea.", "Codos bajo los hombros.", "Core activo.", "Respiración constante."},
		Wrong:   []string{"Dejar caer la zona lumbar.", "Subir mucho la cadera.", "Mirar al frente.", "Aguantar la respiración."},
	},
	"C02": {
		Desc:    "Una plancha corta e intensa centrada en la tensión de todo el cuerpo.",
		HowTo:   []string{"Empieza en plancha de antebrazos.", "Aprieta fuerte el abdomen.", "Aprieta los glúteos.", "Imagina que llevas los codos hacia los pies.", "Mantén con el máximo control."},
		Correct: []string{"Tensión de todo el cuerpo.", "Glúteos activos.", "Core firme.", "Respiración controlada."},
		Wrong:   []string{"Mantenerla como una plancha relajada.", "Dejar que la cadera se hunda.", "Aguantar la respiración.", "Intentar aguantar demasiado."},
	},
	"C03": {
		Desc:    "Un ejercicio de abdomen bajo con movimientos cruzados controlados de las piernas.",
		HowTo:   []string{"Túmbate boca arriba.", "Presiona suavemente la zona lumbar contra el suelo.", "Eleva las piernas.", "Cruza una pierna sobre la otra.", "Continúa alternando el cruce."},
		Correct: []string{"La zona lumbar cerca del suelo.", "Piernas controladas.", "Cuello relajado.", "Movimiento suave."},
		Wrong:   []string{"Arquear la zona lumbar.", "Bajar demasiado las piernas.", "Tirar de la cabeza hacia delante.", "Moverse demasiado rápido."},
	},
	"C04": {
		Desc:    "Un ejercicio de abdomen bajo con patadas cortas y alternas.",
		HowTo:   []string{"Túmbate boca arriba.", "Presiona la zona lumbar contra el suelo.", "Eleva las piernas.", "Mueve las piernas arriba y abajo en patadas cortas y alternas."},
		Correct: []string{"Rango de movimiento pequeño.", "Zona lumbar estable.", "Las piernas no bajan demasiado.", "Respiración constante."},
		Wrong:   []string{"Arquear la zona lumbar.", "Hacer patadas grandes.", "Tensar el cuello.", "Aguantar la respiración."},
	},
	"C05": {
		Desc:    "Un ejercicio abdominal centrado en elevar la pelvis con control.",
		HowTo:   []string{"Túmbate boca arriba.", "Flexiona las rodillas y eleva los pies.", "Usa el abdomen para enrollar la pelvis ligeramente hacia arriba.", "Baja lentamente."},
		Correct: []string{"El movimiento sale del abdomen bajo.", "La pelvis sube con control.", "Cuello relajado.", "La bajada es lenta."},
		Wrong:   []string{"Balancear las piernas.", "Usar impulso.", "Tirar del cuello.", "Dejar caer la cadera rápido."},
	},
	"C06": {
		Desc:    "Un mantenimiento estático del core con palanca acortada para mejor control.",
		HowTo:   []string{"Túmbate boca arriba.", "Presiona la zona lumbar contra el suelo.", "Eleva ligeramente los hombros.", "Rodillas flexionadas o piernas más altas.", "Mantén la posición."},
		Correct: []string{"La zona lumbar en el suelo.", "Abdomen firme.", "Cuello relajado.", "La respiración continúa."},
		Wrong:   []string{"Arquear la zona lumbar.", "Bajar demasiado las piernas.", "Llevar el mentón hacia delante.", "Aguantar la respiración."},
	},
	"C07": {
		Desc:    "Un ejercicio lateral del core para los oblicuos y la estabilidad de hombro y cadera.",
		HowTo:   []string{"Túmbate de lado.", "Coloca el codo bajo el hombro.", "Eleva la cadera.", "El cuerpo en una línea recta.", "Mantén."},
		Correct: []string{"Codo bajo el hombro.", "Cadera arriba.", "Cuerpo largo.", "Cuello neutro."},
		Wrong:   []string{"Dejar caer la cadera.", "Encoger el hombro hacia la oreja.", "Rodar el pecho hacia delante.", "Aguantar la respiración."},
	},
	"C08": {
		Desc:    "Una variación más exigente de la plancha lateral para los oblicuos y la cadera lateral.",
		HowTo:   []string{"Empieza en plancha lateral.", "Baja ligeramente la cadera.", "Vuelve a subir la cadera.", "Repite con control.", "Cambia de lado."},
		Correct: []string{"Movimiento pequeño y controlado.", "El codo bajo el hombro.", "La cadera se mueve en vertical.", "Core firme."},
		Wrong:   []string{"Bajar demasiado.", "Torcer el tronco.", "Hundirse en el hombro.", "Moverse demasiado rápido."},
	},
	"C10": {
		Desc:    "Un ejercicio de core y tríceps alternando entre plancha de antebrazos y plancha alta.",
		HowTo:   []string{"Empieza en plancha de antebrazos.", "Apoya una mano en el suelo y empuja hacia arriba.", "Sube la otra mano a la plancha alta.", "Baja de nuevo a los antebrazos.", "Alterna la mano que lidera."},
		Correct: []string{"La cadera estable.", "Core activo.", "Movimiento controlado.", "Hombros sobre las manos o los codos."},
		Wrong:   []string{"Balancear la cadera de lado a lado.", "Dejar caer la zona lumbar.", "Moverse demasiado rápido.", "Dejar que los hombros se hundan."},
	},
	"C11": {
		Desc:    "Un ejercicio de core controlado y una alternativa más segura a las elevaciones de pierna más duras.",
		HowTo:   []string{"Túmbate boca arriba.", "Eleva las piernas a posición de mesa.", "Baja un talón hacia el suelo.", "Toca ligeramente y vuelve.", "Alterna los lados."},
		Correct: []string{"La zona lumbar estable.", "Movimiento lento.", "El talón toca con suavidad.", "La pelvis quieta."},
		Wrong:   []string{"Dejar caer la pierna demasiado rápido.", "Arquear la zona lumbar.", "Mover las dos piernas a la vez.", "Apresurarse."},
	},
	"C12": {
		Desc:    "Un ejercicio de control del core para el abdomen profundo y la estabilidad de la columna.",
		HowTo:   []string{"Túmbate boca arriba.", "Eleva brazos y piernas a posición de mesa.", "Extiende el brazo y la pierna contrarios.", "Mantén la zona lumbar estable.", "Vuelve y cambia de lado."},
		Correct: []string{"La zona lumbar cerca del suelo.", "Movimiento lento.", "El brazo y la pierna contrarios se mueven juntos.", "El tronco quieto."},
		Wrong:   []string{"Arquear la zona lumbar.", "Moverse demasiado rápido.", "Bajar demasiado la pierna.", "Sacar las costillas."},
	},

	// ---- Flexiones / tríceps / hombros ----
	"P01": {
		Desc:    "Un ejercicio básico de tren superior para el pecho, el tríceps, los hombros y el core.",
		HowTo:   []string{"Empieza en plancha alta.", "Manos bajo los hombros o algo más anchas.", "Baja el cuerpo con control.", "Empuja hacia arriba manteniendo el cuerpo recto."},
		Correct: []string{"Cuerpo en una línea.", "Core firme.", "Codos a unos 30–45 grados.", "Movimiento controlado."},
		Wrong:   []string{"Dejar caer la cadera.", "Subir mucho la cadera.", "Abrir los codos hacia fuera.", "Caer en la posición baja."},
	},
	"P02": {
		Desc:    "Una variación de flexión con más énfasis en el tríceps.",
		HowTo:   []string{"Empieza en plancha alta.", "Coloca las manos más juntas que en una flexión normal.", "Mantén los codos cerca del cuerpo.", "Baja despacio y empuja hacia arriba."},
		Correct: []string{"Codos cerca.", "Cuerpo recto.", "Hombros lejos de las orejas.", "Muñecas cómodas."},
		Wrong:   []string{"Poner las manos demasiado juntas.", "Abrir los codos.", "Dejar caer la zona lumbar.", "Forzar con dolor de muñeca o codo."},
	},
	"P03": {
		Desc:    "Flexiones controladas con una breve pausa en la posición baja.",
		HowTo:   []string{"Empieza en posición de flexión normal.", "Baja despacio.", "Pausa un segundo cerca del fondo.", "Empuja hacia arriba sin perder la posición del cuerpo."},
		Correct: []string{"Pausa controlada.", "Core firme.", "Pecho activo.", "El cuerpo permanece recto."},
		Wrong:   []string{"Relajarse abajo.", "Dejar caer la cadera.", "Aguantar la respiración.", "Rebotar desde el fondo."},
	},
	"P04": {
		Desc:    "Una variación lenta de flexión que aumenta el control y el tiempo bajo tensión.",
		HowTo:   []string{"Empieza en plancha alta.", "Baja durante unos tres segundos.", "Pausa un segundo cerca del fondo.", "Empuja hacia arriba durante unos tres segundos."},
		Correct: []string{"Ritmo lento y uniforme.", "Cuerpo recto.", "Codos controlados.", "Respiración constante."},
		Wrong:   []string{"Apresurar el movimiento.", "Colapsar abajo.", "Dejar caer la cadera.", "Perder la tensión del core."},
	},
	"P05": {
		Desc:    "Una variación asimétrica de flexión que reta al pecho, los hombros y el core.",
		HowTo:   []string{"Empieza en posición de flexión.", "Coloca una mano algo adelantada y la otra algo atrasada.", "Baja con control.", "Empuja hacia arriba.", "Cambia la posición de las manos en el otro lado."},
		Correct: []string{"El cuerpo se mantiene recto.", "Core firme.", "Ambos hombros estables.", "Movimiento controlado."},
		Wrong:   []string{"Torcer el tronco.", "Dejar caer un hombro.", "Poner las manos demasiado anchas.", "Forzar el rango si el hombro molesta."},
	},
	"P06": {
		Desc:    "Una variación de flexión centrada en los hombros.",
		HowTo:   []string{"Empieza en plancha alta.", "Acerca un poco los pies y eleva la cadera.", "La cabeza entre los brazos.", "Baja la cabeza hacia el suelo.", "Empuja de nuevo hacia arriba."},
		Correct: []string{"Cadera alta.", "Los codos se flexionan con control.", "Hombros estables.", "Cuello neutro."},
		Wrong:   []string{"Caer en una flexión normal.", "Llevar la cabeza demasiado adelante.", "Abrir mucho los codos.", "Forzar con dolor de hombro."},
	},
	"P08": {
		Desc:    "Un ejercicio de tríceps y hombros desde una posición de apoyo en antebrazos.",
		HowTo:   []string{"Empieza sobre los antebrazos, el cuerpo largo.", "Codos bajo los hombros o algo por delante.", "Empuja con las manos y los antebrazos.", "Extiende ligeramente los codos.", "Vuelve con control."},
		Correct: []string{"Core firme.", "Los codos se mueven con suavidad.", "Hombros abajo.", "Rango de movimiento cómodo."},
		Wrong:   []string{"Forzar los codos.", "Dejar caer la cadera.", "Encoger los hombros.", "Moverse con dolor de codo o hombro."},
	},

	// ---- Espalda / postura ----
	"B01": {
		Desc:    "Un ejercicio de espalda centrado en la postura para la espalda alta, las escápulas y los extensores de la columna.",
		HowTo:   []string{"Túmbate boca abajo.", "Los brazos a lo largo del cuerpo.", "Gira ligeramente los pulgares hacia fuera.", "Eleva un poco el pecho.", "Lleva las escápulas atrás y abajo.", "Mantén."},
		Correct: []string{"La elevación es pequeña.", "Cuello neutro.", "Escápulas atrás y abajo.", "Glúteos ligeramente activos."},
		Wrong:   []string{"Elevar demasiado.", "Echar la cabeza atrás.", "Hiperextender la zona lumbar.", "Subir los hombros a las orejas."},
	},
	"B02": {
		Desc:    "Un ejercicio de espalda alta y escápulas boca abajo.",
		HowTo:   []string{"Túmbate boca abajo.", "La cabeza mira al suelo.", "Lleva los brazos en un arco amplio desde los lados hacia arriba.", "Vuelve por el mismo camino.", "Movimiento controlado."},
		Correct: []string{"Cuello relajado.", "Las escápulas se mueven con suavidad.", "La elevación del pecho es pequeña o neutra.", "Los brazos sin prisa."},
		Wrong:   []string{"Levantar la cabeza.", "Hiperextender mucho la zona lumbar.", "Balancear los brazos.", "Subir los hombros a las orejas."},
	},
	"B03": {
		Desc:    "Un ejercicio de espalda alta para las escápulas y los hombros posteriores.",
		HowTo:   []string{"Túmbate boca abajo.", "Flexiona los codos en forma de W.", "Eleva un poco los codos y las manos.", "Aprieta las escápulas.", "Baja con control."},
		Correct: []string{"Rango de movimiento pequeño.", "El trabajo lo hacen las escápulas.", "Cuello neutro.", "Hombros lejos de las orejas."},
		Wrong:   []string{"Elevar demasiado.", "Levantar la cabeza.", "Usar impulso.", "Hiperextender la zona lumbar."},
	},
	"B04": {
		Desc:    "Un ejercicio de escápulas para el trapecio inferior y la espalda alta.",
		HowTo:   []string{"Túmbate boca abajo.", "Extiende los brazos en diagonal hacia arriba en forma de Y.", "Eleva un poco los brazos del suelo.", "Baja con control."},
		Correct: []string{"Los brazos suben solo un poco.", "Cuello neutro.", "Hombros abajo.", "Movimiento controlado."},
		Wrong:   []string{"Elevar demasiado.", "Subir los hombros a las orejas.", "Arquear la zona lumbar.", "Moverse demasiado rápido."},
	},
	"B05": {
		Desc:    "Un ejercicio de espalda alta para los hombros posteriores y el control de las escápulas.",
		HowTo:   []string{"Túmbate boca abajo.", "Extiende los brazos a los lados en forma de T.", "Eleva un poco los brazos.", "Aprieta las escápulas.", "Baja lentamente."},
		Correct: []string{"Los brazos en la línea de los hombros.", "Cuello relajado.", "Las escápulas se mueven con suavidad.", "La elevación es pequeña."},
		Wrong:   []string{"Lanzar los brazos hacia arriba.", "Levantar la cabeza.", "Encoger los hombros.", "Forzar un rango amplio."},
	},
	"B06": {
		Desc:    "Un ejercicio de espalda que imita un movimiento de tracción sin equipo.",
		HowTo:   []string{"Túmbate boca abajo con los brazos por encima de la cabeza.", "Eleva un poco el pecho.", "Lleva los codos hacia abajo hacia las costillas.", "Aprieta las escápulas.", "Extiende los brazos de nuevo hacia delante."},
		Correct: []string{"Los codos bajan con control.", "Las escápulas se aprietan con suavidad.", "Cuello neutro.", "Zona lumbar cómoda."},
		Wrong:   []string{"Dar tirones con los brazos.", "Elevar demasiado.", "Echar la cabeza atrás.", "Subir los hombros a las orejas."},
	},
	"B07": {
		Desc:    "Un ejercicio de espalda y coordinación boca abajo.",
		HowTo:   []string{"Túmbate boca abajo.", "Eleva un poco el pecho.", "Mueve el brazo y la pierna contrarios en un patrón de nado controlado.", "El movimiento es pequeño y constante."},
		Correct: []string{"Cuello neutro.", "Movimiento controlado.", "Zona lumbar cómoda.", "La respiración continúa."},
		Wrong:   []string{"Moverse demasiado rápido.", "Elevar demasiado.", "Echar la cabeza atrás.", "Forzar la zona lumbar."},
	},
	"B08": {
		Desc:    "Un ejercicio de cadena posterior para los glúteos, los isquios, los hombros y la espalda.",
		HowTo:   []string{"Siéntate en el suelo con las piernas extendidas.", "Coloca las manos detrás de la cadera.", "Empuja con las manos y los talones.", "Eleva la cadera.", "Mantén el cuerpo en una línea larga."},
		Correct: []string{"Pecho abierto.", "Cadera arriba.", "Hombros estables.", "Cuello neutro."},
		Wrong:   []string{"Dejar caer la cadera.", "Encoger los hombros.", "Echar la cabeza atrás.", "Forzar las muñecas."},
	},
	"B09": {
		Desc:    "Un ejercicio de espalda de rango pequeño para los extensores de la columna y la postura.",
		HowTo:   []string{"Túmbate boca abajo.", "Las manos junto al cuerpo o ligeramente cerca del pecho.", "Eleva un poco el pecho.", "Haz pequeños impulsos controlados.", "Cuello neutro."},
		Correct: []string{"Elevación pequeña.", "Impulsos suaves.", "Cuello largo.", "Zona lumbar controlada."},
		Wrong:   []string{"Impulsar demasiado alto.", "Usar impulso.", "Mirar al frente.", "Sentir dolor agudo en la zona lumbar."},
	},
	"B10": {
		Desc:    "Un ejercicio de espalda más exigente que combina una pequeña elevación de cobra con una tracción de brazos en forma de W.",
		HowTo:   []string{"Túmbate boca abajo con los brazos al frente.", "Eleva un poco el pecho.", "Lleva los codos abajo y atrás en forma de W.", "Aprieta las escápulas.", "Extiende de nuevo hacia delante y baja con control."},
		Correct: []string{"Pequeña elevación del pecho.", "Tracción en W suave.", "Escápulas atrás y abajo.", "Cuello largo."},
		Wrong:   []string{"Hiperextender la zona lumbar.", "Tirar con el cuello.", "Moverse demasiado rápido.", "Subir los hombros a las orejas."},
	},

	// ---- Brazos / bíceps ----
	"A01": {
		Desc:    "Un ejercicio de bíceps sin equipo: una mano opone resistencia a la otra.",
		HowTo:   []string{"De pie o sentado, erguido.", "Flexiona un codo como en un curl.", "Con la otra mano presiona hacia abajo el antebrazo.", "Sube despacio contra la resistencia.", "Cambia de lado."},
		Correct: []string{"Resistencia constante.", "Movimiento lento.", "El codo cerca del cuerpo.", "El hombro abajo."},
		Wrong:   []string{"Moverse sin resistencia real.", "Dar tirones con el brazo.", "Subir el hombro.", "Aguantar la respiración."},
	},
	"A02": {
		Desc:    "Un mantenimiento estático de bíceps con autorresistencia.",
		HowTo:   []string{"Flexiona un codo a unos 90 grados.", "Con la otra mano presiona hacia abajo.", "El brazo que trabaja opone resistencia sin moverse.", "Mantén la tensión.", "Cambia de lado."},
		Correct: []string{"El ángulo del codo estable.", "Tensión constante.", "Hombros relajados.", "La respiración continúa."},
		Wrong:   []string{"Presionar a tirones.", "Torcer el tronco.", "Subir el hombro.", "Aguantar la respiración."},
	},
	"A03": {
		Desc:    "Un ejercicio isométrico de brazo y antebrazo con las manos una contra otra.",
		HowTo:   []string{"Engancha los dedos o las manos.", "Codos ligeramente flexionados.", "Tira de las manos para separarlas.", "Mantén una tensión constante.", "Relaja despacio."},
		Correct: []string{"Tensión controlada.", "Hombros abajo.", "Muñecas cómodas.", "Respiración tranquila."},
		Wrong:   []string{"Tirar de golpe.", "Encoger los hombros.", "Flexionar en exceso las muñecas.", "Aguantar la respiración."},
	},
	"A04": {
		Desc:    "Un ejercicio de bajada lenta para el bíceps con autorresistencia.",
		HowTo:   []string{"Empieza con un codo flexionado.", "Con la otra mano crea resistencia.", "Baja el antebrazo lentamente.", "Mantén la resistencia durante la bajada.", "Cambia de lado."},
		Correct: []string{"La bajada es lenta.", "Resistencia constante.", "El codo cerca del cuerpo.", "El hombro relajado."},
		Wrong:   []string{"Dejar caer el brazo rápido.", "No usar resistencia.", "Torcer el tronco.", "Aguantar la respiración."},
	},

	// ---- Piernas / glúteos ----
	"L01": {
		Desc:    "Un ejercicio de piernas controlado para los muslos, las caderas y la postura.",
		HowTo:   []string{"Pies algo más anchos que las caderas.", "Lleva las caderas hacia atrás.", "Flexiona las rodillas a una profundidad cómoda y poco profunda.", "Vuelve a ponerte de pie lentamente."},
		Correct: []string{"Rodillas en dirección a los dedos.", "Talones abajo.", "Espalda larga.", "Profundidad cómoda."},
		Wrong:   []string{"Bajar demasiado.", "Dejar caer las rodillas hacia dentro.", "Levantar los talones.", "Bajar rápido."},
	},
	"L02": {
		Desc:    "Una variación de sentadilla controlada con una breve pausa en un rango seguro.",
		HowTo:   []string{"De pie y estable.", "Baja a una sentadilla cómoda y poco profunda.", "Pausa breve.", "Vuelve a ponerte de pie con control."},
		Correct: []string{"Pausa estable.", "Rodillas en dirección a los dedos.", "Pecho abierto.", "Talones en el suelo."},
		Wrong:   []string{"Bajar demasiado.", "Rebotar en la posición baja.", "Dejar caer las rodillas hacia dentro.", "Aguantar la respiración."},
	},
	"L03": {
		Desc:    "Un ejercicio de movilidad para las caderas, los isquios y el control lumbar.",
		HowTo:   []string{"Ponte de pie, pies al ancho de caderas.", "Suaviza ligeramente las rodillas.", "Lleva las caderas hacia atrás.", "Baja el torso con la espalda recta.", "Vuelve empujando las caderas hacia delante."},
		Correct: []string{"La espalda se mantiene recta.", "El movimiento empieza en las caderas.", "Rodillas suaves, no muy flexionadas.", "El cuello sigue la línea de la columna."},
		Wrong:   []string{"Redondear la espalda.", "Convertirlo en sentadilla.", "Bajar demasiado.", "Subir mucho la cabeza."},
	},
	"L04": {
		Desc:    "Un ejercicio de glúteos más exigente con una pierna cada vez.",
		HowTo:   []string{"Túmbate boca arriba.", "Flexiona una rodilla con ese pie en el suelo.", "Extiende o eleva la otra pierna.", "Empuja con el talón que trabaja.", "Eleva la cadera y baja despacio."},
		Correct: []string{"La elevación la hace el glúteo que trabaja.", "La cadera se mantiene nivelada.", "Movimiento controlado.", "Zona lumbar neutra."},
		Wrong:   []string{"Torcer la pelvis.", "Empujar con los dedos del pie.", "Arquear la zona lumbar.", "Bajar demasiado rápido."},
	},
	"L05": {
		Desc:    "Un ejercicio de glúteos y estabilidad de la pelvis.",
		HowTo:   []string{"Empieza en posición de puente de glúteos.", "Mantén la cadera elevada.", "Eleva un pie ligeramente del suelo.", "Bájalo y cambia de lado.", "Mantén la pelvis estable."},
		Correct: []string{"La cadera nivelada.", "Glúteos activos.", "Movimiento lento.", "Zona lumbar cómoda."},
		Wrong:   []string{"Dejar caer la cadera.", "Balancearse de lado a lado.", "Moverse demasiado rápido.", "Arquear la zona lumbar."},
	},
	"L06": {
		Desc:    "Un ejercicio de glúteos que fortalece las caderas y la cadena posterior.",
		HowTo:   []string{"Túmbate boca arriba.", "Flexiona las rodillas, pies en el suelo.", "Eleva la cadera.", "Pausa arriba.", "Baja lentamente."},
		Correct: []string{"Empuja con los talones.", "Los glúteos aprietan arriba.", "Las costillas abajo.", "La zona lumbar no se hiperextiende."},
		Wrong:   []string{"Hiperextender la zona lumbar arriba.", "Levantar los talones.", "Dejar caer la cadera demasiado rápido.", "Poner los pies demasiado juntos."},
	},
	"L07": {
		Desc:    "Un ejercicio de pierna baja para las pantorrillas, los tobillos y el control del pie.",
		HowTo:   []string{"De pie, pies al ancho de caderas.", "Sube sobre las puntas de los pies.", "Pausa breve arriba.", "Baja los talones lentamente."},
		Correct: []string{"Cuerpo erguido.", "Movimiento suave.", "Tobillos alineados.", "Bajada controlada."},
		Wrong:   []string{"Dejar caer los talones rápido.", "Balancearse hacia delante.", "Volcar los tobillos hacia fuera.", "Usar impulso."},
	},
	"L09": {
		Desc:    "Un ejercicio de cadena lateral para los glúteos laterales, los oblicuos y la estabilidad de la cadera.",
		HowTo:   []string{"Empieza en plancha lateral.", "El cuerpo en una línea recta.", "Eleva un poco la pierna de arriba.", "Bájala con control.", "Repite y cambia de lado."},
		Correct: []string{"La cadera arriba.", "La pierna de arriba se mueve despacio.", "El tronco estable.", "El hombro fuerte."},
		Wrong:   []string{"Dejar caer la cadera.", "Balancear la pierna.", "Rodar el cuerpo hacia atrás.", "Hundirse en el hombro."},
	},

	// ---- Vuelta a la calma ----
	"CD01": {
		Desc:    "Un estiramiento suave del pecho y la parte frontal de los hombros.",
		HowTo:   []string{"Túmbate boca arriba.", "Abre los brazos a los lados.", "Relaja los hombros.", "Deja que el pecho se abra con suavidad.", "Respira despacio."},
		Correct: []string{"El pecho se abre con suavidad.", "Hombros relajados.", "Zona lumbar cómoda.", "Sin dolor en los hombros."},
		Wrong:   []string{"Forzar los brazos hacia abajo.", "Arquear la zona lumbar.", "Tensar el cuello.", "Estirar con dolor."},
	},
	"CD02": {
		Desc:    "Un estiramiento suave de la parte posterior del muslo.",
		HowTo:   []string{"Túmbate boca arriba.", "Eleva una pierna.", "Sujeta por detrás del muslo o la pantorrilla.", "Lleva la pierna suavemente hacia ti.", "Cambia de lado."},
		Correct: []string{"El estiramiento es suave.", "La rodilla puede quedar algo flexionada.", "Zona lumbar tranquila.", "Respiración lenta."},
		Wrong:   []string{"Tirar demasiado fuerte.", "Forzar la pierna recta.", "Elevar la cadera.", "Sentir dolor agudo detrás de la rodilla."},
	},
	"CD03": {
		Desc:    "Un estiramiento suave de rotación para la columna y el tronco.",
		HowTo:   []string{"Túmbate boca arriba.", "Flexiona las rodillas.", "Deja que las rodillas caigan suavemente a un lado.", "Mantén los hombros en el suelo.", "Vuelve al centro y cambia de lado."},
		Correct: []string{"Movimiento lento.", "Hombros en el suelo.", "Respiración relajada.", "El estiramiento cómodo."},
		Wrong:   []string{"Forzar las rodillas al suelo.", "Girar rápido.", "Levantar el hombro.", "Estirar con dolor."},
	},
	"CD04": {
		Desc:    "Un estiramiento suave de la parte frontal del cuerpo y una leve extensión de la espalda.",
		HowTo:   []string{"Túmbate boca abajo.", "Coloca los antebrazos en el suelo.", "Eleva el pecho con suavidad.", "Mantén el cuello largo.", "Respira despacio."},
		Correct: []string{"La elevación es leve.", "Hombros lejos de las orejas.", "Zona lumbar cómoda.", "Cuello neutro."},
		Wrong:   []string{"Hiperextender la zona lumbar.", "Echar la cabeza atrás.", "Encoger los hombros.", "Forzar la posición."},
	},
	"CD05": {
		Desc:    "Un ejercicio de respiración tranquila para terminar la sesión.",
		HowTo:   []string{"Túmbate boca arriba.", "Relaja los hombros y la mandíbula.", "Si es cómodo, pon una mano en el vientre.", "Inhala suavemente.", "Exhala lentamente."},
		Correct: []string{"Respiración tranquila.", "Hombros relajados.", "Cara relajada.", "El cuerpo se calma."},
		Wrong:   []string{"Respirar con demasiada fuerza.", "Aguantar la respiración.", "Tensar el cuello.", "Arquear la zona lumbar."},
	},
}
