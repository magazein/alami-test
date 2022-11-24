package usecase

import (
	"fmt"
	"strconv"

	"github.com/magazein/alami-test/constant"
)

func (uc *EndOfDayUC) UpdateBenefit(i int, data []string) error {
	// get balance
	blc, err := strconv.Atoi(data[constant.BeforeBalance])
	if err != nil {
		return err
	}

	// get id
	id, err := strconv.Atoi(data[constant.BeforeID])
	if err != nil {
		return err
	}

	// get current afterEod row data
	row, err := uc.afterEodRepo.Find(id)
	if err != nil {
		return err
	}

	if blc > 150 {
		row[constant.AfterBalance] = fmt.Sprint(blc + 25)
		row[constant.AfterThread2B] = fmt.Sprint(i)
		return uc.afterEodRepo.Update(id, row)
	}

	if blc >= 100 {
		row[constant.AfterFreeTrf] = "5"
		row[constant.AfterThread2A] = fmt.Sprint(i)
		return uc.afterEodRepo.Update(id, row)
	}

	return nil
}
