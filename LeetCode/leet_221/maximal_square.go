package leet_221

/*
Given an m x n binary matrix filled with 0's and 1's, find the largest square containing only 1's and return its area.
*/

//------------------------------------------------------------
// brutal force
func maximalSquare(matrix [][]byte) int {
	var res int

	if matrix == nil {
		return 0
	}
	m,n := len(matrix),len(matrix[0])

	for i := 0; i<m;i++ {
		v := make([]int,n)
		for j := i; j < m; j++ {
			for k := 0; k < n; k++ {
				if matrix[j][k] == '1' {
					v[k]++
				}
			}

			res = max(res, getSquareArea(v,j-i+1))
		}
	}

	return res
}

func getSquareArea(v []int, k int)  int{
	n := len(v)
	if n < k {
		return 0
	}

	var count int
	for i := 0; i < n;i++ {
		if v[i] != k {
			count = 0
		} else {
			count++
		}

		if count == k {
			return k*k
		}
	}

	return 0
}


//---------------------------------------
// dynamic programming
//定义dp[i][j]表示到达(i,j)位置所能组成的最大正方形的变长。对于一个点dp[i][j]，因为这个点是作为正方形的最右下角而存在的，所以这个点的状态改变
//只受到dp[i-1][j]、dp[i-1][j-1]和dp[i][j-1]的影响。如果matrix[i][j]=0，那么dp[i][j]必然为0；如果dp[i][j]为1，那么dp[i][j]=min(dp[i-1][j],
//	dp[i-1][j-1],dp[i][j-1])+1。为什么是三者中的最小值呢？其实我们可以这么想，如果这三个位置任何个位置是0，那么即便我们当前位置是1，dp[i][j]也只能是0；
//否则的话只能是这三个位置都非0，试着举一个简单的例子就可以看出来，我们需要的是交集，故而应该取最小值。
func maximalSquare2(matrix [][]byte) int {
	if matrix == nil || matrix[0] == nil {
		return 0
	}

	var res int
	m,n := len(matrix),len(matrix[0])
	dp := make([][]int,m)
	for i := range dp {
		dp[i] = make([]int,n)
	}
	for i := 0; i< m;i++ {
		for j := 0; j <n ;j++ {
			if i == 0 || j == 0 {
				dp[i][j] = int(matrix[i][j] - '0')
			} else if matrix[i][j] == '1' {
				dp[i][j] = min(dp[i-1][j-1],min(dp[i][j-1],dp[i-1][j])) + 1  // core step
			}

			res = max(res,dp[i][j])
		}
	}

	return res*res
}
//---------------------------------------
// util func
func max(a,b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a,b int) int {
	if a < b {
		return a
	}

	return b
}