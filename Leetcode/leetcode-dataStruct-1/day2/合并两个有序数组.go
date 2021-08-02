package main

import "fmt"

// 方法一
// func merge(nums1 []int, m int, nums2 []int, n int) {
// 	nums1 = append(nums1[:m], nums2...)
// 	sort.Ints(nums1)
// 	fmt.Println(nums1)
// }

//方法二
func merge(nums1 []int, m int, nums2 []int, n int) {
	// 存放结果
	res := make([]int, 0, m+n)
	// 两个数组的指针
	point1, point2 := 0, 0

	for {
		// 数组一循环完, 将数组二剩余元素添加到结果里面
		if point1 == m {
			res = append(res, nums2[point2:]...)
			break
		}
		// 与上面同理
		if point2 == n {
			res = append(res, nums1[point1:]...)
			break
		}
		if nums1[point1] < nums2[point2] {
			res = append(res, nums1[point1])
			point1++
		} else {
			res = append(res, nums2[point2])
			point2++
		}
	}
	fmt.Println(res)
	fmt.Println(nums1)
	nums1 = res
	fmt.Println(nums1)
}

// 合并两个有序数组
func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	m, n := 3, 3
	merge(nums1, m, nums2, n)
}
