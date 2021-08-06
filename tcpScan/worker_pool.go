package main

import (
	"fmt"
	"net"
	"sort"
	"time"
)

func worker(ports chan int, results chan int) {
	for p := range ports {
		// fmt.Println(p)
		address := fmt.Sprintf("farmer233.top:%d", p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- p
	}
}

func main() {
	start := time.Now()
	ports := make(chan int, 30000)
	results := make(chan int, 2000)
	scanEndPort := 65535
	var openPorts []int
	var closePorts []int

	// 生成worker
	for i := 0; i < cap(ports); i++ {
		go worker(ports, results)
	}
	// 启动worker
	go func() {
		for i := 1; i < scanEndPort; i++ {
			ports <- i
		}
	}()

	// 处理worker
	// for i := 1; i < 65535; i++ {
	// 	port := <-results
	// 	if port != 0 {
	// 		openPorts = append(openPorts, port)
	// 	} else {
	// 		closePorts = append(closePorts, port)
	// 	}
	// }

	for i := 1; i < scanEndPort; i++ {
		res := <-results
		if res != 0 {
			openPorts = append(openPorts, res)
		} else {
			closePorts = append(closePorts, res)
		}
	}

	close(ports)
	close(results)
	sort.Ints(openPorts)
	// for _, port := range closePorts {
	// 	fmt.Printf("%d closed\n", port)
	// }
	for _, port := range openPorts {
		fmt.Printf("%d opend\n", port)
	}
	elapsed := time.Since(start) / 1e9
	fmt.Printf("\n\n%d seconds\n", elapsed)

}
