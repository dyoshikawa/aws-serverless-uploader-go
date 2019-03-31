package repositories

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/dyoshikawa/aws-serverless-uploader-go/src/domains/defines"
	"github.com/guregu/dynamo"
)

type ImageRepositoryLocalstack struct {
	ImageRepositoryDynamo
}

func NewImageRepositoryLocalstack() ImageRepository {
	db := dynamo.New(session.New(), &aws.Config{
		Region:   aws.String(defines.Region),
		Endpoint: aws.String(defines.LocalStackEndpointURLDynamoDB),
	})

	repo := ImageRepositoryDynamo{
		table: db.Table(defines.TableImage),
	}

	return &repo
}
