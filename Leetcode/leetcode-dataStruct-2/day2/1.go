package main

import (
	"fmt"
	"sort"
)

func sortColors(nums []int) {
	sort.Ints(nums)
}

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)
	fmt.Println(nums)
}
