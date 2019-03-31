package repositories

import "github.com/dyoshikawa/aws-serverless-uploader-go/src/infrastructures/db/entities"

type ImageRepositoryMock struct{}

func NewImageRepositoryMock() ImageRepository {
	return &ImageRepositoryMock{}
}

func (repo *ImageRepositoryMock) GetAll() (entities.Images, error) {
	return entities.Images{}, nil
}

func (repo *ImageRepositoryMock) Put(image entities.Image) error {
	return nil
}

func (repo *ImageRepositoryMock) DeleteAll(image []entities.Image) error {
	return nil
}
