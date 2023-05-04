package migrations

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

var createTicketTableMigration = &Migration{
	Number: 3,
	Name:   "Create tickets table",
	Forwards: func(db *gorm.DB) error {
		const sql = `
		  CREATE TABLE tickets (
			id SERIAL PRIMARY KEY,
			user_id INTEGER NOT NULL,
			airline VARCHAR(255) NOT NULL,
			from_location VARCHAR(255) NOT NULL,
			destination VARCHAR(255) NOT NULL,
			promotion_id INTEGER,
			price FLOAT,
			status VARCHAR(255) NOT NULL,
			created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
			updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
			FOREIGN KEY (user_id) REFERENCES users(id),
			FOREIGN KEY (promotion_id) REFERENCES promotions(id)
			);
		  
		  INSERT INTO tickets (user_id, airline, from_location, destination, status) VALUES (1, 'Airasia', 'CNX', 'BKK', 'draft');
		  `

		err := db.Exec(sql).Error
		return errors.Wrap(err, "unable to create ticket table")
	},
}

func init() {
	Migrations = append(Migrations, createTicketTableMigration)
}
