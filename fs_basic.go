package b2fs

import (
	"io/fs"

	"github.com/go-git/go-billy/v5"
)

// Create implements billy.Filesystem.
func (*FS) Create(filename string) (billy.File, error) {
	panic("unimplemented")
}

// Join implements billy.Filesystem.
func (*FS) Join(elem ...string) string {
	return join(elem...)
}

// Open implements billy.Filesystem.
func (*FS) Open(filename string) (billy.File, error) {
	panic("unimplemented")
}

// OpenFile implements billy.Filesystem.
func (*FS) OpenFile(filename string, flag int, perm fs.FileMode) (billy.File, error) {
	panic("unimplemented")
}

// Remove implements billy.Filesystem.
func (*FS) Remove(filename string) error {
	panic("unimplemented")
}

// Rename implements billy.Filesystem.
func (*FS) Rename(oldpath string, newpath string) error {
	panic("unimplemented")
}

// Stat implements billy.Filesystem.
func (*FS) Stat(filename string) (fs.FileInfo, error) {
	panic("unimplemented")
}
