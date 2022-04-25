package main

import (
	"fmt"
	"strings"
)

func main() {
	// fixedlen-start
	type Book struct {
		ISBN        string
		PublishDate string
		Price       string
		PDF         string
		EPUB        string
		EbookPrice  string
	}

	s := `978-4-87311-865-9201909174620true true 3696
978-4-87311-924-3202010102750falsefalse0000
978-4-87311-878-9201903120000true true 0000`

	for _, line := range strings.Split(s, "\n") {
		r := []rune(line)

		res := Book{
			ISBN:        string(r[0:17]),
			PublishDate: string(r[17:25]),
			Price:       string(r[25:29]),
			PDF:         string(r[29:34]),
			EPUB:        string(r[34:39]),
			EbookPrice:  string(r[39:43]),
		}

		fmt.Printf("%+v\n", res)
		// {ISBN:978-4-87311-865-9 PublishDate:20190917 Price:4620 PDF:true  EPUB:true  EbookPrice:3696}
		// {ISBN:978-4-87311-924-3 PublishDate:20201010 Price:2750 PDF:false EPUB:false EbookPrice:0000}
		// {ISBN:978-4-87311-878-9 PublishDate:20190312 Price:0000 PDF:true  EPUB:true  EbookPrice:0000}
	}
	// fixedlen-end

}
