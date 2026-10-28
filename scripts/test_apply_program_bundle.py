import importlib.util
import copy
import json
from pathlib import Path
import sqlite3
import struct
import tempfile
import unittest

spec = importlib.util.spec_from_file_location("apply_program_bundle", Path(__file__).with_name("apply_program_bundle.py"))
bundle = importlib.util.module_from_spec(spec)
spec.loader.exec_module(bundle)


class BundleContractTests(unittest.TestCase):
    def test_reaction_drill_accepts_both_directions_without_pictures(self):
        source = Path(__file__).resolve().parents[1] / "programs/ilya-tennis"
        docs = json.loads((source / "exercises.json").read_text())
        target = next(doc for doc in docs if doc["unit"] == "seconds" and not doc["perSide"])
        target.pop("poses", None)
        target.pop("animation", None)
        with tempfile.TemporaryDirectory() as tmp:
            folder = Path(tmp)
            (folder / "program.json").write_text((source / "program.json").read_text())
            for directions in (["left", "right"], ["right", "left"]):
                with self.subTest(directions=directions):
                    target["reaction"] = {"directions": directions}
                    (folder / "exercises.json").write_text(json.dumps(docs))
                    _, exercises, images = bundle.load_bundle(folder)
                    self.assertEqual(exercises[target["id"]]["reaction"], target["reaction"])
                    self.assertNotIn("poses", exercises[target["id"]])
                    self.assertEqual(images, [])

    def test_reaction_rejects_invalid_direction_types_and_authored_timing(self):
        valid = {"unit": "seconds", "perSide": False, "reaction": {"directions": ["left", "right"]}}
        cases = [
            {"reaction": []},
            {"reaction": "left"},
            {"reaction": {}},
            {"reaction": {"directions": "left,right"}},
            {"reaction": {"directions": ["left"]}},
            {"reaction": {"directions": ["left", "left"]}},
            {"reaction": {"directions": ["left", "forward"]}},
            {"reaction": {"directions": ["left", "right", "left"]}},
            {"reaction": {"directions": ["left", None]}},
            {"reaction": {"directions": ["left", ["right"]]}},
            {"reaction": {"directions": ["left", "right"], "interval": 3}},
            {"reaction": {"directions": ["left", "right"], "rest": 10}},
            {"reaction": {"directions": ["left", "right"], "sideRest": 3}},
            {"unit": "reps"},
            {"unit": "breaths"},
            {"perSide": True},
            {"perSide": 0},
            {"perSide": "false"},
        ]
        for changes in cases:
            with self.subTest(changes=changes):
                with self.assertRaisesRegex(ValueError, "Invalid reaction"):
                    bundle.validate_reaction(valid | changes)
        bundle.validate_reaction({"unit": "reps", "perSide": True})
        bundle.validate_reaction({"unit": "reps", "perSide": True, "reaction": None})

    def test_load_bundle_validates_reaction_before_install(self):
        source = Path(__file__).resolve().parents[1] / "programs/ilya-tennis"
        docs = json.loads((source / "exercises.json").read_text())
        docs[0]["reaction"] = {"directions": ["left", "right"], "interval": 3}
        with tempfile.TemporaryDirectory() as tmp:
            folder = Path(tmp)
            (folder / "program.json").write_text((source / "program.json").read_text())
            (folder / "exercises.json").write_text(json.dumps(docs))
            with self.assertRaisesRegex(ValueError, "Invalid reaction"):
                bundle.load_bundle(folder)

    def test_new_bundles_cannot_reintroduce_private_pauses_or_missing_phases(self):
        source = Path(__file__).resolve().parents[1] / "programs/ilya-tennis"
        program = json.loads((source / "program.json").read_text())
        with tempfile.TemporaryDirectory() as tmp:
            folder = Path(tmp)
            (folder / "exercises.json").write_text((source / "exercises.json").read_text())
            cases = [
                lambda p: p["warmup"][0].update(rest=15),
                lambda p: p["days"][0]["items"][0].update(sideRest=10),
                lambda p: p["days"][0].update(rounds=1),
                lambda p: p.update(warmupRounds=2),
                lambda p: p.update(warmup=[]),
                lambda p: p["days"][0].update(cooldown=[]),
                lambda p: p["days"][0].update(items=[]),
            ]
            (folder / "program.json").write_text(json.dumps(program))
            bundle.load_bundle(folder)
            for index, alter in enumerate(cases):
                with self.subTest(case=index):
                    invalid = copy.deepcopy(program)
                    alter(invalid)
                    (folder / "program.json").write_text(json.dumps(invalid))
                    with self.assertRaises(ValueError):
                        bundle.load_bundle(folder)


class ApplyBundleTests(unittest.TestCase):
    def setUp(self):
        self.connection = sqlite3.connect(":memory:")
        self.connection.row_factory = sqlite3.Row
        self.connection.executescript("""
            CREATE TABLE users(id INTEGER PRIMARY KEY,name TEXT,program_json TEXT,custom_exercises TEXT,active_program TEXT,position INTEGER,level INTEGER);
            CREATE TABLE poses(id TEXT PRIMARY KEY,kind TEXT,label TEXT);
            CREATE TABLE pose_images(pose_id TEXT,sex TEXT,mime TEXT,data BLOB,PRIMARY KEY(pose_id,sex));
            INSERT INTO users VALUES(44,'Ilya','{}','{"existing":{"id":"existing"}}','custom',39,2);
            INSERT INTO users VALUES(45,'Other','{}','{}','custom',7,1);
        """)
        self.temp = tempfile.TemporaryDirectory()
        self.backup = Path(self.temp.name) / "backup.json"

    def tearDown(self):
        self.connection.close()
        self.temp.cleanup()

    def test_atomic_account_only_update_preserves_progress_and_custom_exercises(self):
        position = bundle.apply_bundle(self.connection, 44, "Ilya", {"name": "New"}, {"IT01": {"id": "IT01"}}, [], self.backup)
        row = self.connection.execute("SELECT * FROM users WHERE id=44").fetchone()
        self.assertEqual((position, row["position"], row["level"]), (39, 39, 2))
        self.assertEqual(set(json.loads(row["custom_exercises"])), {"existing", "IT01"})
        self.assertEqual(json.loads(self.backup.read_text())["program_json"], "{}")
        other = self.connection.execute("SELECT program_json,position FROM users WHERE id=45").fetchone()
        self.assertEqual(tuple(other), ("{}", 7))

    def test_mismatched_identity_refuses_before_backup_or_write(self):
        with self.assertRaisesRegex(ValueError, "Target account"):
            bundle.apply_bundle(self.connection, 44, "Wrong", {}, {}, [], self.backup)
        self.assertFalse(self.backup.exists())
        self.assertEqual(self.connection.execute("SELECT position FROM users WHERE id=44").fetchone()[0], 39)

    def test_invalid_reaction_refuses_before_backup_or_write(self):
        before = tuple(self.connection.execute("SELECT * FROM users WHERE id=44").fetchone())
        invalid = {"id": "IT36", "unit": "seconds", "perSide": False,
                   "reaction": {"directions": ["left", "right"], "rest": 10}}
        with self.assertRaisesRegex(ValueError, "Invalid reaction"):
            bundle.apply_bundle(self.connection, 44, "Ilya", {"name": "New"}, {"IT36": invalid}, [], self.backup)
        self.assertFalse(self.backup.exists())
        self.assertEqual(tuple(self.connection.execute("SELECT * FROM users WHERE id=44").fetchone()), before)

    def test_pose_collision_rolls_back_every_picture_and_account_change(self):
        self.connection.execute("INSERT INTO poses VALUES('existing','test','Existing')")
        self.connection.execute("INSERT INTO pose_images VALUES('existing','m','image/png',?)", (b"old",))
        self.connection.commit()
        images = [{"id": "new", "data": b"new"}, {"id": "existing", "data": b"different"}]
        with self.assertRaisesRegex(ValueError, "Existing pose differs"):
            bundle.apply_bundle(self.connection, 44, "Ilya", {"name": "New"}, {}, images, self.backup)
        self.assertEqual(self.connection.execute("SELECT count(*) FROM poses WHERE id='new'").fetchone()[0], 0)
        self.assertEqual(self.connection.execute("SELECT program_json FROM users WHERE id=44").fetchone()[0], "{}")

    def test_images_only_preserves_program_and_all_other_exercise_fields(self):
        program_before = '{ "name": "Individually edited", "days": [1] }'
        exercises_before = {
            "existing": {"id": "existing", "content": {"ru": {"name": "Personal name"}},
                         "poses": ["old"], "unit": "reps", "value": 17,
                         "extra": {"preserve": [1, 2]}},
            "untouched": {"id": "untouched", "poses": ["leave-alone"]},
        }
        self.connection.execute(
            "UPDATE users SET program_json=?,custom_exercises=?,active_program='demo' WHERE id=44",
            (program_before, json.dumps(exercises_before)),
        )
        self.connection.commit()
        other_before = tuple(self.connection.execute("SELECT * FROM users WHERE id=45").fetchone())
        supplied = {"existing": {"id": "existing", "content": {"ru": {"name": "Overwrite attempt"}},
                                 "unit": "seconds", "poses": ["fresh"]}}
        result = bundle.apply_bundle(
            self.connection, 44, "Ilya", {"name": "Overwrite attempt"}, supplied,
            [{"id": "fresh", "data": b"image"}], self.backup, images_only=True,
        )
        row = self.connection.execute("SELECT * FROM users WHERE id=44").fetchone()
        self.assertEqual(row["program_json"], program_before)
        self.assertEqual((row["active_program"], row["position"], row["level"], result), ("demo", 39, 2, 39))
        exercises_before["existing"]["poses"] = ["fresh"]
        self.assertEqual(json.loads(row["custom_exercises"]), exercises_before)
        self.assertEqual(tuple(self.connection.execute("SELECT * FROM users WHERE id=45").fetchone()), other_before)
        self.assertEqual(self.connection.execute("SELECT data FROM pose_images WHERE pose_id='fresh'").fetchone()[0], b"image")
        backup = json.loads(self.backup.read_text())
        self.assertEqual(backup["program_json"], program_before)
        self.assertEqual(json.loads(backup["custom_exercises"])["existing"]["poses"], ["old"])

    def test_images_only_refuses_missing_exercises_and_empty_poses_before_backup(self):
        before = tuple(self.connection.execute("SELECT * FROM users WHERE id=44").fetchone())
        cases = [
            ({"missing": {"id": "missing", "poses": ["fresh"]}}, "Existing exercise not found"),
            ({"existing": {"id": "existing"}}, "Nonempty pose assignment"),
            ({"existing": {"id": "existing", "poses": []}}, "Nonempty pose assignment"),
            ({"existing": {"id": "existing", "poses": [""]}}, "Nonempty pose assignment"),
            ({"existing": {"id": "existing", "poses": "fresh"}}, "Nonempty pose assignment"),
            ({"existing": {"id": "existing", "poses": ["not-supplied"]}}, "Pose image not found"),
        ]
        for exercises, message in cases:
            with self.subTest(exercises=exercises):
                with self.assertRaisesRegex(ValueError, message):
                    bundle.apply_bundle(
                        self.connection, 44, "Ilya", {}, exercises,
                        [{"id": "fresh", "data": b"image"}], self.backup, images_only=True,
                    )
                self.assertFalse(self.backup.exists())
                self.assertEqual(tuple(self.connection.execute("SELECT * FROM users WHERE id=44").fetchone()), before)
                self.assertEqual(self.connection.execute("SELECT count(*) FROM pose_images").fetchone()[0], 0)

    def test_images_only_collision_rolls_back_assignments_and_new_images(self):
        self.connection.execute("INSERT INTO poses VALUES('collision','test','Existing')")
        self.connection.execute("INSERT INTO pose_images VALUES('collision','m','image/png',?)", (b"old",))
        self.connection.commit()
        before = tuple(self.connection.execute("SELECT * FROM users WHERE id=44").fetchone())
        with self.assertRaisesRegex(ValueError, "Existing pose differs"):
            bundle.apply_bundle(
                self.connection, 44, "Ilya", {}, {"existing": {"poses": ["fresh"]}},
                [{"id": "fresh", "data": b"image"}, {"id": "collision", "data": b"different"}],
                self.backup, images_only=True,
            )
        self.assertTrue(self.backup.exists())
        self.assertEqual(tuple(self.connection.execute("SELECT * FROM users WHERE id=44").fetchone()), before)
        self.assertEqual(self.connection.execute("SELECT count(*) FROM pose_images WHERE pose_id='fresh'").fetchone()[0], 0)
        self.assertEqual(self.connection.execute("SELECT data FROM pose_images WHERE pose_id='collision'").fetchone()[0], b"old")

    def test_animation_only_metadata_preserves_doses_text_and_other_fields(self):
        existing = {"existing": {"id": "existing", "poses": ["old"], "unit": "reps", "content": {"ru": {"name": "Personal"}}}}
        self.connection.execute("UPDATE users SET custom_exercises=? WHERE id=44", (json.dumps(existing),))
        self.connection.commit()
        animation = {"pose": "motion", "otherPose": "other", "frames": 3, "kind": "movement"}
        png = b"\x89PNG\r\n\x1a\n\0\0\0\rIHDR" + struct.pack(">II", 300, 100)
        supplied = {"existing": {"poses": ["old"], "animation": animation, "unit": "seconds"}}
        images = [{"id": "old", "data": b"still"}, {"id": "motion", "data": png}, {"id": "other", "data": png}]
        bundle.apply_bundle(self.connection, 44, "Ilya", {}, supplied, images, self.backup, images_only=True)
        expected = existing
        expected["existing"]["animation"] = animation
        row = self.connection.execute("SELECT * FROM users WHERE id=44").fetchone()
        self.assertEqual(json.loads(row["custom_exercises"]), expected)
        self.assertEqual((row["program_json"], row["position"], row["level"]), ("{}", 39, 2))

    def test_invalid_or_missing_animation_refuses_atomically(self):
        png = b"\x89PNG\r\n\x1a\n\0\0\0\rIHDR" + struct.pack(">II", 300, 100)
        cases = [
            ({"pose": "missing", "frames": 3, "kind": "movement"}, png),
            ({"pose": "motion", "otherPose": "missing", "frames": 3, "kind": "sequence"}, png),
            ({"pose": "motion", "frames": 3, "kind": "alternating"}, png[:-8] + struct.pack(">II", 100, 100)),
            ({"pose": "motion", "frames": 2, "kind": "alternating"}, png),
            ({"pose": "../motion", "frames": 3, "kind": "movement"}, png),
        ]
        before = tuple(self.connection.execute("SELECT * FROM users WHERE id=44").fetchone())
        for animation, data in cases:
            with self.subTest(animation=animation):
                with self.assertRaises(ValueError):
                    bundle.apply_bundle(self.connection, 44, "Ilya", {}, {"existing": {"poses": ["motion"], "animation": animation}}, [{"id": "motion", "data": data}], self.backup, images_only=True)
                self.assertFalse(self.backup.exists())
                self.assertEqual(tuple(self.connection.execute("SELECT * FROM users WHERE id=44").fetchone()), before)
                self.assertEqual(self.connection.execute("SELECT count(*) FROM pose_images").fetchone()[0], 0)


if __name__ == "__main__":
    unittest.main()
