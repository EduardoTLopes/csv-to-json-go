package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func CSVtoJSON() {
	csvFile, _ := os.Open("file.csv")
	reader := csv.NewReader(csvFile)

	var data []map[string]string
	headers, _ := reader.Read()

	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			fmt.Println(error)
		}

		m := make(map[string]string)
		for i, value := range line {
			m[headers[i]] = value
		}
		data = append(data, m)
	}

	jsonData, _ := json.Marshal(data)
	jsonFile, _ := os.Create("file.json")
	jsonFile.Write(jsonData)
	jsonFile.Close()
}

func main() {
	CSVtoJSON()
}
