package main

import (
	"encoding/csv"
	"log"
	"os"
)

func readCsvFile(filePath string) [][]string {
	// open file
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to open file "+filePath, err)
	}
	defer f.Close()

	// create file reader
	reader := csv.NewReader(f)
	reader.Comma = Comma

	// read all file data
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal("Unable to read CSV in "+filePath, err)
	}

	return records
}

func writeCsvFile(name string, data [][]string) {
	// create file
	f, err := os.Create(name)
	if err != nil {
		return
	}
	defer f.Close()

	// create file writter
	w := csv.NewWriter(f)
	w.Comma = Comma

	// write data to file
	if err = w.WriteAll(data); err != nil {
		return
	}
}

func addToChan(channel chan []string, data [][]string) {
	// send each data to channel
	for i := 0; i < len(data); i++ {
		channel <- data[i]
	}

	// close channel
	close(channel)
}

func average(data ...int) int {
	var count int

	// count each element
	for _, x := range data {
		count += x
	}

	// get average
	return count / len(data)
}
