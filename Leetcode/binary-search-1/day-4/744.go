package day4

func nextGreatestLetter(letters []byte, target byte) byte {
	left, right := 0, len(letters)-1
	for left < right {
		mid := left + (right-left)/2
		if letters[mid] <= target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return letters[left%len(letters)]
}
