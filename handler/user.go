package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"multipurpose_api/model"
	"multipurpose_api/service"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type userManagerService interface {
	AddUser(ctx context.Context, user *model.User) error
	GetUser(ctx context.Context, user *model.UserInfo) error
	EditUser(ctx context.Context, user *model.User) error
	DeleteUser(ctx context.Context, user *model.UserInfo) error
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

	pErr, ok := err.(*pq.Error)
	if ok {
		switch string(pErr.Code) {
		case P_ERR_VIOLATES_FOREIGN_KEY:
			w.Header().Add("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(Response{"error": err.Error()})
			return
		}
	}

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

func (u *UserManager) GetUser(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "user_id")

	_, err := uuid.Parse(userID)
	if err != nil {
		w.Header().Add("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{"error": `url param "user_id" is not an uuid`})
		return
	}

	user := model.UserInfo{
		Id: userID,
	}

	err = u.userManager.GetUser(r.Context(), &user)
	if errors.Is(err, service.ErrUserNotFound) {
		w.Header().Add("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusNotFound)
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
	json.NewEncoder(w).Encode(user)
}

func (u *UserManager) EditUser(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "user_id")

	_, err := uuid.Parse(userID)
	if err != nil {
		w.Header().Add("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{"error": `url param "user_id" is not an uuid`})
		return
	}

	var user model.User
	err = json.NewDecoder(r.Body).Decode(&user)

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

	user.Id = userID

	err = u.userManager.EditUser(r.Context(), &user)

	pErr, ok := err.(*pq.Error)
	if ok {
		log.Println("pErr code ---->", string(pErr.Code))
		switch string(pErr.Code) {
		case P_ERR_VIOLATES_FOREIGN_KEY:
			w.Header().Add("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(Response{"error": err.Error()})
			return
		}
	}

	if err != nil {
		w.Header().Add("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Response{"error": err.Error()})
		return
	}

	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response{"message": "success"})
}

func (u *UserManager) DeleteUser(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "user_id")

	_, err := uuid.Parse(userID)
	if err != nil {
		w.Header().Add("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{"error": `url param "user_id" is not an uuid`})
		return
	}

	err = u.userManager.DeleteUser(r.Context(), &model.UserInfo{
		Id: userID,
	})
	if err != nil {
		w.Header().Add("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Response{"error": err.Error()})
		return
	}

	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response{"message": "success"})
}
