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
	DB                 *pg.DB
}

// Open create new Database connection
func (d *Database) Open() *pg.DB {
	d.DB = pg.Connect(&pg.Options{
		Database: d.Name,
		User:     d.Username,
		Password: d.Password,
	})

	var n int
	_, err := d.DB.QueryOne(pg.Scan(&n), "SELECT 1")
	if err != nil {
		return nil
	}

	return d.DB
}

// Migrate will update the database with new add col and tables
func (d *Database) Migrate(entities ...interface{}) error {
	if d.DB == nil {
		d.Open()
	}
	db := d.DB

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
