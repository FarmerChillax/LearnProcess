package day1

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := right - left/2 + left
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			// target在左边
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
