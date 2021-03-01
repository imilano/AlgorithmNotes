package leet_264

/*
Write a program to find the n-th ugly number.

Ugly numbers are positive numbers whose prime factors only include 2, 3, 5.

Example:

Input: n = 10
Output: 12
Explanation: 1, 2, 3, 4, 5, 6, 8, 9, 10, 12 is the sequence of the first 10 ugly numbers.
Note:

1 is typically treated as an ugly number.
n does not exceed 1690.
*/

// 需要一个最小堆，先将1压入堆中，然后弹出堆顶元素，然后依次压入弹出元素的2倍、3倍和5倍，重复上述过程，直到计数值c得到n即可
// 没有最小堆的话可以使用排序算法来代替，但是时间复杂度过高
func nthUglyNumber(n int) int {
	res := []int{1}

	var i2, i3, i5 int
	for len(res) < n {
		m2, m3, m5 := res[i2]*2, res[i3]*3, res[i5]*5
		mn := min(m2, min(m3, m5))
		if mn == m2 {
			i2++
		}
		if mn == m3 {
			i3++
		}
		if mn == m5 {
			i5++
		}
		res = append(res, mn)
	}

	return res[n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
