package day8

// 1 -> 1
// 2 -> 2
// 3 -> 1 + 2
// 4 -> 5
func numWays(n int) int {
	a, b := 1, 1
	for i := 0; i < n; i++ {
		sum := (a + b) % 1000000007
		a, b = b, sum
	}
	return a % 1000000007
}

func NumWays(n int) int {
	return numWays(n)
}
