package landing

import (
	"encoding/json"
	"html/template"
	"net/http"
	"os"
	"strings"

	"morning30/internal/i18n"
)

// Privacy policy (/privacy) and the Android app link file
// (/.well-known/assetlinks.json) — both required for the app stores.

type privacyText struct {
	Title, Updated string
	Sections       []privacySection
}

type privacySection struct {
	Heading string
	Body    []string
}

var privacyTexts = map[string]privacyText{
	"en": {
		Title:   "Privacy policy",
		Updated: "Last updated: 26 November 2026",
		Sections: []privacySection{
			{"Who we are", []string{"morning30.com is a free, open-source home workout app (website and apps for iPhone and Android). The code is public at github.com/sfzorin/morning30."}},
			{"What we store", []string{
				"If you sign up: your email, a one-way hash of your password, the name you enter and the language you choose.",
				"If you continue without signing up: an anonymous guest account tied to a cookie on your device.",
				"For everyone: your settings (voice, figure, rest times, difficulty, rounds), your programs and your own exercises, and the days you worked out, so the app can show your streak and continue your program.",
				"If you fill in \"About you\" or take the fitness test: your age, height, weight, what bothers you (back, knees, blood pressure, feet) and your test results. They are used only to pick a program and its load for you.",
				"A session cookie that keeps you signed in. We use no advertising, no analytics and no tracking cookies, and we do not sell or share your data.",
			}},
			{"On your phone", []string{
				"The apps keep the morning reminder time on your phone and schedule the notification there; it is never sent to us.",
				"Voice cues are spoken by your phone's built-in speech engine.",
				"With your permission the apps write each finished workout (start, end, type) to Apple Health on iPhone or Health Connect on Android. They only write workouts there and never read your health data; nothing from Apple Health or Health Connect is sent to us.",
			}},
			{"AI coach", []string{"If you write to the AI coach, your message, your current program and your \"About you\" answers are sent to the AI model's maker (DeepSeek or Alibaba Qwen, through OpenRouter) to prepare the answer. Your name, email and other programs are not sent. The conversation is kept with the program and deleted with your account."}},
			{"How long we keep it", []string{"Until you delete your account. You can do that at any time in Settings → Delete account, on the website and in the apps. Deleting removes your account and all its data from our database at once."}},
			{"Health information", []string{"Workouts, notes about how they went and your fitness answers are used only to run the app for you. The app does not give medical advice; programs for back and foot problems are not a treatment."}},
			{"Children", []string{"Programs for children are meant to be used with a parent. We do not knowingly collect any personal data from children beyond an anonymous guest account."}},
			{"Contact", []string{"Questions or requests: open an issue at github.com/sfzorin/morning30."}},
		},
	},
	"ru": {
		Title:   "Политика конфиденциальности",
		Updated: "Обновлено 26 ноября 2026 г.",
		Sections: []privacySection{
			{"Кто мы", []string{"morning30.com — бесплатное приложение с открытым кодом для домашних тренировок: сайт и приложения для iPhone и Android. Код открыт на github.com/sfzorin/morning30."}},
			{"Что мы храним", []string{
				"Если вы зарегистрировались: эл. почту, необратимый хеш пароля, имя, которое вы указали, и выбранный язык.",
				"Если вы занимаетесь без регистрации: анонимный гостевой аккаунт, привязанный к cookie на вашем устройстве.",
				"Для всех: ваши настройки (голос, фигура, паузы, сложность, круги), ваши программы и свои упражнения, а также дни тренировок — чтобы показывать серию и продолжать программу.",
				"Если вы заполнили «О себе» или прошли тест формы: возраст, рост, вес, что беспокоит (спина, колени, давление, стопы) и результаты теста. Они нужны только для подбора программы и нагрузки.",
				"Cookie сессии, чтобы вы оставались в аккаунте. Мы не используем рекламу, аналитику и следящие cookie, не продаём и не передаём ваши данные.",
			}},
			{"На вашем телефоне", []string{
				"Время утреннего напоминания хранится на телефоне, и уведомление ставит сам телефон; к нам оно не отправляется.",
				"Голосовые подсказки произносит встроенный синтезатор речи телефона.",
				"С вашего разрешения приложения записывают каждую завершённую тренировку (начало, конец, вид) в «Здоровье» на iPhone или в Health Connect на Android. Они только записывают тренировки и никогда не читают ваши данные о здоровье; из «Здоровья» и Health Connect к нам ничего не попадает.",
			}},
			{"ИИ-тренер", []string{"Если вы пишете ИИ-тренеру, ваше сообщение, текущая программа и ответы из «О себе» отправляются разработчику модели ИИ (DeepSeek или Alibaba Qwen, через OpenRouter), чтобы подготовить ответ. Имя, эл. почта и другие программы не отправляются. Переписка хранится вместе с программой и удаляется вместе с аккаунтом."}},
			{"Сколько мы храним данные", []string{"Пока вы не удалите аккаунт. Это можно сделать в любой момент: Настройки → Удалить аккаунт, на сайте и в приложениях. Удаление сразу стирает аккаунт и все его данные из нашей базы."}},
			{"Данные о здоровье", []string{"Тренировки, отметки о самочувствии и ответы о себе используются только для работы приложения. Приложение не даёт медицинских советов; программы для спины и стоп не заменяют лечение."}},
			{"Дети", []string{"Детские программы рассчитаны на занятия вместе с родителями. Мы сознательно не собираем персональные данные детей, кроме анонимного гостевого аккаунта."}},
			{"Связь", []string{"Вопросы и запросы: откройте issue на github.com/sfzorin/morning30."}},
		},
	},
}

var privacyTmpl = template.Must(template.New("privacy").Parse(`<!doctype html>
<html lang="{{.Lang}}" dir="{{.Dir}}"><head><meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<title>{{.Text.Title}} — morning30.com</title>
<link rel="canonical" href="` + Origin + `/privacy">
<style>body{font-family:system-ui,-apple-system,sans-serif;background:#0f172a;color:#e2e8f0;line-height:1.6;margin:0;padding:24px 16px}
main{max-width:680px;margin:0 auto}h1{font-size:26px}h2{font-size:18px;margin-top:28px}a{color:#fb923c}.muted{color:#94a3b8}</style>
</head><body><main>
<p><a href="/">morning30.com</a></p>
<h1>{{.Text.Title}}</h1><p class="muted">{{.Text.Updated}}</p>
{{range .Text.Sections}}<h2>{{.Heading}}</h2>{{range .Body}}<p>{{.}}</p>{{end}}{{end}}
</main></body></html>`))

func servePrivacy(w http.ResponseWriter, r *http.Request) {
	lang := i18n.FromAcceptLanguage(r.Header.Get("Accept-Language"))
	if q := r.URL.Query().Get("lang"); q != "" {
		lang = i18n.Parse(q)
	}
	text, ok := privacyTexts[string(lang)]
	if !ok {
		lang, text = i18n.EN, privacyTexts["en"]
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Cache-Control", "public, max-age=3600")
	_ = privacyTmpl.Execute(w, map[string]any{"Lang": string(lang), "Dir": i18n.Dir(lang), "Text": text})
}

// serveAssetLinks proves to Android that the app and the site belong together,
// so the app can open the site full screen. ANDROID_CERT_SHA256 holds the
// signing certificate fingerprints (comma-separated) from Google Play.
func serveAssetLinks(w http.ResponseWriter, _ *http.Request) {
	pkg := strings.TrimSpace(os.Getenv("ANDROID_PACKAGE"))
	if pkg == "" {
		pkg = "com.morning30.app"
	}
	var prints []string
	for _, p := range strings.Split(os.Getenv("ANDROID_CERT_SHA256"), ",") {
		if p = strings.TrimSpace(p); p != "" {
			prints = append(prints, p)
		}
	}
	w.Header().Set("Content-Type", "application/json")
	if len(prints) == 0 {
		_, _ = w.Write([]byte("[]"))
		return
	}
	_ = json.NewEncoder(w).Encode([]map[string]any{{
		"relation": []string{"delegate_permission/common.handle_all_urls"},
		"target":   map[string]any{"namespace": "android_app", "package_name": pkg, "sha256_cert_fingerprints": prints},
	}})
}
