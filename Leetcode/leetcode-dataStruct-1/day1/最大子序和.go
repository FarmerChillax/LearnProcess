package main

import "fmt"

func maxSubArray(nums []int) int {
	max := nums[0]
	tmp := max
	for i := 1; i < len(nums); i++ {
		if tmp >= max {
			max = tmp
		}
		if tmp < 0 {
			tmp = 0
		}
		tmp += nums[i]
	}
	return max
}

// 最大子序和
func main() {
	nums := []int{-2, 1}
	res := maxSubArray(nums)
	fmt.Println(res)
}
