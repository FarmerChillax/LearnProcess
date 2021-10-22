package sub_test

import (
	"daily-challenge/sub"
	"log"
	"reflect"
	"testing"
)

func Sub229(t *testing.T) {
	n := []int{3, 2, 3}
	expectd := []int{3}
	res := sub.MajorityElement(n)
	if !reflect.DeepEqual(res, expectd) {
		log.Println("Not match,", res)
	}
}

func Sub229_2(t *testing.T) {
	n := []int{1, 1, 1, 3, 3, 2, 2, 2}
	expectd := []int{1, 2}
	res := sub.MajorityElement(n)
	if !reflect.DeepEqual(res, expectd) {
		log.Println("Not match,", res)
	}

}

func Sub229_3(t *testing.T) {
	n := []int{1}
	expectd := []int{1}
	res := sub.MajorityElement(n)
	if !reflect.DeepEqual(res, expectd) {
		log.Println("Not match,", res)
	}
}

func TestSub229(t *testing.T) {
	t.Run("work=229", Sub229)
	t.Run("work=229-2", Sub229_2)
	t.Run("work=229-3", Sub229_3)
}
