package usecase

import (
	"fmt"
	"strconv"

	"github.com/magazein/alami-test/constant"
	"github.com/magazein/alami-test/helper"
)

func (uc *EndOfDayUC) UpdateAvgBalance(i int, data []string) error {
	// get balance
	blc, err := strconv.Atoi(data[constant.BeforeBalance])
	if err != nil {
		return err
	}

	// get previous balance
	prevBlc, err := strconv.Atoi(data[constant.BeforePrevBalance])
	if err != nil {
		return err
	}

	// get id
	id, err := strconv.Atoi(data[constant.BeforeID])
	if err != nil {
		return err
	}

	// count average
	avg := helper.Average(blc, prevBlc)

	// get current afterEod row data
	row, err := uc.afterEodRepo.Find(id)
	if err != nil {
		return err
	}

	// update data
	row[constant.AfterAvgBalance] = fmt.Sprint(avg)
	row[constant.AfterThread1] = fmt.Sprint(i)

	// execute update
	return uc.afterEodRepo.Update(id, row)
}
