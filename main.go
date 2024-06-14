package main

import (
	"fmt"
	"log"

	"github.com/Project_Restaurant/Auth-Service/api"
	_ "github.com/Project_Restaurant/Auth-Service/docs"
	"github.com/Project_Restaurant/Auth-Service/postgres"
)

func main() {
	db, err := postgres.ConnectDb()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router := api.NewRouter(db)

	fmt.Println("Server is running on port 8081")
	if err := router.Run(":8081"); err != nil {
		log.Fatal(err)
	}

}
