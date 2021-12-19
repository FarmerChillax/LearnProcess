package findthetownjudge

import (
	"log"
	"testing"
)

func TestFindJudge(t *testing.T) {
	n := 2
	trust := [][]int{{1, 2}}
	excepted := 2
	result := findJudge(n, trust)
	if result != excepted {
		log.Fatalf("not match, result: %v; excepted: %v\n", result, excepted)
	}
}

func TestFindJudge2(t *testing.T) {
	n := 3
	trust := [][]int{{1, 3}, {2, 3}}
	excepted := 3
	result := findJudge(n, trust)
	if result != excepted {
		log.Fatalf("not match, result: %v; excepted: %v\n", result, excepted)
	}
}

func TestFindJudge3(t *testing.T) {
	n := 3
	trust := [][]int{{1, 3}, {2, 3}, {3, 1}}
	excepted := -1
	result := findJudge(n, trust)
	if result != excepted {
		log.Fatalf("not match, result: %v; excepted: %v\n", result, excepted)
	}
}

func TestFindJudge4(t *testing.T) {
	n := 3
	trust := [][]int{{1, 2}, {2, 3}}
	excepted := -1
	result := findJudge(n, trust)
	if result != excepted {
		log.Fatalf("not match, result: %v; excepted: %v\n", result, excepted)
	}
}

func TestFindJudge5(t *testing.T) {
	n := 4
	trust := [][]int{{1, 3}, {1, 4}, {2, 3}, {2, 4}, {4, 3}}
	excepted := 3
	result := findJudge(n, trust)
	if result != excepted {
		log.Fatalf("not match, result: %v; excepted: %v\n", result, excepted)
	}
}
