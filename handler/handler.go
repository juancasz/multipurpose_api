package handler

import "net/http"

type CalculatorHandler interface {
	OrderArray(w http.ResponseWriter, r *http.Request)
	BalanceMonths(w http.ResponseWriter, r *http.Request)
}

type Handler struct {
	CalculatorHandler
}

type Response map[string]interface{}
