package main

import (
	"log"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Millisecond * 500)
	go func() {
		for tick := range ticker.C {
			log.Println("Ticker at:", tick)
		}
	}()
	time.Sleep(time.Millisecond * 10000)
	ticker.Stop()
	log.Println("ticker stopped")

}
