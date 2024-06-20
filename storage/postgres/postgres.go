package postgres

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectDb() (*sql.DB, error) {
	psql := "user=sayyidmuhammad password=root dbname=vote sslmode=disable"
	db, err := sql.Open("postgres", psql)
	if err != nil {
		return nil, err
	}
	return db, nil
}
