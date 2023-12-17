package b2fs

import (
	"context"
	"errors"
	"io/fs"
	"sync"
	"time"

	"github.com/Backblaze/blazer/b2"
	"github.com/go-git/go-billy/v5"
)

var (
	ErrConcurrentReadWrite = errors.New("concurrent read/write not supported")
	ErrSeekOpenFile        = errors.New("cannot seek on open file")
)

type File struct {
	b2.Object

	ctx context.Context

	mut sync.Mutex

	offset int64
	length int64

	reader *b2.Reader
	writer *b2.Writer
}

var (
	_ billy.File = (*File)(nil)
)

// getReader returns a reader for the file. If a reader is already open, it is
// returned. If a writer is open, this will return an error. If length is 0, the
// entire file will be read.
func (f *File) getReader(ctx context.Context, offset, length int64) (*b2.Reader, error) {
	f.mut.Lock()
	defer f.mut.Unlock()
	if length == 0 {
		length = -1
	}
	if f.reader == nil {
		r := f.NewRangeReader(ctx, offset, length)
		f.reader = r
		f.offset = offset
		f.length = length
	}
	if length != f.length || offset != f.offset {
		return nil, ErrSeekOpenFile
	}
	return f.reader, nil
}

// Close any open readers and writers. If you've opened a reader or writer, you
// must call Close() before it is garbage collected.
func (f *File) Close() error {
	f.mut.Lock()
	defer f.mut.Unlock()
	if f.reader != nil {
		err := f.reader.Close()
		f.reader = nil
		if err != nil {
			return err
		}
	}
	if f.writer != nil {
		err := f.writer.Close()
		f.writer = nil
		if err != nil {
			return err
		}
	}
	return nil
}

// Read implements billy.File.
func (f *File) Read(p []byte) (n int, err error) {
	reader, err := f.getReader(f.ctx, f.offset, f.length)
	if err != nil {
		return 0, err
	}
	return reader.Read(p)
}

// ReadAt implements billy.File.
func (*File) ReadAt(p []byte, off int64) (n int, err error) {
	panic("unimplemented")
}

// Seek implements billy.File.
func (f *File) Seek(offset int64, whence int) (int64, error) {
	panic("unimplemented")
}

// Truncate implements billy.File.
func (*File) Truncate(size int64) error {
	panic("unimplemented")
}

// Write implements billy.File.
func (*File) Write(p []byte) (n int, err error) {
	panic("unimplemented")
}

// Unsupported by B2

// Lock is not supported by B2.
func (*File) Lock() error {
	return billy.ErrNotSupported
}

// Unlock is not supported by B2.
func (*File) Unlock() error {
	return billy.ErrNotSupported
}

type fileInfo struct {
	ctx context.Context
	obj *b2.Object
}

// IsDir checks if the file is a directory. This is determined by checking if the
// file name ends with a slash.
func (i *fileInfo) IsDir() bool {
	name := i.obj.Name()
	return name[len(name)-1] == '/'
}

// ModTime implements fs.FileInfo.
func (i *fileInfo) ModTime() time.Time {
	panic("unimplemented")
}

// Mode implements fs.FileInfo.
func (i *fileInfo) Mode() fs.FileMode {
	mode := fs.ModePerm
	if i.IsDir() {
		mode |= fs.ModeDir
	}
	return mode
}

// Name implements fs.FileInfo.
func (i *fileInfo) Name() string {
	return i.obj.Name()
}

// Size implements fs.FileInfo.
func (i *fileInfo) Size() int64 {
	panic("unimplemented")
}

// Sys implements fs.FileInfo.
func (i *fileInfo) Sys() any {
	return i.obj
}

var (
	_ fs.FileInfo = (*fileInfo)(nil)
)
