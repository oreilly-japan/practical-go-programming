package main

import (
	"fmt"
	"github.com/gocarina/gocsv"
	"log"
	"os"
)

type record struct {
	Number  int    `csv:"number"`
	Message string `csv:"message"`
}

// chanread-sample-start
func main() {
	f, err := os.Open("large.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	c := make(chan record)
	done := make(chan bool)
	go func() {
		if err := gocsv.UnmarshalToChan(f, c); err != nil {
			log.Fatal(err)
		}
		done <- true
	}()

	for {
		select {
		case v := <-c:
			fmt.Printf("%+v\n", v)
		case <-done:
			return
		}
	}

}

// chanread-sample-end
