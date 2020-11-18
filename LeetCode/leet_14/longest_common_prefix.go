package leet_14

import "sort"

/*
	Write a function to find the longest common prefix string amongst an array of strings.
	If there is no common prefix, return an empty string "".
 */

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
