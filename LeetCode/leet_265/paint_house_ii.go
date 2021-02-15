package leet_265

/*
There are a row of n houses, each house can be painted with one of the k colors. The cost of painting each house with a certain color is different. You have to paint all the houses such that no two adjacent houses have the same color.

The cost of painting each house with a certain color is represented by a n x k cost matrix. For example, costs[0][0] is the cost of painting house 0 with color 0; costs[1][2] is the cost of painting house 1 with color 2, and so on... Find the minimum cost to paint all houses.

Note:
All costs are positive integers.

Example:

Input: [[1,5,3],[2,9,4]]
Output: 5
Explanation: Paint house 0 into color 0, paint house 1 into color 2. Minimum cost: 1 + 4 = 5;
             Or paint house 0 into color 2, paint house 1 into color 0. Minimum cost: 3 + 2 = 5.
Follow up:
Could you solve it in O(nk) runtime?
*/

// 与之前的paint house不同，这次需要使用k个颜色。
// DP思路。在找不同的颜色的最小值不是遍历所有不同颜色，而是用min1和min2里记录之前房子的最小和第二小的花费的颜色，如果当前房子颜色和min1相同
//那么用min2对应的值计算，反之用min1对应的值。
func minCostII(costs [][]int) int {
	if len(costs) == 0 || len(costs[0]) == 0 {
		return 0
	}

	n,k := len(costs),len(costs[0])
	dp := make([][]int,n)
	for i := range dp {
		dp[i] = make([]int,k)
		copy(dp[i],costs[i])
	}

	min1,min2 := -1,-1
	for i := 0; i< n;i++ {
		last1,last2 := min1,min2  // 需要找到当前房子的最小和次小值，故而需要保存历史值
		min1,min2 = -1,-1
		for j := 0; j< k;j++ {
			//if (j != last1) {
			//	dp[i][j] += last1 < 0 ? 0 : dp[i - 1][last1];
			//} else {
			//	dp[i][j] += last2 < 0 ? 0 : dp[i - 1][last2];
			//}
			//if (min1 < 0 || dp[i][j] < dp[i][min1]) {
			//	min2 = min1; min1 = j;
			//} else if (min2 < 0 || dp[i][j] < dp[i][min2]) {
			//	min2 = j;
			//}
			if j != last1 {
				if last1 < 0 {
					dp[i][j] += 0
				} else {
					dp[i][j] += costs[i-1][last1]
				}
			} else {
				if last2 < 0 {
					dp[i][j] += 0
				} else {
					dp[i][j] += costs[i-1][last2]
				}
			}

			if min1 < 0 || dp[i][j] < dp[i][min1] {
				min2=min1
				min1=j
			} else if min2 < 0 || dp[i][j] < dp[i][min2] {
				min2 = j
			}
		}
	}

	return dp[n-1][min1]
}