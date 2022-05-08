package main

import "fmt"

func GetYangHuiTriangleNextLine(inArr []int) []int {
	var out []int
	var i int
	arrLen := len(inArr)
	out = append(out, 1)
	if arrLen == 0 {
		return out
	}
	for i = 0; i < arrLen-1; i++ {
		out = append(out, inArr[i]+inArr[i+1])
	}
	out = append(out, 1)
	return out
}
func main() {
	nums := []int{}
	var i int
	for i = 0; i < 10; i++ {
		nums = GetYangHuiTriangleNextLine(nums)
		fmt.Println(nums)
	}
}
