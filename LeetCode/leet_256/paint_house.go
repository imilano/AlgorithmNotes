package leet_256

import (
	"math"
)

/*
There are a row of n houses, each house can be painted with one of the three colors: red, blue or green. The cost of painting each house with a certain color is different. You have to paint all the houses such that no two adjacent houses have the same color.

The cost of painting each house with a certain color is represented by a n x 3 cost matrix. For example, costs[0][0] is the cost of painting house 0 with color red; costs[1][2] is the cost of painting house 1 with color green, and so on... Find the minimum cost to paint all houses.

Note:
All costs are positive integers.

Example:

Input: [[17,2,17],[16,16,5],[14,3,19]]
Output: 10
Explanation: Paint house 0 into blue, paint house 1 into green, paint house 2 into blue.
             Minimum cost: 2 + 5 + 3 = 10.
*/

// n个房子，每个房子可以涂红绿蓝三个颜色，每个房子涂不同的颜色会有不同的代价，每两个相邻的房子不能同色，求最小代价.
// 递归，比较好理解，不再赘述
func minCost(costs [][]int) int {
	res := math.MaxInt32
	helper(0, -1, 0, &res, costs)
	return res
}

func helper(row, preSelect, curCost int, minCost *int, costs [][]int) {
	if row >= len(costs) {
		if curCost < *minCost {
			*minCost = curCost
		}
		return
	}

	for i := 0; i < 3; i++ {
		if i == preSelect {
			continue
		}
		helper(row+1, i, curCost+costs[row][i], minCost, costs)
	}
}

// 动态规划
//维护一个二维数组dp，dp[i][j]表示表示刷到第i+1个房子用颜色的最小花费，状态转移方程为dp[i][j] += min(dp[i-1][(j+1)%3],dp[i-1][(i+2)%3])。
//整个状态转移方程很好理解：如果当前房子要刷红色，则上一个房子只能用蓝色或者绿色来刷，那么当前房子用红色刷的最小花费就等于当前房子用红色刷的钱加上
//刷到上一个房子用绿色和蓝色中的较小者，这样当结算到最后一个房子时，只要取出三个累计花费的最小值即可。
func minCost2(costs [][]int) int {
	if len(costs) == 0 || len(costs[0]) == 0 {
		return 0
	}

	n := len(costs)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, 3)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < 3; j++ {
			dp[i][j] = costs[i][j]
		}
	}

	for i := 1; i < n; i++ {
		for j := 0; j < 3; j++ {
			dp[i][j] += min(dp[i-1][(j+1)%3], dp[i-1][(j+2)%3])
		}
	}

	return min(dp[n-1][0], min(dp[n-1][1], dp[n-1][2]))
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
