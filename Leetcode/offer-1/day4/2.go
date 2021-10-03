package day4

import "sort"

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		if nums[left] == target && nums[right] == target {
			return right - left + 1
		}
		if nums[left] != target {
			left++
		}
		if nums[right] != target {
			right--
		}
	}
	return 0
}

func Search(nums []int, target int) int {
	return search(nums, target)
}

func searchv2(nums []int, target int) int {
	leftmost := sort.SearchInts(nums, target)
	if leftmost == len(nums) || nums[leftmost] != target {
		return 0
	}
	rightmost := sort.SearchInts(nums, target+1) - 1
	return rightmost - leftmost + 1
}

func SearchV2(nums []int, target int) int {
	return searchv2(nums, target)
}
