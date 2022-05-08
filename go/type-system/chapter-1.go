package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	a := struct {
		name string
		age  int
	}{name: "Farmer", age: 22}

	fmt.Printf("%T\n", a)
	fmt.Printf("%v\n", a)

	b := Person{name: "Farmer.Chillax", age: 23}

	fmt.Printf("%T\n", b)
	fmt.Printf("%v\n", b)

}
