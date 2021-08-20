package main

import "fmt"

func majorityElement(nums []int) int {
	var res int
	rMap := make(map[int]int)
	l := len(nums)

	for _, num := range nums {
		if _, ok := rMap[num]; !ok {
			rMap[num] = 0
		}
		rMap[num]++
	}

	for key, value := range rMap {
		if value > l/2 {
			res = key
		}
	}

	return res
}

func main() {
	raw := []int{3, 2, 3}
	res := majorityElement(raw)
	fmt.Println(res)
}
