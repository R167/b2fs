package b2fs

import (
	"io/fs"

	"github.com/Backblaze/blazer/b2"
)

// MkdirAll implements billy.Filesystem.
func (f *FS) MkdirAll(filename string, perm fs.FileMode) error {
	// TODO: figure out if theres a better thing to do here. b/c B2 is an object
	// store, there's no concept of directories. We could create a 0 byte file
	// to force the "directory" to exist, but that's not really a good solution and
	// requires special filtering on the client side.
	return nil
}

// ReadDir implements billy.Filesystem.
func (f *FS) ReadDir(path string) ([]fs.FileInfo, error) {
	opts := []b2.ListOption{
		b2.ListDelimiter("/"),
		b2.ListPrefix(dirPath(path)),
	}
	iter := f.b.List(f.ctx, opts...)
	var out []fs.FileInfo
	for iter.Next() {
		obj := iter.Object()
		out = append(out, &fileInfo{
			ctx: f.ctx,
			obj: obj,
		})
	}
	if err := iter.Err(); err != nil {
		return nil, err
	}
	return out, nil
}

// dirPath normalizes the path into a normalized directory path.
//
// TODO: reject "non-sense" paths like `../` and `./`?
func dirPath(path string) string {
	if path[0] == '/' {
		path = path[1:]
	}
	if path == "" {
		return ""
	}
	if path[len(path)-1] != '/' {
		return path + "/"
	}
	return path
}
