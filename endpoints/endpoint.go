package endpoint

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strings"
	"time"
	"fmt"
)

const shortUrlLength = 4

type StoreHandler struct {
	Store map[string]string
}

func NewstoreHandlers() *StoreHandler {
	return &StoreHandler{

		Store: map[string]string{},
	}

}
func (u *StoreHandler) ShortenUrl(w http.ResponseWriter, r *http.Request) {
	
	CheckUrl := strings.Split(r.URL.Path, "/")[2]

	if len(CheckUrl) > 1 {
		var short_url string
	
		rand.Seed(time.Now().UnixNano())
		chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZÅÄÖ" +
			"abcdefghijklmnopqrstuvwxyzåäö" +
			"0123456789")

		var b strings.Builder
		for i := 0; i < shortUrlLength ; i++ {
			b.WriteRune(chars[rand.Intn(len(chars))])
		}
		short_url = b.String()

		u.Store[short_url] = CheckUrl

	}

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

			http.Redirect(w, r,fmt.Sprintf("http://%s", val), http.StatusMovedPermanently)
		}

	}

	jsonBytes, err := json.Marshal(u.Store)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	w.Write(jsonBytes)

}
