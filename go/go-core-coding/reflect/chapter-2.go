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
	fmt.Println(ta.Kind().String())
	fmt.Println(tb.Kind().String())

	s1 := A{1}
	s2 := B{"tata"}

	// 实参是具体类型， reflect.TypeOf 返回的是其静态类型
	fmt.Println(reflect.TypeOf(s1).Name())
	fmt.Println(reflect.TypeOf(s2).Name())

	// Type 的 Kind() 方法返回的是基础类型，类型A 和 B 的底层基础类型都是 struct
	fmt.Println(reflect.TypeOf(s1).Kind().String())
	fmt.Println(reflect.TypeOf(s2).Kind().String())

	ita := new(Ita)
	var itb Ita = s2

	// 实参是未绑定具体变量的接口类型，reflect.TypeOf 返回的是接口类型本身
	// 也就是接口的静态类型

	fmt.Println(reflect.TypeOf(ita).Elem().Name())

	fmt.Println(reflect.TypeOf(ita).Elem().Kind().String())

	// 实参是绑定了具体变量的接口类型，reflect.TypeOf 返回的是绑定的具体类型
	// 也就是接口的动态类型
	fmt.Println(reflect.TypeOf(itb).Name())
	fmt.Println(reflect.TypeOf(itb).Kind().String())

}
