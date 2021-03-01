package leet_216

/*
Find all valid combinations of k numbers that sum up to n such that the following conditions are true:

Only numbers 1 through 9 are used.
Each number is used at most once.
Return a list of all possible valid combinations. The list must not contain the same combination twice, and the combinations may be returned in any order.



Example 1:

Input: k = 3, n = 7
Output: [[1,2,4]]
Explanation:
1 + 2 + 4 = 7
There are no other valid combinations.
*/

var res [][]int

func combinationSum3(k int, n int) [][]int {
	var cur []int
	backtrace(1, 0, 0, n, k, cur)
	return res
}

// 对于每个数字，有两种可能，加入或者不加入。加入的话curK+1，不加入则仍为curK
// 根据一个贪心的思路，从大开始到小开始进行遍历，能够更快减枝，会更有效率一些。
func backtrace(start, curSum, curK, sum, k int, cur []int) {
	if curK == k && curSum == sum {
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		res = append(res, tmp)
	}

	if curK > k || start > 9 {
		return
	}

	for i := start; i <= 9; i++ {
		if curSum+i <= sum {
			cur = append(cur, i)
			backtrace(i+1, curSum+i, curK+1, sum, k, cur)
			cur = cur[:len(cur)-1]
		}
	}
}
