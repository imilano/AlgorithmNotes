package main

import (
	"fmt"
)

type Node struct {
	Age int
}
//  Definition for a binary tree node.
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func search(nums []int, target int) int {
	left,right := 0,len(nums)
	for left < right {


		mid := left + (right -left)/2
		fmt.Printf("[%d,%d,%d)\n",left,mid,right)
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid +1
		} else if nums[mid] > target {
			right = mid
		}

	}

	return 0
}

func main() {
	//arr := []int{7,6,4,3,1}
	//fmt.Println(leet_121.MaxProfit(arr))
	//a := make([]*Node,0)
	//fmt.Printf("%v\n",append(a,&Node{12}))
	//
	//fmt.Printf("%b\n",math.MinInt8 << 4)
	//fmt.Printf("%v\n",math.MinInt8 << 4)
	//fmt.Printf("%v\n",24.0/2.0)

	//stack := utils.NewStack()
	//left := &TreeNode{
	//	Val:   4,
	//	Ano: 5,
	//	Left:  nil,
	//	Right: nil,
	//}
	//
	//right := &TreeNode{
	//	Val:   5,
	//	Ano: 4,
	//	Left:  nil,
	//	Right: nil,
	//}
	//root := &TreeNode{
	//	Val:   77,
	//	Ano: 88,
	//	Left:  left,
	//	Right: right,
	//}
	//stack.Push(left)
	//stack.Push(right)
	//stack.Push(root)
	//node := stack.Pop().(*TreeNode)
	//fmt.Printf("%+v",node)

	// fmt.Println(nil == nil)
	//node1 := &TreeNode{
	//	Val:   3,
	//	Left:  nil,
	//	Right: nil,
	//}
	//
	//node2 := &TreeNode{
	//	Val:   3,
	//	Left:  nil,
	//	Right: nil,
	//}
	//
	//root := &TreeNode{
	//	Val:   4,
	//	Left:  node1,
	//	Right: node2,
	//}
	//
	//var que []*TreeNode
	//q1 := node1
	//que = append(que,node1,node2,root)
	//
	//t := que[0]
	//if t == q1 {
	//	fmt.Println("Equal")
	//}

	//var s []int
	//s = append(s,2,3,4)
	//ss := s[2:]
	//fmt.Printf("%+v\n",ss)
	//fmt.Printf("%+v",ss[0:2])

	//arr := []int {
	//	1,2,34,45,6,67,3,2,5,2,5,12,5,2,6,2,6,2,653,43,2,4,234,34,3,2,3,4,5,6,6,35,3,5,35,34,534,5,34,53,445,3,5,345,34,5,34,5,4443,4444,
	//}
	//fmt.
	//	Println(len(arr))
	//sort.Ints(arr)
	//fmt.Println(search(arr,4444))
	//nums := []int{2,3,6,7}
	//target := 7
	////res := combinationSum
	//res := []int {1,2,3,4}
	//b := res
	//
	//fmt.Printf("%p\n",res)
	//sliceCopy(res)
	//fmt.Printf("%p\n",b)
	//fmt.Printf("%p",append(b,4))

	//s := "()("
	//for _,v := range s {
	//	if v == rune('(') {
	//		fmt.Print(v)
	//	}
	//}

	////fmt.Println(gcd(28,28))
	//fmt.Println(string('3') + string('2'))
	//fmt.Println(strconv.Itoa('3') + strconv.Itoa('2'))
	m := make(map[string]bool)
	v := m["hello"]
	fmt.Println(v)
}

func gcd(n1,n2 int) int {
	var gcd int
	k  := 2
	for k <= n1/2 && k <= n2/2 {
		if n1 % k ==0 && n2 % k == 0 {
			gcd = k
		}
		k++
	}
	return gcd
}

func sliceCopy(a []int) {
	fmt.Printf("%p\n",a)
}
