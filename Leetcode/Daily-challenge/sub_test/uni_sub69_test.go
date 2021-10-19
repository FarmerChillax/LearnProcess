package sub_test

import (
	"daily-challenge/sub"
	"fmt"
	"reflect"
	"testing"
)

func Sub69(t *testing.T) {
	arr := []int{3, 4, 5, 1}
	expect := 2
	res := sub.PeakIndexInMountainArray(arr)
	fmt.Println(res)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("result err: %v\n", res)
	}
}

func TestSub69(t *testing.T) {
	t.Run("sub=69", Sub69)
}
