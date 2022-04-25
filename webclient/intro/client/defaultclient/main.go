package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// client-simple-start
	resp, err := http.Get("http://example.com/")
	if err != nil {
		// ...
	}
	// defer でレスポンスボディをクローズする
	defer resp.Body.Close()

	// レスポンスのステータスコードを確認する
	if resp.StatusCode != http.StatusOK {
		// ...
	}
	// client-simple-end
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// ...
	}
	fmt.Println(string(body))
}
