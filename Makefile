up:
	echo "hello"

build:
	GOARCH=amd64 GOOS=linux CGO_ENABLED=0 go build -o bin/main ./main.go

clean:
	rm -rf ./bin ./vendor

local: build
	go run main.go

deploy: clean build
	serverless deploy --verbose --region us-east-1

deploy-dev: clean build
	serverless deploy --verbose --region us-east-1 --stage=dev