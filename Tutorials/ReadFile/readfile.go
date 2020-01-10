package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	//----------Reads all of file----------
	data, err := ioutil.ReadFile("test.txt")
	check(err)
	fmt.Println(data)
	fmt.Println(string(data))

	cummCsv := make([][]string, 0)
	files := []string{"AAL.csv", "AAPL.csv"}
	for _, filename := range files {
		data, _ = ioutil.ReadFile(filename)
		rec := csv.NewReader(strings.NewReader(string(data)))
		records, err := rec.ReadAll()
		if err == nil {
			cummCsv = append(cummCsv, records[1:]...)
		}
	}

	fmt.Println(cummCsv[0])
	fmt.Println(cummCsv[49][0]) //day, id
	fmt.Println(cummCsv[49][1]) //day, lstm
	fmt.Println(cummCsv[49][2]) //day, data
	fmt.Println(cummCsv[49][3]) //day, closing

	//----------Read specific parts of file----------
	//f, err := os.Open("test.txt")
	//check(err)
	//...
}
