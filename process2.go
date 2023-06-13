package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

type API2 struct{}

var statusvar2 string

var x2 int
var k2 int

func (a *API2) ReceiveMsg(msg string, reply *string) error {
	fmt.Printf("P1: ", msg)
	x2 = x2 + 1
	*reply = "P2: Received msg"
	return nil
}

func (a *API2) ReturnStatus(msg string, reply *string) error {
	// fmt.Printf("P1: ", msg)
	if statusvar2 == "Idle" {
		fmt.Print("Taking a local snapshot!")
	}
	*reply = statusvar2
	return nil
}

func main() {
	var reply string

	x2 = 0
	k2 = 0

	statusvar2 = "Active"

	api := new(API2)
	err := rpc.Register(api)
	if err != nil {
		log.Fatal("error registering API", err)
	}

	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":4002")
	if err != nil {
		log.Fatal("Listener error", err)
	}
	log.Printf("serving rpc on port %d", 4002)
	go http.Serve(listener, nil)

	// Retry with timeout
	startTime := time.Now()
	timeout := 10 * time.Second
	connected := false

	for !connected {
		select {
		case <-time.After(time.Second):
			client2, err2 := rpc.DialHTTP("tcp", "localhost:4003")
			if err2 == nil {
				connected = true

				err3 := client2.Call("API3.ReceiveMsg", "Hello P3 from P2", &reply)
				if err3 != nil {
					log.Fatal("Error calling ReceiveMsg: ", err3)
				}

				fmt.Println("Messaged P3 now, Reply from P3:", reply)
				break
			}
		}

		if time.Since(startTime) > timeout {
			// log.Fatal("Timeout: Failed to connect to the RPC server")
		}
	}

	statusvar2 = "Idle"
	x2 = x2 + 1
	k2 = 2
	fmt.Println("Hello Reached Here", statusvar2)

	// timeout1 := 10 * time.Second
	for {
	}

}
