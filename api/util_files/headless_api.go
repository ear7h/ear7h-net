package main

import (
	"net/http"
	_ "github.com/ear7h/ear7h-net/api/endpoints"
	"github.com/ear7h/ear7h-net/api"
)

func main() {
	http.ListenAndServe(":8001", api.Handler)
}
