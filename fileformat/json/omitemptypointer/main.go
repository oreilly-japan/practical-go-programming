package main

import (
	"encoding/json"
	"fmt"
)

// optional-pointer-struct
type Bottle struct {
	Name  string `json:"name"`
	Price int    `json:"price,omitempty"`
	KCal  *int   `json:"kcal,omitempty"` // *intのポインター型で宣言
}

func main() {
	b := Bottle{
		Name:  "ミネラルウォーター",
		Price: 0,
		KCal:  Int(0),
	}

	out, _ := json.Marshal(b)
	fmt.Println(string(out))
	// {"name":"ミネラルウォーター","kcal":0}
}

func Int(v int) *int {
	return &v
}

// optional-pointer-struct
