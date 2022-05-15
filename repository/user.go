package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
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

func (u *UserManager) GetUser(ctx context.Context, user *model.UserInfo) error {
	data, err := json.Marshal(user)
	if err != nil {
		return err
	}

	row := u.DB.QueryRowContext(ctx, callGetUser, string(data), nil)
	if err = row.Err(); err != nil {
		return err
	}

	str := new(sql.NullString)
	if err = row.Scan(str); err != nil {
		return err
	}

	if !str.Valid {
		return errors.New("error p_get_user NULL output")
	}

	if err = json.Unmarshal([]byte(str.String), user); err != nil {
		return err
	}

	return nil
}

func (u *UserManager) EditUser(ctx context.Context, user *model.User) error {
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

	row := tx.QueryRowContext(ctx, callEditUser, string(data))
	if err = row.Err(); err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (u *UserManager) DeleteUser(ctx context.Context, user *model.UserInfo) error {
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

	row := tx.QueryRowContext(ctx, callDeleteUser, string(data))
	if err = row.Err(); err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}
