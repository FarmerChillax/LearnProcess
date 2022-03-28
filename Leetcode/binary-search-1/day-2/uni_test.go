package day2

import (
	"binary-search-1/utils"
	"log"
	"testing"
)

func Test35(t *testing.T) {
	input := []int{1, 3, 5, 6}
	target := 5
	except := 2
	res := searchInsert(input, target)
	utils.HandleResult(res, except)
	log.Println("test -- 2")
	target = 2
	except = 1
	res = searchInsert(input, target)
	utils.HandleResult(res, except)

}

func Test852(t *testing.T) {
	input := []int{0, 1, 23, 33, 100, 2, 1, 0}
	except := 4
	res := peakIndexInMountainArray(input)
	utils.HandleResult(res, except)

	input = []int{0, 10, 5, 2}
	except = 1
	res = peakIndexInMountainArray(input)
	utils.HandleResult(res, except)

	input = []int{3, 4, 5, 1}
	except = 2
	res = peakIndexInMountainArray(input)
	utils.HandleResult(res, except)

	input = []int{24, 69, 100, 99, 79, 78, 67, 36, 26, 19}
	except = 2
	res = peakIndexInMountainArray(input)
	utils.HandleResult(res, except)
}
