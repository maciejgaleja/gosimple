package test

import (
	"testing"

	"github.com/maciejgaleja/gosimple/pkg/storage"
	"github.com/stretchr/testify/assert"
)

func DoTestStorage(t *testing.T, s storage.Storage, cleanup func()) {
	cleanup()
	storageTestCreateDelete(t, s)

	cleanup()
	storageTestWriterReader(t, s)

	cleanup()
	storageTestList(t, s)
}

func storageTestCreateDelete(t *testing.T, s storage.Storage) {
	k := storage.Key("test")

	e := s.Exists(k)
	assert.False(t, e)

	w, err := s.Create(k)
	assert.NoError(t, err)

	err = w.Close()
	assert.NoError(t, err)

	_, err = s.Create(k)
	assert.Error(t, err)

	err = s.Delete(k)
	assert.NoError(t, err)

	e = s.Exists(k)
	assert.False(t, e)
}

func storageTestWriterReader(t *testing.T, s storage.Storage) {
	k := storage.Key("test")
	d := []byte("test")

	_, err := s.Writer(k)
	assert.Error(t, err)

	_, err = s.Reader(k)
	assert.Error(t, err)

	w, err := s.Create(k)
	assert.NoError(t, err)

	err = w.Close()
	assert.NoError(t, err)

	w, err = s.Writer(k)
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
}

func storageTestList(t *testing.T, s storage.Storage) {
	k := storage.Key("test")

	l, err := s.List()
	assert.NoError(t, err)
	assert.Equal(t, 0, len(l))

	w, err := s.Create(k)
	assert.NoError(t, err)

	err = w.Close()
	assert.NoError(t, err)

	l, err = s.List()
	assert.NoError(t, err)
	assert.Equal(t, 1, len(l))
}
