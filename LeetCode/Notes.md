# Notes

字符串匹配类型的题目一般可以使用动态规划和状态机来解决。


## 单调栈

单调栈是在栈先进后出的基础上额外再添加一个特性：从栈顶到栈底的元素是严格递增（或递减的）。

具体过程如下：
- 对于单调递增栈，若当前进栈元素为`e`，从栈顶开始遍历元素，把小于等于`e`的元素弹出栈，直到遇到一个大于`e`的元素或者栈为空为止，然后再把`e`压入栈中。
- 对于单调递减栈，则每次弹出的是大于等于`e`的元素。

**单调递增栈可以找到左起第一个比当前数小的元素，单调递增栈可以找到左起第一个比当前元素大的元素。**

线性的时间复杂度是其最大的优势，每个数字只进栈并处理一次，而解决问题的核心就在处理这块，当前数字如果破坏了单调性，就会触发处理栈顶元素的操作，而触发数字有时候是解决问题的一部分.

很多的单调栈的题目并不是存储实际的数字，而是存储表示该数字的坐标。


## DFS
- 113
```go
func pathSum(root *TreeNode, sum int) [][]int {
	var res [][]int
	if root == nil {
		return res
	}

	helper(root,sum,nil,&res)

	return res
}

func helper(root *TreeNode, sum int, cur []int, res *[][]int) {
	cur = append(cur,root.Val)
	sum -= root.Val
	if root.Left == nil && root.Right == nil && sum == 0 {
		t := make([]int,len(cur))
		copy(t,cur)
		*res = append(*res, t)
		return
	}

	if root.Left != nil {
		helper(root.Left,sum, cur,res)
	}
	if root.Right != nil {
		helper(root.Right, sum, cur,res)
	}
}
```
## BFS
BFS 调用的时候需要注意，由于我们会删除队列中的元素，所以队列的长度经常会在变化，故而不能直接在for循环的条件判断中直接使用len(nums)，而应该提前将其计算好，作为一个常数来使用（代表当前层的节点数量）。


## BS

BS中如何决定是使用left < right 还是left  <= right 呢？ 当你想要查找某个特定target的时候，使用left <= end，一般你会在循环中不断的对mid和target进行检查；
当你不是想找到特定元素的时候，使用left < right，比如说当你想找到最小元素或者边界值的时候，在这种情况下，每一次循环都会reduce space，由于循环在left == right 时候
终止，所以最后剩下的元素就是你想查找的边界。一般循环结束后你还需要对left进行处理，因为left == right时循环就结束了，并没有对left进行处理。

## Go
- string('a')会直接得到a，而strconv.Itoa('a')只会得到98.
- int和int32不是同一个类型。int的实际大小与平台相关，比如在32位上就是32位，在64位上就是64位。官方只保证他和uint有同样的size。也就说说，int可以用来存储比uint32或者int32还要大的数字。

## Palindrome
需要注意字符串的长度是奇数还是偶数，奇数偶数对应着不同的判定方法。
题目：
- 5
1. 中心扩散法
```go
func getPalindrome(s string, left,right int) string {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
    }
    
    // 注意，是left+1
    return s[left+1: end]
}

func longestPalindrome(s string) s string {
	var res string
	for i := 0 ; i < len(s);i++ {
		// handle with odd length
		odd := getPalindrome(s,i,i)
		// handle with even length 
		even := getPalindrome(s,i,i+1)
		
		res = max(res,max(odd,even))
		
    }
    
    return res
}

```

2. 动态规划法
首先，如果一个字符串长度小于2，那么这个字符串肯定是回文串。定义二维数组dp[i][j]，表示从s[i]到s[j]是否为回文串。根据之前的推论，从而引出动态规划的base case， dp[0]=dp[1]=true. 接下来处理状态转移。 
有dp[i][j]为最长回文串的前提是dp[i+1][j-1]是回文串且s[i]==s[j]，因此，在 s[i] == s[j] 成立和 j - i < 3 的前提下，直接可以下结论，dp[i][j] = true，否则才执行状态转移。初始化时，dp[i][i]必然为true。
```go
func longestPalindrome(s string) string {
	var res stirng
	start,maxLen := 0,1  // 非空字符必然会有长度至少为1的子串
	dp := make([][]bool,len(s))
	for i := range dp {
		dp[i] = make([]bool,len(s))
	}
	
    for j := 0; j < n; j++ {
    	// base case
    	dp[j][j] = true
        for i := 0; i < j; i++ {
            //if s[i] != s[j] {
            //    dp[i][j] = false
            //} else {
            //if j-i < 3 {
            //    dp[i][j] = true
            //} else {
            //    dp[i][j] = dp[i+1][j-1]
            //    }
            //}
        	
        	// 上述更精简的写法. 状态转移的核心
        	dp[i][j] = s[i] == s[j] && (j -i < 2 || dp[i+1][j-1])
            if dp[i][j] && j-i+1 > maxLen {
                 start = i
                 maxLen = j - i + 1
                }
            }
        }
    return s[start : start+maxLen]
}
```
   
## Divide
1. MergerKSortedLists

```go
//-----------------------------------------------
// 方法2：使用分治法
// time complexity: O(Nlgk), N 是所有的元素数量，k是数组长度
// space complexity: O(1)
func mergeKListsDivide(lists []*ListNode) *ListNode {
    if lists == nil {
        return nil
    }

    n := len(lists)
    for n > 1{
        // 加1是为了处理奇数个数的情况，加一保证仍然能从尾部开始向前进行合并
        // 如果个数是偶数，那么加1无影响
        k := (n+1)/2
        for i := 0; i< n/2;i++ {
            lists[i] = mergeTwoLists(lists[i],lists[i+k])
        }
        n = k
    }

    return lists[0]
}
```

## LinkedList
1. reverse linkedList
```go
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

```

## Fast and Slow Pointer
```go
	slow,fast := head,head
	for fast.Next != nil && fast.Next.Next != nil{
		slow = slow.Next
		fast = fast.Next.Next
	}
```

##Graph
1. 连通性问题
 - DFS
 - BFS
 - Union Find

## Puzzles
1. Nil Pointer Reference
```go
 type ListNode struct {
     Val int
     Next *ListNode
 }
 
 func main()  {
 	var t1 *ListNode
 	t2 := new(ListNode)
 	t3 := &ListNode{}
 	
 	// why t1.Next will lead to null pointer reference
 	fmt.Println(t1.Next)
}
```
The default pointer value in go is nil. You can pass a nil pointer to a method which takes pointer as its receiver(you can also pass a value, rather than a pointer), but dereference a nil pointer would lead to panic.
If x is nil, evaluate x will cause a run-time panic.

In Go we can create variables that contains the "value of" itself or an address to that value. When the "value of" that variable is an address, the varibale is considered as a pointer. Hence if a pointer is nil, it points to nothing, so evaluate a nil point will lead to panic.

2. Rune and String
可以直接将rune数组转换为string，如`s = string(runeArr)`。