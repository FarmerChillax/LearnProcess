package sub

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if secIndex, ok := numMap[target-nums[i]]; ok {
			return []int{i, secIndex}
		}
		numMap[nums[i]] = i
	}

	return nil
}

func TwoSum(nums []int, target int) []int {
	return twoSum(nums, target)
}
