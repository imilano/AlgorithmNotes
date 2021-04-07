package golang

// 让n个小朋友们围成一个大圈（小朋友编号从0到n-1）。然后，随机指定一个数 m，让编号为 0 的小朋友开始报数。每次喊到 m-1 的那个小朋友要出列唱首歌，
// 然后可以在礼品箱中任意的挑选礼物，并且不再回到圈中，从他的下一个小朋友开始，继续 0...m-1 报数 .... 这样下去 .... 直到剩下最后一个小朋友，
// 可以不用表演。问最后一个小朋友的编号是什么？如果不存在的话，返回-1

//问题简化：给定一个长度为n的序列，每次向后数m个元素并删除，那么最终留下的是第几个元素。这个问题有拆解为子问题的潜质：如果我们知道了n-1个元素中
//会删除的是哪个元素，那么我们就知道n个元素中删除的是哪个元素。我们定义子问题f(n,m)=y，表示从n个元素中向后数m个元素并删除，最终剩下的数的下标是
//y。y是下标就意味着从index=0开始数，数y+1个数。

//方法的核心在于：只关心最终或者的那个人的序号变化情况，从最终结果反推出规律。参见 https://dyfloveslife.github.io/2019/12/27/offer-Josephus/
// 可得到递推公式：f(n,m) = (f(n-1,m) + m) % n. 返回条件，当n=1时，必然返回0
// 时间复杂度O（n)，空间复杂度O(n)
func LastRemaining_Solution( n int ,  m int ) int {
	// write code here
	if n == 0 {
		return -1
	}

	if n == 1 {
		return 0
	}

	x := LastRemaining_Solution(n-1,m)
	return (x+m)%n
}

// 迭代版本
func lastRemaining(n int, m int) int {
	x := -1
	for i := 1; i <n+1; i++ {
		x = (x+m)%i
	}

	return x
}
