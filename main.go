// Rest api for urlshortener

package main

import (
	endpoint "example/urlshortener/endpoints"
	"log"
	"net/http"
)

func main() {

	urlstore := endpoint.NewstoreHandlers()

	http.HandleFunc("/save/", urlstore.ShortenUrl)
	http.HandleFunc("/", urlstore.Redirect)
	log.Println("Listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
