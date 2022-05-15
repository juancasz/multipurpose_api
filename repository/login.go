package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"multipurpose_api/model"
)

func NewLoginManager(db *sql.DB) *LoginManager {
	if db == nil {
		panic("missing db while creating UserManager repository")
	}

	return &LoginManager{
		DB: db,
	}
}

type LoginManager struct {
	*sql.DB
}

func (l *LoginManager) GetUserCredentials(ctx context.Context, username string) (*model.UserCredentials, error) {
	row := l.DB.QueryRowContext(ctx, getUserCredentials, username)
	if err := row.Err(); err != nil {
		return nil, err
	}

	str := new(sql.NullString)
	if err := row.Scan(str); err != nil {
		return nil, err
	}

	if !str.Valid {
		return nil, errors.New("error f_get_user_credentials NULL output")
	}

	var user model.UserCredentials
	if err := json.Unmarshal([]byte(str.String), &user); err != nil {
		return nil, err
	}

	return &user, nil
}
