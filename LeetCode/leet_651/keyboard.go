package leet_651

/*
	用n表示当前还可以进行的按键次数，用num表示当前屏幕上的A个数，用copy表示当前粘贴板中含有的A个数。
	 - 按A： dp(n-1,num+1,copy)
	 - 按Ctrl + V : dp(n-1,num+copy,copy)
	 - 按 Ctrl + A 和 Ctrl + C： dp(n-2,num,num)  // 全选和复制必然联合使用

	base condition应该是dp(0,num,copy),当n为0时，num即为我们想要的答案。

*/

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxA(N int) int {
	return dp(N, 0, 0) // 按N次按键，此时屏幕中还没有A
}

// 可以把函数更改为dp数组
func dp(n int, num int, copy int) int {
	if n < 0 {
		return num
	}

	// 可以用备忘录消除重叠子问题
	return max(max(dp(n-1, num+1, copy), dp(n-1, num+copy, copy)), dp(num-2, num, num)) // Max in four conditions，
}

/*
这次我们只定义一个「状态」，也就是剩余的敲击次数 n。
这个算法基于这样一个事实，最优按键序列一定只有两种情况：
	- 要么一直按 A：A,A,...A（当 N 比较小时）。
	- 要么是这么一个形式：A,A,...C-A,C-C,C-V,C-V,...C-V（当 N 比较大时）。
*/

func dpNew(N int) int {
	dp := make([]int, N+1)
	dp[0] = 0

	for i := 1; i <= N; i++ {
		dp[i] = dp[i-1] + 1
		for j := 2; j < i; j++ {
			dp[i] = max(dp[i], dp[j-2]*(i-j+1))
		}
	}

	return dp[N]
}
