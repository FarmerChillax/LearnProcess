package day4

func findRepeatNumber(nums []int) int {
	cacheMap := make(map[int]int, len(nums)-1)
	for _, value := range nums {
		if _, ok := cacheMap[value]; ok {
			return value
		}
		cacheMap[value] = 1
	}
	return 0
}

func FindRepeatNumber(nums []int) int {
	return findRepeatNumber(nums)
}

func findRepeatNumberV2(nums []int) int {
	for i := 0; i < len(nums); {
		if nums[i] == i {
			i++
		} else if nums[i] != nums[nums[i]] {
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		} else {
			return nums[i]
		}
	}
	return -1
}

func FindRepeatNumberV2(nums []int) int {
	return findRepeatNumberV2(nums)
}
