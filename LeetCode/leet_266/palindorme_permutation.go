package leet_266

/*
Given a string, determine if a permutation of the string could form a palindrome.

Example 1:

Input: "code"
Output: false
Example 2:

Input: "aab"
Output: true
Example 3:

Input: "carerac"
Output: true
Hint:

Consider the palindromes of odd vs even length. What difference do you notice?
Count the frequency of each character.
If each character occurs even number of times, then it must be a palindrome. How about character which occurs odd number of times?
*/

// 根据提示，如果所有字符出现的次数都是偶数次，那么该字符串的排列一定可以构成回文串；如果有字符出现奇数次，那么这样的字符必须只能出现一次
func canPermutePalindrome(s string) bool {
	if s == "" {
		return true
	}

	m := make(map[uint8]int)
	for i := range s {
		m[s[i]]++
	}
	//var even,odd int
	var odd int
	for _,v := range m {
		if v %2 == 1 {
			odd++
		}
		//if v %2 == 0 {
		//	even++
		//} else {
		//	odd++
		//}
	}

	// 不用计算偶数的出现次数，既然是偶数次出现，那么有多少字符偶数次出现根本没有关系
	//return even%2 == 0 && (odd == 0 || len(s) %2 == 1 && odd == 1)
	return odd ==0 || len(s) % 2 == 1 && odd == 1  // 只会有两种情况，一种是没有奇数出现；另一种是字符串长度为奇数，并且奇数只出现一次。
}

// 使用set，遍历字符串，如果一个字符不在set中，则加入该字符，否则删除该字符；最后如果set中没有字符，或者只有一个字符，则说明是回文串。原理与上相同
func canPermutePalindrome2(s string) bool  {
	m := make(map[uint8]bool)
	for i := range s {
		if _,ok := m[s[i]]; !ok {
			m[s[i]] = true
		} else {
			delete(m,s[i])
		}
	}

	return len(m) == 0 || len(m) == 1
}
