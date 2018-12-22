package utils

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

// Database object
type Database struct {
	Username, Password string
	Host               string
	Name               string
	Port               string
	LogMode            bool
}

// Open create new Database connection
func (d *Database) Open() *pg.DB {
	db := pg.Connect(&pg.Options{
		Database: d.Name,
		User:     d.Username,
		Password: d.Password,
	})

	var n int
	_, err := db.QueryOne(pg.Scan(&n), "SELECT 1")
	if err != nil {
		return nil
	}

	return db
}

func (d *Database) Migrate(entities ...interface{}) error {
	db := d.Open()

	for _, entity := range entities {
		err := db.CreateTable(entity, &orm.CreateTableOptions{
			Temp:        false,
			IfNotExists: true,
		})
		if err != nil {
			return err
		}
	}

	return nil
}
