package main

import (
	"fmt"
	"log"
	"strconv"
)

func updateAvgBalance(i int, data []string, target [][]string) {
	// get balance
	blc, err := strconv.Atoi(data[BeforeBalance])
	if err != nil {
		log.Fatal("Invalid data type for balance "+data[BeforeBalance], err)
	}

	// get previous balance
	prevBlc, err := strconv.Atoi(data[BeforePrevBalance])
	if err != nil {
		log.Fatal("Invalid data type for previous balance "+data[BeforePrevBalance], err)
	}

	// get row number
	row, err := strconv.Atoi(data[BeforeID])
	if err != nil {
		log.Fatal("Invalid data type for id "+data[BeforeID], err)
	}

	// count average
	avg := average(blc, prevBlc)

	// update data
	target[row][AfterAvgBalance] = fmt.Sprint(avg)
	target[row][AfterThread1] = fmt.Sprint(i)
}
