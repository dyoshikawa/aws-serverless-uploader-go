package storage

import (
	"testing"

	"github.com/dyoshikawa/aws-serverless-uploader-go/src/domains/utils"
)

func TestPut(t *testing.T) {
	t.Run("S3アップロード PNGファイル", func(t *testing.T) {
		file, err := utils.FileGet("../../test-data/images/450-150.png")
		if err != nil {
			t.Fatal(err.Error())
		}

		storage := NewStorageLocalstack()
		if err := storage.Put("TEST", file); err != nil {
			t.Fatal(err.Error())
		}
	})
}
