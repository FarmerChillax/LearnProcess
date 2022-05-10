package main

import (
	"math/rand"
	"os"
	"runtime/pprof"
	"time"
)

func generate(n int) []int {

	rand.Seed(time.Now().UnixNano())
	nums := make([]int, 0)
	for i := 0; i < n; i++ {
		nums = append(nums, rand.Int())
	}

	return nums

}

func bubbleSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := 1; j < len(nums)-i; j++ {
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			}
		}
	}
}

func main() {
	// 添加 pprof
	pprof.StartCPUProfile(os.Stdout)
	defer pprof.StopCPUProfile()

	// main
	n := 10
	for i := 0; i < 5; i++ {
		nums := generate(n)
		bubbleSort(nums)
		n *= 10
	}
}
