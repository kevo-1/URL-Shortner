package db

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetDomain() string {
	err := godotenv.Load("./envs/.env")

	if err != nil {
		log.Println("Couldn't load env file")
		return ""
	}
	
	domain := os.Getenv("Domain")
	if domain == "" {
		log.Println("Couldn't load domain")
		return ""
	} else {
		return domain
	}
}