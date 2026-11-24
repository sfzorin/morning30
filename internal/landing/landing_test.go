package landing

import (
	"html/template"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"morning30/internal/content"
	"morning30/internal/i18n"
)

// newMux mirrors the server: the app owns "/" and the agent API owns "/api/".
// Registration must not collide with either.
func newMux(t *testing.T) *http.ServeMux {
	t.Helper()
	mux := http.NewServeMux()
	mux.HandleFunc("/api/", func(w http.ResponseWriter, _ *http.Request) { w.WriteHeader(http.StatusTeapot) })
	mux.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) { _, _ = w.Write([]byte("app")) })
	Register(mux)
	return mux
}

func get(mux *http.ServeMux, method, path string, header ...string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, path, nil)
	for i := 0; i+1 < len(header); i += 2 {
		req.Header.Set(header[i], header[i+1])
	}
	rr := httptest.NewRecorder()
	mux.ServeHTTP(rr, req)
	return rr
}

func TestEveryCoursePageInEveryLanguage(t *testing.T) {
	mux := newMux(t)
	for _, l := range i18n.Order {
		lang := string(l)
		rr := get(mux, "GET", IndexPath(lang))
		body := rr.Body.String()
		if rr.Code != http.StatusOK || !strings.Contains(body, `<html lang="`+i18n.BCP47(l)+`"`) {
			t.Fatalf("%s index: %d", lang, rr.Code)
		}
		for _, c := range content.Series30() {
			if !strings.Contains(body, `href="`+CoursePath(lang, c)+`"`) {
				t.Errorf("%s index does not link %s", lang, c.Key)
			}
			rr := get(mux, "GET", CoursePath(lang, c))
			page := rr.Body.String()
			if rr.Code != http.StatusOK {
				t.Fatalf("%s %s: %d", lang, c.Key, rr.Code)
			}
			for _, want := range []string{
				"<h1>" + escaped(c.Title[lang]) + "</h1>",
				`<link rel="canonical" href="` + Origin + CoursePath(lang, c) + `">`,
				`hreflang="x-default" href="` + Origin + CoursePath("en", c) + `"`,
				i18n.T(l, "landing.start"),
			} {
				if !strings.Contains(page, want) {
					t.Errorf("%s %s lacks %q", lang, c.Key, want)
				}
			}
			if n := strings.Count(page, `<link rel="alternate" hreflang=`); n != len(i18n.Order)+1 {
				t.Errorf("%s %s: %d alternates", lang, c.Key, n)
			}
			if n := strings.Count(page, `<div class="head">`); n != content.TotalDays {
				t.Errorf("%s %s: %d day rows", lang, c.Key, n)
			}
			// Every exercise is named in words, never by its internal id.
			for _, d := range c.Resolve().Days {
				for _, it := range d.Items {
					if name := i18n.Exercise(l, it.ID); name == it.ID || !strings.Contains(page, escaped(name)) {
						t.Errorf("%s %s: exercise %s is not named", lang, c.Key, it.ID)
					}
				}
			}
		}
	}
}

// escaped renders text the way the page template does ("+" included).
func escaped(s string) string {
	return strings.ReplaceAll(template.HTMLEscapeString(s), "+", "&#43;")
}

func TestRoutesStayOutOfTheAppsWay(t *testing.T) {
	mux := newMux(t)
	if rr := get(mux, "GET", "/ru/courses/nope-in-30-days"); rr.Code != http.StatusNotFound {
		t.Errorf("unknown course: %d", rr.Code)
	}
	rr := get(mux, "GET", "/courses", "Accept-Language", "de-DE,de;q=0.9")
	if rr.Code != http.StatusFound || rr.Header().Get("Location") != "/de/courses" {
		t.Errorf("redirect: %d %q", rr.Code, rr.Header().Get("Location"))
	}
	// Pages never set cookies: no session is created by a visit.
	if rr := get(mux, "GET", "/en/courses/plank-in-30-days"); len(rr.Result().Cookies()) != 0 {
		t.Error("course page set a cookie")
	}
	for path, want := range map[string]string{"/": "app", "/settings": "app", "/xx/courses": "app", "/day/3": "app"} {
		if rr := get(mux, "GET", path); rr.Body.String() != want {
			t.Errorf("%s no longer reaches the app", path)
		}
	}
	if rr := get(mux, "GET", "/api/courses"); rr.Code != http.StatusTeapot {
		t.Errorf("/api/courses no longer reaches the agent API: %d", rr.Code)
	}
}

func TestSitemapURLs(t *testing.T) {
	urls := SitemapURLs()
	if want := len(i18n.Order) * (1 + len(content.Series30())); len(urls) != want {
		t.Fatalf("%d urls, want %d", len(urls), want)
	}
	mux := newMux(t)
	seen := map[string]bool{}
	for _, u := range urls {
		if seen[u] || !strings.HasPrefix(u, Origin+"/") {
			t.Errorf("bad url %s", u)
		}
		seen[u] = true
		if rr := get(mux, "GET", strings.TrimPrefix(u, Origin)); rr.Code != http.StatusOK {
			t.Errorf("%s: %d", u, rr.Code)
		}
	}
}
