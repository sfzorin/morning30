package content

// detailsFR is the French rich content for every exercise (current library IDs).
var detailsFR = map[string]Detail{
	// ---- Échauffement ----
	"W01": {
		Desc:    "Un échauffement d'épaules simple pour préparer le cou, les épaules et le haut du dos.",
		HowTo:   []string{"Tiens-toi droit, bras détendus.", "Fais de lents cercles d'épaules vers l'arrière.", "Puis lentement vers l'avant.", "Le mouvement reste souple et contrôlé."},
		Correct: []string{"Cou détendu.", "Les épaules bougent en douceur.", "Bras relâchés.", "Corps droit."},
		Wrong:   []string{"Hausser fortement les épaules.", "Bouger trop vite.", "Tendre le cou.", "Arrondir le haut du dos."},
	},
	"W02": {
		Desc:    "Un échauffement des épaules et du haut du dos qui active les deltoïdes et les stabilisateurs de l'épaule.",
		HowTo:   []string{"Tiens-toi droit et lève les bras à hauteur d'épaules.", "Fais de petits cercles vers l'avant.", "Puis de petits cercles vers l'arrière.", "Garde les cercles petits et contrôlés."},
		Correct: []string{"Bras à hauteur d'épaules.", "Les cercles sont petits.", "Épaules basses.", "Cou détendu."},
		Wrong:   []string{"Faire de grands moulinets.", "Lever les épaules vers les oreilles.", "Cambrer le bas du dos.", "Retenir sa respiration."},
	},
	"W03": {
		Desc:    "Un exercice de rotation doux pour le tronc et le haut de la colonne.",
		HowTo:   []string{"Pieds écartés de la largeur des hanches.", "Les hanches restent surtout face à l'avant.", "Tourne le haut du corps d'un côté.", "Reviens au centre et tourne de l'autre côté."},
		Correct: []string{"Les pieds restent stables.", "Le mouvement vient du tronc.", "La rotation est souple.", "Sans torsion des genoux."},
		Wrong:   []string{"Tourner les genoux avec le corps.", "Tourner trop vite.", "Se pencher en arrière.", "Forcer l'amplitude."},
	},
	"W04": {
		Desc:    "Un exercice de mobilité pour les hanches, les ischio-jambiers et le contrôle du bas du dos.",
		HowTo:   []string{"Tiens-toi droit, pieds écartés de la largeur des hanches.", "Assouplis légèrement les genoux.", "Pousse les hanches vers l'arrière.", "Abaisse le buste dos plat.", "Reviens en poussant les hanches vers l'avant."},
		Correct: []string{"Le dos reste droit.", "Le mouvement part des hanches.", "Genoux souples, peu fléchis.", "Le cou suit la ligne de la colonne."},
		Wrong:   []string{"Arrondir le dos.", "En faire un squat.", "Descendre trop bas.", "Lever la tête trop haut."},
	},
	"W05": {
		Desc:    "Un échauffement des jambes contrôlé dans une amplitude sûre et peu profonde.",
		HowTo:   []string{"Pieds un peu plus larges que les hanches.", "Pousse légèrement les hanches vers l'arrière.", "Fléchis les genoux seulement jusqu'à une profondeur confortable.", "Redresse-toi lentement."},
		Correct: []string{"Genoux dans l'axe des orteils.", "Talons au sol.", "Dos long.", "Profondeur modérée."},
		Wrong:   []string{"Descendre trop bas.", "Laisser les genoux rentrer.", "Lever les talons.", "Descendre trop vite."},
	},
	"W06": {
		Desc:    "Un échauffement pour les épaules, les poignets et le tronc.",
		HowTo:   []string{"Commence en planche haute.", "Mains sous les épaules.", "Décale le corps légèrement vers l'avant.", "Reviens à la position de départ.", "Le mouvement est petit et contrôlé."},
		Correct: []string{"Corps en une ligne.", "Épaules stables.", "Tronc gainé.", "Mouvement souple."},
		Wrong:   []string{"Laisser tomber les hanches.", "Pousser les épaules vers les oreilles.", "Aller trop loin vers l'avant.", "Laisser le bas du dos s'affaisser."},
	},

	// ---- Abdos / tronc ----
	"C01": {
		Desc:    "Un exercice de stabilité du tronc pour les abdos, les épaules, les fessiers et le tronc.",
		HowTo:   []string{"Place les coudes sous les épaules.", "Recule les pieds.", "Forme une ligne droite des épaules aux talons.", "Gaine les abdos et les fessiers.", "Tiens la position."},
		Correct: []string{"Corps en une ligne.", "Coudes sous les épaules.", "Tronc gainé.", "Respiration régulière."},
		Wrong:   []string{"Laisser tomber le bas du dos.", "Lever trop les hanches.", "Regarder devant.", "Retenir sa respiration."},
	},
	"C02": {
		Desc:    "Une planche courte et intense axée sur la tension de tout le corps.",
		HowTo:   []string{"Commence en planche sur les avant-bras.", "Gaine fortement les abdos.", "Serre les fessiers.", "Imagine tirer les coudes vers les orteils.", "Tiens avec un contrôle maximal."},
		Correct: []string{"Tension de tout le corps.", "Fessiers actifs.", "Tronc ferme.", "Respiration contrôlée."},
		Wrong:   []string{"La tenir comme une planche relâchée.", "Laisser les hanches s'affaisser.", "Retenir sa respiration.", "Essayer de tenir trop longtemps."},
	},
	"C03": {
		Desc:    "Un exercice du bas des abdos avec des mouvements de croisement contrôlés des jambes.",
		HowTo:   []string{"Allonge-toi sur le dos.", "Plaque doucement le bas du dos au sol.", "Lève les jambes.", "Croise une jambe par-dessus l'autre.", "Continue en alternant le croisement."},
		Correct: []string{"Bas du dos proche du sol.", "Jambes contrôlées.", "Cou détendu.", "Mouvement souple."},
		Wrong:   []string{"Cambrer le bas du dos.", "Descendre trop les jambes.", "Tirer la tête vers l'avant.", "Bouger trop vite."},
	},
	"C04": {
		Desc:    "Un exercice du bas des abdos avec de courts battements alternés des jambes.",
		HowTo:   []string{"Allonge-toi sur le dos.", "Plaque le bas du dos au sol.", "Lève les jambes.", "Monte et descends les jambes en courts battements alternés."},
		Correct: []string{"Petite amplitude.", "Bas du dos stable.", "Les jambes ne descendent pas trop.", "Respiration régulière."},
		Wrong:   []string{"Cambrer le bas du dos.", "Faire de grands battements.", "Tendre le cou.", "Retenir sa respiration."},
	},
	"C05": {
		Desc:    "Un exercice abdominal axé sur l'élévation du bassin avec contrôle.",
		HowTo:   []string{"Allonge-toi sur le dos.", "Fléchis les genoux et lève les pieds.", "Utilise les abdos pour enrouler légèrement le bassin vers le haut.", "Redescends lentement."},
		Correct: []string{"Le mouvement vient du bas des abdos.", "Le bassin monte avec contrôle.", "Cou détendu.", "La descente est lente."},
		Wrong:   []string{"Balancer les jambes.", "Utiliser l'élan.", "Tirer sur le cou.", "Laisser tomber le bassin vite."},
	},
	"C06": {
		Desc:    "Un maintien statique du tronc avec levier raccourci pour un meilleur contrôle.",
		HowTo:   []string{"Allonge-toi sur le dos.", "Plaque le bas du dos au sol.", "Lève légèrement les épaules.", "Genoux fléchis ou jambes plus hautes.", "Tiens la position."},
		Correct: []string{"Bas du dos au sol.", "Abdos fermes.", "Cou détendu.", "La respiration continue."},
		Wrong:   []string{"Cambrer le bas du dos.", "Descendre trop les jambes.", "Tirer le menton vers l'avant.", "Retenir sa respiration."},
	},
	"C07": {
		Desc:    "Un exercice latéral du tronc pour les obliques et la stabilité de l'épaule et de la hanche.",
		HowTo:   []string{"Allonge-toi sur le côté.", "Place le coude sous l'épaule.", "Lève les hanches.", "Le corps en une ligne droite.", "Tiens."},
		Correct: []string{"Coude sous l'épaule.", "Hanches levées.", "Corps long.", "Cou neutre."},
		Wrong:   []string{"Laisser tomber les hanches.", "Coller l'épaule à l'oreille.", "Rouler la poitrine vers l'avant.", "Retenir sa respiration."},
	},
	"C08": {
		Desc:    "Une variante plus exigeante de la planche latérale pour les obliques et la hanche latérale.",
		HowTo:   []string{"Commence en planche latérale.", "Abaisse légèrement les hanches.", "Relève les hanches.", "Répète avec contrôle.", "Change de côté."},
		Correct: []string{"Mouvement petit et contrôlé.", "Coude sous l'épaule.", "Les hanches montent à la verticale.", "Tronc ferme."},
		Wrong:   []string{"Descendre trop bas.", "Tordre le tronc.", "S'affaisser dans l'épaule.", "Bouger trop vite."},
	},
	"C10": {
		Desc:    "Un exercice de tronc et de triceps alternant entre planche sur avant-bras et planche haute.",
		HowTo:   []string{"Commence en planche sur les avant-bras.", "Pose une main au sol et pousse vers le haut.", "Monte l'autre main en planche haute.", "Redescends sur les avant-bras.", "Alterne la main qui mène."},
		Correct: []string{"Hanches stables.", "Tronc gainé.", "Mouvement contrôlé.", "Épaules au-dessus des mains ou des coudes."},
		Wrong:   []string{"Balancer les hanches d'un côté à l'autre.", "Laisser tomber le bas du dos.", "Bouger trop vite.", "Laisser les épaules s'affaisser."},
	},
	"C11": {
		Desc:    "Un exercice de tronc contrôlé et une alternative plus sûre aux relevés de jambes plus durs.",
		HowTo:   []string{"Allonge-toi sur le dos.", "Lève les jambes en position table.", "Abaisse un talon vers le sol.", "Touche légèrement et reviens.", "Alterne les côtés."},
		Correct: []string{"Bas du dos stable.", "Mouvement lent.", "Le talon touche en douceur.", "Bassin immobile."},
		Wrong:   []string{"Laisser tomber la jambe trop vite.", "Cambrer le bas du dos.", "Bouger les deux jambes à la fois.", "Se précipiter."},
	},
	"C12": {
		Desc:    "Un exercice de contrôle du tronc pour les abdos profonds et la stabilité de la colonne.",
		HowTo:   []string{"Allonge-toi sur le dos.", "Lève les bras et les jambes en position table.", "Étends le bras et la jambe opposés.", "Garde le bas du dos stable.", "Reviens et change de côté."},
		Correct: []string{"Bas du dos proche du sol.", "Mouvement lent.", "Le bras et la jambe opposés bougent ensemble.", "Le tronc reste immobile."},
		Wrong:   []string{"Cambrer le bas du dos.", "Bouger trop vite.", "Descendre trop la jambe.", "Laisser ressortir les côtes."},
	},

	// ---- Pompes / triceps / épaules ----
	"P01": {
		Desc:    "Un exercice de base du haut du corps pour la poitrine, les triceps, les épaules et le tronc.",
		HowTo:   []string{"Commence en planche haute.", "Mains sous les épaules ou un peu plus larges.", "Abaisse le corps avec contrôle.", "Pousse vers le haut en gardant le corps droit."},
		Correct: []string{"Corps en une ligne.", "Tronc gainé.", "Coudes à environ 30–45 degrés.", "Mouvement contrôlé."},
		Wrong:   []string{"Laisser tomber les hanches.", "Lever trop les hanches.", "Écarter les coudes droit sur les côtés.", "Tomber dans la position basse."},
	},
	"P02": {
		Desc:    "Une variante de pompe avec plus de travail des triceps.",
		HowTo:   []string{"Commence en planche haute.", "Place les mains plus rapprochées qu'en pompe normale.", "Garde les coudes près du corps.", "Descends lentement et pousse vers le haut."},
		Correct: []string{"Coudes proches.", "Corps droit.", "Épaules loin des oreilles.", "Poignets confortables."},
		Wrong:   []string{"Placer les mains trop près.", "Écarter les coudes.", "Laisser tomber le bas du dos.", "Forcer malgré une douleur au poignet ou au coude."},
	},
	"P03": {
		Desc:    "Des pompes contrôlées avec une courte pause en position basse.",
		HowTo:   []string{"Commence en position de pompe normale.", "Descends lentement.", "Marque une pause d'une seconde près du bas.", "Pousse vers le haut sans perdre la position du corps."},
		Correct: []string{"Pause contrôlée.", "Tronc gainé.", "Poitrine active.", "Le corps reste droit."},
		Wrong:   []string{"Se relâcher en bas.", "Laisser tomber les hanches.", "Retenir sa respiration.", "Rebondir depuis le bas."},
	},
	"P04": {
		Desc:    "Une variante lente de pompe qui augmente le contrôle et le temps sous tension.",
		HowTo:   []string{"Commence en planche haute.", "Descends pendant environ trois secondes.", "Pause d'une seconde près du bas.", "Pousse vers le haut pendant environ trois secondes."},
		Correct: []string{"Tempo lent et régulier.", "Corps droit.", "Coudes contrôlés.", "Respiration régulière."},
		Wrong:   []string{"Précipiter le mouvement.", "S'effondrer en bas.", "Laisser tomber les hanches.", "Perdre le gainage."},
	},
	"P05": {
		Desc:    "Une variante asymétrique de pompe qui sollicite la poitrine, les épaules et le tronc.",
		HowTo:   []string{"Commence en position de pompe.", "Place une main un peu en avant et l'autre un peu en arrière.", "Descends avec contrôle.", "Pousse vers le haut.", "Change la position des mains de l'autre côté."},
		Correct: []string{"Le corps reste droit.", "Tronc gainé.", "Les deux épaules stables.", "Mouvement contrôlé."},
		Wrong:   []string{"Tordre le tronc.", "Laisser tomber une épaule.", "Écarter trop les mains.", "Forcer l'amplitude si l'épaule gêne."},
	},
	"P06": {
		Desc:    "Une variante de pompe axée sur les épaules.",
		HowTo:   []string{"Commence en planche haute.", "Rapproche un peu les pieds et lève les hanches.", "La tête entre les bras.", "Abaisse la tête vers le sol.", "Pousse de nouveau vers le haut."},
		Correct: []string{"Hanches hautes.", "Les coudes se fléchissent avec contrôle.", "Épaules stables.", "Cou neutre."},
		Wrong:   []string{"Retomber en pompe normale.", "Amener la tête trop en avant.", "Écarter trop les coudes.", "Forcer malgré une douleur à l'épaule."},
	},
	"P08": {
		Desc:    "Un exercice de triceps et d'épaules depuis un appui sur les avant-bras.",
		HowTo:   []string{"Commence sur les avant-bras, le corps long.", "Coudes sous ou un peu devant les épaules.", "Pousse avec les mains et les avant-bras.", "Étends légèrement les coudes.", "Reviens avec contrôle."},
		Correct: []string{"Tronc gainé.", "Les coudes bougent en douceur.", "Épaules basses.", "Amplitude confortable."},
		Wrong:   []string{"Forcer les coudes.", "Laisser tomber les hanches.", "Hausser les épaules.", "Bouger malgré une douleur au coude ou à l'épaule."},
	},

	// ---- Dos / posture ----
	"B01": {
		Desc:    "Un exercice de dos axé sur la posture pour le haut du dos, les omoplates et les extenseurs de la colonne.",
		HowTo:   []string{"Allonge-toi sur le ventre.", "Les bras le long du corps.", "Tourne légèrement les pouces vers l'extérieur.", "Soulève un peu la poitrine.", "Tire les omoplates vers l'arrière et le bas.", "Tiens."},
		Correct: []string{"L'élévation est petite.", "Cou neutre.", "Omoplates vers l'arrière et le bas.", "Fessiers légèrement engagés."},
		Wrong:   []string{"Soulever trop haut.", "Rejeter la tête en arrière.", "Hyperextendre le bas du dos.", "Lever les épaules vers les oreilles."},
	},
	"B02": {
		Desc:    "Un exercice pour le haut du dos et les omoplates sur le ventre.",
		HowTo:   []string{"Allonge-toi sur le ventre.", "La tête regarde le sol.", "Déplace les bras en un large arc des côtés vers le haut.", "Reviens par le même chemin.", "Mouvement contrôlé."},
		Correct: []string{"Cou détendu.", "Les omoplates bougent en douceur.", "L'élévation de la poitrine est petite ou neutre.", "Les bras sans hâte."},
		Wrong:   []string{"Lever la tête.", "Hyperextendre trop le bas du dos.", "Balancer les bras.", "Lever les épaules vers les oreilles."},
	},
	"B03": {
		Desc:    "Un exercice du haut du dos pour les omoplates et l'arrière des épaules.",
		HowTo:   []string{"Allonge-toi sur le ventre.", "Fléchis les coudes en forme de W.", "Soulève un peu les coudes et les mains.", "Serre les omoplates.", "Redescends avec contrôle."},
		Correct: []string{"Petite amplitude.", "Le travail vient des omoplates.", "Cou neutre.", "Épaules loin des oreilles."},
		Wrong:   []string{"Soulever trop haut.", "Lever la tête.", "Utiliser l'élan.", "Hyperextendre le bas du dos."},
	},
	"B04": {
		Desc:    "Un exercice d'omoplates pour le trapèze inférieur et le haut du dos.",
		HowTo:   []string{"Allonge-toi sur le ventre.", "Étends les bras en diagonale vers le haut en forme de Y.", "Soulève un peu les bras du sol.", "Redescends avec contrôle."},
		Correct: []string{"Les bras ne montent qu'un peu.", "Cou neutre.", "Épaules basses.", "Mouvement contrôlé."},
		Wrong:   []string{"Soulever trop haut.", "Lever les épaules vers les oreilles.", "Cambrer le bas du dos.", "Bouger trop vite."},
	},
	"B05": {
		Desc:    "Un exercice du haut du dos pour l'arrière des épaules et le contrôle des omoplates.",
		HowTo:   []string{"Allonge-toi sur le ventre.", "Étends les bras sur les côtés en forme de T.", "Soulève un peu les bras.", "Serre les omoplates.", "Redescends lentement."},
		Correct: []string{"Les bras dans la ligne des épaules.", "Cou détendu.", "Les omoplates bougent doucement.", "L'élévation est petite."},
		Wrong:   []string{"Lancer les bras vers le haut.", "Lever la tête.", "Hausser les épaules.", "Forcer une grande amplitude."},
	},
	"B06": {
		Desc:    "Un exercice de dos qui imite un mouvement de tirage sans matériel.",
		HowTo:   []string{"Allonge-toi sur le ventre, bras au-dessus de la tête.", "Soulève un peu la poitrine.", "Tire les coudes vers le bas, vers les côtes.", "Serre les omoplates.", "Étends de nouveau les bras vers l'avant."},
		Correct: []string{"Les coudes descendent avec contrôle.", "Les omoplates se serrent doucement.", "Cou neutre.", "Bas du dos confortable."},
		Wrong:   []string{"Donner des à-coups avec les bras.", "Soulever trop haut.", "Rejeter la tête en arrière.", "Lever les épaules vers les oreilles."},
	},
	"B07": {
		Desc:    "Un exercice de dos et de coordination sur le ventre.",
		HowTo:   []string{"Allonge-toi sur le ventre.", "Soulève un peu la poitrine.", "Bouge le bras et la jambe opposés selon un schéma de nage contrôlé.", "Le mouvement est petit et régulier."},
		Correct: []string{"Cou neutre.", "Mouvement contrôlé.", "Bas du dos confortable.", "La respiration continue."},
		Wrong:   []string{"Bouger trop vite.", "Soulever trop haut.", "Rejeter la tête en arrière.", "Forcer le bas du dos."},
	},
	"B08": {
		Desc:    "Un exercice de chaîne postérieure pour les fessiers, les ischios, les épaules et le dos.",
		HowTo:   []string{"Assieds-toi au sol, jambes tendues.", "Place les mains derrière les hanches.", "Pousse avec les mains et les talons.", "Lève les hanches.", "Tiens le corps en une longue ligne."},
		Correct: []string{"Poitrine ouverte.", "Hanches levées.", "Épaules stables.", "Cou neutre."},
		Wrong:   []string{"Laisser tomber les hanches.", "Hausser les épaules.", "Rejeter la tête en arrière.", "Forcer les poignets."},
	},
	"B09": {
		Desc:    "Un exercice de dos de faible amplitude pour les extenseurs de la colonne et la posture.",
		HowTo:   []string{"Allonge-toi sur le ventre.", "Les mains le long du corps ou légèrement près de la poitrine.", "Soulève un peu la poitrine.", "Fais de petites pulsations contrôlées.", "Cou neutre."},
		Correct: []string{"Petite élévation.", "Pulsations douces.", "Cou long.", "Bas du dos contrôlé."},
		Wrong:   []string{"Pulser trop haut.", "Utiliser l'élan.", "Regarder devant.", "Ressentir une douleur vive dans le bas du dos."},
	},
	"B10": {
		Desc:    "Un exercice de dos plus exigeant combinant une petite élévation cobra avec un tirage des bras en forme de W.",
		HowTo:   []string{"Allonge-toi sur le ventre, bras devant.", "Soulève un peu la poitrine.", "Tire les coudes vers le bas et l'arrière en forme de W.", "Serre les omoplates.", "Étends de nouveau vers l'avant et redescends avec contrôle."},
		Correct: []string{"Petite élévation de la poitrine.", "Tirage en W souple.", "Omoplates vers l'arrière et le bas.", "Cou long."},
		Wrong:   []string{"Hyperextendre le bas du dos.", "Tirer avec le cou.", "Bouger trop vite.", "Hausser les épaules."},
	},

	// ---- Bras / biceps ----
	"A01": {
		Desc:    "Un exercice de biceps sans matériel : une main oppose une résistance à l'autre.",
		HowTo:   []string{"Debout ou assis, bien droit.", "Fléchis un coude comme pour un curl.", "Avec l'autre main, appuie vers le bas sur l'avant-bras.", "Monte lentement contre la résistance.", "Change de côté."},
		Correct: []string{"Résistance constante.", "Mouvement lent.", "Le coude près du corps.", "L'épaule basse."},
		Wrong:   []string{"Bouger sans réelle résistance.", "Donner des à-coups avec le bras.", "Lever l'épaule.", "Retenir sa respiration."},
	},
	"A02": {
		Desc:    "Un maintien statique de biceps en auto-résistance.",
		HowTo:   []string{"Fléchis un coude à environ 90 degrés.", "Avec l'autre main, appuie vers le bas.", "Le bras qui travaille résiste sans bouger.", "Maintiens la tension.", "Change de côté."},
		Correct: []string{"L'angle du coude reste stable.", "Tension constante.", "Épaules détendues.", "La respiration continue."},
		Wrong:   []string{"Appuyer par à-coups.", "Tordre le tronc.", "Lever l'épaule.", "Retenir sa respiration."},
	},
	"A03": {
		Desc:    "Un exercice isométrique de bras et d'avant-bras, les mains l'une contre l'autre.",
		HowTo:   []string{"Accroche les doigts ou les mains.", "Les coudes légèrement fléchis.", "Tire les mains pour les écarter.", "Maintiens une tension régulière.", "Relâche lentement."},
		Correct: []string{"Tension contrôlée.", "Épaules basses.", "Poignets confortables.", "Respiration calme."},
		Wrong:   []string{"Tirer d'un coup.", "Hausser les épaules.", "Trop fléchir les poignets.", "Retenir sa respiration."},
	},
	"A04": {
		Desc:    "Un exercice de descente lente pour le biceps en auto-résistance.",
		HowTo:   []string{"Commence avec un coude fléchi.", "Avec l'autre main, crée une résistance.", "Abaisse lentement l'avant-bras.", "Maintiens la résistance pendant la descente.", "Change de côté."},
		Correct: []string{"La descente est lente.", "Résistance constante.", "Le coude près du corps.", "L'épaule détendue."},
		Wrong:   []string{"Laisser tomber le bras vite.", "N'utiliser aucune résistance.", "Tordre le tronc.", "Retenir sa respiration."},
	},

	// ---- Jambes / fessiers ----
	"L01": {
		Desc:    "Un exercice de jambes contrôlé pour les cuisses, les hanches et la posture.",
		HowTo:   []string{"Pieds un peu plus larges que les hanches.", "Pousse les hanches vers l'arrière.", "Fléchis les genoux jusqu'à une profondeur confortable et peu profonde.", "Redresse-toi lentement."},
		Correct: []string{"Genoux dans l'axe des orteils.", "Talons au sol.", "Dos long.", "Profondeur confortable."},
		Wrong:   []string{"Descendre trop bas.", "Laisser les genoux rentrer.", "Lever les talons.", "Descendre vite."},
	},
	"L02": {
		Desc:    "Une variante de squat contrôlée avec une courte pause dans une amplitude sûre.",
		HowTo:   []string{"Debout et stable.", "Descends dans un squat confortable et peu profond.", "Pause brève.", "Redresse-toi avec contrôle."},
		Correct: []string{"Pause stable.", "Genoux dans l'axe des orteils.", "Poitrine ouverte.", "Talons au sol."},
		Wrong:   []string{"Descendre trop bas.", "Rebondir en position basse.", "Laisser les genoux rentrer.", "Retenir sa respiration."},
	},
	"L03": {
		Desc:    "Un exercice de mobilité pour les hanches, les ischio-jambiers et le contrôle du bas du dos.",
		HowTo:   []string{"Tiens-toi droit, pieds écartés de la largeur des hanches.", "Assouplis légèrement les genoux.", "Pousse les hanches vers l'arrière.", "Abaisse le buste dos plat.", "Reviens en poussant les hanches vers l'avant."},
		Correct: []string{"Le dos reste droit.", "Le mouvement part des hanches.", "Genoux souples, peu fléchis.", "Le cou suit la ligne de la colonne."},
		Wrong:   []string{"Arrondir le dos.", "En faire un squat.", "Descendre trop bas.", "Lever la tête trop haut."},
	},
	"L04": {
		Desc:    "Un exercice de fessiers plus exigeant, une jambe à la fois.",
		HowTo:   []string{"Allonge-toi sur le dos.", "Fléchis un genou, ce pied au sol.", "Étends ou lève l'autre jambe.", "Pousse avec le talon qui travaille.", "Lève les hanches et redescends lentement."},
		Correct: []string{"L'élévation vient du fessier qui travaille.", "Les hanches restent de niveau.", "Mouvement contrôlé.", "Bas du dos neutre."},
		Wrong:   []string{"Tordre le bassin.", "Pousser sur les orteils.", "Cambrer le bas du dos.", "Descendre trop vite."},
	},
	"L05": {
		Desc:    "Un exercice de fessiers et de stabilité du bassin.",
		HowTo:   []string{"Commence en position de pont fessier.", "Garde les hanches levées.", "Lève un pied légèrement du sol.", "Repose-le et change de côté.", "Garde le bassin stable."},
		Correct: []string{"Les hanches restent de niveau.", "Fessiers actifs.", "Mouvement lent.", "Bas du dos confortable."},
		Wrong:   []string{"Laisser tomber les hanches.", "Se balancer d'un côté à l'autre.", "Bouger trop vite.", "Cambrer le bas du dos."},
	},
	"L06": {
		Desc:    "Un exercice de fessiers qui renforce les hanches et la chaîne postérieure.",
		HowTo:   []string{"Allonge-toi sur le dos.", "Fléchis les genoux, pieds au sol.", "Lève les hanches.", "Pause en haut.", "Redescends lentement."},
		Correct: []string{"Pousse par les talons.", "Les fessiers serrent en haut.", "Les côtes vers le bas.", "Le bas du dos ne s'hyperextend pas."},
		Wrong:   []string{"Hyperextendre le bas du dos en haut.", "Lever les talons.", "Laisser tomber les hanches trop vite.", "Placer les pieds trop près."},
	},
	"L07": {
		Desc:    "Un exercice du bas de jambe pour les mollets, les chevilles et le contrôle du pied.",
		HowTo:   []string{"Debout, pieds écartés de la largeur des hanches.", "Monte sur la pointe des pieds.", "Brève pause en haut.", "Abaisse les talons lentement."},
		Correct: []string{"Corps droit.", "Mouvement souple.", "Chevilles alignées.", "Descente contrôlée."},
		Wrong:   []string{"Laisser tomber les talons vite.", "Se balancer vers l'avant.", "Faire basculer les chevilles vers l'extérieur.", "Utiliser l'élan."},
	},
	"L09": {
		Desc:    "Un exercice de chaîne latérale pour les fessiers latéraux, les obliques et la stabilité de la hanche.",
		HowTo:   []string{"Commence en planche latérale.", "Le corps en une ligne droite.", "Lève un peu la jambe du dessus.", "Abaisse-la avec contrôle.", "Répète et change de côté."},
		Correct: []string{"Hanches levées.", "La jambe du dessus bouge lentement.", "Tronc stable.", "Épaule solide."},
		Wrong:   []string{"Laisser tomber les hanches.", "Balancer la jambe.", "Rouler le corps vers l'arrière.", "S'affaisser dans l'épaule."},
	},

	// ---- Retour au calme ----
	"CD01": {
		Desc:    "Un étirement doux de la poitrine et de l'avant des épaules.",
		HowTo:   []string{"Allonge-toi sur le dos.", "Ouvre les bras sur les côtés.", "Détends les épaules.", "Laisse la poitrine s'ouvrir doucement.", "Respire lentement."},
		Correct: []string{"La poitrine s'ouvre en douceur.", "Épaules détendues.", "Bas du dos confortable.", "Pas de douleur aux épaules."},
		Wrong:   []string{"Forcer les bras vers le bas.", "Cambrer le bas du dos.", "Tendre le cou.", "Étirer malgré la douleur."},
	},
	"CD02": {
		Desc:    "Un étirement doux de l'arrière de la cuisse.",
		HowTo:   []string{"Allonge-toi sur le dos.", "Lève une jambe.", "Tiens derrière la cuisse ou le mollet.", "Amène la jambe doucement vers toi.", "Change de côté."},
		Correct: []string{"L'étirement est doux.", "Le genou peut rester légèrement fléchi.", "Bas du dos calme.", "Respiration lente."},
		Wrong:   []string{"Tirer trop fort.", "Forcer la jambe à se tendre.", "Lever les hanches.", "Ressentir une douleur vive derrière le genou."},
	},
	"CD03": {
		Desc:    "Un étirement de rotation doux pour la colonne et le tronc.",
		HowTo:   []string{"Allonge-toi sur le dos.", "Fléchis les genoux.", "Laisse les genoux tomber doucement d'un côté.", "Garde les épaules au sol.", "Reviens au centre et change de côté."},
		Correct: []string{"Mouvement lent.", "Épaules au sol.", "Respiration détendue.", "Étirement confortable."},
		Wrong:   []string{"Forcer les genoux au sol.", "Tourner vite.", "Lever l'épaule.", "Étirer malgré la douleur."},
	},
	"CD04": {
		Desc:    "Un étirement doux de l'avant du corps et une légère extension du dos.",
		HowTo:   []string{"Allonge-toi sur le ventre.", "Place les avant-bras au sol.", "Soulève la poitrine en douceur.", "Garde le cou long.", "Respire lentement."},
		Correct: []string{"L'élévation est légère.", "Épaules loin des oreilles.", "Bas du dos confortable.", "Cou neutre."},
		Wrong:   []string{"Hyperextendre le bas du dos.", "Rejeter la tête en arrière.", "Hausser les épaules.", "Forcer la position."},
	},
	"CD05": {
		Desc:    "Un exercice de respiration calme pour terminer la séance.",
		HowTo:   []string{"Allonge-toi sur le dos.", "Détends les épaules et la mâchoire.", "Si c'est confortable, pose une main sur le ventre.", "Inspire doucement.", "Expire lentement."},
		Correct: []string{"Respiration calme.", "Épaules détendues.", "Visage détendu.", "Le corps s'apaise."},
		Wrong:   []string{"Respirer trop fort.", "Retenir sa respiration.", "Tendre le cou.", "Cambrer le bas du dos."},
	},

	// ---- Ajouts du set Vlad (échauffement / cardio / pliométrie / fentes) ----
	"W07": {
		Desc:    "Un exercice d'ouverture calme : respiration profonde avec un grand étirement de tout le corps pour allonger la colonne.",
		HowTo:   []string{"Tiens-toi droit, pieds écartés de la largeur des hanches.", "Inspire et lève les deux bras au-dessus de la tête.", "Allonge doucement la colonne.", "Expire et redescends les bras en relâchant les épaules."},
		Correct: []string{"La respiration est lente et ample.", "Les épaules se relâchent à l'expiration.", "Les côtes restent basses, sans cambrer.", "Le mouvement est sans hâte."},
		Wrong:   []string{"Retenir sa respiration.", "Cambrer fortement le bas du dos.", "Hausser les épaules vers les oreilles.", "Précipiter l'étirement."},
	},
	"W08": {
		Desc:    "Un exercice de rotation doux pour échauffer la colonne et le tronc.",
		HowTo:   []string{"Pieds écartés de la largeur des hanches.", "Laisse les bras se balancer librement.", "Tourne le haut du corps d'un côté.", "Passe en douceur de l'autre côté."},
		Correct: []string{"Le mouvement vient du tronc.", "Les hanches restent surtout face à l'avant.", "Les pieds restent stables.", "Le rythme est souple et régulier."},
		Wrong:   []string{"Tourner fortement avec les genoux.", "Donner des à-coups avec les bras.", "Tourner trop vite.", "Retenir sa respiration."},
	},
	"W09": {
		Desc:    "Un échauffement dynamique des épaules et de la poitrine pour préparer le haut du corps aux pompes et aux sauts.",
		HowTo:   []string{"Tiens-toi droit, pieds écartés de la largeur des hanches.", "Balance les bras vers l'avant et vers l'arrière de façon contrôlée.", "Laisse la poitrine s'ouvrir naturellement quand les bras reviennent en arrière.", "Garde le mouvement souple et détendu."},
		Correct: []string{"Les bras bougent librement.", "Les épaules restent relâchées.", "La poitrine s'ouvre doucement.", "Le corps reste droit."},
		Wrong:   []string{"Balancer trop brusquement.", "Remonter les épaules vers les oreilles.", "Cambrer le bas du dos.", "Transformer le mouvement en élan incontrôlé."},
	},
	"C13": {
		Desc:    "Un exercice de tronc en rotation pour les obliques et les abdos.",
		HowTo:   []string{"Allonge-toi sur le dos.", "Ramène les genoux.", "Pose les mains légèrement près de la tête.", "Tourne le tronc et amène un coude vers le genou opposé.", "Change de côté dans un rythme contrôlé."},
		Correct: []string{"La rotation vient du tronc.", "Le bas du dos reste contrôlé.", "Le cou reste détendu.", "Le mouvement est souple."},
		Wrong:   []string{"Tirer sur le cou.", "Bouger seulement les coudes.", "Précipiter les répétitions.", "Laisser le bas du dos se cambrer."},
	},
	"B11": {
		Desc:    "Un exercice de dos qui lève le bras et la jambe opposés en alternance, comme une nage.",
		HowTo:   []string{"Allonge-toi sur le ventre, bras tendus au-dessus de la tête.", "Soulève un peu la poitrine et les jambes.", "Lève un bras et la jambe opposée.", "Redescends et change de côté en alternant en douceur."},
		Correct: []string{"L'élévation reste petite.", "Le cou suit la ligne de la colonne.", "Fessiers actifs.", "Mouvement régulier."},
		Wrong:   []string{"Forcer le bas du dos.", "Lever la tête.", "Bouger par à-coups.", "Retenir sa respiration."},
	},
	"L10": {
		Desc:    "Un squat au poids du corps fondamental pour les jambes, les hanches et les fessiers.",
		HowTo:   []string{"Pieds écartés d'environ la largeur des épaules.", "Pousse légèrement les hanches vers l'arrière.", "Fléchis les genoux et descends en squat.", "Garde la poitrine ouverte.", "Redresse-toi en poussant dans tout le pied."},
		Correct: []string{"Genoux dans l'axe des orteils.", "Talons au sol.", "Le dos reste long.", "La profondeur reste contrôlée."},
		Wrong:   []string{"Laisser les genoux rentrer.", "Lever les talons.", "Arrondir le dos.", "Descendre trop vite."},
	},
	"L11": {
		Desc:    "Un exercice du bas du corps pour les jambes, les fessiers et l'équilibre.",
		HowTo:   []string{"Tiens-toi droit.", "Recule un pied.", "Descends en fente avec contrôle.", "Pousse dans le pied avant pour revenir debout.", "Change de côté."},
		Correct: []string{"Le genou avant suit les orteils.", "Le buste reste droit.", "Le pas arrière est contrôlé.", "Le talon avant reste au sol."},
		Wrong:   []string{"Laisser le genou avant rentrer.", "Pencher trop en avant.", "Descendre trop vite.", "Pousser trop fort sur la jambe arrière."},
	},
	"L12": {
		Desc:    "Un squat sans charge fait lentement pour développer le contrôle et la force des jambes.",
		HowTo:   []string{"Pieds écartés de la largeur des épaules.", "Descends sur environ trois secondes.", "Marque une brève pause en bas.", "Redresse-toi avec contrôle."},
		Correct: []string{"Le tempo reste lent.", "Genoux dans l'axe des orteils.", "Talons au sol.", "Dos neutre."},
		Wrong:   []string{"Descendre vite.", "Rebondir depuis le bas.", "Laisser les genoux rentrer.", "Retenir sa respiration."},
	},
	"J01": {
		Desc:    "Un échauffement léger des chevilles et des mollets avec de petits sauts rapides.",
		HowTo:   []string{"Tiens-toi droit, pieds écartés d'environ la largeur des hanches.", "Garde les genoux légèrement souples.", "Fais de petits sauts rapides surtout avec les chevilles.", "Atterris sans bruit et recommence."},
		Correct: []string{"Les sauts restent bas.", "Les atterrissages sont doux et silencieux.", "Les genoux restent légèrement fléchis.", "Le corps reste droit."},
		Wrong:   []string{"Sauter trop haut.", "Atterrir bruyamment.", "Verrouiller les genoux.", "Faire basculer les chevilles vers l'intérieur ou l'extérieur."},
	},
	"J02": {
		Desc:    "Une alternative d'échauffement à faible impact aux montées de genoux.",
		HowTo:   []string{"Tiens-toi droit.", "Lève un genou à hauteur de hanche.", "Redescends-le avec contrôle.", "Change de côté et continue à marcher sur place.", "Garde les bras en mouvement naturellement."},
		Correct: []string{"Le buste reste droit.", "Les genoux montent avec contrôle.", "Les pieds se posent en douceur.", "La respiration reste régulière."},
		Wrong:   []string{"Se pencher en arrière.", "Balancer les jambes sans contrôle.", "Claquer les pieds au sol.", "Retenir sa respiration."},
	},
	"J03": {
		Desc:    "Un exercice cardio pour tout le corps qui élève le rythme cardiaque et échauffe les épaules, les hanches et les jambes.",
		HowTo:   []string{"Tiens-toi droit, bras le long du corps.", "Saute en écartant les pieds tout en levant les bras au-dessus de la tête.", "Reviens en sautant pieds joints et bras en bas.", "Recommence à un rythme régulier."},
		Correct: []string{"Atterris en douceur.", "Garde les genoux légèrement fléchis.", "Les bras bougent en souplesse.", "La respiration reste rythmée."},
		Wrong:   []string{"Atterrir lourdement.", "Verrouiller les genoux.", "Remonter les épaules.", "Aller trop vite et perdre le contrôle."},
	},
	"J04": {
		Desc:    "Un exercice de tronc et de cardio dynamique réalisé depuis la planche haute.",
		HowTo:   []string{"Commence en planche haute.", "Garde les mains sous les épaules.", "Ramène un genou vers la poitrine.", "Change de jambe en rythme.", "Garde les hanches stables."},
		Correct: []string{"Le corps reste en planche solide.", "Tronc gainé.", "Les épaules restent au-dessus des mains.", "Les genoux bougent avec contrôle."},
		Wrong:   []string{"Laisser les hanches rebondir vers le haut.", "Laisser tomber le bas du dos.", "Bouger les pieds de façon désordonnée.", "Laisser les épaules s'effondrer."},
	},
	"J05": {
		Desc:    "Un exercice pliométrique des jambes pour la puissance, le conditionnement et la coordination.",
		HowTo:   []string{"Pieds écartés d'environ la largeur des épaules.", "Descends en squat contrôlé.", "Saute vers le haut.", "Atterris en douceur, genoux légèrement fléchis.", "Reprends ta position avant le saut suivant."},
		Correct: []string{"Atterris sans bruit.", "Genoux dans l'axe des orteils.", "Poitrine ouverte.", "Chaque saut est contrôlé."},
		Wrong:   []string{"Atterrir lourdement.", "Laisser les genoux rentrer.", "Sauter avant que le squat soit stable.", "Se précipiter sans se replacer."},
	},
	"J06": {
		Desc:    "Un saut latéral pour les jambes, les hanches, l'équilibre et la coordination.",
		HowTo:   []string{"Tiens-toi sur une jambe.", "Saute latéralement sur l'autre jambe.", "Atterris en douceur et stabilise-toi.", "Recommence d'un côté à l'autre.", "Utilise les bras pour l'équilibre."},
		Correct: []string{"Les atterrissages sont silencieux.", "Le genou suit les orteils.", "Le buste reste contrôlé.", "Le mouvement d'un côté à l'autre reste fluide."},
		Wrong:   []string{"Atterrir en laissant le genou rentrer.", "Sauter trop loin trop tôt.", "Se laisser tomber dans l'atterrissage.", "Vriller le genou au contact."},
	},
	"J07": {
		Desc:    "Un exercice cardio qui élève le rythme cardiaque et active les hanches, le tronc et les jambes.",
		HowTo:   []string{"Tiens-toi droit.", "Cours sur place en montant les genoux haut.", "Bouge les bras naturellement.", "Garde un rythme rapide mais contrôlé."},
		Correct: []string{"Les genoux montent à hauteur de hanche.", "Les atterrissages restent légers.", "Tronc gainé.", "Le buste reste droit."},
		Wrong:   []string{"Se pencher en arrière.", "Atterrir lourdement.", "Laisser les pieds claquer au sol.", "Perdre la posture quand le rythme s'accélère."},
	},
	"J08": {
		Desc:    "Un exercice de conditionnement pour tout le corps combinant un squat, une planche et un saut. Dans ce programme, utilise la version sans pompe sauf si une pompe est spécifiquement prévue.",
		HowTo:   []string{"Tiens-toi droit.", "Descends en squat et pose les mains au sol.", "Saute ou recule les pieds en planche.", "Saute ou ramène les pieds vers l'avant.", "Redresse-toi et termine par un petit saut."},
		Correct: []string{"Les mains se posent sous les épaules.", "La position de planche est solide.", "Les pieds se posent en douceur.", "Le saut est contrôlé."},
		Wrong:   []string{"Laisser tomber le bas du dos en planche.", "Atterrir lourdement.", "Précipiter des répétitions bâclées.", "Laisser les genoux rentrer au retour des pieds."},
	},
	"J09": {
		Desc:    "Un exercice pliométrique des jambes avancé pour la puissance, le conditionnement et la coordination.",
		HowTo:   []string{"Commence en position de fente.", "Saute vers le haut.", "Change de jambe en l'air.", "Atterris en douceur en fente de l'autre côté.", "Retrouve ton équilibre avant la répétition suivante."},
		Correct: []string{"L'atterrissage est doux et contrôlé.", "Le genou avant suit les orteils.", "Le buste reste droit.", "L'amplitude reste sûre."},
		Wrong:   []string{"Atterrir lourdement.", "Laisser le genou avant rentrer.", "Descendre trop profondément.", "Aller plus vite que ce que tu peux contrôler."},
	},
	"CD07": {
		Desc:    "Un étirement reposant pour le dos, les hanches et les épaules pour décompresser.",
		HowTo:   []string{"Mets-toi à genoux et assieds les hanches vers les talons.", "Étends les bras vers l'avant sur le sol.", "Laisse le front se poser au sol.", "Respire lentement et détends-toi."},
		Correct: []string{"Les hanches descendent vers les talons.", "Le dos s'allonge en douceur.", "Les épaules se relâchent.", "La respiration reste lente."},
		Wrong:   []string{"Forcer les hanches vers le bas.", "Tendre les épaules.", "Retenir sa respiration.", "Forcer les genoux."},
	},
}
