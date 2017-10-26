package main

import (
	"net/http"
	_ "github.com/ear7h/ear7h-net/api/endpoints"
	"github.com/ear7h/ear7h-net/api"
	"os"
	"fmt"
	"github.com/mholt/caddy/caddy/caddymain"
)


const _HEADLESS = `%s {
	proxy /api localhost:8001 {
	    without "/api"
	    header_upstream Host {host}
        header_upstream X-Real-IP {remote}
        header_upstream X-Forwarded-For {remote}
        header_upstream X-Forwarded-Proto {scheme}
        websocket
	}
}`

func makeHeadlessCaddyfile() {
	host := "localhost:8000"

	f, err := os.OpenFile("Caddyfile", os.O_CREATE | os.O_WRONLY | os.O_TRUNC, 0700)
	if err != nil {
		panic(err)
	}


	_, err = fmt.Fprintf(f, _HEADLESS, host)
	if err != nil {
		panic(err)
	}
}


func main() {
	go http.ListenAndServe(":8001", api.Handler)


	makeHeadlessCaddyfile()
	caddymain.Run()
}
