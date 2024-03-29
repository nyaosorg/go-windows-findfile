//go:build !windows
// +build !windows

package findfile

import (
	"context"
	"io/ioutil"
	"os"
	"path/filepath"
)

type _FileInfo struct {
	os.FileInfo
}

func (fi *_FileInfo) IsReadOnly() bool {
	return (fi.Mode() & 0200) == 0
}

func (this *_FileInfo) IsHidden() bool {
	return false
}

func (fi *_FileInfo) IsSystem() bool {
	return false
}

func (fi *_FileInfo) IsReparsePoint() bool {
	return false
}

func walk(ctx context.Context, pattern string, callback func(*_FileInfo) bool) error {
	dir := filepath.Dir(pattern)
	fnamepattern := filepath.Base(pattern)

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, file := range files {
		if ctx != nil {
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
			}
		}
		match, err := filepath.Match(fnamepattern, file.Name())
		if err != nil {
			return err
		}
		if match && !callback(&_FileInfo{file}) {
			break
		}
	}
	return nil
}
