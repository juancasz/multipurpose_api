package router

import (
	"context"
	"multipurpose_api/handler"
	basicauth "multipurpose_api/infrastructure/basic_auth"
	"multipurpose_api/model"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type userCredentialManager interface {
	GetUserCredentials(ctx context.Context, username string) (*model.UserCredentials, error)
}

func New(h handler.Handler, managerAuth userCredentialManager) http.Handler {
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

	r.Mount("/multipurpose-api", routesMultipurposeAPI(h, managerAuth))

	return r
}

func routesMultipurposeAPI(h handler.Handler, managerAuth userCredentialManager) http.Handler {
	r := chi.NewRouter()

	r.Mount("/login", routesLogin(h.LoginManagerHandler))
	r.Mount("/calculator", routesCalculator(h.CalculatorHandler, managerAuth))
	r.Mount("/user", routesUser(h.UserManagerHandler, managerAuth))

	return r
}

func routesLogin(login handler.LoginManagerHandler) http.Handler {
	r := chi.NewRouter()

	r.Post("/", login.Login)

	return r
}

func routesCalculator(calculator handler.CalculatorHandler, managerAuth userCredentialManager) http.Handler {
	r := chi.NewRouter()

	//Basic auth
	r.Use(basicauth.Middleware(managerAuth))

	r.Post("/order-array", calculator.OrderArray)
	r.Post("/balance-months", calculator.BalanceMonths)

	return r
}

func routesUser(userManager handler.UserManagerHandler, managerAuth userCredentialManager) http.Handler {
	r := chi.NewRouter()

	//Basic auth
	r.Use(basicauth.Middleware(managerAuth))

	r.Post("/", userManager.AddUser)
	r.Get("/{user_id}", userManager.GetUser)
	r.Put("/{user_id}", userManager.EditUser)
	r.Delete("/{user_id}", userManager.DeleteUser)

	return r
}
