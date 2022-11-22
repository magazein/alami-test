package main

import (
	"runtime"
)

func main() {
	// read BeforeEod csv
	beforeEod := readCsvFile(BeforeEodCsvName)

	// read AfterEod csv
	afterEod := readCsvFile(AfterEodCsvName)

	// compare length of each csv
	if len(beforeEod) != len(afterEod) {
		return
	}

	// skip header
	beforeEod = beforeEod[1:]

	// define number of concurrency
	macProcs := runtime.GOMAXPROCS(4)

	// question 1: update average balance
	proceedAsync(macProcs, updateAvgBalance, beforeEod, afterEod)

	// question 2: update customer benefit
	proceedAsync(macProcs, updateBenefit, beforeEod, afterEod)

	// question 3: add balance for limited customer
	proceedAsync(8, updateLimitedBalance, beforeEod, afterEod)

	// write to AfterEod csv
	writeCsvFile(AfterEodCsvName, afterEod)
}
