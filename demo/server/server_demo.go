package main

import (
	"github.com/alecthomas/gozmq"
	"log"
)

func main() {
	context, _ := gozmq.NewContext()
	socket, _ := context.NewSocket(gozmq.REP)
	err := socket.Bind("tcp://127.0.0.1:5000")
	if err != nil {
		log.Fatal("socket bind occur err:", err)
	}

	for {
		msg, _ := socket.Recv(0)
		println("Got", string(msg))
		err = socket.Send(msg, 0)
		if err != nil {
			log.Fatal("socket send occur err:", err)
		}
	}
}
