package sub_test

import (
	"daily-challenge/sub"
	"testing"
)

func Sub79(t *testing.T) {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	word := "ABCCED"
	expect := true
	res := sub.Sub79(board, word)
	if res != expect {
		t.Errorf("\nA bad request. word: %v\n", word)
	}
}

func Sub79_s2(t *testing.T) {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	word := "SEE"
	expect := true
	res := sub.Sub79(board, word)
	if res != expect {
		t.Errorf("\nA bad request. word: %v\n", word)
	}

}

func Sub79_s3(t *testing.T) {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	word := "ABCB"
	expect := false
	res := sub.Sub79(board, word)
	if res != expect {
		t.Errorf("\nA bad request. word: %v\n", word)
	}
}

func TestSub(t *testing.T) {
	t.Run("word=ABCCED", Sub79)
	t.Run("word=SEE", Sub79_s2)
	t.Run("word=ABCB", Sub79_s3)
}
