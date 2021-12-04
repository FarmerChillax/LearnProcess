package day3

import "fmt"

func generate(numRows int) [][]int {
	res := [][]int{}
	for i := 0; i < numRows; i++ {
		temp := make([]int, i+1)
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				temp[j] = 1
			} else {
				temp[j] = res[i-1][j-1] + res[i-1][j]
			}
		}
		res = append(res, temp)
	}
	return res
}

func getRow(rowIndex int) []int {
	res := generate(rowIndex + 1)
	fmt.Println(res[len(res)-1])
	return []int{}
}
