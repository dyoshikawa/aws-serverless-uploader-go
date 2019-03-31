package response

import (
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
)

type ResponseBodyError struct {
	Errors []string `json:"errors"`
}

func Response(code int, body interface{}) events.APIGatewayProxyResponse {
	jsonBytes, err := json.Marshal(body)
	if err != nil {
		log.Fatal(err.Error())
	}

	return events.APIGatewayProxyResponse{
		Headers: map[string]string{
			"Access-Control-Allow-Origin":  "*",
			"Access-Control-Allow-Headers": "origin,Accept,Authorization,Content-Type",
			"Content-Type":                 "application/json",
		},
		Body:       string(jsonBytes),
		StatusCode: code,
	}
}
