package leet_53

import "math"

// 将数组递归的划分为左和右两部分，并且在两个部分中递归的查找最大和。然而可能出现的问题是，最大连续数组可能会跨越两个子数组。
// 对于跨越子数组的情况，我们采用线性扫描的方法，从中点开始向左右进行扩展，分别求出中点左右的最大子数组和，故而跨越中点的最大子数组和就是左子数组的最大和
// 加上右子数组的最大和加上中间元素。
// 故而最大和就在左边最大和、右边最大和以及跨越数组最大和中产生

// time complexity: O(nlogn)
// space complexity: O(logn)
func maxSubArrayWithDC(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	return divide(nums,0,len(nums)-1)
}

func divide(nums []int, left int, right int)int {
	if left > right {
		return math.MinInt64
	}

	mid := left + (right-left)/2
	leftMax := divide(nums,left,mid-1)
	rightMax := divide(nums,mid+1,right)

	var lmax,rmax int
	sum := 0
	for i := mid-1; i>= left;i-- {
		sum += nums[i]
		if sum > lmax {
			lmax = sum
		}
	}

	sum = 0
	for i := mid +1; i <= right;i++ {
		sum +=  nums[i]
		if sum > rmax {
			rmax = sum
		}
	}

	return max(max(leftMax,rightMax),lmax+rmax+nums[mid])
}

func max(a,b int) int {
	if a > b {
		return a
	}

	return b
}