package day1

func threeSum(nums []int) [][]int {
	rMap := make(map[int][]int)
	res := [][]int{}
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			// 0- tmp in key
			if _, ok := rMap[-nums[j]]; ok {
				rMap[-nums[j]] = append(rMap[-nums[j]], nums[j])
				res = append(res, rMap[-nums[j]])
			} else {
				rMap[nums[i]+nums[j]] = []int{i, j}
			}
		}
	}
	return res
}
