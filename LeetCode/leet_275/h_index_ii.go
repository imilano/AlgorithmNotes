package leet_275

/*
Given an array of citations sorted in ascending order (each citation is a non-negative integer) of a researcher, write a function to compute the researcher's h-index.

According to the definition of h-index on Wikipedia: "A scientist has index h if h of his/her N papers have at least h citations each, and the other N − h papers have no more than h citations each."

Example:

Input: citations = [0,1,3,5,6]
Output: 3
Explanation: [0,1,3,5,6] means the researcher has 5 papers in total and each of them had
             received 0, 1, 3, 5, 6 citations respectively.
             Since the researcher has 3 papers with at least 3 citations each and the remaining
             two with no more than 3 citations each, her h-index is 3.
Note:

If there are several possible values for h, the maximum one is taken as the h-index.

Follow up:

This is a follow up problem to H-Index, where citations is now guaranteed to be sorted in ascending order.
Could you solve it in logarithmic time complexity?
*/

// 此题与上一题的不同之处在于进行了排序，升序排序，为我们省去了一步。与上一题相似，从后向前进行遍历，找到第一个下标i，使得citation[i] <= n-i，
//返回这个h即可。题目要求我们优化到对数复杂度，那么我们可以使用二分查找
func hIndex(citations []int) int {
	n := len(citations)
	left, right := 0, n-1
	for left <= right {
		mid := left + (right-left)/2
		if citations[mid] == n-mid {
			return n - mid
		}
		if citations[mid] < n-mid {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return n - left
}
