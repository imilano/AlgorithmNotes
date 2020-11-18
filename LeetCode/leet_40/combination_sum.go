package leet_40

import "sort"

/*
	Given a collection of candidate numbers (candidates) and a target number (target), find all unique combinations in candidates where the candidate numbers sum to target.
	Each number in candidates may only be used once in the combination.
	Note: The solution set must not contain duplicate combinations.
 */

func helper(cand []int,cho []int,remain int,start int, res *[][]int) {
	if remain < 0 {
		return
	} else if remain == 0 {
		tmp := make([]int,len(cho))
		copy(tmp,cho)
		*res = append(*res,tmp)
	} else  {
		for i := start; i< len(cand);i++ {
			if i> start && cand[i] == cand[i-1] {
				continue
			}
			cho = append(cho,cand[i])
			helper(cand,cho,remain-cand[i],i+1,res)
			cho = cho[:len(cho)-1]
		}
	}
}

func combinationSum2(candidates []int, target int) [][]int {
	var res [][]int
	if candidates == nil {
		return res
	}

	sort.Ints(candidates)
	helper(candidates,nil,target,0,&res)
	return res
}