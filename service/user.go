package service

import (
	"context"
	"multipurpose_api/model"

	"github.com/google/uuid"
)

type userManagerRepository interface {
	AddUser(ctx context.Context, user *model.User) error
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

func (u *UserManager) AddUser(ctx context.Context, user *model.User) error {
	user.Id = uuid.New().String()
	user.UserIDCreator = "00000000-0000-0000-0000-000000000000"
	return u.userManager.AddUser(ctx, user)
}
