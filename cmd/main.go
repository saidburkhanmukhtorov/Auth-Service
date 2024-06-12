package main

import (
	"auth/api"
	"auth/postgres"
	"fmt"
	"log"
)

func main(){
	db,err := postgres.DbConnection()
	if err != nil{
		log.Fatal("Error whith connection DB",err)
		return
	}
	router := api.NewEngine(db)
	err = router.Run()
	if err != nil{
		log.Fatal("Error whith router",err)
	}
	fmt.Println("Router is running on :8080")
}