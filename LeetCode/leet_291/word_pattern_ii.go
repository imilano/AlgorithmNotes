package leet_291

/*
Given a pattern and a string str, find if strfollows the same pattern.

Here follow means a full match, such that there is a bijection between a letter in pattern and a non-empty substring in str.

Example 1:

Input: pattern = "abab", str = "redblueredblue"
Output: true
Example 2:

Input: pattern = pattern = "aaaa", str = "asdasdasdasd"
Output: true
Example 3:

Input: pattern = "aabb", str = "xyzabcxzyabc"
Output: false
Notes:
You may assume both pattern and str contains only lowercase letters.
*/

// 回溯法
func wordPattern(pattern string, s string) bool {
	m :=  make(map[uint8]string)
	return helper(pattern,0,s,0,&m)
}


func helper(pattern string, p int, str string, r int, m *map[uint8]string) bool {
	if p ==  len(pattern) && len(str) == r {  // 如果二者都匹配到了末尾，则返回true
		return true
	}
	if p == len(pattern)  || r == len(str) {  // 如果二者只有一方匹配到了末尾，则返回false
		return false
	}

	c := pattern[p]
	for i := r; i < len(str);i++ {  // 回溯法，尝试每一种单词组合
		s := str[r:i+1]
		if v,ok := (*m)[c]; ok && s == v {  // 如果c已经存在并且映射值与v相等，那么如果pattern的下一个字符和str的剩下字符也匹配，那么返回true
			if helper(pattern,p+1,str,i+1,m) {
				return true
			}
		} else if _,ok := (*m)[c];!ok { // 如果c不存在映射值，那么检查当前m中s是否已经被匹配，如果被匹配，则什么也不做；
			var existed bool
			for _,v := range *m {
				if v == s {
					existed = true
				}
			}

			if !existed { // 如果没有被匹配，那么将当前c和s作为新的匹配项放入m，然后检查pattern余下字符和str余下字符是否匹配，如果匹配，返回true；如果不匹配，那么不要忘记从m中删除这种错误匹配
				(*m)[c] = s
				if helper(pattern,p+1,str,i+1,m) {
					return true
				}

				delete(*m,c)
			}
		}
	}

	return false
}
