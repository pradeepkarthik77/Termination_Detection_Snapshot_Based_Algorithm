package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type API4 struct{}

var x4 int
var k4 int

var statusvar4 string

func (a *API4) ReceiveMsg(msg string, reply *string) error {
	fmt.Printf("P2: ", msg)
	x4 = x4 + 1
	*reply = "P3: Received msg"
	return nil
}

func (a *API4) ReturnStatus(msg string, reply *string) error {
	// fmt.Printf("P1: ", msg)
	if statusvar4 == "Idle" {
		fmt.Print("Taking a local snapshot!")
	}
	*reply = statusvar4
	return nil
}

func main() {

	statusvar4 = "Active"

	api := new(API4)
	err := rpc.Register(api)
	if err != nil {
		log.Fatal("error registering API", err)
	}

	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":4004")
	if err != nil {
		log.Fatal("Listener error", err)
	}
	log.Printf("serving rpc on port %d", 4004)
	go http.Serve(listener, nil)

	statusvar4 = "Idle"
	fmt.Println("Hello Reached Here", statusvar4)

	// timeout1 := 10 * time.Second
	for {
	}

}
