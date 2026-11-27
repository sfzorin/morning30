// Managed by GoX v0.1.36

//line fitness.gox:1
package pages

import (
	"context"
	"log"
	"fmt"
	"strconv"
	"strings"

	"morning30/assets"
	"morning30/internal/app"
	"morning30/internal/auth"
	"morning30/internal/content"
	"morning30/internal/fitness"
	"morning30/internal/i18n"
	"morning30/path"

	"github.com/doors-dev/doors"
	"github.com/doors-dev/gox"
)

// The fitness profile: "About you", the voice-guided self-test and its result
// with the recommended program. A newcomer lands here right after signing up
// or choosing to continue without an account; later it opens from Settings.

// fitnessSummary is the value of the Settings row: "57 · Good", or nothing.
func fitnessSummary(l i18n.Lang, userID int64) string {
	latest, _, ok, _ := app.LatestFitness(userID)
	if !ok {
		return ""
	}
	return strconv.Itoa(latest.Index) + " · " + i18n.T(l, "fitness.level_"+latest.Level)
}

// aboutForm is the decoded "About you" form.
type aboutForm struct {
	Who      string `form:"who"`
	Sex      string `form:"sex"`
	Age      string `form:"age"`
	Height   string `form:"height"`
	Weight   string `form:"weight"`
	Back     string `form:"c_back"`
	Knees    string `form:"c_knees"`
	Pressure string `form:"c_pressure"`
	Feet     string `form:"c_feet"`
}

type aboutPage struct {
	sess auth.Session
	path doors.Source[path.Path]
}

func numberValue(v int) string {
	if v <= 0 {
		return ""
	}
	return strconv.Itoa(v)
}

//line fitness.gox:60
func (a aboutPage) Main() gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
//line fitness.gox:62
		l := i18n.Lang(a.sess.Lang)
		p, _ := app.FitnessProfile(a.sess.UserID)
		female := p.Female()
		concerns := []struct{ Key, Label string }{
			{fitness.Back, "fitness.c_back"}, {fitness.Knees, "fitness.c_knees"},
			{fitness.Pressure, "fitness.c_pressure"}, {fitness.Feet, "fitness.c_feet"},
		}

		__e = __c.Init("title"); if __e != nil { return }
		{
			__e = __c.Submit(); if __e != nil { return }
//line fitness.gox:70
			__e = __c.Any(i18n.T(l, "fitness.about")); if __e != nil { return }
		}
		__e = __c.Close(); if __e != nil { return }
		__e = __c.Init("main"); if __e != nil { return }
		{
//line fitness.gox:71
			__e = __c.Set("class", "screen"); if __e != nil { return }
			__e = __c.Submit(); if __e != nil { return }
			__e = __c.Init("section"); if __e != nil { return }
			{
//line fitness.gox:72
				__e = __c.Set("class", "fit about"); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
				__e = __c.Init("h1"); if __e != nil { return }
				{
					__e = __c.Submit(); if __e != nil { return }
//line fitness.gox:73
					__e = __c.Any(i18n.T(l, "fitness.about")); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
				__e = __c.Init("p"); if __e != nil { return }
				{
//line fitness.gox:74
					__e = __c.Set("class", "fit-lead"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
//line fitness.gox:74
					__e = __c.Any(i18n.T(l, "fitness.lead")); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
//line fitness.gox:75
				__e = (doors.ASubmit[aboutForm]{On: a.submit}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
					ctx := __c.Context(); _ = ctx
					__e = __c.Init("form"); if __e != nil { return }
					{
//line fitness.gox:75
						__e = __c.Set("class", "fit-form"); if __e != nil { return }
						__e = __c.Set("novalidate", true); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
						__e = __c.Init("p"); if __e != nil { return }
						{
//line fitness.gox:76
							__e = __c.Set("class", "fit-label"); if __e != nil { return }
							__e = __c.Submit(); if __e != nil { return }
//line fitness.gox:76
							__e = __c.Any(i18n.T(l, "fitness.who")); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
						__e = __c.Init("div"); if __e != nil { return }
						{
//line fitness.gox:77
							__e = __c.Set("class", "seg"); if __e != nil { return }
							__e = __c.Submit(); if __e != nil { return }
							__e = __c.Init("label"); if __e != nil { return }
							{
								__e = __c.Submit(); if __e != nil { return }
								__e = __c.InitVoid("input"); if __e != nil { return }
								{
//line fitness.gox:78
									__e = __c.Set("type", "radio"); if __e != nil { return }
//line fitness.gox:78
									__e = __c.Set("name", "who"); if __e != nil { return }
//line fitness.gox:78
									__e = __c.Set("value", "self"); if __e != nil { return }
//line fitness.gox:78
									__e = __c.Set("checked", !p.Child); if __e != nil { return }
								}
								__e = __c.Submit(); if __e != nil { return }
								__e = __c.Init("span"); if __e != nil { return }
								{
									__e = __c.Submit(); if __e != nil { return }
//line fitness.gox:78
									__e = __c.Any(i18n.T(l, "fitness.self")); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
							}
							__e = __c.Close(); if __e != nil { return }
							__e = __c.Init("label"); if __e != nil { return }
							{
								__e = __c.Submit(); if __e != nil { return }
								__e = __c.InitVoid("input"); if __e != nil { return }
								{
//line fitness.gox:79
									__e = __c.Set("type", "radio"); if __e != nil { return }
//line fitness.gox:79
									__e = __c.Set("name", "who"); if __e != nil { return }
//line fitness.gox:79
									__e = __c.Set("value", "child"); if __e != nil { return }
//line fitness.gox:79
									__e = __c.Set("checked", p.Child); if __e != nil { return }
								}
								__e = __c.Submit(); if __e != nil { return }
								__e = __c.Init("span"); if __e != nil { return }
								{
									__e = __c.Submit(); if __e != nil { return }
//line fitness.gox:79
									__e = __c.Any(i18n.T(l, "fitness.child")); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
							}
							__e = __c.Close(); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
						__e = __c.Init("p"); if __e != nil { return }
						{
//line fitness.gox:81
							__e = __c.Set("class", "fit-label"); if __e != nil { return }
							__e = __c.Submit(); if __e != nil { return }
//line fitness.gox:81
							__e = __c.Any(i18n.T(l, "settings.sex")); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
						__e = __c.Init("div"); if __e != nil { return }
						{
//line fitness.gox:82
							__e = __c.Set("class", "figure-grid figure-profile"); if __e != nil { return }
//line fitness.gox:82
							__e = __c.Set("role", "group"); if __e != nil { return }
//line fitness.gox:82
							__e = __c.Set("aria-label", i18n.T(l, "settings.sex")); if __e != nil { return }
							__e = __c.Submit(); if __e != nil { return }
//line fitness.gox:83
							for i, adult := range []string{"m", "f"} {
								__e = __c.Init("label"); if __e != nil { return }
								{
//line fitness.gox:84
									__e = __c.Set("class", "figure-option"); if __e != nil { return }
									__e = __c.Submit(); if __e != nil { return }
									__e = __c.InitVoid("input"); if __e != nil { return }
									{
//line fitness.gox:85
										__e = __c.Set("class", "figure-radio"); if __e != nil { return }
//line fitness.gox:85
										__e = __c.Set("type", "radio"); if __e != nil { return }
//line fitness.gox:85
										__e = __c.Set("name", "sex"); if __e != nil { return }
//line fitness.gox:85
										__e = __c.Set("value", strconv.Itoa(i+1)); if __e != nil { return }
//line fitness.gox:85
										__e = __c.Set("checked", (i == 1) == female); if __e != nil { return }
									}
									__e = __c.Submit(); if __e != nil { return }
									__e = __c.Init("span"); if __e != nil { return }
									{
//line fitness.gox:86
										__e = __c.Set("class", "figure-choice figure-adult"); if __e != nil { return }
										__e = __c.Submit(); if __e != nil { return }
//line fitness.gox:86
										__e = __c.Any(figureTile(l, adult)); if __e != nil { return }
									}
									__e = __c.Close(); if __e != nil { return }
									__e = __c.Init("span"); if __e != nil { return }
									{
//line fitness.gox:87
										__e = __c.Set("class", "figure-choice figure-child"); if __e != nil { return }
										__e = __c.Submit(); if __e != nil { return }
//line fitness.gox:87
										__e = __c.Any(figureTile(l, app.FigureFor(adult, true))); if __e != nil { return }
									}
									__e = __c.Close(); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
							}
						}
						__e = __c.Close(); if __e != nil { return }
						__e = __c.Init("div"); if __e != nil { return }
						{
//line fitness.gox:91
							__e = __c.Set("class", "fit-nums"); if __e != nil { return }
							__e = __c.Submit(); if __e != nil { return }
							__e = __c.Init("label"); if __e != nil { return }
							{
								__e = __c.Submit(); if __e != nil { return }
								__e = __c.Init("span"); if __e != nil { return }
								{
//line fitness.gox:92
									__e = __c.Set("class", "fit-label"); if __e != nil { return }
									__e = __c.Submit(); if __e != nil { return }
//line fitness.gox:92
									__e = __c.Any(i18n.T(l, "fitness.age")); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
								__e = __c.InitVoid("input"); if __e != nil { return }
								{
//line fitness.gox:92
									__e = __c.Set("class", "text-input"); if __e != nil { return }
//line fitness.gox:92
									__e = __c.Set("name", "age"); if __e != nil { return }
//line fitness.gox:92
									__e = __c.Set("type", "number"); if __e != nil { return }
//line fitness.gox:92
									__e = __c.Set("inputmode", "numeric"); if __e != nil { return }
//line fitness.gox:92
									__e = __c.Set("min", "3"); if __e != nil { return }
//line fitness.gox:92
									__e = __c.Set("max", "110"); if __e != nil { return }
//line fitness.gox:92
									__e = __c.Set("value", numberValue(p.Age)); if __e != nil { return }
								}
								__e = __c.Submit(); if __e != nil { return }
							}
							__e = __c.Close(); if __e != nil { return }
							__e = __c.Init("label"); if __e != nil { return }
							{
								__e = __c.Submit(); if __e != nil { return }
								__e = __c.Init("span"); if __e != nil { return }
								{
//line fitness.gox:93
									__e = __c.Set("class", "fit-label"); if __e != nil { return }
									__e = __c.Submit(); if __e != nil { return }
//line fitness.gox:93
									__e = __c.Any(i18n.T(l, "fitness.height")); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
								__e = __c.InitVoid("input"); if __e != nil { return }
								{
//line fitness.gox:93
									__e = __c.Set("class", "text-input"); if __e != nil { return }
//line fitness.gox:93
									__e = __c.Set("name", "height"); if __e != nil { return }
//line fitness.gox:93
									__e = __c.Set("type", "number"); if __e != nil { return }
//line fitness.gox:93
									__e = __c.Set("inputmode", "numeric"); if __e != nil { return }
//line fitness.gox:93
									__e = __c.Set("min", "80"); if __e != nil { return }
//line fitness.gox:93
									__e = __c.Set("max", "250"); if __e != nil { return }
//line fitness.gox:93
									__e = __c.Set("value", numberValue(p.Height)); if __e != nil { return }
								}
								__e = __c.Submit(); if __e != nil { return }
							}
							__e = __c.Close(); if __e != nil { return }
							__e = __c.Init("label"); if __e != nil { return }
							{
								__e = __c.Submit(); if __e != nil { return }
								__e = __c.Init("span"); if __e != nil { return }
								{
//line fitness.gox:94
									__e = __c.Set("class", "fit-label"); if __e != nil { return }
									__e = __c.Submit(); if __e != nil { return }
//line fitness.gox:94
									__e = __c.Any(i18n.T(l, "fitness.weight")); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
								__e = __c.InitVoid("input"); if __e != nil { return }
								{
//line fitness.gox:94
									__e = __c.Set("class", "text-input"); if __e != nil { return }
//line fitness.gox:94
									__e = __c.Set("name", "weight"); if __e != nil { return }
//line fitness.gox:94
									__e = __c.Set("type", "number"); if __e != nil { return }
//line fitness.gox:94
									__e = __c.Set("inputmode", "numeric"); if __e != nil { return }
//line fitness.gox:94
									__e = __c.Set("min", "15"); if __e != nil { return }
//line fitness.gox:94
									__e = __c.Set("max", "300"); if __e != nil { return }
//line fitness.gox:94
									__e = __c.Set("value", numberValue(p.Weight)); if __e != nil { return }
								}
								__e = __c.Submit(); if __e != nil { return }
							}
							__e = __c.Close(); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
						__e = __c.Init("p"); if __e != nil { return }
						{
//line fitness.gox:96
							__e = __c.Set("class", "fit-label"); if __e != nil { return }
							__e = __c.Submit(); if __e != nil { return }
//line fitness.gox:96
							__e = __c.Any(i18n.T(l, "fitness.concerns")); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
						__e = __c.Init("div"); if __e != nil { return }
						{
//line fitness.gox:97
							__e = __c.Set("class", "chips"); if __e != nil { return }
							__e = __c.Submit(); if __e != nil { return }
//line fitness.gox:98
							for _, c := range concerns {
								__e = __c.Init("label"); if __e != nil { return }
								{
//line fitness.gox:99
									__e = __c.Set("class", "fit-chip"); if __e != nil { return }
									__e = __c.Submit(); if __e != nil { return }
									__e = __c.InitVoid("input"); if __e != nil { return }
									{
//line fitness.gox:99
										__e = __c.Set("type", "checkbox"); if __e != nil { return }
//line fitness.gox:99
										__e = __c.Set("name", "c_" + c.Key); if __e != nil { return }
//line fitness.gox:99
										__e = __c.Set("value", "1"); if __e != nil { return }
//line fitness.gox:99
										__e = __c.Set("checked", p.Has(c.Key)); if __e != nil { return }
									}
									__e = __c.Submit(); if __e != nil { return }
									__e = __c.Init("span"); if __e != nil { return }
									{
										__e = __c.Submit(); if __e != nil { return }
//line fitness.gox:99
										__e = __c.Any(i18n.T(l, c.Label)); if __e != nil { return }
									}
									__e = __c.Close(); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
							}
						}
						__e = __c.Close(); if __e != nil { return }
						__e = __c.Init("button"); if __e != nil { return }
						{
//line fitness.gox:102
							__e = __c.Set("class", "btn primary fit-go"); if __e != nil { return }
//line fitness.gox:102
							__e = __c.Set("type", "submit"); if __e != nil { return }
							__e = __c.Submit(); if __e != nil { return }
//line fitness.gox:102
							__e = __c.Any(i18n.T(l, "fitness.next")); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
				return })); if __e != nil { return }
//line fitness.gox:104
				__e = (doors.ALink{Model: path.Path{Page: path.Home}}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
					ctx := __c.Context(); _ = ctx
					__e = __c.Init("a"); if __e != nil { return }
					{
//line fitness.gox:104
						__e = __c.Set("class", "fit-skip"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
//line fitness.gox:104
						__e = __c.Any(i18n.T(l, "workout.skip")); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
				return })); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
		}
		__e = __c.Close(); if __e != nil { return }
	return })
//line fitness.gox:107
}

func atoi(s string) int {
	n, _ := strconv.Atoi(strings.TrimSpace(s))
	return n
}

// submit saves the answers and opens the test.
func (a aboutPage) submit(ctx context.Context, r doors.RequestForm[aboutForm]) bool {
	f := r.Data()
	p := fitness.Profile{Child: f.Who == "child", Sex: "m", Age: atoi(f.Age), Height: atoi(f.Height), Weight: atoi(f.Weight)}
	if f.Sex == "2" {
		p.Sex = "f"
	}
	for key, on := range map[string]string{fitness.Back: f.Back, fitness.Knees: f.Knees, fitness.Pressure: f.Pressure, fitness.Feet: f.Feet} {
		if on != "" {
			p.Concerns = append(p.Concerns, key)
		}
	}
	if err := app.SaveFitnessProfile(a.sess.UserID, p); err != nil {
		log.Printf("about save: %v", err)
		return false
	}
	a.path.Update(ctx, path.Path{Page: path.Test})
	return false
}

// testStep is one exercise of the self-test as the page script sees it.
type testStep struct {
	ID     string   `json:"id"`   // push | squat | plank | pull
	Kind   string   `json:"kind"` // count | minute | hold
	Name   string   `json:"name"`
	Hint   string   `json:"hint"`
	Frames []string `json:"frames"`
}

type testPayload struct {
	BCP47  string            `json:"bcp47"`
	Voice  bool              `json:"voice"`
	Steps  []testStep        `json:"steps"`
	Labels map[string]string `json:"labels"`
}

type testPage struct {
	sess auth.Session
	path doors.Source[path.Path]
}

func (t testPage) payload() testPayload {
	l := i18n.Lang(t.sess.Lang)
	cat := app.UserCatalog(t.sess.UserID)
	sex := app.DB.GetSex(t.sess.UserID)
	step := func(id, kind, exercise, hint string) testStep {
		return testStep{ID: id, Kind: kind, Name: cat.Name(l, exercise), Hint: i18n.T(l, hint), Frames: app.ExerciseFrames(cat, exercise, sex)}
	}
	return testPayload{
		BCP47: i18n.BCP47(l),
		Voice: t.sess.VoiceMode != "off",
		Steps: []testStep{
			step("push", "count", "P01", "fitness.push_hint"),
			step("squat", "minute", "L10", "fitness.squat_hint"),
			step("plank", "hold", "C01", "fitness.plank_hint"),
			step("pull", "count", "PU06", "fitness.pull_hint"),
		},
		Labels: map[string]string{
			"test":  i18n.T(l, "fitness.test"),
			"done":  i18n.T(l, "workout.done"),
			"start": i18n.T(l, "fitness.start"),
			"stop":  i18n.T(l, "fitness.stop"),
			"back":  i18n.T(l, "workout.prev"),
			"skip":  i18n.T(l, "workout.skip"),
			"go":    i18n.Cue(l, "go"),
			"error": i18n.T(l, "fitness.save_failed"),
		},
	}
}

//line fitness.gox:184
func (t testPage) Main() gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
//line fitness.gox:186
		l := i18n.Lang(t.sess.Lang)

		__e = __c.Init("title"); if __e != nil { return }
		{
			__e = __c.Submit(); if __e != nil { return }
//line fitness.gox:188
			__e = __c.Any(i18n.T(l, "fitness.test")); if __e != nil { return }
		}
		__e = __c.Close(); if __e != nil { return }
		__e = __c.Init("div"); if __e != nil { return }
		{
//line fitness.gox:189
			__e = __c.Set("class", "player fit-test"); if __e != nil { return }
			__e = __c.Submit(); if __e != nil { return }
			__e = __c.Init("div"); if __e != nil { return }
			{
//line fitness.gox:190
				__e = __c.Set("class", "player-top"); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
//line fitness.gox:191
				__e = (doors.ALink{Model: path.Path{Page: path.Home}}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
					ctx := __c.Context(); _ = ctx
					__e = __c.Init("a"); if __e != nil { return }
					{
//line fitness.gox:191
						__e = __c.Set("class", "quit"); if __e != nil { return }
//line fitness.gox:191
						__e = __c.Set("aria-label", i18n.T(l, "nav.close")); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
//line fitness.gox:191
						__e = __c.Any(icon("close", false)); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
				return })); if __e != nil { return }
				__e = __c.Init("div"); if __e != nil { return }
				{
//line fitness.gox:192
					__e = __c.Set("class", "pbar fit-bar"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
				__e = __c.Init("div"); if __e != nil { return }
				{
//line fitness.gox:193
					__e = __c.Set("class", "phase"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
//line fitness.gox:193
					__e = __c.Any(i18n.T(l, "fitness.test")); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
			__e = __c.Init("div"); if __e != nil { return }
			{
//line fitness.gox:195
				__e = __c.Set("class", "stage"); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
				__e = __c.Init("div"); if __e != nil { return }
				{
//line fitness.gox:196
					__e = __c.Set("class", "ex-visual"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
					__e = __c.InitVoid("img"); if __e != nil { return }
					{
//line fitness.gox:196
						__e = __c.Set("class", "ex-svg"); if __e != nil { return }
//line fitness.gox:196
						__e = __c.Set("alt", ""); if __e != nil { return }
					}
					__e = __c.Submit(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
				__e = __c.Init("div"); if __e != nil { return }
				{
//line fitness.gox:197
					__e = __c.Set("class", "ex-name"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
				__e = __c.Init("p"); if __e != nil { return }
				{
//line fitness.gox:198
					__e = __c.Set("class", "fit-hint"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
				__e = __c.Init("div"); if __e != nil { return }
				{
//line fitness.gox:199
					__e = __c.Set("class", "fit-count"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
					__e = __c.Init("button"); if __e != nil { return }
					{
//line fitness.gox:200
						__e = __c.Set("class", "lvl-btn fit-minus"); if __e != nil { return }
//line fitness.gox:200
						__e = __c.Set("type", "button"); if __e != nil { return }
//line fitness.gox:200
						__e = __c.Set("aria-label", "−"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
//line fitness.gox:200
						__e = __c.Any(icon("minus", false)); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
					__e = __c.Init("output"); if __e != nil { return }
					{
//line fitness.gox:201
						__e = __c.Set("class", "fit-value"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
						__e = __c.Text("0"); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
					__e = __c.Init("button"); if __e != nil { return }
					{
//line fitness.gox:202
						__e = __c.Set("class", "lvl-btn fit-plus"); if __e != nil { return }
//line fitness.gox:202
						__e = __c.Set("type", "button"); if __e != nil { return }
//line fitness.gox:202
						__e = __c.Set("aria-label", "+"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
//line fitness.gox:202
						__e = __c.Any(icon("plus", false)); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
				__e = __c.Init("div"); if __e != nil { return }
				{
//line fitness.gox:204
					__e = __c.Set("class", "fit-clock hidden"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
				__e = __c.Init("p"); if __e != nil { return }
				{
//line fitness.gox:205
					__e = __c.Set("class", "err fit-err hidden"); if __e != nil { return }
//line fitness.gox:205
					__e = __c.Set("role", "alert"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
			__e = __c.Init("div"); if __e != nil { return }
			{
//line fitness.gox:207
				__e = __c.Set("class", "done-row"); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
				__e = __c.Init("button"); if __e != nil { return }
				{
//line fitness.gox:208
					__e = __c.Set("class", "btn primary fit-main"); if __e != nil { return }
//line fitness.gox:208
					__e = __c.Set("type", "button"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
			__e = __c.Init("div"); if __e != nil { return }
			{
//line fitness.gox:210
				__e = __c.Set("class", "controls"); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
				__e = __c.Init("button"); if __e != nil { return }
				{
//line fitness.gox:211
					__e = __c.Set("class", "btn fit-back"); if __e != nil { return }
//line fitness.gox:211
					__e = __c.Set("type", "button"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
				__e = __c.Init("button"); if __e != nil { return }
				{
//line fitness.gox:212
					__e = __c.Set("class", "btn fit-skip-step"); if __e != nil { return }
//line fitness.gox:212
					__e = __c.Set("type", "button"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
			__e = __c.Init("script"); if __e != nil { return }
			{
//line fitness.gox:214
				__e = __c.Set("src", assets.FitnessTest); if __e != nil { return }
				__e = __c.Set("inline", true); if __e != nil { return }
//line fitness.gox:214
				__e = __c.Set("data:payload", t.payload()); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
				__e = __c.Raw(""); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
		}
		__e = __c.Close(); if __e != nil { return }
	return })
//line fitness.gox:216
}

type fitnessPage struct {
	sess auth.Session
	path doors.Source[path.Path]
}

// clock writes seconds as m:ss.
func clock(sec int) string {
	return fmt.Sprintf("%d:%02d", sec/60, sec%60)
}

func resultValue(v int, seconds bool) string {
	if v < 0 {
		return "—"
	}
	if seconds {
		return clock(v)
	}
	return strconv.Itoa(v)
}

//line fitness.gox:238
func (f fitnessPage) Main() gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
//line fitness.gox:240
		l := i18n.Lang(f.sess.Lang)
		latest, previous, has, hasPrev := app.LatestFitness(f.sess.UserID)
		plan := app.FitnessPlan(f.sess.UserID)
		prog, _ := content.ReadyByKey(plan.Program)
		built := prog.Build()
		first, peak := content.MinutesSpan(built)
		rows := []struct {
			Key     string
			Value   int
			Seconds bool
		}{
			{"fitness.r_push", latest.Results.PushUps, false}, {"fitness.r_squat", latest.Results.Squats, false},
			{"fitness.r_plank", latest.Results.PlankSec, true}, {"fitness.r_pull", latest.Results.PullUps, false},
		}

		__e = __c.Init("title"); if __e != nil { return }
		{
			__e = __c.Submit(); if __e != nil { return }
//line fitness.gox:255
			__e = __c.Any(i18n.T(l, "fitness.title")); if __e != nil { return }
		}
		__e = __c.Close(); if __e != nil { return }
		__e = __c.Init("main"); if __e != nil { return }
		{
//line fitness.gox:256
			__e = __c.Set("class", "screen"); if __e != nil { return }
			__e = __c.Submit(); if __e != nil { return }
			__e = __c.Init("section"); if __e != nil { return }
			{
//line fitness.gox:257
				__e = __c.Set("class", "fit result"); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
				__e = __c.Init("h1"); if __e != nil { return }
				{
					__e = __c.Submit(); if __e != nil { return }
//line fitness.gox:258
					__e = __c.Any(i18n.T(l, "fitness.title")); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
//line fitness.gox:259
				if has {
					__e = __c.Init("div"); if __e != nil { return }
					{
//line fitness.gox:260
						__e = __c.Set("class", "fit-score"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
						__e = __c.Init("span"); if __e != nil { return }
						{
//line fitness.gox:261
							__e = __c.Set("class", "fit-index"); if __e != nil { return }
							__e = __c.Submit(); if __e != nil { return }
//line fitness.gox:261
							__e = __c.Any(latest.Index); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
						__e = __c.Init("span"); if __e != nil { return }
						{
//line fitness.gox:262
							__e = __c.Set("class", "fit-of"); if __e != nil { return }
							__e = __c.Submit(); if __e != nil { return }
//line fitness.gox:262
							__e = __c.Any(i18n.T(l, "fitness.of100")); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
//line fitness.gox:263
						if hasPrev {
							__e = __c.Init("span"); if __e != nil { return }
							{
//line fitness.gox:264
								__e = __c.Set("class", "fit-of"); if __e != nil { return }
								__e = __c.Submit(); if __e != nil { return }
//line fitness.gox:264
								__e = __c.Any(i18n.Tf(l, "fitness.was", previous.Index)); if __e != nil { return }
							}
							__e = __c.Close(); if __e != nil { return }
						}
					}
					__e = __c.Close(); if __e != nil { return }
					__e = __c.Init("p"); if __e != nil { return }
					{
//line fitness.gox:267
						__e = __c.Set("class", "fit-level"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
//line fitness.gox:267
						__e = __c.Any(i18n.T(l, "fitness.level_"+latest.Level)); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
					__e = __c.Init("div"); if __e != nil { return }
					{
//line fitness.gox:268
						__e = __c.Set("class", "settings-card"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
//line fitness.gox:269
						for _, r := range rows {
							__e = __c.Init("div"); if __e != nil { return }
							{
//line fitness.gox:270
								__e = __c.Set("class", "srow"); if __e != nil { return }
								__e = __c.Submit(); if __e != nil { return }
								__e = __c.Init("span"); if __e != nil { return }
								{
									__e = __c.Submit(); if __e != nil { return }
//line fitness.gox:270
									__e = __c.Any(i18n.T(l, r.Key)); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
								__e = __c.Init("span"); if __e != nil { return }
								{
//line fitness.gox:270
									__e = __c.Set("class", "fit-res"); if __e != nil { return }
									__e = __c.Submit(); if __e != nil { return }
//line fitness.gox:270
									__e = __c.Any(resultValue(r.Value, r.Seconds)); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
							}
							__e = __c.Close(); if __e != nil { return }
						}
					}
					__e = __c.Close(); if __e != nil { return }
				}
				__e = __c.Init("div"); if __e != nil { return }
				{
//line fitness.gox:274
					__e = __c.Set("class", "settings-card fit-plan"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
					__e = __c.Init("p"); if __e != nil { return }
					{
//line fitness.gox:275
						__e = __c.Set("class", "fit-plan-name"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
//line fitness.gox:275
						__e = __c.Any(prog.TitleIn(f.sess.Lang)); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
					__e = __c.Init("p"); if __e != nil { return }
					{
//line fitness.gox:276
						__e = __c.Set("class", "fit-plan-meta"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
//line fitness.gox:277
						if !built.FixedDifficulty {
//line fitness.gox:277
							__e = __c.Many(i18n.T(l, "settings.difficulty"), " ", levelLabel(l, plan.Level), " · "); if __e != nil { return }
						}
//line fitness.gox:278
						__e = __c.Many(i18n.Tn(l, "fmt.rounds", plan.Rounds), " · ", i18n.Tf(l, "fmt.min_range", first, peak)); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
//line fitness.gox:281
				__e = (doors.AClick{On: f.start}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
					ctx := __c.Context(); _ = ctx
					__e = __c.Init("button"); if __e != nil { return }
					{
//line fitness.gox:281
						__e = __c.Set("class", "btn primary fit-go"); if __e != nil { return }
//line fitness.gox:281
						__e = __c.Set("type", "button"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
//line fitness.gox:281
						__e = __c.Any(i18n.T(l, "home.start")); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
				return })); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
		}
		__e = __c.Close(); if __e != nil { return }
	return })
//line fitness.gox:284
}

// start switches to the recommended program and begins its next day.
func (f fitnessPage) start(ctx context.Context, r doors.RequestPointer) bool {
	if err := app.ApplyFitnessPlan(f.sess.UserID, f.sess.Lang); err != nil {
		return false
	}
	f.path.Update(ctx, path.Path{Page: path.Start})
	return false
}
