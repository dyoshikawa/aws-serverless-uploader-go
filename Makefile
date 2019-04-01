.PHONY: build clean deploy

dep:
	go mod vendor

build:
	go mod tidy
	env GOOS=linux go build -ldflags="-s -w" -o bin/images-get src/handlers/images/index/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/images-put src/handlers/images/store/main.go

test:
	go test -v ./src/domains/images
	go test -v ./src/infrastructures/storage
	go test -v ./src/infrastructures/db/repositories

clean:
	rm -rf ./bin ./vendor

deploy: build
	yarn run sls deploy --stage dev

remove:
	yarn run sls remove --stage dev

prod-deploy: clean build
	yarn run sls deploy --stage prod

localstack-setup: localstack-setup-s3 localstack-setup-dynamodb

localstack-setup-s3:
	aws --profile localstack --endpoint-url=${LOCALSTACK_ENDPOINT_URL_S3} s3 mb s3://${S3_BUCKET_FILES}
	aws --profile localstack --endpoint-url=${LOCALSTACK_ENDPOINT_URL_S3} s3api put-bucket-acl --bucket ${S3_BUCKET_FILES} --acl public-read

localstack-setup-dynamodb:
	aws --profile localstack --endpoint-url=${LOCALSTACK_ENDPOINT_URL_DYNAMODB} dynamodb create-table \
		--table-name ${APP_NAME}-${STAGE}-Images \
		--cli-input-json file://cli-input-json/dynamodb-table-images-create.json

localstack-s3-uploads:
	aws --endpoint-url ${LOCALSTACK_ENDPOINT_URL_S3} s3 ls s3://${S3_BUCKET_FILES}

localstack-tables:
	aws --endpoint-url=${LOCALSTACK_ENDPOINT_URL_DYNAMODB} dynamodb list-tables

localstack-table-images:
	aws --endpoint-url=${LOCALSTACK_ENDPOINT_URL_DYNAMODB} dynamodb scan --table-name ${APP_NAME}-${STAGE}-Images