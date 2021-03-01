package leet_215

import (
	"sort"
)

/*
Find the kth largest element in an unsorted array. Note that it is the kth largest element in the sorted order, not the kth distinct element.

Example 1:

Input: [3,2,1,5,6,4] and k = 2
Output: 5
Example 2:

Input: [3,2,3,1,2,4,5,5,6] and k = 4
Output: 4
Note:
You may assume k is always valid, 1 ≤ k ≤ array's length.
*/

//--------------------------------------------------------------------------------------
// 先排序，再取值
func findKthLargest(nums []int, k int) int {
	n := len(nums)
	sort.Ints(nums)

	return nums[n-k]
}

//-------------------------------------------------------------------------------------------
//使用堆这个数据结构，然后pop n-k次，接下来pop的那个元素就是我们需要的元素
func findKthLargest2(nums []int, k int) int {
	// 使用container/heap
	return 0
}

//-----------------------------------------------------------------
// 使用计数排序
func findKthLargest3(nums []int, k int) int {
	min, max := getMin(nums), getMax(nums)

	count := make([]int, max-min+1)
	for _, v := range nums {
		count[v-min]++
	}

	for i := max - min; i >= 0; i-- {
		k -= count[i]
		if k <= 0 {
			return i + min
		}
	}

	return 0
}

//--------------------------------------------------------------------------------------------------------------------
// 每次快排都可以确定一个元素的最终位置，准确说，是确定我选定的pivot的最终位置（pivot的选取是任意的），每次partition结束，都会返回这个pivot的最终位置。
// 只需要将partition返回的index与k进行比价并设定边界即可.
func findKthLargest4(nums []int, k int) int {
	n := len(nums)
	left, right := 0, n-1
	for true {
		pos := partition(nums, left, right)
		if pos == n-k {
			return nums[n-k]
		}
		if pos > n-k {
			right = pos - 1
		}
		if pos < n-k {
			left = pos + 1
		}
	}

	return 0
}

func quickSort(nums []int, left, right int) {
	if left < right {
		mid := partition(nums, left, right) // 本次拍好了mid
		quickSort(nums, left, mid-1)        // 排好mid之前的元素
		quickSort(nums, mid+1, right)       // 拍好mid之后的元素
	}
}
func partition(arr []int, left, right int) int {
	pivot := left
	index := pivot + 1

	for i := index; i <= right; i++ {
		if arr[i] < arr[pivot] {
			arr[i], arr[index] = arr[index], arr[i]
			index += 1
		}
	}

	arr[pivot], arr[index-1] = arr[index-1], arr[pivot]
	return index - 1
}

//------------------------------------
// util func
func getMin(arr []int) int {
	res := arr[0]
	for _, v := range arr {
		if v < res {
			res = v
		}
	}

	return res
}

func getMax(arr []int) int {
	res := arr[0]
	for _, v := range arr {
		if v > res {
			res = v
		}
	}

	return res
}
