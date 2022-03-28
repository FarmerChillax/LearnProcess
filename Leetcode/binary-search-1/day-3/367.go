package day3

func isPerfectSquare(num int) bool {
	left, right := 0, num
	for left <= right {
		mid := (left + right) >> 1
		squ := mid * mid
		if num > squ {
			left = mid + 1
		} else if num < squ {
			right = mid - 1
		} else {
			return true
		}
	}
	return false
}
