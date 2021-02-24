package leet_287

/*
Given an array nums containing n + 1 integers where each integer is between 1 and n (inclusive), prove that at least one duplicate number must exist. Assume that there is only one duplicate number, find the duplicate one.

Example 1:

Input: [1,3,4,2,2]
Output: 2
Example 2:

Input: [3,1,3,4,2]
Output: 3
Note:

You must not modify the array (assume the array is read only).
You must use only constant, O(1) extra space.
Your runtime complexity should be less than O(n2).
There is only one duplicate number in the array, but it could be repeated more than once.
*/

// 注意要求：不能修改数组（没法用排序），常数空间复杂度（类计数排序行不通），时间复杂度不超过n^2，数字可能重复出现（不能用求和方法）
// 二分搜索法。先求出1到n之间的中间数mid，然后统计数组中小于等于mid的数的个数，如果小于mid，说明重复值在mid+1到high之间，否则在low到mid之间
func findDuplicate(nums []int) int {
	//low,high := 0, len(nums)
	low, high := 1, len(nums)
	for low < high {
		mid := low + (high-low)/2
		var c int
		for i := range nums {
			if nums[i] <= mid {
				c++
			}
		}

		//if c < mid {
		if c <= mid {
			low = mid + 1
		} else {
			high = mid
		}
	}

	return low
}

// 快慢指针.由于题目限定区间[1,n]，并且有重复数字存在，那么就可以模拟出一个环。
//再次复习一遍，将图依次划分为a、b、c三部分，快指针每次走两步，慢指针走一步，二者相遇时，慢指针走了a+b，快指针走了a+b+c+b，快指针走的路程是
//慢指针的两倍，故而2(a+b) = a+b+c+b,得出a=c
func findDuplicate2(nums []int) int {
	slow, fast := 0, 0
	t := 0
	for true {
		slow = nums[slow]       // 慢指针走一步
		fast = nums[nums[fast]] // 快指针走两步
		if slow == fast {       // 快慢指针相遇。
			break
		}
	}

	for true {
		slow = nums[slow]
		t = nums[t]
		if slow == t { // 二者相遇
			break
		}
	}

	return slow
}

// 位操作，直接看代码，很易懂
func findDuplicate3(nums []int) int {
	var res int
	n := len(nums)
	for i := 0; i < 23; i++ {
		t := 1 << i
		var cnt1, cnt2 int
		for i := 0; i < n; i++ {
			if t&i > 0 {
				cnt1++
			}

			if t&nums[i] > 0 {
				cnt2++
			}

			if cnt2 > cnt1 {
				res += t
			}
		}
	}

	return res
}
