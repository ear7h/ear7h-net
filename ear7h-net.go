package main

import (
	"github.com/ear7h/ear7h-net/api"
	"os/exec"
)

func main() {
	go api.Main()
	exec.Command("caddy").Start()

	hang := make(chan bool)

	<- hang
}