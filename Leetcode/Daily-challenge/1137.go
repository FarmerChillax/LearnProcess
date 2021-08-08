package main

import "fmt"

func main() {
	res := tribonacci(25)
	fmt.Println(res)
}

// ...
func tribonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n <= 2 {
		return 1
	}
	var res int
	t1, t2, t3 := 0, 0, 1
	for i := 3; i <= n+1; i++ {
		res = t1 + t2 + t3
		t1, t2, t3 = t2, t3, res
	}
	return res
}
