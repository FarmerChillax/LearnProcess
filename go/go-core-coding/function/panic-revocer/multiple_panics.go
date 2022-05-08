package main

import "log"

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()

	defer func() {
		panic("first defer panic")
	}()

	defer func() {
		panic("second defer panic")
	}()

	panic("main body panic")

}
