package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

type API3 struct{}

var x3 int
var k3 int

var statusvar3 string

func (a *API3) ReceiveMsg(msg string, reply *string) error {
	fmt.Printf("P2: ", msg)
	x3 = x3 + 1
	*reply = "P3: Received msg"
	return nil
}

func (a *API3) ReturnStatus(msg string, reply *string) error {
	// fmt.Printf("P1: ", msg)
	if statusvar3 == "Idle" {
		fmt.Print("Taking a local snapshot!")
	}
	*reply = statusvar3
	return nil
}

func main() {
	var reply string

	statusvar3 = "Active"

	api := new(API3)
	err := rpc.Register(api)
	if err != nil {
		log.Fatal("error registering API", err)
	}

	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":4003")
	if err != nil {
		log.Fatal("Listener error", err)
	}
	log.Printf("serving rpc on port %d", 4003)
	go http.Serve(listener, nil)

	// Retry with timeout
	startTime := time.Now()
	timeout := 10 * time.Second
	connected := false

	for !connected {
		select {
		case <-time.After(time.Second):
			client2, err2 := rpc.DialHTTP("tcp", "localhost:4004")
			if err2 == nil {
				connected = true

				err3 := client2.Call("API4.ReceiveMsg", "Hello P4 from P3", &reply)
				if err3 != nil {
					log.Fatal("Error calling ReceiveMsg: ", err3)
				}

				fmt.Println("Messaged P4 now, Reply from P4:", reply)
				break
			}
		}

		if time.Since(startTime) > timeout {
			// log.Fatal("Timeout: Failed to connect to the RPC server")
		}
	}
	statusvar3 = "Idle"
	fmt.Println("Hello Reached Here", statusvar3)
	// timeout1 := 10 * time.Second
	for {
	}

}
