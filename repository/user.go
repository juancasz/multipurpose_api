package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"multipurpose_api/model"
)

func NewUserManager(db *sql.DB) *UserManager {
	if db == nil {
		panic("missing db while creating UserManager repository")
	}

	return &UserManager{
		DB: db,
	}
}

type UserManager struct {
	*sql.DB
}

func (u *UserManager) AddUser(ctx context.Context, user *model.User) error {
	data, err := json.Marshal(user)
	if err != nil {
		return err
	}

	tx, err := u.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	row := tx.QueryRowContext(ctx, callAddUser, string(data))
	if err = row.Err(); err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}
