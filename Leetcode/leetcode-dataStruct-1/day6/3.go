package main

import "fmt"

func isAnagram(s string, t string) bool {
	res := [26]int{}
	for i := 0; i < len(s); i++ {
		res[s[i]-'a']++
	}
	for i := 0; i < len(t); i++ {
		if res[t[i]-'a'] == 0 {
			return false
		}
		res[t[i]-'a']--
	}
	for _, v := range res {
		if v != 0 {
			return false
		}
	}
	return true
}

func main() {
	s, t := "anagram", "nagaram"
	res := isAnagram(s, t)
	fmt.Println(res)

}
