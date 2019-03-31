package images

import (
	"sort"

	"github.com/dyoshikawa/aws-serverless-uploader-go/src/infrastructures/db/entities"

	"github.com/dyoshikawa/aws-serverless-uploader-go/src/infrastructures/db/repositories"
)

func GetAll(repo repositories.ImageRepository) ([]entities.Image, error) {
	images, err := repo.GetAll()
	if err != nil {
		return []entities.Image{}, err
	}

	sort.Sort(images)

	return images, nil
}
