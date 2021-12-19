package findthetownjudge

func findJudge(n int, trust [][]int) int {
	tmpTrust := make([]int, n)
	isTrustOther := make([]bool, n)

	for _, item := range trust {
		owner := item[0]
		trustTarget := item[1]
		// 被人信任，则提高信任度
		tmpTrust[trustTarget-1]++
		// 信任别人则不可能是judge
		isTrustOther[owner-1] = true
	}

	for i := 0; i < n; i++ {
		if tmpTrust[i] == n-1 && !isTrustOther[i] {
			return i + 1
		}
	}

	return -1
}
