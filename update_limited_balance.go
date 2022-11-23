package main

import (
	"fmt"
	"strconv"
)

func updateLimitedBalance(i int, data []string, target [][]string) ([][]string, error) {
	// get row number
	row, err := strconv.Atoi(data[BeforeID])
	if err != nil {
		return nil, err
	}

	// get updated balance
	blc, err := strconv.Atoi(target[row][AfterBalance])
	if err != nil {
		return nil, err
	}

	if row <= 100 {
		target[row][AfterBalance] = fmt.Sprint(blc + 10)
	}

	target[row][AfterThread3] = fmt.Sprint(i)

	return target, nil
}
