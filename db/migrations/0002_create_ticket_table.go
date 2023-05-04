package migrations

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

var createTicketTableMigration = &Migration{
	Number: 2,
	Name:   "Create tickets table",
	Forwards: func(db *gorm.DB) error {
		const sql = `
		  CREATE TABLE tickets (
			  id SERIAL PRIMARY KEY,
			  owner_id INTEGER NOT NULL,
			  airline VARCHAR(255) NOT NULL,
			  from_location VARCHAR(255) NOT NULL,
			  destination VARCHAR(255) NOT NULL,
			  promotion VARCHAR(255),
			  status VARCHAR(255) NOT NULL,
			  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
			  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
				FOREIGN KEY (owner_id) REFERENCES users(id)
		  );
		  
		  INSERT INTO tickets (owner_id, airline, from_location, destination, status) VALUES (1, 'Airasia', 'CNX', 'BKK', 'draft');
		  `

		err := db.Exec(sql).Error
		return errors.Wrap(err, "unable to create ticket table")
	},
}

func init() {
	Migrations = append(Migrations, createTicketTableMigration)
}
