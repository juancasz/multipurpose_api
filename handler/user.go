package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"multipurpose_api/model"
	"net/http"
)

type userManagerService interface {
	AddUser(ctx context.Context, user *model.User) error
}

func NewUserManager(
	userManager userManagerService,
) *UserManager {
	if userManager == nil {
		panic("missing userManager while creating UserManager handler")
	}

	return &UserManager{
		userManager: userManager,
	}
}

type UserManager struct {
	userManager userManagerService
}

func (u *UserManager) AddUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.Header().Add("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{"error": fmt.Sprintf("error parsing input data: %s", err.Error())})
		return
	}

	if user.Name == "" {
		w.Header().Add("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{"error": `required field "name" is missing or empty`})
		return
	}

	if user.CountryID == 0 {
		w.Header().Add("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{"error": `required field "country_id" is missing or empty`})
		return
	}

	if user.UniversityID == 0 {
		w.Header().Add("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{"error": `required field "university_id" is missing or empty`})
		return
	}

	if user.Username == "" {
		w.Header().Add("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{"error": `required field "username" is missing or empty`})
		return
	}

	if user.HashPassword == "" {
		w.Header().Add("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{"error": `required field "hash_password" is missing or empty`})
		return
	}

	err = u.userManager.AddUser(r.Context(), &user)
	if err != nil {
		w.Header().Add("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Response{"error": err.Error()})
		return
	}

	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(Response{"message": "success"})
}

func (u *UserManager) GetUser(w http.ResponseWriter, r *http.Request) {}

func (u *UserManager) EditUser(w http.ResponseWriter, r *http.Request) {}

func (u *UserManager) DeleteUser(w http.ResponseWriter, r *http.Request) {}
