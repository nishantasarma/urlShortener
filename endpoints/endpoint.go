package endpoint

import (
	"encoding/json"
	model "example/urlshortener/models"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

const shortUrlLength = 4

type StoreHandler struct {
	Store map[string]string
}

func NewstoreHandlers() *StoreHandler {
	u := StoreHandler{Store: map[string]string{}}
	return &u

}

func (u *StoreHandler) ShortenUrl(w http.ResponseWriter, r *http.Request) {

	longUrl := strings.Split(r.URL.Path, "/")[2]
	var shortUrl string
	if len(longUrl) > 1 {

		rand.Seed(time.Now().UnixNano())
		chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZÅÄÖ" +
			"abcdefghijklmnopqrstuvwxyzåäö" +
			"0123456789")

		var b strings.Builder
		for i := 0; i < shortUrlLength; i++ {
			b.WriteRune(chars[rand.Intn(len(chars))])
		}
		shortUrl = b.String()

	}
	model.Insertdb(shortUrl, longUrl)
	u.Store = *model.Getdb()
	jsonBytes, err := json.Marshal(u.Store)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	w.Write(jsonBytes)

}

func (u *StoreHandler) Redirect(w http.ResponseWriter, r *http.Request) {

	redirect_url := strings.Split(r.URL.Path, "/")[1]

	if len(redirect_url) > 1 {

		if val, ok := u.Store[redirect_url]; ok {

			http.Redirect(w, r, fmt.Sprintf("http://%s", val), http.StatusMovedPermanently)
		}

	}

	u.Store = *model.Getdb()

	jsonBytes, err := json.Marshal(u.Store)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	w.Write(jsonBytes)

}
