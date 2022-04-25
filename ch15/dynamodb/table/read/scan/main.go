package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func main() {
	ctx := context.Background()

	// dynamodb-client
	cfg, err := config.LoadDefaultConfig(ctx) // 環境変数などから設定情報を読み込み
	if err != nil {
		log.Fatalf("failed to load SDK configuration, %v", err)
	}
	db := dynamodb.NewFromConfig(cfg)
	// dynamodb-client

	// dynamodb-scan
	input := &dynamodb.ScanInput{
		TableName: aws.String("example"),
	}
	out, err := db.Scan(ctx, input)
	if err != nil {
		// エラーハンドリング
	}
	var items []Item
	err = attributevalue.UnmarshalListOfMaps(out.Items, &items)
	if err != nil {
		// エラーハンドリング
	}
	for _, v := range items {
		fmt.Printf("query: %+v\n", v)
	}
	// dynamodb-scan
}
