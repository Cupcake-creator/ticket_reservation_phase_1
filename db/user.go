package db

import (
	"context"

	"github.com/pkg/errors"
)

type DBUserInterface interface {
	CreateUser(username, password, email string) (int64, error)
	IsExistUser(username string) (bool, error)
}

func (pgdb *PostgresqlDB) CreateUser(username, password, email string) (int64, error) {
	var userID int64

	err := pgdb.DB.QueryRow(context.Background(), `
		INSERT INTO users (
			"username",
			"password",
			"email"
		)
		VALUES ($1, $2, $3)
		RETURNING id
	`, username, password, email).Scan(&userID)
	if err != nil {
		return 0, errors.Wrap(err, "Unable to create user")
	}

	return userID, nil
}

func (pgdb *PostgresqlDB) IsExistUser(username string) (bool, error) {

	var count int32

	err := pgdb.DB.QueryRow(context.Background(), `
		SELECT count(id) FROM users 
		WHERE username = $1
	`, username).Scan(&count)
	if err != nil {
		return true, errors.Wrap(err, "Unable to create user")
	}

	if count > 0 {
		return true, errors.New("User already exists")
	}

	return false, nil
}
