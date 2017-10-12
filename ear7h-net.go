package main

import (
	"github.com/ear7h/ear7h-net/api"
	"os/exec"
	"os"
)

func main() {
	go api.Main()
	cmd := exec.Command("caddy")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	hang := make(chan bool)

	<- hang
}