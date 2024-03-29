package storage

import (
	"io"
)

type Key string
type MetadataKey string
type MetadataValue []byte

type Storage interface {
	Exists(Key) bool
	Create(Key) (io.WriteCloser, error)
	Delete(Key) error
	Writer(Key) (io.WriteCloser, error)
	Reader(Key) (io.ReadCloser, error)
	List() ([]Key, error)
}
