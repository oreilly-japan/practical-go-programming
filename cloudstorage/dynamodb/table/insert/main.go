// dynamodb-insert
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
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatalf("failed to load SDK configuration, %v", err)
	}
	db := dynamodb.NewFromConfig(cfg)

	item := Item{
		ID:            "00001",
		ProcessDate:   time.Now().Format("2006-01-02"),
		Text:          "example text",
		TextOmitEmpty: "",
	}
	av, err := attributevalue.MarshalMap(item)
	if err != nil {
		log.Fatalf("failed to marshal, item = %v, %v", item, err)
	}
	input := &dynamodb.PutItemInput{
		TableName: aws.String("example"),
		Item:      av,
	}
	_, err = db.PutItem(ctx, input)
	if err != nil {
		log.Fatalf("failed to put item, %v", err)
	}
}

// dynamodb-insert
