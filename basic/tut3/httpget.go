package main

import (
	"io"
	"log"
	"net/http"
)

func main() {

	response, err := http.Get("http://www.google.com/")
	if err != nil {
		log.Fatalln(err)
	}
	defer response.Body.Close()
	
	log.Println("Status:", response.Status)
	log.Println("StatusCode:", response.StatusCode)
	log.Println("ContentLength:", response.ContentLength)

	// Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Body:", string(body))
	
}