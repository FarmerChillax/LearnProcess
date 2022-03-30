package day4

func mySqrt(x int) int {
	left, right := 0, x
	res := -1
	for left <= right {
		mid := (right-left)>>1 + left
		if mid*mid <= x {
			res = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return res
}
