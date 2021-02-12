package leet_249

import "sort"

/*
Given a string, we can "shift" each of its letter to its successive letter, for example: "abc" -> "bcd". We can keep "shifting" which forms the sequence:

"abc" -> "bcd" -> ... -> "xyz"
Given a list of strings which contains only lowercase alphabets, group all strings that belong to the same shifting sequence.

For example, given: ["abc", "bcd", "acef", "xyz", "az", "ba", "a", "z"],
Return:

[
  ["abc","bcd","xyz"],
  ["az","ba"],
  ["acef"],
  ["a","z"]
]
Note: For the return value, each inner list's elements must follow the lexicographic order.
*/
// 要达成偏移字符串，就需要利用一个特点：字符串的每个字母和首字符的相对距离都是相等的。用 abc和efg为例，b到a的距离为1，c到a的距离为2，而f到e的距离为1，g到e的距离为2；
// 需要注意的问题就是，注意这里的偏移如何定义，对于az和yx来说，二者的偏移是一样的，所以计算偏移量时，需要采用(c+26-a[0])%26取余的方法
func groupStrings(strs []string) [][]string {
	var res [][]string
	m := make(map[string][]string)
	for _, s := range strs {
		idx := ""
		for i := range s {
			idx += string((s[i]+26-s[0])%26) + "."
		}
		m[idx] = append(m[idx], s)
	}

	for _, v := range m {
		sort.Strings(v)
		res = append(res, v)
	}

	return res
}
