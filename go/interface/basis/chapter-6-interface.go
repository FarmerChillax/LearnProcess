package main

import "fmt"

type Inter interface {
	Ping()
	Pang()
}

type St struct{}

func (St) Ping() {
	fmt.Println("Ping")
}

func (*St) Pang() {
	fmt.Println("pang")
}

func main() {
	var st *St = nil
	var it Inter = st

	fmt.Printf("%p\n", st)
	fmt.Printf("%p\n", it)

	if it != nil {
		it.Pang()

		// it.Ping()
	}

}
