package main

import (
	"log"
	"os"

	"github.com/gocarina/gocsv"
)

type Country struct {
	Name       string `csv:"国名"`
	ISOCode    string `csv:"ISOコード"`
	Population int    `csv:"人口"`
}

func main() {
	lines := []Country{
		{Name: "アメリカ合衆国", ISOCode: "US/USA", Population: 310232863},
		{Name: "日本", ISOCode: "JP/JPN", Population: 127288000},
		{Name: "中国", ISOCode: "CN/CHN", Population: 1330044000},
	}

	f, err := os.Create("country.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if err := gocsv.MarshalFile(&lines, f); err != nil {
		log.Fatal(err)
	}
}
