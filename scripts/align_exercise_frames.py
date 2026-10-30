#!/usr/bin/env python3
"""Register generated three-frame strips without redrawing any part of the boy.

Requires Pillow and numpy. Each original is split at the gaps between its poses.
Floor exercises use the mat as a fixed camera reference. Standing exercises use
the head for horizontal registration, the planted foot for the ground line, and
the head's dimensions to correct small camera zoom differences. Crucially, we do
not normalize each full-body bounding box: that would enlarge a crouching pose.

All poses of an exercise (including an optional second strip) share one final
crop and scale. The JSON report records the exact affine transforms for review.
"""
import argparse
import json
from pathlib import Path

import numpy as np
from PIL import Image


FLOOR_IDS = {"it09", "it10", "it12", "it14", "it20", "it21", "it22", "it23", "it27", "it28",
             "it39", "it40", "it41", "it43", "it46", "it48", "it50", "it52", "it53"}
FOOT_ANCHOR_IDS = {"it05", "it26", "it44", "it45", "it47", "it51", "it54"}  # keep planted feet fixed as the torso moves
LINE_ANCHOR_IDS = {"it29", "it35"}  # keep the floor line fixed as either foot crosses it
AIR_FRAMES = {"it07": {0}, "it24": {1}}
# A stable camera must still show intentional lateral travel. Use a small,
# explicit excursion within the shared crop, not the generator's panel drift.
SIDE_TRAVEL = {"it07": [0, -.045, .045], "it15": [0, -.045, .045], "it24": [-.045, 0, .045],
               "it34": [0, -.045, .045], "it36": [0, -.045, .045], "it42": [0, .045, -.045]}
NO_ZOOM_IDS = {"it13", "it28", "it33", "it37", "it41", "it45", "it46", "it47", "it51", "it54"}  # changing orientation is not camera zoom
HEAD_ANCHOR_IDS = {"it12", "it13", "it46"}  # stationary head/shoulders while limbs move
WALL_ANCHOR_IDS = {"it51", "it54"}  # keep the wall fixed while the child changes support


def panel_cuts(image):
    """Find actual whitespace between poses; generated thirds can be unequal."""
    pixels = np.asarray(image.convert("RGB"), dtype=float)
    chroma = pixels.max(axis=2)-pixels.min(axis=2)
    occupied = (chroma > 25)*3 + (pixels.mean(axis=2) < 225)
    score = np.convolve(occupied.sum(axis=0), np.ones(7)/7, mode="same")
    cuts = [0]
    cell = image.width/3
    for index in (1, 2):
        left, right = round((index-.25)*cell), round((index+.25)*cell)
        local = score[left:right]
        candidates = np.where(local <= local.min()+.01)[0]
        # Pick the middle of the quietest contiguous gap, not its edge.
        runs = np.split(candidates, np.where(np.diff(candidates) > 1)[0]+1)
        run = max(runs, key=len)
        cuts.append(left+int(round(float(np.median(run)))))
    return cuts+[image.width]


def largest_component(mask):
    """Connected components at quarter resolution; no downloaded vision models."""
    small = np.asarray(Image.fromarray(mask).resize((mask.shape[1] // 4, mask.shape[0] // 4), Image.Resampling.NEAREST)).copy()
    seen = np.zeros(small.shape, dtype=bool)
    best = []
    height, width = small.shape
    for y, x in zip(*np.where(small & ~seen)):
        if seen[y, x]:
            continue
        seen[y, x] = True
        stack, component = [(int(y), int(x))], []
        while stack:
            cy, cx = stack.pop()
            component.append((cy, cx))
            for ny, nx in ((cy-1, cx), (cy+1, cx), (cy, cx-1), (cy, cx+1)):
                if 0 <= ny < height and 0 <= nx < width and small[ny, nx] and not seen[ny, nx]:
                    seen[ny, nx] = True
                    stack.append((ny, nx))
        if len(component) > len(best):
            best = component
    if not best:
        raise ValueError("No reference region found")
    points = np.asarray(best)
    y0, x0 = points.min(axis=0) * 4
    y1, x1 = (points.max(axis=0) + 1) * 4
    # Refine at native resolution so a one-pixel crop offset does not become a
    # four-pixel change in the detected head size after downsampling.
    left, top = max(0, int(x0)-4), max(0, int(y0)-4)
    local = mask[top:min(mask.shape[0], int(y1)+4), left:min(mask.shape[1], int(x1)+4)]
    ys, xs = np.where(local)
    return [int(xs.min())+left, int(ys.min())+top, int(xs.max())+left+1, int(ys.max())+top+1]


def bounds(mask):
    ys, xs = np.where(mask)
    if len(xs) < 100:
        raise ValueError("Subject segmentation failed")
    return [float(np.quantile(xs, .001)), float(np.quantile(ys, .001)),
            float(np.quantile(xs, .999)), float(np.quantile(ys, .999))]


def analyse(image, floor, foot_anchor=False, line_anchor=False, head_anchor=False, support_side=0, wall_anchor=False):
    pixels = np.asarray(image.convert("RGB"), dtype=float)
    r, g, b = pixels.transpose(2, 0, 1)
    chroma = pixels.max(axis=2) - pixels.min(axis=2)
    subject = (chroma > 25) & (pixels.min(axis=2) < 225)
    line_y = None
    if line_anchor:
        rows = np.where(((b > r+40) & (b > g+10)).sum(axis=1) > image.width*.55)[0]
        if not len(rows):
            raise ValueError("Blue floor line was not found")
        line_y = float(np.median(rows))
        subject[rows] = False
    box = bounds(subject)
    # The brown hair is separate from the blue clothes and light skin. Only its
    # largest component is used, so finger shadows cannot shift the reference.
    hair = largest_component((r > g * 1.12) & (g > b * 1.08) & (r < 175) & (b < 115) & (r > 25))
    hw, hh = hair[2] - hair[0], hair[3] - hair[1]
    if floor:
        # A body can divide the visible mat into disconnected islands. Use its
        # outer envelope, not one island whose size changes with the pose.
        mat = bounds((chroma < 24) & (pixels.mean(axis=2) < 218) & (pixels.mean(axis=2) > 75))
        if mat[2]-mat[0] < image.width*.3 or mat[3]-mat[1] < image.height*.04:
            raise ValueError("Mat reference is too small")
        anchor = [(mat[0]+mat[2])/2, mat[3]]
        box = [min(box[0], mat[0]), min(box[1], mat[1]), max(box[2], mat[2]), max(box[3], mat[3])]
        reference = [mat[2]-mat[0], mat[3]-mat[1]]
    else:
        anchor = [(hair[0]+hair[2])/2, box[3]]
        if foot_anchor:
            ys, xs = np.where(subject & (np.indices(subject.shape)[0] >= box[3]-20))
            anchor[0] = (float(xs.min())+float(xs.max()))/2
        reference = [hw, hh]
    if line_y is not None:
        anchor[1] = line_y
    if head_anchor:
        anchor = [(hair[0]+hair[2])/2, (hair[1]+hair[3])/2]
        reference = [hw, hh]
    if support_side:
        yy, xx = np.indices(subject.shape)
        planted = subject & (yy > image.height*.6) & ((xx-anchor[0])*support_side > 0)
        anchor[1] = float(np.quantile(np.where(planted)[0], .999))
    if wall_anchor:
        # Wall strips include a grey wall at the left and a grey mat below.
        # A two-foot setup and a one-foot hold have different foot centres;
        # registering those centres would make the stationary wall jump.
        yy, xx = np.indices(subject.shape)
        grey = (chroma < 15) & (pixels.mean(axis=2) < 242) & (pixels.mean(axis=2) > 100)
        upper_wall = grey & (xx < image.width*.4) & (yy > image.height*.15) & (yy < image.height*.45)
        edges = [np.where(row)[0][-1] for row in upper_wall if row.any()]
        if len(edges) < image.height*.1:
            raise ValueError("Wall reference was not found")
        wall_x = float(np.median(edges))
        lower_mat = grey & (xx > wall_x+12) & (yy > image.height*.65)
        mat = bounds(lower_mat)
        anchor = [wall_x, mat[3]]
    return {"bounds": box, "anchor": anchor, "reference": reference, "hair": hair}


def align_group(paths, output, size=640):
    exercise_id = paths[0].name[:4]
    floor = exercise_id in FLOOR_IDS
    frames = []
    for path in paths:
        with Image.open(path) as original:
            source = original.convert("RGB")
        cuts = panel_cuts(source)
        for index in range(3):
            left, right = cuts[index:index+2]
            image = source.crop((left, 0, right, source.height))
            try:
                features = analyse(image, floor, exercise_id in FOOT_ANCHOR_IDS,
                                   exercise_id in LINE_ANCHOR_IDS, exercise_id in HEAD_ANCHOR_IDS,
                                   (-1 if "other" in path.name else 1) if exercise_id == "it32" else 0,
                                   exercise_id in WALL_ANCHOR_IDS)
            except ValueError as error:
                raise ValueError(f"{path.name} frame {index}: {error}") from error
            frames.append({"file": path.name, "index": index, "image": image,
                           "sourceCrop": [left, 0, right, source.height], **features})
    reference = np.median([frame["reference"] for frame in frames], axis=0)
    if exercise_id in AIR_FRAMES:
        grounded = [f["anchor"][1] for f in frames if f["index"] not in AIR_FRAMES[exercise_id]]
        for frame in frames:
            if frame["index"] in AIR_FRAMES[exercise_id]:
                # Do not move an airborne foot down to the ground. It shares
                # the floor plane of the landing frames, retaining jump height.
                frame["anchor"][1] = float(np.median(grounded))
    extent = []
    for frame in frames:
        ratios = reference / np.asarray(frame["reference"])
        # A head turning changes its silhouette but not camera scale. Correct
        # zoom only when width and height agree; mat width is stable on the floor.
        scale = ratios[0] if floor and exercise_id not in HEAD_ANCHOR_IDS else (float(np.sqrt(ratios.prod())) if abs(np.log(ratios[0]/ratios[1])) < .12 else 1.)
        if exercise_id in NO_ZOOM_IDS:
            scale = 1.
        frame["registrationScale"] = float(np.clip(scale, .88, 1.12))
        x0, y0, x1, y1 = frame["bounds"]
        ax, ay = frame["anchor"]
        s = frame["registrationScale"]
        extent.append([(x0-ax)*s, (y0-ay)*s, (x1-ax)*s, (y1-ay)*s])
    union = np.asarray(extent)
    x0, y0 = union[:, :2].min(axis=0)
    x1, y1 = union[:, 2:].max(axis=0)
    common_scale = min(size*.84/(x1-x0), size*.84/(y1-y0))
    target = [size/2-(x0+x1)*common_scale/2,
              size/2-(y0+y1)*common_scale/2 if floor else size*.92-y1*common_scale]
    report = []
    for frame in frames:
        scale = frame["registrationScale"] * common_scale
        ax, ay = frame["anchor"]
        target_x = target[0] + size*SIDE_TRAVEL.get(exercise_id, [0, 0, 0])[frame["index"]]
        tx, ty = target_x-ax*scale, target[1]-ay*scale
        # Pillow maps output coordinates back to the original. Only uniform
        # scaling and translation are applied, never warps or synthesized pixels.
        frame["aligned"] = frame["image"].transform((size, size), Image.Transform.AFFINE,
            (1/scale, 0, -tx/scale, 0, 1/scale, -ty/scale), Image.Resampling.BICUBIC, fillcolor="white")
        report.append({key: value for key, value in frame.items() if key not in {"image", "aligned"}} |
                      {"scale": float(scale), "translate": [float(tx), float(ty)], "targetAnchor": [target_x, target[1]], "mode": "wall" if exercise_id in WALL_ANCHOR_IDS else ("head" if exercise_id in HEAD_ANCHOR_IDS else ("mat" if floor else "ground"))})
    output.mkdir(parents=True, exist_ok=True)
    for path in paths:
        strip = Image.new("RGB", (size*3, size), "white")
        for frame in frames:
            if frame["file"] == path.name:
                strip.paste(frame["aligned"], (frame["index"]*size, 0))
        strip.save(output/path.name, optimize=True)
    return report


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("source", type=Path)
    parser.add_argument("output", type=Path)
    parser.add_argument("--size", type=int, default=640)
    args = parser.parse_args()
    if args.source.resolve() == args.output.resolve():
        parser.error("Keep the originals in a different directory")
    groups = {}
    for path in sorted(args.source.glob("it*-v1.png")):
        groups.setdefault(path.name[:4], []).append(path)
    report = {key: align_group(paths, args.output, args.size) for key, paths in groups.items()}
    (args.output/"alignment.json").write_text(json.dumps(report, indent=2)+"\n")
    print(f"Aligned {sum(len(paths) for paths in groups.values())} strips, {len(groups)} exercises")


if __name__ == "__main__":
    main()
