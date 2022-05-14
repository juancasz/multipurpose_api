package main

import (
	"log"
	"multipurpose_api/handler"
	"multipurpose_api/router"
	"multipurpose_api/server"
	"multipurpose_api/service"
	"os"
)

func main() {

	calculatorHandler := handler.NewCalculator(service.NewCalculator())

	routes := router.New(handler.Handler{
		CalculatorHandler: calculatorHandler,
	})

	svr := server.New(&server.Options{
		Port:   os.Getenv("PORT"),
		Router: routes,
	})

	if err := svr.Start(); err != nil {
		log.Fatal(err)
	}
}
