package test

import (
	"testing"

	"github.com/maciejgaleja/gosimple/pkg/storage"
	"github.com/stretchr/testify/assert"
)

func DoTest(t *testing.T, s storage.Storage) {
	e := s.Exists("aaa")
	assert.Equal(t, false, e)
}
