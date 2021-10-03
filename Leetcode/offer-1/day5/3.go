package day5

func firstUniqChar(s string) byte {
	cache := [26]int{}
	for i := 0; i < len(s); i++ {
		cache[s[i]-'a']++
	}

	for i := 0; i < len(s); i++ {
		if cache[s[i]-'a'] == 1 {
			return s[i]
		}
	}
	return ' '
}

func FirstUniqChar(s string) byte {
	return firstUniqChar(s)
}
