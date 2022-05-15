package service

import (
	"context"
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

func (u *UserManager) AddUser(ctx context.Context, user *model.User) error {
	user.Id = uuid.New().String()
	user.UserIDCreator = "00000000-0000-0000-0000-000000000000"
	return u.userManager.AddUser(ctx, user)
}

func (u *UserManager) GetUser(ctx context.Context, user *model.UserInfo) error {
	id := user.Id
	if err := u.userManager.GetUser(ctx, user); err != nil {
		return err
	}

	if user.Id == "" {
		return fmt.Errorf("%w: user with id %s not found", ErrUserNotFound, id)
	}

	return nil
}

func (u *UserManager) EditUser(ctx context.Context, user *model.User) error {
	return u.userManager.EditUser(ctx, user)
}

func (u *UserManager) DeleteUser(ctx context.Context, user *model.UserInfo) error {
	return u.userManager.DeleteUser(ctx, user)
}
