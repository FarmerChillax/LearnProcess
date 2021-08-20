package main

import "fmt"

func singleNumber(nums []int) int {
	var res int
	for _, num := range nums {
		res ^= num
	}
	return res
}

func main() {
	raw := []int{2, 2, 1}
	res := singleNumber(raw)
	fmt.Println(res)
}
