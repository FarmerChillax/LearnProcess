package main

import "fmt"

// 组合
type X struct {
	a int
}

func (x X) Print() {
	fmt.Println("In X:", x.a)
}

func (x X) PrintX() {
	fmt.Println("In X:", x.a)
}

type Y struct {
	X
	b int
}

func (y Y) Print() {
	fmt.Println("In Y:", y.a)
}

type Z struct {
	Y
	c int
}

func (z Z) Print() {
	fmt.Println("In Z:", z.a)

	z.Y.Print()
	z.Y.X.Print()
}

// 方法集
type A struct {
	*X
}

func (x X) Get() int {
	return x.a
}

func (x *X) Set(a int) {
	x.a = a
}

func main() {
	x := X{a: 1}
	y := Y{
		X: x,
		b: 2,
	}
	z := Z{
		Y: y,
		c: 3,
	}

	fmt.Println(z.a, z.Y.a, z.Y.X.a)

	z = Z{}

	z.a = 2
	fmt.Println(z.a, z.Y.a, z.Y.X.a)

	// 从外向内查找，因此先找到 Z 的方法
	z.Print()

	// 从外向内查找, 最后找到的是 X 的 XPrint() 方法
	x.PrintX()
	z.Y.PrintX()

	// 方法集测试
	// 编译器做了自动转换
	y.Set(3)
	fmt.Println(y.Get())

	// 为了不让编译器自动转换，使用方法表达式的调用方式
	(*Y).Set(&y, 4)
	fmt.Println(y.Get())

	a := A{
		X: &x,
	}

	// 按照嵌套字段的方法集的规则
	// z 内嵌字段 *X, 所以 type Z 和 type *Z 方法集都包含类型 X 定义的方法
	// Get 和 Set
	// 为了不让编译器自动转换，仍然使用方法表达式的调用方式
	A.Set(a, 5)
	fmt.Println(a.Get())

	(*A).Set(&a, 6)
	fmt.Println(a.Get())

}
