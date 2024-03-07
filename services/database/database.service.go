package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open("postgres", "user=db password=db dbname=db sslmode=disable")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func Migrate() error {
	db, err := Connect()
	if err != nil {
		return err
	}
	defer db.Close()
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS files (id SERIAL PRIMARY KEY, name VARCHAR(255), file_name VARCHAR(255), created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP)")
	return err
}
