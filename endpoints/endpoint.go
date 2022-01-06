package endpoint

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strings"
	"time"
)
type StoreHandler struct {
	Store map[string]string
}

func NewstoreHandlers() *StoreHandler {
	return &StoreHandler{

		Store: map[string]string{},
	}

}
func (u *StoreHandler) Shortenurl(w http.ResponseWriter, r *http.Request) {
	
	check_url := strings.Split(r.URL.Path, "/")[2]

	if len(check_url) > 1 {
		var short_url string
	
		rand.Seed(time.Now().UnixNano())
		chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZÅÄÖ" +
			"abcdefghijklmnopqrstuvwxyzåäö" +
			"0123456789")
		length := 4
		var b strings.Builder
		for i := 0; i < length; i++ {
			b.WriteRune(chars[rand.Intn(len(chars))])
		}
		short_url = b.String()

		u.Store[short_url] = check_url

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

			http.Redirect(w, r, "http://"+val, http.StatusMovedPermanently)
		}

	}

	jsonBytes, err := json.Marshal(u.Store)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	w.Write(jsonBytes)

}
