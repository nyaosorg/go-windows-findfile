package findfile

import (
	"os"
	"syscall"
)

func GetFileAttributes(status os.FileInfo) uint32 {
	if it, ok := status.Sys().(*syscall.Win32FileAttributeData); ok && it != nil {
		return it.FileAttributes
	} else if it, ok := status.(*FileInfo); ok && it != nil {
		return it.FileAttributes
	} else {
		panic("Can not get fileatttribute")
	}
}
