package main

import (
	"encoding/json"
	"fmt"
	"time"
)

// JSONに共通的に含めたい項目をCommon構造体に定義し
// 使う側で構造体を埋め込む
type Individual struct {
	ID string `json:"id"`
	Common
}

type Common struct {
	CreateAt time.Time `json:"create_at"`
}

func main() {
	ind := Individual{
		ID: "hoge",
		Common: Common{
			CreateAt: time.Date(2020, time.September, 9, 9, 35, 00, 0, time.Local),
		},
	}

	b, _ := json.Marshal(ind)
	fmt.Printf("%+v", string(b))
}
