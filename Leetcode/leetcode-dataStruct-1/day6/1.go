package main

import "fmt"

// my timeout code
// func firstUniqChar(s string) int {
// 	stringByte := []byte(s)

// 	for i := 0; i < len(stringByte); i++ {
// 		target := true
// 		for j := 0; j < len(stringByte); j++ {
// 			if stringByte[i] == stringByte[j] && i != j {
// 				target = false
// 			}
// 		}
// 		if target {
// 			return i
// 		}
// 	}
// 	return -1
// }

// my hashmap
// func firstUniqChar(s string) int {
// 	items := map[string]int{}
// 	for _, item := range s {
// 		items[string(item)]++
// 	}
// 	for i := 0; i < len(s); i++ {
// 		if items[string(s[i])] == 1 {
// 			return i
// 		}
// 	}
// 	return -1
// }

func firstUniqChar(s string) int {
	items := [26]int{}
	for i := 0; i < len(s); i++ {
		items[s[i]-'a']++
	}
	for i := 0; i < len(s); i++ {
		if items[s[i]-'a'] == 1 {
			return i
		}
	}
	return -1
}

func main() {
	var str string = "leetcode"
	res := firstUniqChar(str)
	fmt.Println(res)
	// var data []byte = []byte(str)
	// fmt.Println(data, len(data))
}
