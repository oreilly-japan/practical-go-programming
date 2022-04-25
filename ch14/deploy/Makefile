.PHONY: deps clean build

deps:
	go get -u ./...

build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bin/lambda ./cmd/lambda/main.go

zip: build
	zip -j bin/lambda.zip bin/lambda
