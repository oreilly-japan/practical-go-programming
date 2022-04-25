package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sync/atomic"
)

// sample:worker-structure
func Worker(tasks <-chan Task, results chan<- Result) {
	for t := range tasks {

		// ここで何かタスクを行う

		// 結果の送信
		results <- Result{
			Value: v,
			Err:   err,
		}
	}
}

// sample:worker-structure

// sample:pool-task
type Task string

type Result struct {
	Value int64
	Task  Task
	Err   error
}

// sample:pool-task

// sample:pool-worker
func worker(id int, tasks <-chan Task, results chan<- Result) {
	for t := range tasks {
		fmt.Printf("worker: %d task: %s\n", id, t)
		s, err := os.Stat(string(t))
		if err == nil && s.IsDir() {
			err = fmt.Errorf("worker: %d err: %s is dir", id, string(t))
		}
		result := Result{
			Task: t,
		}
		if err != nil {
			result.Err = err
		} else {
			fmt.Printf("worker: %d path: %s size: %d\n", id, string(t), s.Size())
			result.Value = s.Size()
		}
		results <- result
	}
}

// sample:pool-worker

// sample:pool-runner
func TotalFileSize() int64 {
	tasks := make(chan Task)
	results := make(chan Result)

	// ワーカーを起動
	for i := 0; i < runtime.NumCPU(); i++ {
		go worker(i, tasks, results)
	}
	// タスクを非同期でチャネルに投入
	inputDone := make(chan struct{})
	var remainedCount int64
	go func() {
		filepath.Walk(runtime.GOROOT(), func(path string, info os.FileInfo, err error) error {
			atomic.AddInt64(&remainedCount, 1)
			tasks <- Task(path)
			return nil
		})
		close(inputDone)
		close(tasks)
	}()
	// 結果の収集
	var size int64
	for {
		select {
		case result := <-results:
			if result.Err != nil {
				fmt.Printf("err %v for %s\n", result.Err, result.Task)
			} else {
				atomic.AddInt64(&size, result.Value)
			}
			atomic.AddInt64(&remainedCount, -1)
		case <-inputDone:
			if remainedCount == 0 {
				return size
			}
		}
	}
}

// sample:pool-runner

func main() {
	fmt.Println(TotalFileSize())
}
