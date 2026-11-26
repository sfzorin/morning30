package content

import (
	_ "embed"
	"encoding/json"
	"fmt"
)

// More ready programs: arms, pull-ups, 50+, office, tai chi, seven stretching
// programs, five for children (on the kids' exercises with their own
// drawings) and six for athletes. Titles and descriptions in every language
// live in courses_more.json.

//go:embed courses_more.json
var moreCoursesText []byte

func d(id string, from, to int) dose         { return dose{ID: id, From: from, To: to} }
func ds(id string, from, to, since int) dose { return dose{ID: id, From: from, To: to, Since: since} }
func set(pairs ...any) []setDef {
	out := make([]setDef, 0, len(pairs)/2)
	for i := 0; i+1 < len(pairs); i += 2 {
		out = append(out, setDef{pairs[i].(string), pairs[i+1].(int)})
	}
	return out
}

var (
	stretchWarmup = set("J02", 30, "W01", 15, "W03", 6, "W04", 8, "ST15", 6, "W05", 8)
	stretchCool   = set("ST22", 30, "CD03", 30, "CD07", 30, "CD05", 40)
	sportWarmup   = set("J02", 30, "W09", 20, "W08", 20, "W05", 8, "W04", 8, "W10", 6)
	kidsWarmup    = set("IT01", 40, "IT04", 10, "IT02", 10, "IT05", 8, "IT06", 8)
	kidsCool      = set("IT13", 20, "IT52", 20, "IT53", 30, "IT14", 40)
	stretchRecov  = []dose{d("ST13", 20, 30), d("ST03", 20, 30), d("ST22", 20, 30)}
)

var moreCourses = []Course{
	{Key: "arms30", Slug: "arms-in-30-days", Category: CatStrength, plan: coursePlan{
		warmup: set("W01", 20, "W02", 30, "W09", 20, "W03", 6, "W04", 8, "W06", 20), cooldown: set("ST17", 25, "ST18", 25, "CD01", 30, "NK07", 30),
		a:        []dose{d("P09", 6, 14), d("A01", 8, 14), d("A05", 6, 14), d("A02", 15, 30), d("B08", 15, 35)},
		b:        []dose{d("P01", 4, 12), d("A04", 4, 8), d("A03", 15, 30), d("C10", 4, 10), d("B06", 8, 14)},
		c:        []dose{d("P02", 3, 9), d("A05", 8, 16), d("A01", 8, 14), d("C09", 6, 14), ds("P06", 3, 8, 8)},
		recovery: []dose{d("P09", 5, 8), d("A02", 15, 20), d("B01", 20, 30)},
		final:    set("P01", 14, "A05", 16, "P02", 8, "A01", 14, "B08", 35, "C10", 10),
	}},
	{Key: "pullups30", Slug: "pull-ups-in-30-days", Category: CatStrength, plan: coursePlan{
		warmup: set("W01", 20, "W02", 30, "W09", 20, "W03", 6, "W06", 20, "W04", 8), cooldown: set("ST17", 25, "ST18", 25, "CD01", 30, "CD07", 30),
		a:        []dose{d("PU01", 10, 40), d("PU02", 4, 10), d("PU07", 5, 12), d("B06", 8, 14), d("C01", 20, 45)},
		b:        []dose{d("PU05", 3, 8), d("PU04", 10, 25), d("B03", 8, 14), d("C12", 6, 10), ds("PU03", 2, 5, 8)},
		c:        []dose{d("PU02", 5, 12), d("PU07", 6, 14), d("PU01", 15, 45), d("B02", 8, 14), ds("PU03", 2, 6, 12), ds("PU06", 1, 4, 20)},
		recovery: []dose{d("PU01", 10, 20), d("B01", 20, 30), d("C12", 6, 8)},
		final:    set("PU06", 5, "PU04", 25, "PU07", 14, "PU03", 6, "PU01", 45, "B06", 14),
	}},
	{Key: "age50", Slug: "workout-50-plus", Category: CatStart, plan: coursePlan{
		warmup: set("W01", 15, "W09", 20, "W10", 6, "NK04", 8, "W04", 8, "W05", 8), cooldown: set("ST20", 25, "ST16", 25, "ST17", 25, "NK05", 25),
		a:        []dose{d("BL04", 5, 12), d("BL01", 10, 30), d("L07", 8, 16), d("BL05", 6, 12), d("L06", 8, 14)},
		b:        []dose{d("BL02", 20, 60), d("L03", 8, 14), d("L15", 5, 10), d("NK02", 6, 12), ds("BL03", 4, 8, 10)},
		c:        []dose{d("BL04", 6, 14), d("BL01", 15, 35), d("L06", 8, 14), d("A01", 6, 10), d("BL05", 6, 12), d("L07", 10, 18)},
		recovery: []dose{d("BL01", 10, 20), d("NK02", 5, 8), d("L07", 8, 12)},
		final:    set("BL04", 14, "BL01", 35, "BL02", 60, "BL03", 8, "L15", 10, "L06", 14),
	}},
	{Key: "office30", Slug: "neck-and-back-for-desk-workers", Category: CatBackBalance, plan: coursePlan{
		warmup: set("NK03", 4, "NK04", 8, "W01", 15, "W03", 6, "W08", 20, "W07", 20), cooldown: set("NK01", 25, "NK05", 25, "NK06", 25, "NK07", 25),
		a:        []dose{d("NK02", 6, 14), d("B03", 8, 14), d("B01", 20, 35), d("ST19", 20, 35), d("B04", 6, 12)},
		b:        []dose{d("B06", 8, 14), d("NK02", 6, 12), d("B05", 6, 12), d("ST17", 20, 30), d("L03", 8, 14)},
		c:        []dose{d("B10", 6, 12), d("B02", 8, 14), d("NK02", 8, 14), d("C01", 20, 40), d("ST04", 20, 35)},
		recovery: []dose{d("NK02", 5, 8), d("B01", 20, 25), d("ST19", 20, 25)},
		final:    set("NK02", 14, "B10", 12, "B03", 14, "B02", 14, "B01", 35, "ST19", 35),
	}},
	{Key: "taichi30", Slug: "tai-chi-in-30-days", Category: CatBackBalance, plan: coursePlan{
		warmup: set("TC11", 60, "W01", 20, "W03", 6, "NK04", 8, "W04", 8, "W10", 6), cooldown: set("TC10", 60, "ST20", 25, "NK01", 25, "ST17", 25),
		a:        []dose{d("TC01", 30, 60), d("TC02", 30, 60), d("TC03", 20, 40), d("TC06", 15, 30), d("TC09", 20, 45)},
		b:        []dose{d("TC01", 30, 45), d("TC04", 20, 40), d("TC05", 30, 60), d("TC07", 10, 25), d("TC02", 30, 45)},
		c:        []dose{d("TC01", 30, 45), d("TC03", 20, 40), d("TC04", 20, 40), d("TC09", 30, 45), d("TC02", 30, 60), ds("TC08", 20, 40, 10)},
		recovery: []dose{d("TC01", 30, 45), d("TC02", 30, 45), d("TC06", 15, 25)},
		final:    set("TC01", 45, "TC02", 60, "TC03", 40, "TC04", 40, "TC08", 40, "TC09", 45),
	}},
	{Key: "stretch30", Slug: "stretching-in-30-days", Category: CatStretch, plan: coursePlan{
		warmup: stretchWarmup, cooldown: stretchCool,
		a:        []dose{d("ST01", 20, 45), d("ST04", 20, 40), d("ST12", 20, 35), d("ST13", 20, 40), d("ST17", 20, 30), d("ST03", 20, 45)},
		b:        []dose{d("ST02", 20, 45), d("ST08", 20, 40), d("ST19", 20, 35), d("ST16", 20, 35), d("ST20", 15, 30), d("CD04", 20, 40)},
		c:        []dose{d("ST11", 20, 45), d("ST09", 20, 40), d("ST14", 20, 45), d("ST18", 15, 30), d("ST07", 20, 45), ds("ST05", 20, 40, 8)},
		recovery: stretchRecov,
		final:    set("ST01", 45, "ST04", 40, "ST09", 40, "ST05", 40, "ST07", 45, "ST19", 35, "ST14", 45),
	}},
	{Key: "sidesplit30", Slug: "side-splits-in-30-days", Category: CatStretch, plan: coursePlan{
		warmup: stretchWarmup, cooldown: set("ST22", 30, "ST13", 30, "CD07", 30, "CD05", 40),
		a:        []dose{d("ST03", 20, 50), d("ST08", 20, 45), d("ST11", 20, 50), d("ST07", 20, 50), ds("ST06", 15, 40, 6)},
		b:        []dose{d("ST14", 20, 45), d("ST08", 20, 45), d("ST07", 20, 50), d("ST03", 25, 50), d("CD03", 20, 30)},
		c:        []dose{d("ST11", 25, 50), d("ST06", 20, 45), d("ST08", 25, 45), d("ST07", 25, 60), d("L13", 6, 10)},
		recovery: stretchRecov,
		final:    set("ST03", 50, "ST11", 50, "ST08", 45, "ST06", 45, "ST07", 60, "ST14", 45),
	}},
	{Key: "frontsplit30", Slug: "front-splits-in-30-days", Category: CatStretch, plan: coursePlan{
		warmup: stretchWarmup, cooldown: set("ST22", 30, "ST13", 30, "CD07", 30, "CD05", 40),
		a:        []dose{d("ST04", 20, 45), d("ST09", 20, 45), d("ST12", 20, 35), d("ST01", 20, 40), ds("ST10", 15, 40, 8)},
		b:        []dose{d("ST09", 20, 45), d("ST05", 20, 40), d("ST04", 25, 45), d("CD02", 20, 40), d("ST16", 20, 30)},
		c:        []dose{d("ST04", 25, 50), d("ST09", 25, 50), d("ST12", 20, 40), d("ST02", 20, 45), ds("ST10", 15, 45, 10)},
		recovery: []dose{d("ST13", 20, 30), d("ST04", 20, 30), d("ST22", 20, 30)},
		final:    set("ST04", 45, "ST09", 45, "ST10", 40, "ST12", 35, "ST05", 35, "ST02", 40),
	}},
	{Key: "toetouch30", Slug: "touch-your-toes-in-30-days", Category: CatStretch, plan: coursePlan{
		warmup: stretchWarmup, cooldown: stretchCool,
		a:        []dose{d("ST01", 20, 45), d("ST02", 20, 45), d("CD02", 20, 40), d("ST16", 20, 30)},
		b:        []dose{d("ST09", 20, 40), d("ST11", 20, 45), d("ST13", 20, 35), d("CD03", 20, 30)},
		c:        []dose{d("ST01", 25, 50), d("ST07", 20, 45), d("ST09", 20, 45), d("ST22", 20, 30), d("ST02", 25, 50)},
		recovery: []dose{d("ST13", 20, 30), d("ST22", 20, 30), d("CD02", 20, 30)},
		final:    set("ST01", 50, "ST02", 50, "ST09", 45, "ST07", 45, "CD02", 40, "ST11", 45),
	}},
	{Key: "hips30", Slug: "hip-mobility-in-30-days", Category: CatStretch, plan: coursePlan{
		warmup: stretchWarmup, cooldown: stretchCool,
		a:        []dose{d("ST03", 20, 45), d("ST13", 20, 40), d("ST04", 20, 40), d("ST14", 20, 45), d("L06", 8, 14)},
		b:        []dose{d("ST05", 20, 40), d("ST06", 15, 40), d("ST08", 20, 40), d("CD03", 20, 30), d("L04", 5, 10)},
		c:        []dose{d("ST14", 25, 50), d("ST05", 25, 45), d("ST03", 25, 50), d("ST04", 25, 45), d("ST13", 25, 45)},
		recovery: stretchRecov,
		final:    set("ST14", 50, "ST05", 45, "ST06", 40, "ST03", 50, "ST04", 45, "ST08", 40),
	}},
	{Key: "morningstretch", Slug: "morning-stretch", Category: CatStretch, plan: coursePlan{
		warmup: set("W01", 15, "W03", 6, "ST15", 6, "W04", 6, "W09", 15, "W07", 20), cooldown: set("ST20", 20, "ST17", 20, "NK01", 20, "CD07", 25),
		a:        []dose{d("ST01", 20, 35), d("ST12", 15, 30), d("ST20", 15, 25), d("ST04", 15, 30)},
		b:        []dose{d("ST19", 15, 30), d("ST16", 15, 30), d("ST17", 15, 25), d("ST13", 15, 30)},
		c:        []dose{d("ST01", 20, 35), d("ST09", 15, 30), d("ST18", 15, 25), d("ST04", 15, 30)},
		recovery: []dose{d("ST01", 15, 25), d("ST20", 15, 20), d("ST17", 15, 20)},
		final:    set("ST01", 35, "ST09", 30, "ST04", 30, "ST19", 30, "ST12", 30),
	}},
	{Key: "bedtime", Slug: "bedtime-stretch", Category: CatStretch, plan: coursePlan{
		warmup: set("W07", 30, "W01", 15, "W03", 6, "ST15", 6, "NK04", 6, "NK03", 4), cooldown: set("ST21", 60, "ST22", 30, "CD05", 40, "CD03", 30),
		a:        []dose{d("CD07", 30, 60), d("ST13", 20, 40), d("CD03", 20, 40), d("ST22", 20, 40)},
		b:        []dose{d("ST19", 20, 40), d("ST03", 20, 45), d("CD02", 20, 40), d("CD07", 30, 60)},
		c:        []dose{d("ST13", 25, 45), d("ST21", 30, 60), d("CD03", 25, 45), d("ST03", 20, 45)},
		recovery: []dose{d("CD07", 30, 40), d("ST22", 20, 30), d("ST13", 20, 30)},
		final:    set("CD07", 60, "ST13", 45, "ST19", 40, "CD03", 45, "ST21", 60),
	}},
	{Key: "kids30", Slug: "kids-workout", Category: CatKids, Kids: true, plan: coursePlan{
		warmup: kidsWarmup, cooldown: kidsCool,
		a:        []dose{d("IT38", 5, 10), d("IT09", 3, 8), d("IT12", 5, 8), d("IT16", 6, 10), d("IT40", 4, 8)},
		b:        []dose{d("IT15", 20, 40), d("IT27", 5, 8), d("IT28", 10, 25), d("IT46", 6, 10), d("IT44", 6, 12)},
		c:        []dose{d("IT29", 8, 16), d("IT39", 6, 12), d("IT47", 4, 8), d("IT17", 10, 25), d("IT37", 20, 40)},
		recovery: []dose{d("IT20", 6, 8), d("IT22", 6, 8), d("IT17", 10, 20)},
		final:    set("IT38", 10, "IT09", 8, "IT27", 8, "IT39", 12, "IT37", 40, "IT17", 25),
	}},
	{Key: "kidsagility", Slug: "kids-agility-and-coordination", Category: CatKids, Kids: true, plan: coursePlan{
		warmup: kidsWarmup, cooldown: kidsCool,
		a:        []dose{d("IT24", 4, 8), d("IT07", 4, 8), d("IT29", 8, 16), d("IT17", 10, 25), d("IT36", 15, 30)},
		b:        []dose{d("IT34", 15, 30), d("IT30", 6, 12), d("IT32", 4, 8), d("IT35", 15, 30), d("IT16", 6, 10)},
		c:        []dose{d("IT31", 15, 30), d("IT49", 15, 30), d("IT24", 4, 8), d("IT43", 15, 30), d("IT11", 5, 10)},
		recovery: []dose{d("IT17", 10, 20), d("IT29", 6, 10), d("IT18", 8, 12)},
		final:    set("IT24", 8, "IT07", 8, "IT36", 30, "IT34", 30, "IT32", 8, "IT49", 30),
	}},
	{Key: "kidsposture", Slug: "posture-for-schoolchildren", Category: CatKids, Kids: true, plan: coursePlan{
		warmup: kidsWarmup, cooldown: kidsCool,
		a:        []dose{d("IT10", 6, 12), d("IT46", 6, 10), d("IT20", 6, 10), d("IT21", 4, 8), d("IT40", 4, 8)},
		b:        []dose{d("IT22", 6, 10), d("IT12", 5, 8), d("IT10", 6, 12), d("IT28", 10, 25), d("IT54", 15, 25)},
		c:        []dose{d("IT46", 6, 12), d("IT39", 6, 12), d("IT21", 5, 8), d("IT48", 4, 8), d("IT10", 8, 12)},
		recovery: []dose{d("IT20", 6, 8), d("IT22", 6, 8), d("IT21", 4, 6)},
		final:    set("IT10", 12, "IT46", 12, "IT39", 12, "IT48", 8, "IT21", 8, "IT28", 25),
	}},
	{Key: "kidstennis", Slug: "tennis-for-kids", Category: CatKids, Kids: true, plan: coursePlan{
		warmup: kidsWarmup, cooldown: kidsCool,
		a:        []dose{d("IT07", 4, 8), d("IT19", 4, 8), d("IT33", 4, 8), d("IT15", 20, 40), d("IT42", 4, 8)},
		b:        []dose{d("IT34", 15, 30), d("IT11", 5, 10), d("IT36", 15, 30), d("IT03", 6, 10), d("IT41", 10, 20)},
		c:        []dose{d("IT31", 15, 30), d("IT49", 15, 30), d("IT47", 4, 8), d("IT39", 6, 12), d("IT19", 5, 10)},
		recovery: []dose{d("IT17", 10, 20), d("IT19", 4, 6), d("IT20", 6, 8)},
		final:    set("IT07", 8, "IT33", 8, "IT34", 30, "IT42", 8, "IT49", 30, "IT39", 12),
	}},
	{Key: "kidsfootball", Slug: "football-for-kids", Category: CatKids, Kids: true, plan: coursePlan{
		warmup: kidsWarmup, cooldown: kidsCool,
		a:        []dose{d("IT16", 6, 10), d("IT17", 10, 25), d("IT38", 5, 10), d("IT11", 5, 10), d("IT37", 20, 40)},
		b:        []dose{d("IT30", 6, 12), d("IT42", 4, 8), d("IT32", 4, 8), d("IT27", 5, 8), d("IT35", 15, 30)},
		c:        []dose{d("IT43", 15, 30), d("IT24", 4, 8), d("IT45", 4, 8), d("IT47", 4, 8), d("IT34", 15, 30)},
		recovery: []dose{d("IT17", 10, 20), d("IT18", 8, 12), d("IT22", 6, 8)},
		final:    set("IT24", 8, "IT38", 10, "IT45", 8, "IT43", 30, "IT34", 30, "IT17", 25),
	}},
	{Key: "runners", Slug: "strength-for-runners", Category: CatSport, plan: coursePlan{
		warmup: sportWarmup, cooldown: set("ST04", 30, "ST16", 30, "ST12", 30, "ST13", 30),
		a:        []dose{d("L10", 10, 20), d("L04", 5, 10), d("L15", 6, 12), d("C07", 15, 30), d("BL03", 4, 8)},
		b:        []dose{d("L11", 6, 12), d("L06", 10, 16), d("BL01", 15, 35), d("C12", 6, 10), d("L07", 12, 24)},
		c:        []dose{d("L13", 6, 10), d("L05", 6, 12), d("L14", 20, 45), d("C01", 20, 45), d("J15", 20, 40)},
		recovery: []dose{d("L06", 10, 14), d("BL01", 15, 25), d("C12", 6, 8)},
		final:    set("L10", 20, "L04", 10, "L15", 12, "BL03", 8, "L14", 45, "J15", 40),
	}},
	{Key: "football", Slug: "training-for-football-players", Category: CatSport, plan: coursePlan{
		warmup: sportWarmup, cooldown: set("ST04", 30, "ST03", 30, "ST13", 30, "ST16", 30),
		a:        []dose{d("L13", 6, 10), d("C14", 10, 25), d("L04", 5, 10), d("BL03", 4, 8), d("J10", 6, 12)},
		b:        []dose{d("L11", 6, 12), d("BL01", 15, 35), d("C07", 15, 30), d("L15", 6, 12), ds("J14", 3, 6, 10)},
		c:        []dose{d("L10", 10, 20), d("C14", 10, 25), d("L05", 6, 12), d("J10", 6, 14), d("C01", 20, 45)},
		recovery: []dose{d("BL01", 15, 25), d("L06", 10, 14), d("C12", 6, 8)},
		final:    set("L13", 10, "C14", 25, "J10", 14, "BL03", 8, "J14", 6, "C01", 45),
	}},
	{Key: "tennis", Slug: "training-for-tennis-players", Category: CatSport, plan: coursePlan{
		warmup: sportWarmup, cooldown: set("ST17", 30, "ST04", 30, "ST13", 30, "NK06", 30),
		a:        []dose{d("L13", 6, 10), d("J10", 6, 12), d("B03", 8, 14), d("C07", 15, 30), d("P01", 4, 10)},
		b:        []dose{d("L11", 6, 12), d("J06", 15, 30), d("B05", 6, 12), d("C12", 6, 10), d("BL01", 15, 30)},
		c:        []dose{d("L10", 10, 20), d("J10", 6, 14), d("B04", 6, 12), d("C09", 6, 14), d("L04", 5, 10)},
		recovery: []dose{d("B03", 8, 10), d("BL01", 15, 20), d("C12", 6, 8)},
		final:    set("L13", 10, "J10", 14, "J06", 30, "B05", 12, "C07", 30, "P01", 10),
	}},
	{Key: "ski", Slug: "ski-and-snowboard-training", Category: CatSport, plan: coursePlan{
		warmup: sportWarmup, cooldown: set("ST04", 30, "ST12", 30, "ST13", 30, "ST16", 30),
		a:        []dose{d("L14", 20, 60), d("L10", 10, 20), d("BL01", 15, 35), d("C07", 15, 30), d("L13", 6, 10)},
		b:        []dose{d("L11", 6, 12), d("J06", 15, 30), d("L06", 10, 16), d("C01", 20, 45), d("L07", 12, 24)},
		c:        []dose{d("L14", 25, 60), d("J10", 6, 12), d("BL03", 4, 8), d("L02", 5, 10), d("C12", 6, 10)},
		recovery: []dose{d("L14", 20, 30), d("BL01", 15, 20), d("L06", 10, 12)},
		final:    set("L14", 60, "L10", 20, "J06", 30, "J10", 12, "L11", 12, "C01", 45),
	}},
	{Key: "martial", Slug: "training-for-martial-arts", Category: CatSport, plan: coursePlan{
		warmup: sportWarmup, cooldown: set("ST04", 30, "ST03", 30, "ST08", 30, "CD07", 30),
		a:        []dose{d("J12", 30, 60), d("J13", 6, 12), d("L10", 10, 20), d("P01", 4, 12), d("C13", 20, 40)},
		b:        []dose{d("J12", 30, 60), d("J15", 20, 45), d("L13", 6, 10), d("C09", 6, 14), d("C05", 8, 14)},
		c:        []dose{d("J12", 40, 60), d("J10", 6, 12), d("P02", 3, 8), d("J13", 8, 14), d("C01", 20, 45)},
		recovery: []dose{d("J12", 20, 30), d("BL01", 15, 20), d("C12", 6, 8)},
		final:    set("J12", 60, "J13", 14, "J15", 45, "P01", 12, "J10", 12, "C01", 45),
	}},
	{Key: "jumpsports", Slug: "basketball-and-volleyball-training", Category: CatSport, plan: coursePlan{
		warmup: sportWarmup, cooldown: set("ST16", 30, "ST12", 30, "ST04", 30, "ST13", 30),
		a:        []dose{d("J05", 5, 10), d("L10", 10, 20), d("BL01", 15, 30), d("L15", 6, 12), d("C01", 20, 40)},
		b:        []dose{d("J10", 6, 12), d("L04", 5, 10), d("C07", 15, 30), d("L06", 10, 16), ds("J11", 3, 8, 6), ds("J14", 3, 6, 12)},
		c:        []dose{d("J15", 20, 40), d("L13", 6, 10), d("J05", 6, 12), d("BL03", 4, 8), d("L07", 12, 24)},
		recovery: []dose{d("BL01", 15, 20), d("L06", 10, 14), d("C12", 6, 8)},
		final:    set("J05", 12, "J11", 8, "J10", 12, "J14", 6, "L15", 12, "C01", 40),
	}},
}

func init() {
	var text map[string]struct {
		Title map[string]string `json:"title"`
		About map[string]string `json:"about"`
	}
	if err := json.Unmarshal(moreCoursesText, &text); err != nil {
		panic(fmt.Sprintf("content: courses_more.json: %v", err))
	}
	for _, c := range moreCourses {
		t, ok := text[c.Key]
		if !ok {
			panic("content: courses_more.json lacks " + c.Key)
		}
		c.Title, c.About = t.Title, t.About
		series30 = append(series30, c)
	}
	addExtraCourses()
}
