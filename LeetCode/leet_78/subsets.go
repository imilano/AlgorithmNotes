package leet_78

/*
	Given a set of distinct integers, nums, return all possible subsets (the power set).
	Note: The solution set must not contain duplicate subsets.
 */

func backtrace(res *[][]int, nums []int, chosen []int, start int) {
	tmp := make([]int,len(chosen))
	copy(tmp,chosen)
	*res = append(*res, tmp)

	for i:= start;i <len(nums);i++ {
		chosen  = append(chosen,nums[i])
		backtrace(res,nums,chosen,i+1)
		chosen = chosen[:len(chosen)-1]
	}
}

func subsets(nums []int) [][]int {
	var res [][]int
	if len(nums) == 0 {
		return res
	}

	backtrace(&res,nums,nil,0)
	return res
}
