package leet_03

/**
 * Given a string s, find the length of the longest substring without repeating characters.
 **/

/*
	滑动窗口。检查不重复的关键点在于，我们需要确保i和j中间没有重复字符。如果确保呢？我们可以采用集合这个数据结构（集合元素具有唯一性），如果
	滑动窗口的j向前滑动到j+1时，该位置的字符不在set中，那么我们将该字符加入set，更新max值，然后j继续向前滑动；如果该位置的字符应在set中了，
	那么我们呢就收缩窗口，从集合中删除i位置的元素，把i向前滑动。
*/

func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}

// 时间复杂度O（n），最坏情况下，遍历string每个元素2次，所以最多遍历2n次，n为字符串长。
// 空间复杂度O(min(m,k))。空间复杂度受到两个因素影响，一个是字符串长度m，一个是字符集长度k。空间复杂度为而这种较小者。
func lengthOfLongestSubstring(s string) int {
	var ans, right, left int
	m := make(map[uint8]int8)
	l := len(s)

	for left < l && right < l { // 注意循环条件，并非只是简单的right < length 即可，left在不断的增加，会超过length的值
		_, ok := m[s[right]]
		if !ok {
			m[s[right]] = 1
			right++
			// 函数调用会增加一定的开销
			//ans = max(ans,right-left)
			if ans < right-left {
				ans = right - left
			}
		} else {
			delete(m, s[left])
			left++
		}
	}

	return ans
}

/*
	时间空间复杂度均可再度进行优化。
	时间复杂度：
		最坏情况下，遍历2n次。但是，我们无需让i每次只递增1。假设从i开始，恰好到j位置出现了重复，并且假设该重复的位置是k，那么最大不重复的子串必然
		不会出现在i和k之间，我们可以直接让i跳到k+1的位置，然后让j继续向前，重复之前的操作即可。
	空间复杂度：
		如果字符串只由英文字母组成，那么开一个26个元素的数组即可；如果该字符由ASCII码组成，那么开辟一个128大小的数组即可；简言之，如果字符集很小，
		那么我们只需要使用很小的固定数组就可以代替集合类型。
*/

func lengthOfLongestSubstringOptimized(s string) int {
	var left, right, ans, index,gap int
	m := make([]int,128)
	length := len(s)

	for ; right < length; right++ {
		index = m[s[right]]
		if index > left {
			left = index
		}

		gap = right - left + 1
		if gap > ans {
			ans = gap
		}

		m[s[right]] = right + 1
	}
	return ans
}

func LengthOfLongestSubstringOptimized(s string) int {
	return lengthOfLongestSubstringOptimized(s)
}

func LengthOfLongestSubstring(s string) int {
	return lengthOfLongestSubstring(s)
}
