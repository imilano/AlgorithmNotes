package lab5

import (
	"fmt"
	"sort"
	"strconv"
)

/*
 * 给定n个物品，，每个物品的重量为w[i]，价值为v[i]。给定一个容量为c的背包，求该背包可以装下的最大价值。注意，每个物品只有一件，只能选择放或者选择不放。
 * 0/1背包的贪心方案可以有三种方式，
	- 每次装入性价比最高的物品
	- 每次装入的是当前可选择的物品中，价值最高的
	- 每次装入的是当前可选择的物品中，重量最轻的

按照性价比的装入方案，过程如下：
	- 将物品按照性价比从达到小进行排序，依照贪心策略，将尽可能多的单位重量价值最高的物品装入背包。
	- 若将这种物品装入背包之后，背包的内的物品重量没有超过Cap，则选择重量次高的物品尽可能多的装入背包。
	- 依照次策略一直进行下去，一直到背包装满为止。


*/

// util function
type item struct {
	weight    int
	value     int
	valueUnit float32
}

type items []item

func (it items) Len() int {
	return len(it)
}

func (it items) Swap(i int, j int) {
	it[i], it[j] = it[j], it[i]
}

func (it items) Less(i int, j int) bool {
	return it[i].valueUnit > it[j].valueUnit
}

func (i item) String() string {
	return fmt.Sprintf("[%d,%d]", i.weight, i.value)
}

func (it items) String() string {
	res := "{"
	for _, v := range it {
		tmp := "[" + strconv.Itoa(v.weight) + "," + strconv.Itoa(v.value) + "] "
		res += tmp
	}

	return res[:len(res)-1] + "}"
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// 贪心算法只能求出可行解，但是并不一定可以求出最优解。根据不同的贪心方案，所求解出的可行解还不一定相同。使用动态规划可以解决这个问题。
// 按照单位重量的价值量从大到小进行排序，然后优先选择性价比最高的item装入背包。
// Time complexity: O(n)
// Space complexity: O(n)
func KnapsackGreedy(c int, w []int, v []int) (maxValue int, chose []item) {
	var vw items
	var weight int
	for i := 0; i < len(w); i++ {
		vw = append(vw, item{
			weight:    w[i],
			value:     v[i],
			valueUnit: float32(v[i]) / float32(w[i]),
		})
	}

	sort.Sort(vw)
	for _, it := range vw {
		if weight+it.weight <= c {
			chose = append(chose, it)
			weight += it.weight
			maxValue += it.value
		}
	}

	return
}

// DP 算法可以完美解决背包问题，并且可以求解出最优解。
// 定义dp[i][j]表示背包容量为j时，前i个物品能达到的最大价值。则base case是，物品个数为0时，无论背包容量多大，最大价值均为0，即为dp[0][j]=0.
//对于每个物品，如果背包容量j小于物品重量w，则不装入，此时dp[i][j] = dp[i-1][j];而当背包容量大于物品重量w时，有装入和不装入两个选择，
//取这两个选择中的最大值作为当前选择的结果。
// time complexity: O(n*cap)
// Space complexity: O(n*cap)
func KnapsackDP(capacity int, weights []int, values []int) (int, items) {
	var its items
	length := len(weights)
	dp := make([][]int, length+1)
	// dp[i][0] 和 dp[0][i]均为0
	for i := range dp {
		dp[i] = make([]int, capacity+1)
	}

	for i := 1; i <= length; i++ {
		for j := 1; j <= capacity; j++ {
			// 不能装入
			dp[i][j] = dp[i-1][j]
			// 如果当前背包容量装得下w，并且装入后的value大于不装入的value的话，更新dp
			if j >= weights[i-1] && (dp[i-1][j-weights[i-1]]+values[i-1] > dp[i][j]) {
				dp[i][j] = dp[i-1][j-weights[i-1]] + values[i-1]
			}
		}
	}

	return dp[length][capacity], its
}
