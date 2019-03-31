package images

import (
	"encoding/base64"
	"strings"

	"github.com/dyoshikawa/aws-serverless-uploader-go/src/domains/defines"
	"github.com/dyoshikawa/aws-serverless-uploader-go/src/infrastructures/db/entities"

	"github.com/dyoshikawa/aws-serverless-uploader-go/src/infrastructures/db/repositories"
	"github.com/dyoshikawa/aws-serverless-uploader-go/src/infrastructures/storage"
	"github.com/google/uuid"
)

func fileElementsGet(file string) ([]byte, string, error) {
	strs := strings.Split(file, ",")

	var ext string
	if strings.Index(strs[0], "image/png") != -1 {
		ext = ".png"
	} else if strings.Index(strs[0], "image/jpg") != -1 || strings.Index(strs[0], "image/jpeg") != -1 {
		ext = ".jpg"
	}

	bytes, err := base64.StdEncoding.DecodeString(strs[1])
	if err != nil {
		return nil, "", err
	}

	return bytes, ext, err
}

func fileResize(file []byte, ext string) ([]byte, error) {
	var resized []byte
	var err error
	if ext == ".png" {
		resized, err = resizePng(file)
	} else if ext == ".jpg" {
		resized, err = resizeJpeg(file)
	}
	if err != nil {
		return nil, err
	}
	return resized, nil
}

func Put(storage storage.Storage, repo repositories.ImageRepository, b64Str string) error {
	file, ext, err := fileElementsGet(b64Str)
	if err != nil {
		return err
	}

	// 画像のリサイズ
	resized, err := fileResize(file, ext)
	if err != nil {
		return err
	}

	// ファイル名
	name := uuid.New().String() + ext
	// URL
	url := "https://s3-" + defines.Region + ".amazonaws.com/" + defines.S3BucketFiles + "/" + name

	// S3 アップロード
	if err := storage.Put(name, resized); err != nil {
		return err
	}

	// DynamoDB 書き込み
	image := entities.Image{
		Name: name,
		URL:  url,
	}
	if err := repo.Put(image); err != nil {
		return err
	}

	return nil
}
