// Managed by GoX v0.1.36

//line router.gox:1
// Package pages holds the application's page components (router, auth, home,
// workout, settings) and their server handlers.
package pages

import (
	"morning30/assets"
	"morning30/internal/auth"
	"morning30/internal/i18n"
	"morning30/path"

	"github.com/doors-dev/doors"
	"github.com/doors-dev/gox"
)

// Root is the top dispatcher. It binds the shared session: unauthenticated
// visitors see the auth pages; authenticated users see the app.
type Root struct {
	Path doors.Source[path.Path]
	Auth doors.Source[auth.Session]
}

//line router.gox:22
func (r Root) Main() gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
//line router.gox:23
		__e = __c.Any(r.Auth.Bind(func(s auth.Session) gox.Elem {
		return gox.Elem(func(__c gox.Cursor) (__e error) {
			ctx := __c.Context(); _ = ctx
//line router.gox:25
			l := i18n.Lang(s.Lang)

			__e = __c.Init("div"); if __e != nil { return }
			{
//line router.gox:28
				__e = __c.Set("class", "app-lang"); if __e != nil { return }
//line router.gox:28
				__e = __c.Set("lang", i18n.BCP47(l)); if __e != nil { return }
//line router.gox:28
				__e = __c.Set("dir", i18n.Dir(l)); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
				__e = __c.Init("script"); if __e != nil { return }
				{
//line router.gox:30
					__e = __c.Set("src", assets.AppBridge); if __e != nil { return }
					__e = __c.Set("inline", true); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
					__e = __c.Raw(""); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
//line router.gox:31
				if !s.Authorized {
//line router.gox:32
					__e = __c.Any(r.publicRoutes(s)); if __e != nil { return }
				} else  {
//line router.gox:34
					__e = __c.Any(r.privateRoutes(s)); if __e != nil { return }
				}
			}
			__e = __c.Close(); if __e != nil { return }
		return })
//line router.gox:37
	})); if __e != nil { return }
	return })
//line router.gox:38
}

// publicRoutes: login (default) and register.
//line router.gox:41
func (r Root) publicRoutes(s auth.Session) gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
//line router.gox:42
		__e = __c.Any(r.Path.Route(
		doors.RouteMatch(func(p path.Path) bool { return p.Page == path.Register }).
			Comp(authPage{sess: s, auth: r.Auth, path: r.Path, register: true}),
		doors.RouteDefaultComp[path.Path](authPage{sess: s, auth: r.Auth, path: r.Path, register: false}),
	)); if __e != nil { return }
	return })
//line router.gox:47
}

// privateRoutes: settings, the workout player (/day/:n), and home (default).
// Guests can also reach the auth pages to upgrade to a real account.
//line router.gox:51
func (r Root) privateRoutes(s auth.Session) gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
//line router.gox:52
		__e = __c.Any(r.Path.Route(
		doors.RouteMatch(func(p path.Path) bool { return p.Page == path.Register }).
			Comp(authPage{sess: s, auth: r.Auth, path: r.Path, register: true}),
		doors.RouteMatch(func(p path.Path) bool { return p.Page == path.Login }).
			Comp(authPage{sess: s, auth: r.Auth, path: r.Path, register: false}),
		doors.RouteMatch(func(p path.Path) bool { return p.Page == path.Settings }).
			Comp(settingsPage{sess: s, auth: r.Auth, path: r.Path}),
		doors.RouteMatch(func(p path.Path) bool { return p.Page == path.Library }).
			Comp(libraryPage{sess: s, auth: r.Auth, path: r.Path}),
		doors.RouteMatch(func(p path.Path) bool { return p.Page == path.Today }).
			Comp(todayPage{sess: s, path: r.Path}),
		doors.RouteMatch(func(p path.Path) bool { return p.Page == path.Programs }).
			Comp(programsPage{sess: s, auth: r.Auth, path: r.Path}),
		doors.RouteMatch(func(p path.Path) bool { return p.Page == path.Editor }).
			Comp(settingsPage{sess: s, auth: r.Auth, path: r.Path, mode: "editor"}),
		doors.RouteMatch(func(p path.Path) bool { return p.Page == path.Program }).
			Comp(settingsPage{sess: s, auth: r.Auth, path: r.Path, mode: "program"}),
		doors.RouteMatch(func(p path.Path) bool { return p.Page == path.SettingsExercises }).
			Comp(settingsPage{sess: s, auth: r.Auth, path: r.Path, mode: "exercises"}),
		doors.RouteMatch(func(p path.Path) bool { return p.Page == path.SettingsAgent }).
			Comp(settingsPage{sess: s, auth: r.Auth, path: r.Path, mode: "agent"}),
		doors.RouteMatch(func(p path.Path) bool { return p.Page == path.Start }).
			Comp(workoutPage{sess: s, day: 0, path: r.Path}),
		doors.RouteDerive(func(p path.Path) (int, bool) {
			return p.Day, p.Page == path.Day
		}).Bind(func(day int) gox.Elem {
			return workoutPage{sess: s, day: day, path: r.Path}.Main()
		}),
		doors.RouteDefaultComp[path.Path](homePage{sess: s, auth: r.Auth, path: r.Path}),
	)); if __e != nil { return }
	return })
//line router.gox:82
}
