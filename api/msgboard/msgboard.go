package msgboard

import (
	"net/http"
	"fmt"
	"runtime"
	"path"

	"github.com/gorilla/websocket"

	"github.com/ear7h/ear7h-net/api"
)

const STORE = "store.txt"

var msgTower = &tower{
	input: make(chan message, 10),
	output: make([]chan message, 0),
}

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) (ok bool) {
		return true
	},
}


func init() {
	go msgTower.start()
	api.HandleFunc("/msgboard/", ServeHTTP)
}



func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("msg",r.URL)
	if r.URL.Path == "/ws" {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			fmt.Println(err)
			return
		}

		go msgTower.newClient(conn)
		return
	}

	_, f, _, _ := runtime.Caller(0)

	http.ServeFile(w, r, path.Join(path.Dir(f), "/index.html"))
}