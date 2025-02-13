package my_csv

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

func CsvParse(fileName string) [][]string {

	csvTheFreshestFile, err := os.ReadFile(fileName)

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

func CsvSave(fileName string, data [][]string) error {
	fileForSaving, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("error while creating file: %v", err)
	}
	defer fileForSaving.Close()
	writer := csv.NewWriter(fileForSaving)
	writer.Comma = ';'
	writer.UseCRLF = true
	writer.WriteAll(data)
	writer.Flush()
	return nil
}
