package images

import (
	"testing"

	"github.com/dyoshikawa/aws-serverless-uploader-go/src/infrastructures/db/repositories"
)

func TestGetAll(t *testing.T) {
	t.Run("データ取得", func(t *testing.T) {
		repo := repositories.NewImageRepositoryMock()
		_, err := GetAll(repo)
		if err != nil {
			t.Fatal(err.Error())
		}
	})
}
