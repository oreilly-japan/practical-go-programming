package storage

import "fmt"

func Save(data string) {
	fmt.Printf("data %s is saved to CloudWatch Logs\n", data)
}
