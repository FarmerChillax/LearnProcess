package day2

import "sort"

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	x, y := intervals[0][0], intervals[0][1]
	result := make([][]int, 0)
	for i := 1; i < len(intervals); i++ {
		if y >= intervals[i][0] {
			if y < intervals[i][1] {
				y = intervals[i][1]
			}
		} else {
			result = append(result, []int{x, y})
			x = intervals[i][0]
			y = intervals[i][1]
		}
	}
	result = append(result, []int{x, y})
	return result
}

// func merge(intervals [][]int) [][]int {
// 	result := [][]int{}
// 	tmp := [2]int{intervals[0][0], intervals[0][1]}
// 	if len(intervals) == 1 {
// 		return append(result, tmp[:])
// 	}

// 	for i := 1; i < len(intervals); i++ {
// 		if tmp[1] >= intervals[i][0] {
// 			tmp[1] = intervals[i][1]
// 		} else {
// 			result = append(result, []int{tmp[0], tmp[1]})
// 			tmp[0] = intervals[i][0]
// 			tmp[1] = intervals[i][1]
// 		}
// 		if i == len(intervals)-1 {
// 			result = append(result, []int{tmp[0], tmp[1]})
// 		}
// 	}
// 	return result
// }
