package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

	"github.com/aws/aws-sdk-go-v2/config"
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

	// dynamodb-consistent
	input := &dynamodb.GetItemInput{
		TableName: aws.String("example"),
		Key: map[string]types.AttributeValue{
			"id":           &types.AttributeValueMemberS{Value: "00001"},
			"process_date": &types.AttributeValueMemberS{Value: "2020-11-01"},
		},
		// 強い整合性がある読み込みを利用
		ConsistentRead: aws.Bool(true),
	}
	// dynamodb-consistent
	out, err := db.GetItem(ctx, input)
	if err != nil {
		// エラーハンドリング
	}
	var item Item
	err = attributevalue.UnmarshalMap(out.Item, &item)
	if err != nil {
		// エラーハンドリング
	}
	fmt.Printf("get item: %+v", item)
}
