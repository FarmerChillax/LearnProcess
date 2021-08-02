package main

func setZeroes(matrix [][]int) {
	row, col := []int{}, []int{}
	for r := 0; r < len(matrix); r++ {
		for c := 0; c < len(matrix[r]); c++ {
			if matrix[r][c] == 0 {
				// 记录
				row = append(row, r)
				col = append(col, c)
			}
		}
	}

	for _, item := range row {
		for index := range matrix[item] {
			matrix[item][index] = 0
		}
	}

	for _, item := range col {
		for _, value := range matrix {
			value[item] = 0
		}
	}
	// fmt.Println(matrix)
}

func main() {
	matrix := [][]int{
		{0, 1, 2, 0},
		{3, 4, 5, 2},
		{1, 3, 1, 5},
	}

	setZeroes(matrix)

}
