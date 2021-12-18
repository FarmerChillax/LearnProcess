package wordpattern

import (
	"strings"
)

func wordPattern(pattern string, s string) bool {
	strList := strings.Split(s, " ")
	if len(strList) != len(pattern) {
		return false
	}
	tMap := map[byte][]int{}

	for i := 0; i < len(pattern); i++ {
		if _, ok := tMap[pattern[i]]; !ok {
			tMap[pattern[i]] = []int{i}
		} else {
			tMap[pattern[i]] = append(tMap[pattern[i]], i)
		}
	}

	sMap := map[string]bool{}

	for _, value := range tMap {
		key_word := strList[value[0]]
		for i := 1; i < len(value); i++ {
			if strList[value[i]] != key_word {
				return false
			}
		}
		sMap[key_word] = true
	}

	return len(sMap) == len(tMap)
}
