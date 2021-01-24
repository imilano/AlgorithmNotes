package leet_213

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// 只计算[start,end]之间的值
func robInRange(nums []int, start, end int) int {
	n, n1, n2 := 0, 0, 0

	for i := end; i >= start; i-- {
		n = max(n1, n2+nums[i])
		n2, n1 = n1, n
	}

	return n
}

func rob(nums []int) int {
	length := len(nums)

	if length < 1 {
		return 0
	}

	if length == 1 {
		return nums[0]
	}

	return max(
		robInRange(nums, 0, length-2), // 抢第一个不抢最后一个
		robInRange(nums, 1, length-1), // 不抢第一个抢最后一个
	)
}

//-------------------------------------------------------------------
func rob2(nums []int) int {
	n := len(nums)
	if n < 0 {
		return 0
	}

	if n == 1 {
		return nums[0]
	}

	return max(robInRange2(nums, 0, n-1), robInRange2(nums, 1, n))
}

// 使用两个变量表示抢不抢当前的房子，日中rob表示抢当前的房子可以得到的价值，而notRob表示不抢当前的房子可以得到的价值。
func robInRange2(nums []int, left, right int) int {
	var rob, notRob int

	for i := left; i < right; i++ {
		preRob, preNotRob := rob, notRob //  保存之前的状态
		rob = preNotRob + nums[i]        // 如果抢当前的房子，那么不应该抢前一个房子，故而rob的值为前一个不抢的价值加上当前房子的价值
		notRob = max(preRob, preNotRob)  // 如果不抢当前的房子，那么前一个房子可以👍也可以不抢。不抢当前房子的所能所得的价值应该是二者中的较大者
	}

	return max(rob, notRob)
}
