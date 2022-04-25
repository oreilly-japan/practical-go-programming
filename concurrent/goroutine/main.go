package main

import (
	"fmt"
	"time"
)

func goroutine() {
	// sample:goroutine-sample
	fmt.Println("ゴルーチンを実行します")
	go func() {
		fmt.Println("ゴルーチンが実行しています")
	}()

	fmt.Println("ゴルーチンの終了を待ちます")
	time.Sleep(time.Second) // 説明のためにあえて悪いコードを使っています
	fmt.Println("ゴルーチンが終了しました")
	// ゴルーチンを実行します
	// ゴルーチンの終了を待ちます
	// ゴルーチンが実行しています
	// ゴルーチンが終了しました
	// sample:goroutine-sample
}

/*
// sample:goroutine-bug
func main() {
	go func() {
		fmt.Println("goroutineで実行しています")
	}()
	fmt.Println("goroutineを実行しました")
}
// sample:goroutine-bug
*/

func loop() {
	// sample:goroutine-loop-bug
	items := []int{1, 2, 3, 4, 5, 6, 7, 8}
	for _, v := range items {
		go func() {
			fmt.Printf("v = %d\n", v)
		}()
	}
	time.Sleep(time.Second)
	// sample:goroutine-loop-bug
}

func loop2() {
	// sample:goroutine-loop-bug-fix1
	items := []int{1, 2, 3, 4, 5, 6, 7, 8}
	for _, v := range items {
		v2 := v
		go func() {
			fmt.Printf("v2 = %d, address = %p\n", v2, &v2)
		}()
	}
	// sample:goroutine-loop-bug-fix1
	time.Sleep(time.Second)
}

func loop3() {
	// sample:goroutine-loop-bug-fix2
	items := []int{1, 2, 3, 4, 5, 6, 7, 8}
	for _, v := range items {
		go func(v int) {
			fmt.Printf("v = %d\n", v)
		}(v)
	}
	// sample:goroutine-loop-bug-fix2
	time.Sleep(time.Second)
}

func chanloop() {
	// sample:chan-loop
	ic := make(chan int)
	go func() {
		ic <- 10
		ic <- 20
		close(ic) // クローズするとループ解除
	}()
	// チャネルに対してループ
	for v := range ic {
		fmt.Println(v)
	}
	// sample:chan-loop
}

func main() {
	loop2()
}
