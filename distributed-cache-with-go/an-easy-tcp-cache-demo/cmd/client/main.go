package main

import (
	"flag"
	"fmt"
	"tcp-cache-service/tcp"
)

func main() {
	server := flag.String("h", "localhost", "cache server address")
	port := flag.String("p", "2333", "cache server port")
	option := flag.String("c", "set", "command, could be get/set/del")
	key := flag.String("k", "test-key", "key")
	value := flag.String("v", "test-value", "value")
	flag.Parse()

	for i := 0; i < 10; i++ {

		err := tcp.Client(*server, *port, *option, *key, *value)
		if err != nil {
			fmt.Println(err)
		}
	}

}
