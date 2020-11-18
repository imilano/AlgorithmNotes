package leet_213

func max(a,b int) int {
	if a > b {
		return a
	}

	return  b
}


// 只计算[start,end]之间的值
func robInRange(nums []int, start,end int)int{
	n,n1,n2 := 0,0,0

	for i := end; i >= start; i-- {
		n = max(n1,n2+nums[i])
		n2 = n1
		n1 = n
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
		robInRange(nums,0,length-2), // 抢第一个不抢最后一个
		robInRange(nums,1,length-1), // 不抢第一个抢最后一个
		)
}