package sub

import (
	"math"
)

func majorityElement(nums []int) (ans []int) {
	max, min := getMaxMin(nums)
	// 计数排序
	temp := make([]int, max-min+1)
	for i := 0; i < len(nums); i++ {
		temp[nums[i]-min]++
	}

	target := len(nums) / 3
	// 处理结果
	for i := 0; i < len(temp); i++ {
		if temp[i] > target {
			ans = append(ans, i+min)
		}
	}

	return ans
}

func getMaxMin(nums []int) (max, min int) {
	max = math.MinInt64
	min = math.MaxInt64

	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
		if nums[i] < min {
			min = nums[i]
		}
	}
	return max, min
}

func MajorityElement(nums []int) []int {
	return majorityElement(nums)
}
