package postgres

import (
	"database/sql"
	
	_ "github.com/lib/pq"
)


func ConnectDb() (*sql.DB, error){
	psql := "user=postgres password=20005 dbname=restarount sslmode=disable"
	db, err := sql.Open("postgres", psql)
	if err != nil {
		return nil, err
	}
	return db, nil
}
