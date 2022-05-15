package service

import (
	"context"
	"encoding/base64"
	"fmt"
	"multipurpose_api/model"

	"github.com/google/uuid"
)

type userManagerRepository interface {
	AddUser(ctx context.Context, user *model.User) error
	GetUser(ctx context.Context, user *model.UserInfo) error
	EditUser(ctx context.Context, user *model.User) error
	DeleteUser(ctx context.Context, user *model.UserInfo) error
}

func NewUserManager(userManager userManagerRepository) *UserManager {
	if userManager == nil {
		panic("missing userManager while creating UserManager service")
	}

	return &UserManager{
		userManager: userManager,
	}
}

type UserManager struct {
	userManager userManagerRepository
}

func (u *UserManager) AddUser(ctx context.Context, user *model.User) (string, error) {
	user.Id = uuid.New().String()
	key := fmt.Sprintf("%s:%s", user.Username, user.Password)
	user.HashToken = base64.StdEncoding.EncodeToString([]byte(key))
	return user.Id, u.userManager.AddUser(ctx, user)
}

func (u *UserManager) GetUser(ctx context.Context, user *model.UserInfo) error {
	return u.userManager.GetUser(ctx, user)
}

func (u *UserManager) EditUser(ctx context.Context, user *model.User) error {
	return u.userManager.EditUser(ctx, user)
}

func (u *UserManager) DeleteUser(ctx context.Context, user *model.UserInfo) error {
	return u.userManager.DeleteUser(ctx, user)
}
