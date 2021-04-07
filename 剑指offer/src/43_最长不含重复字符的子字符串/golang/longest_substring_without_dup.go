package golang

// 输入一个字符串（只包含 a~z 的字符），求其最长不含重复字符的子字符串的长度。例如对于 arabcacfr，最长不含重复字符的子字符串为 acfr，
// 长度为 4。

// 滑动窗口,使用map保存之前出现的位置
func LongestSubstringWithoutDup(str string) int {
	m := make(map[uint8]int)
	start,res := 0,0
	n := len(str)
	for i := 0; i < n; i++ {
		if m[str[i]] >= start {
			start = m[str[i]] + 1
		}
		m[str[i]] = i
		res = max(res, i-start+1)
	}

	return res
}


// 使用双向队列，把元素不断的加入到队列中，如果有相同的元素，就把队首的元素移除，这样就可以保证队列中永远没有重复的元素
func LongestSubstringWithoutDup2(str string) int {
	var que []uint8
	var res int

	n := len(str)
	for i := 0; i < n;i++ {
		for len(que) != 0 && contains(que,str[i]) {  // 把c之前的字符都出队
			que = que[1:]
		}
		que = append(que,str[i])
		res = max(res, len(que))
	}


	return res
}

func contains(que []uint8,c  uint8) bool {
	for i := range que {
		if que[i] == c {
			return true
		}
	}

	return false
}
func max(a,b int) int {
	if a > b {
		return a
	}

	return b
}