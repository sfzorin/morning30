package app

import (
	"net/http"
	"slices"
	"strings"

	"morning30/internal/content"
	"morning30/internal/db"
)

// ActiveExercisePoses lists uploaded illustrations available for custom exercises.
func ActiveExercisePoses() ([]db.Pose, error) {
	poses, err := DB.ListPoses()
	return slices.DeleteFunc(poses, func(p db.Pose) bool { return retiredPose(p.ID) }), err
}

// Legacy account bundles and the original mannequin stills use reserved pose IDs.
func retiredPose(id string) bool {
	id = db.NormalizePoseID(id)
	if strings.HasPrefix(id, "sergey-man-") || strings.HasPrefix(id, "ilya-boy-") {
		return true
	}
	exercise, frame, ok := strings.Cut(id, "-")
	if !ok || (frame != "1" && frame != "2") {
		return false
	}
	_, builtin := content.Get(strings.ToUpper(exercise))
	return builtin
}

// ServePose streams an uploaded illustration for a custom exercise.
func ServePose(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet && r.Method != http.MethodHead {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	id := r.PathValue("id")
	if DB == nil || retiredPose(id) {
		http.NotFound(w, r)
		return
	}
	mime, data, err := DB.PoseImage(id, r.URL.Query().Get("sex"))
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
