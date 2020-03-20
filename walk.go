package findfile

type FileInfo = _FileInfo

func Walk(pattern string, callback func(*_FileInfo) bool) error {
	return walk(pattern, callback)
}
