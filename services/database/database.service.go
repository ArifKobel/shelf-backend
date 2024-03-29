package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	dsn := "host=localhost user=db password=db dbname=db port=5432 sslmode=disable"
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
