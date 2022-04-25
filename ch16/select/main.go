package main

import (
	"context"
	"log"
	"time"
)

func selecttest() {
	send := make(chan int)
	recv := make(chan int)
	recv2 := make(chan int)
	recv3 := make(chan int)

	go func() {
		time.Sleep(time.Second)
		close(send)
	}()

	// sample:select-sample
	select {
	case <-recv:
		log.Println("受信があったこと or チャネルのクローズを感知 (どちらか判断はできない、値は捨てる)")
	case v := <-recv2:
		log.Println("受信があったこと or チャネルのクローズを感知 (どちらか判断はできない、値は受け取る):", v)
	case v, ok := <-recv3:
		log.Println("受信があったこと or チャネルのクローズを感知 (チャネルの状態と値は受け取る):", v, ok)
	case send <- 10:
		log.Println("送信が成功したことを検知")
	default:
		log.Println("どのチャネルの送受信も行われなかった")
	}
	// sample:select-sample
}

func interrupt() {
	recv := make(chan int)
	ctx1, cancel1 := context.WithCancel(context.Background())
	defer cancel1()

	ctx2, cancel2 := context.WithTimeout(context.Background(), time.Second)
	defer cancel2()

	go func() {
		time.Sleep(time.Millisecond * 500)
		cancel1()
	}()

	select {
	case <-recv:
		log.Println("情報があれば受け取りたいが、いつまでもブロックしたくないチャネル")
	case <-ctx1.Done():
		log.Println("cancel1()が呼ばれて中断")
	case <-ctx2.Done():
		log.Println("タイムアウトで中断")
	}
}

/*
func main() {
	//selecttest()
	interrupt()
}*/
