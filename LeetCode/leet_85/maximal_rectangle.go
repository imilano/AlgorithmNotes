package leet_85

/*
	Given a rows x cols binary matrix filled with 0's and 1's, find the largest rectangle containing only 1's and return its area.
*/
// 解法1，将每一层都看做84题中的直方图，求解最大矩形面积时，只需要对每一层依次计算高度数组，然后在每一层都计算一次最大面积即可
func maximalRectangleHistogram(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}

	row, col := len(matrix), len(matrix[0])
	height := make([]int, col)
	var maxArea int

	for i := 0; i < row; i++ {
		// compute height array of every histogram of current row
		for j := 0; j < col; j++ {
			if matrix[i][j] == '1' {
				height[j] = 1 + height[j]
			} else {
				// Must add this line, cause once we meet 0, there will be no histogram
				height[j] = 0
			}
		}

		// compute maximum area of every histogram
		maxArea = max(maxArea, largestRectangleArea(height))
	}

	return maxArea
}

//
func largestRectangleArea(heights []int) int {
	var maxArea int
	if len(heights) == 0 {
		return maxArea
	}

	length := len(heights)
	lessFromLeft := make([]int, length)
	lessFromRight := make([]int, length)
	lessFromLeft[0], lessFromRight[length-1] = -1, length // used in loop for terminal condition

	for i := 1; i < length; i++ {
		p := i - 1
		for p >= 0 && heights[p] >= heights[i] {
			p = lessFromLeft[p]
		}

		lessFromLeft[i] = p
	}

	for i := length - 2; i >= 0; i-- {
		p := i + 1
		for p < length && heights[p] >= heights[i] {
			p = lessFromRight[p]
		}

		lessFromRight[i] = p
	}

	for i := 0; i < length; i++ {
		maxArea = max(maxArea, heights[i]*(lessFromRight[i]-lessFromLeft[i]-1))
	}

	return maxArea
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
