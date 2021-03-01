package leet_239

import (
	"container/heap"
	"lightsinger.life/algorithm/leetcode/utils"
)

/*
You are given an array of integers nums, there is a sliding window of size k which is moving from the very left of the array to the very right. You can only see the k numbers in the window. Each time the sliding window moves right by one position.

Return the max sliding window.



Example 1:

Input: nums = [1,3,-1,-3,5,3,6,7], k = 3
Output: [3,3,5,5,6,7]
Explanation:
Window position                Max
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7
Example 2:

Input: nums = [1], k = 1
Output: [1]
Example 3:

Input: nums = [1,-1], k = 1
Output: [1,-1]
Example 4:

Input: nums = [9,11], k = 2
Output: [11]
Example 5:

Input: nums = [4,-2], k = 2
Output: [4]
*/

//WRONG method
//func maxSlidingWindow(nums []int, k int) []int {
//	var res []int
//	var curMax int
//
//	n := len(nums)
//	if n < k {
//		res = append(res, findMax(nums))
//	} else { // if n > k
//		curMax = findMax(nums[:k])
//		res = append(res, curMax)
//
//		for i := k; i < n; i++ {
//			if curMax < nums[i] {
//				curMax = nums[i]
//			}
//
//			res = append(res, curMax)
//		}
//	}
//
//	return res
//}
//
//func findMax(nums []int) int {
//	t := nums[0]
//	for _, v := range nums {
//		if v > t {
//			t = v
//		}
//	}
//
//	return t
//}


// 使用一个大小为窗口大小的最大堆。设窗口大小为M，数组长度为N，每次有元素进入窗口，则将该元素压入堆中，每次有元素离开窗口，都将该元素从堆中删除。从堆中插入和删除的时间开销都是
//O(logM)，因此算法的时间复杂度为O(NlogM)
func maxSlidingWindow(nums []int, k int) []int {
	var res []int
	h := &utils.MaxHeap{}
	heap.Init(h)
	n := len(nums)


	//var cnt int
	var idx int
	for true {
		if len(h.MinHeap) == k {
			e := heap.Pop(h).(int)
			res = append(res, e)
			heap.Push(h,e)
			// 不应该是删除堆中最小的元素，因为堆已经排好序了，这个最小的元素并不是我们滑动窗口中将要滑出窗口的那个元素
			//heap.Remove(h,k-1)

			// 正确的做法是找出即将滑出滑动窗口的那个数，然后在堆中找到它的下标，然后将这个下标对应的数从堆中移除
			toRemove := nums[idx-k]
			for i,k := range h.MinHeap {
				if k == toRemove {
					toRemove = i
					break
				}
			}

			heap.Remove(h,toRemove)
		}

		if idx == n {
			break
		}
		if len(h.MinHeap) < k {
			heap.Push(h,nums[idx])
			idx++
		}
	}

	return res
}