package sub_test

import (
	"daily-challenge/sub"
	"reflect"
	"testing"
)

func Sub412(t *testing.T) {
	n := 15
	expect := []string{
		"1",
		"2",
		"Fizz",
		"4",
		"Buzz",
		"Fizz",
		"7",
		"8",
		"Fizz",
		"Buzz",
		"11",
		"Fizz",
		"13",
		"14",
		"FizzBuzz",
	}
	res := sub.FizzBuzz(n)
	// fmt.Println(res)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("result err: %v\n", res)
	}
}

func TestSub412(t *testing.T) {
	t.Run("sub=412", Sub412)
}
