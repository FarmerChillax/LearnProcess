package day3

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	ans := 0
	for i := 0; i < len(arr1); i++ {
		flag := true
		for j := 0; j < len(arr2); j++ {
			if abs(arr1[i]-arr2[j]) <= d {
				flag = false
			}
		}
		if flag {
			ans++
		}
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
