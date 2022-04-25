package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
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

	// dynamodb-query
	input := &dynamodb.QueryInput{
		TableName: aws.String("example"),
		ExpressionAttributeNames: map[string]string{
			"#ID": "id",
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":id":    &types.AttributeValueMemberS{Value: "00001"},
			":month": &types.AttributeValueMemberS{Value: "2020-10"},
		},
		KeyConditionExpression: aws.String("#ID = :id and begins_with(process_date, :month)"),
	}
	out, err := db.Query(ctx, input)
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
	// dynamodb-query
}
