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
	Writer(Key) (io.WriteCloser, error)
	Reader(Key) (io.ReadSeekCloser, error)
	SetMetadata(Key, MetadataKey, MetadataValue) error
	GetMetadata(Key, MetadataKey) (MetadataValue, error)
	List() ([]Key, error)
}
