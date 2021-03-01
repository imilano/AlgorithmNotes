package leet_164

import "sort"

/*
Given an unsorted array, find the maximum difference between the successive elements in its sorted form.

Return 0 if the array contains less than 2 elements.

Example 1:

Input: [3,6,9,1]
Output: 3
Explanation: The sorted form of the array is [1,3,6,9], either
             (3,6) or (6,9) has the maximum difference 3.
*/

//-------------------------------------------
// brute force
func maximumGap(nums []int) int {
	length := len(nums)
	if length < 2 {
		return 0
	}

	sort.Ints(nums)
	var res int
	for i := 1; i < length; i++ {
		res = max(res, nums[i]-nums[i-1])
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

//-------------------------------------------------------------
// 桶排序

//对于有序数组，我们只需要计算两两数字之差即可解决问题。此处我们可以使用宏观上有序的桶排序，只需要O(n)时间的复杂度。在每个桶中，我们维持两个值，
//一个表示当前桶的最大值，一个表示当前桶的最小值。因为各个桶都是宏观有序的，故而我们只需要计算后一个桶的最小值减去前一个桶的最大值，就可以得到连续
//数字的最大差值。那么，问题来了，我们如何确保这个差值不会比桶内的最大最小元素之间的差值小呢？
//
//我们把每个箱子能放的数字个数记为interval，每个箱子中的最大max，最小min。那么箱子的划分范围就是 [min,min + 1 * interval -1]、[min + 1 * interval, min + 2*interval-1]、、、
//划分箱子后，我们很容易把数字映射到(nums[i]-min)/interval的编号的箱子中去。那么，如何确定interval呢？
//
//interval过大，箱子内部的数字可能会产生最大gap；箱子过小的话，需要更多的箱子去存储，比较的次数也增多了。
//所以我们需要找到那个保证箱子内部不会产生最大gap，并且尽量大的interval。
//
//解决方案是，我们的划分需要保证至少有一个空箱子，那么得到的gap一定会大于interval（空箱子后面那个箱子的min减去空箱子前面那个箱子的max），而箱子中的数字最大的gap
//是interval-1。
//
//接下来，如何保证至少有一个空箱子呢？
//鸽巢原理的变形，如果有n个数字，但是箱子数大于n，那么一定会出现空箱子。总范围是max-min，那么interval=（max-min）/箱子数，为使interval最大，箱子数取最小即可，也就是n+1。
//所以interval = （max-min）/(n+1)，除不尽的话，对interval向上取整即可，取整对gap无影响。
