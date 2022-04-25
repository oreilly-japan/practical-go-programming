package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// optional-struct
	type FormInput struct {
		Name        string `json:"name"`
		CompanyName string `json:"company_name,omitempty"`
	}

	in := FormInput{Name: "山田太郎"}

	b, _ := json.Marshal(in)
	fmt.Println(string(b))
	// {"name":"山田太郎"}
	// optional-struct
}
