// body/poses.mjs — one pose per exercise ID for the headless figure renderer (shot.html).
// Pose = joint direction vectors (our space: x=front, y=up, z=left/+L). Omitted joints fall
// back to the flat-foot neutral. `view`: front | side | iso. Feet default to flat rig-rest
// unless overridden (e.g. plank on toes). perSide exercises show one side.
// Authored from exercises.go names + details HowTo + the spec's 18-pose table.

const DN=[0,-1,0];

// ======================= discrete pose vocabulary =======================
// RULES (snap to these, no arbitrary angles):
//  - limb directions come from DIR only (cardinals + 45° diagonals), body frame x=front,y=up,z=left.
//  - foot: POINTED (in line with shin) or FLEXED (~90° to shin). nothing in between.
//  - head: pitch in line with torso (never cranes); yaw ∈ {0,+70,-70} via spec.headYaw.
//  - hand: support = flat on floor (horizontal); free = continues the forearm (omit it).
//  - elbow/knee: straight (~0°) or bent (~90°), not a little.
//  - default L/R mirrored; asymmetry only on purpose (perSide / steps / swings).
const DIR = {
  down:[0,-1,0], up:[0,1,0], fwd:[1,0,0], back:[-1,0,0], left:[0,0,1], right:[0,0,-1],
  fwdDown:[0.71,-0.71,0], fwdUp:[0.71,0.71,0], backDown:[-0.71,-0.71,0], backUp:[-0.71,0.71,0],
  downLeft:[0,-0.71,0.71], downRight:[0,-0.71,-0.71], upLeft:[0,0.71,0.71], upRight:[0,0.71,-0.71],
};
// foot relative to a given shin direction sh: pointed = along shin; flexed = 90° toward the sole.
const norm=v=>{const m=Math.hypot(v[0],v[1],v[2])||1;return[v[0]/m,v[1]/m,v[2]/m];};
const footPointed = sh => norm(sh);                   // toe continues the shin (e.g. tiptoe, lying point)
const footFlexed  = sh => norm([-sh[1], sh[0], 0]);   // 90° to shin toward the toes-forward side (sole flat); standing shin[0,-1,0]→[1,0,0]

// ---------- reusable archetypes ----------
// straight body tilted ~22° (shoulders high on vertical arms, feet low), belly down (back up),
// palms flat forward on the floor, toes down — hands and toes both reach the ground.
const PUSHUP_TOP = {spine:[-0.978,0.21,0],front:[-0.21,-0.978,0],head:[-0.94,0.18,0],
  Luarm:[-0.05,-1,0],Lfarm:[-0.05,-1,0],Lhand:[-1,0,0],
  Ruarm:[-0.05,-1,0],Rfarm:[-0.05,-1,0],Rhand:[-1,0,0],
  Lthigh:[0.978,-0.21,0],Lshin:[0.978,-0.21,0],Lfoot:[0.4,-0.85,0],
  Rthigh:[0.978,-0.21,0],Rshin:[0.978,-0.21,0],Rfoot:[0.4,-0.85,0]};
// forearm plank: upper arm straight DOWN to the elbow, forearm flat FORWARD along the floor,
// body a straight line, toes down. Forearms + toes are the two floor contacts.
const FOREARM_PLANK = {spine:[-0.9995,0.03,0],front:[-0.03,-0.9995,0],head:[-0.999,0.03,0],
  Luarm:[0,-1,0],Lfarm:[-1,0,0],Lhand:[-1,0,0],
  Ruarm:[0,-1,0],Rfarm:[-1,0,0],Rhand:[-1,0,0],
  Lthigh:[0.9995,-0.03,0],Lshin:[0.9995,-0.03,0],Lfoot:[0.4,-0.85,0],
  Rthigh:[0.9995,-0.03,0],Rshin:[0.9995,-0.03,0],Rfoot:[0.4,-0.85,0]};
// side plank (front view): diagonal body on its side; RIGHT arm down = support (hand on
// floor), LEFT arm up; legs stacked sloping to the feet on the floor. Contacts: hand + feet.
const SIDEPLANK = {spine:[0,0.36,0.933],front:[1,0,0],head:[0,0.63,0.78],
  Ruarm:[0,-1,0],Rfarm:[0,-1,0],Rhand:[1,0,0], Luarm:[0,1,0],Lfarm:[0,1,0],Lhand:[0,1,0],   // Rhand flat on floor (full palm)
  Rthigh:[0,-0.36,-0.933],Rshin:[0,-0.36,-0.933],Rfoot:[0,-0.5,-0.87],
  Lthigh:[0,-0.36,-0.933],Lshin:[0,-0.36,-0.933],Lfoot:[0,-0.5,-0.87]};
const SQUAT = {spine:[0.42,0.91,0],front:[1,0,0],
  Luarm:[0.6,0,0],Lfarm:[0.92,0,0],Ruarm:[0.6,0,0],Rfarm:[0.92,0,0],
  Lthigh:[0.95,-0.12,0],Lshin:[-0.6,-0.8,0],Rthigh:[0.95,-0.12,0],Rshin:[-0.6,-0.8,0]};
const HINGE = {spine:[0.75,0.66,0],front:[1,0,0],
  Luarm:[0.2,-0.97,0],Lfarm:[0.2,-0.97,0],Ruarm:[0.2,-0.97,0],Rfarm:[0.2,-0.97,0],
  Lthigh:[0.1,-0.99,0],Lshin:[0,-1,0],Rthigh:[0.1,-0.99,0],Rshin:[0,-1,0]};
// supine = lying on back (head toward -x, face up); legs/arms set per move
const SUPINE = {spine:[-1,0.05,0],front:[0,1,0],
  Luarm:[-0.9,0.2,0],Lfarm:[-0.9,0.2,0],Ruarm:[-0.9,0.2,0],Rfarm:[-0.9,0.2,0],
  Lthigh:[-1,0,0],Lshin:[-1,0,0],Lfoot:[-0.7,-0.7,0],Rthigh:[-1,0,0],Rshin:[-1,0,0],Rfoot:[-0.7,-0.7,0]};
// prone = lying face DOWN along the body axis; head toward +x, belly down (-y), legs extend
// the OPPOSITE way (toward -x) so the body is a flat line, not folded into a crouch.
// PRONE: body lies FLAT face-down (chest+belly+thighs+feet on floor); head +x; legs straight
// back along the floor; feet pointed (tops down, toes back — NOT toes-up). Raises lift arms only.
const PRONE = {spine:[1,0.03,0],front:[0.03,-1,0],head:[0.94,0.08,0],
  Luarm:[0.7,-0.3,0.2],Lfarm:[0.9,-0.1,0.1],Ruarm:[0.7,-0.3,-0.2],Rfarm:[0.9,-0.1,-0.1],
  Lthigh:[-0.99,-0.12,0],Lshin:[-0.99,-0.12,0],Lfoot:[-0.8,-0.55,0],
  Rthigh:[-0.99,-0.12,0],Rshin:[-0.99,-0.12,0],Rfoot:[-0.8,-0.55,0]};
// glute bridge: shoulders/head on the floor, pelvis LIFTED (chest below pelvis), knees bent
// with feet flat on the floor. spine points down-toward-head. contacts: shoulders + feet.
const BRIDGE = {spine:[-0.72,-0.69,0],front:[-0.69,0.72,0],head:[-0.99,-0.04,0],
  Luarm:[0.6,-0.35,0.12],Lfarm:[0.8,-0.25,0.08],Ruarm:[0.6,-0.35,-0.12],Rfarm:[0.8,-0.25,-0.08],
  Lthigh:[0.35,-0.94,0],Lshin:[0.18,-0.98,0],Lfoot:[0.55,-0.65,0],
  Rthigh:[0.35,-0.94,0],Rshin:[0.18,-0.98,0],Rfoot:[0.55,-0.65,0]};
const STAND = {view:'front'};   // plain neutral

export const POSES = {
  _neutral: STAND,

  // ---- Warm-up ----
  W01: {view:'front', pose:{Luarm:[0,-0.95,0.1],Lfarm:[0,-0.95,0.1],Ruarm:[0,-0.95,-0.1],Rfarm:[0,-0.95,-0.1]}}, // shoulder circles
  W02: {view:'front', pose:{Luarm:[0,0.1,1],Lfarm:[0,0.1,1],Ruarm:[0,0.1,-1],Rfarm:[0,0.1,-1]}},                 // arm circles (arms out)
  W03: {view:'front', pose:{front:[0.7,0,0.7],Luarm:[0.3,0,0.6],Lfarm:[0.6,0,0.4],Ruarm:[0.3,0,-0.6],Rfarm:[0.6,0,-0.4]}}, // standing twist
  W04: {view:'side', pose:HINGE},                                                                                  // hip hinge
  W05: {view:'side', pose:SQUAT}, // shallow squat (use the verified squat)
  W06: {view:'side', pose:PUSHUP_TOP},                                                                            // high plank
  W07: {view:'front', pose:{Luarm:[0,1,0.1],Lfarm:[0,1,0.1],Ruarm:[0,1,-0.1],Rfarm:[0,1,-0.1]}},                 // breathing + reach (overhead)
  W08: {view:'front', pose:{front:[0.7,0,0.7],Luarm:[0.2,0,0.7],Lfarm:[0.5,0,0.5],Ruarm:[0.2,0,-0.7],Rfarm:[0.5,0,-0.5]}}, // torso rotations
  W09: {view:'side', pose:{Luarm:[1,0.1,0],Lfarm:[1,0.1,0],Ruarm:[-0.6,-0.5,0],Rfarm:[-0.6,-0.5,0]}},            // arm swings (one fwd/one back)

  // ---- Abs / core ----
  C01: {view:'side', pose:FOREARM_PLANK},                                                                         // forearm plank
  C02: {view:'side', pose:FOREARM_PLANK},                                                                         // RKC plank
  C03: {view:'side', pose:{...SUPINE, Lthigh:[0.45,0.89,0],Lshin:[0.45,0.89,0],Lfoot:[0.45,0.89,0],Rthigh:[0.66,0.75,0],Rshin:[0.66,0.75,0],Rfoot:[0.66,0.75,0]}}, // scissors
  C04: {view:'side', pose:{...SUPINE, Lthigh:[0.35,0.93,0],Lshin:[0.35,0.93,0],Rthigh:[0.7,0.7,0],Rshin:[0.7,0.7,0]}}, // flutter kicks
  C05: {view:'side', pose:{...SUPINE, Lthigh:[0.2,0.6,0],Lshin:[-0.7,0.4,0],Rthigh:[0.2,0.6,0],Rshin:[-0.7,0.4,0]}}, // reverse crunch (knees to chest)
  C06: {view:'side', pose:{...SUPINE, Luarm:[-1,0.1,0],Lfarm:[-1,0.1,0],Ruarm:[-1,0.1,0],Rfarm:[-1,0.1,0], Lthigh:[0.5,0.5,0],Lshin:[0.5,0.5,0],Rthigh:[0.5,0.5,0],Rshin:[0.5,0.5,0]}}, // hollow hold
  C07: {view:'front', pose:SIDEPLANK}, // side plank
  C08: {view:'front', pose:SIDEPLANK}, // side plank hip lift (same hold)
  C09: {view:'side', pose:{...PUSHUP_TOP, Ruarm:[0.3,-0.5,0],Rfarm:[1,0.1,0]}},                                  // plank shoulder taps (one hand up)
  C10: {view:'side', pose:PUSHUP_TOP},                                                                            // up-down plank
  // heel taps: on back (back on floor), knees bent up, hands reach up toward the knees
  C11: {view:'side', pose:{...SUPINE, Lthigh:[0.35,0.6,0],Lshin:[-0.45,0.5,0],Rthigh:[0.35,0.6,0],Rshin:[-0.45,0.5,0], Luarm:[0.3,0.4,0],Lfarm:[0.55,0.55,0],Ruarm:[0.3,0.4,0],Rfarm:[0.55,0.55,0]}}, // heel taps
  C12: {view:'side', pose:{...SUPINE, Luarm:[0,0.9,0],Lfarm:[0,0.9,0], Ruarm:[0,0.9,0],Rfarm:[0,0.9,0], Lthigh:[0.3,0.6,0],Lshin:[-0.6,0.4,0], Rthigh:[0.3,0.6,0],Rshin:[-0.6,0.4,0]}}, // dead bug
  C13: {view:'side', pose:{...SUPINE, Luarm:[-0.3,0.6,0],Lfarm:[0.4,0.5,0], Ruarm:[-0.3,0.6,0],Rfarm:[0.4,0.5,0], Lthigh:[0.3,0.55,0],Lshin:[-0.6,0.3,0], Rthigh:[0.7,0.4,0],Rshin:[0.85,0.2,0]}}, // bicycle crunch

  // ---- Push-ups / shoulders ----
  P01: {view:'side', pose:PUSHUP_TOP}, P02: {view:'side', pose:PUSHUP_TOP}, P03: {view:'side', pose:PUSHUP_TOP},
  P04: {view:'side', pose:PUSHUP_TOP}, P05: {view:'side', pose:{...PUSHUP_TOP, Luarm:[0.3,-0.95,0]}},
  P06: {view:'side', pose:{spine:[-0.5,0.7,0],front:[-0.7,-0.7,0], Luarm:[-0.4,-0.9,0],Lfarm:[-0.4,-0.9,0],Ruarm:[-0.4,-0.9,0],Rfarm:[-0.4,-0.9,0], Lthigh:[0.85,-0.5,0],Lshin:[0.85,-0.5,0],Lfoot:[0.5,-0.86,0],Rthigh:[0.85,-0.5,0],Rshin:[0.85,-0.5,0],Rfoot:[0.5,-0.86,0]}}, // pike push-up
  P08: {view:'side', pose:{...PRONE, Luarm:[0.2,-0.95,0],Lfarm:[0.6,-0.3,0],Ruarm:[0.2,-0.95,0],Rfarm:[0.6,-0.3,0], spine:[0.7,0.66,0],front:[0.2,-0.97,0]}}, // sphinx push-up

  // ---- Back / scapula (prone) ----
  B01: {view:'side', pose:{...PRONE, spine:[0.8,0.55,0],front:[0.4,-0.9,0], Luarm:[0.6,-0.4,0.4],Lfarm:[0.8,-0.2,0.3],Ruarm:[0.6,-0.4,-0.4],Rfarm:[0.8,-0.2,-0.3]}}, // cobra hold
  B02: {view:'side', pose:{...PRONE, Luarm:[0.2,0.1,0.95],Lfarm:[0.2,0.1,0.95],Ruarm:[0.2,0.1,-0.95],Rfarm:[0.2,0.1,-0.95]}}, // reverse snow angels
  B03: {view:'side', pose:{...PRONE, Luarm:[0.3,0.5,0.6],Lfarm:[0.1,0.9,0.3],Ruarm:[0.3,0.5,-0.6],Rfarm:[0.1,0.9,-0.3]}}, // W-raise
  B04: {view:'side', pose:{...PRONE, Luarm:[0.3,0.8,0.4],Lfarm:[0.3,0.9,0.3],Ruarm:[0.3,0.8,-0.4],Rfarm:[0.3,0.9,-0.3]}}, // Y-raise
  B05: {view:'side', pose:{...PRONE, Luarm:[0.1,0.2,1],Lfarm:[0.1,0.2,1],Ruarm:[0.1,0.2,-1],Rfarm:[0.1,0.2,-1]}}, // T-raise
  // superman: the one prone move that lifts arms AND legs off the floor
  B06: {view:'side', pose:{...PRONE, spine:[0.97,0.22,0],front:[0.22,-0.97,0],
    Luarm:[0.9,0.35,0.1],Lfarm:[0.95,0.25,0.05],Ruarm:[0.9,0.35,-0.1],Rfarm:[0.95,0.25,-0.05],
    Lthigh:[-0.9,0.35,0],Lshin:[-0.92,0.3,0],Lfoot:[-0.95,0.15,0],
    Rthigh:[-0.9,0.35,0],Rshin:[-0.92,0.3,0],Rfoot:[-0.95,0.15,0]}}, // superman (arms + legs lifted)
  B07: {view:'side', pose:{...PRONE, Luarm:[0.6,0.5,0.25],Lfarm:[0.5,0.7,0.2],Ruarm:[0.85,-0.2,-0.3],Rfarm:[0.95,-0.1,-0.2], Lthigh:[-0.92,0.32,0],Lshin:[-0.94,0.3,0], Rthigh:[-0.97,-0.18,0],Rshin:[-0.98,-0.12,0]}}, // swimmers (opp arm+leg up)
  // reverse plank: face UP, straight body; arms down to flat hands behind (support), legs
  // extended fwd-down to heels (feet flexed). contacts: hands + heels.
  B08: {view:'side', pose:{spine:[-0.94,0.34,0],front:[0.34,0.94,0],head:[-0.9,0.35,0],
    Luarm:[-0.2,-0.98,0],Lfarm:[-0.2,-0.98,0],Lhand:[1,0,0], Ruarm:[-0.2,-0.98,0],Rfarm:[-0.2,-0.98,0],Rhand:[1,0,0],
    Lthigh:[0.94,-0.34,0],Lshin:[0.94,-0.34,0],Lfoot:[0.34,0.94,0],
    Rthigh:[0.94,-0.34,0],Rshin:[0.94,-0.34,0],Rfoot:[0.34,0.94,0]}}, // reverse plank
  B09: {view:'side', pose:{...PRONE, spine:[0.85,0.5,0],front:[0.3,-0.95,0], Luarm:[0.3,0.1,0.5],Lfarm:[0.5,0.1,0.4],Ruarm:[0.3,0.1,-0.5],Rfarm:[0.5,0.1,-0.4]}}, // back extension pulses
  B10: {view:'side', pose:{...PRONE, Luarm:[0.3,0.5,0.6],Lfarm:[0.1,0.85,0.4],Ruarm:[0.3,0.5,-0.6],Rfarm:[0.1,0.85,-0.4]}}, // cobra to W
  B11: {view:'side', pose:{...PRONE, Luarm:[0.6,0.5,0.25],Lfarm:[0.5,0.7,0.2],Ruarm:[0.85,-0.2,-0.3],Rfarm:[0.95,-0.1,-0.2], Lthigh:[-0.92,0.32,0],Lshin:[-0.94,0.3,0], Rthigh:[-0.97,-0.18,0],Rshin:[-0.98,-0.12,0]}}, // swimmers per side

  // ---- Arms ----
  A01: {view:'front', pose:{Luarm:[0.25,-0.9,0],Lfarm:[0.5,0.75,0], Ruarm:DN,Rfarm:DN}}, // self-resist curl (one arm curled up)
  A02: {view:'front', pose:{Luarm:[0.1,-0.95,0],Lfarm:[1,0,0], Ruarm:DN,Rfarm:DN}},     // biceps isometric 90
  A03: {view:'front', pose:{Luarm:[0.2,-0.6,0],Lfarm:[0.7,0.4,0.2], Ruarm:[0.2,-0.6,0],Rfarm:[0.7,0.4,-0.2]}}, // hand-hook pull (hands at chest)
  A04: {view:'front', pose:{Luarm:[0.2,-0.9,0],Lfarm:[0.85,-0.5,0], Ruarm:DN,Rfarm:DN}}, // slow negative curl (mid)

  // ---- Legs / glutes ----
  L01: {view:'side', pose:SQUAT}, L02: {view:'side', pose:SQUAT}, L03: {view:'side', pose:HINGE},
  // single-leg bridge: R foot on floor (bridge), LEFT leg extended straight in line with the body
  L04: {view:'side', pose:{...BRIDGE, Lthigh:[0.62,0.05,0],Lshin:[0.62,0.05,0],Lfoot:[0.7,0.4,0]}}, // single-leg glute bridge
  // bridge march: R foot on floor (bridge), LEFT knee drawn up
  L05: {view:'side', pose:{...BRIDGE, Lthigh:[0.5,0.4,0],Lshin:[-0.2,0.5,0],Lfoot:[0.5,0.4,0]}}, // bridge march
  L06: {view:'side', pose:BRIDGE}, // glute bridge hold
  L07: {view:'side', pose:{Lfoot:[0.85,-0.5,0],Rfoot:[0.85,-0.5,0], spine:[0,1,0]}},   // calf raises (on toes)
  // side plank + top (left) leg lifted up off the bottom leg
  L09: {view:'front', pose:{...SIDEPLANK, Lthigh:[0,0.25,-0.97],Lshin:[0,0.25,-0.97],Lfoot:[0,-0.2,-0.98]}}, // side plank leg lift
  L10: {view:'side', pose:SQUAT}, L12: {view:'side', pose:SQUAT},
  L11: {view:'side', pose:{spine:[0.05,1,0],front:[1,0,0], Luarm:[0.3,-0.9,0],Lfarm:[0.3,-0.9,0],Ruarm:[0.3,-0.9,0],Rfarm:[0.3,-0.9,0], Lthigh:[0.5,-0.86,0],Lshin:[0.4,-0.9,0],Rthigh:[-0.6,-0.8,0],Rshin:[0.2,-0.98,0],Rfoot:[0.6,-0.8,0]}}, // reverse lunge

  // ---- Cardio / plyometrics ----
  J01: {view:'front', pose:{Luarm:[0.2,-0.9,0],Lfarm:[0.2,-0.9,0],Ruarm:[0.2,-0.9,0],Rfarm:[0.2,-0.9,0], Lfoot:[0.85,-0.5,0],Rfoot:[0.85,-0.5,0]}}, // pogo (on toes, feet together)
  J02: {view:'side', pose:{Luarm:[0.55,-0.55,0],Lfarm:[0.95,0.1,0],Ruarm:[-0.5,-0.7,0],Rfarm:[-0.65,-0.55,0], Lthigh:[0.55,-0.5,0],Lshin:[-0.05,-1,0],Lfoot:[0.6,-0.5,0], Rthigh:[0,-1,0],Rshin:[0,-1,0]}}, // marching high knees (thigh fwd-down, shin hangs)
  J03: {view:'front', pose:{Luarm:[0,1,0.55],Lfarm:[0,1,0.55],Ruarm:[0,1,-0.55],Rfarm:[0,1,-0.55], Lthigh:[0,-1,0.45],Lshin:[0,-1,0.45],Rthigh:[0,-1,-0.45],Rshin:[0,-1,-0.45]}}, // jumping jacks
  J04: {view:'side', pose:{...PUSHUP_TOP, Lthigh:[0.4,-0.4,0],Lshin:[-0.6,-0.5,0],Lfoot:[0.5,-0.5,0]}}, // mountain climbers (one knee in)
  J05: {view:'side', pose:{spine:[0.2,0.97,0],front:[1,0,0], Luarm:[-0.5,-0.7,0],Lfarm:[-0.6,-0.6,0],Ruarm:[-0.5,-0.7,0],Rfarm:[-0.6,-0.6,0], Lthigh:[0.3,-0.95,0],Lshin:[-0.05,-1,0],Lfoot:[0.9,-0.3,0],Rthigh:[0.3,-0.95,0],Rshin:[-0.05,-1,0],Rfoot:[0.9,-0.3,0]}}, // squat jump (takeoff)
  J06: {view:'front', pose:{spine:[0.1,0.99,0.1],front:[1,0,0], Luarm:[0.3,0.2,0.7],Lfarm:[0.3,0.2,0.7],Ruarm:[0.2,-0.6,-0.5],Rfarm:[0.2,-0.6,-0.5], Lthigh:[0.3,-0.9,0.2],Lshin:[0,-1,0], Rthigh:[-0.2,-0.5,-0.7],Rshin:[-0.4,-0.7,-0.5]}}, // skater hop (land one leg)
  J07: {view:'side', pose:{Luarm:[0.6,-0.5,0],Lfarm:[0.95,0.2,0],Ruarm:[-0.6,-0.6,0],Rfarm:[-0.7,-0.5,0], Lthigh:[0.72,-0.32,0],Lshin:[-0.1,-0.99,0],Lfoot:[0.6,-0.5,0], Rthigh:[0,-1,0],Rshin:[0,-1,0],Rfoot:[0.85,-0.5,0]}}, // high knees (thigh up but <horizontal, shin down)
  J08: {view:'side', pose:PUSHUP_TOP},  // burpee (plank phase)
  J09: {view:'side', pose:{spine:[0.1,0.99,0],front:[1,0,0], Luarm:[0.4,-0.85,0],Lfarm:[0.4,-0.85,0],Ruarm:[-0.4,-0.85,0],Rfarm:[-0.4,-0.85,0], Lthigh:[0.55,-0.83,0],Lshin:[0.45,-0.89,0], Rthigh:[-0.6,-0.8,0],Rshin:[0.1,-0.99,0],Rfoot:[0.7,-0.7,0]}}, // jump lunge (air)

  // ---- Cool-down ----
  CD01: {view:'iso', pose:{Luarm:[-0.55,-0.2,0.45],Lfarm:[-0.7,-0.1,0.35],Ruarm:[-0.55,-0.2,-0.45],Rfarm:[-0.7,-0.1,-0.35]}}, // chest stretch (arms back, hands clasped behind)
  CD02: {view:'side', pose:{...SUPINE, Lthigh:[0.2,0.95,0],Lshin:[0.2,0.95,0],Lfoot:[0.2,0.95,0], Rthigh:[1,-0.05,0],Rshin:[1,-0.05,0],Rfoot:[0.7,-0.6,0], Luarm:[0.2,0.9,0],Lfarm:[0.3,0.85,0],Ruarm:[0.2,0.9,0],Rfarm:[0.3,0.85,0]}}, // hamstring stretch (one leg up, other extended)
  CD03: {view:'iso', pose:{...SUPINE, front:[0.4,0.55,0.7], Lthigh:[0.3,0,0.92],Lshin:[0.7,-0.2,0.65], Rthigh:[1,-0.05,0],Rshin:[1,-0.05,0], Luarm:[0,0.1,1],Lfarm:[0,0.1,1],Ruarm:[0,0.1,-1],Rfarm:[0,0.1,-1]}}, // supine twist (one knee across)
  CD04: {view:'side', pose:{...PRONE, spine:[0.75,0.66,0],front:[0.2,-0.97,0], Luarm:[0.3,-0.9,0],Lfarm:[0.6,-0.2,0],Ruarm:[0.3,-0.9,0],Rfarm:[0.6,-0.2,0]}}, // sphinx pose
  CD05: {view:'side', pose:PRONE},  // lying breathing (prone relaxed)
  CD07: {view:'side', pose:{spine:[0.93,0.33,0],front:[0.28,-0.96,0], Luarm:[0.8,-0.25,0],Lfarm:[0.96,-0.12,0],Ruarm:[0.8,-0.25,0],Rfarm:[0.96,-0.12,0], Lthigh:[0.18,-0.98,0],Lshin:[-0.92,-0.32,0],Lfoot:[-0.85,-0.2,0], Rthigh:[0.18,-0.98,0],Rshin:[-0.92,-0.32,0],Rfoot:[-0.85,-0.2,0]}}, // child's pose: kneel, fold torso forward-down, arms reach forward
};
