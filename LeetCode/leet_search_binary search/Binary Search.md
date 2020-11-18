二分查找的边界比较难以控制，经常因为题目的变动而出现不同的变体。所谓「十个二分九个错」，即便是像`Jon Bentley`这样的人也不能免俗。最近刚好在LeetCode上刷到了一些二分查找的题目，今天顺手做个总结。

## 寻找一个数 

### 溢出的问题

这个问题其实是一个简单的问题，说一遍大家也就记住了。常见的二分查找计算`mid`的方式是`mid = (low+high)/2`，如果`low`和`high`的值不大，那这么写也可；但是当`low`和`high`的值很大时，两个有符号数相加，就会出现溢出的问题。更好的写法应该是下面这样：
```go
	mid = low + (high - low)/2
```
### 循环判断用`<`还是用`<=`

这个与判断的区间相关，并不可一概而论。
```go
func search(arrs []int,target int) int {
	low,hight := 0,len(arrs)-1  // 左闭右闭
	for low <= high {  // 注意
		mid := low + (high - low)/ 2
		if arrs[mid] == target {
				return mid
		} else if arrs[mid] <target {
			low = mid + 1 // 注意
		} else {
			high = mid - 1  // 注意
			}
	} 
}
```

上面的代码中，我们定义high为len(arrs)-1，意味着我们的搜索区间是左闭右闭区间，如果我们使用low <high，那么最后的终止条件是low=high，也就是[low,high]，可见还有一个元素没有进行搜索，所以我们需要使用low <= high ，这样循环的终止条件是low = hight+1，此时区间[high+1,high]中没有元素进行搜索了，循环结束。

当然我们也可以使用low < high，只不过这样的话，我们就需要在最后判断一下arrs[low] 是否等于target，等于多了一个Corner case 的判断。

如果我们使用这样的代码：

```go
func search(nums []int, target int) int {
  left,right := 0,len(nums) // 注意，左闭右开
  for left < right {
    mid := left + (right -leftd)/2
    if nums[mid] == target {
      return mid
    } else if nums[mid] < target {
      	left = mid +1  // 注意
    } else if nums[mid] > target {
      	right = mid  // 注意
    }
  }
}
```

很明显可以看到这是左闭右开的区间，所以我们使用[left,right)，所以当更改left和right的值时，我们也采用闭区间的赋值方式。循环的终止条件是left == right，也就是[left,left)，可以看到区间中已经没有元素可以进行搜索了，所说义这样子的方式也是正确的。

### 到底是right  = mid  还是right = mid -1（同理对于left）
上文明确了**搜索区间**的概念，那么，如果arrs[mid]  != target，我们应该怎么缩减区间内。很明显，只需要缩小区间到[low,mid-1]或者[mid+1,high]即可，因为我们已经搜索过这个位置了，那么我们只需要将其从搜索区间去除就好了。

## 寻找左侧边界的二分搜索

在刷LeetCode中，经常会出现这样的情况，要求你查找重复元素的最左边界或者重复元素的最右边界，此时就需要这样的可以对边界值进行bias的查找。
```go
	func search(nums []int, target int) int {
		left,right := 0,len(nums) // 1
		for left < right {  // 2
			mid := left + (right - left)/2
			if nums[mid] == target {
					right = mid		// 3
			} else if nums[mid] < target {
					left = mid + 1 // 4
			} else if nums[mid] > target {
					right = mid // 5
			}
		}
		
    // target 比所有数都大
    if (left == nums.length) return -1; // 循环结束条件left == right, 区间为[left,right)
    // 类似之前算法的处理方式
    return nums[left] == target ? left : -1;
	}
```

同理可以使用区间的概念来解释1、2、4、5。对于3，我们可以这么考虑： 因为数组中存在着很多的重复元素，而我们的目的是找到最左边界，那么当我们找到一个其中的一个元素时，我们想要的肯定是把区间向左进行缩减，从而在[left,right)之内继续进行搜索。

那么为什么程序没有返回-1，而是返回left呢？仔细分析程序可以得到，left的取值区间位于[0,len(nums)]之间，当数组中没有target时，left就会一直往右走，直到等于len(nums)，此时我们只需要检查最后的left的值就可以做到边界检查。
```go
	if left == len(nums) {
		return -1
	}
	
	return nums[left] == target? left:-1
```

## 寻找右侧边界的二分查找

```go
	func search(nums []int,target int) int {
		left,right := 0,len(nums)
		for left < right {
			mid := left + (right -left)/2
			if nums[mid] == target {
				left = mid +1
			} else if nums[mid] > target {
				rigth = mid
			} else if nums[mid] < target {
				left = mid + 1
			}
		}
		
		if left == 0 {
		 return -1
		}
		
    if nums[left-1] == target {  // left == right 循环结束，区间为[left,right)
			return left -1 
		} else {
			return -1
		}
	}

```

终止条件的判断，主要根据循环结束时的终止条件。

### todo

- +1 bias
- 何时取等于，何时加一何时减一
- 最后的情况中，low=mid=high和low=mid=high-1。
