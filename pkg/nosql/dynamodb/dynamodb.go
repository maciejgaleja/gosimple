package dynamodb

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	awsdynamo "github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/maciejgaleja/gosimple/pkg/nosql"
)

type DynamoDb struct {
	d *awsdynamo.DynamoDB
	n string
	k nosql.PrimaryKey
}

func NewDynamoDb(sess *session.Session, tableName string, key nosql.PrimaryKey) DynamoDb {
	return DynamoDb{d: awsdynamo.New(sess), n: tableName, k: key}
}

func (d DynamoDb) Exists(k nosql.PrimaryKey) (bool, error) {
	ak, err := dynamodbattribute.Marshal(k)
	if err != nil {
		return false, err
	}
	input := &awsdynamo.GetItemInput{
		TableName: aws.String(d.n),
		Key:       map[string]*awsdynamo.AttributeValue{string(d.k): ak},
	}
	result, err := d.d.GetItem(input)
	if err != nil {
		return false, err
	}
	return result.Item != nil, nil
}

func (d DynamoDb) Set(doc nosql.Document) error {
	av, err := dynamodbattribute.MarshalMap(doc)
	if err != nil {
		return err
	}

	input := &awsdynamo.PutItemInput{
		Item:      av,
		TableName: aws.String(d.n),
	}

	_, err = d.d.PutItem(input)
	return err
}

func (d DynamoDb) Get(k nosql.PrimaryKey, doc *nosql.Document) error {
	ak, err := dynamodbattribute.Marshal(k)
	if err != nil {
		return err
	}
	input := &awsdynamo.GetItemInput{
		TableName: aws.String(d.n),
		Key:       map[string]*awsdynamo.AttributeValue{string(d.k): ak},
	}
	result, err := d.d.GetItem(input)
	if err != nil {
		return err
	}
	if result.Item == nil {
		return fmt.Errorf("object with key '%s' does not exist in table '%s'", k, d.n)
	}
	return dynamodbattribute.UnmarshalMap(result.Item, doc)
}

func (d DynamoDb) List() ([]nosql.PrimaryKey, error) {
	input := &awsdynamo.ScanInput{
		TableName: aws.String(d.n),
	}

	result, err := d.d.Scan(input)
	if err != nil {
		return nil, err
	}

	ret := []nosql.PrimaryKey{}
	for _, item := range result.Items {
		doc := map[string]interface{}{}
		err = dynamodbattribute.UnmarshalMap(item, &doc)
		if err != nil {
			return nil, err
		}
		ks, ok := doc[string(d.k)].(string)
		if !ok {
			return nil, fmt.Errorf("one of objects in table '%s' does not contain the key named '%s'", d.n, d.k)
		}
		ret = append(ret, nosql.PrimaryKey(ks))
	}
	return ret, nil
}

func (d DynamoDb) Remove(k nosql.PrimaryKey) error {
	input := &awsdynamo.DeleteItemInput{
		TableName: aws.String(d.n),
		Key: map[string]*awsdynamo.AttributeValue{
			string(d.k): {
				S: aws.String(string(k)),
			},
		},
	}

	_, err := d.d.DeleteItem(input)
	return err
}

func (d DynamoDb) Clear() error {
	ks, err := d.List()
	if err != nil {
		return err
	}
	for _, k := range ks {
		err = d.Remove(k)
		if err != nil {
			return err
		}
	}
	return nil
}
