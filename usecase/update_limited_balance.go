package usecase

import (
	"fmt"
	"strconv"

	"github.com/magazein/alami-test/constant"
)

func (uc *EndOfDayUC) UpdateLimitedBalance(i int, data []string) error {
	// get id number
	id, err := strconv.Atoi(data[constant.BeforeID])
	if err != nil {
		return err
	}

	// get current afterEod row data
	row, err := uc.afterEodRepo.Find(id)
	if err != nil {
		return err
	}

	// get updated balance
	blc, err := strconv.Atoi(row[constant.AfterBalance])
	if err != nil {
		return err
	}

	if id <= 100 {
		row[constant.AfterBalance] = fmt.Sprint(blc + 10)
	}

	row[constant.AfterThread3] = fmt.Sprint(i)

	return uc.afterEodRepo.Update(id, row)
}
