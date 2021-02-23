package main

import (
	"github.com/pebbe/zmq3"
	"log"
)

func main() {
	context, _ := zmq3.NewContext()
	socket, _ := context.NewSocket(zmq3.REP)
	err := socket.Bind("tcp://127.0.0.1:5000")
	if err != nil {
		log.Fatal("socket bind occur err:", err)
	}

	for {
		msg, _ := socket.Recv(0)
		println("Got", msg)
		_, err = socket.Send(msg, 0)
		if err != nil {
			log.Fatal("socket send occur err:", err)
		}
	}
}
