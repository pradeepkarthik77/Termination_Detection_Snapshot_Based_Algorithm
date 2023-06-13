package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"
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

	var reply string

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

	localist := []string{"localhost:4001", "localhost:4002", "localhost:4003", "localhost:4004"}
	apilist := []string{"API1.ReturnStatus", "API2.ReturnStatus", "API3.ReturnStatus", "API4.ReturnStatus"}

	timeout1 := 1 * time.Second

	select {
	case <-time.After(timeout1):
		statusvar4 = "Idle"
		x4 = x4 + 1
		k4 = 4
		fmt.Println("Hello Reached Here", statusvar4)
		for i := 0; i <= 3; i++ {
			client, err := rpc.DialHTTP("tcp", localist[i])

			if err != nil {
				log.Fatal("Connection error: ", err)
			}

			client.Call(apilist[i], "", &reply)

			if reply == "Active" {
				fmt.Println("All processes are not yet terminated")
				for {

				}
			}
		}
	}

	fmt.Println("All processes are terminated and all local snapshots are recorded")

	for {

	}

}
