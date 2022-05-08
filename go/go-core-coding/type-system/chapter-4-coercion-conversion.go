package main

import "fmt"

type Map map[string]string

func (m Map) Print() {
	for _, key := range m {
		fmt.Println(key)
	}
}

type iMap Map

func (m iMap) Print() {
	for _, key := range m {
		fmt.Println(key)
	}
}

type slice []int

func (s slice) Print() {
	for _, v := range s {
		fmt.Println(v)
	}
}

func main() {
	mp := make(map[string]string, 10)
	mp["hi"] = "farmer"

	// mp 与 ma 有相同的底层类型 -> map[string]string, 并且 mp 是未命名类型
	// 所以 mp 可以直接赋值给 ma
	var ma Map = mp

	ma.Print()

	// im 与 ma 虽然底层类型都是 -> map[string]string, 但他们都不是未命名类型
	// 不能直接赋值, 比如下面的语句就不能通过编译
	// var im iMap = ma
	// 但可以强制进行类型转换
	var im iMap = (iMap)(ma)
	im.Print()

	// Map 实现了 Print(), 所以其 可以赋值给接口类型变量
	var i interface {
		Print()
	} = ma

	i.Print()

	s1 := []int{1, 2, 3}
	var s2 slice

	s2 = s1
	s2.Print()

}
