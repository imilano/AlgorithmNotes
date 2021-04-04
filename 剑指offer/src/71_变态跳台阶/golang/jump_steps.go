package golang


// 一只青蛙一次可以跳上1级台阶，也可以跳上2级……它也可以跳上n级。求该青蛙跳上一个n级的台阶总共有多少种跳法。

//跳上 n-1 级台阶，可以从 n-2 级跳 1 级上去，也可以从 n-3 级跳 2 级上去...，那么: f(n-1) = f(n-2) + f(n-3) + ... + f(0)
//同样，跳上 n 级台阶，可以从 n-1 级跳 1 级上去，也可以从 n-2 级跳 2 级上去... ，那么: f(n) = f(n-1) + f(n-2) + ... + f(0)
//综上可得 f(n) - f(n-1) = f(n-1), 即 f(n) = 2*f(n-1). 所以 f(n) 是一个等比数列.



func jumpFloorII( number int ) int {
	// write code here
	if number <= 2 {
		return number
	}
	dp := make([]int,number+1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= number;i++ {
		res := 1
		j := i - 1
		for j > 0 {
			res += dp[j]
			j--
		}
		dp[i] = res
	}

	return dp[number]
}
