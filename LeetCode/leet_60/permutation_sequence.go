package leet_60

import (
	"strconv"
	"strings"
)

/*
The set [1, 2, 3, ..., n] contains a total of n! unique permutations.

By listing and labeling all of the permutations in order, we get the following sequence for n = 3:

    "123"
    "132"
    "213"
    "231"
    "312"
    "321"

Given n and k, return the kth permutation sequence.
*/

func getPermutation(n int, k int) string {
	// get factor
	factor := make([]int, n+1)
	factor[0] = 1
	f := 1
	for i := 1; i <= n; i++ {
		f *= i
		factor[i] = f
	}

	// get index
	var nums []string
	for i := 1; i <= n; i++ {
		nums = append(nums, strconv.Itoa(i))
	}

	// index start from 0
	k--
	var res strings.Builder
	for i := 1; i <= n; i++ {
		index := k / factor[n-i]
		res.WriteString(nums[index])
		nums = append(nums[:index], nums[index+1:]...)
		k -= index * factor[n-i]
	}

	return res.String()
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func helper(nums []int, index int, res *[][]int) {
	if index == len(nums)-1 {
		tmp := make([]int, len(nums))
		copy(tmp, nums)
		*res = append(*res, tmp)
	} else {
		for i := index; i < len(nums); i++ {
			swap(nums, index, i)
			tmp := make([]int, len(nums))
			copy(tmp, nums)
			helper(tmp, index+1, res)
		}
	}
}

// wrong answer, too many edge case
func getPermutationOriginal(n int, k int) string {
	if n == 1 && k == 1 {
		return strconv.Itoa(n)
	}
	sum := 1
	for i := 1; i <= n-1; i++ {
		sum *= i
	}

	mul := k / sum
	rem := k % sum

	start := 1 + mul
	var str []int
	for i := 1; i <= n; i++ {
		if i == start {
			continue
		}
		str = append(str, i)
	}

	var res [][]int
	helper(str, 0, &res)

	var tmp string
	for _, v := range res[rem-1] {
		tmp += strconv.Itoa(v)
	}
	return strconv.Itoa(start) + tmp
}
