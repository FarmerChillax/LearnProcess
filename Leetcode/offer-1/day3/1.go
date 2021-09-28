package day3

import "fmt"

func replaceSpace(s string) string {
	var res string
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			res = fmt.Sprintf("%s%s", res, "%20")
		} else {
			res = fmt.Sprintf("%s%c", res, s[i])
		}
	}
	return res
}

func ReplaceSpace(s string) string {
	return replaceSpace(s)
}
