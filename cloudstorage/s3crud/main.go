package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func main() {

	cfg, err := config.LoadDefaultConfig(context.TODO()) // 環境変数などから設定情報を読み込み
	if err != nil {
		log.Fatalf("failed to load SDK configuration, %v", err)
	}

	client := s3.NewFromConfig(cfg)

	var token *string
	for {
		resp, err := client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
			Bucket:            aws.String("bucket-name"),         // バケット名を指定
			Prefix:            aws.String("oreiily/2021/12/31/"), // oreiily/2021/12/31/ から始まるファイルに絞る設定
			ContinuationToken: token,
		})
		if err != nil {
			log.Fatalf("list objects, %v", err)
		}

		for _, c := range resp.Contents {
			fmt.Printf("Name:%s LastModified:%s\n", *c.Key, c.LastModified.Format(time.RFC3339))
		}

		if resp.ContinuationToken == nil { // ページング対応
			break
		}
		token = resp.ContinuationToken
	}

}
