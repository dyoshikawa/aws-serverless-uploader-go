package images

import (
	"testing"

	"github.com/dyoshikawa/aws-serverless-uploader-go/src/infrastructures/db/repositories"

	"github.com/dyoshikawa/aws-serverless-uploader-go/src/domains/utils"

	"github.com/dyoshikawa/aws-serverless-uploader-go/src/infrastructures/storage"
)

func TestPut(t *testing.T) {
	t.Run("データ保存", func(t *testing.T) {
		file, err := utils.FileBase64Get("../../test-data/images/450-150.png")
		if err != nil {
			t.Fatal(err.Error())
		}

		storage := storage.NewStorageMock()
		repo := repositories.NewImageRepositoryMock()
		if err := Put(storage, repo, file); err != nil {
			t.Fatal(err.Error())
		}
	})
}
