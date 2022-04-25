package main

import (
	"log"
	"os"

	"github.com/gocarina/gocsv"
)

// chanwrite-sample-start
type record struct {
	Number  int    `csv:"number"`
	Message string `csv:"message"`
}

func main() {
	c := make(chan interface{})
	go func() {
		defer close(c)
		for i := 0; i < 1000*1000; i++ {
			c <- record{
				Message: "Hello",
				Number:  i + 1,
			}
		}
		return
	}()

	if err := gocsv.MarshalChan(c, gocsv.DefaultCSVWriter(os.Stdout)); err != nil {
		log.Fatal(err)
	}

}

// chanwrite-sample-end
