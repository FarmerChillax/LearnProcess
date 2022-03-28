package day2

func peakIndexInMountainArray(arr []int) int {
	left, right := 0, len(arr)-1
	mid := left
	for left < right {
		mid = (left + right) >> 1
		if arr[mid-1] < arr[mid] && arr[mid] > arr[mid+1] {
			break
		}
		if arr[mid] < arr[mid+1] {
			// 往右
			left = mid
		} else if arr[mid] > arr[mid+1] {
			right = mid
		}
	}
	return mid
}
