package filesystem

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/maciejgaleja/gosimple/pkg/storage"
	"github.com/maciejgaleja/gosimple/pkg/types"
	"github.com/pkg/xattr"
)

const (
	metadataKeyPrefix = "user.metadata."
)

var ErrIncompatibleFilesystem = fmt.Errorf("incompatible filesystem")

type FilesystemStore struct {
	Root types.DirectoryPath
}

func NewFilesystemStore(root types.DirectoryPath) (s FilesystemStore, err error) {
	s.Root = root
	c := s.IsCompatibleFilesystem()
	if !c {
		err = ErrIncompatibleFilesystem
	}
	return
}

func (s FilesystemStore) IsCompatibleFilesystem() bool {
	log.Printf("Checking filesystem for %v", s)

	k := "user.compatibilityCheck"
	v := []byte("test")

	pth := filepath.Join(string(s.Root), "compatibilityCheck")
	f, err := os.Create(pth)
	if err != nil {
		return false
	}

	_, err = f.Write([]byte("Compatibility check for store"))
	if err != nil {
		return false
	}

	err = xattr.FSet(f, "user.compatibilityCheck", v)
	if err != nil {
		return false
	}

	rv, err := xattr.FGet(f, k)
	if err != nil {
		return false
	}

	err = os.Remove(pth)
	if err != nil {
		return false
	}

	return string(rv) == string(v)
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
	return os.Create(pth)
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

func (s FilesystemStore) SetMetadata(h storage.Key, k storage.MetadataKey, v storage.MetadataValue) error {
	if !s.Exists(h) {
		return fmt.Errorf("object does not exist")
	}
	metadataKey := metadataKeyPrefix + k
	pth := filepath.Join(string(s.Root), string(h))
	return xattr.Set(pth, string(metadataKey), v)
}

func (s FilesystemStore) GetMetadata(h storage.Key, k storage.MetadataKey) (storage.MetadataValue, error) {
	if !s.Exists(h) {
		return []byte{}, fmt.Errorf("object does not exist")
	}

	metadataKey := metadataKeyPrefix + k

	pth := filepath.Join(string(s.Root), string(h))
	return xattr.Get(pth, string(metadataKey))
}

func (s FilesystemStore) List() ([]storage.Key, error) {
	fs, err := os.ReadDir(string(s.Root))
	if err != nil {
		return []storage.Key{}, err
	}

	ret := make([]storage.Key, 0)

	for _, e := range fs {
		ret = append(ret, storage.Key(e.Name()))
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
