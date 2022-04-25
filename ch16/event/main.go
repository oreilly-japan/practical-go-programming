package main

import (
	"context"
	"errors"
	"os/exec"
	"time"
)

func event() {
	// sample:empty-chan
	// ゼロバイトの構造体のチャネル
	wait := make(chan struct{})

	// 受信
	<-wait
	// sample:empty-chan

	// sample:send-empty-chan
	// 送信
	wait <- struct{}{}
	// sample:send-empty-chan

	// sample:close-chan
	// 送信の代わりにクローズ
	close(wait)
	// sample:close-chan
}

// sample:run-job
func runJobs(ctx context.Context) error {
	// cancel()関数を作成
	// deferで関数を抜けるときに自動で呼ばれるようにする
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	// エラーと終了情報の受信用チャネル
	ec := make(chan error)
	done := make(chan struct{})
	// sample:run-job

	// sample:start-task
	// タスクを10個起動
	for i := 0; i < 10; i++ {
		go func() {
			// Windowsはtimeoutにしてください
			cmd := exec.CommandContext(ctx, "sleep", "30")
			err := cmd.Run()
			if err != nil {
				ec <- err
			} else {
				done <- struct{}{}
			}
		}()
	}
	// sample:start-task

	// sample:wait-finish
	// 終了をまつ
	for i := 0; i < 10; i++ {
		select {
		case err := <-ec:
			// ここでエラーを返して関数を終了させると
			// deferでcancel()が呼ばれ、実行中の別のタスクが終了する
			return err
		case <-done:
		}
	}
	// sample:wait-finish
	return nil
}

/*
func main() {
	// sample:init-context
	ctx := context.Background()
	err := runJobs(ctx)
	// sample:init-context
ß	log.Println(err)
	// time.Sleep(10 * time.Second) // 終了状態を確認したい場合にコメントアウトしてください
}
*/
func runJobsWithError(ctx context.Context) error {
	// cancel()関数を作成
	// deferで関数を抜けるときに自動で呼ばれるようにする
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	// エラーと終了情報の受信用チャネル
	ec := make(chan error)
	done := make(chan struct{})

	// タスクを10個起動
	for i := 0; i < 10; i++ {
		go func() {
			// Windowsはtimeoutにしてください
			cmd := exec.CommandContext(ctx, "sleep", "30")
			err := cmd.Run()
			if err != nil {
				ec <- err
			} else {
				done <- struct{}{}
			}
		}()
	}
	// sample:abort
	go func() {
		time.Sleep(10 * time.Second)
		ec <- errors.New("accidental error")
	}()
	// sample:abort
	// 終了をまつ
	for i := 0; i < 11; i++ {
		select {
		case err := <-ec:
			// ここでエラーを返して関数を終了させると
			// deferでcancel()が呼ばれ、実行中の別のタスクが終了する
			return err
		case <-done:
		}
	}
	return nil
}
