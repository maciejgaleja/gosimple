package s3_test

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
	simpleS3 "github.com/maciejgaleja/gosimple/pkg/storage/s3"
	"github.com/maciejgaleja/gosimple/test"
	"github.com/stretchr/testify/assert"
)

const (
	bucketName simpleS3.S3BucketName = "gosimple-test"
)

func cleanup(bn simpleS3.S3BucketName) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic(err)
	}

	client := s3.NewFromConfig(cfg)

	input := &s3.ListObjectsV2Input{Bucket: aws.String(string(bn))}

	listObjectsResponse, err := client.ListObjectsV2(context.TODO(), input)
	if err != nil {
		panic(err)
	}
	for _, obj := range listObjectsResponse.Contents {
		deleteInput := &s3.DeleteObjectInput{
			Bucket: aws.String(string(bn)),
			Key:    obj.Key,
		}

		_, err := client.DeleteObject(context.TODO(), deleteInput)
		if err != nil {
			panic(err)
		}
	}
}

func TestS3(t *testing.T) {
	s, err := simpleS3.NewS3Store(bucketName)
	assert.NoError(t, err)
	test.DoTestStorage(t, s, func() { cleanup(bucketName) })
	cleanup(bucketName)
}
