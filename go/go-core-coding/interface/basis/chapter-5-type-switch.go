package main

import (
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.OpenFile("notes.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var i io.Reader = f

	switch v := i.(type) {

	// v 是 io.ReadWriter 接口类型, 所以可以调用 Write 方法
	case io.ReadWriter:
		v.Write([]byte("io.ReadWriter\n"))

	// 由于上一个 case 已匹配, 计算这个 case 也匹配, 也不会走到这里
	case *os.File:
		v.Write([]byte("*os.File\n"))
		v.Sync()
	default:
		return
	}
}
