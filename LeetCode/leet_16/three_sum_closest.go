package leet_16

import (
	"math"
	"sort"
)

/*
	Given an array nums of n integers and an integer target, find three integers in nums such that the sum is closest to target.
	Return the sum of the three integers. You may assume that each input would have exactly one solution.
 */

func abs(a int) int {
	if a >= 0 {
		return a
	}

	return -a
}
// original solution
func threeSumClosestOriginal(nums []int, target int) int {
	var res int
	if len(nums) <=2 {
		return res
	}

	sort.Ints(nums)
	minDiff := math.MaxInt16
	for i := 0 ; i< len(nums)-2;i++ {
		left,right,sum := i+1,len(nums)-1,target-nums[i]
		for left < right {
			tmp := nums[left] + nums[right]
			if tmp == sum {
				return target
			}
			diff := abs(tmp-sum)
			if diff < minDiff {
				minDiff = diff
				res = nums[i] + nums[left] + nums[right]
			}

			if left < right && tmp < sum {
				left++
			}

			if left < right && tmp > sum {
				right--
			}
		}

	}

	return res
}

// optimized
func threeSumClosest(nums []int, target int) int {
	res := math.MaxInt16   // 初始值不能为0，应该尽量大

	sort.Ints(nums)
	for i := 0 ; i< len(nums)-2;i++ {
		// 因为数组已经有序了，所以如果当前nums[i]的三倍都已经比target要大的话，没必要在用比nums[i]还要大的数字继续进行比较了
		if nums[i] *  3 > target {
			return min(res,nums[i]+ nums[i+1]+nums[i+2])
		}
		left,right := i+1,len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == target {
				return target
			}
			if abs(target-sum) < abs(target-res) {
				res = sum
			}
			if sum > target {
				right--
			} else {
				left++
			}
		}
	}

	return res
}


func min(a,b int) int {
	if a > b {
		return b
	}

	return a
}