package main

import (
	"os"
	"fmt"
	"syscall"
	"golang.org/x/crypto/ssh/terminal"
	"time"
	"net/http"
	"io/ioutil"
)

var _EAR7H_URL string
var _PASSWORD string

func setPass() {



	fmt.Println("enter desired password")
	byt, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		fmt.Println("no error pass set: ",err)
	}

	_PASSWORD = string(byt)
}

func init() {
	if os.Getenv("EAR7H_ENV") == "prod" {
		_EAR7H_URL = "https://ear7h.net/api/home"
		setPass()
		return
	}

	_EAR7H_URL = "http://localhost:8000/api/home"
	_PASSWORD = "asd"
}

func makeRequest() {
	c := http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequest(http.MethodPut, _EAR7H_URL, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Set("Authorization", "password " + _PASSWORD)

	res, err := c.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	byt, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(byt))
}

// main loop
func main() {
	for {
		makeRequest()
		time.Sleep(3 * time.Second)
	}
}
