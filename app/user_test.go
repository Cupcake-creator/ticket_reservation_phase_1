package app

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"ticket-reservation/db"
)

type mockDB struct {
	db.DB
	createUser func(username, password, email string) (int64, error)
}

func (mockdb *mockDB) CreateUser(username, password, email string) (int64, error) {
	return mockdb.createUser(username, password, email)
}

func TestRegister(t *testing.T) {
	logger, err := getLoggerForTesting()
	if err != nil {
		t.Fatalf("error get logger: %+v", err)
	}

	userID := int64(1)

	appContext := &Context{
		Logger: logger,

		DB: &mockDB{
			createUser: func(username, password, email string) (int64, error) {
				return userID, nil
			},
		},
	}

	registerResult, err := appContext.Register(RegisterParams{
		Username: "test_user",
	})

	assert.Nil(t, err)
	assert.Equal(t, userID, registerResult.ID)
}
