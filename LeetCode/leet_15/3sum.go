package leet_15

import (
	"sort"
	"strconv"
)

/*
	Given an array nums of n integers, are there elements a, b, c in nums such that a + b + c = 0?
	Find all unique triplets in the array which gives the sum of zero.
	Notice that the solution set must not contain duplicate triplets.
*/

// brutal force
func threeSum1(nums []int) [][]int {
	var res [][]int
	if len(nums) <= 2 {
		return res
	}

	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[left] + nums[right]
			if sum == -nums[i] && !contains(nums[i], nums[left], nums[right], res) {
				res = append(res, []int{nums[i], nums[left], nums[right]})
			}

			if sum < -nums[i] {
				left++
			} else {
				right--
			}
		}
	}

	return res
}

func contains(i, j, k int, res [][]int) bool {
	for _, s := range res {
		if s[0] == i && s[1] == j && s[2] == k {
			return true
		}
	}

	return false
}

//-------------------------------------------------
// using  map to reduce duplication
func threeSum2(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)
	m := make(map[string]int)

	length := len(nums)
	for i := 0; i < length-2; i++ {
		start, end := i+1, length-1
		for start < end {
			sum := nums[start] + nums[end]
			if sum == -nums[i] {
				t := strconv.Itoa(nums[start]) + strconv.Itoa(nums[end])
				if m[t] == 1 {
					start++
					continue
				}

				res = append(res, []int{nums[i], nums[start], nums[end]})
				m[t] = 1
				start++
			} else if sum > -nums[i] {
				end--
			} else {
				start++
			}
		}
	}

	return res
}

//----------------------------------
func threeSum3(nums []int) [][]int {
	var res [][]int
	if len(nums) <= 2 {
		return res
	}

	sort.Ints(nums)
	l := len(nums)
	for i := 0; i < l-2; i++ {
		if i == 0 || (i > 0 && nums[i] != nums[i-1]) { // deduplication
			left, right, sum := i+1, l-1, -nums[i]
			for left < right {
				tmp := nums[left] + nums[right]
				if tmp == sum {
					res = append(res, []int{nums[i], nums[left], nums[right]})

					// skip duplicated num
					for left < right && nums[left] == nums[left+1] {
						left++
					}

					for left < right && nums[right] == nums[right-1] {
						right--
					}
					left++
					right--
				} else if tmp < sum {
					left++
				} else {
					right--
				}
			}

		}
	}

	return res
}
