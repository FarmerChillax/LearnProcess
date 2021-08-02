package main

import "fmt"

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	// resolve 1
	// var res []int

	// for index, item := range nums {
	// 	numMap[item] = index
	// }

	// for i := 0; i < len(nums); i++ {
	// 	tmp := target - nums[i]
	// 	if tmpIndex, ok := numMap[tmp]; ok && numMap[tmp] != i {
	// 		return []int{i, tmpIndex}
	// 	}
	// }

	// resolve 2
	// for index, item := range nums {
	// 	if secIndex, ok := numMap[target-item]; ok {
	// 		return []int{index, secIndex}
	// 	}
	// 	numMap[item] = index
	// }

	for i := 0; i < len(nums); i++ {
		if secIndex, ok := numMap[target-nums[i]]; ok {
			return []int{i, secIndex}
		}
		numMap[nums[i]] = i
	}

	return nil
}

// 两数之和
func main() {
	nums := []int{3, 2, 4}
	target := 6
	res := twoSum(nums, target)
	fmt.Println(res)
}
