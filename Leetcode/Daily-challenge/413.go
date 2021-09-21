package main

import "fmt"

func numberOfArithmeticSlices(nums []int) int {
	n := len(nums)
	var last, last_len, ans int
	for i := 1; i < n; i++ {
		if nums[i]-nums[i-1] == last {
			last_len++
		} else {
			last_len = 1
		}
		if last_len >= 2 {
			ans += last_len - 1
		}
		last = nums[i] - nums[i-1]
	}
	return ans
}

func main() {
	nums := []int{1}
	res := numberOfArithmeticSlices(nums)
	fmt.Println(res)
}
