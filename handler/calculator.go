package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"multipurpose_api/model"
	"multipurpose_api/service"
	"net/http"
)

type calculatorService interface {
	OrderArray(ctx context.Context, input []int) []int
	BalanceMonths(ctx context.Context, input *model.InputBalanceMonths) ([]model.BalanceMonth, error)
}

func NewCalculator(calculator calculatorService) *Calculator {
	if calculator == nil {
		panic("missing calculator while creating Calculator")
	}

	return &Calculator{
		calculator: calculator,
	}
}

type Calculator struct {
	calculator calculatorService
}

func (c *Calculator) OrderArray(w http.ResponseWriter, r *http.Request) {
	var input model.InputOrderArray
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.Header().Add("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{"error": fmt.Sprintf("error parsing input data: %s", err.Error())})
		return
	}

	if len(input.Array) == 0 {
		w.Header().Add("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{"error": `missing required field "sin_clasificar"`})
		return
	}

	outputArray := c.calculator.OrderArray(r.Context(), input.Array)

	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response{
		"sin_clasificar": input.Array,
		"clasificado":    outputArray,
	})
}

func (c *Calculator) BalanceMonths(w http.ResponseWriter, r *http.Request) {
	var input model.InputBalanceMonths
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.Header().Add("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{"error": fmt.Sprintf("error parsing input data: %s", err.Error())})
		return
	}

	if len(input.Months) == 0 {
		w.Header().Add("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{"error": `missing required field "Mes"`})
		return
	}

	if len(input.Sales) == 0 {
		w.Header().Add("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{"error": `missing required field "Gastos"`})
		return
	}

	if len(input.Costs) == 0 {
		w.Header().Add("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{"error": `missing required field "Costos"`})
		return
	}

	balance, err := c.calculator.BalanceMonths(r.Context(), &input)

	if errors.Is(err, service.ErrInvalidInputBalanceMonths) {
		w.Header().Add("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{"error": err.Error()})
		return
	}

	if err != nil {
		w.Header().Add("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Response{"error": err.Error()})
		return
	}

	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response{
		"balance": balance,
	})
}
