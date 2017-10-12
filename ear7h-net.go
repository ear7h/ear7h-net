package main

import (
	_ "github.com/abiosoft/caddy-git"
	"github.com/mholt/caddy/caddy/caddymain"

	"github.com/ear7h/ear7h-net/api"
)

func main() {
	go api.Main()

	makeCaddyfile()
	caddymain.Run()
}