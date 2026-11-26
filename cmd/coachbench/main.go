// Command coachbench compares models for the AI coach on real program edits.
//
// Each scenario signs up a fresh user in a temporary database, picks a
// program, sends one message and checks the program afterwards in code.
//
//	AI_API_KEY=… go run ./cmd/coachbench deepseek/deepseek-v4.1-flash-1017@deepseek qwen/qwen3.7-flash@alibaba
package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"time"

	"morning30/internal/app"
	"morning30/internal/coach"
	"morning30/internal/content"
	"morning30/internal/db"
)

type scenario struct {
	name    string
	program string
	text    string
	check   func(uid int64, before state, res coach.Result) string // "" = passed
}

// state is what checks compare.
type state struct {
	ids      map[string]bool
	jumps    int
	minutes  int // longest day
	avgMin   int // average day
	rounds   int
	restMain int
	name     string
	warmup   []string // day 1 warm-up ids
}

func info(uid int64) map[string]app.ExerciseInfo {
	m := map[string]app.ExerciseInfo{}
	for _, e := range app.Library(uid, "en", "", "", false) {
		m[e.ID] = e
	}
	return m
}

func read(uid int64) state {
	lib := info(uid)
	r := app.UserProgram(uid)
	s := state{ids: map[string]bool{}, rounds: app.DB.GetRounds(uid), name: app.ProgramTitle(uid, "ru")}
	_, s.restMain, _ = app.DB.Rests(uid)
	for i := range r.Days {
		for _, list := range [][]content.RItem{r.WarmupFor(i), r.Days[i].Items, r.CooldownFor(i)} {
			for _, it := range list {
				s.ids[it.ID] = true
			}
		}
		d := app.Day(uid, "ru", i+1)
		s.minutes = max(s.minutes, d.EstMinutes)
		s.avgMin += d.EstMinutes
	}
	if len(r.Days) > 0 {
		s.avgMin /= len(r.Days)
	}
	for id := range s.ids {
		if lib[id].Jumps {
			s.jumps++
		}
	}
	for _, it := range r.WarmupFor(0) {
		s.warmup = append(s.warmup, it.ID)
	}
	return s
}

var scenarios = []scenario{
	{"no jumps", "slim30", "Убери все прыжки, у меня соседи снизу", func(uid int64, b state, _ coach.Result) string {
		if a := read(uid); a.jumps > 0 {
			return fmt.Sprintf("%d jumping exercises left", a.jumps)
		}
		return ""
	}},
	{"shorter", "demo", "Сделай тренировку короче, минут на 10", func(uid int64, b state, _ coach.Result) string {
		if a := read(uid); a.avgMin > 12 {
			return fmt.Sprintf("average day %d min (was %d)", a.avgMin, b.avgMin)
		}
		return ""
	}},
	{"neck warm-up", "demo", "Добавь в разминку что-нибудь для шеи", func(uid int64, b state, _ coach.Result) string {
		lib := info(uid)
		for _, id := range read(uid).warmup {
			if lib[id].Group == "neck" || lib[id].Custom || strings.Contains(strings.ToLower(lib[id].Name), "neck") || strings.Contains(strings.ToLower(lib[id].Name), "head") {
				return ""
			}
		}
		return "no neck exercise in the warm-up"
	}},
	{"easier push-ups", "demo", "Отжимания мне пока тяжело даются, замени на что-то полегче", func(uid int64, b state, _ coach.Result) string {
		lib := info(uid)
		before, after := 0, 0
		for id := range b.ids {
			if lib[id].Group == "push" {
				before = max(before, lib[id].Difficulty)
			}
		}
		for id := range read(uid).ids {
			if lib[id].Group == "push" {
				after = max(after, lib[id].Difficulty)
			}
		}
		if after >= before && before > 0 {
			return fmt.Sprintf("hardest push exercise still difficulty %d", after)
		}
		return ""
	}},
	{"rounds and rests", "demo", "Хочу три круга, а паузы между упражнениями покороче", func(uid int64, b state, _ coach.Result) string {
		a := read(uid)
		if a.rounds != 3 || a.restMain >= b.restMain {
			return fmt.Sprintf("rounds %d, main rest %d→%d", a.rounds, b.restMain, a.restMain)
		}
		return ""
	}},
	{"create exercise", "abs30", "Добавь в основную часть упражнение «лодочка с хлопком» — лёжа на животе поднять руки и ноги и хлопнуть руками за спиной. Если такого нет, создай.", func(uid int64, b state, _ coach.Result) string {
		lib := info(uid)
		for id := range read(uid).ids {
			if lib[id].Custom {
				return ""
			}
		}
		return "no custom exercise in the program"
	}},
	{"rename", "demo", "Переименуй программу в «Утро с Сергеем»", func(uid int64, b state, _ coach.Result) string {
		if a := read(uid); a.name != "Утро с Сергеем" {
			return fmt.Sprintf("name %q", a.name)
		}
		return ""
	}},
	{"off-topic", "demo", "Какая завтра погода в Москве?", func(uid int64, b state, res coach.Result) string {
		if len(res.Changes) > 0 {
			return "changed the program"
		}
		if res.Reply == "" {
			return "no reply"
		}
		return ""
	}},
	{"today only", "demo", "Сегодня хочу без планки", func(uid int64, b state, res coach.Result) string {
		r := app.UserProgram(uid)
		today := app.CurrentDay(uid) - 1
		for _, it := range r.Days[today].Items {
			if strings.HasPrefix(it.ID, "C0") && (it.ID == "C01" || it.ID == "C02") {
				return "plank still on today"
			}
		}
		if len(res.Changes) == 0 {
			return "nothing changed"
		}
		for _, c := range res.Changes {
			if len(c.Days) > 1 {
				return fmt.Sprintf("changed %d days, not only today", len(c.Days))
			}
		}
		return ""
	}},
}

type outcome struct {
	model, scenario, fail, reply string
	cost                         float64
	secs                         float64
	steps                        int
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: coachbench model@provider …")
		os.Exit(2)
	}
	dir, _ := os.MkdirTemp("", "coachbench")
	defer os.RemoveAll(dir)
	d, err := db.Open(filepath.Join(dir, "bench.db"))
	if err != nil {
		panic(err)
	}
	app.DB = d
	app.SeedBuiltins()
	base := coach.FromEnv()
	repeat := 2
	if os.Getenv("BENCH_REPEAT") == "1" {
		repeat = 1
	}
	var mu sync.Mutex
	var results []outcome
	var wg sync.WaitGroup
	for _, spec := range os.Args[1:] {
		cfg := base.WithModel(spec)
		model := cfg.Model
		wg.Add(1)
		go func() {
			defer wg.Done()
			for rep := 0; rep < repeat; rep++ {
				for i, sc := range scenarios {
					if only := os.Getenv("BENCH_ONLY"); only != "" && !strings.Contains(only, sc.name) {
						continue
					}
					u, err := d.CreateUser(fmt.Sprintf("%s-%d-%d@bench", strings.ReplaceAll(model, "/", "_"), rep, i), "x", "ru", "Сергей")
					if err != nil {
						panic(err)
					}
					if sc.program != "demo" {
						if err := app.SelectReady(u.ID, sc.program, "ru"); err != nil {
							panic(err)
						}
					}
					before := read(u.ID)
					start := time.Now()
					ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
					res, _, err := coach.Ask(ctx, cfg, u.ID, "ru", nil, sc.text)
					cancel()
					o := outcome{model: spec, scenario: sc.name, reply: res.Reply, cost: res.Usage.Cost, secs: time.Since(start).Seconds(), steps: res.Steps}
					if err != nil {
						o.fail = "error: " + err.Error()
					} else {
						o.fail = sc.check(u.ID, before, res)
					}
					mu.Lock()
					results = append(results, o)
					mark := "PASS"
					if o.fail != "" {
						mark = "FAIL " + o.fail
					}
					fmt.Printf("%-45s %-18s %5.1fs %d steps $%.5f %s | %s\n", spec, sc.name, o.secs, o.steps, o.cost, mark, strings.ReplaceAll(o.reply, "\n", " "))
					mu.Unlock()
				}
			}
		}()
	}
	wg.Wait()
	type sum struct {
		pass, total int
		cost, secs  float64
	}
	sums := map[string]*sum{}
	for _, r := range results {
		s := sums[r.model]
		if s == nil {
			s = &sum{}
			sums[r.model] = s
		}
		s.total++
		if r.fail == "" {
			s.pass++
		}
		s.cost += r.cost
		s.secs += r.secs
	}
	var names []string
	for n := range sums {
		names = append(names, n)
	}
	sort.Slice(names, func(i, j int) bool { return sums[names[i]].pass > sums[names[j]].pass })
	fmt.Println("\nmodel                                         passed   avg cost/msg   avg time")
	for _, n := range names {
		s := sums[n]
		fmt.Printf("%-45s %2d/%-2d   $%.5f       %.1fs\n", n, s.pass, s.total, s.cost/float64(s.total), s.secs/float64(s.total))
	}
}
