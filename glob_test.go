package findfile

import (
	"testing"
)

func TestGlob(t *testing.T) {
	println("*** part-i ***")
	Glob("*")
	println("*** part-ii ***")
	Glob("")
}
