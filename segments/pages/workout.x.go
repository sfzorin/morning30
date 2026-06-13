// Managed by GoX v0.1.36

//line workout.gox:1
package pages

import (
	"context"
	"encoding/json"
	"strconv"

	"danicc/assets"
	"danicc/internal/app"
	"danicc/internal/auth"
	"danicc/internal/catalog"
	"danicc/internal/content"
	"danicc/internal/i18n"
	"danicc/path"

	"github.com/doors-dev/doors"
	"github.com/doors-dev/gox"
)

// completeReq is the body sent by the client when a day is finished, optionally
// carrying the post-workout feedback.
type completeReq struct {
	Day         int              `json:"day"`
	HasFeedback bool             `json:"hasFeedback"`
	Feedback    content.Feedback `json:"feedback"`
}

// playerRepl is the minimal data for a manual exercise replacement.
type playerRepl struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Unit    string `json:"unit"`
	Value   int    `json:"value"`
	PerSide bool   `json:"perSide"`
	SVG     string `json:"svg"`
	Hint    string `json:"hint"`
	Warning string `json:"warning"`
}

type playerItem struct {
	ID      string `json:"id"`
	Unit    string `json:"unit"`
	Value   int    `json:"value"`
	Slot    string `json:"slot"`
	PerSide bool   `json:"perSide"`
	Round   int    `json:"round"`
	Rest    int    `json:"rest"`
	Name    string `json:"name"`
	Hint    string `json:"hint"`    // short on-screen cue (localized)
	Warning string `json:"warning"` // inline safety warning (localized)
	SVG     string `json:"svg"`
	Video   string `json:"video"`
	// Detail card (Russian canonical content).
	Desc      string   `json:"desc"`
	HowTo     []string `json:"howTo"`
	Correct   []string `json:"correct"`
	Wrong     []string `json:"wrong"`
	Breathing string   `json:"breathing"`
	Safety    string   `json:"safety"`
	// Non-blocking voice cues.
	VStart  string   `json:"vStart"`
	VMid    []string `json:"vMid"`
	VLast   string   `json:"vLast"`
	VFinish string   `json:"vFinish"`
	// Manual replacement (nil if none).
	Repl *playerRepl `json:"repl,omitempty"`
}

type playerPayload struct {
	Day       int               `json:"day"`
	Lang      string            `json:"lang"`
	BCP47     string            `json:"bcp47"`
	Voice     bool              `json:"voice"`
	VoiceMode string            `json:"voiceMode"`
	Name      string            `json:"name"`
	Items     []playerItem      `json:"items"`
	Labels    map[string]string `json:"labels"`
	Cues      map[string]string `json:"cues"`
	Enc       []string          `json:"enc"`
}

// workoutPage renders the client-driven player for one day.
type workoutPage struct {
	sess auth.Session
	day  int
	path doors.Source[path.Path]
}

func (w workoutPage) clampDay() int {
	d := w.day
	if d < 1 {
		d = 1
	}
	if d > content.TotalDays {
		d = content.TotalDays
	}
	return d
}

// voiceFor returns the cues for an exercise. Exercises/languages with per-exercise
// content use those phrases; others fall back to the generic localized cues.
func (w workoutPage) voiceFor(l i18n.Lang, hasVoice bool, d content.Detail) (start string, mid []string, last, finish string) {
	if hasVoice {
		start, mid, last, finish = d.Voice.Start, d.Voice.Mid, d.Voice.Last, d.Voice.Finish
	}
	if start == "" {
		start = i18n.Cue(l, "go")
	}
	if finish == "" {
		finish = i18n.Cue(l, "well_done")
	}
	if last == "" {
		last = i18n.Cue(l, "five")
	}
	return
}

// replFor returns the manual-replacement item for an exercise (nil if none).
func (w workoutPage) replFor(l i18n.Lang, cat catalog.Catalog, id string, value int) *playerRepl {
	ex, ok := cat.Def(id)
	if !ok || ex.Replacement == "" {
		return nil
	}
	rid := ex.Replacement
	r, ok := cat.Def(rid)
	if !ok {
		return nil
	}
	return &playerRepl{
		ID:      rid,
		Name:    cat.Name(l, rid),
		Unit:    string(r.Unit),
		Value:   content.ConvertValue(ex.Unit, r.Unit, value),
		PerSide: r.PerSide,
		SVG:     cat.SVG(rid),
		Hint:    cat.Hint(l, rid),
		Warning: cat.Warning(l, rid),
	}
}

// payloadJSON serializes everything the client player needs.
func (w workoutPage) payloadJSON() string {
	l := i18n.Lang(w.sess.Lang)
	day := w.clampDay()
	adapt, _ := app.DB.GetAdaptState(w.sess.UserID)
	streak, _ := app.DB.Streak(w.sess.UserID, app.Today())
	wk := app.UserWorkout(w.sess.UserID, day, w.sess.Rest, adapt.ToAdapt().SuppressIncreaseOnGap(streak > 0))
	cat := app.UserCatalog(w.sess.UserID)

	items := make([]playerItem, 0, len(wk.Items))
	for _, it := range wk.Items {
		id := it.ExerciseID
		d, _ := cat.Detail(l, id)
		vs, vm, vl, vf := w.voiceFor(l, cat.HasVoiceLang(l, id), d)
		items = append(items, playerItem{
			ID:        id,
			Unit:      string(it.Unit),
			Value:     it.Value,
			Slot:      string(it.Slot),
			PerSide:   it.PerSide,
			Round:     it.Round,
			Rest:      it.Rest,
			Name:      cat.Name(l, id),
			Hint:      cat.Hint(l, id),
			Warning:   cat.Warning(l, id),
			SVG:       cat.SVG(id),
			Video:     content.Pool[id].Video,
			Desc:      d.Desc,
			HowTo:     d.HowTo,
			Correct:   d.Correct,
			Wrong:     d.Wrong,
			Breathing: d.Breathing,
			Safety:    d.Safety,
			VStart:    vs,
			VMid:      vm,
			VLast:     vl,
			VFinish:   vf,
			Repl:      w.replFor(l, cat, id, it.Value),
		})
	}

	labels := map[string]string{
		"warmup":     i18n.T(l, "workout.warmup"),
		"main":       i18n.T(l, "workout.main"),
		"cooldown":   i18n.T(l, "workout.cooldown"),
		"round":      i18n.T(l, "workout.round"),
		"rest":       i18n.T(l, "workout.rest"),
		"ready":      i18n.T(l, "workout.get_ready"),
		"next":       i18n.T(l, "workout.next"),
		"pause":      i18n.T(l, "workout.pause"),
		"resume":     i18n.T(l, "workout.resume"),
		"skip":       i18n.T(l, "workout.skip"),
		"prev":       i18n.T(l, "workout.prev"),
		"done":       i18n.T(l, "workout.done"),
		"sec":        i18n.T(l, "workout.sec"),
		"reps":       i18n.T(l, "workout.x_reps"),
		"breaths":    i18n.T(l, "workout.breaths"),
		"side":       i18n.T(l, "workout.side"),
		"done_title": i18n.T(l, "done.title"),
		"bravo":      i18n.T(l, "done.bravo"),
		"day_done":   i18n.T(l, "done.day_done"),
		"streak":     i18n.T(l, "done.streak"),
		"back_home":  i18n.T(l, "done.back_home"),
	}
	cues := map[string]string{
		"ready":       i18n.Cue(l, "ready"),
		"go":          i18n.Cue(l, "go"),
		"rest":        i18n.Cue(l, "rest"),
		"next":        i18n.Cue(l, "next"),
		"halfway":     i18n.Cue(l, "halfway"),
		"last":        i18n.Cue(l, "last"),
		"well_done":   i18n.Cue(l, "well_done"),
		"workout_end": i18n.Cue(l, "workout_end"),
		"switch_side": i18n.Cue(l, "switch_side"),
	}
	enc := make([]string, 0, 6)
	for n := 0; n < 6; n++ {
		enc = append(enc, i18n.T(l, "enc."+strconv.Itoa(n)))
	}

	mode := w.sess.VoiceMode
	if mode == "" {
		mode = "normal"
	}
	p := playerPayload{
		Day:       day,
		Lang:      w.sess.Lang,
		BCP47:     i18n.BCP47(l),
		Voice:     w.sess.Voice,
		VoiceMode: mode,
		Name:      w.sess.Name,
		Items:     items,
		Labels: labels,
		Cues:   cues,
		Enc:    enc,
	}
	b, _ := json.Marshal(p)
	return string(b)
}

// complete is the JS hook fired when a day finishes. It records today's
// activity (streak/calendar), advances the looping 30-day cycle when the
// finished day is the current one, and folds the optional post-workout feedback
// into the adaptation state that shapes the next workout. Returns the streak.
func (w workoutPage) complete(ctx context.Context, r doors.RequestHook[completeReq]) (any, bool) {
	uid := w.sess.UserID
	day := w.clampDay()
	today := app.Today()
	data := r.Data()

	_ = app.DB.RecordActivity(uid, today)

	// One session passed: decay block counters and clear the one-shot
	// level/force-R that applied to this workout.
	st, _ := app.DB.GetAdaptState(uid)
	st = st.Decay()

	if pos, err := app.DB.Position(uid); err == nil {
		if day == pos%content.TotalDays+1 {
			_ = app.DB.AdvancePosition(uid)
		}
	}

	if data.HasFeedback {
		st = content.NextAdapt(st, data.Feedback)
		_ = app.DB.AddFeedback(uid, today, day, data.Feedback)
	}
	_ = app.DB.SetAdaptState(uid, st)

	streak, _ := app.DB.Streak(uid, today)
	return streak, false
}

// painGroup renders a 3-chip pain question (None / A bit / Pain → 0 / 3 / 6).
//line workout.gox:275
func (w workoutPage) painGroup(l i18n.Lang, key, labelKey string) gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
		__e = __c.Init("div"); if __e != nil { return }
		{
//line workout.gox:276
			__e = __c.Set("class", "fb-q fb-group"); if __e != nil { return }
//line workout.gox:276
			__e = __c.Set("data-key", key); if __e != nil { return }
			__e = __c.Submit(); if __e != nil { return }
			__e = __c.Init("label"); if __e != nil { return }
			{
//line workout.gox:277
				__e = __c.Set("class", "fb-label"); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
//line workout.gox:277
				__e = __c.Any(i18n.T(l, labelKey)); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
			__e = __c.Init("div"); if __e != nil { return }
			{
//line workout.gox:278
				__e = __c.Set("class", "fb-chips"); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
				__e = __c.Init("button"); if __e != nil { return }
				{
//line workout.gox:279
					__e = __c.Set("type", "button"); if __e != nil { return }
//line workout.gox:279
					__e = __c.Set("class", "fb-chip active"); if __e != nil { return }
//line workout.gox:279
					__e = __c.Set("data-val", "0"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
//line workout.gox:279
					__e = __c.Any(i18n.T(l, "fb.none")); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
				__e = __c.Init("button"); if __e != nil { return }
				{
//line workout.gox:280
					__e = __c.Set("type", "button"); if __e != nil { return }
//line workout.gox:280
					__e = __c.Set("class", "fb-chip"); if __e != nil { return }
//line workout.gox:280
					__e = __c.Set("data-val", "3"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
//line workout.gox:280
					__e = __c.Any(i18n.T(l, "fb.some")); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
				__e = __c.Init("button"); if __e != nil { return }
				{
//line workout.gox:281
					__e = __c.Set("type", "button"); if __e != nil { return }
//line workout.gox:281
					__e = __c.Set("class", "fb-chip"); if __e != nil { return }
//line workout.gox:281
					__e = __c.Set("data-val", "6"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
//line workout.gox:281
					__e = __c.Any(i18n.T(l, "fb.pain")); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
		}
		__e = __c.Close(); if __e != nil { return }
	return })
//line workout.gox:284
}

//line workout.gox:286
func (w workoutPage) Main() gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
//line workout.gox:288
		l := i18n.Lang(w.sess.Lang)
		payloadStr := w.payloadJSON()

		__e = __c.Init("title"); if __e != nil { return }
		{
			__e = __c.Submit(); if __e != nil { return }
//line workout.gox:291
			__e = __c.Any(i18n.T(l, "app.name")); if __e != nil { return }
		}
		__e = __c.Close(); if __e != nil { return }
		__e = __c.Init("div"); if __e != nil { return }
		{
//line workout.gox:292
			__e = __c.Set("class", "player"); if __e != nil { return }
			__e = __c.Submit(); if __e != nil { return }
			__e = __c.Init("div"); if __e != nil { return }
			{
//line workout.gox:293
				__e = __c.Set("class", "player-top"); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
//line workout.gox:294
				__e = (doors.ALink{Model: path.Path{Page: path.Home}}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
					ctx := __c.Context(); _ = ctx
					__e = __c.Init("a"); if __e != nil { return }
					{
//line workout.gox:294
						__e = __c.Set("class", "quit"); if __e != nil { return }
//line workout.gox:294
						__e = __c.Set("aria-label", i18n.T(l, "workout.quit")); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
						__e = __c.Text("✕"); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
				return })); if __e != nil { return }
				__e = __c.Init("div"); if __e != nil { return }
				{
//line workout.gox:295
					__e = __c.Set("class", "pbar"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
					__e = __c.Init("div"); if __e != nil { return }
					{
//line workout.gox:295
						__e = __c.Set("class", "pbar-fill"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
				__e = __c.Init("div"); if __e != nil { return }
				{
//line workout.gox:296
					__e = __c.Set("class", "phase"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
				__e = __c.Init("button"); if __e != nil { return }
				{
//line workout.gox:297
					__e = __c.Set("class", "info-btn"); if __e != nil { return }
//line workout.gox:297
					__e = __c.Set("type", "button"); if __e != nil { return }
//line workout.gox:297
					__e = __c.Set("aria-label", i18n.T(l, "info.technique")); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
					__e = __c.Text("ⓘ"); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
			__e = __c.Init("div"); if __e != nil { return }
			{
//line workout.gox:300
				__e = __c.Set("class", "stage"); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
				__e = __c.Init("div"); if __e != nil { return }
				{
//line workout.gox:301
					__e = __c.Set("class", "ex-index"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
				__e = __c.InitVoid("img"); if __e != nil { return }
				{
//line workout.gox:302
					__e = __c.Set("class", "ex-svg"); if __e != nil { return }
//line workout.gox:302
					__e = __c.Set("alt", ""); if __e != nil { return }
				}
				__e = __c.Submit(); if __e != nil { return }
				__e = __c.Init("div"); if __e != nil { return }
				{
//line workout.gox:303
					__e = __c.Set("class", "ex-name"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
				__e = __c.Init("div"); if __e != nil { return }
				{
//line workout.gox:304
					__e = __c.Set("class", "ex-side"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
				__e = __c.Init("div"); if __e != nil { return }
				{
//line workout.gox:305
					__e = __c.Set("class", "ex-value"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
				__e = __c.Init("div"); if __e != nil { return }
				{
//line workout.gox:306
					__e = __c.Set("class", "ex-hint"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
				__e = __c.Init("div"); if __e != nil { return }
				{
//line workout.gox:307
					__e = __c.Set("class", "ex-warn"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
			__e = __c.Init("div"); if __e != nil { return }
			{
//line workout.gox:310
				__e = __c.Set("class", "controls"); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
				__e = __c.Init("button"); if __e != nil { return }
				{
//line workout.gox:311
					__e = __c.Set("class", "ctl-prev btn"); if __e != nil { return }
//line workout.gox:311
					__e = __c.Set("type", "button"); if __e != nil { return }
//line workout.gox:311
					__e = __c.Set("aria-label", i18n.T(l, "workout.prev")); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
					__e = __c.Text("◀"); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
				__e = __c.Init("button"); if __e != nil { return }
				{
//line workout.gox:312
					__e = __c.Set("class", "ctl-pause btn"); if __e != nil { return }
//line workout.gox:312
					__e = __c.Set("type", "button"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
				__e = __c.Init("button"); if __e != nil { return }
				{
//line workout.gox:313
					__e = __c.Set("class", "ctl-skip btn"); if __e != nil { return }
//line workout.gox:313
					__e = __c.Set("type", "button"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
				__e = __c.Init("button"); if __e != nil { return }
				{
//line workout.gox:314
					__e = __c.Set("class", "ctl-done btn primary"); if __e != nil { return }
//line workout.gox:314
					__e = __c.Set("type", "button"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
			__e = __c.Init("div"); if __e != nil { return }
			{
//line workout.gox:317
				__e = __c.Set("class", "rest-overlay hidden"); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
				__e = __c.Init("div"); if __e != nil { return }
				{
//line workout.gox:318
					__e = __c.Set("class", "rest-title"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
				__e = __c.Init("div"); if __e != nil { return }
				{
//line workout.gox:319
					__e = __c.Set("class", "rest-count"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
				__e = __c.Init("div"); if __e != nil { return }
				{
//line workout.gox:320
					__e = __c.Set("class", "rest-next"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
				__e = __c.Init("div"); if __e != nil { return }
				{
//line workout.gox:321
					__e = __c.Set("class", "rest-actions"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
					__e = __c.Init("button"); if __e != nil { return }
					{
//line workout.gox:322
						__e = __c.Set("class", "rest-back"); if __e != nil { return }
//line workout.gox:322
						__e = __c.Set("type", "button"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
					__e = __c.Init("button"); if __e != nil { return }
					{
//line workout.gox:323
						__e = __c.Set("class", "rest-skip"); if __e != nil { return }
//line workout.gox:323
						__e = __c.Set("type", "button"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
			__e = __c.Init("div"); if __e != nil { return }
			{
//line workout.gox:328
				__e = __c.Set("class", "info-overlay hidden"); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
				__e = __c.Init("div"); if __e != nil { return }
				{
//line workout.gox:329
					__e = __c.Set("class", "info-card"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
					__e = __c.InitVoid("img"); if __e != nil { return }
					{
//line workout.gox:330
						__e = __c.Set("class", "info-media"); if __e != nil { return }
//line workout.gox:330
						__e = __c.Set("alt", ""); if __e != nil { return }
					}
					__e = __c.Submit(); if __e != nil { return }
					__e = __c.Init("h2"); if __e != nil { return }
					{
//line workout.gox:331
						__e = __c.Set("class", "info-name"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
					__e = __c.Init("p"); if __e != nil { return }
					{
//line workout.gox:332
						__e = __c.Set("class", "info-desc"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
					__e = __c.Init("div"); if __e != nil { return }
					{
//line workout.gox:333
						__e = __c.Set("class", "info-sec info-how-sec"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
						__e = __c.Init("h3"); if __e != nil { return }
						{
							__e = __c.Submit(); if __e != nil { return }
//line workout.gox:334
							__e = __c.Any(i18n.T(l, "info.how")); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
						__e = __c.Init("ol"); if __e != nil { return }
						{
//line workout.gox:335
							__e = __c.Set("class", "info-how"); if __e != nil { return }
							__e = __c.Submit(); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
					__e = __c.Init("div"); if __e != nil { return }
					{
//line workout.gox:337
						__e = __c.Set("class", "info-sec info-correct-sec"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
						__e = __c.Init("h3"); if __e != nil { return }
						{
							__e = __c.Submit(); if __e != nil { return }
//line workout.gox:338
							__e = __c.Any(i18n.T(l, "info.correct")); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
						__e = __c.Init("ul"); if __e != nil { return }
						{
//line workout.gox:339
							__e = __c.Set("class", "info-correct"); if __e != nil { return }
							__e = __c.Submit(); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
					__e = __c.Init("div"); if __e != nil { return }
					{
//line workout.gox:341
						__e = __c.Set("class", "info-sec info-wrong-sec"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
						__e = __c.Init("h3"); if __e != nil { return }
						{
							__e = __c.Submit(); if __e != nil { return }
//line workout.gox:342
							__e = __c.Any(i18n.T(l, "info.wrong")); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
						__e = __c.Init("ul"); if __e != nil { return }
						{
//line workout.gox:343
							__e = __c.Set("class", "info-wrong"); if __e != nil { return }
							__e = __c.Submit(); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
					__e = __c.Init("div"); if __e != nil { return }
					{
//line workout.gox:345
						__e = __c.Set("class", "info-sec info-breathing-sec"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
						__e = __c.Init("h3"); if __e != nil { return }
						{
							__e = __c.Submit(); if __e != nil { return }
//line workout.gox:346
							__e = __c.Any(i18n.T(l, "info.breathing")); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
						__e = __c.Init("p"); if __e != nil { return }
						{
//line workout.gox:347
							__e = __c.Set("class", "info-breathing"); if __e != nil { return }
							__e = __c.Submit(); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
					__e = __c.Init("div"); if __e != nil { return }
					{
//line workout.gox:349
						__e = __c.Set("class", "info-sec info-warn-sec"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
						__e = __c.Init("h3"); if __e != nil { return }
						{
							__e = __c.Submit(); if __e != nil { return }
							__e = __c.Text("⚠️ "); if __e != nil { return }
//line workout.gox:350
							__e = __c.Any(i18n.T(l, "info.warning")); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
						__e = __c.Init("p"); if __e != nil { return }
						{
//line workout.gox:351
							__e = __c.Set("class", "info-warn"); if __e != nil { return }
							__e = __c.Submit(); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
					__e = __c.Init("button"); if __e != nil { return }
					{
//line workout.gox:353
						__e = __c.Set("class", "info-replace btn"); if __e != nil { return }
//line workout.gox:353
						__e = __c.Set("type", "button"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
//line workout.gox:353
						__e = __c.Any(i18n.T(l, "info.replace")); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
					__e = __c.Init("button"); if __e != nil { return }
					{
//line workout.gox:354
						__e = __c.Set("class", "info-close btn primary"); if __e != nil { return }
//line workout.gox:354
						__e = __c.Set("type", "button"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
//line workout.gox:354
						__e = __c.Any(i18n.T(l, "info.close")); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
			__e = __c.Init("div"); if __e != nil { return }
			{
//line workout.gox:359
				__e = __c.Set("class", "feedback-overlay hidden"); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
				__e = __c.Init("h2"); if __e != nil { return }
				{
					__e = __c.Submit(); if __e != nil { return }
//line workout.gox:360
					__e = __c.Any(i18n.T(l, "fb.title")); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
				__e = __c.Init("div"); if __e != nil { return }
				{
//line workout.gox:361
					__e = __c.Set("class", "fb-q"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
					__e = __c.Init("label"); if __e != nil { return }
					{
//line workout.gox:362
						__e = __c.Set("class", "row-between"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
						__e = __c.Init("span"); if __e != nil { return }
						{
							__e = __c.Submit(); if __e != nil { return }
//line workout.gox:363
							__e = __c.Any(i18n.T(l, "fb.difficulty")); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
						__e = __c.Init("output"); if __e != nil { return }
						{
//line workout.gox:364
							__e = __c.Set("class", "fb-diff-out"); if __e != nil { return }
							__e = __c.Submit(); if __e != nil { return }
							__e = __c.Text("6"); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
					__e = __c.InitVoid("input"); if __e != nil { return }
					{
//line workout.gox:366
						__e = __c.Set("class", "fb-diff"); if __e != nil { return }
//line workout.gox:366
						__e = __c.Set("type", "range"); if __e != nil { return }
//line workout.gox:366
						__e = __c.Set("min", "1"); if __e != nil { return }
//line workout.gox:366
						__e = __c.Set("max", "10"); if __e != nil { return }
//line workout.gox:366
						__e = __c.Set("value", "6"); if __e != nil { return }
					}
					__e = __c.Submit(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
//line workout.gox:368
				__e = __c.Any(w.painGroup(l, "knee", "fb.knee")); if __e != nil { return }
//line workout.gox:369
				__e = __c.Any(w.painGroup(l, "back", "fb.back")); if __e != nil { return }
//line workout.gox:370
				__e = __c.Any(w.painGroup(l, "shoulder", "fb.shoulder")); if __e != nil { return }
				__e = __c.Init("div"); if __e != nil { return }
				{
//line workout.gox:371
					__e = __c.Set("class", "fb-q fb-group"); if __e != nil { return }
//line workout.gox:371
					__e = __c.Set("data-key", "energy"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
					__e = __c.Init("label"); if __e != nil { return }
					{
//line workout.gox:372
						__e = __c.Set("class", "fb-label"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
//line workout.gox:372
						__e = __c.Any(i18n.T(l, "fb.energy")); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
					__e = __c.Init("div"); if __e != nil { return }
					{
//line workout.gox:373
						__e = __c.Set("class", "fb-chips"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
						__e = __c.Init("button"); if __e != nil { return }
						{
//line workout.gox:374
							__e = __c.Set("type", "button"); if __e != nil { return }
//line workout.gox:374
							__e = __c.Set("class", "fb-chip"); if __e != nil { return }
//line workout.gox:374
							__e = __c.Set("data-val", "worse"); if __e != nil { return }
							__e = __c.Submit(); if __e != nil { return }
//line workout.gox:374
							__e = __c.Any(i18n.T(l, "fb.worse")); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
						__e = __c.Init("button"); if __e != nil { return }
						{
//line workout.gox:375
							__e = __c.Set("type", "button"); if __e != nil { return }
//line workout.gox:375
							__e = __c.Set("class", "fb-chip active"); if __e != nil { return }
//line workout.gox:375
							__e = __c.Set("data-val", "normal"); if __e != nil { return }
							__e = __c.Submit(); if __e != nil { return }
//line workout.gox:375
							__e = __c.Any(i18n.T(l, "fb.normal")); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
						__e = __c.Init("button"); if __e != nil { return }
						{
//line workout.gox:376
							__e = __c.Set("type", "button"); if __e != nil { return }
//line workout.gox:376
							__e = __c.Set("class", "fb-chip"); if __e != nil { return }
//line workout.gox:376
							__e = __c.Set("data-val", "better"); if __e != nil { return }
							__e = __c.Submit(); if __e != nil { return }
//line workout.gox:376
							__e = __c.Any(i18n.T(l, "fb.better")); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
				__e = __c.Init("div"); if __e != nil { return }
				{
//line workout.gox:379
					__e = __c.Set("class", "fb-actions"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
					__e = __c.Init("button"); if __e != nil { return }
					{
//line workout.gox:380
						__e = __c.Set("type", "button"); if __e != nil { return }
//line workout.gox:380
						__e = __c.Set("class", "fb-save btn primary"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
//line workout.gox:380
						__e = __c.Any(i18n.T(l, "fb.save")); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
					__e = __c.Init("button"); if __e != nil { return }
					{
//line workout.gox:381
						__e = __c.Set("type", "button"); if __e != nil { return }
//line workout.gox:381
						__e = __c.Set("class", "fb-skip"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
//line workout.gox:381
						__e = __c.Any(i18n.T(l, "fb.skip")); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
			__e = __c.Init("div"); if __e != nil { return }
			{
//line workout.gox:385
				__e = __c.Set("class", "done-overlay hidden"); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
				__e = __c.Init("div"); if __e != nil { return }
				{
//line workout.gox:386
					__e = __c.Set("class", "applause"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
					__e = __c.Text("👏"); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
				__e = __c.Init("div"); if __e != nil { return }
				{
//line workout.gox:387
					__e = __c.Set("class", "done-title"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
				__e = __c.Init("div"); if __e != nil { return }
				{
//line workout.gox:388
					__e = __c.Set("class", "done-enc"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
				__e = __c.Init("div"); if __e != nil { return }
				{
//line workout.gox:389
					__e = __c.Set("class", "done-streak"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
//line workout.gox:390
				__e = (doors.ALink{Model: path.Path{Page: path.Home}}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
					ctx := __c.Context(); _ = ctx
					__e = __c.Init("a"); if __e != nil { return }
					{
//line workout.gox:390
						__e = __c.Set("class", "btn primary back-home"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
//line workout.gox:390
						__e = __c.Any(i18n.T(l, "done.back_home")); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
				return })); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
			__e = __c.Init("script"); if __e != nil { return }
			{
//line workout.gox:394
				__e = __c.Set("src", assets.Player); if __e != nil { return }
				__e = __c.Set("inline", true); if __e != nil { return }
//line workout.gox:396
				__e = __c.Set("data:payload", payloadStr); if __e != nil { return }
//line workout.gox:397
				__e = __c.Modify(doors.AHook[completeReq]{Name: "complete", On: w.complete}); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
				__e = __c.Raw(""); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
		}
		__e = __c.Close(); if __e != nil { return }
	return })
//line workout.gox:400
}
