package leet_14

import "sort"

/*
	Write a function to find the longest common prefix string amongst an array of strings.
	If there is no common prefix, return an empty string "".
 */

// 给字符串数组排序，由于排序是按照字母相对顺序进行的，而不是按照长度进行，故而具有相同字符多的字符串会被聚到一起，而跟大家相同字符越少的就会
// 被挤到两端。那么如果有共同前缀的话，一定会出现在首尾两端的字符串中。故而只需要找到首尾两端字符串的共同长度即可。
func min(a,b int) int {
	if a > b {
		return b
	}

	return a
}

func longestCommonPrefix2(strs []string) string {
	length := len(strs)
	if length == 0 {
		return ""
	}

	sort.Strings(strs)
	var index int
	first,last := strs[0],strs[length-1]
	minLen := min(len(first),len(last))   // 不知道到底谁才是最短的
	for index < minLen &&  first[index] == last[index] {
		index++
	}

	return first[:index]
}

func longestCommonPrefix(strs []string) string {
	var index int

	if len(strs) == 0 {
		return ""
	}

	sort.Strings(strs)
	for i := 0; i< len(strs)-1;i++ {
		if len(strs[i]) == 0 && len(strs[i+1]) != 0{
			return ""
		}
	}

	for index != len(strs[0]){
		c := strs[0][index]
		for i := range strs {
			if strs[i][index] != c {
				return strs[0][:index]
			}
		}

		index++
	}

	return strs[0]
}
