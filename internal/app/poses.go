package app

import (
	"net/http"
)

// ServePose streams a catalog still. Public, like /static/ex — the player <img>
// cannot send a Bearer token. Missing id → 404.
func ServePose(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet && r.Method != http.MethodHead {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if DB == nil {
		http.NotFound(w, r)
		return
	}
	mime, data, err := DB.PoseImage(r.PathValue("id"), r.URL.Query().Get("sex"))
	if err != nil || len(data) == 0 {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", mime)
	w.Header().Set("Cache-Control", "public, max-age=86400")
	if r.Method == http.MethodHead {
		w.WriteHeader(http.StatusOK)
		return
	}
	_, _ = w.Write(data)
}
