package nosql_test

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/maciejgaleja/gosimple/pkg/keyvalue"
	kvnosql "github.com/maciejgaleja/gosimple/pkg/keyvalue/nosql"
	nsqldynamo "github.com/maciejgaleja/gosimple/pkg/nosql/dynamodb"
	"github.com/maciejgaleja/gosimple/test"
	"github.com/stretchr/testify/assert"
)

func TestKeyValueDynamodb(t *testing.T) {
	test.DoTestKeyValue(t, func() keyvalue.Store {
		sess, err := session.NewSessionWithOptions(session.Options{
			SharedConfigState: session.SharedConfigEnable,
			Profile:           "default",
		})
		assert.NoError(t, err)

		nsql := nsqldynamo.NewDynamoDb(sess, "gosimple-test", "key")

		ddb := kvnosql.NewNoSql(nsql, "key", "value")
		err = ddb.Clear()
		assert.NoError(t, err)
		return ddb
	})
}
