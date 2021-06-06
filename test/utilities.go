package test

import (
	"encoding/csv"
	"os"
	"strings"

	model "github.com/phuockhanhle/familytree/model"
)

func ParseCSVtoListPeople(filePath string) []model.Person {
	var res []model.Person

	lines, _ := CSVFileToMap(filePath)
	for _, line := range lines {
		tmp := model.MapToStruct(line, model.Person{}).(model.Person)
		res = append(res, tmp)
	}

	return res
}

func CSVFileToMap(filePath string) ([]map[string]interface{}, error) {

	var returnMap []map[string]interface{}
	// read csv file
	csvfile, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	defer csvfile.Close()

	reader := csv.NewReader(csvfile)

	rawCSVdata, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	header := []string{} // holds first row (header)
	for lineNum, record := range rawCSVdata {

		// for first row, build the header slice
		if lineNum == 0 {
			for i := 0; i < len(record); i++ {
				header = append(header, strings.TrimSpace(record[i]))
			}
		} else {
			// for each cell, map[string]string k=header v=value
			line := map[string]interface{}{}
			for i := 0; i < len(record); i++ {
				line[header[i]] = record[i]
			}
			returnMap = append(returnMap, line)
		}
	}

	return returnMap, nil
}
