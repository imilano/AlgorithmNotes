package leet_xxx_eggdrop


func getMax(i,j int) int {
	if i > j {
		return  i
	}

	return i
}

// 递归办法
func eggDrop(K int, N int) int {
	if K == 1 {
		return N
	}

	if N == 0 {
		return 0
	}

	var res int

	for i := 1; i<= K; i++ {
		res = getMax(res,getMax(eggDrop(i-1,N-1),eggDrop(K,N-i))+1)  // 在第i层扔鸡蛋的结果
	}

	return res
}

// 备忘录方法，使用伪代码演示
/*
def superEggDrop(K: int, N: int):

    memo = dict()
    def dp(K, N) -> int:
        # base case
        if K == 1: return N
        if N == 0: return 0
        # 避免重复计算
        if (K, N) in memo:
            return memo[(K, N)]

        res = float('INF')
        # 穷举所有可能的选择
        for i in range(1, N + 1):
            res = min(res,
                      max(
                            dp(K, N - i),
                            dp(K - 1, i - 1)
                         ) + 1
                  )
        # 记入备忘录
        memo[(K, N)] = res
        return res

    return dp(K, N)
 */

