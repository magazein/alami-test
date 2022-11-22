package main

import (
	"fmt"
	"log"
	"strconv"
)

func updateLimitedBalance(i int, data []string, target [][]string) {
	// get balance
	blc, err := strconv.Atoi(data[BeforeBalance])
	if err != nil {
		log.Fatal("Invalid data type for balance "+data[BeforeBalance], err)
	}

	// get row number
	row, err := strconv.Atoi(data[BeforeID])
	if err != nil {
		log.Fatal("Invalid data type for id "+data[BeforeID], err)
	}

	if row <= 100 {
		target[row][AfterBalance] = fmt.Sprint(blc + 10)
	}

	target[row][AfterThread3] = fmt.Sprint(i)
}
