package leet_45

/*
Given an array of non-negative integers nums, you are initially positioned at the first index of the array.
Each element in the array represents your maximum jump length at that position.
Your goal is to reach the last index in the minimum number of jumps.
You can assume that you can always reach the last index.
 */

// Utility function
func max(a,b int) int {
	if a > b {
		return a
	}

	return b
}



// traverse from right to left, find the leftmost position that can reach the current
// time complexity: O(n^2)
// space complexity: O(1)
func jump(nums []int) int {
	var jump int
	length := len(nums)
	if length <= 1 {
		return jump
	}

	position := length-1
	for position != 0 {
		for i := 0; i< position;i++ {
			if i + nums[i] >= position {
				position = i
				jump++
				break
			}
		}
	}

	return jump
}


// ---------------------------------------------------------
// Use BFS
// time complexity
// space complexity
func jumpWithBFS(nums []int) int {
	length := len(nums)
	if length <= 1{
		return 0
	}

	curLevelNum,curNode,nextLevelNum,level := 0,0,0,0
	for curLevelNum - curNode + 1 > 0 {
		level++
		for ; curNode <= curLevelNum;curNode++ {
			nextLevelNum = max(nextLevelNum,nums[curNode] + curNode)
			if nextLevelNum >= length-1 {
				return level
			}
		}

		curLevelNum = nextLevelNum
	}

	return -1
}

// -----------------------------
// Greedy solution
// See it as bfs algorithm
func jumpGreedy(nums []int)int {
	length := len(nums)
	if length <= 1 {
		return 0
	}


	var curEnd,curFurthest,level int
	for i := 0; i< length;i++ {
		curFurthest = max(curFurthest,nums[i] + i)
		if i == curEnd {
			level++
			curEnd = curFurthest
		}
	}

	return level
}