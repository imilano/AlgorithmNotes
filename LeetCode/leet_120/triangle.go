package leet_120

import "math"


/*
Given a triangle, find the minimum path sum from top to bottom. Each step you may move to adjacent numbers on the row below.

For example, given the following triangle

[
     [2],
    [3,4],
   [6,5,7],
  [4,1,8,3]
]
The minimum path sum from top to bottom is 11 (i.e., 2 + 3 + 5 + 1 = 11).

Note:

Bonus point if you are able to do this using only O(n) extra space, where n is the total number of rows in the triangle.
 */

// Dynamic Programming
// 由于没有说不能破坏原数组，所以可以直接复用原数组，直接一层一层的累加下去
// 转态转移方程如下： dp[i][j] = dp[i-1][j] + dp[i-1][j-1]
// 需要注意的是，对于两侧边缘的情况，直接将上一层的边缘值赋值给当前层的边缘值即可
func minimumTotal(triangle [][]int) int {
	row := len(triangle)
	for i :=1  ; i< row;i++ {  // 直接从第一层开始遍历
		for j := 0; j< len(triangle[i]); j++ {
			if j == 0 {
				triangle[i][j] += triangle[i-1][j]
			} else if j == len(triangle[i])-1 {
				triangle[i][j] += triangle[i-1][j-1]
			} else {
				triangle[i][j] += min(triangle[i-1][j], triangle[i-1][j-1])
			}
		}
	}

	minVal := math.MaxInt32
	for i := 0; i< len(triangle[row-1]) ; i++ {
		minVal = min(minVal, triangle[row-1][i])
	}
	return minVal
}

func min(a,b int) int {
	if a > b {
		return b
	}

	return a
}



// DP optomized
func minimumTotalConcise(triangle [][]int) int {
	length := len(triangle)
	dp := make([]int, len(triangle[length-1]))
	copy(dp,triangle[length-1])


	for  i := length-2; i>=0; i-- {
		for  j := 0; j<= i;j++ {
			dp[j] =  min(dp[j], dp[j+1]) + triangle[i][j]
		}
	}

	return dp[0]
}


// greedy algorithm
// WRONG， 局部最优解无法保证全局最优解，过早剪纸
//func minimumTotalGreedy(triangle [][]int) int {
//	length := len(triangle)
//	if length == 0 {
//		return 0
//	}
//	minVal := triangle[0][0]
//	curStep := 0
//
//	    for i := 1; i< len(triangle);i++ {
//	        curRow := triangle[i]
//
//	        if len(curRow) > 1 && curStep+1 < len(curRow) {
//	            curStep = min(curRow, curStep,curStep+1)
//	        }
//
//	        minVal += curRow[curStep]
//	    }
//
//
//	return minVal
//}
//
//func helper(arr [][]int, row,col int) int {
//	if row == len(arr)
//}
//
//
//func min(arr []int, i, j int) int {
//	if  arr[i] > arr[j] {
//		return j
//	}
//
//	return i
//}
