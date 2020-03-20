package findfile

import (
	"os"
)

func GetFileAttributes(status os.FileInfo) uint32 {
	return getFileAttributes(status)
}
