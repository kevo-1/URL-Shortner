package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func StartDB() {
	cfg, err := loadEnv()
	if err != nil {
		log.Fatal("Failed to load environment variables:", err)
	}

	db, err = sql.Open("mysql", cfg)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("Database connection failed:", err)
	}

	log.Println("Connected successfully to the database!")
}


func loadEnv() (string, error) {
	if err := godotenv.Load("./envs/.env"); err != nil {
		log.Println("Failed to load .env file, currently using system env variables")
	}

	User := os.Getenv("DB_USER")
	Pswd := os.Getenv("DB_PASSWORD")
	Host := os.Getenv("DB_HOST")
	Port := os.Getenv("DB_PORT")
	Name := os.Getenv("DB_NAME")

	if User == "" || Pswd == "" || Host == "" || Port == "" || Name == "" {
		return "", fmt.Errorf("missing required database configuration")
	}

	cfg := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", User, Pswd, Host, Port, Name)
	return cfg, nil
}

func GetDB() *sql.DB {
	if db == nil {
		log.Fatal("Database connection has not been initialized. Call StartDB first.")
	}
	return db
}
