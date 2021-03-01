package leet_47

import "sort"

/*
	Given a collection of numbers that might contain duplicates, return all possible unique permutations.
*/

/*
	Given a collection of numbers that might contain duplicates, return all possible unique permutations.
*/

func permuteUnique(nums []int) [][]int {
	res := make([][]int, 0)
	sort.Ints(nums)
	helper(nums, 0, &res)
	return res
}

func helper(nums []int, index int, res *[][]int) {
	if index == len(nums)-1 {
		tmp := make([]int, len(nums))
		copy(tmp, nums)
		*res = append(*res, tmp)
	} else {
		for i := index; i < len(nums); i++ {
			if i == index || nums[i] != nums[index] {
				swap(nums, index, i)
				tmp := make([]int, len(nums))
				copy(tmp, nums)
				helper(tmp, index+1, res)
			}
		}
	}
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

// Original wrong solution
//func permuteUnique(nums []int) [][]int {
//	added := make(map[string]bool)
//	res := make([][]int,0)
//	sort.Ints(nums)
//	helper(nums,nil,0,&res,added)
//	return res
//}
//
//func helper(nums []int,cur []int,index int, res *[][]int,added map[string]bool) {
//	if index == len(nums) {
//		s := toString(cur)
//		if !added[s] {
//			tmp := make([]int, len(cur))
//			copy(tmp,cur)
//			*res = append(*res,tmp)
//			added[s] = true
//		}
//		return
//	}
//
//
//	for i := index;i < len(nums);i++ {
//		if i != index && nums[i] == nums[index] {
//			continue
//		}
//
//		swap(nums,index,i)
//		cur = append(cur,nums[i])
//		tmp := make([]int, len(cur))
//		copy(tmp,cur)
//		helper(nums,tmp,i+1,res,added)
//	}
//}
//
//func swap(arr []int,i,j int) {
//	arr[i],arr[j] = arr[j],arr[i]
//}
//
//func toString(arr []int) string {
//	var res string
//	for _,v := range arr {
//		res += strconv.Itoa(v)
//	}
//
//	return res
//}
