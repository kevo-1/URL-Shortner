package routing

import (
	"UrlShortner/db"
	"log"
	"net/http"
)

func RedirectUrl(writer http.ResponseWriter, request *http.Request) {
	shortUrl := request.URL.Path[1:]

	log.Printf("Extracted Short URL Hash: %s", shortUrl)

	original, err := db.RetrieveData(shortUrl)
	if err != nil {
		log.Printf("Error retrieving original URL for hash '%s': %v", shortUrl, err)
		http.Error(writer, "Internal server error", http.StatusInternalServerError)
		return
	}

	if original == "" {
		http.Error(writer, "Short URL not valid", http.StatusNotFound)
		return
	}

	log.Printf("Redirecting to: %s", original)
	http.Redirect(writer, request, "http://"+original, http.StatusFound)
}



