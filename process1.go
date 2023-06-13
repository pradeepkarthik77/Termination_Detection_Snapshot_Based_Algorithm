package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

type API1 struct{}

var statusvar string

func (a *API1) ReceiveMsg(msg string, reply *string) error {
	fmt.Printf("P1: ", msg)
	*reply = "P1: Received msg"
	return nil
}

func (a *API1) ReturnStatus(msg string, reply *string) error {
	if statusvar == "Idle" {
		fmt.Print("Taking a local snapshot!")
	}
	*reply = statusvar
	return nil
}

func main() {
	var reply string

	statusvar = "Active"

	api := new(API1)
	err := rpc.Register(api)
	if err != nil {
		log.Fatal("error registering API", err)
	}

	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":4001")
	if err != nil {
		log.Fatal("Listener error", err)
	}
	log.Printf("serving rpc on port %d", 4001)
	go http.Serve(listener, nil)

	// Retry with timeout
	startTime := time.Now()
	timeout := 10 * time.Second
	connected := false

	for !connected {
		select {
		case <-time.After(time.Second):
			client2, err2 := rpc.DialHTTP("tcp", "localhost:4002")
			if err2 == nil {
				connected = true

				err3 := client2.Call("API2.ReceiveMsg", "Hello P2 from P1", &reply)
				if err3 != nil {
					log.Fatal("Error calling ReceiveMsg: ", err3)
				}

				fmt.Println("Messaged P2 now, Reply from P2:", reply)
			}
		}

		if time.Since(startTime) > timeout {
			// log.Fatal("Timeout: Failed to connect to the RPC server")
		}
	}

	statusvar = "Idle"

	localist := []string{"localhost:4001", "localhost:4002", "localhost:4003", "localhost:4004"}
	apilist := []string{"API1.ReturnStatus", "API2.ReturnStatus", "API3.ReturnStatus", "API4.ReturnStatus"}

	timeout1 := 20 * time.Second

	select {
	case <-time.After(timeout1):
		for i := 0; i <= 3; i++ {
			client, err := rpc.DialHTTP("tcp", localist[i])

			if err != nil {
				log.Fatal("Connection error: ", err)
			}

			client.Call(apilist[i], "", &reply)

			if reply == "Active" {
				fmt.Println("All processes are not yet terminated")
				return
			}
		}
	}

	fmt.Println("All processes are terminated and all local snapshots are recorded")

}
