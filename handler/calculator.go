package handler

import (
	"encoding/json"
	"fmt"
	"multipurpose_api/model"
	"net/http"
)

type calculatorService interface {
	OrderArray(input []int) []int
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

	outputArray := c.calculator.OrderArray(input.Array)

	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response{
		"sin_clasificar": input.Array,
		"clasificado":    outputArray,
	})
}
