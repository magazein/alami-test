package usecase

type EndOfDayUCItf interface {
	Proceed() error
	UpdateAvgBalance(i int, data []string) error
	UpdateBenefit(i int, data []string) error
	UpdateLimitedBalance(i int, data []string) error
}
