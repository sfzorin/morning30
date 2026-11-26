package coach

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"morning30/internal/app"
	"morning30/internal/content"
	"morning30/internal/db"
	"morning30/internal/i18n"
	"morning30/internal/langs"
)

// DailyLimit is how many messages one account may send the coach per day.
const DailyLimit = 40

// maxSteps bounds the tool-calling loop of one reply.
const maxSteps = 12

// ErrLimit is returned when the daily limit is used up.
var ErrLimit = errors.New("daily coach limit reached")

// ErrDisabled is returned when no model key is configured.
var ErrDisabled = errors.New("the AI coach is not configured")

// session is one reply in progress.
type session struct {
	userID  int64
	lang    string
	edited  bool
	created []string // ids of exercises created during the reply
}

// Result is the coach's reply to one message.
type Result struct {
	Reply   string
	Changes []Change
	Usage   Usage
	Steps   int
}

// snapshot is everything an undo brings back.
type snapshot struct {
	Program      string `json:"program"`
	Custom       bool   `json:"custom"`
	Level        int    `json:"level"`
	Rounds       int    `json:"rounds"`
	RestWarmup   int    `json:"restWarmup"`
	RestMain     int    `json:"restMain"`
	RestCooldown int    `json:"restCooldown"`
	Name         string `json:"name"`
}

func takeSnapshot(userID int64) snapshot {
	w, m, c := app.DB.Rests(userID)
	return snapshot{
		Program: app.UserProgram(userID).Marshal(), Custom: app.HasCustomProgram(userID),
		Level: app.DB.GetLevel(userID), Rounds: app.DB.GetRounds(userID), RestWarmup: w, RestMain: m, RestCooldown: c,
		Name: app.ProgramTitle(userID, app.LangOf(userID)),
	}
}

// Ask sends one message to the coach, lets it edit the current program, and
// returns its reply with the list of changes. history is the earlier
// conversation of this program (oldest first).
func Ask(ctx context.Context, cfg Config, userID int64, lang string, history []db.CoachMessage, text string) (Result, snapshot, error) {
	before := takeSnapshot(userID)
	if !cfg.Enabled() {
		return Result{}, before, ErrDisabled
	}
	s := &session{userID: userID, lang: lang}
	msgs := []Message{{Role: "system", Content: SystemPrompt(userID, lang)}}
	for _, h := range history {
		role := "user"
		if h.Role == "coach" {
			role = "assistant"
		}
		msgs = append(msgs, Message{Role: role, Content: h.Text})
	}
	msgs = append(msgs, Message{Role: "user", Content: text})
	var res Result
	tools := toolDefs()
	for step := 0; step < maxSteps; step++ {
		msg, usage, err := cfg.chat(ctx, msgs, tools)
		if err != nil && cfg.Fallback != nil && ctx.Err() == nil {
			// The main model is down or gone: the fallback carries on.
			cfg = *cfg.Fallback
			msg, usage, err = cfg.chat(ctx, msgs, tools)
		}
		res.Usage.PromptTokens += usage.PromptTokens
		res.Usage.CompletionTokens += usage.CompletionTokens
		res.Usage.Cost += usage.Cost
		res.Steps = step + 1
		if err != nil {
			return res, before, err
		}
		if len(msg.ToolCalls) == 0 {
			res.Reply = cleanReply(msg.Content)
			break
		}
		msg.Role = "assistant"
		msgs = append(msgs, msg)
		for _, call := range msg.ToolCalls {
			out := s.run(call.Function.Name, call.Function.Arguments)
			msgs = append(msgs, Message{Role: "tool", ToolCallID: call.ID, Content: out})
		}
	}
	if s.edited {
		res.Changes = Diff(userID, lang, before, s.created)
	}
	return res, before, nil
}

// Undo brings back the program and its workout settings from before a reply.
func Undo(userID int64, raw string) error {
	var snap snapshot
	if err := json.Unmarshal([]byte(raw), &snap); err != nil || snap.Program == "" {
		return errors.New("nothing to undo")
	}
	r, err := content.ParseResolved(snap.Program)
	if err != nil {
		return err
	}
	if snap.Custom || app.HasCustomProgram(userID) {
		if !snap.Custom {
			r.Name = snap.Name
		}
		if err := app.SaveCurrent(userID, r); err != nil {
			return err
		}
	}
	_, err = app.UpdateSettings(userID, app.SettingsPatch{
		Level: &snap.Level, Rounds: &snap.Rounds, RestWarmup: &snap.RestWarmup, RestMain: &snap.RestMain, RestCooldown: &snap.RestCooldown,
	})
	return err
}

// SnapshotJSON serializes a snapshot for storage with the reply.
func SnapshotJSON(s snapshot) string {
	b, _ := json.Marshal(s)
	return string(b)
}

// SystemPrompt tells the model where it is, who it talks to and what the
// current program looks like, so the person never has to explain it.
func SystemPrompt(userID int64, lang string) string {
	l := i18n.Parse(lang)
	langName := string(l)
	if m, ok := langs.Get(string(l)); ok {
		langName = m.Name + " (" + m.Code + ")"
	}
	var b strings.Builder
	b.WriteString(`You are the AI coach inside morning30.com — a free app (website and iPhone/Android apps) with voice-guided 30-day home workout programs. The person has opened "AI coach" from the page of their CURRENT program and wants to change that program. You are a friendly, competent personal trainer.

What you can do (with tools): change the current program — add, replace and remove exercises on chosen days, change doses, change this program's rounds, difficulty and rests, rename it, and create new exercises in the person's own library. You cannot switch or delete programs, change other settings, see other people's data, or do anything outside this program. If asked for something else, say briefly what you can do here.

How a program works: 30 days; each day = warm-up, then the main part repeated N rounds, then cool-down. The app speaks every exercise, counts reps or seconds and pauses between exercises. After day 30 the cycle starts again. Doses: reps (count), seconds (hold/timed) or breaths. "Per side" exercises are done on each side. The difficulty level scales every main-part dose by 10% per step.

Rules a good program follows (keep them when you edit):
- warm-up 4–7 exercises, cool-down about 4, main part 4–7 exercises per day;
- a day takes 5–25 minutes; keep the length the person asked for;
- the load grows gradually through the 30 days; recovery days are lighter;
- no jumps when the person mentions neighbours, knees, weight or quiet; no equipment unless they have it;
- warm-up exercises go to warm-up, stretches to cool-down;
- never change things the person did not ask for.

How to work:
- Do it, don't ask, unless the request is truly unclear. Use the program below and search_exercises to pick exercise ids; never invent ids.
- "Today" means day ` + fmt.Sprint(app.CurrentDay(userID)) + `. When no days are named, change all days where the exercise appears.
- Create an exercise only when the library has nothing close; then add it.
- Health: you are not a doctor. With pain, injury, pregnancy or illness, make the program gentler and suggest seeing a doctor; never diagnose.
- Answer in ` + langName + `, in 1–3 short sentences, plain text without markdown. The app speaks to the person like this: «` + i18n.T(l, "coach.lead") + `» — use the same form of address (formal or informal) and the same tone. Do not list every change: the app shows a list of the changes under your answer. Use the exercises' names, never ids.
`)
	b.WriteString("\nThe person:\n")
	p, hasProfile := app.FitnessProfile(userID)
	figure := map[string]string{"m": "man", "f": "woman", "b": "boy", "g": "girl"}[p.Sex]
	fmt.Fprintf(&b, "- figure in pictures: %s\n", figure)
	if hasProfile {
		if p.Child {
			b.WriteString("- this program is for a child; keep it playful and safe\n")
		}
		if p.Age > 0 {
			fmt.Fprintf(&b, "- age %d\n", p.Age)
		}
		if bmi := p.BMI(); bmi > 0 {
			fmt.Fprintf(&b, "- height %d cm, weight %d kg\n", p.Height, p.Weight)
		}
		if len(p.Concerns) > 0 {
			fmt.Fprintf(&b, "- told us about: %s\n", strings.Join(p.Concerns, ", "))
		}
	}
	if latest, _, ok, _ := app.LatestFitness(userID); ok {
		fmt.Fprintf(&b, "- fitness index %d of 100 (%s)\n", latest.Index, latest.Level)
	}
	if n, _ := app.DB.Position(userID); n > 0 {
		fmt.Fprintf(&b, "- workouts done in this program: %d\n", n)
	}
	b.WriteString("\n")
	b.WriteString(Overview(userID, lang))
	return b.String()
}

// Overview writes the current program day by day, with ids and doses.
func Overview(userID int64, lang string) string {
	l := i18n.Parse(lang)
	r := app.UserProgram(userID)
	cat := app.UserCatalog(userID)
	w, m, c := app.DB.Rests(userID)
	var b strings.Builder
	fmt.Fprintf(&b, "Current program: %q. Next workout: day %d of 30. Main part: %d rounds. Difficulty level %d (−3..+3). Rests: warm-up %d s, main %d s, cool-down %d s.\n",
		app.ProgramTitle(userID, lang), app.CurrentDay(userID), app.DB.GetRounds(userID), app.DB.GetLevel(userID), w, m, c)
	if r.FixedDifficulty {
		b.WriteString("This program has prescribed doses: the difficulty level does not apply.\n")
	}
	item := func(it content.RItem) string {
		dose := fmt.Sprintf("%d reps", it.Value)
		switch it.Unit {
		case content.Seconds:
			dose = fmt.Sprintf("%d s", it.Value)
		case content.Breaths:
			dose = fmt.Sprintf("%d breaths", it.Value)
		}
		if it.PerSide {
			dose += " per side"
		}
		return fmt.Sprintf("%s %s — %s", it.ID, cat.Name(l, it.ID), dose)
	}
	list := func(items []content.RItem) string {
		parts := make([]string, len(items))
		for i, it := range items {
			parts[i] = item(it)
		}
		return strings.Join(parts, "; ")
	}
	var days []string
	for i := range r.Days {
		days = append(days, "warm-up: "+list(r.WarmupFor(i))+" | main: "+list(r.Days[i].Items)+" | cool-down: "+list(r.CooldownFor(i)))
	}
	// Consecutive identical days are written once.
	for i := 0; i < len(days); {
		j := i
		for j+1 < len(days) && days[j+1] == days[i] {
			j++
		}
		if i == j {
			fmt.Fprintf(&b, "Day %d: %s\n", i+1, days[i])
		} else {
			fmt.Fprintf(&b, "Days %d–%d: %s\n", i+1, j+1, days[i])
		}
		i = j + 1
	}
	return b.String()
}

// startOfDay is the beginning of today in UTC (the daily limit's window).
func startOfDay() time.Time {
	now := time.Now().UTC()
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
}

// LimitReached reports whether the user has used up today's messages.
func LimitReached(userID int64) bool {
	return app.DB.CoachRequestsSince(userID, startOfDay()) >= DailyLimit
}

// cleanReply drops reasoning some models leak into the answer
// ("…</think> answer") and markdown emphasis the screen does not render.
func cleanReply(s string) string {
	if i := strings.LastIndex(s, "</think>"); i >= 0 {
		s = s[i+len("</think>"):]
	}
	s = strings.NewReplacer("**", "", "__", "").Replace(s)
	return strings.TrimSpace(s)
}
