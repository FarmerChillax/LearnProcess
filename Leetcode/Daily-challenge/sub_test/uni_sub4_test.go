package sub_test

import (
	"daily-challenge/sub"
	"testing"
)

func Sub1(t *testing.T) {
	n1 := []int{1, 3, 6}
	n2 := []int{2, 4, 5}
	_ = sub.FindMedianSortedArrays(n1, n2)
	// fmt.Println(res)
}

func TestSub4(t *testing.T) {
	t.Run("work=1", Sub1)
}
