package test

import (
	"testing"

	"github.com/maciejgaleja/gosimple/pkg/keyvalue"
	"github.com/stretchr/testify/assert"
)

func DoTestKeyValue(t *testing.T, k func() keyvalue.Store) {
	keyvalueTestExists(t, k())
	keyvalueTestSetGet(t, k())
	keyvalueTestList(t, k())
	keyvalueTestComplexType(t, k())
}

func keyvalueTestExists(t *testing.T, kv keyvalue.Store) {
	k := keyvalue.Key("test")
	v := keyvalue.Value("test")

	e := kv.Exists(k)
	assert.False(t, e)

	err := kv.Set(k, v)
	assert.NoError(t, err)

	e = kv.Exists(k)
	assert.True(t, e)

	err = kv.Remove(k)
	assert.NoError(t, err)

	e = kv.Exists(k)
	assert.False(t, e)
}

func keyvalueTestSetGet(t *testing.T, kv keyvalue.Store) {
	k := keyvalue.Key("test")
	v := keyvalue.Value("test")

	err := kv.Set(k, v)
	assert.NoError(t, err)

	var vr string
	err = kv.Get(k, &vr)
	assert.NoError(t, err)
	assert.Equal(t, v, vr)
}

func keyvalueTestList(t *testing.T, kv keyvalue.Store) {
	ks := []string{"test1", "test2", "test3", "test4"}
	for i, k := range ks {
		err := kv.Set(keyvalue.Key(k), keyvalue.Value(k))
		assert.NoError(t, err)

		l, err := kv.List()
		assert.NoError(t, err)
		assert.Equal(t, i+1, len(l))
	}
	for _, k := range ks {
		assert.True(t, kv.Exists(keyvalue.Key(k)))
	}
}

func keyvalueTestComplexType(t *testing.T, kv keyvalue.Store) {
	type complex struct {
		Name   string
		Number int
		Data   []byte
	}
	k := keyvalue.Key("test")
	v := keyvalue.Value(complex{Name: "a", Number: 1, Data: []byte("test")})

	err := kv.Set(k, v)
	assert.NoError(t, err)

	var vr complex
	err = kv.Get(k, &vr)
	assert.NoError(t, err)
	assert.Equal(t, v, vr)
}
