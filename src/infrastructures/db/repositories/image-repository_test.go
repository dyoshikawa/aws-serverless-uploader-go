package repositories

import (
	"testing"
	"time"

	"github.com/dyoshikawa/aws-serverless-uploader-go/src/domains/defines"

	"github.com/dyoshikawa/aws-serverless-uploader-go/src/infrastructures/db/entities"
	"github.com/google/uuid"
)

func TestPut(t *testing.T) {
	t.Run("DynamoDB 書き込み", func(t *testing.T) {
		name := uuid.New().String()
		url := defines.LocalStackEndpointURLS3 + "/" + defines.S3BucketFiles + "/" + name
		repo := NewImageRepositoryLocalstack()
		image := entities.Image{
			Name:      name,
			URL:       url,
			CreatedAt: time.Now().Format("20060102150405"),
		}
		if err := repo.Put(image); err != nil {
			t.Fatal(err.Error())
		}
	})

	t.Run("DynamoDB レコード取得", func(t *testing.T) {
		name := uuid.New().String()
		url := defines.LocalStackEndpointURLS3 + "/" + defines.S3BucketFiles + "/" + name
		repo := NewImageRepositoryLocalstack()
		image := entities.Image{
			Name:      name,
			URL:       url,
			CreatedAt: time.Now().Format("20060102150405"),
		}
		if err := repo.Put(image); err != nil {
			t.Fatal(err.Error())
		}

		images, err := repo.GetAll()
		if err != nil {
			t.Fatal(err.Error())
		}
		if images.Len() == 0 {
			t.Fatal("レコードを取得できません。")
		}
	})
}
