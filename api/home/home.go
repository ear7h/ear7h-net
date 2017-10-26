package home

import (
	"net/http"
	"strings"
	"net/http/httputil"
	"strconv"

	"github.com/ear7h/ear7h-net/api"
	"github.com/ear7h/ear7h-net/api/auth"
	"fmt"
)

var _HOME_ADDR string


//requests to /home
func handleHome(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPut:
		if !auth.IsHeaderCorrect(r) {
			http.Error(w, "auth problem", http.StatusBadRequest)
			return
		}

		// set the address
		// this header is set by the caddy proxy
		_HOME_ADDR = r.Header.Get("X-Forwarded-For")
		// for testing
		if _HOME_ADDR == "::1" {
			_HOME_ADDR = "127.0.0.1"
		}

		w.WriteHeader(http.StatusOK)
	case http.MethodGet:
		if !auth.IsHeaderCorrect(r) {
			http.Error(w, "auth problem", http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(_HOME_ADDR))
	default:
		http.Error(w, "get or put only", http.StatusMethodNotAllowed)
	}
}

// get port number from /home/{int} or /home/{int}/*
func getPort(path string) (port string, ok bool) {
	arr := strings.Split(path,"/")[1:]
	fmt.Println(arr)
	if len(arr) < 2 {
		return
	}

	if _, err := strconv.ParseUint(arr[0], 10, 32); err != nil {
		return
	}

	port, ok = arr[0], true
	return
}

func init() {

	/*
	 * journal in code 1
	 *
	 * I hate this community college shit
	 * but it's what I get for the shit I did when
	 * I should have been studying
	 *
	 * stats quiz: section 5 and 6
	 */

	// set and get address
	api.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
			handleHome(w, r)
	})

	// set the proxy to redirect requests to the global addr
	// with the port based on the request url
	proxy := httputil.ReverseProxy{
		Director: func (r *http.Request) {
			// get the port string from the path
			// ignore error because it was checked in the handler
			port, _ := getPort(r.URL.Path)

			fmt.Println("path: ", r.URL.Path)
			fmt.Println("parsed path: ", r.URL.Path[len(port) + 1:])

			r.URL.Host += _HOME_ADDR + ":" + port
			r.URL.Path = r.URL.Path[len(port) + 1:]
			r.URL.Scheme = r.Header.Get("X-Forwarded-Proto")

			fmt.Println("URL: ", r.URL)
		},
	}

	// request proxy
	api.HandleFunc("/home/", func (w http.ResponseWriter, r *http.Request) {
		if _HOME_ADDR == "" {
			http.Error(w, "proxy not set up", http.StatusBadGateway)
			return
		}

		if _, ok := getPort(r.URL.Path); !ok {
			http.Error(w, "couldn't parse port number", http.StatusBadRequest)
			return
		}

		fmt.Println("proxying")
		proxy.ServeHTTP(w, r)
	})
}
