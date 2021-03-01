package leet_75

/*
	Given an array nums with n objects colored red, white, or blue, sort them in-place so that objects of the same color are adjacent,
	with the colors in the order red, white, and blue.
	Here, we will use the integers 0, 1, and 2 to represent the color red, white, and blue respectively.

	Follow up:
		Could you solve this problem without using the library's sort function?
		Could you come up with a one-pass algorithm using only O(1) constant space?
*/

func sortColors(nums []int) {
	if nums == nil || len(nums) == 1 {
		return
	}

	length := len(nums)
	index := 0
	for i := 1; i < length; i++ {
		index = i - 1
		tmp := nums[i]
		for index >= 0 && tmp < nums[index] {
			nums[index+1] = nums[index]
			index--
		}
		nums[index+1] = tmp
	}
}

// 这是一个荷兰国旗问题，我们初始将数组划分为四个组：红、白、蓝和未划分。我们维持三个指针，指针red指向红组的结尾，指针white指向白组的结尾，指针blue指向蓝组的开头。
// 接下来从头开始迭代数组，首先red和white都初始化为0，然后将blue指向数组最后的位置。接下来逐步开始迭代，循环结束条件是white尾指针小于等于blue头指针
// 		- 如果白指针指向红色，那么将白指针的值和红指针的值交换，然后二者同时加加
//		- 如果白指针指向蓝色，那么将白指针和蓝指针交换（也就是选择一个位分类的值和白指针交换），然后减少蓝指针
//		- 如果白指针指向白色，白指针自增即可
// time complexity: O(n)
// space complexity: O(1)
func sortColorsDutch(nums []int) {
	if nums == nil || len(nums) == 1 {
		return
	}

	red, white, blue := 0, 0, len(nums)-1
	for white <= blue {
		if nums[white] == 0 {
			nums[white], nums[red] = nums[red], nums[white]
			white++
			red++
		} else if nums[white] == 2 {
			nums[white], nums[blue] = nums[blue], nums[white]
			blue--
		} else if nums[white] == 1 {
			white++
		}
	}
}
