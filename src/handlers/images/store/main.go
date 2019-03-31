package main

import (
	"encoding/json"
	"errors"
	"log"
	"strings"

	"github.com/dyoshikawa/aws-serverless-uploader-go/src/infrastructures/db/repositories"
	validator "gopkg.in/go-playground/validator.v9"

	"github.com/dyoshikawa/aws-serverless-uploader-go/src/infrastructures/storage"

	"github.com/dyoshikawa/aws-serverless-uploader-go/src/domains/defines"
	"github.com/dyoshikawa/aws-serverless-uploader-go/src/domains/images"
	"github.com/dyoshikawa/aws-serverless-uploader-go/src/domains/response"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

const size = 300

type requestBody struct {
	Data string `json:"data" validate:"required"`
}

type responseBody struct {
	Message string `json:"message"`
}

func requestMarshal(body string) (*requestBody, error) {
	jsonBytes := ([]byte)(body)
	reqBody := new(requestBody)
	if err := json.Unmarshal(jsonBytes, reqBody); err != nil {
		return &requestBody{}, err
	}
	return reqBody, nil
}

func requestValidate(reqBody *requestBody) error {
	if err := validator.New().Struct(reqBody); err != nil {
		return err
	}

	strs := strings.Split(reqBody.Data, ",")
	if strings.Index(strs[0], "image/png") != -1 || strings.Index(strs[0], "image/jpg") != -1 || strings.Index(strs[0], "image/jpeg") != -1 {
	} else {
		return errors.New("ファイル形式が不正です。")
	}

	return nil
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	reqBody, err := requestMarshal(request.Body)
	if err != nil {
		log.Println(err.Error())
		return response.Response(500, response.ResponseBodyError{
			Errors: []string{"JSON形式が不正です。"},
		}), nil
	}
	if err := requestValidate(reqBody); err != nil {
		log.Println(err.Error())
		return response.Response(422, response.ResponseBodyError{
			Errors: []string{err.Error()},
		}), nil
	}

	b64Str := reqBody.Data
	storage := storage.NewStorage()
	repo := repositories.NewImageRepository()
	if err := images.Put(storage, repo, b64Str); err != nil {
		log.Println(err.Error())
		return response.Response(500, response.ResponseBodyError{
			Errors: []string{"データの保存に失敗しました。"},
		}), nil
	}

	return response.Response(200, responseBody{
		Message: defines.Success,
	}), nil
}

func main() {
	lambda.Start(handler)
}
