package main

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

// demo
// func generate(numRows int) [][]int {
// 	ans := make([][]int, numRows)
// 	for i := range ans {
// 		ans[i] = make([]int, i+1)
// 		ans[i][0] = 1
// 		ans[i][i] = 1
// 		for j := 1; j < i; j++ {
// 			ans[i][j] = ans[i-1][j] + ans[i-1][j-1]
// 		}
// 	}
// 	return ans
// }

func main() {
	res := generate(5)
	fmt.Println(res)
}
