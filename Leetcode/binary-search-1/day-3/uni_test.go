package day3

import (
	"binary-search-1/utils"
	"testing"
)

func Test367(t *testing.T) {
	input := 16
	except := true
	res := isPerfectSquare(input)
	utils.HandleResult(res, except)

	input = 14
	except = false
	res = isPerfectSquare(input)
	utils.HandleResult(res, except)

}

func Test1385(t *testing.T) {
	input1 := []int{4, 5, 8}
	input2 := []int{10, 9, 1, 8}
	d := 2
	except := 2
	res := findTheDistanceValue(input1, input2, d)
	utils.HandleResult(res, except)

	input1 = []int{1, 4, 2, 3}
	input2 = []int{-4, -3, 6, 10, 20, 30}
	d = 3
	except = 2
	res = findTheDistanceValue(input1, input2, d)
	utils.HandleResult(res, except)

	input1 = []int{2, 1, 100, 3}
	input2 = []int{-5, -2, 10, -3, 7}
	d = 6
	except = 1
	res = findTheDistanceValue(input1, input2, d)
	utils.HandleResult(res, except)
}
