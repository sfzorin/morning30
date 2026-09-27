# Neutral mannequin, two frames per exercise, one locked camera.
# Blender --background --python body/blender/render_shots.py -- --only P01,L10
import math
import os
import shutil
import sys

import bpy
from mathutils import Matrix, Vector

ROOT = os.path.abspath(os.path.join(os.path.dirname(__file__), "..", "out"))
os.makedirs(ROOT, exist_ok=True)
CHAR_DIR = os.path.abspath(os.path.join(os.path.dirname(__file__), "..", "chars", "ubc"))

F = Vector((0, -1, 0))   # face / forward
B = Vector((0, 1, 0))    # back
U = Vector((0, 0, 1))
D = Vector((0, 0, -1))
LFT = Vector((1, 0, 0))  # character left
RGT = Vector((-1, 0, 0))

SKIN = (0.90, 0.78, 0.68, 1)
SHIRT = (0.29, 0.46, 0.66, 1)
SHORTS = (0.24, 0.29, 0.36, 1)
HAIR = (0.18, 0.15, 0.13, 1)
DBCOL = (0.55, 0.57, 0.60, 1)

# head, tail, parent. Facing -Y, Z up, character-left is +X.
BONES = [
    ("hips", (0, 0, 0.96), (0, 0, 1.08), None),
    ("spine", (0, 0, 1.08), (0, 0, 1.28), "hips"),
    ("chest", (0, 0, 1.28), (0, 0, 1.50), "spine"),
    ("neck", (0, 0, 1.50), (0, 0, 1.58), "chest"),
    ("head", (0, 0, 1.58), (0, 0, 1.78), "neck"),
    ("clavicle.L", (0.04, 0, 1.46), (0.20, 0, 1.46), "chest"),
    ("clavicle.R", (-0.04, 0, 1.46), (-0.20, 0, 1.46), "chest"),
    ("upper_arm.L", (0.20, 0, 1.46), (0.24, 0, 1.14), "clavicle.L"),
    ("upper_arm.R", (-0.20, 0, 1.46), (-0.24, 0, 1.14), "clavicle.R"),
    ("forearm.L", (0.24, 0, 1.14), (0.22, 0, 0.84), "upper_arm.L"),
    ("forearm.R", (-0.24, 0, 1.14), (-0.22, 0, 0.84), "upper_arm.R"),
    ("hand.L", (0.22, 0, 0.84), (0.20, -0.02, 0.70), "forearm.L"),
    ("hand.R", (-0.22, 0, 0.84), (-0.20, -0.02, 0.70), "forearm.R"),
    ("thigh.L", (0.10, 0, 0.94), (0.11, 0, 0.50), "hips"),
    ("thigh.R", (-0.10, 0, 0.94), (-0.11, 0, 0.50), "hips"),
    ("shin.L", (0.11, 0, 0.50), (0.10, 0, 0.08), "thigh.L"),
    ("shin.R", (-0.11, 0, 0.50), (-0.10, 0, 0.08), "thigh.R"),
    ("foot.L", (0.10, 0.0, 0.08), (0.10, -0.20, 0.03), "shin.L"),
    ("foot.R", (-0.10, 0.0, 0.08), (-0.10, -0.20, 0.03), "shin.R"),
]

TORSO = ("hips", "spine", "chest", "neck", "head")
ARM_B = ("upper_arm", "forearm", "hand")
LEG_B = ("thigh", "shin", "foot")


def nvec(*a):
    v = Vector(a)
    return v.normalized() if v.length else Vector((0, 0, 1))


def mirror(v):
    return Vector((-v.x, v.y, v.z))


def basis(direction, up):
    y = Vector(direction).normalized()
    x = y.cross(Vector(up))
    if x.length < 1e-6:
        x = y.cross(Vector((1, 0, 0)))
    if x.length < 1e-6:
        x = y.cross(Vector((0, 1, 0)))
    x.normalize()
    z = x.cross(y).normalized()
    x = y.cross(z).normalized()
    return Matrix(((x.x, y.x, z.x), (x.y, y.y, z.y), (x.z, y.z, z.z)))


def aim(pb, direction, up):
    up = Vector(up)
    # Thighs and feet are rolled so local Z points back. Flip the hint so the knee stays forward.
    rest_z = pb.bone.matrix_local.to_3x3().col[2]
    if rest_z.normalized().dot(F) < -0.2:
        up = -up
    desired = basis(direction, up)
    loc = pb.matrix.to_translation()
    pb.matrix = Matrix.Translation(loc) @ desired.to_4x4()


def new_mat(name, rgba):
    m = bpy.data.materials.new(name)
    m.use_nodes = True
    b = m.node_tree.nodes["Principled BSDF"]
    b.inputs["Base Color"].default_value = rgba
    b.inputs["Roughness"].default_value = 0.62
    return m


def clear_scene():
    bpy.ops.wm.read_factory_settings(use_empty=True)


def build_armature():
    data = bpy.data.armatures.new("Mannequin")
    arm = bpy.data.objects.new("Mannequin", data)
    bpy.context.collection.objects.link(arm)
    bpy.context.view_layer.objects.active = arm
    bpy.ops.object.mode_set(mode="EDIT")
    for name, head, tail, parent in BONES:
        b = data.edit_bones.new(name)
        b.head = head
        b.tail = tail
        if parent:
            b.parent = data.edit_bones[parent]
    bpy.ops.object.mode_set(mode="OBJECT")
    return arm


def unit_sphere():
    bpy.ops.mesh.primitive_uv_sphere_add(segments=24, ring_count=16, radius=1)
    ob = bpy.context.active_object
    bpy.ops.object.shade_smooth()
    return ob


def unit_cylinder_y():
    bpy.ops.mesh.primitive_cylinder_add(vertices=20, radius=1, depth=2)
    ob = bpy.context.active_object
    ob.rotation_euler = (math.radians(-90), 0, 0)
    bpy.ops.object.transform_apply(rotation=True)
    bpy.ops.object.shade_smooth()
    return ob


def build_visuals():
    mats = {
        "skin": new_mat("skin", SKIN),
        "shirt": new_mat("shirt", SHIRT),
        "shorts": new_mat("shorts", SHORTS),
        "hair": new_mat("hair", HAIR),
        "db": new_mat("db", DBCOL),
        "mat": new_mat("matcol", (0.80, 0.83, 0.86, 1)),
        "floor": new_mat("floor", (0.90, 0.91, 0.93, 1)),
    }
    parts = []

    def add(kind, name, bone, scale, color):
        ob = unit_sphere() if kind == "sphere" else unit_cylinder_y()
        ob.name = name
        ob.data.materials.append(mats[color])
        parts.append((ob, bone, scale))
        return ob

    # torso volumes, scale in bone space: X width, Y along bone, Z depth
    add("sphere", "pelvis", "hips", (0.15, 0.12, 0.11), "shorts")
    add("sphere", "waist", "spine", (0.125, 0.13, 0.10), "shirt")
    add("sphere", "ribs", "chest", (0.175, 0.16, 0.12), "shirt")
    add("cylinder", "neck", "neck", (0.042, None, 0.042), "skin")
    add("sphere", "head", "head", (0.105, 0.115, 0.12), "skin")
    add("sphere", "hair", "head", (0.108, 0.07, 0.11), "hair")
    add("sphere", "nose", "head", (0.018, 0.018, 0.018), "skin")
    for side in ("L", "R"):
        add("cylinder", f"clav_{side}", f"clavicle.{side}", (0.035, None, 0.035), "shirt")
        add("sphere", f"sho_{side}", f"upper_arm.{side}", (0.055, 0.055, 0.055), "shirt")
        add("cylinder", f"ua_{side}", f"upper_arm.{side}", (0.042, None, 0.042), "skin")
        add("sphere", f"el_{side}", f"forearm.{side}", (0.04, 0.04, 0.04), "skin")
        add("cylinder", f"fa_{side}", f"forearm.{side}", (0.036, None, 0.036), "skin")
        add("sphere", f"ha_{side}", f"hand.{side}", (0.034, 0.055, 0.02), "skin")
        add("sphere", f"hipj_{side}", f"thigh.{side}", (0.07, 0.07, 0.07), "shorts")
        add("cylinder", f"th_{side}", f"thigh.{side}", (0.062, None, 0.062), "shorts")
        add("sphere", f"kn_{side}", f"shin.{side}", (0.05, 0.05, 0.05), "skin")
        add("cylinder", f"sh_{side}", f"shin.{side}", (0.046, None, 0.046), "skin")
        add("sphere", f"ft_{side}", f"foot.{side}", (0.038, None, 0.028), "skin")
    dbs = []
    for side in ("L", "R"):
        bar = unit_cylinder_y()
        bar.name = f"db_{side}"
        bar.data.materials.append(mats["db"])
        bar.hide_render = True
        dbs.append((bar, side))
    bpy.ops.mesh.primitive_plane_add(size=1)
    matp = bpy.context.active_object
    matp.name = "Mat"
    matp.data.materials.append(mats["mat"])
    bpy.ops.mesh.primitive_plane_add(size=12, location=(0, 0, 0))
    floor = bpy.context.active_object
    floor.name = "Floor"
    floor.data.materials.append(mats["floor"])
    return parts, dbs, matp


def bone_len(arm, name):
    pb = arm.pose.bones[name]
    return (pb.tail - pb.head).length


# Joint spheres sit on the bone head. Everything else is centered on the bone,
# because a unit mesh is centered on its origin and the bone origin is the head.
AT_HEAD = ("sho_", "el_", "kn_", "hipj_")


def place_parts(arm, parts, dbs, props, nose_shift=True):
    bpy.context.view_layer.update()
    for ob, bone, scale in parts:
        pb = arm.pose.bones[bone]
        length = (pb.tail - pb.head).length
        sx, sy, sz = scale
        if sy is None:
            sy = length / 2.0
        mw = arm.matrix_world @ pb.matrix
        at_head = ob.name.startswith(AT_HEAD)
        if ob.name == "hair":
            local = Matrix.Translation((0, length * 0.72, 0.0))
        elif ob.name == "nose":
            local = Matrix.Translation((0, length * 0.42, 0.105))
        elif at_head:
            local = Matrix.Identity(4)
        else:
            local = Matrix.Translation((0, length * 0.5, 0))
        ob.matrix_world = mw @ local @ Matrix.Diagonal((sx, sy, sz, 1))
    propset = set(props or ())
    for bar, side in dbs:
        show = side in propset or f"hand.{side}" in propset or side.lower() in propset
        # props are "L" / "R"
        show = side in propset
        bar.hide_render = not show
        if not show:
            continue
        pb = arm.pose.bones[f"hand.{side}"]
        mw = arm.matrix_world @ pb.matrix
        # across the palm (bone local X), at the hand
        bar.matrix_world = mw @ Matrix.Translation((0, 0.06, 0)) @ Matrix.Diagonal((0.045, 0.09, 0.045, 1))


# Our pose names -> MakeHuman game-engine rig.
BONE_MAP = {
    "hips": "pelvis",
    "spine": "spine_01",
    "chest": "spine_03",
    "neck": "neck_01",
    "head": "head",
    "clavicle.L": "clavicle_l",
    "clavicle.R": "clavicle_r",
    "upper_arm.L": "upperarm_l",
    "upper_arm.R": "upperarm_r",
    "forearm.L": "lowerarm_l",
    "forearm.R": "lowerarm_r",
    "hand.L": "hand_l",
    "hand.R": "hand_r",
    "thigh.L": "thigh_l",
    "thigh.R": "thigh_r",
    "shin.L": "calf_l",
    "shin.R": "calf_r",
    "foot.L": "foot_l",
    "foot.R": "foot_r",
}
AIM_ORDER = [
    "hips", "spine", "chest", "neck", "head",
    "clavicle.L", "clavicle.R",
    "upper_arm.L", "forearm.L", "hand.L",
    "upper_arm.R", "forearm.R", "hand.R",
    "thigh.L", "shin.L", "foot.L",
    "thigh.R", "shin.R", "foot.R",
]


def hand_up_hint(direction, side, palm_normal):
    """Roll the hand so the palm normal matches. Left palm is local +X, right palm is local -X."""
    y = Vector(direction).normalized()
    sign = 1 if side == "L" else -1
    local_x = Vector(palm_normal) * sign
    local_x = local_x - y * local_x.dot(y)
    if local_x.length < 0.2:
        fallback = Vector(B) if abs(y.dot(B)) < 0.8 else Vector(LFT)
        local_x = fallback * sign
        local_x = local_x - y * local_x.dot(y)
    local_x.normalize()
    return local_x.cross(y).normalized()


def default_palm(direction, side):
    d = Vector(direction).normalized()
    if abs(d.z) > 0.72:
        # Arm is vertical: palm faces the thigh, not the camera.
        return RGT if side == "L" else LFT
    return D


def relax_fingers(arm, grip=False):
    # Local +X bends the finger toward the palm in the rest T-pose.
    curl = {"01": 68, "02": 76, "03": 52} if grip else {"01": 10, "02": 14, "03": 8}
    thumb = (("01", 46), ("02", 34), ("03", 24)) if grip else (("01", 12), ("02", 10), ("03", 8))
    for side in ("l", "r"):
        for finger in ("index", "middle", "ring", "pinky"):
            for joint, deg in curl.items():
                name = f"{finger}_{joint}_{side}"
                pb = arm.pose.bones.get(name)
                if pb is None:
                    continue
                pb.rotation_mode = "XYZ"
                pb.rotation_euler.x += math.radians(deg)
        for joint, deg in thumb:
            pb = arm.pose.bones.get(f"thumb_{joint}_{side}")
            if pb is None:
                continue
            pb.rotation_mode = "XYZ"
            pb.rotation_euler.x += math.radians(deg)


def apply_pose(arm, pose):
    """Local bends off the rest T-pose. Absolute aim() was tearing this rig."""
    bpy.context.view_layer.objects.active = arm
    if bpy.context.object.mode != "POSE":
        bpy.ops.object.mode_set(mode="POSE")
    for pb in arm.pose.bones:
        pb.rotation_mode = "XYZ"
        pb.rotation_euler = (0, 0, 0)
        pb.location = (0, 0, 0)
    arm.rotation_mode = "XYZ"
    rx, ry, rz = pose.get("root", (0, 0, 0))
    arm.rotation_euler = (math.radians(rx), math.radians(ry), math.radians(rz))
    arm.location = Vector(pose.get("loc", (0, 0, 0)))
    for name, eul in (pose.get("bones") or {}).items():
        pb = arm.pose.bones.get(name)
        if pb is None:
            continue
        pb.rotation_euler = tuple(math.radians(float(v)) for v in eul)
    if pose.get("fingers") != "straight":
        relax_fingers(arm, grip=bool(pose.get("props")))
    bpy.ops.object.mode_set(mode="OBJECT")
    bpy.context.view_layer.update()


_SKIP_BONES = ("index", "middle", "ring", "pinky", "thumb", "ball_")


def world_points(arm):
    pts = []
    for pb in arm.pose.bones:
        if any(s in pb.name for s in _SKIP_BONES):
            continue
        pts.append(arm.matrix_world @ pb.head)
        pts.append(arm.matrix_world @ pb.tail)
    return pts


def foot_bone(arm, side):
    name = BONE_MAP.get(f"foot.{side}", f"foot.{side}")
    return arm.pose.bones[name]


def body_mesh():
    return bpy.data.objects.get("SuperHero_Male") or bpy.data.objects.get("Human")


def _posed_vertices(ob):
    """Armature-deformed vertices. ob.data stays in the bind pose."""
    deps = bpy.context.evaluated_depsgraph_get()
    ev = ob.evaluated_get(deps)
    mesh = ev.to_mesh()
    mw = ev.matrix_world
    try:
        return [mw @ v.co for v in mesh.vertices]
    finally:
        ev.to_mesh_clear()


def lowest_vertex_z(ob):
    verts = _posed_vertices(ob)
    if not verts:
        return 0.0
    return min(v.z for v in verts)


def plant(arm, mode, gap=None):
    """Drop the mesh onto the floor. Extremity-only planting was burying the torso."""
    bpy.context.view_layer.update()
    if mode == "feet":
        toes = []
        for side in ("l", "r"):
            pb = arm.pose.bones[f"foot_{side}"]
            toes.append(arm.matrix_world @ pb.head)
        avg = sum(toes, Vector()) / len(toes)
        arm.location.x -= avg.x
        arm.location.y -= avg.y
    else:
        pts = character_points()
        if pts:
            xs = [p.x for p in pts]
            ys = [p.y for p in pts]
            arm.location.x -= (min(xs) + max(xs)) * 0.5
            arm.location.y -= (min(ys) + max(ys)) * 0.5
    bpy.context.view_layer.update()
    body = body_mesh()
    zmin = lowest_vertex_z(body) if body else min(p.z for p in world_points(arm))
    if gap is None:
        gap = 0.62 if mode == "air" else 0.006
    arm.location.z -= zmin - gap
    bpy.context.view_layer.update()
    if os.environ.get("POSE_DEBUG"):
        z2 = lowest_vertex_z(body) if body else 0
        bz = min(p.z for p in world_points(arm))
        print(f"PLANT {mode} zmin {zmin:.3f} after {z2:.3f} bones {bz:.3f} loc {arm.location.z:.3f}", flush=True)


def extremity_z(arm):
    names = (
        "hand_l", "hand_r", "lowerarm_l", "lowerarm_r",
        "foot_l", "foot_r", "calf_l", "calf_r",
    )
    zs = []
    for name in names:
        pb = arm.pose.bones.get(name)
        if pb is None:
            continue
        zs.append((arm.matrix_world @ pb.head).z)
        zs.append((arm.matrix_world @ pb.tail).z)
    if not zs:
        return body_low_z()
    return min(zs)


def body_low_z():
    ob = bpy.data.objects.get("Human")
    if ob is None:
        return min(p.z for p in world_points(bpy.data.objects["Human.rig"]))
    deps = bpy.context.evaluated_depsgraph_get()
    ev = ob.evaluated_get(deps)
    mw = ev.matrix_world
    verts = ev.data.vertices
    step = max(1, len(verts) // 2500)
    zmin = 1e9
    for i in range(0, len(verts), step):
        zmin = min(zmin, (mw @ verts[i].co).z)
    return zmin


def hip_axes(arm):
    pb = arm.pose.bones["pelvis"] if "pelvis" in arm.pose.bones else arm.pose.bones["hips"]
    mw = (arm.matrix_world @ pb.matrix).to_3x3()
    bone_y = (mw @ Vector((0, 1, 0))).normalized()
    bone_z = (mw @ Vector((0, 0, 1))).normalized()
    # Standing: the pelvis bone leans, so its Z dips. The face is the horizontal part.
    if abs(bone_y.z) > 0.75:
        face = Vector((bone_z.x, bone_z.y, 0))
        if face.length < 1e-6:
            face = Vector(F)
        face.normalize()
        up = Vector(U)
    else:
        face = bone_z.normalized()
        up = bone_y
    right = face.cross(up)
    if right.length < 1e-6:
        right = Vector((1, 0, 0))
    right.normalize()
    return face, up, right


def look_at(obj, target):
    direction = (Vector(target) - obj.location).normalized()
    obj.rotation_mode = "QUATERNION"
    obj.rotation_quaternion = direction.to_track_quat("-Z", "Y")


def view_offset(view, face, up, right):
    face = Vector(face).normalized()
    up = Vector(up).normalized()
    right = Vector(right).normalized()
    world_up = Vector((0, 0, 1))
    view = {
        "профиль": "profile",
        "анфас": "front",
        "три четверти": "threeq",
        "сзади-сверху": "back_high",
        "сверху-сбоку": "side_high",
        "спереди-сверху": "front_high",
    }.get(view, view)
    if view == "profile":
        offset = Vector(right)
    elif view == "front":
        offset = Vector(face)
    elif view == "threeq":
        offset = (face + right).normalized()
    elif view == "back_high":
        offset = (-face * math.cos(math.radians(40)) + world_up * math.sin(math.radians(40))).normalized()
    elif view == "side_high":
        offset = (right * math.cos(math.radians(35)) + world_up * math.sin(math.radians(35))).normalized()
    elif view == "front_high":
        offset = (face * math.cos(math.radians(25)) + world_up * math.sin(math.radians(25))).normalized()
    else:
        offset = Vector(face)
    # On the floor the chest points into the ground. Lift the camera so the body stays in frame.
    if abs(up.z) < 0.5 and offset.z < 0.15:
        offset = Vector((offset.x, offset.y, 0.55))
        if offset.length < 1e-4:
            offset = Vector((right.x, right.y, 0.55))
        offset.normalize()
    return offset


def camera_for(view, face, up, right, center, radius):
    face = Vector(face).normalized()
    up = Vector(up).normalized()
    right = Vector(right).normalized()
    world_up = Vector((0, 0, 1))
    offset = view_offset(view, face, up, right)
    dist = max(radius * 2.4, 2.2)
    cam = bpy.data.objects["Cam"]
    cam.location = Vector(center) + offset * dist
    # ortho scale from projected span is computed by the caller via radius
    look_at(cam, center)
    # force camera up toward world up by rebuilding the look with a track
    direction = (Vector(center) - cam.location).normalized()
    cam.rotation_quaternion = direction.to_track_quat("-Z", "Y")
    cam.data.type = "PERSP"
    cam.data.lens = 70
    cam.data.sensor_width = 36
    bpy.context.view_layer.update()
    half = max(radius, 0.35)
    dist = half / math.tan(cam.data.angle * 0.5) * 1.08
    cam.location = Vector(center) + offset.normalized() * dist
    look = Vector(center)
    if look.z < 0.1:
        look.z = 0.22
    if cam.location.z < 0.28:
        cam.location.z = 0.42
    direction = (look - cam.location).normalized()
    cam.rotation_quaternion = direction.to_track_quat("-Z", "Y")
    return cam


def projected_radius(points, offset, center=None):
    offset = Vector(offset).normalized()
    xs3 = [p.x for p in points]
    ys3 = [p.y for p in points]
    zs3 = [p.z for p in points]
    if center is None:
        center = Vector((
            (min(xs3) + max(xs3)) * 0.5,
            (min(ys3) + max(ys3)) * 0.5,
            (min(zs3) + max(zs3)) * 0.5,
        ))
    else:
        center = Vector(center)
        center.z = (min(zs3) + max(zs3)) * 0.5
    # One head of air above the crown, a little floor under the feet.
    padded = list(points)
    padded.append(Vector((center.x, center.y, max(zs3) + 0.26)))
    padded.append(Vector((center.x, center.y, min(zs3) - 0.06)))
    forward = (-offset).normalized()
    world_up = Vector((0, 0, 1))
    right = forward.cross(world_up)
    if right.length < 1e-6:
        right = forward.cross(Vector((0, 1, 0)))
    right.normalize()
    up = right.cross(forward).normalized()
    xs, ys = [], []
    for p in padded:
        d = p - center
        xs.append(d.dot(right))
        ys.append(d.dot(up))
    span = max(max(xs) - min(xs), max(ys) - min(ys))
    return span * 0.64, center


def render_still(path):
    bpy.context.scene.render.filepath = path
    bpy.ops.render.render(write_still=True)


def photo_setup():
    sc = bpy.context.scene
    sc.render.engine = "CYCLES"
    sc.render.resolution_x = 800
    sc.render.resolution_y = 800
    sc.render.resolution_percentage = 100
    sc.render.film_transparent = False
    sc.render.image_settings.file_format = "PNG"
    sc.render.image_settings.color_mode = "RGB"
    sc.cycles.samples = 64
    sc.cycles.use_adaptive_sampling = True
    sc.cycles.adaptive_threshold = 0.015
    sc.cycles.use_denoising = True
    try:
        sc.cycles.denoiser = "OPENIMAGEDENOISE"
    except TypeError:
        pass
    prefs = bpy.context.preferences.addons["cycles"].preferences
    prefs.compute_device_type = "METAL"
    prefs.get_devices()
    for dev in prefs.devices:
        dev.use = True
    sc.cycles.device = "GPU"
    sc.view_settings.view_transform = "AgX"
    sc.view_settings.look = "AgX - Punchy"
    sc.view_settings.exposure = 0.05
    sc.view_settings.gamma = 1.0

    world = sc.world or bpy.data.worlds.new("PhotoWorld")
    sc.world = world
    world.use_nodes = True
    bg = world.node_tree.nodes["Background"]
    bg.inputs[0].default_value = (0.93, 0.94, 0.95, 1)
    bg.inputs[1].default_value = 0.55

    def area(name, loc, energy, size, color):
        old = bpy.data.objects.get(name)
        if old:
            bpy.data.objects.remove(old, do_unlink=True)
        data = bpy.data.lights.new(name, "AREA")
        ob = bpy.data.objects.new(name, data)
        bpy.context.collection.objects.link(ob)
        ob.location = loc
        direction = (-Vector(loc)).normalized()
        ob.rotation_mode = "QUATERNION"
        ob.rotation_quaternion = direction.to_track_quat("-Z", "Y")
        data.energy = energy
        data.size = size
        data.color = color
        data.shape = "DISK"
        return ob

    area("Key", (1.6, -2.4, 2.6), 280, 2.2, (1.0, 0.96, 0.92))
    area("Fill", (-2.2, -1.6, 1.8), 70, 3.0, (0.90, 0.94, 1.0))
    area("Rim", (0.4, 2.6, 2.4), 140, 1.6, (1.0, 0.97, 0.94))

    if "Cam" not in bpy.data.objects:
        cam = bpy.data.objects.new("Cam", bpy.data.cameras.new("Cam"))
        bpy.context.collection.objects.link(cam)
    sc.camera = bpy.data.objects["Cam"]

    floor = bpy.data.objects.get("Floor")
    if floor is None:
        bpy.ops.mesh.primitive_plane_add(size=14, location=(0, 0, 0))
        floor = bpy.context.active_object
        floor.name = "Floor"
        mat = new_mat("floor", (0.86, 0.87, 0.88, 1))
        mat.node_tree.nodes["Principled BSDF"].inputs["Roughness"].default_value = 0.72
        floor.data.materials.append(mat)
    floor.hide_render = False

    for ob in bpy.data.objects:
        if ob.type != "MESH":
            continue
        for mod in ob.modifiers:
            if mod.type == "SUBSURF":
                mod.render_levels = 1
                mod.levels = 1
        low = ob.name.lower()
        if "polo" in low or "trouser_short" in low or "swim_short" in low:
            ob.hide_render = True
    paint_shorts()


def paint_shorts():
    """Shorts are faces of the body itself, so they bend with the legs and do not tear."""
    ob = bpy.data.objects.get("Human")
    if ob is None:
        return
    mesh = ob.data
    mat = bpy.data.materials.get("shorts_cloth") or new_mat("shorts_cloth", (0.11, 0.14, 0.18, 1))
    mat.node_tree.nodes["Principled BSDF"].inputs["Roughness"].default_value = 0.58
    if mat.name not in mesh.materials:
        mesh.materials.append(mat)
    slot = mesh.materials.find(mat.name)
    sks = mesh.shape_keys
    basis = sks.reference_key if sks else None

    def co(i):
        p = mesh.vertices[i].co.copy()
        if not sks:
            return p
        for kb in sks.key_blocks:
            if kb == basis or kb.value == 0:
                continue
            p += (kb.data[i].co - basis.data[i].co) * kb.value
        return p

    for poly in mesh.polygons:
        c = sum((co(i) for i in poly.vertices), Vector()) / len(poly.vertices)
        if 0.72 < c.z < 1.12 and abs(c.x) < 0.22:
            poly.material_index = slot


def ensure_mat():
    ob = bpy.data.objects.get("Mat")
    if ob is None:
        bpy.ops.mesh.primitive_plane_add(size=1, location=(0, 0, 0.004))
        ob = bpy.context.active_object
        ob.name = "Mat"
        mat = new_mat("matcol", (0.45, 0.52, 0.58, 1))
        mat.node_tree.nodes["Principled BSDF"].inputs["Roughness"].default_value = 0.8
        ob.data.materials.append(mat)
    return ob


def ensure_dbs():
    made = []
    for side, bone in (("L", "hand_l"), ("R", "hand_r")):
        name = f"db_{side}"
        ob = bpy.data.objects.get(name)
        if ob is None:
            bpy.ops.mesh.primitive_cylinder_add(vertices=24, radius=1, depth=2)
            ob = bpy.context.active_object
            ob.name = name
            mat = new_mat(name + "_m", (0.22, 0.23, 0.25, 1))
            bsdf = mat.node_tree.nodes["Principled BSDF"]
            bsdf.inputs["Roughness"].default_value = 0.28
            bsdf.inputs["Metallic"].default_value = 0.85
            ob.data.materials.append(mat)
        ob.hide_render = True
        made.append((ob, bone))
    return made


def place_dbs(arm, dbs, props):
    bpy.context.view_layer.update()
    want = set(props or ())
    for ob, bone in dbs:
        side = "L" if bone.endswith("_l") else "R"
        ob.hide_render = side not in want
        if ob.hide_render:
            continue
        pb = arm.pose.bones[bone]
        mw = arm.matrix_world @ pb.matrix
        # Cylinder axis is local Z, across the palm. Palm normal is +X left, -X right.
        sign = 1.0 if side == "L" else -1.0
        ob.matrix_world = mw @ Matrix.Translation((sign * 0.020, 0.034, 0.0)) @ Matrix.Diagonal((0.015, 0.015, 0.055, 1))


def setup_scene():
    clear_scene()
    sc = bpy.context.scene
    sc.render.engine = "BLENDER_EEVEE_NEXT"
    sc.render.resolution_x = 640
    sc.render.resolution_y = 640
    sc.render.film_transparent = False
    sc.render.image_settings.file_format = "PNG"
    if hasattr(sc, "eevee"):
        if hasattr(sc.eevee, "taa_render_samples"):
            sc.eevee.taa_render_samples = 16
    world = bpy.data.worlds.new("W")
    sc.world = world
    world.use_nodes = True
    world.node_tree.nodes["Background"].inputs[0].default_value = (1, 1, 1, 1)
    world.node_tree.nodes["Background"].inputs[1].default_value = 1.0
    sun_data = bpy.data.lights.new("Key", "SUN")
    sun = bpy.data.objects.new("Key", sun_data)
    bpy.context.collection.objects.link(sun)
    sun.rotation_euler = (math.radians(42), math.radians(-8), math.radians(28))
    sun.data.energy = 3.2
    sun.data.angle = math.radians(18)
    fill_data = bpy.data.lights.new("Fill", "AREA")
    fill = bpy.data.objects.new("Fill", fill_data)
    bpy.context.collection.objects.link(fill)
    fill.location = (-1.5, -2.0, 2.2)
    fill.data.energy = 250
    fill.data.size = 3
    cam_data = bpy.data.cameras.new("Cam")
    cam = bpy.data.objects.new("Cam", cam_data)
    bpy.context.collection.objects.link(cam)
    sc.camera = cam


# ---------- pose helpers ----------
# Degrees are local bends off the rest T-pose.
# Arm: flex forward, down from horizontal toward the hip, elbow fold, twist around the bone.
# Leg: hip flexion forward, knee flexion, abduct out. Foot stays flat unless foot= is set.

PRONE = (90, 0, 0)
SUPINE = (-90, 0, 0)
RIGHT_SIDE = (0, -90, 0)


def add(*parts):
    out = {}
    for part in parts:
        if part:
            out.update(part)
    return out


def arm(side, flex=0, down=0, elbow=0, twist=0):
    s = "l" if side == "L" else "r"
    z = -down if side == "L" else down
    return {
        f"upperarm_{s}": (flex, twist, z),
        f"lowerarm_{s}": (elbow, 0, 0),
    }


def arms(flex=0, down=0, elbow=0, twist=0):
    return add(
        arm("L", flex, down, elbow, twist),
        arm("R", flex, down, elbow, twist),
    )


def leg(side, hip=0, knee=0, abduct=0, foot=None, twist=0, tilt=0):
    s = "l" if side == "L" else "r"
    abd = -abduct if side == "L" else abduct
    tw = -twist if side == "L" else twist
    # tilt is the pelvis X rotation the legs must cancel so the feet stay down.
    if foot is None:
        foot = hip - knee - tilt
    return {
        f"thigh_{s}": (-hip, tw, abd),
        f"calf_{s}": (knee, 0, 0),
        f"foot_{s}": (foot, 0, 0),
    }


def legs(hip=0, knee=0, abduct=0, foot=None, twist=0, tilt=0):
    return add(
        leg("L", hip, knee, abduct, foot, twist, tilt),
        leg("R", hip, knee, abduct, foot, twist, tilt),
    )


def spine(flex=0, twist=0, pelvis=0):
    return {
        "pelvis": (-pelvis, 0, 0),
        "spine_01": (flex * 0.28, twist * 0.22, 0),
        "spine_02": (flex * 0.40, twist * 0.42, 0),
        "spine_03": (flex * 0.32, twist * 0.36, 0),
    }


def head(nod=0, turn=0):
    return {
        "neck_01": (nod * 0.4, turn * 0.35, 0),
        "Head": (nod * 0.6, turn * 0.65, 0),
    }


def shrug(lift=0, back=0):
    return {
        "clavicle_l": (-lift, 0, back),
        "clavicle_r": (-lift, 0, -back),
    }


def stand():
    return add(arms(down=76, elbow=10), legs(hip=1, knee=3, abduct=6))


def P(bones, root=(0, 0, 0), plant="feet", props=(), fingers=None, gap=None):
    pose = {
        "bones": bones,
        "root": tuple(root),
        "plant": plant,
        "props": tuple(props),
    }
    if fingers:
        pose["fingers"] = fingers
    if gap is not None:
        pose["gap"] = gap
    return pose


def shot(sid, view, frames, mat=False, props=(), copies=()):
    return {"id": sid, "view": view, "frames": frames, "mat": mat, "props": props, "copies": copies}


def squat_pose(hip, knee):
    # Pelvis bone X is the negative of this argument. A large positive value tips the chest back.
    pel = 4
    return add(
        legs(hip=hip, knee=knee, abduct=9, tilt=-pel),
        spine(flex=8, pelvis=pel),
        arms(flex=68, down=14, elbow=14),
        head(nod=6),
    )


def hinge_pose(arms_down=36, arms_flex=0, elbow=10):
    # Pelvis folds the chest forward. Thighs cancel that swing so the legs stay planted.
    tilt = 46
    return add(
        spine(pelvis=-tilt),
        legs(hip=tilt + 8, knee=20, abduct=4, tilt=tilt),
        arms(flex=arms_flex, down=arms_down, elbow=elbow),
        head(nod=16),
    )


def march(lift, hip=86, knee=90, heel=0):
    other = "L" if lift == "R" else "R"
    support_foot = -heel if heel else None
    return add(
        leg(lift, hip=hip, knee=knee),
        leg(other, hip=4, knee=6, abduct=3, foot=support_foot),
        arm(other, flex=58, down=28, elbow=52),
        arm(lift, flex=-22, down=36, elbow=38),
        spine(flex=6),
    )


def chest_turn(sign):
    return add(
        legs(hip=2, knee=4, abduct=6),
        spine(twist=sign * 36),
        head(turn=sign * 30),
        arms(flex=58, down=32, elbow=92),
    )


def flat_hands():
    # Palms on the floor. Right side mirrors Y and Z; the same euler points the right thumb down.
    return {"hand_l": (20, -60, 60), "hand_r": (20, 60, -60)}


def high_plank(elbow=0, flex=75, shift=0, stagger=0, narrow=0, root=72):
    bones = add(
        legs(hip=0, knee=2, foot=70),
        arm("L", flex=flex + shift + stagger, elbow=elbow),
        arm("R", flex=flex + shift - stagger, elbow=elbow),
        spine(flex=-2),
        head(nod=14),
        flat_hands(),
    )
    if narrow:
        bones["clavicle_l"] = (0, 0, -narrow)
        bones["clavicle_r"] = (0, 0, narrow)
    return P(bones, root=(root, 0, 0), plant="body", fingers="straight")


def forearm_plank():
    return P(add(
        legs(hip=0, knee=2, foot=75),
        arms(flex=74, elbow=92),
        head(nod=12),
        flat_hands(),
    ), root=(76, 0, 0), plant="body", fingers="straight")


def prone_body(chest=-8, **arm_kw):
    return add(legs(hip=0, knee=2, foot=0), spine(flex=chest), head(nod=10 + chest * 0.15), arms(**arm_kw))


def press(elbow, left):
    """Right arm does the curl. Left palm stays on that forearm."""
    return P(add(
        legs(1, 3, 6),
        arm("R", flex=6, down=74, elbow=elbow, twist=-80),
        arm("L", **left),
    ))


def side_plank(sag=0):
    bones = add(
        legs(hip=1, knee=2, abduct=1),
        arm("R", flex=0, down=0, elbow=0),
        arm("L", flex=10, down=85, elbow=6),
        spine(flex=-2 + sag, pelvis=-sag * 0.35),
        head(turn=4),
    )
    # Elbow under the shoulder, forearm flat on the floor toward the camera.
    bones["lowerarm_r"] = (90, 0, 0)
    return P(bones, root=RIGHT_SIDE, plant="body", fingers="straight")


def skate(support):
    free = "R" if support == "L" else "L"
    sign = -1 if support == "L" else 1
    return P(add(
        leg(support, hip=8, knee=24, abduct=3),
        leg(free, hip=-80, knee=12, abduct=-24, foot=10),
        spine(flex=10, twist=sign * 12),
        arm(free, flex=56, down=16, elbow=24),
        arm(support, flex=-16, down=38, elbow=18),
    ))


def supine_body(**leg_kw):
    return add(legs(**leg_kw), arms(down=78, elbow=6))


SHOTS = [
    shot("W01", "threeq", [
        P(add(stand(), shrug(26))),
        P(add(arms(down=74, elbow=8, flex=-14), legs(hip=1, knee=3, abduct=6), shrug(back=16), spine(flex=-8))),
    ]),
    shot("W02", "threeq", [
        P(add(legs(hip=1, knee=3, abduct=6), arms(elbow=2))),
        P(add(legs(hip=1, knee=3, abduct=6), arms(flex=18, elbow=2))),
    ]),
    shot("W03", "front", [
        P(chest_turn(-1)),
        P(chest_turn(1)),
    ]),
    shot("W09", "profile", [
        P(add(legs(1, 3, 6), arm("R", flex=86, elbow=8), arm("L", flex=-38, down=18, elbow=8))),
        P(add(legs(1, 3, 6), arm("L", flex=86, elbow=8), arm("R", flex=-38, down=18, elbow=8))),
    ]),
    shot("J02", "profile", [
        P(march("R")),
        P(march("L")),
    ]),
    shot("J07", "profile", [
        P(march("R", hip=108, knee=62, heel=32)),
        P(march("L", hip=108, knee=62, heel=32)),
    ]),
    shot("J03", "front", [
        P(add(arms(down=78, elbow=8), legs(hip=0, knee=2, abduct=0))),
        P(add(arms(down=-102, flex=4, elbow=8), legs(hip=8, knee=10, abduct=24))),
    ]),
    shot("J06", "front", [
        skate("L"),
        skate("R"),
    ]),
    shot("L07", "profile", [
        P(stand()),
        P(add(arms(down=76, elbow=8), legs(hip=1, knee=2, abduct=6, foot=-34))),
    ]),
    shot("A01", "threeq", [
        press(16, dict(flex=100, down=0, elbow=60, twist=85)),
        press(120, dict(flex=110, down=-15, elbow=84, twist=75)),
    ]),
    shot("A04", "threeq", [
        press(120, dict(flex=110, down=-15, elbow=84, twist=75)),
        press(40, dict(flex=100, down=0, elbow=54, twist=75)),
    ]),
    shot("A02", "threeq", [
        (hold_a02 := press(90, dict(flex=110, down=-10, elbow=78, twist=75))),
        hold_a02,
    ]),
    shot("A03", "front", [
        P(add(legs(1, 3, 6), arms(flex=52, down=22, elbow=108))),
        P(add(legs(1, 3, 6), arms(flex=28, down=-6, elbow=62))),
    ]),
    shot("D01", "front", [
        P(add(legs(1, 3, 6), arm("L", down=78, elbow=6), arm("R", flex=-6, down=83, elbow=92, twist=-25)), props=("R",)),
        P(add(legs(1, 3, 6), arm("L", down=78, elbow=6), arm("R", flex=-8, down=90, elbow=92, twist=22)), props=("R",)),
    ], props=("R",)),
    shot("D04", "front", [
        (hold_d04 := P(add(legs(1, 3, 5), spine(flex=-6), arms(down=-84, flex=8, elbow=4)), props=("L", "R"))),
        hold_d04,
    ], props=("L", "R")),
    shot("D05", "front", [
        (hold_d05 := P(add(legs(1, 3, 6), arms(down=76, elbow=6)), props=("R",))),
        hold_d05,
    ], props=("R",)),
    shot("L01", "profile", [P(stand()), P(squat_pose(38, 50))], copies=("W05", "L02")),
    shot("L10", "profile", [P(stand()), P(squat_pose(78, 102))], copies=("L12",)),
    shot("L11", "profile", [
        P(add(arms(down=76, elbow=8), legs(hip=0, knee=2, abduct=1))),
        P(add(
            leg("L", hip=58, knee=96, abduct=3),
            leg("R", hip=-16, knee=90, abduct=4, foot=-20),
            spine(flex=2, pelvis=6),
            arms(down=68, elbow=12),
        )),
    ]),
    shot("J05", "profile", [
        P(squat_pose(78, 102)),
        P(add(legs(hip=8, knee=12, abduct=4, foot=-24), arms(flex=-28, down=48, elbow=16), spine(flex=4)), plant="air"),
    ]),
    shot("D03", "profile", [
        P(add(legs(1, 3, 6), arms(flex=40, down=26, elbow=112), spine(flex=-4)), props=("L", "R")),
        P(add(squat_pose(76, 100), arms(flex=40, down=26, elbow=112)), props=("L", "R")),
    ], props=("L", "R")),
    shot("W04", "profile", [P(stand()), P(hinge_pose())], copies=("L03",)),
    shot("D02", "front_high", [
        P(hinge_pose(arms_down=78, arms_flex=4, elbow=8), props=("L", "R")),
        P(hinge_pose(arms_down=4, arms_flex=0, elbow=18), props=("L", "R")),
    ], props=("L", "R")),
    shot("W06", "profile", [high_plank(0), high_plank(0, shift=8)], mat=True),
    shot("P01", "profile", [high_plank(0), high_plank(100, flex=30, root=76)], mat=True, copies=("P03", "P04")),
    shot("P02", "threeq", [
        high_plank(0, narrow=18),
        high_plank(100, flex=30, root=76, narrow=14),
    ], mat=True),
    shot("P05", "threeq", [
        high_plank(0, stagger=18),
        high_plank(100, flex=30, root=76, stagger=16),
    ], mat=True),
    shot("P06", "profile", [
        P(add(legs(hip=55, knee=4, foot=-40), spine(pelvis=-18), arms(flex=85, elbow=2), head(nod=12), flat_hands()), root=(82, 0, 0), plant="body", fingers="straight"),
        P(add(legs(hip=55, knee=4, foot=-40), spine(pelvis=-16), arms(flex=55, elbow=80), head(nod=24), flat_hands()), root=(82, 0, 0), plant="body", fingers="straight"),
    ], mat=True),
    shot("P08", "profile", [
        P(add(legs(hip=0, knee=2, foot=78), arms(flex=62, elbow=100), spine(flex=10), head(nod=12), flat_hands()), root=(74, 0, 0), plant="body", fingers="straight"),
        P(add(legs(hip=0, knee=2, foot=65), arms(flex=48, elbow=62), spine(flex=-4), head(nod=12), flat_hands()), root=(76, 0, 0), plant="body", fingers="straight"),
    ], mat=True),
    shot("C01", "profile", [forearm_plank(), forearm_plank()], mat=True, copies=("C02",)),
    shot("C09", "front_high", [
        high_plank(0),
        P(add(
            legs(hip=0, knee=2, foot=70),
            arm("L", flex=75, elbow=0),
            arm("R", flex=140, down=20, elbow=90, twist=40),
            spine(flex=-2),
            head(nod=10),
            {"hand_l": (20, -60, 60)},
        ), root=(72, 0, 0), plant="body", fingers="straight"),
    ], mat=True),
    shot("C10", "profile", [forearm_plank(), high_plank(0)], mat=True),
    shot("C07", "front", [
        (side := side_plank()),
        side,
    ], mat=True),
    shot("C08", "front", [
        side_plank(16),
        side,
    ], mat=True),
    shot("C03", "side_high", [
        P(add(leg("R", hip=74, knee=2), leg("L", hip=40, knee=2), arms(down=74, elbow=6)), root=SUPINE, plant="body"),
        P(add(leg("L", hip=74, knee=2), leg("R", hip=40, knee=2), arms(down=74, elbow=6)), root=SUPINE, plant="body"),
    ], mat=True),
    shot("C04", "profile", [
        P(add(leg("R", hip=16, knee=2), leg("L", hip=6, knee=2), arms(down=76, elbow=6)), root=SUPINE, plant="body"),
        P(add(leg("L", hip=16, knee=2), leg("R", hip=6, knee=2), arms(down=76, elbow=6)), root=SUPINE, plant="body"),
    ], mat=True),
    shot("C05", "profile", [
        P(add(legs(hip=94, knee=90), arms(down=76, elbow=6)), root=SUPINE, plant="body"),
        P(add(legs(hip=118, knee=96), arms(down=74, elbow=6), spine(flex=18, pelvis=12)), root=SUPINE, plant="body"),
    ], mat=True),
    shot("C06", "profile", [
        P(add(legs(hip=48, knee=62, abduct=8, foot=70), arms(down=78, elbow=6)), root=SUPINE, plant="body"),
        P(add(legs(hip=88, knee=82, foot=20), arms(flex=12, down=62, elbow=4), spine(flex=10), head(nod=16)), root=SUPINE, plant="body"),
    ], mat=True),
    shot("C11", "profile", [
        P(add(legs(hip=98, knee=92), arms(down=76, elbow=6)), root=SUPINE, plant="body"),
        P(add(leg("R", hip=102, knee=14), leg("L", hip=98, knee=92), arms(down=76, elbow=6)), root=SUPINE, plant="body"),
    ], mat=True),
    shot("C12", "side_high", [
        P(add(legs(hip=118, knee=105), arms(flex=72, elbow=4)), root=SUPINE, plant="body", fingers="straight"),
        P(add(
            leg("L", hip=30, knee=6),
            leg("R", hip=118, knee=105),
            arm("R", flex=8, down=-64, elbow=4),
            arm("L", flex=72, elbow=4),
        ), root=SUPINE, plant="body", fingers="straight"),
    ], mat=True),
    shot("C13", "side_high", [
        P(add(
            leg("L", hip=118, knee=102),
            leg("R", hip=26, knee=4),
            arms(flex=62, down=12, elbow=108),
            spine(twist=16),
            head(turn=-12, nod=8),
        ), root=SUPINE, plant="body"),
        P(add(
            leg("R", hip=118, knee=102),
            leg("L", hip=26, knee=4),
            arms(flex=62, down=12, elbow=108),
            spine(twist=-16),
            head(turn=12, nod=8),
        ), root=SUPINE, plant="body"),
    ], mat=True),
    shot("L06", "profile", [
        P(add(legs(hip=55, knee=100, abduct=6, foot=50), arms(down=76, elbow=6)), root=SUPINE, plant="body"),
        P(add(legs(hip=25, knee=90, abduct=6, foot=40), arms(flex=-40, down=60, elbow=20), head(nod=40), {"spine_01": (-42, 0, 0)}), root=SUPINE, plant="body"),
    ], mat=True),
    shot("L05", "profile", [
        P(add(legs(hip=25, knee=90, abduct=6, foot=40), arms(flex=-40, down=60, elbow=20), head(nod=40), {"spine_01": (-42, 0, 0)}), root=SUPINE, plant="body"),
        P(add(
            leg("L", hip=25, knee=90, abduct=6, foot=40),
            leg("R", hip=110, knee=78, foot=8),
            arms(flex=-40, down=60, elbow=20),
            head(nod=40),
            {"spine_01": (-42, 0, 0)},
        ), root=SUPINE, plant="body"),
    ], mat=True),
    shot("L04", "profile", [
        P(add(leg("R", hip=55, knee=100, abduct=6, foot=50), leg("L", hip=108, knee=2, foot=10), arms(down=76, elbow=6)), root=SUPINE, plant="body"),
        P(add(
            leg("R", hip=25, knee=90, abduct=6, foot=40),
            leg("L", hip=12, knee=2, foot=0),
            arms(flex=-40, down=60, elbow=20),
            head(nod=40),
            {"spine_01": (-42, 0, 0)},
        ), root=SUPINE, plant="body"),
    ], mat=True),
    shot("CD01", "side_high", [
        P(add(legs(hip=1, knee=2), arm("L", down=80, elbow=3, twist=35), arm("R", down=80, elbow=3, twist=-35)), root=SUPINE, plant="body", fingers="straight"),
        P(add(legs(hip=1, knee=2), arm("L", flex=-18, down=0, elbow=2, twist=20), arm("R", flex=-18, down=0, elbow=2, twist=-20)), root=SUPINE, plant="body", fingers="straight"),
    ], mat=True),
    shot("CD02", "profile", [
        P(add(
            leg("R", hip=58, knee=2),
            leg("L", hip=2, knee=2),
            arm("R", flex=48, down=16, elbow=28),
            arm("L", flex=40, down=22, elbow=24),
        ), root=SUPINE, plant="body"),
        P(add(
            leg("R", hip=88, knee=2),
            leg("L", hip=2, knee=2),
            arm("R", flex=62, down=8, elbow=22),
            arm("L", flex=55, down=12, elbow=18),
        ), root=SUPINE, plant="body"),
    ], mat=True),
    shot("CD03", "side_high", [
        P(add(legs(hip=92, knee=96), arms(down=12, elbow=3, twist=40)), root=SUPINE, plant="body", fingers="straight"),
        P(add(legs(hip=88, knee=96, twist=55), arms(down=12, elbow=3, twist=40), head(turn=-12)), root=SUPINE, plant="body", fingers="straight"),
    ], mat=True),
    shot("B01", "profile", [
        P(add(legs(hip=0, knee=0, foot=-50), spine(flex=-2), head(nod=8), arms(down=78, elbow=4)), root=(88, 0, 0), plant="body"),
        P(add(legs(hip=0, knee=0, foot=-50), spine(flex=-12), head(nod=6), arms(down=78, elbow=4)), root=(88, 0, 0), plant="body"),
    ], mat=True),
    shot("B02", "back_high", [
        P(prone_body(-8, down=96, elbow=0), root=PRONE, plant="body"),
        P(prone_body(-10, down=-100, flex=6, elbow=3), root=PRONE, plant="body"),
    ], mat=True),
    shot("B03", "back_high", [
        P(prone_body(-4, flex=22, down=-18, elbow=82), root=PRONE, plant="body"),
        P(prone_body(-16, flex=28, down=-28, elbow=74), root=PRONE, plant="body"),
    ], mat=True),
    shot("B04", "back_high", [
        P(prone_body(-4, down=-48, flex=12, elbow=4), root=PRONE, plant="body"),
        P(prone_body(-14, down=-52, flex=14, elbow=4), root=PRONE, plant="body"),
    ], mat=True),
    shot("B05", "back_high", [
        P(prone_body(-4, down=6, elbow=2), root=PRONE, plant="body"),
        P(prone_body(-14, down=-6, elbow=2), root=PRONE, plant="body"),
    ], mat=True),
    shot("B06", "back_high", [
        P(prone_body(-12, down=-80, flex=6, elbow=4), root=PRONE, plant="body"),
        P(prone_body(-14, flex=-8, down=36, elbow=102), root=PRONE, plant="body"),
    ], mat=True),
    shot("B10", "back_high", [
        P(prone_body(-12, down=-80, flex=6, elbow=4), root=PRONE, plant="body"),
        P(prone_body(-14, flex=26, down=-22, elbow=78), root=PRONE, plant="body"),
    ], mat=True),
    shot("CD04", "profile", [
        P(add(legs(hip=0, knee=2, foot=78), arms(flex=62, elbow=100), spine(flex=2), head(nod=16), flat_hands()), root=(74, 0, 0), plant="body", fingers="straight"),
        P(add(legs(hip=0, knee=2, foot=-45), arms(flex=52, elbow=84), spine(flex=-42), head(nod=2), flat_hands()), root=PRONE, plant="body", fingers="straight"),
    ], mat=True),
    shot("CD07", "profile", [
        P(add(legs(hip=70, knee=155, foot=80), arms(down=70, elbow=24), spine(flex=8)), plant="body"),
        P(add(
            legs(hip=70, knee=155, foot=80),
            spine(flex=95, pelvis=6),
            arms(flex=70, down=-6, elbow=4),
            head(nod=55),
            flat_hands(),
        ), plant="body", fingers="straight"),
    ], mat=True),
    shot("W07", "front", [
        P(stand()),
        P(add(legs(1, 3, 6), spine(flex=-4), arms(down=-96, flex=6, elbow=6), head(nod=-4))),
    ]),
    shot("W08", "front", [
        P(chest_turn(-1)),
        P(chest_turn(1)),
    ]),
    shot("B07", "back_high", [
        P(add(
            prone_body(-10, down=20, elbow=6),
            arm("L", down=-80, flex=16, elbow=10),
            arm("R", down=50, flex=-16, elbow=24),
            leg("R", hip=-55, knee=12, foot=20),
            leg("L", hip=4, knee=4, foot=0),
        ), root=PRONE, plant="body"),
        P(add(
            prone_body(-10, down=20, elbow=6),
            arm("R", down=-80, flex=16, elbow=10),
            arm("L", down=50, flex=-16, elbow=24),
            leg("L", hip=-55, knee=12, foot=20),
            leg("R", hip=4, knee=4, foot=0),
        ), root=PRONE, plant="body"),
    ], mat=True, copies=("B11",)),
    shot("B08", "profile", [
        (rev := P(add(
            legs(hip=6, knee=4, foot=-30),
            arms(flex=-70, down=10, elbow=4),
            spine(flex=-4),
            head(nod=-6),
        ), root=(-64, 0, 0), plant="body", fingers="straight")),
        rev,
    ], mat=True),
    shot("B09", "profile", [
        P(add(legs(hip=0, knee=0, foot=-50), spine(flex=-8), head(nod=6), arms(down=78, elbow=4)), root=(88, 0, 0), plant="body"),
        P(add(legs(hip=0, knee=0, foot=-50), spine(flex=-16), head(nod=4), arms(down=78, elbow=6)), root=(88, 0, 0), plant="body"),
    ], mat=True),
    shot("J01", "profile", [
        P(add(arms(down=76, elbow=8), legs(hip=1, knee=8, abduct=4, foot=-30))),
        P(add(legs(hip=4, knee=16, abduct=4, foot=-18), arms(down=68, elbow=16)), plant="air", gap=0.55),
    ]),
    shot("J04", "profile", [
        high_plank(0),
        P(add(
            leg("R", hip=110, knee=145),
            leg("L", hip=2, knee=4, foot=-8),
            arm("L", flex=75, elbow=0),
            arm("R", flex=75, elbow=0),
            spine(flex=-2),
            head(nod=14),
            flat_hands(),
        ), root=(72, 0, 0), plant="body", fingers="straight"),
    ], mat=True),
    shot("J08", "profile", [
        P(squat_pose(78, 102)),
        high_plank(0),
    ], mat=True),
    shot("J09", "profile", [
        P(add(
            leg("L", hip=58, knee=96, abduct=3),
            leg("R", hip=-16, knee=90, abduct=4, foot=-20),
            spine(flex=2, pelvis=6),
            arms(down=48, elbow=28),
        )),
        P(add(
            leg("R", hip=50, knee=80, abduct=8),
            leg("L", hip=-10, knee=70, abduct=6),
            arms(down=36, elbow=24),
            spine(flex=4),
        ), plant="air", gap=0.45),
    ]),
    shot("L09", "front", [
        side_plank(),
        P(add(
            leg("R", hip=1, knee=2, abduct=1),
            leg("L", hip=6, knee=4, abduct=28),
            arm("R", flex=0, down=0, elbow=0),
            arm("L", flex=10, down=85, elbow=6),
            spine(flex=-2),
            head(turn=4),
            {"lowerarm_r": (90, 0, 0)},
        ), root=RIGHT_SIDE, plant="body", fingers="straight"),
    ], mat=True),
    shot("CD05", "profile", [
        (breath := P(add(
            legs(hip=12, knee=28, abduct=8, foot=8),
            arms(down=78, elbow=8),
            head(nod=6),
        ), root=SUPINE, plant="body")),
        breath,
    ], mat=True),
]




def only_ids():
    argv = sys.argv
    if "--" in argv:
        argv = argv[argv.index("--") + 1 :]
    else:
        argv = []
    if "--only" in argv:
        raw = argv[argv.index("--only") + 1]
        return set(raw.split(","))
    return None


def character_points():
    skip_prefix = ("db_",)
    skip = {"Floor", "Mat"}
    pts = []
    deps = bpy.context.evaluated_depsgraph_get()
    for ob in bpy.data.objects:
        if ob.type != "MESH" or ob.name in skip or ob.name.startswith(skip_prefix):
            continue
        if ob.hide_render:
            continue
        ev = ob.evaluated_get(deps)
        mesh = ev.to_mesh()
        mw = ev.matrix_world
        try:
            step = max(1, len(mesh.vertices) // 500)
            for i in range(0, len(mesh.vertices), step):
                pts.append(mw @ mesh.vertices[i].co)
        finally:
            ev.to_mesh_clear()
    return pts


def render_shot(arm, dbs, mat_ob, shot):
    view = shot["view"]
    baked = []
    all_pts = []
    anchors = []
    hip = None
    for frame in shot["frames"]:
        apply_pose(arm, frame)
        plant(arm, frame.get("plant", "feet"), frame.get("gap"))
        bpy.context.view_layer.update()
        pts = world_points(arm) + character_points()
        all_pts.extend(pts)
        pel = arm.matrix_world @ arm.pose.bones["pelvis"].head
        anchors.append(pel)
        for side in ("l", "r"):
            anchors.append(arm.matrix_world @ arm.pose.bones[f"foot_{side}"].head)
        if frame.get("plant", "feet") == "feet":
            feet = sum((arm.matrix_world @ arm.pose.bones[f"foot_{s}"].head for s in ("l", "r")), Vector()) / 2
            back = (pel - feet).dot(B)
            if back > 0.16:
                print(f"LEAN {shot['id']} back {back:.2f}", flush=True)
        if hip is None:
            hip = hip_axes(arm)
            if view == "front_high":
                head_p = arm.matrix_world @ arm.pose.bones["Head"].head
                pel_p = arm.matrix_world @ arm.pose.bones["pelvis"].head
                along = Vector((head_p.x - pel_p.x, head_p.y - pel_p.y, 0))
                if along.length > 0.05:
                    face = along.normalized()
                    up = Vector(U)
                    right = face.cross(up).normalized()
                    hip = (face, up, right)
        baked.append(frame)
    face, up, right = hip
    offset = view_offset(view, face, up, right)
    radius, center = projected_radius(all_pts, offset)
    camera_for(view, face, up, right, center, radius)
    # lock camera: ortho already set
    mat_ob.hide_render = not shot.get("mat")
    if shot.get("mat"):
        xs = [p.x for p in all_pts]
        ys = [p.y for p in all_pts]
        mat_ob.location = ((min(xs) + max(xs)) / 2, (min(ys) + max(ys)) / 2, 0.005)
        mat_ob.scale = (max(0.7, (max(xs) - min(xs)) / 2 + 0.28), max(0.45, (max(ys) - min(ys)) / 2 + 0.22), 1)
        mat_ob.rotation_euler = (0, 0, 0)
    props = shot.get("props") or ()
    paths = []
    for i, frame in enumerate(baked, start=1):
        apply_pose(arm, frame)
        plant(arm, frame.get("plant", "feet"), frame.get("gap"))
        place_dbs(arm, dbs, frame.get("props", props))
        path = os.path.join(ROOT, f"{shot['id']}-{i}.png")
        render_still(path)
        paths.append(path)
        print("wrote", path, flush=True)
    for copy_id in shot.get("copies") or ():
        for i in (1, 2):
            src = os.path.join(ROOT, f"{shot['id']}-{i}.png")
            dst = os.path.join(ROOT, f"{copy_id}-{i}.png")
            shutil.copyfile(src, dst)
            print("copy", dst, flush=True)
    return paths


def load_character():
    clear_scene()
    body = os.path.join(CHAR_DIR, "Superhero_Male_FullBody.gltf")
    hair_path = os.path.join(CHAR_DIR, "Hair_SimpleParted.gltf")
    bpy.ops.import_scene.gltf(filepath=body)
    bpy.ops.import_scene.gltf(filepath=hair_path)
    keep = {"Eyebrows", "Eyes", "SuperHero_Male", "Hair_SimpleParted"}
    for ob in list(bpy.data.objects):
        if ob.type == "MESH" and ob.name not in keep:
            bpy.data.objects.remove(ob, do_unlink=True)
    arm = next(ob for ob in bpy.data.objects if ob.type == "ARMATURE")
    if arm.parent:
        world = arm.matrix_world.copy()
        arm.parent = None
        arm.matrix_world = world
    hair = bpy.data.objects["Hair_SimpleParted"]
    bpy.context.view_layer.objects.active = arm
    bpy.ops.object.mode_set(mode="POSE")
    arm.data.bones.active = arm.data.bones["Head"]
    bpy.ops.object.mode_set(mode="OBJECT")
    bpy.ops.object.select_all(action="DESELECT")
    hair.select_set(True)
    arm.select_set(True)
    bpy.context.view_layer.objects.active = arm
    bpy.ops.object.parent_set(type="BONE")
    for mat in bpy.data.materials:
        if "Hair" not in mat.name or not mat.node_tree:
            continue
        bsdf = next((n for n in mat.node_tree.nodes if n.type == "BSDF_PRINCIPLED"), None)
        if bsdf is None:
            continue
        for link in list(mat.node_tree.links):
            if link.to_node == bsdf and link.to_socket.name == "Base Color":
                mat.node_tree.links.remove(link)
        bsdf.inputs["Base Color"].default_value = (0.07, 0.04, 0.028, 1)
        bsdf.inputs["Roughness"].default_value = 0.62
    return arm


def setup_game():
    sc = bpy.context.scene
    sc.render.engine = "BLENDER_EEVEE_NEXT"
    sc.render.resolution_x = 760
    sc.render.resolution_y = 760
    sc.render.resolution_percentage = 100
    sc.render.film_transparent = False
    sc.render.image_settings.file_format = "PNG"
    sc.render.image_settings.color_mode = "RGB"
    sc.eevee.taa_render_samples = 16
    sc.view_settings.view_transform = "AgX"
    sc.view_settings.look = "AgX - Medium High Contrast"
    sc.view_settings.exposure = 0.0

    world = sc.world or bpy.data.worlds.new("Studio")
    sc.world = world
    world.use_nodes = True
    bg = world.node_tree.nodes["Background"]
    bg.inputs[0].default_value = (1, 1, 1, 1)
    bg.inputs[1].default_value = 1.0

    def area(name, loc, energy, size, color):
        data = bpy.data.lights.new(name, "AREA")
        ob = bpy.data.objects.new(name, data)
        bpy.context.collection.objects.link(ob)
        ob.location = loc
        ob.rotation_mode = "QUATERNION"
        ob.rotation_quaternion = (-Vector(loc)).to_track_quat("-Z", "Y")
        data.energy = energy
        data.size = size
        data.color = color
        return ob

    area("Key", (1.4, -2.4, 2.5), 180, 2.4, (1.0, 0.95, 0.90))
    area("Fill", (-2.0, -1.5, 1.6), 60, 3.0, (0.88, 0.92, 1.0))
    area("Rim", (0.3, 2.2, 2.3), 90, 1.6, (1.0, 0.97, 0.94))

    bpy.ops.mesh.primitive_plane_add(size=12, location=(0, 0, 0))
    floor = bpy.context.active_object
    floor.name = "Floor"
    mat = new_mat("floor", (0.86, 0.86, 0.87, 1))
    mat.node_tree.nodes["Principled BSDF"].inputs["Roughness"].default_value = 0.9
    floor.data.materials.append(mat)

    cam = bpy.data.objects.new("Cam", bpy.data.cameras.new("Cam"))
    bpy.context.collection.objects.link(cam)
    sc.camera = cam


def cli_arg(flag):
    argv = sys.argv
    if "--" in argv:
        argv = argv[argv.index("--") + 1 :]
    else:
        argv = []
    if flag in argv:
        return argv[argv.index(flag) + 1]
    return None


def main():
    global ROOT
    out = cli_arg("--out")
    if out:
        ROOT = os.path.abspath(out)
        os.makedirs(ROOT, exist_ok=True)
    arm = load_character()
    setup_game()
    mat_ob = ensure_mat()
    dbs = ensure_dbs()
    want = only_ids()
    for s in SHOTS:
        if want and s["id"] not in want:
            continue
        render_shot(arm, dbs, mat_ob, s)
    print("done", flush=True)


if __name__ == "__main__":
    main()
