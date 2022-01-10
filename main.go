package main

import (

	"net/http"
	"example/urlshortener/endpoints"
	"log"

)

func main() {

	urlstore := endpoint.NewstoreHandlers()

	http.HandleFunc("/save/", urlstore.ShortenUrl)
	http.HandleFunc("/", urlstore.Redirect)
	log.Println("Listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
