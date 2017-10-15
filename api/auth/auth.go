package auth

import (
	"os"
	"fmt"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

var password string

func init() {
	if os.Getenv("EAR7H_ENV") == "prod" {
		setPass()
		return
	}

	password = "asd"
}


func setPass() {
	fmt.Println("enter desired password")
	byt, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		fmt.Println("no error pass set: ",err)
	}

	password = string(byt)
}

func IsCorrect(str string) bool {
	return str == password
}