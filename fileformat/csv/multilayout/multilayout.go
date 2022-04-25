package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"strings"

	"github.com/gocarina/gocsv"
)

// multilayout-sample-start
type Summary struct {
	RecordType string
	Summary    string
}

type Country struct {
	RecordType string
	Name       string
	ISOCode    string
	Population int
}

// csv.CSVReader を満たした構造体
type singleCSVReader struct {
	record []string
}

func (r singleCSVReader) Read() ([]string, error) {
	return r.record, nil
}
func (r singleCSVReader) ReadAll() ([][]string, error) {
	return [][]string{r.record}, nil
}

func main() {
	s := `summary,3件
country,アメリカ合衆国,US/USA,310232863
country,日本,JP/JPN,127288000
country,中国,CN/CHN,1330044000`

	r := csv.NewReader(strings.NewReader(s)) // 標準パッケージでまず読み込む
	r.FieldsPerRecord = -1                   // レコード数が異なるのでチェックを無効化する
	all, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	for _, record := range all {
		if record[0] == "summary" { // カラムによって読み込み先を分岐
			var summaries []Summary
			if err := gocsv.UnmarshalCSVWithoutHeaders(singleCSVReader{record}, &summaries); err != nil {
				log.Fatal(err)
			}
			fmt.Printf("summary行の読み込み: %+v\n", summaries[0])
		} else {
			var countries []Country
			if err := gocsv.UnmarshalCSVWithoutHeaders(singleCSVReader{record}, &countries); err != nil {
				log.Fatal(err)
			}
			fmt.Printf("country行の読み込み: %+v\n", countries[0])
		}

	}
}

// multilayout-sample-end
