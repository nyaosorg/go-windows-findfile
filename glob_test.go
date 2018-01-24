package findfile

import (
	"os"
	"testing"
)

func TestExpandEnv(t *testing.T) {
	temp := os.Getenv("TEMP")
	if ExpandEnv("%TEMP%") != temp {
		t.Fatal(`Fail: ExpandEnv("%TEMP%")`)
		return
	}
	if ExpandEnv("$TEMP") != temp {
		t.Fatal(`Fail: ExpandEnv("$TEMP")`)
		return
	}
	if ExpandEnv("${TEMP}") != temp {
		t.Fatal(`Fail: ExpandEnv("${TEMP}")`)
		return
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
