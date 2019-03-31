package defines

import "os"

const (
	Success = "Success."
)

var (
	AppName       = os.Getenv("APP_NAME")
	Stage         = os.Getenv("STAGE")
	Region        = os.Getenv("REGION")
	S3BucketFiles = os.Getenv("S3_BUCKET_FILES")

	// DynamoDB テーブル
	TableImage = AppName + "-" + Stage + "-" + "Images"

	// Localstack
	LocalStackEndpointURLS3       = os.Getenv("LOCALSTACK_ENDPOINT_URL_S3")
	LocalStackEndpointURLDynamoDB = os.Getenv("LOCALSTACK_ENDPOINT_URL_DYNAMODB")
)
