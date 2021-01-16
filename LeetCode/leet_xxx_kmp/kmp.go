package leet_xxx_kmp

// 在src中查找子串pat，如果找到，就返回src中该子串的起始index，否则，返回-1。

// 暴力算法，时间复杂度O(MN)，空间复杂度O(1)
func brutalForceKMP(pat, src string) int {
	M, N := len(pat), len(src)
	for i := 0; i <= N-M; i++ { // 暴力匹配
		var j int
		for j := 0; j < M; j++ { // 逐一和模式串进行比较
			if pat[j] != src[i+j] {
				break
			}
		}
		if j == M {
			return i
		}
	}
	return -1
}

// 使用DFA的KMP。
// 算法分为两个部分，一个部分是计算出dp数组，这个函数只需要用到模式串pat；第二部分是待查找字符串str中是否存在匹配的pat，此时只需要用到str和dp

// 计算dp数组
var dp [][]int

func dfa(pat string) {
	M := len(pat)

	// dp[i][j]=m; i表示当前状态，j表示遇到的字符串,m表示应该转移到的状态
	dp = make([][]int, M)
	for i := 'A'; i < 'A'+26; i++ {
		dp[i] = make([]int, 26)
	}

	// base condition
	// 状态0碰上pat[0]时应该转移到转态1
	dp[0][pat[0]] = 1

	// 标识转移转态
	x := 0

	// 转态转移由状态和字符造成，并且状态的计算只和pat有关。
	// dp[i][j]=m，其中i表示状态，j表示字符，m表示j状态碰到j字符应该转移到状态m
	var j uint8
	for i := 1; i < M; i++ {
		for j = 'A'; j < 'A'+26; j++ { // 大写字符
			// 如果匹配成功
			if pat[i] == j {
				dp[i][j] = i + 1
			} else {
				// 当前字符匹配失败，交给具有公共前缀的状态处去处理
				dp[i][j] = dp[x][j]
			}
			// 更新x的状态，x总是比匹配的字符串j慢一个位置，因而x总是会有值
			x = dp[x][j]
		}
	}
}

func search(s string) int {
	N := len(s)

	var j int
	for i := 0; i < N; i++ { // 逐一进行匹配
		j = dp[j][s[i]] // 计算下一个状态
		if j == N {     // 匹配成功，返回起始位置
			return i - N + 1
		}
	}

	// 匹配失败，返回-1
	return -1
}
