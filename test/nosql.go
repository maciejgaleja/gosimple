package test

import (
	"testing"

	"github.com/maciejgaleja/gosimple/pkg/nosql"
	"github.com/stretchr/testify/assert"
)

func DoTestNoSql(t *testing.T, k func() nosql.Store) {
	t.Run("exists", func(t *testing.T) {
		nosqlTestExists(t, k())
	})
	t.Run("setGet", func(t *testing.T) {
		nosqlTestSetGet(t, k())
	})
	t.Run("list", func(t *testing.T) {
		nosqlTestList(t, k())
	})
	t.Run("clear", func(t *testing.T) {
		nosqlTestClear(t, k())
	})
}

func nosqlTestExists(t *testing.T, ns nosql.Store) {
	k := nosql.PrimaryKey("test")

	e, err := ns.Exists(k)
	assert.NoError(t, err)
	assert.False(t, e)

	err = ns.Set(nosql.Document{"key": k})
	assert.NoError(t, err)

	e, err = ns.Exists(k)
	assert.NoError(t, err)
	assert.True(t, e)

	err = ns.Remove(k)
	assert.NoError(t, err)

	e, err = ns.Exists(k)
	assert.NoError(t, err)
	assert.False(t, e)
}

func nosqlTestSetGet(t *testing.T, ns nosql.Store) {
	k := nosql.PrimaryKey("test")
	v := nosql.Document{"key": string(k)}

	err := ns.Set(v)
	assert.NoError(t, err)

	var vr nosql.Document
	err = ns.Get(k, &vr)
	assert.NoError(t, err)
	assert.Equal(t, v, vr)
}

func nosqlTestList(t *testing.T, ns nosql.Store) {
	ks := []string{"test1", "test2", "test3", "test4"}
	for i, k := range ks {
		err := ns.Set(nosql.Document{"key": k, "value": k})
		assert.NoError(t, err)

		l, err := ns.List()
		assert.NoError(t, err)
		assert.Equal(t, i+1, len(l))
	}
	for _, k := range ks {
		b, err := ns.Exists(nosql.PrimaryKey(k))
		assert.True(t, b)
		assert.NoError(t, err)
	}
}

func nosqlTestClear(t *testing.T, ns nosql.Store) {
	err := ns.Clear()
	assert.NoError(t, err)

	ii, err := ns.List()
	assert.NoError(t, err)
	assert.Equal(t, 0, len(ii))
}
