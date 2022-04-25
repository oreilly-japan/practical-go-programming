package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gocarina/gocsv"
)

type Country struct {
	Name       string `csv:"国名"`
	ISOCode    string `csv:"ISOコード"`
	Population int    `csv:"人口"`
}

// gocsv-sample-start
func main() {
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
		// {Name:アメリカ合衆国 ISOCode:US/USA Population:310232863}
		// {Name:日本 ISOCode:JP/JPN Population:127288000}
		// {Name:中国 ISOCode:CN/CHN Population:1330044000}
	}

}

// gocsv-sample-end
