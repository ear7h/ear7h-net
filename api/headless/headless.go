package main

import (
	"net/http"

	"github.com/ear7h/ear7h-net/api"
)

func main() {
	http.ListenAndServe(":8001", api.Handler)
}
