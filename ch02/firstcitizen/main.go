package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"reflect"
)

func main() {
	// first-citizen
	var int int = 1
	// first-citizen
	log.Println(int)
}

func typeOf() {
	// typeof
	writerType := reflect.TypeOf((*io.Writer)(nil)).Elem()
	
	fileType := reflect.TypeOf((*os.File)(nil))
	fmt.Println(fileType.Implements(writerType))
	// typeof
}