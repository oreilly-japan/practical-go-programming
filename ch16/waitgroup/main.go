package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
)

func waitgroup() {
	// sample:waitgroup-sample
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		time.Sleep(time.Second)
		fmt.Println("Done: 1")
		wg.Done()
	}()
	go func() {
		time.Sleep(time.Second * 2)
		fmt.Println("Done: 2")
		wg.Done()
	}()
	go func() {
		time.Sleep(time.Second * 3)
		fmt.Println("Done: 3")
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("Done all tasks")
	// sample:waitgroup-sample
}

func eg() {
	ctx := context.Background()
	// sample:errgroup-sample
	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		time.Sleep(time.Second)
		fmt.Println("Done: 1")
		return nil
	})
	eg.Go(func() error {
		time.Sleep(time.Second * 2)
		fmt.Println("Done: 2")
		return nil
	})
	eg.Go(func() error {
		time.Sleep(time.Second * 3)
		fmt.Println("Done: 3")
		return nil
	})
	err := eg.Wait()
	fmt.Println("Done all tasks", err)
	// sample:errgroup-sample
}

func egctx() {
	ctx := context.Background()
	eg, ctx := errgroup.WithContext(ctx)
	log.Println(ctx, eg)
}

/*
func main() {
	eg()
}
*/
