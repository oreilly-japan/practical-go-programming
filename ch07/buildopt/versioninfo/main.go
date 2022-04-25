package main

import (
	"fmt"
	"log"
	"runtime/debug"
)

func main() {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return
	}

	buf, err := info.MarshalText()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(buf))
}
