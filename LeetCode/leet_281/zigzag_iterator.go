package leet_281

import "math"

/*
Given two 1d vectors, implement an iterator to return their elements alternately.

Example:

Input:
v1 = [1,2]
v2 = [3,4,5,6]

Output: [1,3,2,4,5,6]

Explanation: By calling next repeatedly until hasNext returns false,
             the order of elements returned by next should be: [1,3,2,4,5,6].
Follow up: What if you are given k 1d vectors? How well can your code be extended to such cases?

Clarification for the follow up question:
The "Zigzag" order is not clearly defined and is ambiguous for k > 2 cases. If "Zigzag" does not look right to you, replace "Zigzag" with "Cyclic". For example:

Input:
[1,2,3]
[4,5,6,7]
[8,9]

Output: [1,4,8,2,5,9,3,6,7].
*/
type zigzag struct {
	nums [][]int
	idx  []int
	cur  int
}

func (z *zigzag) NewIterator(v1 []int, v2 []int) {
	if v1 != nil {
		z.nums = append(z.nums, v1)
		z.idx = append(z.idx, 0)
	}
	if v2 != nil {
		z.nums = append(z.nums, v2)
		z.idx = append(z.idx, 0)
	}

	z.cur = 0
}

func (z *zigzag) next() int {
	if z.hasNext() {
		if z.cur == len(z.nums) {  // 掉头
			z.cur = 0
		}
		for z.idx[z.cur] >= len(z.nums[z.cur]) {  // 跳过那些idx超过数组长度的数组
			z.cur++
		}
		t := z.nums[z.cur][z.idx[z.cur]]
		z.idx[z.cur]++
		z.cur++
		return t
	} else {
		return math.MinInt32  // how to return none
	}
}

func (z *zigzag) hasNext() bool {
	var res bool
	for i := range z.nums {
		t := z.idx[i] < len(z.nums[i])
		res = res || t
	}

	return res
}
