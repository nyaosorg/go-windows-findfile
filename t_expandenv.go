// +build ignore

package main

import (
	"github.com/zetamatta/go-findfile"
)

func main() {
	t := []string{
		`~`,
		`~\foo`,
		`~guest\hoge`,
		`~Administrator/foo`,
		`~public\foo`,
	}
	for _, s := range t {
		println(s, "=", findfile.ExpandEnv(s))
	}
}
