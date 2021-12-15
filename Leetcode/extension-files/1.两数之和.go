/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */
package main

// @lc code=start
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

// @lc code=end
