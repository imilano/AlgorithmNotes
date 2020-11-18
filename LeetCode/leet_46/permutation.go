package leet_46

/*
	Given a collection of distinct integers, return all possible permutations.
 */

func contains(nums []int, target int) bool{
	for _,v := range nums {
		if v == target {
			return true
		}
	}
	 return false
}

func backtrace(res *[][]int,nums []int, chosen []int) {
	if len(chosen) == len(nums) {
		tmp := make([]int, len(chosen))
		copy(tmp, chosen)
		*res = append(*res, tmp)
	} else {
		for i := 0; i < len(nums); i++ {
			if contains(chosen,nums[i]) { // every element is unique, so there wont be any duplication
				continue
			}
			chosen = append(chosen,nums[i])
			backtrace(res,nums,chosen)
			chosen = chosen[:len(chosen)-1]
		}
	}
}

func permute(nums []int) [][]int {
	var res [][]int
	if len(nums) == 0 {
		return res
	}

	backtrace(&res,nums,nil)
	return res
}
