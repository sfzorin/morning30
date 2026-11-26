package app

import "morning30/internal/db"

// Figures shown in pictures: a man, a woman, a boy and a girl. A child figure
// without its own picture falls back to the adult of the same sex.
var Figures = []string{"m", "f", "b", "g"}

// childFigure and adultFigure pair each figure with its counterpart.
var childFigure = map[string]string{"m": "b", "f": "g", "b": "b", "g": "g"}
var adultFigure = map[string]string{"b": "m", "g": "f", "m": "m", "f": "f"}

// FigureFor returns the figure to show for a program: a children's program
// shows the child of the same sex, any other program the adult.
func FigureFor(current string, kids bool) string {
	current = db.NormalizeSex(current)
	if kids {
		return childFigure[current]
	}
	return adultFigure[current]
}

// adjustFigure switches the user's figure when they pick a ready program for
// children or leave one. They can change it back in settings.
func adjustFigure(userID int64, kids bool) {
	current := DB.GetSex(userID)
	if next := FigureFor(current, kids); next != current {
		_ = DB.SetSex(userID, next)
	}
}
