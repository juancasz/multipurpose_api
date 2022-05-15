package main

import (
	"log"
	"multipurpose_api/handler"
	"multipurpose_api/infrastructure/postgres"
	"multipurpose_api/infrastructure/server"
	"multipurpose_api/repository"
	"multipurpose_api/router"
	"multipurpose_api/service"
	"os"
)

func main() {
	db, err := postgres.Init(postgres.Options{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Name:     os.Getenv("DB_NAME"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		log.Fatal(err)
	}

	routes := router.New(handler.Handler{
		CalculatorHandler:  handler.NewCalculator(service.NewCalculator()),
		UserManagerHandler: handler.NewUserManager(service.NewUserManager(repository.NewUserManager(db))),
	})

	svr := server.New(&server.Options{
		Port:   os.Getenv("PORT"),
		Router: routes,
	})

	if err := svr.Start(); err != nil {
		log.Fatal(err)
	}
}
