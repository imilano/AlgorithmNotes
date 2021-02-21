package leet_277

/*
Suppose you are at a party with n people (labeled from 0 to n - 1) and among them, there may exist one celebrity. The definition of a celebrity is that all the other n - 1 people know him/her but he/she does not know any of them.

Now you want to find out who the celebrity is or verify that there is not one. The only thing you are allowed to do is to ask questions like: "Hi, A. Do you know B?" to get information of whether A knows B. You need to find out the celebrity (or verify there is not one) by asking as few questions as possible (in the asymptotic sense).

You are given a helper function bool knows(a, b)which tells you whether A knows B. Implement a function int findCelebrity(n). There will be exactly one celebrity if he/she is in the party. Return the celebrity's label if there is a celebrity in the party. If there is no celebrity, return -1.
*/

// 暴力解法。由题意可知，celebrity就是有向图中那个入度为n-1，出度为0的点
func findCelebrity(n int) int {
	if n == 0 {
		return -1
	}

	in := make([]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if knows(i, j) {
				in[j]++
			}
		}
	}

	m, res := in[0], 0
	for k, v := range in {
		if v > m {
			m = v
			res = k
		}
	}

	return res
}

// 初始一个一维数组，假定每个人都是名人，然后开始遍历
func findCelebrity2(n int) int {
	if n == 0 {
		return -1
	}

	for i := 0; i < n; i++ {
		j := 0
		for ; j < n; j++ {
			if i != j && (knows(i, j) || !knows(j, i)) { // 对于i，如果i认识j或者j不认识i,那么i都不可能是名人，直接跳过i，从i+1开始继续遍历
				break
			}

			if j == n { // 对于i，如果对于每个j，如果i都不认识j，并且j都认识i，那么i就是名人，直接返回
				return i
			}
		}
	}

	return -1 // 所有人都检查完毕，没有名人
}

func max(a, b int) int {
	if a < b {
		return b
	}

	return a
}

func knows(a, b int) bool {
	return true
}
