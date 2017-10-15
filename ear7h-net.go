package main

import (
	_ "github.com/abiosoft/caddy-git"
	"github.com/mholt/caddy/caddy/caddymain"

	"github.com/ear7h/ear7h-net/api"
	"net/http"
)

func main() {
	go http.ListenAndServe(":8001", api.Handler)

	makeCaddyfile()
	caddymain.Run()
}