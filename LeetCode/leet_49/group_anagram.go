package leet_49

import (
	"sort"
	"strings"
)

/*
Given an array of strings strs, group the anagrams together. You can return the answer in any order.
An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

Example 1:
Input: strs = ["eat","tea","tan","ate","nat","bat"]
Output: [["bat"],["nat","tan"],["ate","eat","tea"]]


将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。
*/

// NOTE golang is definitely lack of data structure
// time complexity: O(N KlogK), where N is the number of input string, K is maximum length of all the strings
// space complexity: O(N K)
func groupAnagrams(strs []string) [][]string {
	var res [][]string
	if len(strs) == 0 {
		return res
	}

	m := make(map[string][]string)
	for i := range strs {
		tmp := strings.Split(strs[i], "")
		// can use counting sort to reduce time complexity from O(KlgK) to O(K)
		sort.Strings(tmp)
		s := strings.Join(tmp, "")

		if _, ok := m[s]; !ok {
			m[s] = make([]string, 0)
		}
		m[s] = append(m[s], strs[i])
	}

	for _, v := range m {
		res = append(res, v)
	}

	return res
}

// Original wrong ans
//func groupAnagrams(strs []string) [][]string {
//	var res [][]string
//	length := len(strs)
//	if length == 0 {
//		return res
//	}
//
//
//	sort.Strings(strs)
//	var dedup []string
//	for i:=1; i< length;i++ {
//		if strs[i] == strs[i-1] {
//			continue
//		}
//		dedup = append(dedup,strs[i])
//	}
//
//	for i := 0; i< len(dedup);i++ {
//		t := []string{dedup[i]}
//		for j := i+1; j<len(dedup);j++ {
//			if isSame(dedup[i],dedup[j]) {
//				t = append(t,dedup[j])
//			}
//		}
//
//		res = append(res, t)
//	}
//
//	return res
//}

//func isSame(s1,s2 string) bool {
//	if len(s1) != len(s2) {
//		return false
//	}
//
//	m := make(map[uint8]bool)
//	for i := 0; i< len(s1);i++ {
//		m[s1[i]] = true
//	}
//
//	for i := range s2 {
//		if m[s2[i]] == false {
//			return false
//		}
//	}
//
//	return true
//}
