package main

import (
	"UrlShortner/utils"
	"UrlShortner/Db"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("./envs/.env")

	if err != nil {
		fmt.Println(err)
		return
	}

	domain := os.Getenv("DOMAIN")
	if domain == "" {
		fmt.Println("Not loaded")
	}

	db.StartDB()
	fmt.Println("Enter the link you want to shorten: ")
	var URL string
	fmt.Scanln(&URL)
	hashed := utils.MD5Hash(URL)
	shortLink := fmt.Sprintf("%s%s", domain, hashed[:6])
	fmt.Println("Original Link: ", URL,"\nShortened Link: ", shortLink)
	
	db.InsertData(hashed[:8], URL)
}