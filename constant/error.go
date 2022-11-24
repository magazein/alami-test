package constant

import "errors"

var (
	ErrOutOfRange           = errors.New("out of range")
	ErrInvalidData          = errors.New("invalid data")
	ErrUpdateAvgBalance     = errors.New("error update avg balance")
	ErrUpdateBenefit        = errors.New("error update benefit")
	ErrUpdateLimitedBalance = errors.New("error update limited balance")
)
