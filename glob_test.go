package findfile

import (
	"os"
	"path/filepath"
	"testing"
)

func TestExpandEnv(t *testing.T) {
	temp := os.Getenv("TEMP")
	home := os.Getenv("HOME")
	if home == "" {
		home = os.Getenv("USERPROFILE")
	}
	tests := [][2]string{
		{"%TEMP%", temp},
		{"$TEMP", temp},
		{"${TEMP}", temp},
		{`~`, home},
		{`~` + string(os.PathSeparator) + `foo`, filepath.Join(home, "foo")},
	}

	for _, p := range tests {
		if result := ExpandEnv(p[0]); result != p[1] {
			t.Fatalf("Fail: ExpandEnv('%s'):'%s' != '%s'", p[0], result, p[1])
		}
	}
}

func TestGlob(t *testing.T) {
	for _, pattern1 := range []string{"*", "", ".*"} {
		println("<test for '" + pattern1 + "'>")
		list, err := Glob(pattern1)
		if err != nil {
			t.Fatal(err)
		}
		if list != nil {
			for _, p := range list {
				println(p)
			}
		}
	}
}
