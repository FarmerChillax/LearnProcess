package day5

import "fmt"

func minArray(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	mid := len(numbers) >> 1
	fmt.Println(numbers[:mid], numbers[mid:])

	return 0
}

func MinArray(numbers []int) int {
	return minArray(numbers)
}
