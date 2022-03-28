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
