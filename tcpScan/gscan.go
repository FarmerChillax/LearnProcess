package main

// func main() {
// 	start := time.Now()
// 	var wg sync.WaitGroup
// 	for i := 21; i < 1024; i++ {
// 		wg.Add(1)
// 		go func(port int) {
// 			defer wg.Done()
// 			address := fmt.Sprintf("192.168.2.123:%d", port)
// 			conn, err := net.Dial("tcp", address)
// 			if err != nil {
// 				log.Printf("%s 关闭了\n", address)
// 				return
// 			}
// 			conn.Close()
// 			log.Printf("%s 打开了！!!!!!!!\n", address)
// 		}(i)
// 	}
// 	wg.Wait()
// 	elapsed := time.Since(start) / 1e9
// 	log.Printf("\n\n%d seconds", elapsed)
// }

// func main() {
// 	for i := 20; i < 1024; i++ {
// 		address := fmt.Sprintf("farmer233.top:%d", i)
// 		fmt.Println(address)
// 		conn, err := net.Dial("tcp", address)
// 		if err != nil {
// 			log.Printf("%s 关闭了\n", address)
// 		} else {
// 			conn.Close()
// 			log.Printf("%s 打开了！\n", address)
// 		}
// 	}
// }

// func main() {
// 	var wg sync.WaitGroup
// 	var mutex sync.Mutex
// 	ports := make([]int, 0)
// 	for i := 1; i < 65535; i++ {
// 		wg.Add(1)
// 		go func(port int) {
// 			defer wg.Done()
// 			conn, err := net.DialTimeout("tcp", fmt.Sprintf("192.168.2.123:%d", port), time.Second)
// 			if err != nil {
// 				log.Printf("Error: %v.Port:[%d]\n", err, port)
// 			} else {
// 				conn.Close()
// 				log.Printf("Connection success.port:[%d]\n", port)
// 				mutex.Lock()
// 				ports = append(ports, port)
// 				mutex.Unlock()
// 			}
// 		}(i)
// 	}
// 	wg.Wait()
// 	fmt.Printf("Opened ports: %v\n", ports)
// }
