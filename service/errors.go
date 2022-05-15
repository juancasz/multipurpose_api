package service

import "errors"

var (
	ErrInvalidInputBalanceMonths = errors.New("the provided input is not coherent")
	ErrInvalidUsernameOrPassword = errors.New("invalid username or password")
)
