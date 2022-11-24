package usecase

import (
	"log"
	"runtime"

	"github.com/magazein/alami-test/constant"
	"github.com/magazein/alami-test/helper"
	"github.com/magazein/alami-test/repository"
)

type EndOfDayUC struct {
	beforeEodRepo repository.BeforeEodRepoItf
	afterEodRepo  repository.AfterEodRepoItf
	self          EndOfDayUCItf
}

func NewEndOfDayUC(
	beforeEodRepo repository.BeforeEodRepoItf,
	afterEodRepo repository.AfterEodRepoItf,
) *EndOfDayUC {
	uc := &EndOfDayUC{
		beforeEodRepo: beforeEodRepo,
		afterEodRepo:  afterEodRepo,
	}

	uc.self = uc

	return uc
}

func (uc *EndOfDayUC) Proceed() error {
	// get before eod data
	beforeEod, err := uc.beforeEodRepo.Get()
	if err != nil {
		return err
	}

	// validate length of data
	if len(beforeEod) < 2 {
		return constant.ErrInvalidData
	}

	// construct AfterEod basic data
	afterEod := uc.constructAfterEod(beforeEod)
	uc.afterEodRepo.Replace(afterEod)

	// skip header
	beforeEod = beforeEod[1:]

	// define number of concurrency
	macProcs := runtime.GOMAXPROCS(4)

	// question 1: update average balance
	errs := helper.ProceedAsync(macProcs, uc.self.UpdateAvgBalance, beforeEod)
	if len(errs) > 0 {
		log.Println(errs)
		return constant.ErrUpdateAvgBalance
	}

	// question 2: update customer benefit
	errs = helper.ProceedAsync(macProcs, uc.self.UpdateBenefit, beforeEod)
	if len(errs) > 0 {
		log.Println(errs)
		return constant.ErrUpdateBenefit
	}

	// question 3: add balance for limited customer
	errs = helper.ProceedAsync(8, uc.self.UpdateLimitedBalance, beforeEod)
	if len(errs) > 0 {
		log.Println(errs)
		return constant.ErrUpdateLimitedBalance
	}

	return uc.afterEodRepo.WriteCSV()
}

func (uc *EndOfDayUC) constructAfterEod(beforeEod [][]string) [][]string {
	rows := make([][]string, len(beforeEod))

	header := []string{
		"id",
		"Nama",
		"Age",
		"Balanced",
		"No 2b Thread-No",
		"No 3 Thread-No",
		"Previous Balanced",
		"Average Balanced",
		"No 1 Thread-No",
		"Free Transfer",
		"No 2a Thread-No",
	}

	rows[0] = header

	for i := 1; i < len(beforeEod); i++ {
		row := make([]string, 11)
		data := beforeEod[i]

		row[constant.AfterID] = data[constant.BeforeID]
		row[constant.AfterName] = data[constant.BeforeName]
		row[constant.AfterAge] = data[constant.BeforeAge]
		row[constant.AfterBalance] = data[constant.BeforeBalance]
		row[constant.AfterPrevBalance] = data[constant.BeforePrevBalance]
		row[constant.AfterAvgBalance] = data[constant.BeforeAvgBalance]
		row[constant.AfterFreeTrf] = data[constant.BeforeFreeTrf]

		rows[i] = row
	}

	return rows
}
