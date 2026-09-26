package content

// StandardProgram is a named, read-only built-in exercise set.
type StandardProgram struct {
	Key   string          // stable id, stored as the user's active program
	Name  string          // display name (i18n overlay may replace this in the UI)
	Build func() Resolved // produces the program (regenerated on demand)
}

// StandardPrograms is the public built-in set. There is only the demo; every
// other plan is a per-user JSON file.
var StandardPrograms = []StandardProgram{
	{Key: "demo", Name: "Demo", Build: ResolveDemo},
}

// StandardByKey returns the Resolved program for a standard key.
func StandardByKey(key string) (Resolved, bool) {
	for _, p := range StandardPrograms {
		if p.Key == key {
			return p.Build(), true
		}
	}
	return Resolved{}, false
}

// DefaultStandardKey is the key of the default standard program (the demo).
func DefaultStandardKey() string { return StandardPrograms[0].Key }

// LegacyNamed returns a previously selectable named program so existing
// accounts can be snapshotted into a personal file. These are not in
// StandardPrograms.
func LegacyNamed(key string) (Resolved, bool) {
	switch key {
	case "ilya":
		return resolveIlya(), true
	case "vlad":
		return resolveVlad(), true
	case "sergey", "builtin":
		return ResolveBuiltin(), true
	}
	return Resolved{}, false
}
