package main

// dynamodb-model
type Item struct {
	ID            string `dynamodbav:"id"`
	ProcessDate   string `dynamodbav:"process_date"`
	Text          string `dynamodbav:"text"`
	TextOmitEmpty string `dynamodbav:"text_omit_empty,omitempty"`
}

// dynamodb-model
