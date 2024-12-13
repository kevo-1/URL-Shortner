package main

import (
	"fmt"
	"UrlShortner/utils"
)

func main() {
	fmt.Println("Enter the link you want to shorten: ")
	var URL string
	fmt.Scanln(&URL)
	hashed := utils.MD5Hash(URL)
	fmt.Println("Original Link: ", URL,"\nShortened Link: ", hashed)
}