package findfile

import (
	"testing"
)

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
