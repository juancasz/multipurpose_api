package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"multipurpose_api/model"
	"multipurpose_api/service"
	"net/http"

	"github.com/lib/pq"
)

type loginManagerService interface {
	Login(ctx context.Context, user *model.UserLogin) (*model.UserCredentials, error)
}

func NewLoginManager(login loginManagerService) *LoginManager {
	if login == nil {
		panic("missing loginManager while creating LoginManager handler")

	}

	return &LoginManager{
		login: login,
	}
}

type LoginManager struct {
	login loginManagerService
}

func (l *LoginManager) Login(w http.ResponseWriter, r *http.Request) {
	var user model.UserLogin
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.Header().Add("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{"error": fmt.Sprintf("error parsing input data: %s", err.Error())})
		return
	}

	if user.Username == "" {
		w.Header().Add("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{"error": `required field "username" is missing or empty`})
		return
	}

	if user.Password == "" {
		w.Header().Add("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{"error": `required field "password" is missing or empty`})
		return
	}

	userCredential, err := l.login.Login(r.Context(), &user)

	pErr, ok := err.(*pq.Error)
	if ok {
		switch string(pErr.Code) {
		case P_ERR_DATA_NOT_FOUND:
			w.Header().Add("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(Response{"error": err.Error()})
			return
		}
	}

	if errors.Is(err, service.ErrInvalidUsernameOrPassword) {
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
	json.NewEncoder(w).Encode(*userCredential)

}
