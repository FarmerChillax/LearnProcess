package day5

// 左边大于右边时，右边的元素为目标
func minArray(numbers []int) int {
	low := 0
	high := len(numbers) - 1
	for low < high {
		mid := low + (high-low)>>1
		if numbers[mid] < numbers[high] {
			high = mid
		} else if numbers[mid] > numbers[high] {
			low = mid + 1
		} else {
			high--
		}
	}
	return numbers[low]
}

func MinArray(numbers []int) int {
	return minArray(numbers)
}
