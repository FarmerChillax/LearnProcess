package main

import (
	"fmt"
	"log"
	"net/rpc"
	"sync"
)

var wg sync.WaitGroup

func main() {

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go Client(i)
	}
	wg.Wait()
}

func Client(id int) {

	client, err := rpc.Dial("tcp", "localhost:9980")
	if err != nil {
		log.Fatal("rpcDial err:", err)
	}

	var userName, reply string
	userName = fmt.Sprintf("caller %d", id)
	err = client.Call("HelloService.Hello", userName, &reply)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)
	wg.Done()

}
