package findfile

import (
	"os"
	"syscall"
	"time"
)

// FileInfo is the struct for Walk()'s parameter
type FileInfo struct {
	syscall.Win32finddata
	handle syscall.Handle
}

func (fi *FileInfo) clone() *FileInfo {
	return &FileInfo{fi.Win32finddata, fi.handle}
}

func findFirst(pattern string) (*FileInfo, error) {
	pattern16, err := syscall.UTF16PtrFromString(pattern)
	if err != nil {
		return nil, err
	}
	this := new(FileInfo)
	this.handle, err = syscall.FindFirstFile(pattern16, &this.Win32finddata)
	if err != nil {
		return nil, err
	}
	return this, nil
}

// Name returns file's name
func (fi *FileInfo) Name() string {
	return syscall.UTF16ToString(fi.FileName[:])
}

// Size returns file's size by bytes
func (fi *FileInfo) Size() int64 {
	return int64((int64(fi.FileSizeHigh) << 32) | int64(fi.FileSizeLow))
}

// ModTime returns timestamp when the file was last updated.
func (fi *FileInfo) ModTime() time.Time {
	return time.Unix(0, fi.LastWriteTime.Nanoseconds())
}

// Mode emurates os.FileMode.
func (fi *FileInfo) Mode() os.FileMode {
	m := os.FileMode(0444)
	if fi.IsDir() {
		m |= 0111 | os.ModeDir
	}
	if !fi.IsReadOnly() {
		m |= 0222
	}
	return m
}

// Sys returns underlying data source like os.FileInfo.Sys()
func (fi *FileInfo) Sys() interface{} {
	return &fi.Win32finddata
}

func (fi *FileInfo) findNext() error {
	return syscall.FindNextFile(fi.handle, &fi.Win32finddata)
}

func (fi *FileInfo) close() {
	syscall.FindClose(fi.handle)
}

// Attribute returns attributes-bit of the file.
func (fi *FileInfo) Attribute() uint32 {
	return fi.FileAttributes
}

// IsReparsePoint returns true when the file has a reparse point attribute.
func (fi *FileInfo) IsReparsePoint() bool {
	return (fi.Attribute() & FILE_ATTRIBUTE_REPARSE_POINT) != 0
}

// IsReadOnly returns true when the file has a readonly attribute.
func (fi *FileInfo) IsReadOnly() bool {
	return (fi.Attribute() & syscall.FILE_ATTRIBUTE_READONLY) != 0
}

// IsDir returns true when the file is directory.
func (fi *FileInfo) IsDir() bool {
	return (fi.Attribute() & syscall.FILE_ATTRIBUTE_DIRECTORY) != 0
}

// IsHidden returns true when the file is hidden.
func (fi *FileInfo) IsHidden() bool {
	return (fi.Attribute() & syscall.FILE_ATTRIBUTE_HIDDEN) != 0
}

// IsSystem returns true when the file has a system attribute.
func (fi *FileInfo) IsSystem() bool {
	return (fi.Attribute() & syscall.FILE_ATTRIBUTE_SYSTEM) != 0
}

// Walk enumerates the files matching patterns.
// It uses Win32's-findfile-API.
func Walk(pattern string, callback func(*FileInfo) bool) error {
	this, err := findFirst(pattern)
	if err != nil {
		return err
	}
	defer this.close()
	for {
		if !callback(this.clone()) {
			return nil
		}
		if err := this.findNext(); err != nil {
			return nil
		}
	}
}
