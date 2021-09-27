package sub

func lengthOfLastWord(s string) (res int) {
	i := len(s) - 1
	for s[i] == ' ' {
		i--
	}
	for i >= 0 && s[i] != ' ' {
		i--
		res++
	}
	return res
}

// func main() {
// 	fmt.Println(lengthOfLastWord("   fly me   to   the moon  "))
// }
