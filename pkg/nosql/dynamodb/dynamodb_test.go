package dynamodb_test

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/maciejgaleja/gosimple/pkg/nosql"
	"github.com/maciejgaleja/gosimple/pkg/nosql/dynamodb"
	"github.com/maciejgaleja/gosimple/test"
	"github.com/stretchr/testify/assert"
)

func TestDynamoDb(t *testing.T) {
	test.DoTestNoSql(t, func() nosql.Store {
		sess, err := session.NewSessionWithOptions(session.Options{
			SharedConfigState: session.SharedConfigEnable,
			Profile:           "default",
		})
		assert.NoError(t, err)

		ddb := dynamodb.NewDynamoDb(sess, "gosimple-test", "key")
		err = ddb.Clear()
		assert.NoError(t, err)
		return ddb
	})
}
