package content

// detailsEN is the English rich content for every exercise, keyed by the current
// library IDs. It is the authoritative card text; other languages fall back here
// until translated.
var detailsEN = map[string]Detail{
	// ---- Warm-up ----
	"W01": {
		Desc:    "A simple shoulder warm-up to prepare the neck, shoulders and upper back.",
		HowTo:   []string{"Stand tall with the arms relaxed.", "Make slow circles with the shoulders backward.", "Then make slow circles forward.", "Keep the movement smooth and controlled."},
		Correct: []string{"Relax your neck.", "Move your shoulders smoothly.", "Keep your arms loose.", "Stand up tall."},
		Wrong:   []string{"Don't shrug hard.", "Don't move too fast.", "Don't tense your neck.", "Don't round your upper back."},
	},
	"W02": {
		Desc:    "A shoulder and upper-back warm-up that activates the deltoids and shoulder stabilizers.",
		HowTo:   []string{"Stand tall and raise the arms to shoulder height.", "Make small circles forward.", "Then make small circles backward.", "Keep the circles tight and controlled."},
		Correct: []string{"Keep your arms near shoulder height.", "Keep the circles small.", "Keep your shoulders down.", "Relax your neck."},
		Wrong:   []string{"Don't make large swinging motions.", "Don't raise your shoulders to your ears.", "Don't arch your lower back.", "Don't hold your breath."},
	},
	"W03": {
		Desc:    "A gentle rotation drill for the trunk and upper spine.",
		HowTo:   []string{"Stand with feet hip-width apart.", "Keep the hips mostly facing forward.", "Rotate the upper body to one side.", "Return to center and rotate to the other side."},
		Correct: []string{"Keep your feet stable.", "Turn from your torso.", "Rotate smoothly.", "Keep your knees still as you turn."},
		Wrong:   []string{"Don't turn your knees with your body.", "Don't rotate too fast.", "Don't lean backward.", "Don't force the range of motion."},
	},
	"W04": {
		Desc:    "A movement drill for the hips, hamstrings and lower back control.",
		HowTo:   []string{"Stand tall with feet hip-width apart.", "Soften the knees slightly.", "Push the hips back.", "Lower the torso with a straight back.", "Return to standing by driving the hips forward."},
		Correct: []string{"Keep your back straight.", "Move from your hips.", "Keep your knees soft, not deeply bent.", "Let your neck follow your spine."},
		Wrong:   []string{"Don't round your back.", "Don't turn it into a squat.", "Don't reach too low.", "Don't lift your head too high."},
	},
	"W05": {
		Desc:    "A controlled leg warm-up using a safe, shallow range.",
		HowTo:   []string{"Stand with feet slightly wider than hip-width.", "Push the hips slightly back.", "Bend the knees only to a comfortable depth.", "Return to standing slowly."},
		Correct: []string{"Keep your knees over your toes.", "Keep your heels on the floor.", "Keep your back long.", "Keep the depth moderate."},
		Wrong:   []string{"Don't squat too deep.", "Don't let your knees collapse inward.", "Don't lift your heels.", "Don't drop down too fast."},
	},
	"W06": {
		Desc:    "A warm-up for the shoulders, wrists and core.",
		HowTo:   []string{"Start in a high plank position.", "Hands are under the shoulders.", "Shift the body slightly forward.", "Shift back to the starting position.", "Keep the movement small and controlled."},
		Correct: []string{"Keep your body in one line.", "Keep your shoulders stable.", "Keep your core engaged.", "Move smoothly."},
		Wrong:   []string{"Don't drop your hips.", "Don't push your shoulders up to your ears.", "Don't move too far forward.", "Don't let your lower back sag."},
	},

	// ---- Abs / core ----
	"C01": {
		Desc:    "A core stability exercise for the abs, shoulders, glutes and trunk.",
		HowTo:   []string{"Place the elbows under the shoulders.", "Step the feet back.", "Make a straight line from shoulders to heels.", "Tighten the abs and glutes.", "Hold the position."},
		Correct: []string{"Keep your body in one line.", "Keep your elbows under your shoulders.", "Keep your core active.", "Keep your breathing steady."},
		Wrong:   []string{"Don't let your lower back drop.", "Don't lift your hips too high.", "Don't look forward.", "Don't hold your breath."},
	},
	"C02": {
		Desc:    "A short, intense plank focused on full-body tension.",
		HowTo:   []string{"Start in a forearm plank.", "Tighten the abs strongly.", "Squeeze the glutes.", "Imagine pulling the elbows toward the toes.", "Hold with maximum control."},
		Correct: []string{"Tense your whole body.", "Keep your glutes active.", "Keep your core tight.", "Keep your breathing controlled."},
		Wrong:   []string{"Don't hold it like a relaxed plank.", "Don't let your hips sag.", "Don't hold your breath.", "Don't try to hold too long."},
	},
	"C03": {
		Desc:    "A lower-ab exercise with controlled crossing movements of the legs.",
		HowTo:   []string{"Lie on the back.", "Press the lower back gently toward the floor.", "Lift the legs.", "Cross one leg over the other.", "Continue alternating the crossing motion."},
		Correct: []string{"Keep your lower back close to the floor.", "Keep your legs controlled.", "Relax your neck.", "Move smoothly."},
		Wrong:   []string{"Don't arch your lower back.", "Don't lower your legs too much.", "Don't pull your head forward.", "Don't move too fast."},
	},
	"C04": {
		Desc:    "A lower-ab exercise with short alternating leg kicks.",
		HowTo:   []string{"Lie on the back.", "Press the lower back toward the floor.", "Lift the legs.", "Move the legs up and down in short alternating kicks."},
		Correct: []string{"Keep the range of motion small.", "Keep your lower back stable.", "Don't let your legs drop too low.", "Keep your breathing steady."},
		Wrong:   []string{"Don't arch your lower back.", "Don't make large kicks.", "Don't tense your neck.", "Don't hold your breath."},
	},
	"C05": {
		Desc:    "An abdominal exercise focused on lifting the pelvis with control.",
		HowTo:   []string{"Lie on the back.", "Bend the knees and lift the feet.", "Use the abs to curl the pelvis slightly upward.", "Lower back down slowly."},
		Correct: []string{"Move from your lower abs.", "Lift your pelvis with control.", "Relax your neck.", "Lower slowly."},
		Wrong:   []string{"Don't swing your legs.", "Don't use momentum.", "Don't pull on your neck.", "Don't drop your hips quickly."},
	},
	"C06": {
		Desc:    "A static core hold with a shortened lever for better control.",
		HowTo:   []string{"Lie on the back.", "Press the lower back toward the floor.", "Lift the shoulders slightly.", "Keep the knees bent or legs raised higher.", "Hold the position."},
		Correct: []string{"Keep your lower back on the floor.", "Keep your abs tight.", "Relax your neck.", "Keep breathing."},
		Wrong:   []string{"Don't arch your lower back.", "Don't lower your legs too much.", "Don't pull your chin forward.", "Don't hold your breath."},
	},
	"C07": {
		Desc:    "A lateral core exercise for the obliques, shoulder and hip stability.",
		HowTo:   []string{"Lie on one side.", "Place the elbow under the shoulder.", "Lift the hips.", "Keep the body in one straight line.", "Hold the position."},
		Correct: []string{"Keep your elbow under your shoulder.", "Keep your hips lifted.", "Keep your body long.", "Keep your neck neutral."},
		Wrong:   []string{"Don't drop your hips.", "Don't shrug your shoulder.", "Don't roll your chest forward.", "Don't hold your breath."},
	},
	"C08": {
		Desc:    "A stronger side-plank variation for the obliques and lateral hip.",
		HowTo:   []string{"Start in a side plank.", "Lower the hips slightly.", "Lift the hips back up.", "Repeat with control.", "Switch sides."},
		Correct: []string{"Keep the movement small and controlled.", "Keep your elbow under your shoulder.", "Move your hips straight up and down.", "Keep your core tight."},
		Wrong:   []string{"Don't drop too low.", "Don't twist your torso.", "Don't collapse into your shoulder.", "Don't move too fast."},
	},
	"C09": {
		Desc:    "A core and shoulder exercise: tapping each shoulder while holding a plank.",
		HowTo:   []string{"Start in a high plank.", "Keep the hips level and stable.", "Tap one hand to the opposite shoulder.", "Return it and tap with the other hand.", "Keep the body still."},
		Correct: []string{"Keep your hips level.", "Keep your core tight.", "Keep your shoulders stable.", "Move slowly."},
		Wrong:   []string{"Don't rock your hips side to side.", "Don't move too fast.", "Don't let your lower back drop.", "Don't twist your torso."},
	},
	"C10": {
		Desc:    "A core and triceps exercise moving between forearm plank and high plank.",
		HowTo:   []string{"Start in a forearm plank.", "Place one hand on the floor and press up.", "Bring the other hand up into a high plank.", "Lower back down to the forearms.", "Alternate the leading arm."},
		Correct: []string{"Keep your hips stable.", "Keep your core engaged.", "Move with control.", "Keep your shoulders over your hands or elbows."},
		Wrong:   []string{"Don't rock your hips side to side.", "Don't let your lower back drop.", "Don't move too fast.", "Don't let your shoulders collapse."},
	},
	"C11": {
		Desc:    "A controlled core exercise and a safer alternative to harder leg-lift movements.",
		HowTo:   []string{"Lie on the back.", "Lift the legs into a tabletop position.", "Lower one heel toward the floor.", "Tap lightly and return.", "Alternate sides."},
		Correct: []string{"Keep your lower back stable.", "Move slowly.", "Touch your heel down softly.", "Keep your pelvis steady."},
		Wrong:   []string{"Don't drop your leg too fast.", "Don't arch your lower back.", "Don't move both legs at once.", "Don't rush the movement."},
	},
	"C12": {
		Desc:    "A core control exercise for the deep abs and spinal stability.",
		HowTo:   []string{"Lie on the back.", "Lift the arms and legs into tabletop position.", "Extend the opposite arm and leg.", "Keep the lower back stable.", "Return and switch sides."},
		Correct: []string{"Keep your lower back close to the floor.", "Move slowly.", "Move your opposite arm and leg together.", "Keep your torso still."},
		Wrong:   []string{"Don't arch your lower back.", "Don't move too fast.", "Don't drop your leg too low.", "Don't let your ribs flare."},
	},

	// ---- Push-ups / triceps / shoulders ----
	"P01": {
		Desc:    "A basic upper-body exercise for the chest, triceps, shoulders and core.",
		HowTo:   []string{"Start in a high plank.", "Place the hands under the shoulders or slightly wider.", "Lower the body with control.", "Push back up while keeping the body straight."},
		Correct: []string{"Keep your body in one line.", "Keep your core tight.", "Keep your elbows at about 30 to 45 degrees.", "Move with control."},
		Wrong:   []string{"Don't drop your hips.", "Don't raise your hips too high.", "Don't flare your elbows straight out.", "Don't fall into the bottom position."},
	},
	"P02": {
		Desc:    "A push-up variation with more focus on the triceps.",
		HowTo:   []string{"Start in a high plank.", "Place the hands closer than in a regular push-up.", "Keep the elbows near the body.", "Lower slowly and press back up."},
		Correct: []string{"Keep your elbows close.", "Keep your body straight.", "Keep your shoulders away from your ears.", "Keep your wrists comfortable."},
		Wrong:   []string{"Don't place your hands too close.", "Don't flare your elbows.", "Don't let your lower back drop.", "Don't push through wrist or elbow pain."},
	},
	"P03": {
		Desc:    "A controlled push-up with a short pause in the bottom position.",
		HowTo:   []string{"Start in a regular push-up position.", "Lower slowly.", "Pause for one second near the bottom.", "Push back up without losing body position."},
		Correct: []string{"Keep the pause controlled.", "Keep your core tight.", "Keep your chest active.", "Keep your body straight."},
		Wrong:   []string{"Don't relax at the bottom.", "Don't drop your hips.", "Don't hold your breath.", "Don't bounce out of the bottom position."},
	},
	"P04": {
		Desc:    "A slow push-up variation that increases control and time under tension.",
		HowTo:   []string{"Start in a high plank.", "Lower for about three seconds.", "Pause for one second near the bottom.", "Push up for about three seconds."},
		Correct: []string{"Keep a slow, even tempo.", "Keep your body straight.", "Keep your elbows controlled.", "Keep your breathing steady."},
		Wrong:   []string{"Don't rush the movement.", "Don't collapse at the bottom.", "Don't let your hips sag.", "Don't lose tension in your core."},
	},
	"P05": {
		Desc:    "An asymmetric push-up variation that challenges the chest, shoulders and core.",
		HowTo:   []string{"Start in a push-up position.", "Place one hand slightly forward and the other slightly back.", "Lower with control.", "Push back up.", "Switch hand position on the next side."},
		Correct: []string{"Keep your body square.", "Keep your core tight.", "Keep both shoulders stable.", "Move with control."},
		Wrong:   []string{"Don't twist your torso.", "Don't drop one shoulder.", "Don't go too wide with your hands.", "Don't force the range if your shoulder feels uncomfortable."},
	},
	"P06": {
		Desc:    "A shoulder-focused push-up variation.",
		HowTo:   []string{"Start in a high plank.", "Walk the feet slightly closer and lift the hips.", "Keep the head between the arms.", "Lower the head toward the floor.", "Press back up."},
		Correct: []string{"Keep your hips high.", "Bend your elbows with control.", "Keep your shoulders stable.", "Keep your neck neutral."},
		Wrong:   []string{"Don't drop into a regular push-up.", "Don't let your head lead too far forward.", "Don't flare your elbows too much.", "Don't push through shoulder pain."},
	},
	"P08": {
		Desc:    "A triceps and shoulder exercise performed from a forearm support position.",
		HowTo:   []string{"Start on the forearms with the body long.", "Keep the elbows under or slightly ahead of the shoulders.", "Press through the hands and forearms.", "Extend the elbows slightly.", "Return with control."},
		Correct: []string{"Keep your core tight.", "Move your elbows smoothly.", "Keep your shoulders down.", "Keep the range of motion comfortable."},
		Wrong:   []string{"Don't force your elbows.", "Don't drop your hips.", "Don't shrug your shoulders.", "Don't move through elbow or shoulder pain."},
	},

	// ---- Back / posture ----
	"B01": {
		Desc:    "A posture-focused back exercise for the upper back, shoulder blades and spinal extensors.",
		HowTo:   []string{"Lie face down.", "Place the arms along the body.", "Turn the thumbs slightly outward.", "Lift the chest slightly.", "Pull the shoulder blades back and down.", "Hold."},
		Correct: []string{"Keep the lift small.", "Keep your neck neutral.", "Draw your shoulder blades back and down.", "Lightly engage your glutes."},
		Wrong:   []string{"Don't lift too high.", "Don't throw your head back.", "Don't over-arch your lower back.", "Don't shrug your shoulders."},
	},
	"B02": {
		Desc:    "An upper-back and shoulder-blade exercise performed lying face down.",
		HowTo:   []string{"Lie on the stomach.", "Keep the head facing down.", "Move the arms in a wide arc from the sides toward overhead.", "Return along the same path.", "Keep the movement controlled."},
		Correct: []string{"Relax your neck.", "Move your shoulder blades smoothly.", "Keep your chest lift small or neutral.", "Move your arms without rushing."},
		Wrong:   []string{"Don't lift your head.", "Don't arch your lower back too much.", "Don't swing your arms.", "Don't shrug your shoulders."},
	},
	"B03": {
		Desc:    "An upper-back exercise for the shoulder blades and rear shoulders.",
		HowTo:   []string{"Lie face down.", "Bend the elbows into a W shape.", "Lift the elbows and hands slightly.", "Squeeze the shoulder blades.", "Lower with control."},
		Correct: []string{"Keep the range of motion small.", "Let your shoulder blades do the work.", "Keep your neck neutral.", "Keep your shoulders away from your ears."},
		Wrong:   []string{"Don't lift too high.", "Don't pull your head up.", "Don't use momentum.", "Don't over-arch your lower back."},
	},
	"B04": {
		Desc:    "A shoulder-blade exercise for the lower traps and upper back.",
		HowTo:   []string{"Lie face down.", "Reach the arms diagonally overhead into a Y shape.", "Lift the arms slightly from the floor.", "Lower with control."},
		Correct: []string{"Lift your arms only a little.", "Keep your neck neutral.", "Keep your shoulders down.", "Move with control."},
		Wrong:   []string{"Don't lift too high.", "Don't shrug your shoulders.", "Don't arch your lower back.", "Don't move too fast."},
	},
	"B05": {
		Desc:    "An upper-back exercise for the rear shoulders and shoulder-blade control.",
		HowTo:   []string{"Lie face down.", "Extend the arms out to the sides in a T shape.", "Lift the arms slightly.", "Squeeze the shoulder blades.", "Lower slowly."},
		Correct: []string{"Keep your arms in line with your shoulders.", "Relax your neck.", "Move your shoulder blades gently.", "Keep the lift small."},
		Wrong:   []string{"Don't throw your arms upward.", "Don't lift your head.", "Don't shrug.", "Don't force a large range of motion."},
	},
	"B06": {
		Desc:    "A back exercise that mimics a pulling motion without equipment.",
		HowTo:   []string{"Lie on the stomach with arms overhead.", "Lift the chest slightly.", "Pull the elbows down toward the ribs.", "Squeeze the shoulder blades.", "Reach the arms forward again."},
		Correct: []string{"Pull your elbows down with control.", "Squeeze your shoulder blades gently.", "Keep your neck neutral.", "Keep your lower back comfortable."},
		Wrong:   []string{"Don't jerk your arms.", "Don't lift too high.", "Don't throw your head back.", "Don't let your shoulders rise to your ears."},
	},
	"B07": {
		Desc:    "A back and coordination exercise performed lying face down.",
		HowTo:   []string{"Lie on the stomach.", "Lift the chest slightly.", "Move opposite arm and leg in a controlled swimming pattern.", "Keep the movement small and steady."},
		Correct: []string{"Keep your neck neutral.", "Move with control.", "Keep your lower back comfortable.", "Keep breathing."},
		Wrong:   []string{"Don't move too fast.", "Don't lift too high.", "Don't throw your head back.", "Don't force your lower back."},
	},
	"B08": {
		Desc:    "A posterior-chain exercise for the glutes, hamstrings, shoulders and back.",
		HowTo:   []string{"Sit on the floor with legs extended.", "Place the hands behind the hips.", "Press through the hands and heels.", "Lift the hips.", "Hold the body in a long line."},
		Correct: []string{"Keep your chest open.", "Keep your hips lifted.", "Keep your shoulders stable.", "Keep your neck neutral."},
		Wrong:   []string{"Don't drop your hips.", "Don't shrug your shoulders.", "Don't throw your head back.", "Don't force your wrists."},
	},
	"B09": {
		Desc:    "A small-range back exercise for the spinal extensors and posture.",
		HowTo:   []string{"Lie face down.", "Place the hands beside the body or lightly near the chest.", "Lift the chest slightly.", "Make small controlled pulses.", "Keep the neck neutral."},
		Correct: []string{"Keep the lift small.", "Keep your pulses smooth.", "Keep your neck long.", "Keep your lower back controlled."},
		Wrong:   []string{"Don't pulse too high.", "Don't use momentum.", "Don't look forward.", "Don't push into sharp lower-back pain."},
	},
	"B10": {
		Desc:    "A stronger back exercise combining a small cobra lift with a W-shaped arm pull.",
		HowTo:   []string{"Lie face down with arms forward.", "Lift the chest slightly.", "Pull the elbows down and back into a W shape.", "Squeeze the shoulder blades.", "Reach forward again and lower with control."},
		Correct: []string{"Keep your chest lift small.", "Keep the W-pull smooth.", "Draw your shoulder blades back and down.", "Keep your neck long."},
		Wrong:   []string{"Don't over-arch your lower back.", "Don't pull with your neck.", "Don't move too fast.", "Don't shrug your shoulders."},
	},

	// ---- Arms / biceps ----
	"A01": {
		Desc:    "A no-equipment biceps exercise using one hand to resist the other.",
		HowTo:   []string{"Stand or sit tall.", "Bend one elbow like a curl.", "Use the opposite hand to press down on the working forearm.", "Curl upward slowly against resistance.", "Switch sides."},
		Correct: []string{"Keep the resistance steady.", "Move slowly.", "Keep your elbow near your body.", "Keep your shoulder down."},
		Wrong:   []string{"Don't move without real resistance.", "Don't jerk your arm.", "Don't raise your shoulder.", "Don't hold your breath."},
	},
	"A02": {
		Desc:    "A static biceps hold using self-resistance.",
		HowTo:   []string{"Bend one elbow to about 90 degrees.", "Use the opposite hand to press down.", "The working arm resists without moving.", "Hold the tension.", "Switch sides."},
		Correct: []string{"Keep your elbow angle stable.", "Keep the tension steady.", "Relax your shoulders.", "Keep breathing."},
		Wrong:   []string{"Don't press in short jerks.", "Don't twist your torso.", "Don't raise your shoulder.", "Don't hold your breath."},
	},
	"A03": {
		Desc:    "An isometric arm and forearm exercise using the hands against each other.",
		HowTo:   []string{"Hook the fingers or hands together.", "Keep the elbows slightly bent.", "Pull the hands away from each other.", "Hold steady tension.", "Relax slowly."},
		Correct: []string{"Keep the tension controlled.", "Keep your shoulders down.", "Keep your wrists comfortable.", "Keep your breathing calm."},
		Wrong:   []string{"Don't pull with a jerk.", "Don't shrug your shoulders.", "Don't over-bend your wrists.", "Don't hold your breath."},
	},
	"A04": {
		Desc:    "A slow lowering exercise for the biceps using self-resistance.",
		HowTo:   []string{"Start with one elbow bent.", "Use the opposite hand to create resistance.", "Slowly lower the working forearm.", "Keep resisting during the lowering phase.", "Switch sides."},
		Correct: []string{"Lower slowly.", "Keep the resistance constant.", "Keep your elbow close to your body.", "Relax your shoulder."},
		Wrong:   []string{"Don't let your arm drop quickly.", "Don't use no resistance.", "Don't twist your body.", "Don't hold your breath."},
	},

	// ---- Legs / glutes ----
	"L01": {
		Desc:    "A controlled leg exercise for the thighs, hips and posture.",
		HowTo:   []string{"Stand with feet slightly wider than hip-width.", "Push the hips back.", "Bend the knees to a comfortable shallow depth.", "Stand back up slowly."},
		Correct: []string{"Keep your knees over your toes.", "Keep your heels down.", "Keep your back long.", "Keep the depth comfortable."},
		Wrong:   []string{"Don't squat too deep.", "Don't let your knees collapse inward.", "Don't lift your heels.", "Don't drop quickly."},
	},
	"L02": {
		Desc:    "A controlled squat variation with a short hold in a safe range.",
		HowTo:   []string{"Stand with feet stable.", "Lower into a comfortable shallow squat.", "Pause briefly.", "Stand back up with control."},
		Correct: []string{"Keep the pause stable.", "Keep your knees over your toes.", "Keep your chest open.", "Keep your heels on the floor."},
		Wrong:   []string{"Don't go too deep.", "Don't bounce in the bottom position.", "Don't let your knees cave in.", "Don't hold your breath."},
	},
	"L03": {
		Desc:    "A movement drill for the hips, hamstrings and lower back control.",
		HowTo:   []string{"Stand tall with feet hip-width apart.", "Soften the knees slightly.", "Push the hips back.", "Lower the torso with a straight back.", "Return to standing by driving the hips forward."},
		Correct: []string{"Keep your back straight.", "Move from your hips.", "Keep your knees soft, not deeply bent.", "Let your neck follow your spine."},
		Wrong:   []string{"Don't round your back.", "Don't turn it into a squat.", "Don't reach too low.", "Don't lift your head too high."},
	},
	"L04": {
		Desc:    "A stronger glute exercise using one leg at a time.",
		HowTo:   []string{"Lie on the back.", "Bend one knee and keep that foot on the floor.", "Extend or lift the other leg.", "Push through the working heel.", "Lift the hips and lower slowly."},
		Correct: []string{"Lift with your working glute.", "Keep your hips level.", "Move with control.", "Keep your lower back neutral."},
		Wrong:   []string{"Don't twist your pelvis.", "Don't push through your toes.", "Don't arch your lower back.", "Don't drop too fast."},
	},
	"L05": {
		Desc:    "A glute and pelvis-stability exercise.",
		HowTo:   []string{"Start in a glute bridge position.", "Keep the hips lifted.", "Lift one foot slightly from the floor.", "Put it down and switch sides.", "Keep the pelvis stable."},
		Correct: []string{"Keep your hips level.", "Keep your glutes active.", "Move slowly.", "Keep your lower back comfortable."},
		Wrong:   []string{"Don't let your hips drop.", "Don't rock side to side.", "Don't move too fast.", "Don't arch your lower back."},
	},
	"L06": {
		Desc:    "A glute exercise that strengthens the hips and posterior chain.",
		HowTo:   []string{"Lie on the back.", "Bend the knees and place the feet on the floor.", "Lift the hips.", "Pause at the top.", "Lower slowly."},
		Correct: []string{"Push through your heels.", "Squeeze your glutes at the top.", "Keep your ribs down.", "Don't let your lower back over-arch."},
		Wrong:   []string{"Don't arch your lower back at the top.", "Don't lift your heels.", "Don't drop your hips too fast.", "Don't place your feet too close."},
	},
	"L07": {
		Desc:    "A lower-leg exercise for the calves, ankles and foot control.",
		HowTo:   []string{"Stand tall with feet hip-width apart.", "Rise onto the balls of the feet.", "Pause briefly at the top.", "Lower the heels slowly."},
		Correct: []string{"Stand up tall.", "Move smoothly.", "Keep your ankles aligned.", "Lower with control."},
		Wrong:   []string{"Don't drop your heels quickly.", "Don't rock forward.", "Don't roll your ankles outward.", "Don't use momentum."},
	},
	"L09": {
		Desc:    "A lateral-chain exercise for the side glutes, obliques and hip stability.",
		HowTo:   []string{"Start in a side plank.", "Keep the body in a straight line.", "Lift the top leg slightly.", "Lower it with control.", "Repeat and switch sides."},
		Correct: []string{"Keep your hips lifted.", "Move your top leg slowly.", "Keep your torso stable.", "Keep your shoulder strong."},
		Wrong:   []string{"Don't drop your hips.", "Don't swing your leg.", "Don't roll your body backward.", "Don't collapse into your shoulder."},
	},

	// ---- Cool-down ----
	"CD01": {
		Desc:    "A gentle stretch for the chest and front of the shoulders.",
		HowTo:   []string{"Lie on the back.", "Open the arms out to the sides.", "Relax the shoulders.", "Let the chest open gently.", "Breathe slowly."},
		Correct: []string{"Let your chest open softly.", "Keep your shoulders relaxed.", "Keep your lower back comfortable.", "Keep your shoulders pain-free."},
		Wrong:   []string{"Don't force your arms down.", "Don't arch your lower back.", "Don't tense your neck.", "Don't stretch through pain."},
	},
	"CD02": {
		Desc:    "A gentle stretch for the back of the thigh.",
		HowTo:   []string{"Lie on the back.", "Lift one leg.", "Hold behind the thigh or calf.", "Pull the leg gently toward you.", "Switch sides."},
		Correct: []string{"Keep the stretch mild.", "Let your knee stay slightly bent.", "Keep your lower back calm.", "Keep your breathing slow."},
		Wrong:   []string{"Don't pull too hard.", "Don't force your leg straight.", "Don't lift your hips.", "Don't push into sharp pain behind your knee."},
	},
	"CD03": {
		Desc:    "A gentle rotation stretch for the spine and trunk.",
		HowTo:   []string{"Lie on the back.", "Bend the knees.", "Let the knees move gently to one side.", "Keep the shoulders on the floor.", "Return to center and switch sides."},
		Correct: []string{"Move slowly.", "Keep your shoulders down.", "Keep your breathing relaxed.", "Keep the stretch comfortable."},
		Wrong:   []string{"Don't force your knees to the floor.", "Don't twist quickly.", "Don't lift your shoulder.", "Don't stretch through pain."},
	},
	"CD04": {
		Desc:    "A gentle front-body stretch and mild back extension.",
		HowTo:   []string{"Lie on the stomach.", "Place the forearms on the floor.", "Lift the chest gently.", "Keep the neck long.", "Breathe slowly."},
		Correct: []string{"Keep the lift mild.", "Keep your shoulders away from your ears.", "Keep your lower back comfortable.", "Keep your neck neutral."},
		Wrong:   []string{"Don't over-arch your lower back.", "Don't throw your head back.", "Don't shrug your shoulders.", "Don't force the position."},
	},
	"CD05": {
		Desc:    "A calm breathing exercise to finish the session.",
		HowTo:   []string{"Lie on the back.", "Relax the shoulders and jaw.", "Place one hand on the belly if comfortable.", "Inhale gently.", "Exhale slowly."},
		Correct: []string{"Keep your breathing calm.", "Keep your shoulders relaxed.", "Keep your face soft.", "Let your body settle down."},
		Wrong:   []string{"Don't breathe too forcefully.", "Don't hold your breath.", "Don't tense your neck.", "Don't arch your lower back."},
	},

	// ---- Vlad set additions (warm-up / cardio / plyometrics / lunges) ----
	"W07": {
		Desc:    "A calm opening drill: deep breathing with a gentle full-body reach to lengthen the spine.",
		HowTo:   []string{"Stand tall with feet hip-width apart.", "Inhale and reach both arms overhead.", "Lengthen gently through the spine.", "Exhale and lower the arms, dropping the shoulders."},
		Correct: []string{"Breathe slowly and fully.", "Relax your shoulders as you exhale.", "Keep your ribs down, don't over-arch.", "Move unhurried."},
		Wrong:   []string{"Don't hold your breath.", "Don't arch your lower back hard.", "Don't shrug your shoulders to your ears.", "Don't rush the reach."},
	},
	"W08": {
		Desc:    "A gentle rotation drill to warm up the spine and trunk.",
		HowTo:   []string{"Stand with feet hip-width apart.", "Let the arms swing loosely.", "Rotate the upper body to one side.", "Flow smoothly to the other side."},
		Correct: []string{"Turn from your trunk.", "Keep your hips mostly forward.", "Keep your feet planted.", "Keep a smooth, even pace."},
		Wrong:   []string{"Don't twist hard through your knees.", "Don't yank your arms.", "Don't rotate too fast.", "Don't hold your breath."},
	},
	"W09": {
		Desc:    "A dynamic shoulder and chest warm-up to prepare the upper body for push-ups and jumping exercises.",
		HowTo:   []string{"Stand tall, feet hip-width apart.", "Swing the arms forward and backward in a controlled way.", "Let the chest open naturally as the arms move back.", "Keep the movement smooth and relaxed."},
		Correct: []string{"Let your arms move freely.", "Keep your shoulders relaxed.", "Let your chest open gently.", "Stand up tall."},
		Wrong:   []string{"Don't swing too aggressively.", "Don't raise your shoulders to your ears.", "Don't arch your lower back.", "Don't let the movement turn into uncontrolled momentum."},
	},
	"C13": {
		Desc:    "A rotational abdominal exercise for the obliques and front abs.",
		HowTo:   []string{"Lie on the back.", "Bring the knees up.", "Place the hands lightly near the head.", "Rotate the torso and bring one elbow toward the opposite knee.", "Switch sides in a controlled rhythm."},
		Correct: []string{"Rotate from your torso.", "Keep your lower back controlled.", "Relax your neck.", "Move smoothly."},
		Wrong:   []string{"Don't pull on your neck.", "Don't move only your elbows.", "Don't rush the reps.", "Don't let your lower back arch."},
	},
	"B11": {
		Desc:    "A back exercise lifting opposite arm and leg in an alternating, swimming-like motion.",
		HowTo:   []string{"Lie face down, arms reaching overhead.", "Lift the chest and legs slightly.", "Raise one arm and the opposite leg.", "Lower and switch sides, alternating smoothly."},
		Correct: []string{"Keep the lift small.", "Let your neck follow your spine.", "Keep your glutes engaged.", "Move steadily."},
		Wrong:   []string{"Don't crank your lower back.", "Don't lift your head.", "Don't move in jerks.", "Don't hold your breath."},
	},
	"L10": {
		Desc:    "A basic bodyweight squat for the legs, hips and glutes.",
		HowTo:   []string{"Stand with feet about shoulder-width apart.", "Push the hips back slightly.", "Bend the knees and lower into a squat.", "Keep the chest open.", "Stand back up through the whole foot."},
		Correct: []string{"Keep your knees over your toes.", "Keep your heels on the floor.", "Keep your back long.", "Keep the depth controlled."},
		Wrong:   []string{"Don't let your knees collapse inward.", "Don't lift your heels.", "Don't round your back.", "Don't drop too fast."},
	},
	"L11": {
		Desc:    "A lower-body exercise for the legs, glutes and balance.",
		HowTo:   []string{"Stand tall.", "Step one leg back.", "Lower into a lunge with control.", "Push through the front foot to return to standing.", "Switch sides."},
		Correct: []string{"Keep your front knee over your toes.", "Keep your torso upright.", "Step back with control.", "Keep your front heel grounded."},
		Wrong:   []string{"Don't let your front knee collapse inward.", "Don't lean too far forward.", "Don't drop quickly.", "Don't push off your back leg too hard."},
	},
	"L12": {
		Desc:    "An air squat done slowly to build control and leg strength.",
		HowTo:   []string{"Stand with feet shoulder-width apart.", "Lower over about three seconds.", "Pause briefly at the bottom.", "Stand up with control."},
		Correct: []string{"Keep the tempo slow.", "Keep your knees over your toes.", "Keep your heels grounded.", "Keep your back neutral."},
		Wrong:   []string{"Don't drop down fast.", "Don't bounce out of the bottom.", "Don't let your knees cave in.", "Don't hold your breath."},
	},
	"J01": {
		Desc:    "A light ankle and calf warm-up using small, quick jumps.",
		HowTo:   []string{"Stand tall with feet about hip-width apart.", "Keep the knees slightly soft.", "Make small quick jumps using mostly the ankles.", "Land quietly and repeat."},
		Correct: []string{"Keep your jumps low.", "Land soft and quiet.", "Keep your knees slightly bent.", "Stand up tall."},
		Wrong:   []string{"Don't jump too high.", "Don't land loudly.", "Don't lock your knees.", "Don't let your ankles roll inward or outward."},
	},
	"J02": {
		Desc:    "A low-impact warm-up alternative to high knees.",
		HowTo:   []string{"Stand tall.", "Lift one knee toward hip height.", "Lower it with control.", "Switch sides and continue marching.", "Keep the arms moving naturally."},
		Correct: []string{"Keep your torso upright.", "Lift your knees under control.", "Land your feet softly.", "Keep your breathing steady."},
		Wrong:   []string{"Don't lean backward.", "Don't swing your legs without control.", "Don't slam your feet down.", "Don't hold your breath."},
	},
	"J03": {
		Desc:    "A full-body cardio exercise that raises the heart rate and warms up the shoulders, hips and legs.",
		HowTo:   []string{"Stand tall with arms by the sides.", "Jump the feet out while raising the arms overhead.", "Jump the feet back in while lowering the arms.", "Repeat with a steady rhythm."},
		Correct: []string{"Land softly.", "Keep your knees slightly bent.", "Move your arms smoothly.", "Keep your breathing rhythmic."},
		Wrong:   []string{"Don't land hard.", "Don't lock your knees.", "Don't shrug your shoulders.", "Don't move too fast and lose control."},
	},
	"J04": {
		Desc:    "A dynamic core and cardio exercise performed from a high plank.",
		HowTo:   []string{"Start in a high plank.", "Keep the hands under the shoulders.", "Drive one knee toward the chest.", "Switch legs rhythmically.", "Keep the hips stable."},
		Correct: []string{"Hold a strong plank.", "Keep your core tight.", "Keep your shoulders over your hands.", "Move your knees under control."},
		Wrong:   []string{"Don't let your hips bounce high.", "Don't let your lower back drop.", "Don't move your feet too wildly.", "Don't let your shoulders collapse."},
	},
	"J05": {
		Desc:    "A plyometric leg exercise for power, conditioning and coordination.",
		HowTo:   []string{"Start with feet about shoulder-width apart.", "Lower into a controlled squat.", "Jump upward.", "Land softly with knees slightly bent.", "Reset your position before the next jump."},
		Correct: []string{"Land quietly.", "Keep your knees over your toes.", "Keep your chest open.", "Keep each jump controlled."},
		Wrong:   []string{"Don't land hard.", "Don't let your knees collapse inward.", "Don't jump before the squat is stable.", "Don't rush without resetting."},
	},
	"J06": {
		Desc:    "A lateral jumping exercise for legs, hips, balance and coordination.",
		HowTo:   []string{"Stand on one leg.", "Hop sideways to the other leg.", "Land softly and stabilize.", "Repeat side to side.", "Use the arms for balance."},
		Correct: []string{"Land quietly.", "Keep your knee over your toes.", "Keep your torso controlled.", "Move smoothly side to side."},
		Wrong:   []string{"Don't let your knee collapse inward as you land.", "Don't jump too far too soon.", "Don't fall into the landing.", "Don't twist your knee on contact."},
	},
	"J07": {
		Desc:    "A cardio exercise that raises the heart rate and activates the hips, core and legs.",
		HowTo:   []string{"Stand tall.", "Run in place while lifting the knees high.", "Move the arms naturally.", "Keep the rhythm quick but controlled."},
		Correct: []string{"Lift your knees toward hip height.", "Keep your landings light.", "Keep your core engaged.", "Keep your torso tall."},
		Wrong:   []string{"Don't lean backward.", "Don't land heavily.", "Don't let your feet slap the floor.", "Don't lose your posture as the pace picks up."},
	},
	"J08": {
		Desc:    "A full-body conditioning exercise combining a squat, plank and jump. In this program, use the no-push-up version unless a push-up is specifically assigned.",
		HowTo:   []string{"Stand tall.", "Squat down and place the hands on the floor.", "Jump or step the feet back into a plank.", "Jump or step the feet forward.", "Stand up and finish with a small jump."},
		Correct: []string{"Land your hands under your shoulders.", "Hold a strong plank.", "Land your feet softly.", "Keep the jump controlled."},
		Wrong:   []string{"Don't let your lower back drop in the plank.", "Don't land heavily.", "Don't rush through sloppy reps.", "Don't let your knees collapse inward when jumping forward."},
	},
	"J09": {
		Desc:    "An advanced plyometric leg exercise for power, conditioning and coordination.",
		HowTo:   []string{"Start in a lunge position.", "Jump upward.", "Switch legs in the air.", "Land softly in the opposite lunge.", "Reset your balance before the next rep."},
		Correct: []string{"Land soft and controlled.", "Keep your front knee over your toes.", "Keep your torso upright.", "Keep the range of motion safe."},
		Wrong:   []string{"Don't land hard.", "Don't let your front knee cave inward.", "Don't drop too deep.", "Don't move faster than you can control."},
	},
	"CD07": {
		Desc:    "A restful stretch for the back, hips and shoulders to wind down.",
		HowTo:   []string{"Kneel and sit the hips back toward the heels.", "Reach the arms forward on the floor.", "Let the forehead rest down.", "Breathe slowly and relax."},
		Correct: []string{"Let your hips settle toward your heels.", "Lengthen your back gently.", "Relax your shoulders.", "Keep your breathing slow."},
		Wrong:   []string{"Don't force your hips down.", "Don't tense your shoulders.", "Don't hold your breath.", "Don't strain your knees."},
	},
}
