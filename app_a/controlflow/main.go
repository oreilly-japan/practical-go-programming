package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	statusCode := 1
	cache := map[int]int{
		2: 4,
		3: 9,
		4: 16,
	}
	input := 2
	// control-flow-if1
	// if文/if else/elseの基本的な書き方
	if statusCode == 200 {
		fmt.Println("no error")
	} else if statusCode < 500 {
		fmt.Println("client error")
	} else {
		fmt.Println("server error")
	}
	// control-flow-if1
	// control-flow-if2
	// データ取得とチェックを同時に行う
	if result, ok := cache[input]; ok {
		fmt.Println("cached value", result)
	}
	// control-flow-if2

	// control-flow-for1
	// スライスやマップの書く要素に対してループ
	scketches := []string{"Dead Parrot", "Killer joke", "Spanish Inquisition", "Spam"}
	for i, s := range scketches {
		fmt.Println(i, s)
	}
	// control-flow-for1

	// control-flow-for2
	// 1変数だけ書けばインデックスのみを受け取れる
	for i := range scketches {
		fmt.Println(i)
	}

	// ブランク識別子でインデックスを読み飛ばして値だけを使う
	for _, s := range scketches {
		fmt.Println(s)
	}
	// control-flow-for2

	// control-flow-for3
	for _, s := range scketches {
		// もしスケッチ名がKから始まっていたら読み飛ばす
		if strings.HasPrefix(s, "K") {
			continue
		}
		// もしスケッチ名がnで終わっていたらループ終了
		if strings.HasSuffix(s, "n") {
			break
		}
	}
	// control-flow-for3

	// control-flow-for4
	counter := 0
	for counter < 10 {
		fmt.Println("ブール値がtrueの間回り続けるループ")
		counter += 1
	}

	end := time.Now().Add(time.Second)
	for {
		fmt.Println("breakやreturnで抜けないと終わらないループ")
		if end.Before(time.Now()) {
			break
		}
	}
	// control-flow-for4

	// control-flow-for5
	for i := 0; i < 10; i++ {
		fmt.Println("10回繰り返す")
	}
	// control-flow-for5

	s := "run"
	// control-flow-switch1
	switch s {
	case "running":
		fmt.Println("実行中")
	case "stop":
		fmt.Println("停止中")
	default:
		fmt.Println("その他")
	}
	// control-flow-switch1

	// control-flow-switch2
	switch s {
	case "running":
		fallthrough // "runningの時も実行中と表示される"
	case "run":
		fmt.Println("実行中")
	case "stop":
		fmt.Println("停止中")
	default:
		fmt.Println("その他")
	}
	// control-flow-switch2
}
