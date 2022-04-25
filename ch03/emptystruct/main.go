package main

import (
	"fmt"
	"reflect"
)

func send() {
	// empty-struct-use
	wait := make(chan struct{})
	go func() {
		// 空の構造体のインスタンスを送信
		fmt.Println("送信")
		wait <- struct{}{}
	}()
	// データサイズゼロのインスタンスを受け取る
	fmt.Println("受信待ち")
	<-wait
	fmt.Println("受信完了")
	// empty-struct-use
}

func size() {
	// empty-struct-size
	emptyStructType := reflect.TypeOf(&struct{}{}).Elem()
	fmt.Println(emptyStructType.Size())
	// 0
	// empty-struct-size
}

func main() {
	size()
	send()
}
