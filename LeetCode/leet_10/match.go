package leet_10

// 根据递归方法来进行匹配
// 实际上就是一个状态机的思想。两个要点，.可以匹配任意一个字符，*可以匹配前一个字符0次或多次。从而总共有三种情况：
//	- pat[i] == txt[j]， 这是正常匹配，那么继续进行匹配：recMatch(i+1,j+1)
//	- pat[i] == .,由于点可以匹配任意字符，所以直接后移即可: recMatch(i+1,j+1)
//	- pat[i] == *, 这里实际可以分为两种情况，一种是匹配0个字符，一种是匹配多个字符。而由于我们采用了递归的方法，匹配多个字符的情况等价于匹配一个字符的情况：
// 			- 匹配0个字符的情况： recMatch(i+1,j)
//			- 匹配一个字符的情况： recMatch(i,j+1)
func recMatch(t,p string) bool {
	// 结束条件，pat遍历结束。此时如果txt也遍历结束，那么匹配匹配成功，否则匹配失败
	if len(p) == 0 {
		return len(t) == 0
	}


	// 直接相等或者模式串出现.字符
	first := len(t) != 0 && (p[0] == t[0] || p[0] == '.')
	// 处理*的零次或多次匹配
	if len(p) >= 2 && p[1] == '*' {
		return recMatch(t,p[2:]) || // 0次匹配，txt不动，pat前移2,跳过*和*前面的字符
			first && recMatch(t[1:],p)  // 1次匹配，txt前移，pat不动，*留下来继续进行下一次匹配
	} else {
		// 正常匹配
		return first && recMatch(t[1:],p[1:])
	}

}

// TAG NFA
// TAG DP
// 递归的方法中存在很多的重复计算，所以可以采用带备忘录的方法进行优化
func matchWithMemo() {
/*
	def isMatch(text, pattern) -> bool:
       memo = dict() # 备忘录
       def dp(i, j):
           if (i, j) in memo: return memo[(i, j)]
           if j == len(pattern): return i == len(text)

           first = i < len(text) and pattern[j] in {text[i], '.'}

           if j <= len(pattern) - 2 and pattern[j + 1] == '*':
               ans = dp(i, j + 2) or \
                       first and dp(i + 1, j)
           else:
               ans = first and dp(i + 1, j + 1)

           memo[(i, j)] = ans
           return ans

       return dp(0, 0)
 */
}

func isMatch(s string, p string) bool {
	return recMatch(s,p)
}

// For testing
func IsMatch(s string, p string) bool {
	return recMatch(s,p)
}