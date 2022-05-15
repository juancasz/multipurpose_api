package service

import "errors"

var (
	ErrInvalidInputBalanceMonths = errors.New("the provided input is not coherent")
	ErrUserNotFound              = errors.New("user not found")
)
