package db

import (
	"github.com/asdine/storm/v3"
)

var DB *storm.DB

func InitDB() error {
	var err error
	DB, err = storm.Open("database.db") // Persistent DB
	return err
}

func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}
