package leet_39

/*
	Given an array of distinct integers candidates and a target integer target, return a list of all unique combinations of
	candidates where the chosen numbers sum to target. You may return the combinations in any order. The same number may be
	chosen from candidates an unlimited number of times. Two combinations are unique if the frequency of at least one of the
	chosen numbers is different.
	It is guaranteed that the number of unique combinations that sum up to target is less than 150 combinations for the given input.
*/

import "sort"

func helper(cand []int, cho []int, remain int, start int, res *[][]int) {
	if remain < 0 {
		return
	} else if remain == 0 {
		tmp := make([]int, len(cho))
		copy(tmp, cho)
		*res = append(*res, tmp)
		// you can't add cho directly to result, cause slice cho still point to the same array as before
		// when you say slice a := b, which means address of a is equal to b, unless a or b has enlarged
		//*res = append(*res, cho)
	} else {
		for i := start; i < len(cand); i++ {
			cho = append(cho, cand[i])
			helper(cand, cho, remain-cand[i], i, res)
			cho = cho[:len(cho)-1]
		}
	}
}

func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	if candidates == nil {
		return res
	}

	sort.Ints(candidates)
	helper(candidates, nil, target, 0, &res)
	return res
}

//func CombinationSum(cand []int, target int) [][]int {
//	return combinationSum(cand,target)
//}
