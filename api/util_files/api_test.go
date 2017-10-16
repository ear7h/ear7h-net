package main

import (
	"testing"
	"net/http"

	"github.com/ear7h/ear7h-net/api"

	_ "net/http/pprof"
)

func TestApi(t *testing.T) {
	// api server
	go http.ListenAndServe(":8001", api.Handler)

	// for http pprof
	http.ListenAndServe(":8080", nil)
}
