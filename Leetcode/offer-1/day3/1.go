package day3

import (
	"net/url"
	"strings"
)

func replaceSpaceV1(s string) string {
	// for i := 0; i < len(s); i++ {
	// 	if s[i] == ' ' {
	// 		s = s[:i] + "%20" + s[i+1:]
	// 	}
	// }
	return url.PathEscape(s)
	// return s
}

func replaceSpaceV2(s string) string {
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			s = s[:i] + "%20" + s[i+1:]
		}
	}
	return s
}

func replaceSpaceV3(s string) string {
	var b strings.Builder
	b.Grow(len(s))
	last := 0
	for k, v := range s {
		if v == ' ' {
			// 当遍历到了空格的时候，把0到空格下标-1的切片写入
			// 然后把 “%20”写入，
			// 然后下标偏移到 %20 +1位置
			b.WriteString(s[last:k])
			b.WriteString("%20")
			last = k + 1
		}
	}
	// 把%20 到末尾的所有字符全部写入到 b
	b.WriteString(s[last:])
	return b.String()
}

func ReplaceSpaceV1(s string) string {
	return replaceSpaceV1(s)
}

func ReplaceSpaceV2(s string) string {
	return replaceSpaceV2(s)
}

func ReplaceSpaceV3(s string) string {
	return replaceSpaceV3(s)
}
