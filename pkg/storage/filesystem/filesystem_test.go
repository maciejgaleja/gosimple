package filesystem_test

import (
	"os"
	"testing"

	"github.com/maciejgaleja/gosimple/pkg/storage"
	"github.com/maciejgaleja/gosimple/pkg/storage/filesystem"
	"github.com/maciejgaleja/gosimple/pkg/types"
	"github.com/maciejgaleja/gosimple/test"
	"github.com/stretchr/testify/assert"
)

const (
	rootPath = types.DirectoryPath("../../../test/workspace/storage")
)

func cleanup(rootPath types.DirectoryPath) {
	if err := os.RemoveAll(string(rootPath)); err != nil {
		panic(err)
	}
	if err := os.MkdirAll(string(rootPath), 0755); err != nil {
		panic(err)
	}
}

func TestFilesystem(t *testing.T) {
	fs, err := filesystem.NewFilesystemStore(rootPath)
	assert.NoError(t, err)
	test.DoTestStorage(t, fs, func() { cleanup(rootPath) })
}

func TestReadOnly(t *testing.T) {
	s, err := filesystem.NewFilesystemStore(rootPath)
	assert.NoError(t, err)

	cleanup(rootPath)

	k := storage.Key("test")
	d := []byte("test")

	err = s.MakeReadOnly(k)
	assert.Error(t, err)

	w, err := s.Create(k)
	assert.NoError(t, err)

	n, err := w.Write(d)
	assert.NoError(t, err)
	assert.Equal(t, len(d), n)

	err = w.Close()
	assert.NoError(t, err)

	r, err := s.Reader(k)
	assert.NoError(t, err)

	rd := make([]byte, len(d))
	n, err = r.Read(rd)
	assert.NoError(t, err)
	assert.Equal(t, len(d), n)
	assert.EqualValues(t, d, rd)

	err = r.Close()
	assert.NoError(t, err)

	err = s.MakeReadOnly(k)
	assert.NoError(t, err)

	_, err = w.Write(d)
	assert.Error(t, err)
}
