package findfile

import (
	"os"
)

// GetFileAttributes returns STATUS's FileAttributes member.
// dummy
func getFileAttributes(status os.FileInfo) uint32 {
	return 0
}
