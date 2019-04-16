package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// CreateConnection ...
func CreateConnection() (*gorm.DB, error) {
	return gorm.Open("sqlite3", "/tmp/user.db")
}
