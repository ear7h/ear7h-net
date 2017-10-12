package api

import (
	"net/http"
	"github.com/ear7h/ear7h-net/api/home"
	"io/ioutil"
	"fmt"
	"strings"
	"golang.org/x/crypto/ssh/terminal"
	"syscall"
)

var PASSWORD string

func init() {
	fmt.Println("enter desired password")
	byt, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		panic(err)
	}

	PASSWORD = string(byt)
}

func getString(r *http.Request) string {
	byt, _ := ioutil.ReadAll(r.Body)
	return string(byt)
}

func Main() {
	m := http.NewServeMux()

	m.HandleFunc("/whereishome", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPut:
			if getString(r) != PASSWORD {
				http.Error(w, "oops", http.StatusBadRequest)
				return
			}

			home.Put(strings.Split(r.RemoteAddr, ":")[0])
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("okay"))
		case http.MethodGet:
			if r.Header.Get("Auth") != PASSWORD {
				http.Error(w, "oops", http.StatusBadRequest)
				return
			}

			home.Put(strings.Split(r.RemoteAddr, ":")[0])
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(home.Get()))
		default:
			http.Error(w, "oops", http.StatusMethodNotAllowed)
		}
	})

	m.HandleFunc("/em", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
<!doctype html>
<body>
<p>this isn't the most extraordinary piece of code</p><br>
		<p style="display: inline">but click </p>
		<button style="display: inline" onclick="goof()">here</button>
		<p style="display: inline">repeatedly</p>
		<div id="a-div">
		</div>
		<script>
			function goof() {
			let d = document.getElementById("a-div")

			let c = document.createElement("p")
			c.style.color = "#" + ((1 << 24) * Math.random() | 0).toString(16)
			c.innerHTML = "i can't wait to see you, goof"

			d.appendChild(c)
		}
		</script>
		</body>
		`))
	})

	http.ListenAndServe(":8001", m)
}
