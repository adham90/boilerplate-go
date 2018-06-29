package database

import (
	"fmt"

	"database/sql"

	_ "github.com/lib/pq"
)

// Config database configration struct
type Config struct {
	Username, Password string
	Host               string
	Name               string
	Port               string
	LogMode            bool
}

func (c Config) connStr() string {
	return fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		c.Username,
		c.Password,
		c.Name,
		c.Host,
		c.Port,
	)
}

// Database object
type Database struct {
	DB *sql.DB
}

// New Database object
func New(c *Config) (*Database, error) {
	db, err := sql.Open("postgres", c.connStr())
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("pg: could not get a connection: %v", err)
	}

	return &Database{DB: db}, nil
}

// Migrate will update the database with new add col and tables
func (db *Database) Migrate() {
	// db.DB.AutoMigrate(entity.User{})
}

// Close the database connection
func (db *Database) Close() {
	db.DB.Close()
}
