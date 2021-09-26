package sorts

func NewArray() []int {
	return []int{5, 8, 6, 3, 9, 2, 1, 7}
}

// func main() {
// 	targetList := NewArray()
// 	bubbleSort(targetList)
// 	fmt.Println(targetList)
// 	targetList = NewArray()
// 	optimizationBubbleSort(targetList)
// 	fmt.Println(targetList)
// }

func bubbleSort(array []int) {
	var flag bool = false
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				tmp := array[j]
				array[j] = array[j+1]
				array[j+1] = tmp
				flag = true
			}
		}
		if !flag {
			break
		}
	}
}

func BubbleSort(array []int) {
	bubbleSort(array)
}

func optimizationBubbleSort(array []int) {
	var is_sort bool = true
	var lastExchangeIndex int
	sort_border := len(array) - 1
	for i := 0; i < len(array); i++ {
		is_sort = true
		for j := 0; j < sort_border; j++ {
			if array[j] > array[j+1] {
				tmp := array[j]
				array[j] = array[j+1]
				array[j+1] = tmp
				is_sort = false
				lastExchangeIndex = j
			}
		}
		sort_border = lastExchangeIndex
		if is_sort {
			break
		}
	}
}

func OptimizationBubbleSort(array []int) {
	optimizationBubbleSort(array)
}
