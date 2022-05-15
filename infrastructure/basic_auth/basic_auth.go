package basicauth

import (
	"context"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"multipurpose_api/model"
	"net/http"
	"strings"
)

type loginManager interface {
	GetUserCredentials(ctx context.Context, username string) (*model.UserCredentials, error)
}

type userID struct{}

func Middleware(manager loginManager) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			username, password, ok := r.BasicAuth()
			if ok {
				userCredentials, err := manager.GetUserCredentials(r.Context(), username)
				if err != nil {
					w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
					http.Error(w, "Unauthorized", http.StatusUnauthorized)
					return
				}

				tokenBytes, err := base64.StdEncoding.DecodeString(userCredentials.HashToken)
				if err != nil {
					w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
					http.Error(w, "Unauthorized", http.StatusUnauthorized)
					return
				}

				arrToken := strings.Split(string(tokenBytes), ":")
				if len(arrToken) != 2 {
					w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
					http.Error(w, "Unauthorized", http.StatusUnauthorized)
					return
				}

				usernameHash := sha256.Sum256([]byte(username))
				passwordHash := sha256.Sum256([]byte(password))
				expectedUsernameHash := sha256.Sum256([]byte(arrToken[0]))
				expectedPasswordHash := sha256.Sum256([]byte(arrToken[1]))

				usernameMatch := (subtle.ConstantTimeCompare(usernameHash[:], expectedUsernameHash[:]) == 1)
				passwordMatch := (subtle.ConstantTimeCompare(passwordHash[:], expectedPasswordHash[:]) == 1)

				if usernameMatch && passwordMatch {
					newCtx := context.WithValue(r.Context(), userID{}, userCredentials.Id)
					next.ServeHTTP(w, r.WithContext(newCtx))
					return
				}
			}

			w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
		})
	}
}

func UserIDFromRequest(r *http.Request) (string, error) {
	userID, ok := r.Context().Value(userID{}).(string)
	if !ok {
		return "", errors.New("invalid user_id in context request")
	}

	return userID, nil
}
