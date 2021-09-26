package sorts

// func main() {
// 	arr := []int{5, 4, 3, 2, 1, 7, 6, 8, 9}
// 	arr = selectionSort(arr)
// 	fmt.Println(arr)
// }

func findSmallest(arr []int) int {
	smallest := arr[0]
	smallest_index := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] < smallest {
			smallest = arr[i]
			smallest_index = i
		}
	}
	return smallest_index
}

func FindSmallest(arr []int) int { return findSmallest(arr) }

func selectionSort(arr []int) (result []int) {
	count := len(arr)
	for i := 0; i < count; i++ {
		smallest_index := findSmallest(arr)
		result = append(result, arr[smallest_index])
		arr = append(arr[:smallest_index], arr[smallest_index+1:]...)
	}
	return result
}

func SelectionSort(arr []int) []int { return selectionSort(arr) }
