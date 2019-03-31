package storage

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/dyoshikawa/aws-serverless-uploader-go/src/domains/defines"
)

func NewStorageLocalstack() Storage {
	svc := s3.New(session.New(), &aws.Config{
		Region:           aws.String(defines.Region),
		S3ForcePathStyle: aws.Bool(true),
		Endpoint:         aws.String(defines.LocalStackEndpointURLS3),
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
