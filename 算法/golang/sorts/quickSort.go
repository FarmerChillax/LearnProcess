package sorts

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	pivot := arr[0]
	var left, right []int

	for _, item := range arr[1:] {
		if item <= pivot {
			left = append(left, item)
		} else {
			right = append(right, item)
		}
	}
	return append(quickSort(left),
		append([]int{pivot}, quickSort(right)...)...)
}

func QuickSort(arr []int) []int {
	return quickSort(arr)
}

// func main() {
// 	arr := []int{12, 87, 1, 66, 38, 213, 126, 1233, 6542, 456, 973, 102, 3}
// 	result := quickSort(arr)
// 	fmt.Println(result)
// }
