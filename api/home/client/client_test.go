package main

import (
	"testing"
	"net/http"
	"io/ioutil"
	"fmt"
)

const _TEST_SERVER_PORT  = "8081"

func TestAll(t *testing.T) {
	go main()

	http.HandleFunc("/asd", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("test server hit")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("asd"))
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("catch all", r.URL.Path)
	})


	go http.ListenAndServe(":" + _TEST_SERVER_PORT, http.DefaultServeMux)

	res, err := http.Get(_EAR7H_URL + "/" + _TEST_SERVER_PORT +"/asd")
	if err != nil {
		panic(err)
	}

	byt, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("res: ", string(byt))
	if string(byt) != "asd" {
		t.Fail()
	}
}
