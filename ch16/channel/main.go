package main

import "log"

func channel() {
	// sample:chan-basic
	// バッファなしintチャネル作成
	ic := make(chan int)

	// バッファありstringチャネル作成(バッファ量は10)
	sc := make(chan string, 10)
	// sample:chan-basic

	// sample:chan-send-receive
	// チャネルに送信
	ic <- 100

	// チャネルから受信（送信結果は捨てる）
	<-sc

	// チャネルから受信（送信結果を変数に入れる）
	r := <-sc

	// チャネルから受信（送信結果とチャネルの状態を変数に入れる）
	r, ok := <-sc
	// sample:chan-send-receive

	// 待ち用のゼロバイト（空構造体）チャネル
	wc := make(chan struct{})

	// ゼロバイトの送信
	wc <- struct{}{}

	// 送信されるのを待つ（情報は来ないので変数に受信情報を入れる必要はない）
	<-wc

	// sample:chan-close
	c := make(chan int)
	// チャネルのクローズ
	close(c)
	// sample:chan-close

	// チャネルから値と、クローズしているかどうかを取得
	v, ok := <-c

	log.Println(r)
}

/*
// sample:chan-error
func recv(r <-chan string) {
	v := <-r
	r <- "送信はNG" // エラー
}

func send(s chan<- string) {
	s <- "送信は可能"
	v := <-s // エラー
}
// sample:chan-error
*/
/*
func main() {
	channel()
}
*/
