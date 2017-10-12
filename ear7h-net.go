package main

import (
	"github.com/ear7h/ear7h-net/api"
	"os/exec"
	"os"
	"bufio"
	"fmt"
	"os/signal"
	"syscall"
)


func main() {
	go api.Main()

	caddyCmd := exec.Command("caddy")
	caddyCmd.Stdout = os.Stdout
	caddyCmd.Stderr = os.Stderr
	go caddyCmd.Start()

	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt)
	<- sig

	caddyCmd.Process.Kill()
	os.Exit(0)

}