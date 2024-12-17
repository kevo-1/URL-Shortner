package main

import (
	db "UrlShortner/db"
	routing "UrlShortner/Routing"
	"fmt"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("./envs/.env")

	if err != nil {
		fmt.Println(err)
		return
	}

	db.StartDB()
	routing.Run()	
}