package api

import (
	"testing"
	"net/http"

	_ "net/http/pprof"
)

func TestApi(t *testing.T) {
	// api server
	go http.ListenAndServe(":8001", Handler)

	// for http pprof
	http.ListenAndServe(":8080", nil)
}
