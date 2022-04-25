package main

import (
	"context"
	"net/http"
	"sampleapp"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
)

func LambdaHandler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", sampleapp.Handler)

	adapter := httpadapter.New(mux)
	return adapter.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(LambdaHandler)
}
