package leet_121

func getMax(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// 二维动态规划
// 时间复杂度O(N),空间复杂度O(N)
func maxProfit(prices []int) int {
	l := len(prices)
	if l < 1 {
		return 0
	}

	dp := make([][]int, l)
	for i := 0; i < l; i++ {
		dp[i] = make([]int, 2)
	}


	// base condition
	// 第一天只有不买和买入两种情况
	dp[0][0] = 0          // 不买，当然是0
	dp[0][1] = -prices[0] // 买入，当前收入为负

	for i := 1; i < l; i++ {
		// 手头不持有股票
		dp[i][0] =  getMax(dp[i-1][0], dp[i-1][1] + prices[i])

		// 手头持有股票
		// dp[i][1] = getMax(dp[i-1][0] - prices[i],dp[i-1][1])
		dp[i][1] =  getMax(-prices[i],dp[i-1][1]) // 今天刚买的，或者昨天就买了
	}

	// 最后一天不持有股票肯定比最后一天持有股票所获得的收益要高
	//return dp[l-1][1]
	return dp[l-1][0]
}

// 二维DP优化。时间复杂度O(N),空间复杂度O(1)
// 注意到，新状态只和前一个状态相关，而前一个状态只会出现两种结果，就是dp[i][0]和dp[i][1]。
// 也就是说，每一次求值，都只会用到两个状态，所以只需要在迭代的时候使用两个变量来保存值就好了。

func maxProfitOptimized(prices []int) int {
	length :=  len(prices)
	if length < 1 {
		return 0
	}

	p0,p1 := 0,-prices[0]
	for i := 1 ; i< length; i++ {
		p0 = getMax(p0,p1+prices[i])
		p1 = getMax(-prices[i],p1)
	}

	return p0
}

// 模拟DP
// 时间复杂度O(N)，空间复杂度O(1)
func getMaxProfit(prices []int) int {
	if len(prices) < 1 {
		return 0
	}

	min,maxProfit,length := prices[0],0,len(prices)
	for i := 1; i< length; i++ {
		if prices[i] - min > maxProfit {
			maxProfit = prices[i] - min
		}
		if prices[i] < min {
			min = prices[i]
		}
	}

	return maxProfit
}
// For test
func MaxProfit(prices []int) int {
	return maxProfit(prices)
}