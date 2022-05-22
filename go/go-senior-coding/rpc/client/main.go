package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:9980")
	if err != nil {
		log.Fatal("rpcDial err:", err)
	}

	var reply string
	err = client.Call("HelloService.Hello", "farmer", &reply)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)

}
