package keyvalue

type Key string
type Value any

type Store interface {
	Exists(Key) (bool, error)
	Set(Key, Value) error
	Get(Key, any) error
	List() ([]Key, error)
	Remove(Key) error
	Clear() error
}
