package main

import (
	"net/http"
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"syscall"
	"io/ioutil"
)

func main() {
	fmt.Println("enter password")
	byt, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		panic(err)
	}

	password := string(byt)


	r, err := http.NewRequest(http.MethodGet, "https://ear7h.net/api/home", nil)

	if err != nil {
		panic(err)
	}

	r.Header.Set("Authorization", "password " + password)

	res, err := http.DefaultClient.Do(r)
	if err != nil {
		panic(err)
	}

	byt, err = ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()


	fmt.Println(string(byt))
}
