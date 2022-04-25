package main

import (
	"fmt"
	"log"

	"github.com/xuri/excelize/v2"
)

func main() {
	// ファイルの書き込み
	out := excelize.NewFile()
	out.SetCellValue("Sheet1", "A1", "Hello Excel")
	if err := out.SaveAs("Book1.xlsx"); err != nil {
		fmt.Println(err)
	}

	// ファイルの読み込み
	in, err := excelize.OpenFile("Book1.xlsx")
	if err != nil {
		log.Fatal(err)
	}
	cell, err := in.GetCellValue("Sheet1", "A1")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cell)
	// Hello Excel

}
