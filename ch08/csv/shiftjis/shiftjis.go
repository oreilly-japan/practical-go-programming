package main

import (
	"encoding/csv"
	"fmt"
	"github.com/gocarina/gocsv"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
	"io"
	"log"
	"os"
)

func main() {
	fn := func(in io.Reader) gocsv.CSVReader {
		return csv.NewReader(transform.NewReader(in, japanese.ShiftJIS.NewDecoder()))
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
	Code       int    `csv:"国コード"`
	ISOCode    string `csv:"ISOコード"`
	Population int    `csv:"人口"`
}
