package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/ianlopshire/go-fixedwidth"
)

func main() {

	type Book struct {
		F1 bool `fixed:"1,5,left, "`
		F2 bool `fixed:"6,10,left, "`
		F3 int  `fixed:"11,12"`
	}

	s := `true true 1
falsefalse2
true false3`

	for _, s := range strings.Split(s, "\n") {
		var b Book
		if err := fixedwidth.Unmarshal([]byte(s), &b); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%+v\n", b)
	}

}
