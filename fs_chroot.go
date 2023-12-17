package b2fs

import (
	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-billy/v5/helper/chroot"
)

// Root implements billy.Filesystem.
func (*FS) Root() string {
	return ""
}

// Chroot returns a new filesystem rooted at the given path.
func (f *FS) Chroot(path string) (billy.Filesystem, error) {
	return chroot.New(f, path), nil
}
