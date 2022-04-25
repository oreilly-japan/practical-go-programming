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

	// dynamodb-update
	input := &dynamodb.UpdateItemInput{
		Key: map[string]types.AttributeValue{
			"id":           &types.AttributeValueMemberS{Value: "00001"},
			"process_date": &types.AttributeValueMemberS{Value: "2020-11-01"},
		},
		TableName: aws.String("example"),
		ExpressionAttributeNames: map[string]string{
			"#Text": "text",
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":text": &types.AttributeValueMemberS{Value: "new text"},
		},
		UpdateExpression: aws.String("SET #Text = :text"),
	}
	if _, err := db.UpdateItem(ctx, input); err != nil {
		// エラーハンドリング
	}
	// dynamodb-update
}
