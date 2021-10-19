package sub

type WordDictionary struct {
	Val map[int][]string
}

func Constructor() WordDictionary {
	return WordDictionary{
		Val: make(map[int][]string),
	}
}

func (this *WordDictionary) AddWord(word string) {
	this.Val[len(word)] = append(this.Val[len(word)], word)
}

func (this *WordDictionary) Search(word string) bool {
	if words, ok := this.Val[len(word)]; ok {
		for _, val := range words {
			if checkWord(word, val) {
				return true
			}
		}
	}
	return false
}

func checkWord(word, comp string) bool {
	if len(word) != len(comp) {
		return false
	}

	for i := 0; i < len(word); i++ {
		if word[i] != comp[i] && word[i] != '.' {
			return false
		}
	}

	return true
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
