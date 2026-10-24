import importlib.util
from pathlib import Path
import tempfile
import unittest

from PIL import Image, ImageDraw

spec = importlib.util.spec_from_file_location("align_exercise_frames", Path(__file__).with_name("align_exercise_frames.py"))
align = importlib.util.module_from_spec(spec)
spec.loader.exec_module(align)


class AlignmentTests(unittest.TestCase):
    def test_unequal_panels_are_split_in_whitespace_without_cutting_a_head(self):
        strip = Image.new("RGB", (900, 300), "white")
        draw = ImageDraw.Draw(strip)
        for box in [(20, 20, 235, 280), (270, 20, 580, 280), (620, 20, 880, 280)]:
            draw.rectangle(box, fill=(30, 130, 220))
        cuts = align.panel_cuts(strip)
        self.assertTrue(235 < cuts[1] < 270)
        self.assertTrue(580 < cuts[2] < 620)

    def test_occluded_mat_uses_both_visible_halves(self):
        for offset in [-30, 30]:
            frame = Image.new("RGB", (300, 300), "white")
            draw = ImageDraw.Draw(frame)
            draw.rectangle((20, 100, 280, 280), fill=(180, 180, 180))
            draw.rectangle((100+offset, 80, 200+offset, 285), fill=(30, 130, 220))
            draw.rectangle((110+offset, 30, 190+offset, 75), fill=(100, 60, 30))
            features = align.analyse(frame, True)
            self.assertEqual(features["reference"][0], 260)
            self.assertEqual(features["anchor"], [150, 280])

    def test_line_is_the_reference_when_feet_step_across_it(self):
        for foot_y in [220, 270]:
            frame = Image.new("RGB", (300, 300), "white")
            draw = ImageDraw.Draw(frame)
            draw.rectangle((0, 235, 299, 241), fill=(50, 150, 240))
            draw.rectangle((110, 30, 190, 75), fill=(100, 60, 30))
            draw.rectangle((100, 80, 200, 190), fill=(30, 130, 220))
            draw.rectangle((100, 190, 120, foot_y), fill=(240, 160, 110))
            self.assertEqual(align.analyse(frame, False, line_anchor=True)["anchor"][1], 238)

    def test_camera_shift_is_removed_without_enlarging_crouching_pose(self):
        strip = Image.new("RGB", (900, 300), "white")
        for index, (offset, crouch) in enumerate([(0, 0), (18, 35), (-14, 70)]):
            frame = Image.new("RGB", (300, 300), "white")
            draw = ImageDraw.Draw(frame)
            # Constant head size, changing posture, and unwanted camera drift.
            draw.rectangle((110+offset, 20+crouch, 190+offset, 70+crouch), fill=(100, 60, 30))
            draw.rectangle((100+offset, 75+crouch, 200+offset, 220), fill=(30, 130, 220))
            draw.rectangle((100+offset, 221, 120+offset, 280), fill=(240, 160, 110))
            draw.rectangle((180+offset, 221, 200+offset, 280), fill=(240, 160, 110))
            strip.paste(frame, (index*300, 0))
        with tempfile.TemporaryDirectory() as folder:
            source = Path(folder)/"it06-v1.png"
            strip.save(source)
            output = Path(folder)/"aligned"
            report = align.align_group([source], output, 320)
            self.assertTrue(all(abs(row["registrationScale"]-1) < .001 for row in report))
            result = Image.open(output/source.name)
            features = [align.analyse(result.crop((i*320, 0, (i+1)*320, 320)), False) for i in range(3)]
            centres = [(row["hair"][0]+row["hair"][2])/2 for row in features]
            self.assertLessEqual(max(centres)-min(centres), 4)
            floors = [row["bounds"][3] for row in features]
            self.assertLessEqual(max(floors)-min(floors), 1)
            self.assertLess(features[0]["hair"][1], features[1]["hair"][1])
            self.assertLess(features[1]["hair"][1], features[2]["hair"][1])
            self.assertEqual(result.size, (960, 320))
            result.close()
            with Image.open(source) as original:
                self.assertEqual(original.size, (900, 300))


if __name__ == "__main__":
    unittest.main()
