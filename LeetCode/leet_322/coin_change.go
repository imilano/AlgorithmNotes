package leet_322

/*
	You are given coins of different denominations and a total amount of money amount.
	Write a function to compute the fewest number of coins that you need to make up that amount.
	If that amount of money cannot be made up by any combination of the coins, return -1.
 */

/*
	题解：
	动态规划需要注意的问题：状态转移方程、base condition。
	先说状态转移。对于数组dp，假设n为当前总额，dp[n]为凑出amount为n时所需要的硬币数。由于要求最小的硬币数，则有，dp[n] = min(cur,dp[i-cj]+1).
	其中cur表示不选面值为cj的硬币时，达到amount n所需要的硬币数；而dp[i-cj]+1则表示选择了面值为cj时的硬币时，当前所需的硬币数。
	对于base condition， 可以看到，当n为0时，dp[i]=0；当n<0时，dp[i] =-1。
 */


// Solution 1, brute force
func getMin(a,b int) int {
	if a > b {
		return b
	}

	return a
}

func coinChange(coins []int, amount int) int {
	dp := make([]int,amount+1)
	for i,_ := range dp {
		dp[i] = amount+1
	}
	// base condition
	dp[0] = 0
	for i := 0; i< len(dp); i++ {
		for _,coin := range coins {
			if i - coin < 0 {
				continue
			}

			dp[i] = getMin(dp[i],dp[i-coin]+1)
		}
	}

	if dp[amount] == amount + 1 {
		return -1
	}

	return dp[amount]
}