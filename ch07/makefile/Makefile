.PHONY: install test run build zip deploy-dev deploy-prod

#############
# Development
#############
.EXPORT_ALL_VARIABLES:
AWS_ACCESS_KEY_ID=test
AWS_SECRET_ACCESS_KEY=test
AWS_REGION=ap-northeast-1

clean:
	rm -rf ./bin/*

install:
	go install golang.org/x/tools/cmd/stringer@v0.1.8
	go generate
	go mod download

test:
	go test -race $(option) ./...

run:
	go run cmd/main/main.go --port 8000

#############
# Deployment
#############
build: install clean
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w -buildid=" -trimpath -o bin/lambda cmd/lambda/lambda.go

zip: build
	zip -j bin/lambda.zip bin/lambda

deploy-dev: zip
	aws lambda update-function-code --profile my_profile_dev --region ap-northeast-1 --function-name dev-example-api --zip-file fileb://bin/lambda.zip

deploy-prod: zip
	aws lambda update-function-code --profile my_profile_prod --region ap-northeast-1 --function-name prod-example-api --zip-file fileb://bin/lambda.zip
