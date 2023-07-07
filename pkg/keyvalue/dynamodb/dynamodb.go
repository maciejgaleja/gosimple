package dynamodb

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	awsdynamo "github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/maciejgaleja/gosimple/pkg/keyvalue"
)

type DynamoDb struct {
	d *awsdynamo.DynamoDB
	n string
	k string
}

func NewDynamoDb(sess *session.Session, tableName string, key string) DynamoDb {
	return DynamoDb{d: awsdynamo.New(sess), n: tableName, k: key}
}

// func (d DynamoDb) Exists(keyvalue.Key) bool {

// }

func (d DynamoDb) Set(k keyvalue.Key, v keyvalue.Value) error {
	av, err := dynamodbattribute.MarshalMap(v)
	if err != nil {
		return err
	}
	av[d.k], err = dynamodbattribute.Marshal(k)
	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(d.n),
	}

	_, err = d.d.PutItem(input)
	return err
}

// func (d DynamoDb) Get(keyvalue.Key, any) error {

// }

// func (d DynamoDb) List() ([]keyvalue.Key, error) {

// }

// func (d DynamoDb) Remove(keyvalue.Key) error {

// }

// func (d DynamoDb) Clear() error {

// }
