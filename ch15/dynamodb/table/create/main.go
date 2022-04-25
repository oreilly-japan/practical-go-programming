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

	// dynamodb-create-table

	attributeDefinitions := []types.AttributeDefinition{
		{AttributeName: aws.String("id"), AttributeType: types.ScalarAttributeTypeS},
		{AttributeName: aws.String("process_date"), AttributeType: types.ScalarAttributeTypeS},
	}
	keySchema := []types.KeySchemaElement{
		{AttributeName: aws.String("id"), KeyType: types.KeyTypeHash},
		{AttributeName: aws.String("process_date"), KeyType: types.KeyTypeRange},
	}
	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: attributeDefinitions,
		BillingMode:          types.BillingModePayPerRequest,
		KeySchema:            keySchema,
		TableName:            aws.String("example"),
	}
	if _, err := db.CreateTable(ctx, input); err != nil {
		// エラーハンドリング
	}
	// dynamodb-create-table
}
