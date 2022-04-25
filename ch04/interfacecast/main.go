package main

import (
	"context"
	"io"
	"log"
)

func Parse(r io.Reader) {
	// cast-to-closer
	if c, ok := r.(io.Closer); ok {
		c.Close()
	}
	// cast-to-closer
}

func main() {
	// use-context-value
	// "favorite"というキーに対して、"銭形平次"という値を設定する
	// context.Contextについては並行処理の章で詳しく説明します
	ctx := context.WithValue(context.Background(), "favorite", "銭形平次")

	// ctx.Value()はinterface{]なので変換が必要
	// 必ず、okで成功可否を確認すること
	if s, ok := ctx.Value("favorite").(string); ok {
		// sはstring型
		log.Printf("私の好きなものは%sです\n", s)
	}

	// あるいは型スイッチ
	// 同じvだが、case節ごとにvの型が変わる
	switch v := ctx.Value("favorite").(type) {
	case string:
		log.Printf("好きなものは: %s\n", v)
	case int:
		log.Printf("好きな数値は: %d\n", v)
	case complex128:
		log.Printf("好きな複素数は: %f\n", v)
	default: // どれにもマッチしない場合
		log.Printf("好きなものは: %v\n", v)
	}
	// use-context-value

	// cast-success
	// anyの値を型のあるスライスに代入
	var fishList = []any{"鯖", "鰤", "鮪"}
	fishNames := make([]string, len(fishList))
	for i, f := range fishList {
		// ダウンキャストは型アサーションが必要
		if fn, ok := f.(string); ok {
			fishNames[i] = fn
		}
	}

	// 型のあるスライスをanyのスライスに代入
	fibonacciNumbers := []int{1, 1, 2, 3, 5, 8}
	anyValues := make([]any, len(fibonacciNumbers))
	for i, fn := range fibonacciNumbers {
		// アップキャストは型アサーション不要だが要素ごとに入れる必要あり
		anyValues[i] = fn
	}
	// cast-success
}

/*
	// cast-error
	type Fish interface{}
	var fishList = []Fish{"鯖", "鰤", "鮪"}
	var fishNameList = fishList.([]string) // このように一括変換はできない
	var anyList []any = fishList           // anyへの一括アップキャストもできない
	// cast-error
	log.Println(fishList)
*/
