package golang


// 把只包含质因子2、3和5的数称作丑数（Ugly Number）。例如6、8都是丑数，但14不是，因为它包含质因子7。 习惯上我们把1当做是第一个丑数。
// 求按从小到大的顺序的第N个丑数。

// 方法1：使用最小堆进行递推
// 方法2： 分别使用三个变量模拟三个队列，每个队列分别按照2、3、5倍的顺序递增
func GetUglyNumber_Solution( index int ) int {
	if index <= 6 {
		return index
	}

	dp := make([]int,index)
	var i2,i3,i5 int
	dp[0] = 1
	for i := 1; i < index;i++ {
		next2,next3,next5 := dp[i2]*2,dp[i3]*3,dp[i5]*5  // 分别模拟三个队列当前的元素
		dp[i] = min(next2,min(next3,next5)) // 选中最小的哪一个
		if dp[i] == next2 {  // 如果选中的是2，那么递增
			i2++
		}
		if dp[i] == next3 { // 如果选中的是3，那么递增
			i3++
		}
		if dp[i] == next5 { // 如果选中的是5，那么递增
			i5++
		}
	}

	return dp[index-1]
}

func min(a,b int) int {
	if a < b {
		return a
	}

	return b
}