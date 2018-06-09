package database

import (
	"fmt"

	"github.com/adham90/boilerplate/pkg/entity"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// PGConfig database config
type Config struct {
	// Optional.
	Username, Password string

	// Host of the MySQL instance.
	//
	// If set, UnixSocket should be unset.
	Host string

	// DatabaseName as string
	DatabaseName string

	// Port of the MySQL instance.
	//
	// If set, UnixSocket should be unset.
	Port string

	LogMode bool
}

func (c Config) connString() string {
	return fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", c.Username, c.Password, c.DatabaseName, c.Host, c.Port)
}

// Database object
type Database struct {
	Conn *gorm.DB
}

// New Database object
func New(c Config) (*Database, error) {
	conn, err := gorm.Open("postgres", c.connString())
	conn.LogMode(c.LogMode)
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("pg: could not get a connection: %v", err)
	}

	db := &Database{
		Conn: conn,
	}

	return db, nil
}

// Migrate will update the database with new add col and tables
func (db *Database) Migrate() {
	db.Conn.AutoMigrate(entity.User{})
}

// Close the database connection
func (db *Database) Close() {
	db.Conn.Close()
}
