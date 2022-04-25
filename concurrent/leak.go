package main

import (
	"bytes"
	"io"
	"os"
	"time"
)

/*
func leak() {
	ic := make(chan int)
	tasks := make(chan int)
	// sample:infinite-recv1
	// NGなケース
	go func() {
		for {
			task := <-tasks
			// タスクを実行
		}
	}()
	// sample:infinite-recv1

	// sample:infinite-recv2
	// チャネルクローズで終了できるようにする
	go func() {
		// forはチャネルの終了で自然に抜ける
		for task := range tasks {
			// タスクを実行
		}
	}()
	// sample:infinite-recv2

	// sample:infinite-send1
	// NGなケース
	go func() {
		count := 0
		for {
			ic <- count
			count++
		}
	}()
	// sample:infinite-send1

	// sample:infinite-send2
	// OKなケース
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func() {
		count := 0
		for {
			select {
			case ic <- count:
				ic <- count
				count++
			case <-ctx.Done():
				return // タイムアウトなど
			}
		}
	}()
	// sample:infinite-send2
}
*/

// sample:infinite-counter
type InfiniteCounter struct {
	C    chan int
	exit chan struct{}
}

func NewInfiniteCounter() *InfiniteCounter {
	r := &InfiniteCounter{
		C:    make(chan int),
		exit: make(chan struct{}),
	}
	go func() {
		count := 0
		for {
			select {
			case r.C <- count:
				count++
			case <-r.exit:
				close(r.C)
				return
			}
		}
	}()
	return r
}

func (ic *InfiniteCounter) Close() {
	close(ic.exit)
}

// sample:infinite-counter

/*
func main() {
	c := NewInfiniteCounter()
	log.Println(<-c.C)
	log.Println(<-c.C)
	log.Println(<-c.C)
	c.Close()
	log.Println(<-c.C)
}*/

func leak2() {
	go func() {
		for {
			time.Sleep(time.Second)
		}
	}()
	// sample:infinite-wait
	// NGなケース
	wait := make(chan struct{})
	go func() {
		f, err := os.Open("important.txt")
		if err != nil {
			return // closeせずにエラーで関数抜けてしまった！
		}
		var buffer bytes.Buffer
		io.Copy(&buffer, f)
		close(wait)
	}()
	<-wait
	// sample:infinite-wait
}

func main() {
	leak2()
}
