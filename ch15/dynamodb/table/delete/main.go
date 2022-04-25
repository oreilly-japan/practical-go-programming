package main

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
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

	// dynamodb-delete
	input := &dynamodb.DeleteItemInput{
		Key: map[string]types.AttributeValue{
			"id":           &types.AttributeValueMemberS{Value: "00001"},
			"process_date": &types.AttributeValueMemberS{Value: "2020-11-01"},
		},
		TableName: aws.String("example"),
	}
	if _, err := db.DeleteItem(ctx, input); err != nil {
		// エラーハンドリング
	}
	// dynamodb-delete
}
