package home

import (
	"net/http"
	"strings"
	"io/ioutil"

	"github.com/ear7h/ear7h-net/api"
	"github.com/ear7h/ear7h-net/api/auth"
)

var addr string

func setaddr(r *http.Request) {
	addr = strings.Split(r.RemoteAddr, ":")[0]
}

func getString(r *http.Request) string {
	byt, _ := ioutil.ReadAll(r.Body)
	return string(byt)
}


func init() {
	api.HandleFunc("/whereishome", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPut:
			if !auth.IsCorrect(getString(r)) {
				http.Error(w, "oops", http.StatusBadRequest)
				return
			}

			setaddr(r)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("okay"))
		case http.MethodGet:
			if !auth.IsCorrect(r.Header.Get("Auth")) {
				http.Error(w, "oops", http.StatusBadRequest)
				return
			}

			w.WriteHeader(http.StatusOK)
			w.Write([]byte(addr))
		default:
			http.Error(w, "oops", http.StatusMethodNotAllowed)
		}
	})
}