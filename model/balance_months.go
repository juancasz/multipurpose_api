package model

type (
	InputBalanceMonths struct {
		Months []string `json:"Mes"`
		Sales  []int    `json:"Ventas"`
		Costs  []int    `json:"Gastos"`
	}

	BalanceMonth struct {
		Month   string `json:"Mes"`
		Sales   int    `json:"Ventas"`
		Costs   int    `json:"Gastos"`
		Balance int    `json:"Balance"`
	}
)
