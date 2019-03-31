.PHONY: build clean deploy

build:
	dep ensure -v
	env GOOS=linux go build -ldflags="-s -w" -o bin/images-get src/handlers/images/index/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/images-put src/handlers/images/store/main.go

test:
	go test -v ./src/domains/images
	go test -v ./src/infrastructures/storage
	go test -v ./src/infrastructures/db/repositories

test-no-cache:
	go clean -cache
	make test

clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy: build
	sls deploy --stage dev

remove:
	sls remove --stage dev

prod-deploy: clean build
	sls deploy --stage prod

localstack-setup: localstack-setup-s3 localstack-setup-dynamodb

localstack-setup-s3:
	aws --profile localstack --endpoint-url=http://localhost:4572 s3 mb s3://${S3_BUCKET_FILES}
	aws --profile localstack --endpoint-url=http://localhost:4572 s3api put-bucket-acl --bucket ${S3_BUCKET_FILES} --acl public-read

localstack-setup-dynamodb:
	aws --profile localstack --endpoint-url=http://localhost:4569 dynamodb create-table \
		--table-name ${APP_NAME}-${STAGE}-Images \
		--cli-input-json file://cli-input-json/dynamodb-table-images-create.json

localstack-s3-uploads:
	aws --endpoint-url http://localhost:4572 s3 ls s3://test

localstack-tables:
	aws --endpoint-url=http://localhost:4569 dynamodb list-tables

localstack-table-images:
	aws --endpoint-url=http://localhost:4569 dynamodb scan --table-name ${APP_NAME}-${STAGE}-Images