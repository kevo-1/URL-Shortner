package main

import (
	"UrlShortner/Db"
	"UrlShortner/utils"
	"fmt"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("./envs/.env")

	if err != nil {
		fmt.Println(err)
		return
	}

	domain := db.GetDomain()
	if domain == "" {
		fmt.Println("Domain Not loaded")
	}

	db.StartDB()
	fmt.Println("Enter the link you want to shorten: ")
	var URL string
	fmt.Scanln(&URL)
	hashed := utils.MD5Hash(URL)
	shortLink := fmt.Sprintf("%s%s", domain, hashed[:8])
	fmt.Println("Original Link: ", URL,"\nShortened Link: ", shortLink)
	
	db.InsertData(hashed[:8], URL)
}