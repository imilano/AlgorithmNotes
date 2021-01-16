package leet_163

import "strconv"

/*
Given a sorted integer array nums, where the range of elements are in the inclusive range [lower, upper], return its missing ranges.

Example:

Input: nums = [0, 1, 3, 50, 75], lower = 0 and upper = 99,
Output: ["2", "4->49", "51->74", "76->99"]
*/

func findMissingRanges(nums []int, lower int, upper int) []string {
	var res []string
	for _, v := range nums {
		if v > lower {
			if v-1 > lower {
				res = append(res, strconv.Itoa(lower)+"->"+strconv.Itoa(v-1))
			} else {
				res = append(res, strconv.Itoa(lower))
			}
		}

		if v == upper {
			return res
		}

		lower = v + 1
	}

	if lower <= upper {
		if upper > lower {
			res = append(res, strconv.Itoa(lower)+"->"+strconv.Itoa(upper))
		} else {
			res = append(res, strconv.Itoa(lower))
		}
	}

	return res
}
