// Package landing serves the public, crawlable pages of the "… in 30 days"
// courses: one index and one page per course in every supported language.
// They are plain HTML outside the app: no session is read or created and no
// account is touched. Their call to action only links into the app.
package landing

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"morning30/internal/catalog"
	"morning30/internal/content"
	"morning30/internal/i18n"
	"morning30/internal/langs"
)

// Origin is the canonical public address used in links for crawlers.
const Origin = "https://morning30.com"

// pageRests are the rest settings of a new account; durations shown publicly
// match what a newcomer sees in the app.
var pageRests = content.Rests{Warmup: 10, Main: 20, Cooldown: 10}

// Register adds the public course routes. Paths are listed per language so
// they never overlap the app's own routes.
func Register(mux *http.ServeMux) {
	mux.HandleFunc("GET /privacy", servePrivacy)
	mux.HandleFunc("GET /.well-known/assetlinks.json", serveAssetLinks)
	mux.HandleFunc("GET /courses", func(w http.ResponseWriter, r *http.Request) {
		l := i18n.FromAcceptLanguage(r.Header.Get("Accept-Language"))
		http.Redirect(w, r, IndexPath(string(l)), http.StatusFound)
	})
	for _, l := range i18n.Order {
		mux.HandleFunc("GET "+IndexPath(string(l)), func(w http.ResponseWriter, r *http.Request) {
			render(w, indexPage(l))
		})
		mux.HandleFunc("GET "+IndexPath(string(l))+"/{slug}", func(w http.ResponseWriter, r *http.Request) {
			c, ok := content.CourseBySlug(r.PathValue("slug"))
			if !ok {
				http.NotFound(w, r)
				return
			}
			render(w, coursePage(l, c))
		})
	}
}

// IndexPath is the course list for a language, e.g. "/ru/courses".
func IndexPath(lang string) string { return "/" + lang + "/courses" }

// CoursePath is one course's page for a language.
func CoursePath(lang string, c content.Course) string { return IndexPath(lang) + "/" + c.Slug }

// SitemapURLs lists every public course URL, absolute.
func SitemapURLs() []string {
	var out []string
	for _, l := range i18n.Order {
		out = append(out, Origin+IndexPath(string(l)))
		for _, c := range content.Series30() {
			out = append(out, Origin+CoursePath(string(l), c))
		}
	}
	return out
}

type link struct {
	Lang, Label, URL string
	Current          bool
}

type card struct {
	Title, About, URL, Minutes string
}

type dayRow struct {
	Label, Minutes string // "Day 3", "≈ 14 min"
	Light          bool
	Items          []string
}

type page struct {
	Lang, HTMLLang, Dir string
	Title, Desc        string
	Canonical, Default string
	Alternates         []link
	Heading, Lead      string
	Minutes, Days30    string
	Cards              []card
	Days               []dayRow
	IsCourse           bool
	T                  map[string]string
	IndexURL           string
}

var uiKeys = []string{
	"app.name", "home.courses", "home.day", "home.min", "landing.intro", "landing.how", "landing.start",
	"landing.start_note", "landing.plan", "landing.plan_note", "landing.light", "landing.all", "landing.a_day",
}

func base(l i18n.Lang, path func(string) string) page {
	p := page{Lang: string(l), HTMLLang: i18n.BCP47(l), Dir: i18n.Dir(l), T: map[string]string{}, IndexURL: IndexPath(string(l))}
	for _, k := range uiKeys {
		p.T[strings.ReplaceAll(k, ".", "_")] = i18n.T(l, k)
	}
	p.Canonical = Origin + path(string(l))
	p.Default = Origin + path(string(i18n.EN))
	for _, o := range i18n.Order {
		p.Alternates = append(p.Alternates, link{Lang: i18n.BCP47(o), Label: langName(o), URL: path(string(o)), Current: o == l})
	}
	return p
}

// langName is the language's own name for the picker ("Deutsch", "日本語").
func langName(l i18n.Lang) string {
	if m, ok := langs.Get(string(l)); ok {
		return m.Name
	}
	return strings.ToUpper(string(l))
}

func minutes(sec int) int {
	if m := (sec + 59) / 60; m > 1 {
		return m
	}
	return 1
}

// span is the length of a session on day 1 and on the peak day, e.g.
// "12–18 min a day".
func span(l i18n.Lang, r content.Resolved) string {
	first, peak := minutes(r.Workout(1, 0, pageRests).EstSec), minutes(r.Workout(content.TotalDays-1, 0, pageRests).EstSec)
	return i18n.Tf(l, "fmt.min_range", first, peak) + " " + i18n.T(l, "landing.a_day")
}

func indexPage(l i18n.Lang) page {
	p := base(l, IndexPath)
	p.Heading = i18n.T(l, "home.courses")
	p.Title = p.Heading + " — " + i18n.T(l, "app.name")
	p.Desc = i18n.T(l, "landing.intro")
	p.Lead = p.Desc
	for _, c := range content.Series30() {
		p.Cards = append(p.Cards, card{
			Title: c.TitleIn(string(l)), About: c.AboutIn(string(l)), URL: CoursePath(string(l), c),
			Minutes: span(l, c.Resolve()),
		})
	}
	return p
}

func coursePage(l i18n.Lang, c content.Course) page {
	p := base(l, func(lang string) string { return CoursePath(lang, c) })
	r := c.Resolve()
	p.IsCourse = true
	p.Days30 = i18n.Tn(l, "fmt.days", content.TotalDays)
	p.Heading = c.TitleIn(string(l))
	p.Title = p.Heading + " — " + i18n.T(l, "app.name")
	p.Desc = c.AboutIn(string(l))
	p.Lead = p.Desc
	p.Minutes = span(l, r)
	for _, d := range r.Days {
		row := dayRow{Label: i18n.Tf(l, "fmt.day", d.Day), Light: content.IsRecoveryDay(d.Day),
			Minutes: "≈ " + i18n.Tf(l, "fmt.min", minutes(r.Workout(d.Day, 0, pageRests).EstSec))}
		for _, it := range d.Items {
			row.Items = append(row.Items, dose(l, it))
		}
		p.Days = append(p.Days, row)
	}
	for _, o := range content.Series30() {
		if o.Key != c.Key {
			p.Cards = append(p.Cards, card{Title: o.TitleIn(string(l)), URL: CoursePath(string(l), o), Minutes: span(l, o.Resolve())})
		}
	}
	return p
}

// dose formats one station, e.g. "Forearm plank · 20 sec" or "Push-ups × 6".
func dose(l i18n.Lang, it content.RItem) string {
	name := it.ID
	if doc, ok := catalog.Builtin(it.ID); ok {
		if loc, ok := doc.Content[string(l)]; ok && loc.Name != "" {
			name = loc.Name
		} else if loc, ok := doc.Content["en"]; ok && loc.Name != "" {
			name = loc.Name
		}
	}
	var s string
	if it.Unit == content.Seconds {
		s = name + " · " + i18n.Tf(l, "fmt.sec", it.Value)
	} else {
		s = fmt.Sprintf("%s × %d", name, it.Value)
	}
	if it.PerSide {
		s += " (" + i18n.T(l, "workout.per_side") + ")"
	}
	return s
}

func render(w http.ResponseWriter, p page) {
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, p); err != nil {
		http.Error(w, "page unavailable", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Content-Language", p.HTMLLang)
	w.Header().Set("Cache-Control", "public, max-age=3600")
	_, _ = w.Write(buf.Bytes())
}

var tmpl = template.Must(template.New("page").Parse(`<!doctype html>
<html lang="{{.HTMLLang}}" dir="{{.Dir}}">
<head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<title>{{.Title}}</title>
<meta name="description" content="{{.Desc}}">
<link rel="canonical" href="{{.Canonical}}">
{{range .Alternates}}<link rel="alternate" hreflang="{{.Lang}}" href="` + Origin + `{{.URL}}">
{{end}}<link rel="alternate" hreflang="x-default" href="{{.Default}}">
<meta property="og:type" content="website">
<meta property="og:title" content="{{.Title}}">
<meta property="og:description" content="{{.Desc}}">
<meta property="og:url" content="{{.Canonical}}">
<meta property="og:image" content="` + Origin + `/static/icon-512.png">
<meta name="theme-color" content="#0f172a">
<link rel="icon" href="/static/icon.svg">
<style>
:root{--bg:#0f172a;--card:#1e293b;--card2:#273449;--fg:#f1f5f9;--muted:#94a3b8;--accent:#fb923c;--accent2:#f97316;--ok:#34d399;--line:#334155}
*{box-sizing:border-box}
body{margin:0;background:var(--bg);color:var(--fg);font:16px/1.5 -apple-system,BlinkMacSystemFont,"Segoe UI",Roboto,sans-serif}
a{color:inherit}
.wrap{max-width:720px;margin:0 auto;padding:16px 16px 48px}
.top{display:flex;justify-content:space-between;align-items:center;gap:12px;margin-bottom:20px}
.brand{font-weight:800;text-decoration:none}
.lang{appearance:none;-webkit-appearance:none;background:var(--card) url("data:image/svg+xml;utf8,<svg xmlns='http://www.w3.org/2000/svg' width='12' height='8' viewBox='0 0 12 8' fill='none' stroke='%2394a3b8' stroke-width='1.6' stroke-linecap='round' stroke-linejoin='round'><path d='M1 1.5 6 6.5 11 1.5'/></svg>") no-repeat right 12px center;color:var(--fg);border:1px solid var(--line);border-radius:10px;padding:8px 32px 8px 12px;font:inherit;font-size:14px;max-width:50%}
[dir=rtl] .lang{background-position:left 12px center;padding:8px 12px 8px 32px}
h1{font-size:30px;line-height:1.15;letter-spacing:-.5px;margin:0 0 12px;text-wrap:balance}
h2{font-size:14px;color:var(--muted);text-transform:uppercase;letter-spacing:.5px;margin:32px 0 10px}
.lead{font-size:17px;margin:0 0 12px}
.muted{color:var(--muted);font-size:14px;margin:0 0 12px}
.facts{display:flex;flex-wrap:wrap;gap:8px;margin:16px 0}
.facts span{background:var(--card);border-radius:999px;padding:6px 12px;font-size:14px}
.cta{display:inline-block;margin:8px 0 6px;padding:14px 26px;border-radius:999px;font-weight:700;font-size:18px;text-decoration:none;color:#1f1300;background:linear-gradient(135deg,var(--accent),var(--accent2))}
.cards{display:grid;grid-template-columns:repeat(auto-fill,minmax(210px,1fr));gap:10px}
.card{display:block;background:var(--card);border-radius:14px;padding:14px;text-decoration:none}
.card b{display:block;font-size:16px;line-height:1.25}
.card small{display:block;color:var(--muted);margin-top:4px}
.card p{margin:8px 0 0;font-size:14px;color:var(--muted)}
.days{list-style:none;margin:0;padding:0}
.days li{background:var(--card);border-radius:12px;padding:12px 14px;margin-bottom:8px}
.days .head{display:flex;justify-content:space-between;gap:12px;align-items:baseline}
.days .n{font-weight:700}
.days .light{color:var(--ok);font-size:13px;font-weight:600}
.days .min{color:var(--muted);font-size:13px;white-space:nowrap}
.days .items{margin:6px 0 0;font-size:14px;color:var(--muted)}
</style>
</head>
<body>
<div class="wrap">
<div class="top">
<a class="brand" href="/">🌅 {{.T.app_name}}</a>
<select class="lang" aria-label="Language" onchange="location.href=this.value">{{range .Alternates}}<option value="{{.URL}}" lang="{{.Lang}}"{{if .Current}} selected{{end}}>{{.Label}}</option>{{end}}</select>
</div>
<main>
<h1>{{.Heading}}</h1>
<p class="lead">{{.Lead}}</p>
{{if .IsCourse}}
<div class="facts"><span>{{.Days30}}</span><span>{{.Minutes}}</span></div>
<p class="muted">{{.T.landing_how}}</p>
<a class="cta" href="/">{{.T.landing_start}}</a>
<p class="muted">{{.T.landing_start_note}}</p>
<h2>{{.T.landing_plan}}</h2>
<p class="muted">{{.T.landing_plan_note}}</p>
<ol class="days">
{{range .Days}}<li>
<div class="head"><span class="n">{{.Label}}{{if .Light}} <span class="light">· {{$.T.landing_light}}</span>{{end}}</span><span class="min">{{.Minutes}}</span></div>
<p class="items">{{range $i, $it := .Items}}{{if $i}}, {{end}}{{$it}}{{end}}</p>
</li>
{{end}}</ol>
<a class="cta" href="/">{{.T.landing_start}}</a>
<h2>{{.T.landing_all}}</h2>
<div class="cards">
{{range .Cards}}<a class="card" href="{{.URL}}"><b>{{.Title}}</b><small>{{.Minutes}}</small></a>
{{end}}<a class="card" href="{{.IndexURL}}"><b>{{.T.landing_all}} →</b></a>
</div>
{{else}}
<p class="muted">{{.T.landing_how}}</p>
<div class="cards">
{{range .Cards}}<a class="card" href="{{.URL}}"><b>{{.Title}}</b><small>{{.Minutes}}</small><p>{{.About}}</p></a>
{{end}}</div>
<p><a class="cta" href="/">{{.T.landing_start}}</a></p>
{{end}}
</main>
</div>
</body>
</html>
`))
