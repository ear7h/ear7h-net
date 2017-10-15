package api

import (
	"fmt"
	"net/http"
	"strings"

	// endpoints
	_ "github.com/ear7h/ear7h-net/api/em"
	_ "github.com/ear7h/ear7h-net/api/home"
	_ "github.com/ear7h/ear7h-net/api/msgboard"
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

	return
}

func isTopLevel(pattern string) (b bool) {
	// valid : /asd
	// invalid : as/d /asd/a
	return pattern[0] == '/' && !strings.Contains(pattern[1:], "/")
}

func HandleFunc(pattern string, h http.HandlerFunc) {
	if IsHandled(pattern) {
		panic(fmt.Errorf("%s already handled", pattern))
	}

	if !isTopLevel(pattern) {
		panic(fmt.Errorf("%s invalid, only top level directories can be registered", pattern))
	}

	mux.HandleFunc(pattern, h)
}

func Handle(pattern string, h http.Handler) {
	if IsHandled(pattern) {
		panic(fmt.Errorf("%s already handled", pattern))
	}

	if !isTopLevel(pattern) {
		panic(fmt.Errorf("%s invalid, only top level directories can be registered", pattern))
	}

	mux.Handle(pattern, h)
}
