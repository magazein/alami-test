package main

import (
	"encoding/csv"
	"os"
)

func readCsvFile(filePath string) ([][]string, error) {
	// open file
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// create file reader
	reader := csv.NewReader(f)
	reader.Comma = Comma

	// read all file data
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}

func writeCsvFile(name string, data [][]string) error {
	// create file
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()

	// create file writter
	w := csv.NewWriter(f)
	w.Comma = Comma

	// write data to file
	if err = w.WriteAll(data); err != nil {
		return err
	}

	return nil
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
