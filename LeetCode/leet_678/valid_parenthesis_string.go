package leet_678

/*
Given a string containing only three types of characters: '(', ')' and '*', write a function to check whether this string is valid. We define the validity of a string by these rules:

Any left parenthesis '(' must have a corresponding right parenthesis ')'.
Any right parenthesis ')' must have a corresponding left parenthesis '('.
Left parenthesis '(' must go before the corresponding right parenthesis ')'.
'*' could be treated as a single right parenthesis ')' or a single left parenthesis '(' or an empty string.
An empty string is also valid.
Example 1:
Input: "()"
Output: True
Example 2:
Input: "(*)"
Output: True
Example 3:
Input: "(*))"
Output: True
*/
//----------------------------
//两次遍历，采用类似于括号合法性检测的方法。第一次将所有的*都当做左括号，然后，遇到左括号，left++，遇到右括号，left--。检测过程中，一旦出现left < 0 的情况，
//说明此时即便将所偶的*都当做左括号，右括号的数量还是比左括号多，则非法，返回false。检测到最后，如果left=0，返回true；如果left>0，说明左括号太多了，
//但我们还不知道应该返回true还是false，因为此时会有两种情况，多出来的左括号原本要么是左括号，要么是星号变的。接下来需要排除二者之一。
//
//这次从右边开始向左遍历，把所有的星号都当做右括号。遍历时，遇到右括号，right++，遇到左括号，right--；如果某个时刻，right小于0了，那么说明到目前为止左括号太多了，
//远多于右括号加上星号的值，则返回false；最后，如果right为0，那么返回true；如果right大于0，说明右括号多了，此时要么原本右括号就多，要么多出来的右括号都是星号变的。
//由于上一次遍历时，left大于0，可知右括号并不多，此时right也大于0，说明第一次遍历时多出来的左括号都是星号变的，只需要将这些星号都视为0，那么就合法了，故而可以直接返回true。

func checkValidString(s string) bool {
	var left, right int
	n := len(s)

	for i := 0; i < n; i++ {
		if s[i] == '(' || s[i] == '*' {
			left++
		} else {
			left--
		}

		if left < 0 {
			return false
		}
	}

	for i := n - 1; i >= 0; i-- {
		if s[i] == ')' || s[i] == '*' {
			right++
		} else {
			right--
		}

		if right < 0 {
			return false
		}
	}

	return true
}

//-------------------------------
//递归，遇到星号时，分别将星号视为左括号和右括号以及空，只要三者中有一个为true，就可以视为true。详见代码
func checkValidString2(s string)  bool {
	return helper(s,0,0)
}

func helper(s string, start int, cnt int) bool {
	if cnt < 0 {
		return false
	}

	for i := start; i< len(s);i++ {
		if s[i] == '(' {
			cnt++
		} else if s[i] == ')' {
			if cnt <= 0 {
				return false
			}

			cnt--
		} else {
			return helper(s,i+1,cnt) || // 将星号视为空
				helper(s,i+1,cnt+1) || // 将星号视为左括号
				helper(s,i+1,cnt-1)  // 将星号视为右括号
		}
	}

	return cnt == 0  // 循环退出后，如果cnt为0，则返回true
}