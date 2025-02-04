package my_csv

import (
	"encoding/csv"
	"io/ioutil"
	"log"
	"strings"
)

func CsvParse(fileName string) [][]string {
	csvTheFreshestFile, err := ioutil.ReadFile(fileName)

	if err != nil {
		log.Fatal(err)
	}

	preData := csv.NewReader(strings.NewReader(string(csvTheFreshestFile)))
	preData.Comma = ';'
	preData.Comment = '#'

	theFreshestData, err := preData.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	return theFreshestData
}
