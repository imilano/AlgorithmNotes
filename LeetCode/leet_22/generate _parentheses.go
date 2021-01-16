package leet_22

/*
	Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.
*/

//核心要点就在于，尽量先放置左括号，任意时刻，左括号的数量一定要大于等于右括号的数量，不然不可能形成合法字符

//-------------------------
// time limit exceeded
// 采用枚举法
func generateParenthesis(n int) []string {
	var res []string
	if n == 0 {
		return res
	}

	generate("", n*2, &res)
	return res
}

// 枚举办法，枚举每个位置出现的可能
func generate(cur string, num int, res *[]string) {
	if len(cur) == num {
		if isValidate(cur) {
			*res = append(*res, cur)
		}
	} else {
		cur += "("
		generate(cur, num, res)
		cur += ")"
		generate(cur, num, res)
	}
}

// 一个括号序列是合理的，可以通过这个方法进行验证：把"（"看做+1，"）"看做-1，任一时刻，括号序列的前缀和应该大于0，如果不大于0，那么当前序列就不是合法序列
func isValidate(s string) bool {
	var count int
	for _, c := range s {
		if c == rune('(') {
			count++
		} else {
			count--
		}

		if count < 0 {
			return false
		}
	}

	return true
}

//-----------------------------
// 采用回溯法
// 时间复杂度与卡特兰数相关
// We keep track of opening and closing brackets we have placed so far. We can start an opening bracket if we still have one remaining "(" to place,
// And we start a closing brackets if it would not exceed the number of opening brackets
func generateParenthesis2(n int) []string {
	var res []string
	if n == 0 {
		return res
	}

	backtrace("", 0, 0, n, &res)
	return res
}

func backtrace(cur string, open int, close int, max int, res *[]string) {
	if len(cur) == max*2 {
		*res = append(*res, cur)
		return
	}

	// 优先添加左括号
	if open < max {
		backtrace(cur+"(", open+1, close, max, res)
	}

	if close < open {
		backtrace(cur+")", open, close+1, max, res)
	}
}

//-----------------------
// dfs
func generateParenthesis3(n int) []string {
	var res []string
	dfs(n, n, "", &res)
	return res
}

// left,right 分别代表剩余的左右括号数
// 尽可能先放左括号，任意时刻，左括号的数量一定要大于等于右括号数，才可能会出现合法的情况
func dfs(left int, right int, cur string, res *[]string) {
	// 如果左括号剩余数大于右括号剩余数，说明字符串中右括号数大于左括号数，从而出现非法情况，剪枝
	if left > right {
		return
	}

	if left == 0 && right == 0 {
		*res = append(*res, cur)
	}

	if left > 0 {
		dfs(left-1, right, cur+"(", res)
	}

	if right > 0 {
		dfs(left, right-1, cur+")", res)
	}
}
