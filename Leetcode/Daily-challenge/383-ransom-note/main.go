package ransomnote

func canConstruct(ransomNote string, magazine string) bool {
	// conMap := make(map[byte]int, len(magazine))
	conMap := [26]int{}

	for i := 0; i < len(magazine); i++ {
		conMap[magazine[i]-'a']++
	}
	for i := 0; i < len(ransomNote); i++ {
		if conMap[ransomNote[i]-'a'] == 0 {
			return false
		}
		conMap[ransomNote[i]-'a']--
	}
	return true

}
