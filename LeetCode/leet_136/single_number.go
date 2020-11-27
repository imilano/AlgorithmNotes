package leet_136

/*
Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.

Follow up: Could you implement a solution with a linear runtime complexity and without using extra memory?
*/

// solution 1: set
func singleNumber(nums []int) int {
	set := make(map[int]int)
	for i,v := range nums {
		if _,ok := set[v]; ok {
			delete(set,v)
		} else {
			set[v] = i
		}
	}

	// 最后只剩唯一一个元素
	var res int
	for i := range set {
		res = i
	}

	return res
}

// bitwise, xor
func singleNumber2(num []int)int {
	var res int // 初始值为0，而不能为1,为0能够很好的维持最后一个数的值
	for _,v := range num {
		res ^= v
	}

	return res
}