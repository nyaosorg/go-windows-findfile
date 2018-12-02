package findfile

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

type FileInfo struct {
	os.FileInfo
}

func (fi *FileInfo) IsReadOnly() bool {
	return (fi.Mode() & 0200) == 0
}

func (this *FileInfo) IsHidden() bool {
	return false
}

func (fi *FileInfo) IsSystem() bool {
	return false
}

func (fi *FileInfo) IsReparsePoint() bool {
	return false
}

func Walk(pattern string, callback func(*FileInfo) bool) error {
	dir := filepath.Dir(pattern)
	fnamepattern := filepath.Base(pattern)

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, file := range files {
		match, err := filepath.Match(fnamepattern, file.Name())
		if err != nil {
			return err
		}
		if match && !callback(&FileInfo{file}) {
			break
		}
	}
	return nil
}
