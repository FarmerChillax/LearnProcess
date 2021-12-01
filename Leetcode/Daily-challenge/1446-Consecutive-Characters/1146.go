package consecutivecharacters

// 计算最长非空子字符串长度
func maxPower(s string) int {
	result, tmp := 1, 1
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			tmp++
			if tmp > result {
				result = tmp
			}
		} else {
			tmp = 1
		}
	}
	return result
}
