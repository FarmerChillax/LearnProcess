package main

import "fmt"

func main() {
	s1 := NewServer()
	s2 := NewServer(WithAddr("localhost"), WithPort(8080))
	s3 := NewServer(WithPort(8000))

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
}
