package main

import (
	"fmt"
	"strconv"
)

func updateBenefit(i int, data []string, target [][]string) ([][]string, error) {
	// get balance
	blc, err := strconv.Atoi(data[BeforeBalance])
	if err != nil {
		return nil, err
	}

	// get row number
	row, err := strconv.Atoi(data[BeforeID])
	if err != nil {
		return nil, err
	}

	if blc > 150 {
		target[row][AfterBalance] = fmt.Sprint(blc + 25)
		target[row][AfterThread2B] = fmt.Sprint(i)
		return target, nil
	}

	if blc >= 100 {
		target[row][AfterFreeTrf] = "5"
		target[row][AfterThread2A] = fmt.Sprint(i)
	}

	return target, nil
}
