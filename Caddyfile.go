package main

import (
	"fmt"
	"os"
)

const FSTR = `%s {
	root front/_site
	proxy /api localhost:8001 {
	    without "/api"
	    header_upstream Host {host}
        header_upstream X-Real-IP {remote}
        header_upstream X-Forwarded-For {remote}
        header_upstream X-Forwarded-Proto {scheme}
        websocket
	}
	git https://github.com/ear7h/ear7h.github.io ./.. {
		then bundle exec jekyll build
	}
}`

func makeCaddyfile() {
	host := "localhost:8080"

	if os.Getenv("EAR7H_ENV") == "prod" {
		host = "ear7h.net"
	}

	f, err := os.OpenFile("Caddyfile", os.O_CREATE | os.O_WRONLY | os.O_TRUNC, 0700)
	if err != nil {
		panic(err)
	}


	_, err = fmt.Fprintf(f, FSTR, host)
	if err != nil {
		panic(err)
	}
}