package handler

import "net/http"

type CalculatorHandler interface {
	OrderArray(w http.ResponseWriter, r *http.Request)
	BalanceMonths(w http.ResponseWriter, r *http.Request)
}

type UserManagerHandler interface {
	AddUser(w http.ResponseWriter, r *http.Request)
	GetUser(w http.ResponseWriter, r *http.Request)
	EditUser(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
}

type Handler struct {
	CalculatorHandler
	UserManagerHandler
}

type Response map[string]interface{}

const (
	P_ERR_VIOLATES_FOREIGN_KEY = "23503"
)
