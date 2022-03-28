package day2

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	res := len(nums)
	for left <= right {
		mid := (right-left)>>1 + left
		if target <= nums[mid] {
			res = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return res
}
