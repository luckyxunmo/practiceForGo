package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main() {
	//writeCsv()
	openCsv()
	//fmt.Println(len("\xEF\xBB\xBF"))

}
func writeCsv() {
	f, err := os.Create("test2.csv")
	if err != nil {
		panic(err)
	}

	defer f.Close()
	f.WriteString("\xEF\xBB\xBF")
	w := csv.NewWriter(f)
	data := [][]string{
		{"1", "中国", "23"},
		{"2", "美国", "23"},
		{"3", "bb", "23"},
		{"4", "bb", "23"},
		{"5", "bb", "23"},
	}
	w.WriteAll(data)
	w.Flush()
}
func openCsv() {
	file, err := os.Open("test2.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	reader.Comma = ','
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {

			fmt.Println("Error:", err)
			break
		}
		fmt.Println("record is ", record[0])
	}
}
