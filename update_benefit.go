package main

import (
	"fmt"
	"log"
	"strconv"
)

func updateBenefit(i int, data []string, target [][]string) {
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

	if blc > 150 {
		target[row][AfterBalance] = fmt.Sprint(blc + 25)
		target[row][AfterThread2B] = fmt.Sprint(i)
		return
	}

	if blc >= 100 {
		target[row][AfterFreeTrf] = "5"
		target[row][AfterThread2A] = fmt.Sprint(i)
	}
}
