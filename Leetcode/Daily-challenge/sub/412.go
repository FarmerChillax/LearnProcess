package sub

import (
	"strconv"
	"strings"
)

const (
	FIZZ = "Fizz"
	BUZZ = "Buzz"
)

func fizzBuzz(n int) (res []string) {
	for i := 1; i <= n; i++ {
		tmpStr := &strings.Builder{}
		if i%3 == 0 {
			tmpStr.WriteString(FIZZ)
		}
		if i%5 == 0 {
			tmpStr.WriteString(BUZZ)
		}
		if tmpStr.Len() == 0 {
			tmpStr.WriteString(strconv.Itoa(i))
		}
		res = append(res, tmpStr.String())
	}
	return res
}

func FizzBuzz(n int) []string {
	return fizzBuzz(n)
}
