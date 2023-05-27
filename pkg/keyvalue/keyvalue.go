package keyvalue

type Key string
type Value any

type Store interface {
	Exists(Key) bool
	Set(Key, Value) error
	Get(Key) (Value, error)
	List() ([]Key, error)
	Remove(Key) error
	Clear() error
}
