package leet_877

type pair struct {
	first  int
	second int
}

// 2D DP
func stoneGame(piles []int) bool {
	length := len(piles)

	dp := make([][]pair, length)
	for i := 0; i < length; i++ {
		dp[i] = make([]pair, length)
	}

	// base condition
	// if there is only a pile of stone, and someone select first, the second get 0 stone.
	for i := 0; i < length; i++ {
		dp[i][i].first = piles[i]
		dp[i][i].second = 0
	}

	var j,left,right int
	for l := 2; l <= length; l++ {
		for i := 0; i <= length-l; i++ {
			j = l + i - 1
			left = piles[i] + dp[i+1][j].second
			right = piles[j] + dp[i][j-1].second

			if left > right {
				dp[i][j].first = left
				dp[i][j].second = dp[i+1][j].first
			} else {
				dp[i][j].first = right
				dp[i][j].second = dp[i][j-1].first
			}
		}
	}

	return dp[0][length-1].first > dp[0][length-1].second
}

// mathematical method
func stoneGameMath(piles []int) bool {
	return true
}

// 1D DP
func stoneGame1D(piles []int) bool {
	// TODO
	return true
}