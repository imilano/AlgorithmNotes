package leet_84

/*
	Given n non-negative integers representing the histogram's bar height where the width of each bar is 1,
	find the area of largest rectangle in the histogram.
 */

func max(a,b int) int {
	if a > b {
		return a
	}

	return b
}

// space complexity: O(n)
// time complexity: O(n)
// For array of length n, at every position i,there are two special positions, let's say its index is left and right,
// where  array[left] <= array[i] and array[right] <= array[i]. So, for position i, we can get the area of max rectangle,
// which is  array[i] * (array[right] - array[left] - ).
// So our goal is to find the left and right position for every index i.
// we create two array to store it. A straight-forward solution is to scan the array, and in that scan, use a loop to find the position, which leads to O(n^2) time complexity.
// here is a way to shift O(n^2) to O(n).
//
// for example, when search left of index i, if array[i] <= array[i-1], we can reuse the left of i-1 to avoid another loop , which means we could directly jump to the left of i-1.
// same for right.
func largestRectangleArea(heights []int) int {
	var maxArea int
	if len(heights) == 0 {
		return maxArea
	}

	length := len(heights)
	lessFromLeft := make([]int,length)
	lessFromRight := make([]int,length)
	lessFromLeft[0],lessFromRight[length-1]=-1,length // used in loop for terminal condition

	for i := 1; i < length;i++ {
		p := i-1
		for p >= 0 && heights[p] >= heights[i] {
			p = lessFromLeft[p]
		}

		lessFromLeft[i] = p
	}

	for i := length-2; i>=0;i-- {
		p := i + 1
		for p < length && heights[p] >= heights[i] {
			p = lessFromRight[p]
		}

		lessFromRight[i] = p
	}

	for i := 0; i< length;i++ {
		maxArea = max(maxArea, heights[i] * (lessFromRight[i]-lessFromLeft[i]-1))
	}

	return maxArea
}
