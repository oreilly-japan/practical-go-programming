package main

import (
	"runtime"

	"go.uber.org/ratelimit"
)

type Task string

type Result struct {
	Value int64
	Task  Task
	Err   error
}

func main() {
	poolWithRateLimit()
}

// sample:rate-limited-worker
func workerWithRateLimit(rt ratelimit.Limiter, tasks <-chan Task, results chan<- Result) {
	for t := range tasks {
		rt.Take() // 待つ
		// 秒間呼び出し回数の限界があるなにかを実行
		result := Result{
			Task: t,
		}
		results <- result
	}
}

// sample:rate-limited-worker

func poolWithRateLimit() {
	tasks := make(chan Task, 100)
	results := make(chan Result)
	// sample:rate-limited-worker-launcher
	rl := ratelimit.New(100) // 秒間100回
	// ワーカーを起動
	for i := 0; i < runtime.NumCPU(); i++ {
		go workerWithRateLimit(rl, tasks, results)
	}
	// sample:rate-limited-worker-launcher
}
