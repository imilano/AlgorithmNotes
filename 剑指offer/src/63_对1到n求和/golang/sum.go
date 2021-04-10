package golang


// 求1+2+3+...+n，要求不能使用乘除法、for、while、if、else、switch、case等关键字及条件判断语句（A?B:C）。

func Sum_Solution( n int ) int {
	// write code here
	var f func(ret *int, n int)bool
	f = func (ret *int, n int)  bool {
		*ret += n
		return n != 0 && f(ret,n-1)
	}

	var ans int
	f(&ans,n)
	return ans
}