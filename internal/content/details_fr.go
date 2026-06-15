package content

// detailsFR is the French rich content for every exercise (current library IDs).
var detailsFR = map[string]Detail{
	// ---- Échauffement ----
	"W01": {
		Desc:    "Un échauffement d'épaules simple pour préparer le cou, les épaules et le haut du dos.",
		HowTo:   []string{"Tiens-toi droit, bras détendus.", "Fais de lents cercles d'épaules vers l'arrière.", "Puis lentement vers l'avant.", "Le mouvement reste souple et contrôlé."},
		Correct: []string{"Garde le cou détendu.", "Fais bouger les épaules en douceur.", "Relâche bien les bras.", "Tiens-toi droit."},
		Wrong:   []string{"Ne hausse pas fortement les épaules.", "Ne bouge pas trop vite.", "Ne tends pas le cou.", "N'arrondis pas le haut du dos."},
	},
	"W02": {
		Desc:    "Un échauffement des épaules et du haut du dos qui active les deltoïdes et les stabilisateurs de l'épaule.",
		HowTo:   []string{"Tiens-toi droit et lève les bras à hauteur d'épaules.", "Fais de petits cercles vers l'avant.", "Puis de petits cercles vers l'arrière.", "Garde les cercles petits et contrôlés."},
		Correct: []string{"Garde les bras à hauteur d'épaules.", "Fais des cercles bien petits.", "Garde les épaules basses.", "Détends le cou."},
		Wrong:   []string{"Ne fais pas de grands moulinets.", "Ne lève pas les épaules vers les oreilles.", "Ne cambre pas le bas du dos.", "Ne retiens pas ta respiration."},
	},
	"W03": {
		Desc:    "Un exercice de rotation doux pour le tronc et le haut de la colonne.",
		HowTo:   []string{"Pieds écartés de la largeur des hanches.", "Les hanches restent surtout face à l'avant.", "Tourne le haut du corps d'un côté.", "Reviens au centre et tourne de l'autre côté."},
		Correct: []string{"Garde les pieds stables.", "Laisse le mouvement venir du tronc.", "Tourne en souplesse.", "Ne fais pas pivoter les genoux."},
		Wrong:   []string{"Ne fais pas tourner les genoux avec le corps.", "Ne tourne pas trop vite.", "Ne te penche pas en arrière.", "Ne force pas l'amplitude."},
	},
	"W04": {
		Desc:    "Un exercice de mobilité pour les hanches, les ischio-jambiers et le contrôle du bas du dos.",
		HowTo:   []string{"Tiens-toi droit, pieds écartés de la largeur des hanches.", "Assouplis légèrement les genoux.", "Pousse les hanches vers l'arrière.", "Abaisse le buste dos plat.", "Reviens en poussant les hanches vers l'avant."},
		Correct: []string{"Garde le dos droit.", "Pars bien des hanches.", "Garde les genoux souples, à peine fléchis.", "Laisse le cou suivre la ligne de la colonne."},
		Wrong:   []string{"N'arrondis pas le dos.", "N'en fais pas un squat.", "Ne descends pas trop bas.", "Ne lève pas la tête trop haut."},
	},
	"W05": {
		Desc:    "Un échauffement des jambes contrôlé dans une amplitude sûre et peu profonde.",
		HowTo:   []string{"Pieds un peu plus larges que les hanches.", "Pousse légèrement les hanches vers l'arrière.", "Fléchis les genoux seulement jusqu'à une profondeur confortable.", "Redresse-toi lentement."},
		Correct: []string{"Garde les genoux dans l'axe des orteils.", "Garde les talons au sol.", "Allonge bien le dos.", "Reste à une profondeur modérée."},
		Wrong:   []string{"Ne descends pas trop bas.", "Ne laisse pas les genoux rentrer vers l'intérieur.", "Ne lève pas les talons.", "Ne descends pas trop vite."},
	},
	"W06": {
		Desc:    "Un échauffement pour les épaules, les poignets et le tronc.",
		HowTo:   []string{"Commence en planche haute.", "Mains sous les épaules.", "Décale le corps légèrement vers l'avant.", "Reviens à la position de départ.", "Le mouvement est petit et contrôlé."},
		Correct: []string{"Garde le corps bien aligné.", "Stabilise les épaules.", "Gaine le tronc.", "Bouge en souplesse."},
		Wrong:   []string{"Ne laisse pas tomber les hanches.", "Ne pousse pas les épaules vers les oreilles.", "Ne va pas trop loin vers l'avant.", "Ne laisse pas le bas du dos s'affaisser."},
	},

	// ---- Abdos / tronc ----
	"C01": {
		Desc:    "Un exercice de stabilité du tronc pour les abdos, les épaules, les fessiers et le tronc.",
		HowTo:   []string{"Place les coudes sous les épaules.", "Recule les pieds.", "Forme une ligne droite des épaules aux talons.", "Gaine les abdos et les fessiers.", "Tiens la position."},
		Correct: []string{"Garde le corps bien aligné.", "Place les coudes sous les épaules.", "Gaine le tronc.", "Respire régulièrement."},
		Wrong:   []string{"Ne laisse pas tomber le bas du dos.", "Ne lève pas trop les hanches.", "Ne regarde pas devant toi.", "Ne retiens pas ta respiration."},
	},
	"C02": {
		Desc:    "Une planche courte et intense axée sur la tension de tout le corps.",
		HowTo:   []string{"Commence en planche sur les avant-bras.", "Gaine fortement les abdos.", "Serre les fessiers.", "Imagine tirer les coudes vers les orteils.", "Tiens avec un contrôle maximal."},
		Correct: []string{"Mets tout le corps en tension.", "Active les fessiers.", "Garde le tronc ferme.", "Contrôle ta respiration."},
		Wrong:   []string{"Ne la tiens pas comme une planche relâchée.", "Ne laisse pas les hanches s'affaisser.", "Ne retiens pas ta respiration.", "N'essaie pas de tenir trop longtemps."},
	},
	"C03": {
		Desc:    "Un exercice du bas des abdos avec des mouvements de croisement contrôlés des jambes.",
		HowTo:   []string{"Allonge-toi sur le dos.", "Plaque doucement le bas du dos au sol.", "Lève les jambes.", "Croise une jambe par-dessus l'autre.", "Continue en alternant le croisement."},
		Correct: []string{"Garde le bas du dos près du sol.", "Contrôle bien tes jambes.", "Détends le cou.", "Bouge en souplesse."},
		Wrong:   []string{"Ne cambre pas le bas du dos.", "Ne descends pas trop les jambes.", "Ne tire pas la tête vers l'avant.", "Ne bouge pas trop vite."},
	},
	"C04": {
		Desc:    "Un exercice du bas des abdos avec de courts battements alternés des jambes.",
		HowTo:   []string{"Allonge-toi sur le dos.", "Plaque le bas du dos au sol.", "Lève les jambes.", "Monte et descends les jambes en courts battements alternés."},
		Correct: []string{"Garde une petite amplitude.", "Garde le bas du dos stable.", "Ne descends pas trop les jambes.", "Respire régulièrement."},
		Wrong:   []string{"Ne cambre pas le bas du dos.", "Ne fais pas de grands battements.", "Ne tends pas le cou.", "Ne retiens pas ta respiration."},
	},
	"C05": {
		Desc:    "Un exercice abdominal axé sur l'élévation du bassin avec contrôle.",
		HowTo:   []string{"Allonge-toi sur le dos.", "Fléchis les genoux et lève les pieds.", "Utilise les abdos pour enrouler légèrement le bassin vers le haut.", "Redescends lentement."},
		Correct: []string{"Laisse le mouvement venir du bas des abdos.", "Fais monter le bassin avec contrôle.", "Détends le cou.", "Redescends lentement."},
		Wrong:   []string{"Ne balance pas les jambes.", "N'utilise pas l'élan.", "Ne tire pas sur le cou.", "Ne laisse pas le bassin retomber d'un coup."},
	},
	"C06": {
		Desc:    "Un maintien statique du tronc avec levier raccourci pour un meilleur contrôle.",
		HowTo:   []string{"Allonge-toi sur le dos.", "Plaque le bas du dos au sol.", "Lève légèrement les épaules.", "Genoux fléchis ou jambes plus hautes.", "Tiens la position."},
		Correct: []string{"Garde le bas du dos au sol.", "Garde les abdos fermes.", "Détends le cou.", "Continue de respirer."},
		Wrong:   []string{"Ne cambre pas le bas du dos.", "Ne descends pas trop les jambes.", "Ne tire pas le menton vers l'avant.", "Ne retiens pas ta respiration."},
	},
	"C07": {
		Desc:    "Un exercice latéral du tronc pour les obliques et la stabilité de l'épaule et de la hanche.",
		HowTo:   []string{"Allonge-toi sur le côté.", "Place le coude sous l'épaule.", "Lève les hanches.", "Le corps en une ligne droite.", "Tiens."},
		Correct: []string{"Place le coude sous l'épaule.", "Garde les hanches levées.", "Allonge bien le corps.", "Garde le cou neutre."},
		Wrong:   []string{"Ne laisse pas tomber les hanches.", "Ne colle pas l'épaule à l'oreille.", "Ne roule pas la poitrine vers l'avant.", "Ne retiens pas ta respiration."},
	},
	"C08": {
		Desc:    "Une variante plus exigeante de la planche latérale pour les obliques et la hanche latérale.",
		HowTo:   []string{"Commence en planche latérale.", "Abaisse légèrement les hanches.", "Relève les hanches.", "Répète avec contrôle.", "Change de côté."},
		Correct: []string{"Garde un mouvement petit et contrôlé.", "Place le coude sous l'épaule.", "Fais monter les hanches à la verticale.", "Garde le tronc ferme."},
		Wrong:   []string{"Ne descends pas trop bas.", "Ne tords pas le tronc.", "Ne t'affaisse pas dans l'épaule.", "Ne bouge pas trop vite."},
	},
	"C10": {
		Desc:    "Un exercice de tronc et de triceps alternant entre planche sur avant-bras et planche haute.",
		HowTo:   []string{"Commence en planche sur les avant-bras.", "Pose une main au sol et pousse vers le haut.", "Monte l'autre main en planche haute.", "Redescends sur les avant-bras.", "Alterne la main qui mène."},
		Correct: []string{"Garde les hanches stables.", "Gaine le tronc.", "Bouge avec contrôle.", "Garde les épaules au-dessus des mains ou des coudes."},
		Wrong:   []string{"Ne balance pas les hanches d'un côté à l'autre.", "Ne laisse pas tomber le bas du dos.", "Ne bouge pas trop vite.", "Ne laisse pas les épaules s'affaisser."},
	},
	"C11": {
		Desc:    "Un exercice de tronc contrôlé et une alternative plus sûre aux relevés de jambes plus durs.",
		HowTo:   []string{"Allonge-toi sur le dos.", "Lève les jambes en position table.", "Abaisse un talon vers le sol.", "Touche légèrement et reviens.", "Alterne les côtés."},
		Correct: []string{"Garde le bas du dos stable.", "Bouge lentement.", "Pose le talon en douceur.", "Garde le bassin immobile."},
		Wrong:   []string{"Ne laisse pas tomber la jambe trop vite.", "Ne cambre pas le bas du dos.", "Ne bouge pas les deux jambes à la fois.", "Ne te précipite pas."},
	},
	"C12": {
		Desc:    "Un exercice de contrôle du tronc pour les abdos profonds et la stabilité de la colonne.",
		HowTo:   []string{"Allonge-toi sur le dos.", "Lève les bras et les jambes en position table.", "Étends le bras et la jambe opposés.", "Garde le bas du dos stable.", "Reviens et change de côté."},
		Correct: []string{"Garde le bas du dos près du sol.", "Bouge lentement.", "Fais bouger le bras et la jambe opposés ensemble.", "Garde le tronc immobile."},
		Wrong:   []string{"Ne cambre pas le bas du dos.", "Ne bouge pas trop vite.", "Ne descends pas trop la jambe.", "Ne laisse pas ressortir les côtes."},
	},

	// ---- Pompes / triceps / épaules ----
	"P01": {
		Desc:    "Un exercice de base du haut du corps pour la poitrine, les triceps, les épaules et le tronc.",
		HowTo:   []string{"Commence en planche haute.", "Mains sous les épaules ou un peu plus larges.", "Abaisse le corps avec contrôle.", "Pousse vers le haut en gardant le corps droit."},
		Correct: []string{"Garde le corps bien aligné.", "Gaine le tronc.", "Garde les coudes à environ 30 à 45 degrés.", "Bouge avec contrôle."},
		Wrong:   []string{"Ne laisse pas tomber les hanches.", "Ne lève pas trop les hanches.", "N'écarte pas les coudes droit sur les côtés.", "Ne te laisse pas tomber en position basse."},
	},
	"P02": {
		Desc:    "Une variante de pompe avec plus de travail des triceps.",
		HowTo:   []string{"Commence en planche haute.", "Place les mains plus rapprochées qu'en pompe normale.", "Garde les coudes près du corps.", "Descends lentement et pousse vers le haut."},
		Correct: []string{"Garde les coudes proches du corps.", "Tiens le corps bien droit.", "Garde les épaules loin des oreilles.", "Garde les poignets confortables."},
		Wrong:   []string{"Ne place pas les mains trop près.", "N'écarte pas les coudes.", "Ne laisse pas tomber le bas du dos.", "Ne force pas malgré une douleur au poignet ou au coude."},
	},
	"P03": {
		Desc:    "Des pompes contrôlées avec une courte pause en position basse.",
		HowTo:   []string{"Commence en position de pompe normale.", "Descends lentement.", "Marque une pause d'une seconde près du bas.", "Pousse vers le haut sans perdre la position du corps."},
		Correct: []string{"Marque une pause contrôlée.", "Gaine le tronc.", "Garde la poitrine active.", "Garde le corps bien droit."},
		Wrong:   []string{"Ne te relâche pas en bas.", "Ne laisse pas tomber les hanches.", "Ne retiens pas ta respiration.", "Ne rebondis pas depuis le bas."},
	},
	"P04": {
		Desc:    "Une variante lente de pompe qui augmente le contrôle et le temps sous tension.",
		HowTo:   []string{"Commence en planche haute.", "Descends pendant environ trois secondes.", "Pause d'une seconde près du bas.", "Pousse vers le haut pendant environ trois secondes."},
		Correct: []string{"Garde un tempo lent et régulier.", "Tiens le corps bien droit.", "Contrôle bien les coudes.", "Respire régulièrement."},
		Wrong:   []string{"Ne précipite pas le mouvement.", "Ne t'effondre pas en bas.", "Ne laisse pas tomber les hanches.", "Ne perds pas le gainage."},
	},
	"P05": {
		Desc:    "Une variante asymétrique de pompe qui sollicite la poitrine, les épaules et le tronc.",
		HowTo:   []string{"Commence en position de pompe.", "Place une main un peu en avant et l'autre un peu en arrière.", "Descends avec contrôle.", "Pousse vers le haut.", "Change la position des mains de l'autre côté."},
		Correct: []string{"Garde le corps bien droit.", "Gaine le tronc.", "Garde les deux épaules stables.", "Bouge avec contrôle."},
		Wrong:   []string{"Ne tords pas le tronc.", "Ne laisse pas tomber une épaule.", "N'écarte pas trop les mains.", "Ne force pas l'amplitude si l'épaule gêne."},
	},
	"P06": {
		Desc:    "Une variante de pompe axée sur les épaules.",
		HowTo:   []string{"Commence en planche haute.", "Rapproche un peu les pieds et lève les hanches.", "La tête entre les bras.", "Abaisse la tête vers le sol.", "Pousse de nouveau vers le haut."},
		Correct: []string{"Garde les hanches hautes.", "Fléchis les coudes avec contrôle.", "Garde les épaules stables.", "Garde le cou neutre."},
		Wrong:   []string{"Ne retombe pas en pompe normale.", "N'amène pas la tête trop en avant.", "N'écarte pas trop les coudes.", "Ne force pas malgré une douleur à l'épaule."},
	},
	"P08": {
		Desc:    "Un exercice de triceps et d'épaules depuis un appui sur les avant-bras.",
		HowTo:   []string{"Commence sur les avant-bras, le corps long.", "Coudes sous ou un peu devant les épaules.", "Pousse avec les mains et les avant-bras.", "Étends légèrement les coudes.", "Reviens avec contrôle."},
		Correct: []string{"Gaine le tronc.", "Fais bouger les coudes en douceur.", "Garde les épaules basses.", "Reste dans une amplitude confortable."},
		Wrong:   []string{"Ne force pas sur les coudes.", "Ne laisse pas tomber les hanches.", "Ne hausse pas les épaules.", "Ne bouge pas malgré une douleur au coude ou à l'épaule."},
	},

	// ---- Dos / posture ----
	"B01": {
		Desc:    "Un exercice de dos axé sur la posture pour le haut du dos, les omoplates et les extenseurs de la colonne.",
		HowTo:   []string{"Allonge-toi sur le ventre.", "Les bras le long du corps.", "Tourne légèrement les pouces vers l'extérieur.", "Soulève un peu la poitrine.", "Tire les omoplates vers l'arrière et le bas.", "Tiens."},
		Correct: []string{"Soulève juste un peu la poitrine.", "Garde le cou neutre.", "Tire les omoplates vers l'arrière et le bas.", "Engage légèrement les fessiers."},
		Wrong:   []string{"Ne soulève pas trop haut.", "Ne rejette pas la tête en arrière.", "N'hyperextends pas le bas du dos.", "Ne lève pas les épaules vers les oreilles."},
	},
	"B02": {
		Desc:    "Un exercice pour le haut du dos et les omoplates sur le ventre.",
		HowTo:   []string{"Allonge-toi sur le ventre.", "La tête regarde le sol.", "Déplace les bras en un large arc des côtés vers le haut.", "Reviens par le même chemin.", "Mouvement contrôlé."},
		Correct: []string{"Détends le cou.", "Fais bouger les omoplates en douceur.", "Ne soulève la poitrine que très peu.", "Bouge les bras sans hâte."},
		Wrong:   []string{"Ne lève pas la tête.", "N'hyperextends pas trop le bas du dos.", "Ne balance pas les bras.", "Ne lève pas les épaules vers les oreilles."},
	},
	"B03": {
		Desc:    "Un exercice du haut du dos pour les omoplates et l'arrière des épaules.",
		HowTo:   []string{"Allonge-toi sur le ventre.", "Fléchis les coudes en forme de W.", "Soulève un peu les coudes et les mains.", "Serre les omoplates.", "Redescends avec contrôle."},
		Correct: []string{"Garde une petite amplitude.", "Laisse le travail venir des omoplates.", "Garde le cou neutre.", "Garde les épaules loin des oreilles."},
		Wrong:   []string{"Ne soulève pas trop haut.", "Ne lève pas la tête.", "N'utilise pas l'élan.", "N'hyperextends pas le bas du dos."},
	},
	"B04": {
		Desc:    "Un exercice d'omoplates pour le trapèze inférieur et le haut du dos.",
		HowTo:   []string{"Allonge-toi sur le ventre.", "Étends les bras en diagonale vers le haut en forme de Y.", "Soulève un peu les bras du sol.", "Redescends avec contrôle."},
		Correct: []string{"Ne fais monter les bras qu'un peu.", "Garde le cou neutre.", "Garde les épaules basses.", "Bouge avec contrôle."},
		Wrong:   []string{"Ne soulève pas trop haut.", "Ne lève pas les épaules vers les oreilles.", "Ne cambre pas le bas du dos.", "Ne bouge pas trop vite."},
	},
	"B05": {
		Desc:    "Un exercice du haut du dos pour l'arrière des épaules et le contrôle des omoplates.",
		HowTo:   []string{"Allonge-toi sur le ventre.", "Étends les bras sur les côtés en forme de T.", "Soulève un peu les bras.", "Serre les omoplates.", "Redescends lentement."},
		Correct: []string{"Garde les bras dans la ligne des épaules.", "Détends le cou.", "Fais bouger les omoplates doucement.", "Soulève juste un peu."},
		Wrong:   []string{"Ne lance pas les bras vers le haut.", "Ne lève pas la tête.", "Ne hausse pas les épaules.", "Ne force pas une grande amplitude."},
	},
	"B06": {
		Desc:    "Un exercice de dos qui imite un mouvement de tirage sans matériel.",
		HowTo:   []string{"Allonge-toi sur le ventre, bras au-dessus de la tête.", "Soulève un peu la poitrine.", "Tire les coudes vers le bas, vers les côtes.", "Serre les omoplates.", "Étends de nouveau les bras vers l'avant."},
		Correct: []string{"Fais descendre les coudes avec contrôle.", "Serre les omoplates doucement.", "Garde le cou neutre.", "Garde le bas du dos confortable."},
		Wrong:   []string{"Ne donne pas d'à-coups avec les bras.", "Ne soulève pas trop haut.", "Ne rejette pas la tête en arrière.", "Ne lève pas les épaules vers les oreilles."},
	},
	"B07": {
		Desc:    "Un exercice de dos et de coordination sur le ventre.",
		HowTo:   []string{"Allonge-toi sur le ventre.", "Soulève un peu la poitrine.", "Bouge le bras et la jambe opposés selon un schéma de nage contrôlé.", "Le mouvement est petit et régulier."},
		Correct: []string{"Garde le cou neutre.", "Bouge avec contrôle.", "Garde le bas du dos confortable.", "Continue de respirer."},
		Wrong:   []string{"Ne bouge pas trop vite.", "Ne soulève pas trop haut.", "Ne rejette pas la tête en arrière.", "Ne force pas sur le bas du dos."},
	},
	"B08": {
		Desc:    "Un exercice de chaîne postérieure pour les fessiers, les ischios, les épaules et le dos.",
		HowTo:   []string{"Assieds-toi au sol, jambes tendues.", "Place les mains derrière les hanches.", "Pousse avec les mains et les talons.", "Lève les hanches.", "Tiens le corps en une longue ligne."},
		Correct: []string{"Garde la poitrine ouverte.", "Garde les hanches levées.", "Garde les épaules stables.", "Garde le cou neutre."},
		Wrong:   []string{"Ne laisse pas tomber les hanches.", "Ne hausse pas les épaules.", "Ne rejette pas la tête en arrière.", "Ne force pas sur les poignets."},
	},
	"B09": {
		Desc:    "Un exercice de dos de faible amplitude pour les extenseurs de la colonne et la posture.",
		HowTo:   []string{"Allonge-toi sur le ventre.", "Les mains le long du corps ou légèrement près de la poitrine.", "Soulève un peu la poitrine.", "Fais de petites pulsations contrôlées.", "Cou neutre."},
		Correct: []string{"Garde une petite élévation.", "Fais des pulsations douces.", "Allonge bien le cou.", "Garde le bas du dos contrôlé."},
		Wrong:   []string{"Ne pulse pas trop haut.", "N'utilise pas l'élan.", "Ne regarde pas devant toi.", "Ne continue pas si tu ressens une douleur vive dans le bas du dos."},
	},
	"B10": {
		Desc:    "Un exercice de dos plus exigeant combinant une petite élévation cobra avec un tirage des bras en forme de W.",
		HowTo:   []string{"Allonge-toi sur le ventre, bras devant.", "Soulève un peu la poitrine.", "Tire les coudes vers le bas et l'arrière en forme de W.", "Serre les omoplates.", "Étends de nouveau vers l'avant et redescends avec contrôle."},
		Correct: []string{"Soulève juste un peu la poitrine.", "Fais un tirage en W souple.", "Tire les omoplates vers l'arrière et le bas.", "Allonge bien le cou."},
		Wrong:   []string{"N'hyperextends pas le bas du dos.", "Ne tire pas avec le cou.", "Ne bouge pas trop vite.", "Ne hausse pas les épaules."},
	},

	// ---- Bras / biceps ----
	"A01": {
		Desc:    "Un exercice de biceps sans matériel : une main oppose une résistance à l'autre.",
		HowTo:   []string{"Debout ou assis, bien droit.", "Fléchis un coude comme pour un curl.", "Avec l'autre main, appuie vers le bas sur l'avant-bras.", "Monte lentement contre la résistance.", "Change de côté."},
		Correct: []string{"Maintiens une résistance constante.", "Bouge lentement.", "Garde le coude près du corps.", "Garde l'épaule basse."},
		Wrong:   []string{"Ne bouge pas sans réelle résistance.", "Ne donne pas d'à-coups avec le bras.", "Ne lève pas l'épaule.", "Ne retiens pas ta respiration."},
	},
	"A02": {
		Desc:    "Un maintien statique de biceps en auto-résistance.",
		HowTo:   []string{"Fléchis un coude à environ 90 degrés.", "Avec l'autre main, appuie vers le bas.", "Le bras qui travaille résiste sans bouger.", "Maintiens la tension.", "Change de côté."},
		Correct: []string{"Garde l'angle du coude stable.", "Maintiens une tension constante.", "Détends les épaules.", "Continue de respirer."},
		Wrong:   []string{"N'appuie pas par à-coups.", "Ne tords pas le tronc.", "Ne lève pas l'épaule.", "Ne retiens pas ta respiration."},
	},
	"A03": {
		Desc:    "Un exercice isométrique de bras et d'avant-bras, les mains l'une contre l'autre.",
		HowTo:   []string{"Accroche les doigts ou les mains.", "Les coudes légèrement fléchis.", "Tire les mains pour les écarter.", "Maintiens une tension régulière.", "Relâche lentement."},
		Correct: []string{"Garde une tension contrôlée.", "Garde les épaules basses.", "Garde les poignets confortables.", "Respire calmement."},
		Wrong:   []string{"Ne tire pas d'un coup.", "Ne hausse pas les épaules.", "Ne fléchis pas trop les poignets.", "Ne retiens pas ta respiration."},
	},
	"A04": {
		Desc:    "Un exercice de descente lente pour le biceps en auto-résistance.",
		HowTo:   []string{"Commence avec un coude fléchi.", "Avec l'autre main, crée une résistance.", "Abaisse lentement l'avant-bras.", "Maintiens la résistance pendant la descente.", "Change de côté."},
		Correct: []string{"Descends lentement.", "Maintiens une résistance constante.", "Garde le coude près du corps.", "Garde l'épaule détendue."},
		Wrong:   []string{"Ne laisse pas tomber le bras vite.", "Ne te passe pas de résistance.", "Ne tords pas le tronc.", "Ne retiens pas ta respiration."},
	},

	// ---- Jambes / fessiers ----
	"L01": {
		Desc:    "Un exercice de jambes contrôlé pour les cuisses, les hanches et la posture.",
		HowTo:   []string{"Pieds un peu plus larges que les hanches.", "Pousse les hanches vers l'arrière.", "Fléchis les genoux jusqu'à une profondeur confortable et peu profonde.", "Redresse-toi lentement."},
		Correct: []string{"Garde les genoux dans l'axe des orteils.", "Garde les talons au sol.", "Allonge bien le dos.", "Reste à une profondeur confortable."},
		Wrong:   []string{"Ne descends pas trop bas.", "Ne laisse pas les genoux rentrer vers l'intérieur.", "Ne lève pas les talons.", "Ne descends pas vite."},
	},
	"L02": {
		Desc:    "Une variante de squat contrôlée avec une courte pause dans une amplitude sûre.",
		HowTo:   []string{"Debout et stable.", "Descends dans un squat confortable et peu profond.", "Pause brève.", "Redresse-toi avec contrôle."},
		Correct: []string{"Marque une pause stable.", "Garde les genoux dans l'axe des orteils.", "Garde la poitrine ouverte.", "Garde les talons au sol."},
		Wrong:   []string{"Ne descends pas trop bas.", "Ne rebondis pas en position basse.", "Ne laisse pas les genoux rentrer vers l'intérieur.", "Ne retiens pas ta respiration."},
	},
	"L03": {
		Desc:    "Un exercice de mobilité pour les hanches, les ischio-jambiers et le contrôle du bas du dos.",
		HowTo:   []string{"Tiens-toi droit, pieds écartés de la largeur des hanches.", "Assouplis légèrement les genoux.", "Pousse les hanches vers l'arrière.", "Abaisse le buste dos plat.", "Reviens en poussant les hanches vers l'avant."},
		Correct: []string{"Garde le dos droit.", "Pars bien des hanches.", "Garde les genoux souples, à peine fléchis.", "Laisse le cou suivre la ligne de la colonne."},
		Wrong:   []string{"N'arrondis pas le dos.", "N'en fais pas un squat.", "Ne descends pas trop bas.", "Ne lève pas la tête trop haut."},
	},
	"L04": {
		Desc:    "Un exercice de fessiers plus exigeant, une jambe à la fois.",
		HowTo:   []string{"Allonge-toi sur le dos.", "Fléchis un genou, ce pied au sol.", "Étends ou lève l'autre jambe.", "Pousse avec le talon qui travaille.", "Lève les hanches et redescends lentement."},
		Correct: []string{"Laisse l'élévation venir du fessier qui travaille.", "Garde les hanches de niveau.", "Bouge avec contrôle.", "Garde le bas du dos neutre."},
		Wrong:   []string{"Ne tords pas le bassin.", "Ne pousse pas sur les orteils.", "Ne cambre pas le bas du dos.", "Ne descends pas trop vite."},
	},
	"L05": {
		Desc:    "Un exercice de fessiers et de stabilité du bassin.",
		HowTo:   []string{"Commence en position de pont fessier.", "Garde les hanches levées.", "Lève un pied légèrement du sol.", "Repose-le et change de côté.", "Garde le bassin stable."},
		Correct: []string{"Garde les hanches de niveau.", "Active les fessiers.", "Bouge lentement.", "Garde le bas du dos confortable."},
		Wrong:   []string{"Ne laisse pas tomber les hanches.", "Ne te balance pas d'un côté à l'autre.", "Ne bouge pas trop vite.", "Ne cambre pas le bas du dos."},
	},
	"L06": {
		Desc:    "Un exercice de fessiers qui renforce les hanches et la chaîne postérieure.",
		HowTo:   []string{"Allonge-toi sur le dos.", "Fléchis les genoux, pieds au sol.", "Lève les hanches.", "Pause en haut.", "Redescends lentement."},
		Correct: []string{"Pousse bien par les talons.", "Serre les fessiers en haut.", "Garde les côtes vers le bas.", "Ne laisse pas le bas du dos s'hyperextendre."},
		Wrong:   []string{"N'hyperextends pas le bas du dos en haut.", "Ne lève pas les talons.", "Ne laisse pas tomber les hanches trop vite.", "Ne place pas les pieds trop près."},
	},
	"L07": {
		Desc:    "Un exercice du bas de jambe pour les mollets, les chevilles et le contrôle du pied.",
		HowTo:   []string{"Debout, pieds écartés de la largeur des hanches.", "Monte sur la pointe des pieds.", "Brève pause en haut.", "Abaisse les talons lentement."},
		Correct: []string{"Tiens le corps bien droit.", "Bouge en souplesse.", "Garde les chevilles alignées.", "Redescends avec contrôle."},
		Wrong:   []string{"Ne laisse pas tomber les talons vite.", "Ne te balance pas vers l'avant.", "Ne fais pas basculer les chevilles vers l'extérieur.", "N'utilise pas l'élan."},
	},
	"L09": {
		Desc:    "Un exercice de chaîne latérale pour les fessiers latéraux, les obliques et la stabilité de la hanche.",
		HowTo:   []string{"Commence en planche latérale.", "Le corps en une ligne droite.", "Lève un peu la jambe du dessus.", "Abaisse-la avec contrôle.", "Répète et change de côté."},
		Correct: []string{"Garde les hanches levées.", "Fais bouger la jambe du dessus lentement.", "Garde le tronc stable.", "Garde l'épaule solide."},
		Wrong:   []string{"Ne laisse pas tomber les hanches.", "Ne balance pas la jambe.", "Ne roule pas le corps vers l'arrière.", "Ne t'affaisse pas dans l'épaule."},
	},

	// ---- Retour au calme ----
	"CD01": {
		Desc:    "Un étirement doux de la poitrine et de l'avant des épaules.",
		HowTo:   []string{"Allonge-toi sur le dos.", "Ouvre les bras sur les côtés.", "Détends les épaules.", "Laisse la poitrine s'ouvrir doucement.", "Respire lentement."},
		Correct: []string{"Laisse la poitrine s'ouvrir en douceur.", "Détends les épaules.", "Garde le bas du dos confortable.", "Arrête si tu sens une douleur aux épaules."},
		Wrong:   []string{"Ne force pas les bras vers le bas.", "Ne cambre pas le bas du dos.", "Ne tends pas le cou.", "Ne t'étire pas malgré la douleur."},
	},
	"CD02": {
		Desc:    "Un étirement doux de l'arrière de la cuisse.",
		HowTo:   []string{"Allonge-toi sur le dos.", "Lève une jambe.", "Tiens derrière la cuisse ou le mollet.", "Amène la jambe doucement vers toi.", "Change de côté."},
		Correct: []string{"Étire-toi en douceur.", "Tu peux garder le genou légèrement fléchi.", "Garde le bas du dos calme.", "Respire lentement."},
		Wrong:   []string{"Ne tire pas trop fort.", "Ne force pas la jambe à se tendre.", "Ne lève pas les hanches.", "Ne continue pas si tu sens une douleur vive derrière le genou."},
	},
	"CD03": {
		Desc:    "Un étirement de rotation doux pour la colonne et le tronc.",
		HowTo:   []string{"Allonge-toi sur le dos.", "Fléchis les genoux.", "Laisse les genoux tomber doucement d'un côté.", "Garde les épaules au sol.", "Reviens au centre et change de côté."},
		Correct: []string{"Bouge lentement.", "Garde les épaules au sol.", "Respire de façon détendue.", "Reste dans un étirement confortable."},
		Wrong:   []string{"Ne force pas les genoux au sol.", "Ne tourne pas vite.", "Ne lève pas l'épaule.", "Ne t'étire pas malgré la douleur."},
	},
	"CD04": {
		Desc:    "Un étirement doux de l'avant du corps et une légère extension du dos.",
		HowTo:   []string{"Allonge-toi sur le ventre.", "Place les avant-bras au sol.", "Soulève la poitrine en douceur.", "Garde le cou long.", "Respire lentement."},
		Correct: []string{"Soulève juste un peu la poitrine.", "Garde les épaules loin des oreilles.", "Garde le bas du dos confortable.", "Garde le cou neutre."},
		Wrong:   []string{"N'hyperextends pas le bas du dos.", "Ne rejette pas la tête en arrière.", "Ne hausse pas les épaules.", "Ne force pas la position."},
	},
	"CD05": {
		Desc:    "Un exercice de respiration calme pour terminer la séance.",
		HowTo:   []string{"Allonge-toi sur le dos.", "Détends les épaules et la mâchoire.", "Si c'est confortable, pose une main sur le ventre.", "Inspire doucement.", "Expire lentement."},
		Correct: []string{"Respire calmement.", "Détends les épaules.", "Détends le visage.", "Laisse le corps s'apaiser."},
		Wrong:   []string{"Ne respire pas trop fort.", "Ne retiens pas ta respiration.", "Ne tends pas le cou.", "Ne cambre pas le bas du dos."},
	},

	// ---- Ajouts du set Vlad (échauffement / cardio / pliométrie / fentes) ----
	"W07": {
		Desc:    "Un exercice d'ouverture calme : respiration profonde avec un grand étirement de tout le corps pour allonger la colonne.",
		HowTo:   []string{"Tiens-toi droit, pieds écartés de la largeur des hanches.", "Inspire et lève les deux bras au-dessus de la tête.", "Allonge doucement la colonne.", "Expire et redescends les bras en relâchant les épaules."},
		Correct: []string{"Respire lentement et amplement.", "Relâche les épaules à l'expiration.", "Garde les côtes basses, sans cambrer.", "Bouge sans hâte."},
		Wrong:   []string{"Ne retiens pas ta respiration.", "Ne cambre pas fortement le bas du dos.", "Ne hausse pas les épaules vers les oreilles.", "Ne précipite pas l'étirement."},
	},
	"W08": {
		Desc:    "Un exercice de rotation doux pour échauffer la colonne et le tronc.",
		HowTo:   []string{"Pieds écartés de la largeur des hanches.", "Laisse les bras se balancer librement.", "Tourne le haut du corps d'un côté.", "Passe en douceur de l'autre côté."},
		Correct: []string{"Laisse le mouvement venir du tronc.", "Garde les hanches surtout face à l'avant.", "Garde les pieds stables.", "Garde un rythme souple et régulier."},
		Wrong:   []string{"Ne tourne pas fortement avec les genoux.", "Ne donne pas d'à-coups avec les bras.", "Ne tourne pas trop vite.", "Ne retiens pas ta respiration."},
	},
	"W09": {
		Desc:    "Un échauffement dynamique des épaules et de la poitrine pour préparer le haut du corps aux pompes et aux sauts.",
		HowTo:   []string{"Tiens-toi droit, pieds écartés de la largeur des hanches.", "Balance les bras vers l'avant et vers l'arrière de façon contrôlée.", "Laisse la poitrine s'ouvrir naturellement quand les bras reviennent en arrière.", "Garde le mouvement souple et détendu."},
		Correct: []string{"Fais bouger les bras librement.", "Garde les épaules relâchées.", "Laisse la poitrine s'ouvrir doucement.", "Garde le corps bien droit."},
		Wrong:   []string{"Ne balance pas trop brusquement.", "Ne remonte pas les épaules vers les oreilles.", "Ne cambre pas le bas du dos.", "Ne transforme pas le mouvement en élan incontrôlé."},
	},
	"C13": {
		Desc:    "Un exercice de tronc en rotation pour les obliques et les abdos.",
		HowTo:   []string{"Allonge-toi sur le dos.", "Ramène les genoux.", "Pose les mains légèrement près de la tête.", "Tourne le tronc et amène un coude vers le genou opposé.", "Change de côté dans un rythme contrôlé."},
		Correct: []string{"Laisse la rotation venir du tronc.", "Garde le bas du dos contrôlé.", "Garde le cou détendu.", "Bouge en souplesse."},
		Wrong:   []string{"Ne tire pas sur le cou.", "Ne bouge pas seulement les coudes.", "Ne précipite pas les répétitions.", "Ne laisse pas le bas du dos se cambrer."},
	},
	"B11": {
		Desc:    "Un exercice de dos qui lève le bras et la jambe opposés en alternance, comme une nage.",
		HowTo:   []string{"Allonge-toi sur le ventre, bras tendus au-dessus de la tête.", "Soulève un peu la poitrine et les jambes.", "Lève un bras et la jambe opposée.", "Redescends et change de côté en alternant en douceur."},
		Correct: []string{"Garde une petite élévation.", "Laisse le cou suivre la ligne de la colonne.", "Active les fessiers.", "Bouge régulièrement."},
		Wrong:   []string{"Ne force pas sur le bas du dos.", "Ne lève pas la tête.", "Ne bouge pas par à-coups.", "Ne retiens pas ta respiration."},
	},
	"L10": {
		Desc:    "Un squat au poids du corps fondamental pour les jambes, les hanches et les fessiers.",
		HowTo:   []string{"Pieds écartés d'environ la largeur des épaules.", "Pousse légèrement les hanches vers l'arrière.", "Fléchis les genoux et descends en squat.", "Garde la poitrine ouverte.", "Redresse-toi en poussant dans tout le pied."},
		Correct: []string{"Garde les genoux dans l'axe des orteils.", "Garde les talons au sol.", "Allonge bien le dos.", "Garde la profondeur sous contrôle."},
		Wrong:   []string{"Ne laisse pas les genoux rentrer vers l'intérieur.", "Ne lève pas les talons.", "N'arrondis pas le dos.", "Ne descends pas trop vite."},
	},
	"L11": {
		Desc:    "Un exercice du bas du corps pour les jambes, les fessiers et l'équilibre.",
		HowTo:   []string{"Tiens-toi droit.", "Recule un pied.", "Descends en fente avec contrôle.", "Pousse dans le pied avant pour revenir debout.", "Change de côté."},
		Correct: []string{"Garde le genou avant dans l'axe des orteils.", "Garde le buste bien droit.", "Recule le pied avec contrôle.", "Garde le talon avant au sol."},
		Wrong:   []string{"Ne laisse pas le genou avant rentrer vers l'intérieur.", "Ne te penche pas trop en avant.", "Ne descends pas trop vite.", "Ne pousse pas trop fort sur la jambe arrière."},
	},
	"L12": {
		Desc:    "Un squat sans charge fait lentement pour développer le contrôle et la force des jambes.",
		HowTo:   []string{"Pieds écartés de la largeur des épaules.", "Descends sur environ trois secondes.", "Marque une brève pause en bas.", "Redresse-toi avec contrôle."},
		Correct: []string{"Garde un tempo lent.", "Garde les genoux dans l'axe des orteils.", "Garde les talons au sol.", "Garde le dos neutre."},
		Wrong:   []string{"Ne descends pas vite.", "Ne rebondis pas depuis le bas.", "Ne laisse pas les genoux rentrer vers l'intérieur.", "Ne retiens pas ta respiration."},
	},
	"J01": {
		Desc:    "Un échauffement léger des chevilles et des mollets avec de petits sauts rapides.",
		HowTo:   []string{"Tiens-toi droit, pieds écartés d'environ la largeur des hanches.", "Garde les genoux légèrement souples.", "Fais de petits sauts rapides surtout avec les chevilles.", "Atterris sans bruit et recommence."},
		Correct: []string{"Garde les sauts bas.", "Atterris en douceur et sans bruit.", "Garde les genoux légèrement fléchis.", "Garde le corps bien droit."},
		Wrong:   []string{"Ne saute pas trop haut.", "N'atterris pas bruyamment.", "Ne verrouille pas les genoux.", "Ne fais pas basculer les chevilles vers l'intérieur ou l'extérieur."},
	},
	"J02": {
		Desc:    "Une alternative d'échauffement à faible impact aux montées de genoux.",
		HowTo:   []string{"Tiens-toi droit.", "Lève un genou à hauteur de hanche.", "Redescends-le avec contrôle.", "Change de côté et continue à marcher sur place.", "Garde les bras en mouvement naturellement."},
		Correct: []string{"Garde le buste bien droit.", "Fais monter les genoux avec contrôle.", "Pose les pieds en douceur.", "Respire régulièrement."},
		Wrong:   []string{"Ne te penche pas en arrière.", "Ne balance pas les jambes sans contrôle.", "Ne claque pas les pieds au sol.", "Ne retiens pas ta respiration."},
	},
	"J03": {
		Desc:    "Un exercice cardio pour tout le corps qui élève le rythme cardiaque et échauffe les épaules, les hanches et les jambes.",
		HowTo:   []string{"Tiens-toi droit, bras le long du corps.", "Saute en écartant les pieds tout en levant les bras au-dessus de la tête.", "Reviens en sautant pieds joints et bras en bas.", "Recommence à un rythme régulier."},
		Correct: []string{"Atterris en douceur.", "Garde les genoux légèrement fléchis.", "Fais bouger les bras en souplesse.", "Garde une respiration rythmée."},
		Wrong:   []string{"N'atterris pas lourdement.", "Ne verrouille pas les genoux.", "Ne remonte pas les épaules.", "Ne va pas trop vite au point de perdre le contrôle."},
	},
	"J04": {
		Desc:    "Un exercice de tronc et de cardio dynamique réalisé depuis la planche haute.",
		HowTo:   []string{"Commence en planche haute.", "Garde les mains sous les épaules.", "Ramène un genou vers la poitrine.", "Change de jambe en rythme.", "Garde les hanches stables."},
		Correct: []string{"Garde le corps en planche solide.", "Gaine le tronc.", "Garde les épaules au-dessus des mains.", "Fais bouger les genoux avec contrôle."},
		Wrong:   []string{"Ne laisse pas les hanches rebondir vers le haut.", "Ne laisse pas tomber le bas du dos.", "Ne bouge pas les pieds de façon désordonnée.", "Ne laisse pas les épaules s'effondrer."},
	},
	"J05": {
		Desc:    "Un exercice pliométrique des jambes pour la puissance, le conditionnement et la coordination.",
		HowTo:   []string{"Pieds écartés d'environ la largeur des épaules.", "Descends en squat contrôlé.", "Saute vers le haut.", "Atterris en douceur, genoux légèrement fléchis.", "Reprends ta position avant le saut suivant."},
		Correct: []string{"Atterris sans bruit.", "Garde les genoux dans l'axe des orteils.", "Garde la poitrine ouverte.", "Contrôle bien chaque saut."},
		Wrong:   []string{"N'atterris pas lourdement.", "Ne laisse pas les genoux rentrer vers l'intérieur.", "Ne saute pas avant que le squat soit stable.", "Ne te précipite pas sans te replacer."},
	},
	"J06": {
		Desc:    "Un saut latéral pour les jambes, les hanches, l'équilibre et la coordination.",
		HowTo:   []string{"Tiens-toi sur une jambe.", "Saute latéralement sur l'autre jambe.", "Atterris en douceur et stabilise-toi.", "Recommence d'un côté à l'autre.", "Utilise les bras pour l'équilibre."},
		Correct: []string{"Atterris en silence.", "Garde le genou dans l'axe des orteils.", "Garde le buste contrôlé.", "Reste fluide d'un côté à l'autre."},
		Wrong:   []string{"Ne laisse pas le genou rentrer à l'atterrissage.", "Ne saute pas trop loin trop tôt.", "Ne te laisse pas tomber dans l'atterrissage.", "Ne vrille pas le genou au contact."},
	},
	"J07": {
		Desc:    "Un exercice cardio qui élève le rythme cardiaque et active les hanches, le tronc et les jambes.",
		HowTo:   []string{"Tiens-toi droit.", "Cours sur place en montant les genoux haut.", "Bouge les bras naturellement.", "Garde un rythme rapide mais contrôlé."},
		Correct: []string{"Fais monter les genoux à hauteur de hanche.", "Garde des atterrissages légers.", "Gaine le tronc.", "Garde le buste bien droit."},
		Wrong:   []string{"Ne te penche pas en arrière.", "N'atterris pas lourdement.", "Ne laisse pas les pieds claquer au sol.", "Ne perds pas la posture quand le rythme s'accélère."},
	},
	"J08": {
		Desc:    "Un exercice de conditionnement pour tout le corps combinant un squat, une planche et un saut. Dans ce programme, utilise la version sans pompe sauf si une pompe est spécifiquement prévue.",
		HowTo:   []string{"Tiens-toi droit.", "Descends en squat et pose les mains au sol.", "Saute ou recule les pieds en planche.", "Saute ou ramène les pieds vers l'avant.", "Redresse-toi et termine par un petit saut."},
		Correct: []string{"Pose les mains sous les épaules.", "Garde une position de planche solide.", "Pose les pieds en douceur.", "Contrôle bien le saut."},
		Wrong:   []string{"Ne laisse pas tomber le bas du dos en planche.", "N'atterris pas lourdement.", "Ne précipite pas des répétitions bâclées.", "Ne laisse pas les genoux rentrer au retour des pieds."},
	},
	"J09": {
		Desc:    "Un exercice pliométrique des jambes avancé pour la puissance, le conditionnement et la coordination.",
		HowTo:   []string{"Commence en position de fente.", "Saute vers le haut.", "Change de jambe en l'air.", "Atterris en douceur en fente de l'autre côté.", "Retrouve ton équilibre avant la répétition suivante."},
		Correct: []string{"Atterris en douceur et avec contrôle.", "Garde le genou avant dans l'axe des orteils.", "Garde le buste bien droit.", "Reste dans une amplitude sûre."},
		Wrong:   []string{"N'atterris pas lourdement.", "Ne laisse pas le genou avant rentrer vers l'intérieur.", "Ne descends pas trop profondément.", "Ne va pas plus vite que ce que tu peux contrôler."},
	},
	"CD07": {
		Desc:    "Un étirement reposant pour le dos, les hanches et les épaules pour décompresser.",
		HowTo:   []string{"Mets-toi à genoux et assieds les hanches vers les talons.", "Étends les bras vers l'avant sur le sol.", "Laisse le front se poser au sol.", "Respire lentement et détends-toi."},
		Correct: []string{"Laisse les hanches descendre vers les talons.", "Laisse le dos s'allonger en douceur.", "Relâche les épaules.", "Respire lentement."},
		Wrong:   []string{"Ne force pas les hanches vers le bas.", "Ne tends pas les épaules.", "Ne retiens pas ta respiration.", "Ne force pas sur les genoux."},
	},
}
