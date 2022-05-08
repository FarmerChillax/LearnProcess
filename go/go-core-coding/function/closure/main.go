package main

import "log"

func fa(a int) func(i int) int {
	return func(i int) int {
		log.Println(&a, a)
		a = a + i
		return a
	}
}

func main() {
	f := fa(1)
	g := fa(1)

	log.Println(f(1))
	// 多次调用一个闭包函数，f引用的还是同一个a(地址相同)
	log.Println(f(1))

	// f中的 a 不会影响 g 中的 a(因为地址不同)
	log.Println(g(1))
	log.Println(g(1))
}
