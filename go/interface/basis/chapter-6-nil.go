package main

import "fmt"

type Inter interface {
	Ping()
	Pang()
}

type St struct {
	Name string
}

func (St) Ping() {
	fmt.Println("ping")
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

		// 下面的语句会导致 panic
		// 方法转换为函数调用, 第一个参数是 St 类型, 由于 *St 是 nil, 无法获取指针所指的
		// 对象值, 所以导致 panic
		it.Ping()
	}

}
