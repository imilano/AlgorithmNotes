package leet_55

/*
	Given an array of non-negative integers, you are initially positioned at the first index of the array.
	Each element in the array represents your maximum jump length at that position.
	Determine if you are able to reach the last index.
*/

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}

// time limit exceeded
const (
	_ = iota
	GOOD
	BAD
	UNKNOWN
)

func canJumpFromPositionRec(pos int, nums []int, memo []int) bool {
	if memo[pos] != UNKNOWN {
		if memo[pos] == GOOD {
			return true
		} else {
			return false
		}
	}

	furtherJump := min(pos+nums[pos], len(nums)-1)
	// we always try to make the biggest jump such that we reach the end as soon as possible
	//for start := pos+1; start <= furtherJump; start++ {
	for start := furtherJump; start > pos; start-- {
		if canJumpFromPositionRec(start, nums, memo) {
			memo[pos] = GOOD
			return true
		}
	}
	memo[pos] = BAD
	return false

}

func canJumpRec(nums []int) bool {
	memo := make([]int, len(nums))
	for i := range memo {
		memo[i] = UNKNOWN
	}
	memo[len(memo)-1] = GOOD
	return canJumpFromPositionRec(0, nums, memo)
}

// TODO refer to solution and write a post to learn dynamic programming
// greedy algorithm
func canJump(nums []int) bool {
	maxJump := len(nums) - 1
	for i := len(nums) - 2; i >= 0; i-- {
		if i+nums[i] >= maxJump {
			maxJump = i
		}
	}

	return maxJump == 0
}
