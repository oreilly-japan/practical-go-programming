package main

import (
	"encoding/csv"
	"fmt"
	"github.com/gocarina/gocsv"
	"github.com/spkg/bom"
	"io"
	"log"
	"os"
)

func main() {

	fn := func(in io.Reader) gocsv.CSVReader {
		r := csv.NewReader(bom.NewReader(in)) // BOMの回避
		return r
	}
	gocsv.SetCSVReader(fn)

	f, err := os.Open("country.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var lines []Country
	if err := gocsv.UnmarshalFile(f, &lines); err != nil {
		log.Fatal(err)
	}

	for _, v := range lines {
		fmt.Printf("%+v\n", v)
	}

}

type Country struct {
	Name       string `csv:"国名"`
	ISOCode    string `csv:"ISOコード"`
	Population int    `csv:"人口"`
}
