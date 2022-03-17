package pages

import (
	"net/http"
	"os"
	"strings"
)

func Accets(w http.ResponseWriter, r *http.Request) {
	url := r.URL.String()
	url = strings.Replace(url, "/accets/", "", 1)

	bytes, err := os.ReadFile("accets/" + url)
	if err != nil {
		w.WriteHeader(404)
		w.Write([]byte("Requested accet not found of " + url))
		return
	}

	w.WriteHeader(200)
	w.Write(bytes)
}
