package findfile

import (
	"context"
)

type FileInfo = _FileInfo

func Walk(pattern string, callback func(*_FileInfo) bool) error {
	return walk(nil, pattern, callback)
}

func WalkContext(ctx context.Context, pattern string, callback func(*_FileInfo) bool) error {
	return walk(ctx, pattern, callback)
}
