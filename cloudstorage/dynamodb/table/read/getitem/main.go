package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

	"github.com/aws/aws-sdk-go-v2/aws"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func main() {

	// dynamodb-get-item
	ctx := context.Background()
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatalf("failed to load SDK configuration, %v", err)
	}
	db := dynamodb.NewFromConfig(cfg)

	input := &dynamodb.GetItemInput{
		TableName: aws.String("example"),
		Key: map[string]types.AttributeValue{
			"id":           &types.AttributeValueMemberS{Value: "00001"},
			"process_date": &types.AttributeValueMemberS{Value: "2020-11-01"},
		},
	}
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

	// dynamodb-get-item

}
