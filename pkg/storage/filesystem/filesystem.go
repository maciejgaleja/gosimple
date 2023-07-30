package filesystem

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/maciejgaleja/gosimple/pkg/storage"
	"github.com/maciejgaleja/gosimple/pkg/types"
)

type FilesystemStore struct {
	Root types.DirectoryPath
}

func NewFilesystemStore(root types.DirectoryPath) (s FilesystemStore, err error) {
	s.Root = root
	return
}

func (s FilesystemStore) Exists(h storage.Key) bool {
	pth := filepath.Join(string(s.Root), string(h))
	_, err := os.Stat(pth)
	return !errors.Is(err, os.ErrNotExist)
}

func (s FilesystemStore) Create(h storage.Key) (io.WriteCloser, error) {
	if s.Exists(h) {
		return nil, fmt.Errorf("object already exists")
	}

	pth := filepath.Join(string(s.Root), string(h))

	d := filepath.Dir(pth)
	if _, err := os.Stat(d); err != nil {
		if err := os.MkdirAll(d, os.ModePerm); err != nil {
			return nil, err
		}
	}

	return os.Create(pth)
}

func (s FilesystemStore) Delete(h storage.Key) error {
	if !s.Exists(h) {
		return fmt.Errorf("object does not exist")
	}
	pth := filepath.Join(string(s.Root), string(h))
	return os.Remove(pth)
}

func (s FilesystemStore) Writer(h storage.Key) (io.WriteCloser, error) {
	if !s.Exists(h) {
		return nil, fmt.Errorf("object does not exist")
	}

	pth := filepath.Join(string(s.Root), string(h))
	return os.OpenFile(pth, os.O_WRONLY|os.O_TRUNC, 0644)
}

func (s FilesystemStore) Reader(h storage.Key) (io.ReadSeekCloser, error) {
	if !s.Exists(h) {
		return nil, fmt.Errorf("object does not exist")
	}

	pth := filepath.Join(string(s.Root), string(h))
	return os.Open(pth)
}

func (s FilesystemStore) List() ([]storage.Key, error) {
	ret := []storage.Key{}

	err := filepath.Walk(string(s.Root), func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		k, err := filepath.Rel(string(s.Root), path)
		if err != nil {
			return err
		}
		ret = append(ret, storage.Key(k))
		return nil
	})
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (s FilesystemStore) MakeReadOnly(h storage.Key) error {
	if !s.Exists(h) {
		return fmt.Errorf("object does not exist")
	}
	pth := filepath.Join(string(s.Root), string(h))
	return os.Chmod(pth, 0400)
}
