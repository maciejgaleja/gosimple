package s3

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/maciejgaleja/gosimple/pkg/storage"
)

type S3BucketName string

type S3Store struct {
	bn string
	s3 *s3.Client
}

type ObjectWriteCache struct {
	buf bytes.Buffer
	s3  *S3Store
	key string
}

func (o *ObjectWriteCache) Write(d []byte) (int, error) {
	return o.buf.Write(d)
}

func (o *ObjectWriteCache) Close() error {
	input := &s3.PutObjectInput{
		Bucket: &o.s3.bn,
		Key:    &o.key,
		Body:   &o.buf,
	}

	_, err := o.s3.s3.PutObject(context.TODO(), input)
	return err
}

type ObjectReadCache struct {
	d io.ReadCloser
}

func (o *ObjectReadCache) Read(d []byte) (int, error) {
	n, err := o.d.Read(d)
	if err == io.EOF {
		err = nil
	}
	return n, err
}

func (o *ObjectReadCache) Close() error {
	return o.d.Close()
}

func NewS3Store(bn S3BucketName) (*S3Store, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, err
	}
	return &S3Store{s3: s3.NewFromConfig(cfg), bn: string(bn)}, nil
}

func (s *S3Store) Exists(h storage.Key) bool {
	input := &s3.HeadObjectInput{Bucket: &s.bn, Key: (*string)(&h)}

	_, err := s.s3.HeadObject(context.TODO(), input)
	return err == nil
}

func (s *S3Store) Create(h storage.Key) (io.WriteCloser, error) {
	if s.Exists(h) {
		return nil, fmt.Errorf("object already exists")
	}
	return &ObjectWriteCache{buf: bytes.Buffer{}, s3: s, key: string(h)}, nil
}

func (s *S3Store) Delete(h storage.Key) error {
	if !s.Exists(h) {
		return fmt.Errorf("object does not exist")
	}
	deleteInput := &s3.DeleteObjectInput{
		Bucket: &s.bn,
		Key:    (*string)(&h),
	}
	_, err := s.s3.DeleteObject(context.TODO(), deleteInput)
	return err
}

func (s *S3Store) Writer(h storage.Key) (io.WriteCloser, error) {
	if !s.Exists(h) {
		return nil, fmt.Errorf("object does not exist")
	}
	return &ObjectWriteCache{buf: bytes.Buffer{}, s3: s, key: string(h)}, nil
}

func (s *S3Store) Reader(h storage.Key) (io.ReadCloser, error) {
	if !s.Exists(h) {
		return nil, fmt.Errorf("object does not exist")
	}
	input := &s3.GetObjectInput{
		Bucket: &s.bn,
		Key:    (*string)(&h),
	}

	resp, err := s.s3.GetObject(context.TODO(), input)
	if err != nil {
		return nil, err
	}
	return &ObjectReadCache{d: resp.Body}, nil
}

func (s *S3Store) List() ([]storage.Key, error) {
	ret := []storage.Key{}
	input := &s3.ListObjectsV2Input{Bucket: &s.bn}

	listObjectsResponse, err := s.s3.ListObjectsV2(context.TODO(), input)
	if err != nil {
		return nil, err
	}
	for _, obj := range listObjectsResponse.Contents {
		ret = append(ret, storage.Key(*obj.Key))
	}
	return ret, nil
}
