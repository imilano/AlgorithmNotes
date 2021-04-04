package _5_第一个只出现一次的字符

/*
请实现一个函数用来找出字符流中第一个只出现一次的字符。例如，当从字符流中只读出前两个字符 "go" 时，第一个只出现一次的字符是 "g"。当从该字符流中读出前六个字符“google" 时，第一个只出现一次的字符是 "l"。
*/

var cnt map[uint8]int
var que []uint8


// 使用一个数组/字典记录字符出现的次数，使用一个队列记录当前已经出现的字符。当插入字符时，如果一个字符出现次数大于1，那么应该在队列的头部删除这些元素，
//这样在进行查找时，只需要队列的第一个元素即可
func Insert(c uint8) {
	cnt[c]++
	que = append(que,c)

	for len(que) != 0 && cnt[que[0]] > 1 {
		que = que[1:] // 注意这里不应该在cnt中删除c，因为字符流是连续的，一个字符一旦出现次数大于1，那么剧再也不是只出现一次的字符了
	}
}

func FirstAppearingOnce() uint8 {
	if len(que) == 0 {
		return ' '
	} else {
		return que[0]
	}
}