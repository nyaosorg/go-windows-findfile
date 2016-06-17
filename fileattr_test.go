package findfile

import (
	"testing"
)

func TestGetFileAttributes(t *testing.T) {
	Walk("*", func(f *FileInfo) bool {
		print(f.Name(), "...", GetFileAttributes(f), "\n")
		return true
	})
}
