package router

import (
	"multipurpose_api/handler"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func New(h handler.Handler) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Basic CORS
	// for more ideas, see: https://developer.github.com/v3/#cross-origin-resource-sharing
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Mount("/multipurpose-api", routesMultipurposeAPI(h))

	return r
}

func routesMultipurposeAPI(h handler.Handler) http.Handler {
	r := chi.NewRouter()

	r.Post("/order-array", h.CalculatorHandler.OrderArray)
	r.Post("/balance-months", h.CalculatorHandler.BalanceMonths)

	r.Post("/user", h.UserManagerHandler.AddUser)
	r.Get("/user/{user_id}", h.UserManagerHandler.GetUser)
	r.Put("/user/{user_id}", h.UserManagerHandler.EditUser)
	r.Delete("/user/{user_id}", h.UserManagerHandler.DeleteUser)

	return r
}
