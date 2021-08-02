package main

import "fmt"

func matrixReshape(mat [][]int, r int, c int) [][]int {
	m, n := len(mat), len(mat[0])
	if m*n != r*c {
		return mat
	}
	res := make([][]int, r)
	for r := range res {
		res[r] = make([]int, c)
	}
	for i := 0; i < m*n; i++ {
		res[i/c][i%c] = mat[i/n][i%n]
	}
	return res
}

func main() {
	mat := [][]int{
		{1, 2},
		{3, 4},
	}
	r, c := 1, 4
	fmt.Println(matrixReshape(mat, r, c))
}
