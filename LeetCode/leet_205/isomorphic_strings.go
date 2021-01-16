package leet_205

/*
Given two strings s and t, determine if they are isomorphic.

Two strings are isomorphic if the characters in s can be replaced to get t.

All occurrences of a character must be replaced with another character while preserving the order of characters.
No two characters may map to the same character but a character may map to itself.

Example 1:

Input: s = "egg", t = "add"
Output: true
Example 2:

Input: s = "foo", t = "bar"
Output: false
*/

// wrong
func isIsomorphic(s string, t string) bool {
	len1, len2 := len(s), len(t)
	if len1 != len2 {
		return false
	}
	m := make(map[uint8]uint8)

	for i := 0; i < len1; i++ {
		v, ok := m[t[i]]
		if ok && v == t[i]-s[i] {
			continue
		} else if ok && v != t[i]-s[i] {
			return false
		} else if !ok {
			m[t[i]] = t[i] - s[i]
		}
	}

	return true
}

// 采用定长数组
func isIsomorphic2(s string, t string) bool {
	length := len(s)
	m1, m2 := make([]int, 256), make([]int, 256)
	for i := 0; i < length; i++ {
		if m1[s[i]] != m2[t[i]] {
			return false
		}

		m1[s[i]] = i + 1
		m2[t[i]] = i + 1
	}

	return true
}
