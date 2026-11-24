package mcp

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"

	"morning30/internal/app"
	"morning30/internal/catalog"
	"morning30/internal/content"
	"morning30/internal/db"
)

type principalKey struct{}

type principal struct {
	Admin bool
	User  db.Account
}

// Handler serves Bearer-protected /api/*. A per-user token from settings is
// scoped to that account. An optional process-wide admin token (MORNING_MCP_TOKEN)
// may still list and edit any account.
func Handler(adminToken string) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/accounts", func(w http.ResponseWriter, r *http.Request) {
		p := mustPrincipal(r)
		if !p.Admin {
			writeJSON(w, []db.Account{p.User})
			return
		}
		app.SnapshotAll()
		accs, err := app.DB.ListAccounts()
		if err != nil {
			writeErr(w, http.StatusInternalServerError, err)
			return
		}
		if accs == nil {
			accs = []db.Account{}
		}
		writeJSON(w, accs)
	})
	mux.HandleFunc("GET /api/accounts/{who}/program", func(w http.ResponseWriter, r *http.Request) {
		acc, ok := lookupAccount(w, r)
		if !ok {
			return
		}
		writeJSON(w, programView{
			Account:  acc,
			Personal: app.ActiveProgramKey(acc.ID) == "custom",
			Program:  app.UserProgram(acc.ID),
		})
	})
	mux.HandleFunc("PUT /api/accounts/{who}/program", func(w http.ResponseWriter, r *http.Request) {
		acc, ok := lookupAccount(w, r)
		if !ok {
			return
		}
		body, err := io.ReadAll(io.LimitReader(r.Body, 1<<20))
		if err != nil {
			writeErr(w, http.StatusBadRequest, err)
			return
		}
		if err := app.ReplaceUserProgram(acc.ID, string(body)); err != nil {
			writeErr(w, http.StatusBadRequest, err)
			return
		}
		acc, _ = app.DB.AccountByID(acc.ID)
		writeJSON(w, programView{Account: acc, Personal: true, Program: app.UserProgram(acc.ID)})
	})
	mux.HandleFunc("PUT /api/accounts/{who}/program/name", func(w http.ResponseWriter, r *http.Request) {
		acc, ok := lookupAccount(w, r)
		if !ok {
			return
		}
		body, err := io.ReadAll(io.LimitReader(r.Body, 1<<12))
		if err != nil {
			writeErr(w, http.StatusBadRequest, err)
			return
		}
		// The name may be sent as plain text or as {"name": "..."}.
		name := string(body)
		if trimmed := strings.TrimSpace(name); strings.HasPrefix(trimmed, "{") {
			var a renameArgs
			if err := json.Unmarshal([]byte(trimmed), &a); err != nil {
				writeErr(w, http.StatusBadRequest, err)
				return
			}
			name = a.Name
		}
		view, err := renameProgram(acc, name)
		if err != nil {
			writeErr(w, http.StatusBadRequest, err)
			return
		}
		writeJSON(w, view)
	})
	mux.HandleFunc("POST /api/accounts/{who}/course/{key}", func(w http.ResponseWriter, r *http.Request) {
		acc, ok := lookupAccount(w, r)
		if !ok {
			return
		}
		view, err := useCourse(acc, r.PathValue("key"), r.URL.Query().Get("lang"))
		if errors.Is(err, app.ErrUnknownCourse) {
			writeErr(w, http.StatusNotFound, err)
			return
		}
		if err != nil {
			writeErr(w, http.StatusBadRequest, err)
			return
		}
		writeJSON(w, view)
	})
	mux.HandleFunc("GET /api/courses", func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, courseList())
	})
	mux.HandleFunc("GET /api/courses/{key}", func(w http.ResponseWriter, r *http.Request) {
		view, err := courseDetail(r.PathValue("key"))
		if err != nil {
			writeErr(w, http.StatusNotFound, err)
			return
		}
		writeJSON(w, view)
	})
	mux.HandleFunc("GET /api/accounts/{who}/saved-programs", func(w http.ResponseWriter, r *http.Request) {
		acc, ok := lookupAccount(w, r)
		if !ok {
			return
		}
		list, err := app.DB.SavedPrograms(acc.ID)
		if err != nil {
			writeErr(w, http.StatusInternalServerError, err)
			return
		}
		writeJSON(w, list)
	})
	mux.HandleFunc("POST /api/accounts/{who}/saved-programs/{id}/restore", func(w http.ResponseWriter, r *http.Request) {
		acc, ok := lookupAccount(w, r)
		if !ok {
			return
		}
		id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
		if err != nil {
			writeErr(w, http.StatusBadRequest, err)
			return
		}
		view, err := restoreSaved(acc, id)
		if errors.Is(err, db.ErrSavedProgramNotFound) {
			writeErr(w, http.StatusNotFound, err)
			return
		}
		if err != nil {
			writeErr(w, http.StatusInternalServerError, err)
			return
		}
		writeJSON(w, view)
	})
	mux.HandleFunc("POST /api/accounts/{who}/restore-demo", func(w http.ResponseWriter, r *http.Request) {
		acc, ok := lookupAccount(w, r)
		if !ok {
			return
		}
		if err := app.RestoreDemo(acc.ID); err != nil {
			writeErr(w, http.StatusInternalServerError, err)
			return
		}
		acc, _ = app.DB.AccountByID(acc.ID)
		writeJSON(w, programView{Account: acc, Personal: false, Program: app.UserProgram(acc.ID)})
	})
	mux.HandleFunc("GET /api/accounts/{who}/exercises", func(w http.ResponseWriter, r *http.Request) {
		acc, ok := lookupAccount(w, r)
		if !ok {
			return
		}
		j, err := app.DB.GetCustomExercises(acc.ID)
		if err != nil {
			writeErr(w, http.StatusInternalServerError, err)
			return
		}
		writeJSON(w, catalog.New(j).Custom())
	})
	mux.HandleFunc("GET /api/accounts/{who}/exercises/{id}", func(w http.ResponseWriter, r *http.Request) {
		acc, ok := lookupAccount(w, r)
		if !ok {
			return
		}
		d, found := app.UserCatalog(acc.ID).Doc(r.PathValue("id"))
		if !found {
			http.NotFound(w, r)
			return
		}
		writeJSON(w, d)
	})
	mux.HandleFunc("PUT /api/accounts/{who}/exercises/{id}", func(w http.ResponseWriter, r *http.Request) {
		acc, ok := lookupAccount(w, r)
		if !ok {
			return
		}
		body, err := io.ReadAll(io.LimitReader(r.Body, 1<<20))
		if err != nil {
			writeErr(w, http.StatusBadRequest, err)
			return
		}
		var d catalog.Doc
		if err := json.Unmarshal(body, &d); err != nil {
			writeErr(w, http.StatusBadRequest, err)
			return
		}
		if d.ID == "" {
			d.ID = r.PathValue("id")
		}
		if err := app.SaveCustomExercise(acc.ID, d); err != nil {
			writeErr(w, http.StatusBadRequest, err)
			return
		}
		got, _ := app.UserCatalog(acc.ID).Doc(d.ID)
		writeJSON(w, got)
	})
	mux.HandleFunc("GET /api/figures", func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, content.FigureIDs())
	})
	mux.HandleFunc("GET /api/poses", func(w http.ResponseWriter, r *http.Request) {
		poses, err := app.DB.ListPoses()
		if err != nil {
			writeErr(w, http.StatusInternalServerError, err)
			return
		}
		if poses == nil {
			poses = []db.Pose{}
		}
		writeJSON(w, poses)
	})
	mux.HandleFunc("GET /api/poses/{id}", func(w http.ResponseWriter, r *http.Request) {
		p, err := app.DB.GetPose(r.PathValue("id"))
		if err != nil {
			http.NotFound(w, r)
			return
		}
		writeJSON(w, p)
	})
	mux.HandleFunc("PUT /api/poses/{id}", func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(io.LimitReader(r.Body, 6<<20))
		if err != nil {
			writeErr(w, http.StatusBadRequest, err)
			return
		}
		id := r.PathValue("id")
		kind := r.URL.Query().Get("kind")
		label := r.URL.Query().Get("label")
		sex := r.URL.Query().Get("sex")
		ct := r.Header.Get("Content-Type")
		var data []byte
		if strings.HasPrefix(ct, "image/") {
			data = body
		} else {
			var a poseArgs
			if err := json.Unmarshal(body, &a); err != nil {
				writeErr(w, http.StatusBadRequest, err)
				return
			}
			if a.ID != "" {
				id = a.ID
			}
			if a.Kind != "" {
				kind = a.Kind
			}
			if a.Label != "" {
				label = a.Label
			}
			if a.Sex != "" {
				sex = a.Sex
			}
			data, err = decodePoseImage(a.Image)
			if err != nil {
				writeErr(w, http.StatusBadRequest, err)
				return
			}
		}
		if err := app.DB.PutPose(id, kind, label, sex, data); err != nil {
			writeErr(w, http.StatusBadRequest, err)
			return
		}
		p, err := app.DB.GetPose(id)
		if err != nil {
			writeErr(w, http.StatusInternalServerError, err)
			return
		}
		writeJSON(w, p)
	})
	mux.HandleFunc("PUT /api/exercises/{id}/poses", func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(io.LimitReader(r.Body, 1<<16))
		if err != nil {
			writeErr(w, http.StatusBadRequest, err)
			return
		}
		var a setPosesArgs
		if err := json.Unmarshal(body, &a); err != nil {
			writeErr(w, http.StatusBadRequest, err)
			return
		}
		id := r.PathValue("id")
		if a.ID != "" {
			id = a.ID
		}
		if err := app.SetLibraryPoses(id, a.Poses); err != nil {
			writeErr(w, http.StatusBadRequest, err)
			return
		}
		d, ok := catalog.NewLayered(mustAllExercises(), "").Doc(id)
		if !ok {
			http.NotFound(w, r)
			return
		}
		writeJSON(w, d)
	})
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		got := strings.TrimSpace(strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer "))
		if adminToken != "" && got == adminToken {
			r = r.WithContext(context.WithValue(r.Context(), principalKey{}, principal{Admin: true}))
			mux.ServeHTTP(w, r)
			return
		}
		acc, err := app.DB.AccountByMCPToken(got)
		if err != nil {
			w.Header().Set("WWW-Authenticate", `Bearer realm="morning30-mcp"`)
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}
		r = r.WithContext(context.WithValue(r.Context(), principalKey{}, principal{User: acc}))
		mux.ServeHTTP(w, r)
	})
}

func mustPrincipal(r *http.Request) principal {
	p, _ := r.Context().Value(principalKey{}).(principal)
	return p
}

func lookupAccount(w http.ResponseWriter, r *http.Request) (db.Account, bool) {
	app.SnapshotAll()
	p := mustPrincipal(r)
	who := strings.TrimSpace(r.PathValue("who"))
	if !p.Admin {
		if who != "" && who != "me" && !sameAccount(p.User, who) {
			writeErr(w, http.StatusForbidden, errors.New("token is scoped to this account"))
			return db.Account{}, false
		}
		acc, err := app.DB.AccountByID(p.User.ID)
		if err != nil {
			writeErr(w, http.StatusNotFound, err)
			return db.Account{}, false
		}
		return acc, true
	}
	acc, err := app.DB.FindAccount(who)
	if err != nil {
		code := http.StatusNotFound
		if errors.Is(err, db.ErrAccountAmbiguous) {
			code = http.StatusConflict
		}
		writeErr(w, code, err)
		return db.Account{}, false
	}
	return acc, true
}

func sameAccount(a db.Account, who string) bool {
	if who == strconv.FormatInt(a.ID, 10) {
		return true
	}
	if a.Email != "" && strings.EqualFold(who, a.Email) {
		return true
	}
	return a.Name != "" && strings.EqualFold(who, a.Name)
}

func writeJSON(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	_ = enc.Encode(v)
}

func writeErr(w http.ResponseWriter, code int, err error) {
	http.Error(w, err.Error(), code)
}
