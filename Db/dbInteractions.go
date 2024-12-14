package db

import (
	"database/sql"
	"errors"
	"log"
)

func InsertData(hash, originalUrl string) error{

	if test, err := RetrieveData(hash); test != "" && err == nil{
		log.Println("Url already exists with provided hash")
		return errors.New("already exists")
	}

	query := "INSERT INTO urldata (shortened, original) VALUES (?, ?)"
	_, err := db.Exec(query, hash, originalUrl)
	if err != nil {
		log.Println("Failed while saving Url: ",err)
		return err
	}

	log.Println("URL successfully saved")
	return nil;
}


func RetrieveData(hash string) (string, error) {
	query := "Select Original from urldata where shortened = ?"
	var originalUrl string
	err := db.QueryRow(query, hash).Scan(&originalUrl)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", nil
		}
		log.Println("Failed to retrieve data:", err)
		return "", err
	}

	return originalUrl, nil
}