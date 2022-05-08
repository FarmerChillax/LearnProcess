package main

import (
	"log"
	"time"
)

func do() {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()
	go da()
	go db()
	time.Sleep(3 * time.Second)
}

func da() {
	panic("panic da")
	for i := 0; i < 10; i++ {
		log.Println(i)
	}
}

func db() {
	for i := 0; i < 10; i++ {
		log.Println(i)
	}
}

func main() {
	do()
}
