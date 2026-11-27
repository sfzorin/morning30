// Managed by GoX v0.1.36

//line app.gox:1
// App is the root component: HTML shell, head (PWA meta + styles), and the router.
package main

import (
	"morning30/assets"
	"morning30/components"
	"morning30/internal/auth"
	"morning30/path"
	"morning30/segments/pages"

	"github.com/doors-dev/doors"
	"github.com/doors-dev/gox"
)

// App holds the shared reactive session created by the page factory.
type App struct {
	auth doors.Source[auth.Session]
}

//line app.gox:20
func (a App) Main() gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
		__e = __c.Raw("<!doctype html>"); if __e != nil { return }
		__e = __c.Init("html"); if __e != nil { return }
		{
//line app.gox:22
			__e = __c.Set("lang", "en"); if __e != nil { return }
			__e = __c.Submit(); if __e != nil { return }
			__e = __c.Init("head"); if __e != nil { return }
			{
				__e = __c.Submit(); if __e != nil { return }
				__e = __c.InitVoid("meta"); if __e != nil { return }
				{
//line app.gox:24
					__e = __c.Set("charset", "utf-8"); if __e != nil { return }
				}
				__e = __c.Submit(); if __e != nil { return }
				__e = __c.InitVoid("meta"); if __e != nil { return }
				{
//line app.gox:26
					__e = __c.Set("name", "viewport"); if __e != nil { return }
//line app.gox:27
					__e = __c.Set("content", "width=device-width, initial-scale=1, maximum-scale=1, viewport-fit=cover"); if __e != nil { return }
				}
				__e = __c.Submit(); if __e != nil { return }
				__e = __c.InitVoid("meta"); if __e != nil { return }
				{
//line app.gox:28
					__e = __c.Set("name", "theme-color"); if __e != nil { return }
//line app.gox:28
					__e = __c.Set("content", "#0f172a"); if __e != nil { return }
				}
				__e = __c.Submit(); if __e != nil { return }
				__e = __c.InitVoid("meta"); if __e != nil { return }
				{
//line app.gox:29
					__e = __c.Set("name", "mobile-web-app-capable"); if __e != nil { return }
//line app.gox:29
					__e = __c.Set("content", "yes"); if __e != nil { return }
				}
				__e = __c.Submit(); if __e != nil { return }
				__e = __c.InitVoid("meta"); if __e != nil { return }
				{
//line app.gox:30
					__e = __c.Set("name", "apple-mobile-web-app-capable"); if __e != nil { return }
//line app.gox:30
					__e = __c.Set("content", "yes"); if __e != nil { return }
				}
				__e = __c.Submit(); if __e != nil { return }
				__e = __c.InitVoid("meta"); if __e != nil { return }
				{
//line app.gox:31
					__e = __c.Set("name", "apple-mobile-web-app-status-bar-style"); if __e != nil { return }
//line app.gox:31
					__e = __c.Set("content", "black-translucent"); if __e != nil { return }
				}
				__e = __c.Submit(); if __e != nil { return }
				__e = __c.InitVoid("meta"); if __e != nil { return }
				{
//line app.gox:32
					__e = __c.Set("name", "apple-mobile-web-app-title"); if __e != nil { return }
//line app.gox:32
					__e = __c.Set("content", "morning30.com"); if __e != nil { return }
				}
				__e = __c.Submit(); if __e != nil { return }
				__e = __c.InitVoid("meta"); if __e != nil { return }
				{
//line app.gox:35
					__e = __c.Set("name", "description"); if __e != nil { return }
//line app.gox:35
					__e = __c.Set("content", "A free, voice-guided 30-day morning workout you can do anywhere with just a mat. Daily progression, streaks, male/female figures, and 28 languages. Open-source PWA."); if __e != nil { return }
				}
				__e = __c.Submit(); if __e != nil { return }
				__e = __c.InitVoid("meta"); if __e != nil { return }
				{
//line app.gox:36
					__e = __c.Set("name", "application-name"); if __e != nil { return }
//line app.gox:36
					__e = __c.Set("content", "morning30.com"); if __e != nil { return }
				}
				__e = __c.Submit(); if __e != nil { return }
				__e = __c.InitVoid("meta"); if __e != nil { return }
				{
//line app.gox:37
					__e = __c.Set("name", "robots"); if __e != nil { return }
//line app.gox:37
					__e = __c.Set("content", "index, follow"); if __e != nil { return }
				}
				__e = __c.Submit(); if __e != nil { return }
				__e = __c.InitVoid("link"); if __e != nil { return }
				{
//line app.gox:38
					__e = __c.Set("rel", "canonical"); if __e != nil { return }
//line app.gox:38
					__e = __c.Set("href", "https://morning30.com/"); if __e != nil { return }
				}
				__e = __c.Submit(); if __e != nil { return }
				__e = __c.InitVoid("meta"); if __e != nil { return }
				{
//line app.gox:39
					__e = __c.Set("property", "og:type"); if __e != nil { return }
//line app.gox:39
					__e = __c.Set("content", "website"); if __e != nil { return }
				}
				__e = __c.Submit(); if __e != nil { return }
				__e = __c.InitVoid("meta"); if __e != nil { return }
				{
//line app.gox:40
					__e = __c.Set("property", "og:site_name"); if __e != nil { return }
//line app.gox:40
					__e = __c.Set("content", "morning30.com"); if __e != nil { return }
				}
				__e = __c.Submit(); if __e != nil { return }
				__e = __c.InitVoid("meta"); if __e != nil { return }
				{
//line app.gox:41
					__e = __c.Set("property", "og:title"); if __e != nil { return }
//line app.gox:41
					__e = __c.Set("content", "morning30.com — 30-day morning workout"); if __e != nil { return }
				}
				__e = __c.Submit(); if __e != nil { return }
				__e = __c.InitVoid("meta"); if __e != nil { return }
				{
//line app.gox:42
					__e = __c.Set("property", "og:description"); if __e != nil { return }
//line app.gox:42
					__e = __c.Set("content", "A free, voice-guided 30-day morning workout you can do anywhere with just a mat. Daily progression, streaks, and 28 languages."); if __e != nil { return }
				}
				__e = __c.Submit(); if __e != nil { return }
				__e = __c.InitVoid("meta"); if __e != nil { return }
				{
//line app.gox:43
					__e = __c.Set("property", "og:url"); if __e != nil { return }
//line app.gox:43
					__e = __c.Set("content", "https://morning30.com/"); if __e != nil { return }
				}
				__e = __c.Submit(); if __e != nil { return }
				__e = __c.InitVoid("meta"); if __e != nil { return }
				{
//line app.gox:44
					__e = __c.Set("property", "og:image"); if __e != nil { return }
//line app.gox:44
					__e = __c.Set("content", "https://morning30.com/static/icon-512.png"); if __e != nil { return }
				}
				__e = __c.Submit(); if __e != nil { return }
				__e = __c.InitVoid("meta"); if __e != nil { return }
				{
//line app.gox:45
					__e = __c.Set("property", "og:image:width"); if __e != nil { return }
//line app.gox:45
					__e = __c.Set("content", "512"); if __e != nil { return }
				}
				__e = __c.Submit(); if __e != nil { return }
				__e = __c.InitVoid("meta"); if __e != nil { return }
				{
//line app.gox:46
					__e = __c.Set("property", "og:image:height"); if __e != nil { return }
//line app.gox:46
					__e = __c.Set("content", "512"); if __e != nil { return }
				}
				__e = __c.Submit(); if __e != nil { return }
				__e = __c.InitVoid("meta"); if __e != nil { return }
				{
//line app.gox:47
					__e = __c.Set("name", "twitter:card"); if __e != nil { return }
//line app.gox:47
					__e = __c.Set("content", "summary"); if __e != nil { return }
				}
				__e = __c.Submit(); if __e != nil { return }
				__e = __c.InitVoid("meta"); if __e != nil { return }
				{
//line app.gox:48
					__e = __c.Set("name", "twitter:title"); if __e != nil { return }
//line app.gox:48
					__e = __c.Set("content", "morning30.com — 30-day morning workout"); if __e != nil { return }
				}
				__e = __c.Submit(); if __e != nil { return }
				__e = __c.InitVoid("meta"); if __e != nil { return }
				{
//line app.gox:49
					__e = __c.Set("name", "twitter:description"); if __e != nil { return }
//line app.gox:49
					__e = __c.Set("content", "A free, voice-guided 30-day morning workout you can do anywhere with just a mat."); if __e != nil { return }
				}
				__e = __c.Submit(); if __e != nil { return }
				__e = __c.InitVoid("meta"); if __e != nil { return }
				{
//line app.gox:50
					__e = __c.Set("name", "twitter:image"); if __e != nil { return }
//line app.gox:50
					__e = __c.Set("content", "https://morning30.com/static/icon-512.png"); if __e != nil { return }
				}
				__e = __c.Submit(); if __e != nil { return }
				__e = __c.InitVoid("link"); if __e != nil { return }
				{
//line app.gox:52
					__e = __c.Set("rel", "manifest"); if __e != nil { return }
//line app.gox:52
					__e = __c.Set("href", "/static/manifest.webmanifest"); if __e != nil { return }
				}
				__e = __c.Submit(); if __e != nil { return }
				__e = __c.InitVoid("link"); if __e != nil { return }
				{
//line app.gox:53
					__e = __c.Set("rel", "icon"); if __e != nil { return }
//line app.gox:53
					__e = __c.Set("type", "image/svg+xml"); if __e != nil { return }
//line app.gox:53
					__e = __c.Set("href", "/static/icon.svg"); if __e != nil { return }
				}
				__e = __c.Submit(); if __e != nil { return }
				__e = __c.InitVoid("link"); if __e != nil { return }
				{
//line app.gox:54
					__e = __c.Set("rel", "icon"); if __e != nil { return }
//line app.gox:54
					__e = __c.Set("type", "image/png"); if __e != nil { return }
//line app.gox:54
					__e = __c.Set("sizes", "192x192"); if __e != nil { return }
//line app.gox:54
					__e = __c.Set("href", "/static/icon-192.png"); if __e != nil { return }
				}
				__e = __c.Submit(); if __e != nil { return }
				__e = __c.InitVoid("link"); if __e != nil { return }
				{
//line app.gox:55
					__e = __c.Set("rel", "apple-touch-icon"); if __e != nil { return }
//line app.gox:55
					__e = __c.Set("href", "/static/icon-180.png"); if __e != nil { return }
				}
				__e = __c.Submit(); if __e != nil { return }
				__e = __c.InitVoid("link"); if __e != nil { return }
				{
//line app.gox:56
					__e = __c.Set("rel", "stylesheet"); if __e != nil { return }
//line app.gox:56
					__e = __c.Set("href", assets.Style); if __e != nil { return }
//line app.gox:56
					__e = __c.Set("name", "main.css"); if __e != nil { return }
				}
				__e = __c.Submit(); if __e != nil { return }
				__e = __c.Init("title"); if __e != nil { return }
				{
					__e = __c.Submit(); if __e != nil { return }
					__e = __c.Text("morning30.com — 30-day morning workout"); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
				__e = __c.Init("script"); if __e != nil { return }
				{
//line app.gox:58
					__e = __c.Set("type", "application/ld+json"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
					__e = __c.Raw("{\"@context\":\"https://schema.org\",\"@type\":\"WebApplication\",\"name\":\"morning30.com\",\"url\":\"https://morning30.com/\",\"applicationCategory\":\"HealthApplication\",\"operatingSystem\":\"Web\",\"browserRequirements\":\"Requires JavaScript\",\"offers\":{\"@type\":\"Offer\",\"price\":\"0\",\"priceCurrency\":\"USD\"},\"description\":\"A free, voice-guided 30-day morning workout you can do anywhere with just a mat.\",\"inLanguage\":[\"en\",\"ru\",\"tr\",\"de\",\"es\",\"fr\",\"it\",\"pt\",\"pl\",\"nl\",\"sv\",\"da\",\"nb\",\"fi\",\"cs\",\"ro\",\"hu\",\"el\",\"ja\",\"ko\",\"zh-Hans\",\"zh-Hant\",\"id\",\"vi\",\"th\",\"ar\",\"he\",\"fa\"]}"); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
			__e = __c.Init("body"); if __e != nil { return }
			{
				__e = __c.Submit(); if __e != nil { return }
//line app.gox:63
				__e = __c.Any(doors.Route(
				doors.RouteModel(func(p doors.Source[path.Path]) gox.Comp {
					return pages.Root{Path: p, Auth: a.auth}
				}),
				doors.RouteLocationDefault(components.NotFound),
			)); if __e != nil { return }
				__e = __c.Init("script"); if __e != nil { return }
				{
					__e = __c.Submit(); if __e != nil { return }
					__e = __c.Raw("if (\"serviceWorker\" in navigator) {\n\t\t\t\t\tnavigator.serviceWorker.register(\"/static/sw.js\").catch(function () {})\n\t\t\t\t}\n\t\t\t\t// A page left in the background for long has been forgotten by the\n\t\t\t\t// server (doors drops idle pages after ~40 min, and deploys restart it);\n\t\t\t\t// its first tap would only reload it. Reload as soon as it is back.\n\t\t\t\t(function () {\n\t\t\t\t\tvar hiddenAt = document.hidden ? Date.now() : 0\n\t\t\t\t\tvar stale = 15 * 60 * 1000\n\t\t\t\t\t// Never during a workout or the fitness test: those keep their own\n\t\t\t\t\t// state, and a reload there would interrupt.\n\t\t\t\t\tfunction busy() { return /^\\/(day\\/|start$|test$)/.test(location.pathname) }\n\t\t\t\t\tdocument.addEventListener(\"visibilitychange\", function () {\n\t\t\t\t\t\tif (document.hidden) { hiddenAt = Date.now(); return }\n\t\t\t\t\t\tif (hiddenAt && Date.now() - hiddenAt > stale && !busy()) location.reload()\n\t\t\t\t\t\thiddenAt = 0\n\t\t\t\t\t})\n\t\t\t\t\twindow.addEventListener(\"pageshow\", function (e) {\n\t\t\t\t\t\tif (e.persisted && !busy()) location.reload()\n\t\t\t\t\t})\n\t\t\t\t})()"); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
		}
		__e = __c.Close(); if __e != nil { return }
	return })
//line app.gox:95
}
