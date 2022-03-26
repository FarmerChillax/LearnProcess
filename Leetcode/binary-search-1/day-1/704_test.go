package day1

import (
	"log"
	"testing"
)

func HandleResult(res, expcet int) {
	if res != expcet {
		log.Fatalf("expcet: %d; result: %d\n", expcet, res)
	}
}

func Test704(t *testing.T) {
	input := []int{-1, 0, 3, 5, 9, 12}
	target := 9
	expcet := 4
	res := search(input, target)
	HandleResult(res, expcet)
	target = 2
	expcet = -1
	res = search(input, target)
	HandleResult(res, expcet)
}
