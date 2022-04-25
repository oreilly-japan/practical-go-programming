package main

// dynamodb-model
type Item struct {
	ID          string `dynamodbav:"id"`
	ProcessDate string `dynamodbav:"process_date"`
	Text        string `dynamodbav:"text"`
	DeleteTTL   int64  `dynamodbav:"delete_ttl"`
}

// dynamodb-model
