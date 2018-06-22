package findfile

import (
	"os"
	"syscall"
)

// GetFileAttributes returns STATUS's FileAttributes member.
// ( status.Sys().Win32FileAttributeData or status.FileAttributes )
func GetFileAttributes(status os.FileInfo) uint32 {
	if it, ok := status.Sys().(*syscall.Win32FileAttributeData); ok && it != nil {
		return it.FileAttributes
	} else if it, ok := status.(*FileInfo); ok && it != nil {
		return it.FileAttributes
	} else {
		panic("Can not get fileatttribute")
	}
}
