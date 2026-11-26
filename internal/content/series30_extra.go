package content

import (
	_ "embed"
	"encoding/json"
	"fmt"
)

// Even more ready programs: nine more sports, three for common back and foot
// problems, and eight more for children. Titles and descriptions in every
// language live in courses_extra.json.

//go:embed courses_extra.json
var extraCoursesText []byte

var (
	swimWarmup = set("J02", 30, "W11", 20, "W01", 15, "W09", 20, "W03", 6, "ST15", 6)
	backWarmup = set("W07", 20, "ST15", 6, "W04", 6, "NK04", 8, "W10", 6)
	kidsBallUp = set("IT01", 40, "IT78", 20, "IT04", 10, "IT02", 10, "IT06", 8)
)

var extraCourses = []Course{
	// ---- Sports ----
	{Key: "swim", Slug: "dryland-training-for-swimmers", Category: CatSport, plan: coursePlan{
		warmup: swimWarmup, cooldown: set("ST17", 30, "ST25", 30, "ST27", 30, "CD08", 40),
		a:        []dose{d("B13", 15, 40), d("C06", 15, 30), d("B07", 20, 40), d("P01", 4, 10), d("L06", 10, 16)},
		b:        []dose{d("B14", 8, 14), d("C04", 15, 30), d("B12", 6, 12), d("J05", 5, 10), d("C07", 15, 30)},
		c:        []dose{d("B13", 20, 45), d("C16", 6, 10), d("B03", 8, 14), d("P08", 6, 12), d("L10", 10, 20)},
		recovery: []dose{d("B12", 6, 8), d("C16", 6, 8), d("B07", 20, 25)},
		final:    set("B13", 45, "B14", 14, "C06", 30, "B07", 40, "P01", 10, "J05", 10),
	}},
	{Key: "wrestling", Slug: "training-for-wrestling-and-judo", Category: CatSport, plan: coursePlan{
		warmup: sportWarmup, cooldown: set("ST08", 30, "ST03", 30, "ST23", 30, "CD07", 30),
		a:        []dose{d("J16", 4, 10), d("J17", 4, 8), d("NK08", 20, 40), d("C18", 15, 35), d("L06", 10, 16)},
		b:        []dose{d("L17", 4, 8), d("J16", 5, 10), d("NK09", 15, 30), d("C10", 4, 10), d("L04", 5, 10)},
		c:        []dose{d("J17", 5, 10), d("P01", 5, 12), d("NK08", 20, 40), d("C07", 15, 30), d("L10", 10, 20)},
		recovery: []dose{d("NK08", 15, 20), d("C16", 6, 8), d("L06", 10, 12)},
		final:    set("J16", 10, "J17", 8, "L17", 8, "C18", 35, "NK09", 30, "P01", 12),
	}},
	{Key: "cycling", Slug: "training-for-cyclists", Category: CatSport, plan: coursePlan{
		warmup:   set("J02", 30, "W04", 8, "W05", 8, "W10", 6, "W08", 20, "ST15", 6),
		cooldown: set("ST04", 30, "ST12", 30, "ST28", 30, "NK06", 30),
		a:        []dose{d("L18", 6, 12), d("L04", 5, 10), d("C16", 6, 10), d("L14", 20, 50), d("B15", 8, 14)},
		b:        []dose{d("L10", 10, 20), d("L05", 6, 12), d("C17", 15, 30), d("L07", 12, 24), d("B12", 6, 12)},
		c:        []dose{d("L11", 6, 12), d("L06", 10, 16), d("C01", 20, 45), d("L18", 8, 14), d("B04", 8, 14)},
		recovery: []dose{d("L06", 10, 14), d("C16", 6, 8), d("B12", 6, 8)},
		final:    set("L18", 14, "L10", 20, "L14", 50, "C01", 45, "B15", 14, "L05", 12),
	}},
	{Key: "hockey", Slug: "hockey-and-skating-training", Category: CatSport, plan: coursePlan{
		warmup: sportWarmup, cooldown: set("ST04", 30, "ST08", 30, "ST13", 30, "ST16", 30),
		a:        []dose{d("J06", 15, 30), d("L17", 4, 8), d("J18", 15, 30), d("C19", 10, 20), d("L04", 5, 10)},
		b:        []dose{d("L16", 6, 12), d("BL01", 15, 35), d("J10", 6, 12), d("C07", 15, 30), d("L14", 20, 45)},
		c:        []dose{d("J18", 20, 40), d("L13", 6, 10), d("J05", 5, 10), d("C18", 15, 30), d("L07", 12, 24)},
		recovery: []dose{d("BL01", 15, 25), d("L06", 10, 14), d("C12", 6, 8)},
		final:    set("J06", 30, "L17", 8, "J18", 40, "L16", 12, "C19", 20, "L14", 45),
	}},
	{Key: "climbing", Slug: "training-for-climbers", Category: CatSport, plan: coursePlan{
		warmup:   set("J02", 30, "W01", 15, "W09", 20, "NK04", 8, "W03", 6, "W06", 20),
		cooldown: set("NK07", 30, "ST17", 30, "ST25", 30, "ST18", 30),
		a:        []dose{d("PU01", 10, 30), d("PU02", 4, 8), d("A07", 15, 30), d("C06", 15, 30), d("P08", 6, 10)},
		b:        []dose{d("PU08", 3, 8), d("B12", 6, 12), d("A06", 6, 10), d("C18", 15, 30), d("L04", 5, 10)},
		c:        []dose{d("PU04", 10, 25), d("PU07", 5, 10), d("B05", 8, 14), d("C07", 15, 30), d("A07", 20, 40)},
		recovery: []dose{d("PU01", 10, 20), d("A06", 6, 8), d("B12", 6, 8)},
		final:    set("PU01", 30, "PU04", 25, "PU08", 8, "A07", 40, "C18", 30, "B05", 14),
	}},
	{Key: "golf", Slug: "training-for-golfers", Category: CatSport, plan: coursePlan{
		warmup:   set("J02", 30, "W14", 6, "W03", 6, "W04", 8, "W01", 15, "W10", 6),
		cooldown: set("ST23", 30, "ST28", 30, "ST17", 30, "ST12", 30),
		a:        []dose{d("BL06", 15, 30), d("C16", 6, 10), d("L03", 8, 14), d("B15", 8, 14), d("C07", 15, 30)},
		b:        []dose{d("L04", 5, 10), d("C19", 8, 16), d("B12", 6, 12), d("L13", 4, 8), d("BL01", 15, 35)},
		c:        []dose{d("L11", 6, 12), d("C17", 15, 30), d("BL03", 4, 8), d("B04", 6, 12), d("L06", 10, 16)},
		recovery: []dose{d("C16", 6, 8), d("BL01", 15, 20), d("B12", 6, 8)},
		final:    set("BL06", 30, "C19", 16, "L11", 12, "B15", 14, "C07", 30, "BL03", 8),
	}},
	{Key: "racket", Slug: "padel-badminton-table-tennis-training", Category: CatSport, plan: coursePlan{
		warmup: sportWarmup, cooldown: set("ST17", 30, "ST04", 30, "ST13", 30, "NK07", 30),
		a:        []dose{d("J19", 6, 12), d("L20", 5, 10), d("J18", 15, 30), d("B03", 8, 14), d("C07", 15, 30)},
		b:        []dose{d("L13", 6, 10), d("J10", 6, 12), d("B05", 6, 12), d("A06", 6, 10), d("C12", 6, 10)},
		c:        []dose{d("J19", 8, 14), d("L16", 6, 12), d("J06", 15, 30), d("B04", 6, 12), d("C09", 6, 14)},
		recovery: []dose{d("B03", 8, 10), d("BL01", 15, 20), d("C12", 6, 8)},
		final:    set("J19", 14, "L20", 10, "J18", 30, "B05", 12, "C09", 14, "J06", 30),
	}},
	{Key: "hiking", Slug: "hiking-and-trekking-training", Category: CatSport, plan: coursePlan{
		warmup:   set("J02", 30, "W10", 6, "W04", 8, "W05", 8, "W12", 20, "W13", 20),
		cooldown: set("ST16", 30, "ST12", 30, "ST13", 30, "ST04", 30),
		a:        []dose{d("L18", 6, 12), d("L19", 8, 16), d("BL01", 15, 35), d("C16", 6, 10), d("L07", 12, 24)},
		b:        []dose{d("L21", 4, 8), d("L11", 6, 12), d("B09", 8, 14), d("C17", 15, 30), d("L15", 6, 12)},
		c:        []dose{d("L14", 20, 50), d("L18", 8, 14), d("BL02", 20, 40), d("L05", 6, 12), d("B12", 6, 12)},
		recovery: []dose{d("L06", 10, 14), d("BL01", 15, 25), d("L19", 8, 10)},
		final:    set("L18", 14, "L21", 8, "L14", 50, "L19", 16, "BL01", 35, "C16", 10),
	}},
	{Key: "rowing", Slug: "rowing-kayak-sup-training", Category: CatSport, plan: coursePlan{
		warmup:   set("J02", 30, "W09", 20, "W08", 20, "W04", 8, "ST15", 6, "W01", 15),
		cooldown: set("ST25", 30, "ST01", 30, "ST13", 30, "NK06", 30),
		a:        []dose{d("B15", 8, 14), d("L03", 8, 14), d("C19", 8, 16), d("L10", 10, 20), d("C01", 20, 45)},
		b:        []dose{d("B14", 8, 14), d("L06", 10, 16), d("C18", 15, 30), d("B12", 6, 12), d("A07", 15, 30)},
		c:        []dose{d("B06", 8, 14), d("L04", 5, 10), d("C13", 15, 30), d("B13", 15, 30), d("P01", 4, 10)},
		recovery: []dose{d("B12", 6, 8), d("C16", 6, 8), d("L06", 10, 12)},
		final:    set("B15", 14, "C19", 16, "L10", 20, "B14", 14, "C18", 30, "C01", 45),
	}},

	// ---- Back and feet ----
	{Key: "backcare", Slug: "back-care-for-disc-protrusions", Category: CatBackBalance, plan: coursePlan{
		warmup: backWarmup, cooldown: set("ST24", 30, "ST13", 30, "ST04", 30, "CD08", 40),
		a:        []dose{d("C15", 4, 8), d("C16", 4, 8), d("C17", 10, 25), d("L06", 8, 14), d("B03", 6, 12)},
		b:        []dose{d("C12", 4, 8), d("L03", 6, 12), d("B12", 5, 10), d("BL01", 15, 30), d("NK02", 6, 10)},
		c:        []dose{d("C15", 5, 10), d("C16", 5, 10), d("L04", 4, 8), d("B01", 10, 25), d("L07", 10, 20)},
		recovery: []dose{d("C16", 4, 6), d("L06", 8, 10), d("NK02", 6, 8)},
		final:    set("C15", 10, "C16", 10, "C17", 25, "L06", 14, "B03", 12, "BL01", 30),
	}},
	{Key: "scoliosis", Slug: "back-symmetry-for-scoliosis", Category: CatBackBalance, plan: coursePlan{
		warmup:   set("W07", 20, "ST15", 6, "W01", 15, "NK04", 8, "W04", 6),
		cooldown: set("ST19", 30, "ST20", 30, "ST22", 30, "CD08", 40),
		a:        []dose{d("C16", 5, 10), d("C17", 10, 30), d("B12", 6, 12), d("C12", 5, 10), d("L06", 8, 14)},
		b:        []dose{d("B04", 6, 12), d("B05", 6, 12), d("C07", 10, 30), d("BL01", 15, 30), d("NK02", 6, 10)},
		c:        []dose{d("B13", 10, 30), d("C16", 6, 10), d("L04", 4, 8), d("B03", 6, 12), d("C17", 15, 30)},
		recovery: []dose{d("C16", 4, 6), d("B12", 5, 8), d("L06", 8, 10)},
		final:    set("C16", 10, "C17", 30, "B12", 12, "B13", 30, "BL01", 30, "C12", 10),
	}},
	{Key: "feet", Slug: "healthy-feet-flat-feet", Category: CatBackBalance, plan: coursePlan{
		warmup:   set("W10", 6, "W12", 20, "W13", 20, "W05", 6, "W04", 6),
		cooldown: set("ST26", 30, "ST16", 30, "ST13", 30, "ST12", 30),
		a:        []dose{d("L23", 10, 30), d("L24", 20, 40), d("L07", 10, 20), d("L25", 6, 12), d("BL01", 15, 30)},
		b:        []dose{d("L19", 8, 16), d("L26", 8, 16), d("L15", 5, 10), d("L23", 15, 30), d("L06", 8, 14)},
		c:        []dose{d("L24", 20, 45), d("L25", 8, 14), d("BL02", 20, 40), d("L07", 12, 24), d("L04", 4, 8)},
		recovery: []dose{d("L23", 10, 20), d("L26", 8, 10), d("BL01", 15, 20)},
		final:    set("L23", 30, "L24", 40, "L19", 16, "L25", 14, "BL01", 30, "L15", 10),
	}},

	// ---- Children ----
	{Key: "kidsswim", Slug: "swimming-for-kids", Category: CatKids, Kids: true, plan: coursePlan{
		warmup: set("IT01", 40, "IT56", 20, "IT04", 10, "IT78", 20, "IT02", 10), cooldown: kidsCool,
		a:        []dose{d("IT55", 10, 25), d("IT57", 10, 25), d("IT59", 10, 20), d("IT12", 5, 8), d("IT40", 4, 8)},
		b:        []dose{d("IT58", 10, 20), d("IT60", 6, 12), d("IT46", 6, 10), d("IT27", 5, 8), d("IT37", 20, 40)},
		c:        []dose{d("IT55", 15, 30), d("IT57", 15, 30), d("IT39", 6, 12), d("IT10", 6, 12), d("IT76", 4, 8)},
		recovery: []dose{d("IT55", 10, 15), d("IT60", 6, 8), d("IT17", 10, 20)},
		final:    set("IT55", 30, "IT57", 30, "IT58", 20, "IT60", 12, "IT59", 20, "IT39", 12),
	}},
	{Key: "kidswrestling", Slug: "wrestling-and-judo-for-kids", Category: CatKids, Kids: true, plan: coursePlan{
		warmup: kidsWarmup, cooldown: kidsCool,
		a:        []dose{d("IT63", 5, 10), d("IT65", 15, 30), d("IT38", 5, 10), d("IT28", 10, 25), d("IT17", 10, 20)},
		b:        []dose{d("IT64", 3, 6), d("IT66", 15, 30), d("IT09", 3, 8), d("IT27", 5, 8), d("IT70", 10, 25)},
		c:        []dose{d("IT63", 6, 12), d("IT65", 20, 40), d("IT47", 4, 8), d("IT39", 6, 12), d("IT36", 15, 30)},
		recovery: []dose{d("IT63", 4, 6), d("IT17", 10, 20), d("IT20", 6, 8)},
		final:    set("IT63", 12, "IT64", 6, "IT65", 40, "IT66", 30, "IT09", 8, "IT70", 25),
	}},
	{Key: "kidskarate", Slug: "karate-and-taekwondo-for-kids", Category: CatKids, Kids: true, plan: coursePlan{
		warmup: kidsWarmup, cooldown: kidsCool,
		a:        []dose{d("IT67", 6, 12), d("IT68", 4, 8), d("IT70", 10, 25), d("IT17", 10, 25), d("IT12", 5, 8)},
		b:        []dose{d("IT69", 5, 10), d("IT36", 15, 30), d("IT61", 10, 20), d("IT28", 10, 25), d("IT38", 5, 10)},
		c:        []dose{d("IT68", 5, 10), d("IT67", 8, 14), d("IT70", 15, 30), d("IT39", 6, 12), d("IT45", 4, 8)},
		recovery: []dose{d("IT17", 10, 20), d("IT67", 6, 8), d("IT20", 6, 8)},
		final:    set("IT67", 14, "IT68", 10, "IT69", 10, "IT70", 30, "IT61", 20, "IT39", 12),
	}},
	{Key: "kidsgym", Slug: "gymnastics-for-kids", Category: CatKids, Kids: true, plan: coursePlan{
		warmup: kidsWarmup, cooldown: kidsCool,
		a:        []dose{d("IT58", 10, 20), d("IT59", 10, 20), d("IT61", 10, 20), d("IT62", 10, 20), d("IT38", 5, 10)},
		b:        []dose{d("IT73", 6, 12), d("IT17", 10, 25), d("IT66", 15, 30), d("IT40", 4, 8), d("IT21", 4, 8)},
		c:        []dose{d("IT58", 15, 25), d("IT62", 10, 25), d("IT44", 6, 12), d("IT61", 15, 25), d("IT28", 10, 25)},
		recovery: []dose{d("IT17", 10, 20), d("IT20", 6, 8), d("IT22", 6, 8)},
		final:    set("IT58", 25, "IT59", 20, "IT61", 25, "IT62", 25, "IT73", 12, "IT66", 30),
	}},
	{Key: "kidsdance", Slug: "dance-and-rhythm-for-kids", Category: CatKids, Kids: true, plan: coursePlan{
		warmup: set("IT01", 40, "IT04", 10, "IT78", 20, "IT02", 10, "IT05", 8), cooldown: kidsCool,
		a:        []dose{d("IT71", 20, 40), d("IT73", 6, 12), d("IT72", 15, 30), d("IT17", 10, 25), d("IT11", 5, 10)},
		b:        []dose{d("IT31", 15, 30), d("IT30", 6, 12), d("IT61", 10, 20), d("IT44", 6, 12), d("IT12", 5, 8)},
		c:        []dose{d("IT71", 20, 40), d("IT72", 20, 40), d("IT73", 8, 14), d("IT29", 8, 16), d("IT28", 10, 25)},
		recovery: []dose{d("IT71", 15, 20), d("IT17", 10, 20), d("IT20", 6, 8)},
		final:    set("IT71", 40, "IT72", 40, "IT73", 14, "IT31", 30, "IT61", 20, "IT17", 25),
	}},
	{Key: "kidshockey", Slug: "hockey-and-skating-for-kids", Category: CatKids, Kids: true, plan: coursePlan{
		warmup: kidsWarmup, cooldown: kidsCool,
		a:        []dose{d("IT74", 4, 8), d("IT75", 10, 25), d("IT15", 20, 40), d("IT17", 10, 25), d("IT28", 10, 25)},
		b:        []dose{d("IT42", 4, 8), d("IT34", 15, 30), d("IT11", 5, 10), d("IT39", 6, 12), d("IT47", 4, 8)},
		c:        []dose{d("IT74", 6, 10), d("IT75", 15, 30), d("IT36", 15, 30), d("IT45", 4, 8), d("IT12", 5, 8)},
		recovery: []dose{d("IT17", 10, 20), d("IT75", 10, 15), d("IT20", 6, 8)},
		final:    set("IT74", 10, "IT75", 30, "IT42", 8, "IT34", 30, "IT45", 8, "IT17", 25),
	}},
	{Key: "kidsball", Slug: "basketball-and-volleyball-for-kids", Category: CatKids, Kids: true, plan: coursePlan{
		warmup: kidsBallUp, cooldown: kidsCool,
		a:        []dose{d("IT76", 4, 8), d("IT34", 15, 30), d("IT77", 6, 12), d("IT16", 6, 10), d("IT28", 10, 25)},
		b:        []dose{d("IT24", 4, 8), d("IT11", 5, 10), d("IT36", 15, 30), d("IT39", 6, 12), d("IT38", 5, 10)},
		c:        []dose{d("IT76", 5, 10), d("IT43", 15, 30), d("IT77", 8, 14), d("IT30", 6, 12), d("IT12", 5, 8)},
		recovery: []dose{d("IT17", 10, 20), d("IT77", 6, 8), d("IT20", 6, 8)},
		final:    set("IT76", 10, "IT24", 8, "IT77", 14, "IT34", 30, "IT36", 30, "IT39", 12),
	}},
	{Key: "kidsfeet", Slug: "flat-feet-exercises-for-kids", Category: CatKids, Kids: true, plan: coursePlan{
		warmup: set("IT01", 40, "IT79", 20, "IT80", 20, "IT81", 20, "IT04", 10), cooldown: kidsCool,
		a:        []dose{d("IT82", 15, 30), d("IT85", 6, 12), d("IT86", 6, 12), d("IT84", 10, 20), d("IT17", 10, 20)},
		b:        []dose{d("IT83", 4, 8), d("IT11", 5, 10), d("IT79", 20, 40), d("IT85", 6, 12), d("IT12", 5, 8)},
		c:        []dose{d("IT82", 20, 40), d("IT84", 15, 25), d("IT80", 20, 40), d("IT86", 8, 14), d("IT17", 15, 25)},
		recovery: []dose{d("IT85", 6, 8), d("IT84", 10, 15), d("IT17", 10, 15)},
		final:    set("IT82", 40, "IT83", 8, "IT85", 12, "IT86", 14, "IT84", 25, "IT17", 25),
	}},
}

// addExtraCourses appends these programs after the ones in series30_more.go;
// it is called from there so the catalog order never depends on file names.
func addExtraCourses() {
	var text map[string]struct {
		Title map[string]string `json:"title"`
		About map[string]string `json:"about"`
	}
	if err := json.Unmarshal(extraCoursesText, &text); err != nil {
		panic(fmt.Sprintf("content: courses_extra.json: %v", err))
	}
	for _, c := range extraCourses {
		t, ok := text[c.Key]
		if !ok {
			panic("content: courses_extra.json lacks " + c.Key)
		}
		c.Title, c.About = t.Title, t.About
		series30 = append(series30, c)
	}
}
