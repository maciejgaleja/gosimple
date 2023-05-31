package json

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/maciejgaleja/gosimple/pkg/keyvalue"
	"github.com/maciejgaleja/gosimple/pkg/types"
)

func ErrNoSuchKey(k keyvalue.Key) error {
	return fmt.Errorf("no such key: %s", k)
}

type JsonStore struct {
	f types.FilePath
	d map[keyvalue.Key]keyvalue.Value
}

func NewJsonStore(f types.FilePath) (*JsonStore, error) {
	j := JsonStore{f: f}

	if _, err := os.Stat(string(f)); errors.Is(err, os.ErrNotExist) {
		j.d = map[keyvalue.Key]keyvalue.Value{}
	} else {
		d, err := os.ReadFile(string(j.f))
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(d, &j.d)
		if err != nil {
			return nil, err
		}
	}

	return &j, nil
}

func (j *JsonStore) sync() error {
	d, err := json.Marshal(j.d)
	if err != nil {
		return err
	}

	return os.WriteFile(string(j.f), d, 0644)
}

func (j *JsonStore) Exists(k keyvalue.Key) bool {
	_, ok := j.d[k]
	return ok
}

func (j *JsonStore) Set(k keyvalue.Key, v keyvalue.Value) error {
	j.d[k] = v
	return j.sync()
}

func (j *JsonStore) Get(k keyvalue.Key, v any) error {
	vv, ok := j.d[k]
	if ok {
		b, err := json.Marshal(vv)
		if err != nil {
			return err
		}
		err = json.Unmarshal(b, v)
		if err != nil {
			return err
		}
		return nil
	} else {
		return ErrNoSuchKey(k)
	}
}

func (j *JsonStore) List() ([]keyvalue.Key, error) {
	ret := make([]keyvalue.Key, len(j.d))
	i := 0
	for k := range j.d {
		ret[i] = k
		i++
	}
	return ret, nil
}

func (j *JsonStore) Remove(k keyvalue.Key) error {
	delete(j.d, k)
	return j.sync()
}

func (j *JsonStore) Clear() error {
	j.d = map[keyvalue.Key]keyvalue.Value{}
	return j.sync()
}
