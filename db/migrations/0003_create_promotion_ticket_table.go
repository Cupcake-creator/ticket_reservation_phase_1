package migrations

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

var createPromotionTableMigration = &Migration{
	Number: 2,
	Name:   "Create promotion table",
	Forwards: func(db *gorm.DB) error {
		const sql = `
		  CREATE TABLE promotions (
			id SERIAL PRIMARY KEY,
			discount INTEGER DEFAULT 0,
			quota INTEGER ,
			airline VARCHAR(255) NOT NULL,
			status VARCHAR(255) NOT NULL,
			expire_at TIMESTAMPTZ,
			created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
			updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
			);
		  
		  INSERT INTO promotions (quota, airline, status) VALUES (10, 'Airasia', 'draft');
		  `

		err := db.Exec(sql).Error
		return errors.Wrap(err, "unable to create promotions table")
	},
}

func init() {
	Migrations = append(Migrations, createPromotionTableMigration)
}
