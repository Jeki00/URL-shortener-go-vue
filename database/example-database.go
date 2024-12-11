package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var db *sql.DB

func Connect() (*sql.DB, error) {
	var err error
	db, err := sql.Open("postgres", fmt.Sprintf("postgres://<username>:<password>@<port ex:127.0.0.5432>/<mydb>?sslmode=disable"))
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
