package main

import (
	"log"
	"runtime"
)

func main() {
	// read BeforeEod csv
	beforeEod, err := readCsvFile(BeforeEodCsvName)
	if err != nil {
		log.Println(err)
		return
	}

	// read AfterEod csv
	afterEod, err := readCsvFile(AfterEodCsvName)
	if err != nil {
		log.Println(err)
		return
	}

	// compare length of each csv
	if len(beforeEod) != len(afterEod) {
		return
	}

	// validate length of row
	if len(beforeEod) < 2 {
		return
	}

	// skip header
	beforeEod = beforeEod[1:]

	// define number of concurrency
	macProcs := runtime.GOMAXPROCS(4)

	// question 1: update average balance
	errs := proceedAsync(macProcs, updateAvgBalance, beforeEod, afterEod)
	if len(errs) > 0 {
		log.Println(errs)
		return
	}

	// question 2: update customer benefit
	errs = proceedAsync(macProcs, updateBenefit, beforeEod, afterEod)
	if len(errs) > 0 {
		log.Println(errs)
		return
	}

	// question 3: add balance for limited customer
	errs = proceedAsync(8, updateLimitedBalance, beforeEod, afterEod)
	if len(errs) > 0 {
		log.Println(errs)
		return
	}

	// write to AfterEod csv
	err = writeCsvFile(AfterEodCsvName, afterEod)
	if err != nil {
		log.Println(err)
	}
}
