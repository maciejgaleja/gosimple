package keyvalue

import "encoding/json"

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

func Cast[T any](v Value) (T, error) {
	var t T
	b, err := json.Marshal(v)
	if err != nil {
		return t, err
	}
	err = json.Unmarshal(b, &t)
	if err != nil {
		return t, err
	}
	return t, nil
}
