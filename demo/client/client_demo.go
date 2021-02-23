package main

import (
	"github.com/pebbe/zmq3"
	"log"
)

func main() {
	context, _ := zmq3.NewContext()
	socket, _ := context.NewSocket(zmq3.REQ)

	err := socket.Connect("tcp://127.0.0.1:5000")
	if err != nil {
		log.Fatal("socket connect occur err:", err)
	}

	_, err = socket.Send("hello world", 0)
	if err != nil {
		log.Fatal("socket send occur err:", err)
	}
}
