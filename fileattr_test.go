//go:build windows
// +build windows

package findfile

import (
	"os"
	"testing"

	"golang.org/x/sys/windows"
)

func TestGetFileAttributes(t *testing.T) {
	stat, err := os.Stat("fileattr_test.go")
	if err != nil {
		t.Fatal(err.Error())
	}
	attr := GetFileAttributes(stat)
	offbits := []uint32{
		windows.FILE_ATTRIBUTE_READONLY,
		windows.FILE_ATTRIBUTE_HIDDEN,
		windows.FILE_ATTRIBUTE_SYSTEM,
		windows.FILE_ATTRIBUTE_DIRECTORY,
	}
	for _, bit := range offbits {
		if (attr & bit) != 0 {
			t.Fatalf("normal file has invalid bit: %X", bit)
		}
	}
	stat, err = os.Stat(".")
	if err != nil {
		t.Fatal(err.Error())
	}
	attr = GetFileAttributes(stat)
	if (attr & windows.FILE_ATTRIBUTE_DIRECTORY) == 0 {
		t.Fatal("directory file does not have DIRECTORY-BIT")
	}
}
