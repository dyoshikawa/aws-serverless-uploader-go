package images

import (
	"testing"

	"github.com/dyoshikawa/aws-serverless-uploader-go/src/domains/utils"
)

func TestResizePng(t *testing.T) {
	t.Run("リサイズ PNG", func(t *testing.T) {
		file, err := utils.FileGet("../../test-data/images/450-150.png")
		if err != nil {
			t.Fatal(err.Error())
		}

		resized, err := resizePng(file)
		if err != nil {
			t.Fatal(err.Error())
		}

		decoded, err := utils.FileBase64Decode(resized)
		if err != nil {
			t.Fatal(err.Error())
		}

		if decoded.Width != 300 {
			t.Fatal("リサイズ後の横幅が300になっていません。")
		}
	})
}
