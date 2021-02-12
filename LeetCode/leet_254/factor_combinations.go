package leet_254

/*
Numbers can be regarded as product of its factors. For example,

8 = 2 x 2 x 2;
  = 2 x 4.
Write a function that takes an integer n and return all possible combinations of its factors.

Note:

You may assume that n is always positive.
Factors should be greater than 1 and less than n.
Example 1:

Input: 1
Output: []
Example 2:

Input: 37
Output:[]
Example 3:

Input: 12
Output:
[
  [2, 6],
  [2, 2, 3],
  [3, 4]
]
Example 4:

Input: 32
Output:
[
  [2, 16],
  [2, 2, 8],
  [2, 2, 2, 4],
  [2, 2, 2, 2, 2],
  [2, 4, 4],
  [4, 8]
]
*/

// 回溯法。注意题目中要求的是，按照因子的大小进行排列
func getFactors(n int) [][]int {
	var res [][]int
	helper(2, n, []int{}, &res)
	return res
}

func helper(start, n int, out []int, res *[][]int) { // TODO why transfer parameter out using value rather than pointer
	if n == 1 {
		if len(out) > 1 {
			*res = append(*res, out)
		}
		return
	}

	for i := start; i <= n; i++ {
		if n%i != 0 {
			continue
		}

		out = append(out, i)
		helper(i, n/i, out, res) // 注意这里为什么从i开始进行计算：因为题目要求对因子进行升序排列，故而后面的因子只会大于等于当前已经计算出的因子
		out = (out)[:len(out)-1]
	}
}

//------------------------------------------------
func getFactors2(n int) [][]int {
	var res [][]int
	helper2(2, n, []int{}, &res)
	return res
}

func helper2(start, end int, out []int, res *[][]int) {
	for i := start; i*i <= end; i++ {
		if end%i != 0 {
			continue
		}

		t := make([]int, len(out)) // TODO why make a copy here
		copy(t, out)
		t = append(t, i)
		helper2(i, end/i, t, res) // 这一步相当于一个更精细化的划分，进一步划分大因子。
		t = append(t, end/i)      // 直接添加大因子，不对大因子继续进行划分
		*res = append(*res, t)
	}
}
