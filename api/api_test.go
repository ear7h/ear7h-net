package api

import (
	"testing"
	"net/http"

	_ "net/http/pprof"
)

func TestApi(t *testing.T) {
	go Main()

	// for http pprof
	http.ListenAndServe(":8080", nil)
}
