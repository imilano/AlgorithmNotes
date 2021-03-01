package ch15_dynamic_programming

import (
	"math"
)

// 对自身的递归调用导致子问题的重复求解
func cutRod(profit []int, n int) int {
	if n == 0 {
		return 0
	}

	p := math.MinInt32
	for i := 0; i < n; i++ {
		p = max(p, profit[i]+cutRod(profit, n-i))
	}

	return p
}

// 带备忘录的自顶向下方法
func curRodWithMemo(profit []int, n int) int {
	memo := make([]int, len(profit)+1)
	for i := range memo {
		memo[i] = math.MinInt32
	}

	return topDownMemo(profit, n, memo)
}

func topDownMemo(profit []int, n int, memo []int) int {
	p := math.MinInt32
	if n == 0 {
		p = 0
	}
	if memo[n] != math.MinInt32 {
		return memo[n]
	}

	for i := 0; i < n; i++ {
		p = max(p, profit[i]+topDownMemo(profit, n-i, memo))
	}

	memo[n] = p
	return p
}

// bottom Up
func cutRodBottomUp(profit []int, n int) int {
	memo := make([]int, len(profit)+1)

	memo[0] = 0
	for i := 1; i <= n; i++ {
		q := math.MinInt32
		for j := 1; j <= i; j++ {
			q = max(q, profit[j]+memo[i-j])
		}

		memo[i] = q
	}

	return memo[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
