package leet_123


func max(a,b int) int {
	if a > b {
		return a
	}

	return b
}

// 动态规划，时间复杂度O(N)，空间复杂度O(1)
func maxProfit(prices []int) int {
	length := len(prices)

	if length < 1 {
		return 0
	}

	profitOne0, profitOne1 , profitTwo0 , profitTwo1 := 0,-prices[0],0,-prices[0]
	for i := 1 ; i< length; i++ {
		profitTwo0 = max(profitTwo0, profitTwo1 + prices[i])
		profitTwo1 =max(profitTwo1, profitOne0 - prices[i])
		profitOne0 =max(profitOne0, profitOne1 + prices[i])
		profitOne1 =max(profitOne1, -prices[i])

	}

	return profitTwo0

}