package main

import (
	"fmt"
	"strconv"
)

func updateAvgBalance(i int, data []string, target [][]string) ([][]string, error) {
	// get balance
	blc, err := strconv.Atoi(data[BeforeBalance])
	if err != nil {
		return nil, err
	}

	// get previous balance
	prevBlc, err := strconv.Atoi(data[BeforePrevBalance])
	if err != nil {
		return nil, err
	}

	// get row number
	row, err := strconv.Atoi(data[BeforeID])
	if err != nil {
		return nil, err
	}

	// count average
	avg := average(blc, prevBlc)

	// update data
	target[row][AfterAvgBalance] = fmt.Sprint(avg)
	target[row][AfterThread1] = fmt.Sprint(i)

	return target, nil
}
