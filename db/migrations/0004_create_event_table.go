package migrations

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

var createTicketEventTableMigration = &Migration{
	Number: 4,
	Name:   "Create ticketevents table",
	Forwards: func(db *gorm.DB) error {
		const sql = `
		  CREATE TABLE ticketevents (
			id SERIAL PRIMARY KEY,
			user_id INTEGER NOT NULL,
			ticket_id INTEGER NOT NULL,
			method VARCHAR(255) NOT NULL,
			status VARCHAR(255) NOT NULL,
			created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
			updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
			FOREIGN KEY (user_id) REFERENCES users(id),
			FOREIGN KEY (ticket_id) REFERENCES tickets(id)
			);
		  `

		err := db.Exec(sql).Error
		return errors.Wrap(err, "unable to create ticketevents table")
	},
}

func init() {
	Migrations = append(Migrations, createTicketEventTableMigration)
}
