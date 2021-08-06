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
		address := fmt.Sprintf("xiaotao2333.top:%d", p)
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
	ports := make(chan int, 100)
	results := make(chan int, 100)
	var openPorts []int
	var closePorts []int

	// 生成worker
	for i := 0; i < cap(ports); i++ {
		go worker(ports, results)
	}

	// 启动worker
	go func() {
		for i := 1; i < 65535; i++ {
			ports <- i
		}
	}()

	// 处理worker
	for i := 1; i < 65535; i++ {
		port := <-results
		if port != 0 {
			openPorts = append(openPorts, port)
		} else {
			closePorts = append(closePorts, port)
		}
	}
	// wg.Wait()
	close(ports)
	close(results)
	sort.Ints(openPorts)
	for _, port := range closePorts {
		fmt.Printf("%d closed\n", port)
	}
	for _, port := range openPorts {
		fmt.Printf("%d opend\n", port)
	}
	elapsed := time.Since(start) / 1e9
	fmt.Printf("\n\n%d seconds\n", elapsed)

}
