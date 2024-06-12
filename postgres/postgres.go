package postgres

import (
	"auth/config"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func DbConnection() (*sql.DB,error) {
	cfg := config.Load()
	
	con := fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%d sslmode=disable",cfg.DBHOST,cfg.DBUSER,cfg.DBNAME,cfg.DBPASSWORD,cfg.DBPORT)
	fmt.Println(con)
	db, err := sql.Open("postgres",con)
	if err != nil{
		log.Fatal("Error with dbconnection in postgres",err)
		return nil,err
	}
	err = db.Ping()
	if err != nil{
		log.Fatal("Error with dbconnection in postgres",err)
		return nil,err
	}
	return db,nil
}

