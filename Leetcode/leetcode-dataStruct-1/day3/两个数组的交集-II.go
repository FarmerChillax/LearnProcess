package main

func intersect(nums1 []int, nums2 []int) []int {
	res := []int{}
	tmp := map[int]int{}

	if len(nums2) > len(nums1) {
		nums1, nums2 = nums2, nums1
	}

	for _, num := range nums1 {
		tmp[num]++
	}

	for _, item := range nums2 {
		if _, ok := tmp[item]; ok && tmp[item] != 0 {
			res = append(res, item)
			tmp[item]--
		}
	}

	return res
}

func main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	intersect(nums1, nums2)
}
