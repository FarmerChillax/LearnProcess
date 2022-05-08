package main

import "fmt"

// chain 函数的输入参数和输出参数类型相同, 都是 chan int 类型
// chain 函数的功能是将 chan 内的数据统一 +1
func chain(in chan int) chan int {
	out := make(chan int)
	go func() {
		for v := range in {
			out <- v + 1
		}
		close(out)
	}()

	return out
}

func main() {
	in := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			in <- i
		}
		close(in)
	}()

	// 连续调用 3 次 chan, 相当于 in 中的每个元素都 +3
	out := chain(chain(chain(in)))
	for v := range out {
		fmt.Println(v)
	}
}
