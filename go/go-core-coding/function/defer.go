package main

import "fmt"

func main() {
	defer fmt.Println("registry defer - 1")
	defer fmt.Println("registry defer - 2")
	defer fmt.Println("registry defer - 3")
}
