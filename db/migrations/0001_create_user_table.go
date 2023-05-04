package migrations

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

var createUserTableMigration = &Migration{
	Number: 1,
	Name:   "Create users table",
	Forwards: func(db *gorm.DB) error {
		const sql = `
		CREATE TABLE users (
			id SERIAL PRIMARY KEY,
			username VARCHAR(255) NOT NULL,
			password VARCHAR(255) NOT NULL,
			email VARCHAR(255) NOT NULL,
			role VARCHAR(255) NOT NULL DEFAULT 'admin',
			active BOOLEAN NOT NULL DEFAULT TRUE,
			created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
			updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
		  );
		  
		  INSERT INTO users (username, password, email) VALUES ('Jezy01', '1234', 'jew1@gmail.com');
		  INSERT INTO users (username, password, email) VALUES ('Jezy02', '1234', 'jew2@gmail.com');
		`

		err := db.Exec(sql).Error
		return errors.Wrap(err, "unable to create users table")
	},
}

func init() {
	Migrations = append(Migrations, createUserTableMigration)
}
