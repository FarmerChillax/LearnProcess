package main

import (
	"fmt"
	"net"
	"sort"
	"time"
)

func worker(ports chan int, results chan int) {
	for p := range ports {
		address := fmt.Sprintf("172.16.90.%d:80", p)
		fmt.Printf("start %s\n", address)
		conn, err := net.DialTimeout("tcp", address, time.Second*1)
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
	ports := make(chan int, 100)
	results := make(chan int, 100)
	scanEndPort := 255
	var openPorts []int
	var closePorts []int

	// 生成worker
	for i := 0; i < cap(ports); i++ {
		go worker(ports, results)
	}
	// push port to channel
	go func() {
		for i := 1; i <= scanEndPort; i++ {
			ports <- i
		}
		close(ports)
	}()

	// 处理worker
	for i := 1; i <= scanEndPort; i++ {
		res := <-results
		if res != 0 {
			openPorts = append(openPorts, res)
		} else {
			closePorts = append(closePorts, res)
		}
	}

	close(results)
	sort.Ints(openPorts)
	for _, port := range openPorts {
		fmt.Printf("%d opend\n", port)
	}
	elapsed := time.Since(start) / 1e9
	fmt.Printf("\n\n%d seconds\n", elapsed)

}
