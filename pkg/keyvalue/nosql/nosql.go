package nosql

import (
	"encoding/json"
	"fmt"

	"github.com/maciejgaleja/gosimple/pkg/keyvalue"
	"github.com/maciejgaleja/gosimple/pkg/nosql"
)

type NoSql struct {
	db nosql.Store
	k  nosql.PrimaryKey
	v  string
}

func NewNoSql(nsql nosql.Store, keyName nosql.PrimaryKey, valueName string) NoSql {
	return NoSql{db: nsql, k: keyName, v: valueName}
}

func (d NoSql) Exists(k keyvalue.Key) (bool, error) {
	return d.db.Exists(nosql.PrimaryKey(k))
}

func (d NoSql) Set(k keyvalue.Key, v keyvalue.Value) error {
	bs, err := json.Marshal(v)
	if err != nil {
		return err
	}
	doc := nosql.Document{}
	doc[string(d.k)] = k
	doc[d.v] = bs
	return d.db.Set(doc)
}

func (d NoSql) Get(k keyvalue.Key, v any) error {
	var doc nosql.Document
	err := d.db.Get(nosql.PrimaryKey(k), &doc)
	if err != nil {
		return err
	}
	bs, ok := doc[d.v].([]byte)
	if !ok {
		return fmt.Errorf("error while parsing value of key '%s'", k)
	}
	return json.Unmarshal(bs, v)
}

func (d NoSql) List() ([]keyvalue.Key, error) {
	ks, err := d.db.List()
	if err != nil {
		return nil, err
	}
	ret := make([]keyvalue.Key, len(ks))
	for i, k := range ks {
		ret[i] = keyvalue.Key(k)
	}
	return ret, nil
}

func (d NoSql) Remove(k keyvalue.Key) error {
	return d.db.Remove(nosql.PrimaryKey(k))
}

func (d NoSql) Clear() error {
	return d.db.Clear()
}
