package main

import "fmt"

func isMatch(s string, p string) bool {
	reChar := p[0]
	for sIndex, pIndex := 0, 0; sIndex < len(s); sIndex++ {
		// fmt.Println("当前re字符为:", reChar)
		if pIndex == len(p) {
			return false
		}
		if p[pIndex] != '*' {
			reChar = p[pIndex]
			pIndex++
		}
		if s[sIndex] != reChar {
			if reChar == '.' {
				continue
			}
			return false
		}

	}

	return true
}

func main() {
	result := isMatch("abdasd", "a*")
	fmt.Println(result)
}
