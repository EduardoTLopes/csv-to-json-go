package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Data struct {
	Header1 string
	Header2 string
	Header3 string
	Header4 string
	Header5 string
}

func main() {
	csvFile, _ := os.Open("file.csv")
	reader := csv.NewReader(csvFile)

	var data []Data
	headers, _ := reader.Read()

	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			fmt.Println(error)
		}

		var d Data
		fmt.Println(line)
		for i, value := range line {
			switch headers[i] {
			case "Header1":
				d.Header1 = value
			case "Header2":
				d.Header2 = value
			case "Header3":
				d.Header3 = value
			case "Header4":
				d.Header4 = value
			case "Header5":
				d.Header5 = value
			}
		}
		data = append(data, d)
	}

	jsonData, _ := json.Marshal(data)
	jsonFile, _ := os.Create("file.json")
	jsonFile.Write(jsonData)
	jsonFile.Close()
}
