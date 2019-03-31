package storage

import (
	"bytes"

	"github.com/dyoshikawa/aws-serverless-uploader-go/src/domains/defines"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type Storage interface {
	Put(key string, data []byte) error
	Destroy() error
}

type StorageS3 struct {
	svc            *s3.S3
	iter           s3manager.BatchDeleteIterator
	putObjectInput *s3.PutObjectInput
}

func NewStorage() Storage {
	svc := s3.New(session.New(), &aws.Config{
		Region: aws.String(defines.Region),
	})
	input := &s3.PutObjectInput{
		Bucket:               aws.String(defines.S3BucketFiles),
		ACL:                  aws.String("public-read"),
		ServerSideEncryption: aws.String("AES256"),
	}
	iter := s3manager.NewDeleteListIterator(svc, &s3.ListObjectsInput{
		Bucket: aws.String(defines.S3BucketFiles),
	})
	return &StorageS3{
		svc:            svc,
		iter:           iter,
		putObjectInput: input,
	}
}

func (storage *StorageS3) Put(key string, file []byte) error {
	input := storage.putObjectInput
	input.Key = aws.String(key)
	input.Body = bytes.NewReader(file)
	_, err := storage.svc.PutObject(input)
	if err != nil {
		return err
	}
	return nil
}

func (storage *StorageS3) Destroy() error {
	if err := s3manager.NewBatchDeleteWithClient(storage.svc).Delete(aws.BackgroundContext(), storage.iter); err != nil {
		return err
	}
	return nil
}
