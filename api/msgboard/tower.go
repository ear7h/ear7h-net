package msgboard

import (
	"github.com/gorilla/websocket"
	"fmt"
)

type message struct {
	Name string `json:"name"`
	Data string `json:"data"`
}
type tower struct {
	input  chan message
	output []chan message
}

func (t *tower) newClient(conn *websocket.Conn) {
	clientListen := make(chan message, 10)

	t.output = append(t.output, clientListen)

	//listen
	go func() {
		defer conn.Close()

		for {
			var msg message

			err := conn.ReadJSON(&msg)
			if err != nil {
				conn.WriteMessage(websocket.BinaryMessage, []byte("bad json"))
				fmt.Println(err)
				break
			}

			fmt.Println("new message", msg)
			t.input <- msg
		}
	}()

	//serve
	go func() {
		for {
			select {
			case msg := <- clientListen:
				fmt.Println("broadcasting", msg)

				err := conn.WriteJSON(msg)
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	}()
}

func (t *tower) start() {
	for {
		val := <- t.input
		for _, v := range t.output{
			v <- val
		}
	}
}