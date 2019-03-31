package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/dyoshikawa/aws-serverless-uploader-go/src/domains/images"
	"github.com/dyoshikawa/aws-serverless-uploader-go/src/domains/response"
	"github.com/dyoshikawa/aws-serverless-uploader-go/src/infrastructures/db/repositories"
)

const size = 300

type responseBody struct {
	Message string `json:"message"`
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	repo := repositories.NewImageRepository()
	images, err := images.GetAll(repo)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	return response.Response(200, images), nil
}

func main() {
	lambda.Start(handler)
}
