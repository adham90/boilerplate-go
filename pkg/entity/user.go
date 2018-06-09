package entity

import (
	"time"
)

//User data
type User struct {
	ID        uint   `gorm:"primary_key"`
	UUID      string `gorm:"index"`
	Name      string
	Email     string `gorm:"type:varchar(100);unique_index"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
