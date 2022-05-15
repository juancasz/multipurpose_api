package service

import (
	"context"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/base64"
	"multipurpose_api/model"
	"strings"
)

type loginManagerRepository interface {
	GetUserCredentials(ctx context.Context, username string) (*model.UserCredentials, error)
}

func NewLoginManager(repository loginManagerRepository) *LoginManager {
	if repository == nil {
		panic("missing repository while creating LoginManager service")

	}

	return &LoginManager{
		repository: repository,
	}
}

type LoginManager struct {
	repository loginManagerRepository
}

func (l *LoginManager) Login(ctx context.Context, user *model.UserLogin) (*model.UserCredentials, error) {
	userCredentials, err := l.repository.GetUserCredentials(ctx, user.Username)
	if err != nil {
		return nil, err
	}

	tokenBytes, err := base64.StdEncoding.DecodeString(userCredentials.HashToken)
	if err != nil {
		return nil, ErrInvalidUsernameOrPassword
	}

	arrToken := strings.Split(string(tokenBytes), ":")
	if len(arrToken) != 2 {
		return nil, ErrInvalidUsernameOrPassword
	}

	usernameHash := sha256.Sum256([]byte(user.Username))
	passwordHash := sha256.Sum256([]byte(user.Password))
	expectedUsernameHash := sha256.Sum256([]byte(arrToken[0]))
	expectedPasswordHash := sha256.Sum256([]byte(arrToken[1]))

	usernameMatch := (subtle.ConstantTimeCompare(usernameHash[:], expectedUsernameHash[:]) == 1)
	passwordMatch := (subtle.ConstantTimeCompare(passwordHash[:], expectedPasswordHash[:]) == 1)

	if usernameMatch && passwordMatch {
		return userCredentials, nil
	}

	return nil, ErrInvalidUsernameOrPassword

}
