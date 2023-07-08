package dynamodb_test

// func TestKeyValue(t *testing.T) {
// 	test.DoTestKeyValue(t, func() keyvalue.Store {
// 		sess, err := session.NewSessionWithOptions(session.Options{
// 			SharedConfigState: session.SharedConfigEnable,
// 			Profile:           "default",
// 		})
// 		assert.NoError(t, err)

// 		ddb := dynamodb.NewDynamoDb(sess, "gosimple-test", "key")
// 		err = ddb.Clear()
// 		assert.NoError(t, err)
// 		return ddb
// 	})
// }
