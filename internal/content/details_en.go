package content

// detailsEN is the English rich content for every exercise, keyed by the current
// library IDs. It is the authoritative card text; other languages fall back here
// until translated.
var detailsEN = map[string]Detail{
	// ---- Warm-up ----
	"W01": {
		Desc:    "A simple shoulder warm-up to prepare the neck, shoulders and upper back.",
		HowTo:   []string{"Stand tall with the arms relaxed.", "Make slow circles with the shoulders backward.", "Then make slow circles forward.", "Keep the movement smooth and controlled."},
		Correct: []string{"Neck relaxed.", "Shoulders move smoothly.", "Arms stay loose.", "Body stays upright."},
		Wrong:   []string{"Shrugging hard.", "Moving too fast.", "Tensing the neck.", "Rounding the upper back."},
	},
	"W02": {
		Desc:    "A shoulder and upper-back warm-up that activates the deltoids and shoulder stabilizers.",
		HowTo:   []string{"Stand tall and raise the arms to shoulder height.", "Make small circles forward.", "Then make small circles backward.", "Keep the circles tight and controlled."},
		Correct: []string{"Arms stay near shoulder height.", "Circles are small.", "Shoulders stay down.", "Neck stays relaxed."},
		Wrong:   []string{"Making large swinging motions.", "Raising the shoulders to the ears.", "Arching the lower back.", "Holding the breath."},
	},
	"W03": {
		Desc:    "A gentle rotation drill for the trunk and upper spine.",
		HowTo:   []string{"Stand with feet hip-width apart.", "Keep the hips mostly facing forward.", "Rotate the upper body to one side.", "Return to center and rotate to the other side."},
		Correct: []string{"Feet stay stable.", "Movement comes from the torso.", "Rotation is smooth.", "No twisting through the knees."},
		Wrong:   []string{"Turning the knees with the body.", "Rotating too fast.", "Leaning backward.", "Forcing the range of motion."},
	},
	"W04": {
		Desc:    "A movement drill for the hips, hamstrings and lower back control.",
		HowTo:   []string{"Stand tall with feet hip-width apart.", "Soften the knees slightly.", "Push the hips back.", "Lower the torso with a straight back.", "Return to standing by driving the hips forward."},
		Correct: []string{"Back stays straight.", "Movement starts from the hips.", "Knees are soft, not deeply bent.", "Neck follows the line of the spine."},
		Wrong:   []string{"Rounding the back.", "Turning it into a squat.", "Reaching too low.", "Lifting the head too high."},
	},
	"W05": {
		Desc:    "A controlled leg warm-up using a safe, shallow range.",
		HowTo:   []string{"Stand with feet slightly wider than hip-width.", "Push the hips slightly back.", "Bend the knees only to a comfortable depth.", "Return to standing slowly."},
		Correct: []string{"Knees follow the toes.", "Heels stay on the floor.", "Back stays long.", "Depth stays moderate."},
		Wrong:   []string{"Squatting too deep.", "Letting the knees collapse inward.", "Lifting the heels.", "Dropping down too fast."},
	},
	"W06": {
		Desc:    "A warm-up for the shoulders, wrists and core.",
		HowTo:   []string{"Start in a high plank position.", "Hands are under the shoulders.", "Shift the body slightly forward.", "Shift back to the starting position.", "Keep the movement small and controlled."},
		Correct: []string{"Body stays in one line.", "Shoulders stay stable.", "Core stays engaged.", "Movement is smooth."},
		Wrong:   []string{"Dropping the hips.", "Pushing the shoulders up to the ears.", "Moving too far forward.", "Letting the lower back sag."},
	},

	// ---- Abs / core ----
	"C01": {
		Desc:    "A core stability exercise for the abs, shoulders, glutes and trunk.",
		HowTo:   []string{"Place the elbows under the shoulders.", "Step the feet back.", "Make a straight line from shoulders to heels.", "Tighten the abs and glutes.", "Hold the position."},
		Correct: []string{"Body stays in one line.", "Elbows stay under the shoulders.", "Core stays active.", "Breathing stays steady."},
		Wrong:   []string{"Dropping the lower back.", "Lifting the hips too high.", "Looking forward.", "Holding the breath."},
	},
	"C02": {
		Desc:    "A short, intense plank focused on full-body tension.",
		HowTo:   []string{"Start in a forearm plank.", "Tighten the abs strongly.", "Squeeze the glutes.", "Imagine pulling the elbows toward the toes.", "Hold with maximum control."},
		Correct: []string{"Full-body tension.", "Glutes stay active.", "Core feels tight.", "Breathing stays controlled."},
		Wrong:   []string{"Holding it like a relaxed plank.", "Letting the hips sag.", "Holding the breath.", "Trying to hold too long."},
	},
	"C03": {
		Desc:    "A lower-ab exercise with controlled crossing movements of the legs.",
		HowTo:   []string{"Lie on the back.", "Press the lower back gently toward the floor.", "Lift the legs.", "Cross one leg over the other.", "Continue alternating the crossing motion."},
		Correct: []string{"Lower back stays close to the floor.", "Legs stay controlled.", "Neck stays relaxed.", "Movement stays smooth."},
		Wrong:   []string{"Arching the lower back.", "Lowering the legs too much.", "Pulling the head forward.", "Moving too fast."},
	},
	"C04": {
		Desc:    "A lower-ab exercise with short alternating leg kicks.",
		HowTo:   []string{"Lie on the back.", "Press the lower back toward the floor.", "Lift the legs.", "Move the legs up and down in short alternating kicks."},
		Correct: []string{"Small range of motion.", "Lower back stays stable.", "Legs do not drop too low.", "Breathing stays steady."},
		Wrong:   []string{"Arching the lower back.", "Making large kicks.", "Tensing the neck.", "Holding the breath."},
	},
	"C05": {
		Desc:    "An abdominal exercise focused on lifting the pelvis with control.",
		HowTo:   []string{"Lie on the back.", "Bend the knees and lift the feet.", "Use the abs to curl the pelvis slightly upward.", "Lower back down slowly."},
		Correct: []string{"Movement comes from the lower abs.", "Pelvis lifts with control.", "Neck stays relaxed.", "Lowering is slow."},
		Wrong:   []string{"Swinging the legs.", "Using momentum.", "Pulling on the neck.", "Dropping the hips quickly."},
	},
	"C06": {
		Desc:    "A static core hold with a shortened lever for better control.",
		HowTo:   []string{"Lie on the back.", "Press the lower back toward the floor.", "Lift the shoulders slightly.", "Keep the knees bent or legs raised higher.", "Hold the position."},
		Correct: []string{"Lower back stays on the floor.", "Abs stay tight.", "Neck stays relaxed.", "Breathing continues."},
		Wrong:   []string{"Arching the lower back.", "Lowering the legs too much.", "Pulling the chin forward.", "Holding the breath."},
	},
	"C07": {
		Desc:    "A lateral core exercise for the obliques, shoulder and hip stability.",
		HowTo:   []string{"Lie on one side.", "Place the elbow under the shoulder.", "Lift the hips.", "Keep the body in one straight line.", "Hold the position."},
		Correct: []string{"Elbow under shoulder.", "Hips stay lifted.", "Body stays long.", "Neck stays neutral."},
		Wrong:   []string{"Dropping the hips.", "Shrugging the shoulder.", "Rolling the chest forward.", "Holding the breath."},
	},
	"C08": {
		Desc:    "A stronger side-plank variation for the obliques and lateral hip.",
		HowTo:   []string{"Start in a side plank.", "Lower the hips slightly.", "Lift the hips back up.", "Repeat with control.", "Switch sides."},
		Correct: []string{"Movement is small and controlled.", "Elbow stays under the shoulder.", "Hips move vertically.", "Core stays tight."},
		Wrong:   []string{"Dropping too low.", "Twisting the torso.", "Collapsing into the shoulder.", "Moving too fast."},
	},
	"C09": {
		Desc:    "A core and shoulder exercise: tapping each shoulder while holding a plank.",
		HowTo:   []string{"Start in a high plank.", "Keep the hips level and stable.", "Tap one hand to the opposite shoulder.", "Return it and tap with the other hand.", "Keep the body still."},
		Correct: []string{"Hips stay level.", "Core stays tight.", "Shoulders stay stable.", "Movement is slow."},
		Wrong:   []string{"Rocking the hips side to side.", "Moving too fast.", "Dropping the lower back.", "Twisting the torso."},
	},
	"C10": {
		Desc:    "A core and triceps exercise moving between forearm plank and high plank.",
		HowTo:   []string{"Start in a forearm plank.", "Place one hand on the floor and press up.", "Bring the other hand up into a high plank.", "Lower back down to the forearms.", "Alternate the leading arm."},
		Correct: []string{"Hips stay stable.", "Core stays engaged.", "Movement is controlled.", "Shoulders stay over the hands or elbows."},
		Wrong:   []string{"Rocking the hips side to side.", "Dropping the lower back.", "Moving too fast.", "Letting the shoulders collapse."},
	},
	"C11": {
		Desc:    "A controlled core exercise and a safer alternative to harder leg-lift movements.",
		HowTo:   []string{"Lie on the back.", "Lift the legs into a tabletop position.", "Lower one heel toward the floor.", "Tap lightly and return.", "Alternate sides."},
		Correct: []string{"Lower back stays stable.", "Movement is slow.", "Heel touches softly.", "Pelvis stays steady."},
		Wrong:   []string{"Dropping the leg too fast.", "Arching the lower back.", "Moving both legs at once.", "Rushing the movement."},
	},
	"C12": {
		Desc:    "A core control exercise for the deep abs and spinal stability.",
		HowTo:   []string{"Lie on the back.", "Lift the arms and legs into tabletop position.", "Extend the opposite arm and leg.", "Keep the lower back stable.", "Return and switch sides."},
		Correct: []string{"Lower back stays close to the floor.", "Movement is slow.", "Opposite arm and leg move together.", "Torso stays still."},
		Wrong:   []string{"Arching the lower back.", "Moving too fast.", "Dropping the leg too low.", "Letting the ribs flare."},
	},

	// ---- Push-ups / triceps / shoulders ----
	"P01": {
		Desc:    "A basic upper-body exercise for the chest, triceps, shoulders and core.",
		HowTo:   []string{"Start in a high plank.", "Place the hands under the shoulders or slightly wider.", "Lower the body with control.", "Push back up while keeping the body straight."},
		Correct: []string{"Body stays in one line.", "Core stays tight.", "Elbows move at about 30–45 degrees.", "Movement is controlled."},
		Wrong:   []string{"Dropping the hips.", "Raising the hips too high.", "Flaring the elbows straight out.", "Falling into the bottom position."},
	},
	"P02": {
		Desc:    "A push-up variation with more focus on the triceps.",
		HowTo:   []string{"Start in a high plank.", "Place the hands closer than in a regular push-up.", "Keep the elbows near the body.", "Lower slowly and press back up."},
		Correct: []string{"Elbows stay close.", "Body stays straight.", "Shoulders stay away from the ears.", "Wrists feel comfortable."},
		Wrong:   []string{"Placing the hands too close.", "Flaring the elbows.", "Dropping the lower back.", "Forcing through wrist or elbow pain."},
	},
	"P03": {
		Desc:    "A controlled push-up with a short pause in the bottom position.",
		HowTo:   []string{"Start in a regular push-up position.", "Lower slowly.", "Pause for one second near the bottom.", "Push back up without losing body position."},
		Correct: []string{"Pause is controlled.", "Core stays tight.", "Chest stays active.", "Body remains straight."},
		Wrong:   []string{"Relaxing at the bottom.", "Dropping the hips.", "Holding the breath.", "Bouncing out of the bottom position."},
	},
	"P04": {
		Desc:    "A slow push-up variation that increases control and time under tension.",
		HowTo:   []string{"Start in a high plank.", "Lower for about three seconds.", "Pause for one second near the bottom.", "Push up for about three seconds."},
		Correct: []string{"Slow and even tempo.", "Body stays straight.", "Elbows stay controlled.", "Breathing stays steady."},
		Wrong:   []string{"Rushing the movement.", "Collapsing at the bottom.", "Letting the hips sag.", "Losing tension in the core."},
	},
	"P05": {
		Desc:    "An asymmetric push-up variation that challenges the chest, shoulders and core.",
		HowTo:   []string{"Start in a push-up position.", "Place one hand slightly forward and the other slightly back.", "Lower with control.", "Push back up.", "Switch hand position on the next side."},
		Correct: []string{"Body stays square.", "Core stays tight.", "Both shoulders stay stable.", "Movement is controlled."},
		Wrong:   []string{"Twisting the torso.", "Dropping one shoulder.", "Going too wide with the hands.", "Forcing the range if the shoulder feels uncomfortable."},
	},
	"P06": {
		Desc:    "A shoulder-focused push-up variation.",
		HowTo:   []string{"Start in a high plank.", "Walk the feet slightly closer and lift the hips.", "Keep the head between the arms.", "Lower the head toward the floor.", "Press back up."},
		Correct: []string{"Hips stay high.", "Elbows bend with control.", "Shoulders stay stable.", "Neck stays neutral."},
		Wrong:   []string{"Dropping into a regular push-up.", "Letting the head lead too far forward.", "Flaring the elbows too much.", "Forcing through shoulder pain."},
	},
	"P08": {
		Desc:    "A triceps and shoulder exercise performed from a forearm support position.",
		HowTo:   []string{"Start on the forearms with the body long.", "Keep the elbows under or slightly ahead of the shoulders.", "Press through the hands and forearms.", "Extend the elbows slightly.", "Return with control."},
		Correct: []string{"Core stays tight.", "Elbows move smoothly.", "Shoulders stay down.", "Range of motion stays comfortable."},
		Wrong:   []string{"Forcing the elbows.", "Dropping the hips.", "Shrugging the shoulders.", "Moving through elbow or shoulder pain."},
	},

	// ---- Back / posture ----
	"B01": {
		Desc:    "A posture-focused back exercise for the upper back, shoulder blades and spinal extensors.",
		HowTo:   []string{"Lie face down.", "Place the arms along the body.", "Turn the thumbs slightly outward.", "Lift the chest slightly.", "Pull the shoulder blades back and down.", "Hold."},
		Correct: []string{"Lift is small.", "Neck stays neutral.", "Shoulder blades move back and down.", "Glutes lightly engaged."},
		Wrong:   []string{"Lifting too high.", "Throwing the head back.", "Over-arching the lower back.", "Shrugging the shoulders."},
	},
	"B02": {
		Desc:    "An upper-back and shoulder-blade exercise performed lying face down.",
		HowTo:   []string{"Lie on the stomach.", "Keep the head facing down.", "Move the arms in a wide arc from the sides toward overhead.", "Return along the same path.", "Keep the movement controlled."},
		Correct: []string{"Neck stays relaxed.", "Shoulder blades move smoothly.", "Chest lift is small or neutral.", "Arms move without rushing."},
		Wrong:   []string{"Lifting the head.", "Arching the lower back too much.", "Swinging the arms.", "Shrugging the shoulders."},
	},
	"B03": {
		Desc:    "An upper-back exercise for the shoulder blades and rear shoulders.",
		HowTo:   []string{"Lie face down.", "Bend the elbows into a W shape.", "Lift the elbows and hands slightly.", "Squeeze the shoulder blades.", "Lower with control."},
		Correct: []string{"Small range of motion.", "Shoulder blades do the work.", "Neck stays neutral.", "Shoulders stay away from the ears."},
		Wrong:   []string{"Lifting too high.", "Pulling the head up.", "Using momentum.", "Over-arching the lower back."},
	},
	"B04": {
		Desc:    "A shoulder-blade exercise for the lower traps and upper back.",
		HowTo:   []string{"Lie face down.", "Reach the arms diagonally overhead into a Y shape.", "Lift the arms slightly from the floor.", "Lower with control."},
		Correct: []string{"Arms lift only a little.", "Neck stays neutral.", "Shoulders stay down.", "Movement is controlled."},
		Wrong:   []string{"Lifting too high.", "Shrugging the shoulders.", "Arching the lower back.", "Moving too fast."},
	},
	"B05": {
		Desc:    "An upper-back exercise for the rear shoulders and shoulder-blade control.",
		HowTo:   []string{"Lie face down.", "Extend the arms out to the sides in a T shape.", "Lift the arms slightly.", "Squeeze the shoulder blades.", "Lower slowly."},
		Correct: []string{"Arms stay in line with the shoulders.", "Neck stays relaxed.", "Shoulder blades move gently.", "Lift is small."},
		Wrong:   []string{"Throwing the arms upward.", "Lifting the head.", "Shrugging.", "Forcing a large range of motion."},
	},
	"B06": {
		Desc:    "A back exercise that mimics a pulling motion without equipment.",
		HowTo:   []string{"Lie on the stomach with arms overhead.", "Lift the chest slightly.", "Pull the elbows down toward the ribs.", "Squeeze the shoulder blades.", "Reach the arms forward again."},
		Correct: []string{"Elbows pull down with control.", "Shoulder blades squeeze gently.", "Neck stays neutral.", "Lower back stays comfortable."},
		Wrong:   []string{"Jerking the arms.", "Lifting too high.", "Throwing the head back.", "Letting the shoulders rise to the ears."},
	},
	"B07": {
		Desc:    "A back and coordination exercise performed lying face down.",
		HowTo:   []string{"Lie on the stomach.", "Lift the chest slightly.", "Move opposite arm and leg in a controlled swimming pattern.", "Keep the movement small and steady."},
		Correct: []string{"Neck stays neutral.", "Movement is controlled.", "Lower back stays comfortable.", "Breathing continues."},
		Wrong:   []string{"Moving too fast.", "Lifting too high.", "Throwing the head back.", "Forcing the lower back."},
	},
	"B08": {
		Desc:    "A posterior-chain exercise for the glutes, hamstrings, shoulders and back.",
		HowTo:   []string{"Sit on the floor with legs extended.", "Place the hands behind the hips.", "Press through the hands and heels.", "Lift the hips.", "Hold the body in a long line."},
		Correct: []string{"Chest stays open.", "Hips stay lifted.", "Shoulders stay stable.", "Neck stays neutral."},
		Wrong:   []string{"Dropping the hips.", "Shrugging the shoulders.", "Throwing the head back.", "Forcing the wrists."},
	},
	"B09": {
		Desc:    "A small-range back exercise for the spinal extensors and posture.",
		HowTo:   []string{"Lie face down.", "Place the hands beside the body or lightly near the chest.", "Lift the chest slightly.", "Make small controlled pulses.", "Keep the neck neutral."},
		Correct: []string{"Small lift.", "Smooth pulses.", "Neck stays long.", "Lower back feels controlled."},
		Wrong:   []string{"Pulsing too high.", "Using momentum.", "Looking forward.", "Feeling sharp pain in the lower back."},
	},
	"B10": {
		Desc:    "A stronger back exercise combining a small cobra lift with a W-shaped arm pull.",
		HowTo:   []string{"Lie face down with arms forward.", "Lift the chest slightly.", "Pull the elbows down and back into a W shape.", "Squeeze the shoulder blades.", "Reach forward again and lower with control."},
		Correct: []string{"Small chest lift.", "Smooth W-pull.", "Shoulder blades move back and down.", "Neck stays long."},
		Wrong:   []string{"Over-arching the lower back.", "Pulling with the neck.", "Moving too fast.", "Shrugging the shoulders."},
	},

	// ---- Arms / biceps ----
	"A01": {
		Desc:    "A no-equipment biceps exercise using one hand to resist the other.",
		HowTo:   []string{"Stand or sit tall.", "Bend one elbow like a curl.", "Use the opposite hand to press down on the working forearm.", "Curl upward slowly against resistance.", "Switch sides."},
		Correct: []string{"Resistance is steady.", "Movement is slow.", "Elbow stays near the body.", "Shoulder stays down."},
		Wrong:   []string{"Moving without real resistance.", "Jerking the arm.", "Raising the shoulder.", "Holding the breath."},
	},
	"A02": {
		Desc:    "A static biceps hold using self-resistance.",
		HowTo:   []string{"Bend one elbow to about 90 degrees.", "Use the opposite hand to press down.", "The working arm resists without moving.", "Hold the tension.", "Switch sides."},
		Correct: []string{"Elbow angle stays stable.", "Tension is steady.", "Shoulders stay relaxed.", "Breathing continues."},
		Wrong:   []string{"Pressing in short jerks.", "Twisting the torso.", "Raising the shoulder.", "Holding the breath."},
	},
	"A03": {
		Desc:    "An isometric arm and forearm exercise using the hands against each other.",
		HowTo:   []string{"Hook the fingers or hands together.", "Keep the elbows slightly bent.", "Pull the hands away from each other.", "Hold steady tension.", "Relax slowly."},
		Correct: []string{"Tension is controlled.", "Shoulders stay down.", "Wrists stay comfortable.", "Breathing stays calm."},
		Wrong:   []string{"Pulling with a jerk.", "Shrugging the shoulders.", "Over-bending the wrists.", "Holding the breath."},
	},
	"A04": {
		Desc:    "A slow lowering exercise for the biceps using self-resistance.",
		HowTo:   []string{"Start with one elbow bent.", "Use the opposite hand to create resistance.", "Slowly lower the working forearm.", "Keep resisting during the lowering phase.", "Switch sides."},
		Correct: []string{"Lowering is slow.", "Resistance stays constant.", "Elbow stays close to the body.", "Shoulder remains relaxed."},
		Wrong:   []string{"Letting the arm drop quickly.", "Using no resistance.", "Twisting the body.", "Holding the breath."},
	},

	// ---- Legs / glutes ----
	"L01": {
		Desc:    "A controlled leg exercise for the thighs, hips and posture.",
		HowTo:   []string{"Stand with feet slightly wider than hip-width.", "Push the hips back.", "Bend the knees to a comfortable shallow depth.", "Stand back up slowly."},
		Correct: []string{"Knees follow the toes.", "Heels stay down.", "Back stays long.", "Depth stays comfortable."},
		Wrong:   []string{"Squatting too deep.", "Letting the knees collapse inward.", "Lifting the heels.", "Dropping quickly."},
	},
	"L02": {
		Desc:    "A controlled squat variation with a short hold in a safe range.",
		HowTo:   []string{"Stand with feet stable.", "Lower into a comfortable shallow squat.", "Pause briefly.", "Stand back up with control."},
		Correct: []string{"Pause is stable.", "Knees track with the toes.", "Chest stays open.", "Heels stay on the floor."},
		Wrong:   []string{"Going too deep.", "Bouncing in the bottom position.", "Letting the knees cave in.", "Holding the breath."},
	},
	"L03": {
		Desc:    "A movement drill for the hips, hamstrings and lower back control.",
		HowTo:   []string{"Stand tall with feet hip-width apart.", "Soften the knees slightly.", "Push the hips back.", "Lower the torso with a straight back.", "Return to standing by driving the hips forward."},
		Correct: []string{"Back stays straight.", "Movement starts from the hips.", "Knees are soft, not deeply bent.", "Neck follows the line of the spine."},
		Wrong:   []string{"Rounding the back.", "Turning it into a squat.", "Reaching too low.", "Lifting the head too high."},
	},
	"L04": {
		Desc:    "A stronger glute exercise using one leg at a time.",
		HowTo:   []string{"Lie on the back.", "Bend one knee and keep that foot on the floor.", "Extend or lift the other leg.", "Push through the working heel.", "Lift the hips and lower slowly."},
		Correct: []string{"Working glute does the lift.", "Hips stay level.", "Movement is controlled.", "Lower back stays neutral."},
		Wrong:   []string{"Twisting the pelvis.", "Pushing through the toes.", "Arching the lower back.", "Dropping too fast."},
	},
	"L05": {
		Desc:    "A glute and pelvis-stability exercise.",
		HowTo:   []string{"Start in a glute bridge position.", "Keep the hips lifted.", "Lift one foot slightly from the floor.", "Put it down and switch sides.", "Keep the pelvis stable."},
		Correct: []string{"Hips stay level.", "Glutes stay active.", "Movement is slow.", "Lower back stays comfortable."},
		Wrong:   []string{"Letting the hips drop.", "Rocking side to side.", "Moving too fast.", "Arching the lower back."},
	},
	"L06": {
		Desc:    "A glute exercise that strengthens the hips and posterior chain.",
		HowTo:   []string{"Lie on the back.", "Bend the knees and place the feet on the floor.", "Lift the hips.", "Pause at the top.", "Lower slowly."},
		Correct: []string{"Push through the heels.", "Glutes squeeze at the top.", "Ribs stay down.", "Lower back does not over-arch."},
		Wrong:   []string{"Arching the lower back at the top.", "Lifting the heels.", "Dropping the hips too fast.", "Placing the feet too close."},
	},
	"L07": {
		Desc:    "A lower-leg exercise for the calves, ankles and foot control.",
		HowTo:   []string{"Stand tall with feet hip-width apart.", "Rise onto the balls of the feet.", "Pause briefly at the top.", "Lower the heels slowly."},
		Correct: []string{"Body stays upright.", "Movement is smooth.", "Ankles stay aligned.", "Lowering is controlled."},
		Wrong:   []string{"Dropping the heels quickly.", "Rocking forward.", "Rolling the ankles outward.", "Using momentum."},
	},
	"L09": {
		Desc:    "A lateral-chain exercise for the side glutes, obliques and hip stability.",
		HowTo:   []string{"Start in a side plank.", "Keep the body in a straight line.", "Lift the top leg slightly.", "Lower it with control.", "Repeat and switch sides."},
		Correct: []string{"Hips stay lifted.", "Top leg moves slowly.", "Torso stays stable.", "Shoulder stays strong."},
		Wrong:   []string{"Dropping the hips.", "Swinging the leg.", "Rolling the body backward.", "Collapsing into the shoulder."},
	},

	// ---- Cool-down ----
	"CD01": {
		Desc:    "A gentle stretch for the chest and front of the shoulders.",
		HowTo:   []string{"Lie on the back.", "Open the arms out to the sides.", "Relax the shoulders.", "Let the chest open gently.", "Breathe slowly."},
		Correct: []string{"Chest opens softly.", "Shoulders stay relaxed.", "Lower back stays comfortable.", "No pain in the shoulders."},
		Wrong:   []string{"Forcing the arms down.", "Arching the lower back.", "Tensing the neck.", "Stretching through pain."},
	},
	"CD02": {
		Desc:    "A gentle stretch for the back of the thigh.",
		HowTo:   []string{"Lie on the back.", "Lift one leg.", "Hold behind the thigh or calf.", "Pull the leg gently toward you.", "Switch sides."},
		Correct: []string{"Stretch feels mild.", "Knee can stay slightly bent.", "Lower back stays calm.", "Breathing stays slow."},
		Wrong:   []string{"Pulling too hard.", "Trying to force the leg straight.", "Lifting the hips.", "Feeling sharp pain behind the knee."},
	},
	"CD03": {
		Desc:    "A gentle rotation stretch for the spine and trunk.",
		HowTo:   []string{"Lie on the back.", "Bend the knees.", "Let the knees move gently to one side.", "Keep the shoulders on the floor.", "Return to center and switch sides."},
		Correct: []string{"Movement is slow.", "Shoulders stay down.", "Breathing is relaxed.", "Stretch stays comfortable."},
		Wrong:   []string{"Forcing the knees to the floor.", "Twisting quickly.", "Lifting the shoulder.", "Stretching through pain."},
	},
	"CD04": {
		Desc:    "A gentle front-body stretch and mild back extension.",
		HowTo:   []string{"Lie on the stomach.", "Place the forearms on the floor.", "Lift the chest gently.", "Keep the neck long.", "Breathe slowly."},
		Correct: []string{"Lift is mild.", "Shoulders stay away from the ears.", "Lower back feels comfortable.", "Neck stays neutral."},
		Wrong:   []string{"Over-arching the lower back.", "Throwing the head back.", "Shrugging the shoulders.", "Forcing the position."},
	},
	"CD05": {
		Desc:    "A calm breathing exercise to finish the session.",
		HowTo:   []string{"Lie on the back.", "Relax the shoulders and jaw.", "Place one hand on the belly if comfortable.", "Inhale gently.", "Exhale slowly."},
		Correct: []string{"Breathing is calm.", "Shoulders stay relaxed.", "Face stays soft.", "Body settles down."},
		Wrong:   []string{"Breathing too forcefully.", "Holding the breath.", "Tensing the neck.", "Arching the lower back."},
	},
}
