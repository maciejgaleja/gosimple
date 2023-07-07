package main

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/maciejgaleja/gosimple/pkg/keyvalue/dynamodb"
)

// Create struct to hold info about new item
type Item struct {
	Year   int
	Title  string
	Plot   string
	Rating float64
}

func main() {
	sess, err := session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Profile:           "default",
	})
	if err != nil {
		panic(err)
	}

	item := Item{
		Year:   2015,
		Title:  "423",
		Plot:   "Nothing happens at all.",
		Rating: 0.0,
	}

	ddb := dynamodb.NewDynamoDb(sess, "gosimple-test", "key")
	err = ddb.Set("1", item)
	if err != nil {
		panic(err)
	}
}
