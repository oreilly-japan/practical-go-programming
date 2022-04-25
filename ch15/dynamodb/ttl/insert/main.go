package main

import (
	"context"
	"log"
	"time"

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

	// dynamodb-insert
	item := Item{
		ID:          "00001",
		ProcessDate: time.Now().Format("2006-01-02"),
		Text:        "example text",
		DeleteTTL:   time.Now().Add(5 * 24 * time.Hour).Unix(),
	}
	av, err := attributevalue.MarshalMap(item)
	if err != nil {
		// エラーハンドリング
	}
	input := &dynamodb.PutItemInput{
		TableName: aws.String("example"),
		Item:      av,
	}
	if _, err = db.PutItem(ctx, input); err != nil {
		// エラーハンドリング
	}
	// dynamodb-insert
}
