package leet_274

import "sort"

/*
Given an array of citations (each citation is a non-negative integer) of a researcher, write a function to compute the researcher's h-index.

According to the definition of h-index on Wikipedia: "A scientist has index h if h of his/her N papers have at least h citations each, and the other N − h papers have no more than h citations each."

Example:

Input: citations = [3,0,6,1,5]
Output: 3
Explanation: [3,0,6,1,5] means the researcher has 5 papers in total and each of them had
             received 3, 0, 6, 1, 5 citations respectively.
             Since the researcher has 3 papers with at least 3 citations each and the remaining
             two with no more than 3 citations each, her h-index is 3.
Note: If there are several possible values for h, the maximum one is taken as the h-index.
*/

// 由题意可以得出，要找出h，使得h篇文章的引用次数都大于等于h，那么对于降序排序后的数组，只需要找到第一个下标i，使得引用次数citation[i]小于i即可
func hIndex(citations []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(citations)))
	n := len(citations)
	for i := 0; i < n; i++ {
		if i >= citations[i] {
			return i
		}
	}

	return n
}
