package main

import (
	"net/http"

	_ "github.com/abiosoft/caddy-git"
	"github.com/mholt/caddy/caddy/caddymain"

	_ "github.com/ear7h/ear7h-net/api/endpoints"
	"github.com/ear7h/ear7h-net/api"
)

func main() {
	go http.ListenAndServe(":8001", api.Handler)

	makeCaddyfile()
	caddymain.Run()
}