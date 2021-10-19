package sub

func peakIndexInMountainArray(arr []int) int {
	max := -1
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
		if arr[i] < max {
			return i - 1
		}
	}
	return 0
}

func PeakIndexInMountainArray(arr []int) int {
	return peakIndexInMountainArray(arr)
}

func peakIndexInMountainArrayV2(arr []int) int {
	low, high := 1, len(arr)-2
	for low < high {
		mid := (low + high) >> 1
		if arr[mid] < arr[mid+1] {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return low
}
func PeakIndexInMountainArrayV2(arr []int) int {
	return peakIndexInMountainArrayV2(arr)
}
