package main

import (
	"log"

	"github.com/xuri/excelize/v2"
)

// exelize-write-start
type Country struct {
	Name       string
	ISOCode    string
	Population int
}

func (c Country) HeaderColumns() []interface{} {
	return []interface{}{"国名", "ISOコード", "人工"}
}

func (c Country) Columns() []interface{} {
	return []interface{}{c.Name, c.ISOCode, c.Population}
}

func main() {
	lines := []Country{
		{Name: "アメリカ合衆国", ISOCode: "US/USA", Population: 310232863},
		{Name: "日本", ISOCode: "JP/JPN", Population: 127288000},
		{Name: "中国", ISOCode: "CN/CHN", Population: 1330044000},
	}

	f := excelize.NewFile()
	sw, err := f.NewStreamWriter("Sheet1")
	if err != nil {
		log.Fatal(err)
	}

	for i, line := range lines {
		if i == 0 {
			cell, _ := excelize.CoordinatesToCellName(1, i+1)
			sw.SetRow(cell, line.HeaderColumns())
		}
		cell, _ := excelize.CoordinatesToCellName(1, i+2)
		sw.SetRow(cell, line.Columns())
	}

	if err := sw.Flush(); err != nil {
		log.Fatal(err)
	}

	if err := f.SaveAs("Book2.xlsx"); err != nil {
		log.Fatal(err)
	}

}

// exelize-write-end
