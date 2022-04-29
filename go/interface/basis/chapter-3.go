package main

import "fmt"

type ReaderWriterTest interface {
	Read() ([]byte, error)
	Write([]byte) (int, error)
}

type RWT struct{}

func (this RWT) Read() ([]byte, error) {
	fmt.Println("rwt read...")
	return []byte("RWT Read"), nil
}

func (this RWT) Write(p []byte) (int, error) {
	fmt.Println("RWT Write:", p)
	return len(p), nil
}

type ReaderWriterDemo interface {
	Read() ([]byte, error)
	Write([]byte) (int, error)
}

type RWD struct{}

func (this RWD) Read() ([]byte, error) {
	fmt.Println("RWD read...")
	return []byte("RWD Read"), nil
}

func (this RWD) Write(p []byte) (int, error) {
	fmt.Println("RWD Write:", p)
	return len(p), nil
}

func main() {
	var rwTest ReaderWriterTest

	rwTest = RWD{}

	rwTest.Read()

}
