package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/xuri/excelize/v2"
)

type Country struct {
	Name       string
	ISOCode    string
	Population int
}

// exelize-start
func main() {
	f, err := excelize.OpenFile("Book1.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	rows, err := f.Rows("Sheet1")
	if err != nil {
		log.Fatal(err)
	}

	var countries []Country
	for i := 0; rows.Next(); i++ {
		row, err := rows.Columns()
		if err != nil {
			log.Fatal(err)
		}
		if i == 0 { // ヘッダー行のためスキップ
			continue
		}

		population, err := strconv.Atoi(row[2]) // 型変換が必要
		if err != nil {
			log.Fatal(err)
		}

		countries = append(countries, Country{
			Name:       row[0],
			ISOCode:    row[1],
			Population: population,
		})
	}

	fmt.Println(countries)
	// [{アメリカ合衆国 US/USA 310242863} {日本 JP/JPN 127288000} {中国 CN/CHN 1330044000}]
}

// exelize-end
