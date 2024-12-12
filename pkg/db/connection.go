package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // For PostgreSQL
)

// Connect initializes a new database connection.
func Connect(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %w", err)
	}
	db.AutoMigrate(&Inventory{}) // Automate schema migrations
	return db, nil
}
