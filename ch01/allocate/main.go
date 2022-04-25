package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"strconv"
)

//go:noinline
func waistSlice() {
	// allocate-test
	var a []int
	var lastPtr *int
	for i := 0; i < 10000; i++ {
		a = append(a, i)
		if &a[0] != lastPtr {
			fmt.Println(i)
		}
		lastPtr = &a[0]
	}
	// allocate-test
}

//go:noinline
func waistMap() {
	a := make(map[string]int)
	for i := 0; i < 10000; i++ {
		a[strconv.FormatInt(int64(i), 10)] = i
	}
}

func main() {
	alloc()
}

func alloc() {
	{
		f, err := os.Create("cpu.pprof")
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer f.Close() // error handling omitted for example
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}
	{
		f, err := os.Create("mem.pprof")
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		defer f.Close() // error handling omitted for examplep
		runtime.GC()    // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
	}

	// alloc-slice
	// 正確な長さがわかっている場合
	s1 := make([]int, 1000)
	fmt.Println(len(s1)) // 1000
	fmt.Println(cap(s1)) // 1000

	// 正確な長さがわからないが最大量の見込みがつく場合
	// キャパシティだけ増やす
	s2 := make([]int, 0, 1000)
	fmt.Println(len(s2)) // 0
	fmt.Println(cap(s2)) // 1000
	// alloc-slice

	// alloc-map
	m := make(map[string]string, 1000)
	fmt.Println(len(m)) // 0
	// alloc-map

	// waist
	waistSlice()
	waistMap()
}

func deferSample() error {
	var files []string
	var result [][]byte
	// defer-loop
	for _, fname := range files {
		f, err := os.Open(fname)
		if err != nil {
			return err
		}
		// この書き方だとClose()は全部のループを抜けるまで実行されない
		// deferを使わずにファイルを使った後に自分でf.Close()を呼ぶ
		defer f.Close()
		data, _ := io.ReadAll(f)
		result = append(result, data)
	}
	// defer-loop
	return nil
}

// defer-close
func deferReturnSample(fname string) (err error) {
	var f *os.File
	f, err = os.Create(fname)
	if err != nil {
		return fmt.Errorf("ファイルオープンのエラー %w", err)
	}
	defer func() {
		// Closeのエラーを拾って名前付き返り値に代入
		// すでにerrに別のものが入る可能性があるときは
		// さらに要注意
		err = f.Close()
	}()
	io.WriteString(f, "deferのエラーを拾うサンプル")
	return
}

// defer-close
