package findfile

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var envPattern = regexp.MustCompile(`%\w+%`)

// ExpandEnv replace ~/ and ~\ to %HOME% or %USERPROFILE%,
// and %ENVIRONMENTVARIABLE% to its' value.
func ExpandEnv(pattern string) string {
	if strings.HasPrefix(pattern, `~/`) || strings.HasPrefix(pattern, `~\`) {
		home := os.Getenv("HOME")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		if home != "" {
			pattern = home + pattern[1:]
		}
	}
	pattern = envPattern.ReplaceAllStringFunc(pattern, func(m string) string {
		name := m[1 : len(m)-1]
		return os.Getenv(name)
	})
	return os.ExpandEnv(pattern)
}

// Glob expands filenames matching with wildcard-pattern.
func Glob(pattern string) ([]string, error) {
	pname := filepath.Base(pattern)
	if strings.IndexAny(pname, "*?") < 0 {
		return nil, nil
	}
	match := make([]string, 0, 100)
	dirname := filepath.Dir(pattern)
	pattern = ExpandEnv(pattern)
	err := Walk(pattern, func(findf *FileInfo) bool {
		name := findf.Name()
		if (!strings.HasPrefix(name, ".") || strings.HasPrefix(pname, ".")) && !findf.IsHidden() {
			match = append(match, filepath.Join(dirname, name))
		}
		return true
	})
	return match, err
}

// Globs expands filenames matching with wildcard-patterns.
func Globs(patterns []string) []string {
	result := make([]string, 0, len(patterns))
	for _, pattern1 := range patterns {
		matches, err := Glob(pattern1)
		if matches == nil || len(matches) <= 0 || err != nil {
			result = append(result, pattern1)
		} else {
			for _, s := range matches {
				result = append(result, s)
			}
		}
	}
	return result
}
