package repositories

import (
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/dyoshikawa/aws-serverless-uploader-go/src/domains/defines"
	"github.com/dyoshikawa/aws-serverless-uploader-go/src/infrastructures/db/entities"
	"github.com/guregu/dynamo"
)

type ImageRepository interface {
	GetAll() (entities.Images, error)
	Put(image entities.Image) error
	DeleteAll(image []entities.Image) error
}

type ImageRepositoryDynamo struct {
	table dynamo.Table
}

func NewImageRepository() ImageRepository {
	db := dynamo.New(session.New(), &aws.Config{Region: aws.String(defines.Region)})

	repo := ImageRepositoryDynamo{
		table: db.Table(defines.TableImage),
	}

	return &repo
}

func (repo *ImageRepositoryDynamo) GetAll() (entities.Images, error) {
	images := entities.Images{}
	err := repo.table.Scan().All(&images)
	if err != nil {
		return entities.Images{}, err
	}

	return images, nil
}

func (repo *ImageRepositoryDynamo) Put(image entities.Image) error {
	if image.CreatedAt == "" {
		image.CreatedAt = time.Now().Format("20060102150405")
	}
	if err := repo.table.Put(image).Run(); err != nil {
		return err
	}

	return nil
}

func (repo *ImageRepositoryDynamo) DeleteAll(image []entities.Image) error {
	for _, v := range image {
		repo.table.Delete("Name", v.Name).Run()
	}

	return nil
}
