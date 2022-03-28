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
