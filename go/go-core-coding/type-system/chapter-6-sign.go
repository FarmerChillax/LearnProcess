package main

import "fmt"

// 有名函数定义，函数名是 add
// add 类型是函数字面量类型 func (int, int) int
func add(a, b int) int {
	return a + b
}

// add 函数签名，实际上就是 add 的字面量类型
// func (int, int) int

// 匿名函数不能独立存在，常作为函数的参数、返回值，或者赋值给某个变量
// 匿名函数可以直接显示初始化
// 匿名函数的类型也是函数字面量类型 func (int, int) int
// > func (a, b int) int {
// > 	return a + b
// > }

// 新定义函数类型 ADD
// ADD 底层类型是函数字面量类型 func (int, int) int
type ADD func(int, int) int

// add 和 ADD 的底层类型相同，并且 add 是字面量类型
// 所以 add 可直接赋值给 ADD 类型的变量 g
var g ADD = add

func main() {
	f := func(a, b int) int {
		return a + b
	}

	g(1, 2)
	f(1, 2)

	fmt.Printf("%T\n", f)
	fmt.Printf("%T\n", add)

}
