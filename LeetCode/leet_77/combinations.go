package leet_77

/*
Given two integers n and k, return all possible combinations of k numbers out of 1 ... n.
You may return the answer in any order.
 */

//--------------------------------------
// backtracing
func combine(n int, k int) [][]int {
	var res [][]int
	if k > n {
		return res
	}

	getCombination(nil,1,n,k,&res)
	return  res
}

func getCombination(nums []int, start int, n int, k int, res *[][]int) {
	if k == 0 {
		*res = append(*res,nums)
		return
	}

	for i := start; i<= n-k+1;i++ {
		nums = append(nums,i)
		tmp := make([]int, len(nums))
		copy(tmp,nums)
		getCombination(tmp,i+1,n,k-1,res)
		nums = nums[:len(nums)-1]
	}
}

//-------------------------------------------
// 递归法，利用公式 C(n,k) = C(n-1,k-1) + C(n-1,k)。
// 含义，从n个数里面选k个数，有两种方法：
//		- 从n-1个数里选k-1个数，然后每个结果加上n
// 		- 从n-1个数里选k个数，不包含n
// 将上述二者的结果加起来就是我们想要的结果
func getCombinationRec(n int, k int) [][]int {
	if k == n || k == 0 {
		var row []int
		for i := 1; i<= k; i++ {
			row = append(row, i)
		}

		return [][]int{row}
	}

	// 从n-1个数里面选k-1个数
	res := getCombinationRec(n-1,k-1)
	// 加上n
	for i := range res {
		res[i] = append(res[i],n)
	}

	// 结果再加上不包括n的选择
	res = append(res,combine(n-1,k)...)
	return res
}

//---------------------------------------------
// DFS
func combineDFS(n int, k int) [][]int {
	var res [][]int
	if k > n || n == 0 || k == 0 {
		return res
	}

	dfs(1,n,k,nil,&res)
	return res
}

func dfs(start int, n int, k int, current []int, res *[][]int) {
	if k == 0 {
		tmp  := make([]int, len(current))
		copy(tmp,current)
		*res = append(*res,tmp)
		return
	}


	for i := start; i <= n-k+1; i++ {
		dfs(i+1,n,k-1,append(current,i),res)
	}
}

//
//func getCombination(nums []int, start int,k int,  res *[][]int) {
//	if len(nums) == 1 {
//		*res = append(*res,nums)
//		return
//	}
//
//	if index == len(nums) {
//		*res = append(*res,nums)
//		return
//	}
//
//	for i := index; i<len(nums);i++ {
//		nums[index],nums[i] = nums[i],nums[index]
//		tmp := make([]int, len(nums))
//		copy(tmp,nums)
//		getCombination(tmp,index+1,res)
//	}
//}
//
