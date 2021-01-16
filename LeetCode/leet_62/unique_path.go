package leet_62

/*
	A robot is located at the top-left corner of a m x n grid (marked 'Start' in the diagram below).
	The robot can only move either down or right at any point in time. The robot is trying to reach the bottom-right corner of the grid (marked 'Finish' in the diagram below).
	How many possible unique paths are there?
*/

// dp[i][j] represent the number of unique path to point(i,j)
// time complexity: O(m*n)
// space complexity: O(m*n)
func uniquePathsV1(m int, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	//base case,dp[0][j] = dp[i][0] = 1
	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}
	for i := 0; i < n; i++ {
		dp[0][i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m-1][n-1]
}

// every time we update dp[i][j], we just need dp[i-1][j](the previous row) and dp[i][j-1](the current row),
// hence we just need two rows to store them, and cutting space to O(n)
// time complexity: O(m*n)
// space complexity: O(n)
func uniquePathsV2(m int, n int) int {
	pre := make([]int, m)
	for i := range pre {
		pre[i] = 1
	}

	cur := make([]int, n)
	for i := range cur {
		cur[i] = 1
	}

	//TODO  figure out how to optimize this code, also check the discussion in leetcode
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			cur[j] = pre[j] + cur[i-1]
		}
		swap(&pre, &cur)
	}

	return cur[n-1]
}

func swap(a, b *[]int) {
	a, b = b, a
}

// 使用排列组合的方法。 假设一个m行n列的数组，一个点要从（0，0）移动到（m-1，n-1），意味着他需要向下走m-1步，向右走n-1步，一共需要走m+n-2步。
// 也就是所我们只需要从m+n-2步中选出m-1步向下走和n-1步向右走的方法并且进行一个排列组合即可
func uniquePathsWithPermutation(m int, n int) int {
	N := m + n - 2
	k := m - 1
	res := 1
	for i := 1; i <= k; i++ {
		res = res * (N - k + i) / i
	}

	return res
}

func uniquePaths(m int, n int) int {
	//return uniquePathsV2(m,n)
	return uniquePathsWithPermutation(m, n)
}
