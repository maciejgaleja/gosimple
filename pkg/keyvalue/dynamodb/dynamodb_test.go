package dynamodb_test

import (
	"os"
	"testing"

	"github.com/maciejgaleja/gosimple/pkg/keyvalue"
	"github.com/maciejgaleja/gosimple/pkg/keyvalue/json"
	"github.com/maciejgaleja/gosimple/test"
)

func TestKeyValue(t *testing.T) {
	test.DoTestKeyValue(t, func() keyvalue.Store {
		os.MkdirAll(string(R), 0755)
		j, err := json.NewJsonStore(F)
		if err != nil {
			panic(err)
		}
		if err := j.Clear(); err != nil {
			panic(err)
		}
		return j
	})
}
