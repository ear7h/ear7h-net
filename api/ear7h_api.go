package api

import (
	"fmt"
	"net/http"
	"strings"
)

var mux = http.NewServeMux()
var alreadyHandled = []string{}

var Handler = http.Handler(mux)

func IsHandled(pattern string) (b bool) {
	for _, v := range alreadyHandled {
		if v == pattern {
			return true
		}
	}

	alreadyHandled = append(alreadyHandled, pattern)

	return
}

func isTopLevel(pattern string) (b bool) {
	// valid : /asd, /asd/
	// invalid : a, as/d, /asd/a
	return pattern[0] == '/' && !strings.Contains(pattern[1:len(pattern) - 1], "/")
}


// HandleFunc makes sure the handler isn't already handled and the
// pattern is for a top level dir, and adds the pattern and handler to
// the global handler. The matched pattern is stripped request's url
// before passing it to registered handler.
func HandleFunc(pattern string, h http.HandlerFunc) {
	if IsHandled(pattern) {
		panic(fmt.Errorf("%s already handled", pattern))
	}

	if !isTopLevel(pattern) {
		panic(fmt.Errorf("%s invalid, only top level directories can be registered", pattern))
	}

	mux.HandleFunc(pattern, func (w http.ResponseWriter, r *http.Request) {
		r.URL.Path = r.URL.Path[len(pattern) - 1:]
		h(w, r)
	})
}

// Handle delegates to Handle func
func Handle(pattern string, h http.Handler) {
	HandleFunc(pattern, h.ServeHTTP)
}

func init() {
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("Endpoints: %s", alreadyHandled)))
	})
}