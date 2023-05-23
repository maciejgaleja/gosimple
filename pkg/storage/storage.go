package storage

import (
	"io"
)

type Key string
type MetadataKey string
type MetadataValue string

type Storage interface {
	Exists(Key) (bool, error)
	Create(Key) error
	Writer(Key) io.WriteCloser
	Reader(Key) io.ReadSeekCloser
	SetMetadata(MetadataKey, MetadataValue) error
	GetMetadata(MetadataKey) (MetadataValue, error)
	List() ([]Key, error)
}
