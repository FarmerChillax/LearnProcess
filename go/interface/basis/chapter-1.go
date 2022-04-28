package main

import "fmt"

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

// Go 编译器会自动进行展开，最终展开成第 3 种格式
type ReaderWriter interface {
	Reader
	Writer
}

type ReaderWriter_2 interface {
	Reader
	Write(p []byte) (n int, err error)
}

type ReaderWriterFinally interface {
	Read(p []byte) (n int, err error)
	Write(p []byte) (n int, err error)
}

func main() {
	var i Reader
	fmt.Printf("%T\n", i) // <nil>
}
