package leet_70

/*
You are climbing a stair case. It takes n steps to reach to the top.
Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?
*/

// we can see from climbStairs, that F(n) = F(n-1) + F(n-2), which actually is a Fibonacci number.
// So when we need to count ith(which is n), we just need two element.
func climbStairsFib(n int) int {
	if n == 1 {
		return 1
	}
	// 为什么这么赋值，按照双指针滑动的方向就可以看出来。
	MinusOne,MinusTwo := 1,2
	for i := 3; i<= n;i++ {
		N := MinusOne + MinusTwo
		MinusOne = MinusTwo
		MinusTwo = N
	}

	return MinusTwo
}
//func climbStairsFib(n int) int {
//	if n == 1 {
//		return 1
//	}
//
//	MinusOne,MinusTwo := 1,2
//	for i := 3; i<= n;i++ {
//		N := MinusOne + MinusTwo
//		MinusTwo = MinusOne
//		MinusOne = N
//	}
//
//	return MinusOne
//}

func climbStairs(n int) int {
	dp := make([]int,n+1)
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	dp[1],dp[2] = 1,2
	for i := 3; i<= n; i++ {
		dp[i] = dp[i-1]  + dp[i-2]
	}

	return dp[n]
}
