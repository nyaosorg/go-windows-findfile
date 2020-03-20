package findfile

import (
	"os"
	"path/filepath"
	"strings"
	"time"

	"golang.org/x/sys/windows"
)

// _FileInfo is the struct for Walk()'s parameter
type _FileInfo struct {
	windows.Win32finddata
	handle windows.Handle
}

func (fi *_FileInfo) clone() *FileInfo {
	return &_FileInfo{fi.Win32finddata, fi.handle}
}

func findFirst(pattern string) (*_FileInfo, error) {
	pattern16, err := windows.UTF16PtrFromString(pattern)
	if err != nil {
		return nil, err
	}
	this := new(_FileInfo)
	this.handle, err = windows.FindFirstFile(pattern16, &this.Win32finddata)
	if err != nil {
		return nil, err
	}
	return this, nil
}

// Name returns file's name
func (fi *_FileInfo) Name() string {
	return windows.UTF16ToString(fi.FileName[:])
}

// Size returns file's size by bytes
func (fi *_FileInfo) Size() int64 {
	return int64((int64(fi.FileSizeHigh) << 32) | int64(fi.FileSizeLow))
}

// ModTime returns timestamp when the file was last updated.
func (fi *_FileInfo) ModTime() time.Time {
	return time.Unix(0, fi.LastWriteTime.Nanoseconds())
}

// Mode emurates os.FileMode.
func (fi *_FileInfo) Mode() os.FileMode {
	m := os.FileMode(0444)
	if fi.IsDir() {
		m |= 0111 | os.ModeDir
	}
	if !fi.IsReadOnly() {
		m |= 0222
	}
	return m
}

// Sys returns underlying data source like os._FileInfo.Sys()
func (fi *_FileInfo) Sys() interface{} {
	return &fi.Win32finddata
}

func (fi *_FileInfo) findNext() error {
	return windows.FindNextFile(fi.handle, &fi.Win32finddata)
}

func (fi *_FileInfo) close() {
	windows.FindClose(fi.handle)
}

// Attribute returns attributes-bit of the file.
func (fi *_FileInfo) Attribute() uint32 {
	return fi.FileAttributes
}

// IsReparsePoint returns true when the file has a reparse point attribute.
func (fi *_FileInfo) IsReparsePoint() bool {
	return (fi.Attribute() & windows.FILE_ATTRIBUTE_REPARSE_POINT) != 0
}

// IsReadOnly returns true when the file has a readonly attribute.
func (fi *_FileInfo) IsReadOnly() bool {
	return (fi.Attribute() & windows.FILE_ATTRIBUTE_READONLY) != 0
}

// IsDir returns true when the file is directory.
func (fi *_FileInfo) IsDir() bool {
	return (fi.Attribute() & windows.FILE_ATTRIBUTE_DIRECTORY) != 0
}

// IsHidden returns true when the file is hidden.
func (fi *_FileInfo) IsHidden() bool {
	return (fi.Attribute() & windows.FILE_ATTRIBUTE_HIDDEN) != 0
}

// IsSystem returns true when the file has a system attribute.
func (fi *_FileInfo) IsSystem() bool {
	return (fi.Attribute() & windows.FILE_ATTRIBUTE_SYSTEM) != 0
}

// Walk enumerates the files matching patterns.
// It uses Win32's-findfile-API.
func walk(pattern string, callback func(*_FileInfo) bool) error {
	this, err := findFirst(pattern)
	if err != nil {
		return err
	}
	_pattern := strings.ToUpper(filepath.Base(pattern))
	defer this.close()
	for {
		_name := strings.ToUpper(this.Name())
		matched, err := filepath.Match(_pattern, _name)
		if err == nil && matched {
			if !callback(this.clone()) {
				return nil
			}
		}
		if err := this.findNext(); err != nil {
			return nil
		}
	}
}
