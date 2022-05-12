package main

import (
	"fmt"
	"reflect"
)

type INT int

type A struct {
	a int
}

type B struct {
	b string
}

type Ita interface {
	String() string
}

func (b B) String() string {
	return b.b
}

func main() {
	var a INT = 12
	var b int = 14

	// 实参是具体类型，reflect.TypeOf返回是其静态类型
	ta := reflect.TypeOf(a)
	tb := reflect.TypeOf(b)

	// INT 和 int 是两个类型，二者并不相等
	if ta == tb {
		fmt.Println("ta == tb")
	} else {
		fmt.Println("ta != tb") // ta != tb
	}

	fmt.Println(ta.Name())
	fmt.Println(tb.Name())

	// 底层基础类型
	

}
