package main

import (

	"net/http"
	"example/urlshortener/endpoints"

)

func main() {

	urlstore := endpoint.NewstoreHandlers()

	http.HandleFunc("/save/", urlstore.Shortenurl)
	http.HandleFunc("/", urlstore.Redirect)
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		panic(err)
	}

}
