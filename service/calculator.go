package service

import (
	"context"
	"fmt"
	"multipurpose_api/model"
	"sort"
	"strings"
)

func NewCalculator() *Calculator {
	return &Calculator{}
}

type Calculator struct{}

func (c *Calculator) OrderArray(ctx context.Context, input []int) []int {

	output := make([]int, len(input))
	copy(output, input)

	//sort array
	sort.Ints(output)

	//move repeated elements to the end of the array
	read := 0
	write := 0

	for read < len(output) {

		// Swap the values pointed at by read and write.
		pointerWrite := output[write]
		output[write], output[read] = output[read], output[write]

		/*
			Advance the read pointer forward to the next unique value.  Since we
			moved the unique value to the write location, we compare values
			against input[write] instead of input[read].
		*/
		for read < len(output) && (output[read] == output[write] || output[read] == pointerWrite) {
			read++
		}

		write++
	}

	return output
}

var monthsSpanish map[string]bool = map[string]bool{
	"enero":      true,
	"febrero":    true,
	"marzo":      true,
	"abril":      true,
	"mayo":       true,
	"junio":      true,
	"julio":      true,
	"agosto":     true,
	"septiembre": true,
	"octubre":    true,
	"noviembre":  true,
	"diciembre":  true,
}

func (c *Calculator) BalanceMonths(ctx context.Context, input *model.InputBalanceMonths) ([]model.BalanceMonth, error) {
	if len(input.Months) != len(input.Sales) || len(input.Months) != len(input.Costs) {
		return nil, fmt.Errorf("%w : %d meses , %d ventas, %d gastos", ErrInvalidInputBalanceMonths, len(input.Months), len(input.Sales), len(input.Costs))
	}

	months := len(input.Months)

	balanceMonths := make([]model.BalanceMonth, 0)
	var balanceMonth model.BalanceMonth

	var validMonth bool

	for i := 0; i < months; i++ {
		_, validMonth = monthsSpanish[strings.ToLower(input.Months[i])]
		if !validMonth {
			return nil, fmt.Errorf("%w: mes %s no válido", ErrInvalidInputBalanceMonths, input.Months[i])
		}

		balanceMonth = model.BalanceMonth{
			Month:   input.Months[i],
			Sales:   input.Sales[i],
			Costs:   input.Costs[i],
			Balance: input.Sales[i] - input.Costs[i],
		}
		balanceMonths = append(balanceMonths, balanceMonth)
	}

	return balanceMonths, nil
}
