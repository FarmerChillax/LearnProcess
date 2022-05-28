package main

import (
	"fmt"
	"log"
	"net/rpc"
	"rpc-service-advance/service"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var reply string

	err = client.Call(service.HelloServiceName+".Hello", "farmer", &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)

}
