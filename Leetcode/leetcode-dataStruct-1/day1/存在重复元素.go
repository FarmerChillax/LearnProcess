package main

import "fmt"

func containsDuplicate(nums []int) bool {
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return false
		}
	}
	return true
}

//  存在重复元素
func main() {
	nums := []int{1, 2, 3, 4, 5, 6}
	res := containsDuplicate(nums)
	fmt.Println(res)
}
