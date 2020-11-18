package leet_122

import (
	"lightsinger.life/algorithm/leetcode/utils"
)

// 二维DP
// 时间复杂度O(N),空间复杂度O(N)
func maxProfit(prices []int) int {
	length := len(prices)
	if length < 1 {
		return 0
	}

	dp := make([][]int, length)
	for i := 0; i < length; i++ {
		dp[i] = make([]int, 2)
	}

	// base condition
	dp[0][0] = 0
	dp[0][1] = -prices[0]

	for i := 1; i < length; i++ {
		dp[i][0] = utils.Max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = utils.Max(dp[i-1][1], dp[i-1][0]-prices[i])
	}

	return dp[length-1][0]
}

// 二维DP优化，时间复杂度O(N),空间复杂度O(1)
// 当前状态只和前一个状态有关，因而无需使用数组，只需使用简单使用两个变量即可
func maxProfitOptimized(prices []int) int {
	length := len(prices)
	if length < 1 {
		return 0
	}

	p0, p1 := 0, -prices[0]
	for i := 1; i < length; i++ {
		p0 = utils.Max(p0, p1+prices[i])
		p1 = utils.Max(p1, p0-prices[i])
	}

	return p0
}

// 时间复杂度O(N),空间复杂度O(1)
// 涨幅才会带来收入。收入的来源必然是每次跌落的时候买入，然后涨到顶峰时候就售出。只要有张峰就有收入，就可以计算赚的钱。连续涨可以用两两相减来计算，两两相减再累加，就相当于涨到波峰的最大值减去谷底的值。
//
// 可能的情况下，在每个局部最小值买入股票，然后在之后遇到的第一个局部最大值卖出股票。这个做法等价于找到股票价格数组中的递增子数组，对于每个递增子数组，
// 在开始位置买入并在结束位置卖出。可以看到，这和累计收益是相同的，只要这样的操作的收益为正。
//

func maxProfits(prices []int) int {
	length := len(prices)

	if length < 1 {
		return 1
	}

	var profits int
	for i := 1; i < length; i++ {
		if prices[i] > prices[i-1] {
			profits += prices[i] - prices[i-1]
		}
	}

	return profits
}
