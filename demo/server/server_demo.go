package main

import (
	"github.com/pebbe/zmq3"
	"log"
)

func main() {
	socket, err := zmq3.NewSocket(zmq3.SUB)
	if err != nil {
		log.Fatal("zmq new socket occur err:", err)
	}
	defer closeZmq3Socket(socket)

	err = socket.Bind("tcp://localhost:1234")
	if err != nil {
		log.Fatal("socket bind occur err:", err)
	}

	res, err := socket.Recv(0)
	if err != nil {
		log.Fatal("socket recv occur err:", err)
	}
	log.Println(res)

}

func closeZmq3Socket(socket *zmq3.Socket) {
	if socket != nil {
		err := socket.Close()
		if err != nil {
			log.Fatal("socket close occur err:", err)
		}
	}
}
