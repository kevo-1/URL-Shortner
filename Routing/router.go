package routing

import (
	db "UrlShortner/db"
	"UrlShortner/utils"
	"fmt"
	"net/http"
	"log"
)

func generateShortUrl(writer http.ResponseWriter, request *http.Request) {
	original := request.URL.Query().Get("url")
	if original == "" {
		http.Error(writer, "Provide a URL in the 'url' parameter", http.StatusBadRequest)
		return
	}

	var shortUrl string
	hash := utils.MD5Hash(original)
	if temp, err := db.RetrieveData(hash[:8]); temp != "" && err == nil {
		shortUrl = fmt.Sprintf("http://localhost:8080/%s", hash[:8])
	} else if err != nil {
		log.Printf("Error retrieving data: %v", err)
		http.Error(writer, "Internal server error", http.StatusInternalServerError)
		return
	} else {
		err := db.InsertData(hash[:8], original)
		if err != nil {
			log.Printf("Error inserting data: %v", err)
			http.Error(writer, "Internal server error", http.StatusInternalServerError)
			return
		}
		shortUrl = fmt.Sprintf("http://localhost:8080/%s", hash[:8])
	}
	fmt.Fprintf(writer, "Shortened Link: %s", shortUrl)
}



func Run(){ 
	http.HandleFunc("/generate", generateShortUrl)
	http.HandleFunc("/", RedirectUrl)
	log.Printf("Server started at https://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println("Error starting server: ", err)
	}
}