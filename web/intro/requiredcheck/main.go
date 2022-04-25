package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/go-playground/validator/v10"
)

// go-validator-required-start
type Book struct {
	Title string `validate:"required"`
	Price int    `validate:"required"`
}

func main() {
	s := `{"Title":"Real World HTTP ミニ版", "Price":0}` // 無料版のためPriceは0円
	var b Book
	if err := json.Unmarshal([]byte(s), &b); err != nil {
		log.Fatal(err)
	}

	if err := validator.New().Struct(b); err != nil {
		var ve validator.ValidationErrors // validatorの独自型に変換
		if errors.As(err, &ve) {
			for _, fe := range ve {
				fmt.Printf("フィールド %s が %s 違反です（値:%v）\n", fe.Field(), fe.Tag(), fe.Value())
				// フィールド Price が required 違反です（値:0）
			}
		}
	}

}

// go-validator-required-end
