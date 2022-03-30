package day4

import (
	"binary-search-1/utils"
	"testing"
)

func Test69(t *testing.T) {
	input := 4
	except := 2
	res := mySqrt(input)
	utils.HandleResult(res, except)

	input = 8
	except = 2
	res = mySqrt(input)
	utils.HandleResult(res, except)
}

func Test744(t *testing.T) {
	input := []byte{'c', 'f', 'j'}
	target := 'a'
	except := 'c'
	res := nextGreatestLetter(input, byte(target))
	utils.HandleResult(res, except)

	input = []byte{'c', 'f', 'j'}
	target = 'c'
	except = 'f'
	res = nextGreatestLetter(input, byte(target))
	utils.HandleResult(res, except)

	input = []byte{'c', 'f', 'j'}
	target = 'd'
	except = 'f'
	res = nextGreatestLetter(input, byte(target))
	utils.HandleResult(res, except)
}
