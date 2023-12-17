package b2fs

import (
	"context"
	"path/filepath"
	"strings"

	"github.com/Backblaze/blazer/b2"
	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-billy/v5/helper/polyfill"
)

const sep = "/"

type FS struct {
	// billy.Filesystem doesn't accept context on calls, so we have to store it
	// here. For this reason, it's generally recommended to use run a single
	// filesystem per context.
	ctx context.Context

	b *b2.Bucket
}

func join(elem ...string) string {
	// copy from filepath.Join unix logic
	for i, e := range elem {
		if e != "" {
			return filepath.Clean(strings.Join(elem[i:], string(sep)))
		}
	}
	return ""
}

// Billy returns a billy.Filesystem for the given bucket. This automatically polyfills
// any missing interfaces.
func (f *FS) Billy() billy.Filesystem {
	return polyfill.New(f)
}

// type assertion
var (
	_ billy.Basic  = (*FS)(nil) // fs_basic.go
	_ billy.Dir    = (*FS)(nil) // fs_dir.go
	_ billy.Chroot = (*FS)(nil) // fs_chroot.go

	// Not supported by B2. These are polyfilled by [FS.Billy].
	// _ billy.Symlink  = (*B2FS)(nil)
	// _ billy.TempFile = (*B2FS)(nil)
)
