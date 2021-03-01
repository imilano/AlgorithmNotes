package leet_179

import (
	"sort"
	"strconv"
)

/*
Given a list of non-negative integers nums, arrange them such that they form the largest number.
Note: The result may be very large, so you need to return a string instead of an integer.
*/

func largestNumber(nums []int) string {
	var res string
	var tmp []string
	for _, v := range nums {
		tmp = append(tmp, strconv.Itoa(v))
	}

	//sort.Strings(tmp)
	// 需要自定义排序
	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i]+tmp[j] > tmp[j]+tmp[i]
	})
	for _, v := range tmp {
		res += v
	}

	// 处理全0的情况
	if res[0] == '0' {
		return "0"
	}
	return res
}
