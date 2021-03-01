package utils


// 使用golang 的container包构造一个最小堆
type MinHeap []int


// container/heap 接口首先内嵌了一个sort.Interface,所以需要先实现sort接口
func (h MinHeap) Less(i,j int) bool {
	return h[i] < h[j]
}
func (h MinHeap) Len()int {
	return len(h)
}
func (h MinHeap) Swap(i,j int) {
	h[i],h[j] = h[j],h[i]
}

// 如下两个方法是heap的特有方法
func (h *MinHeap) Push(i interface{}) {
	*h = append(*h,i.(int))
}

func (h *MinHeap)Pop() interface{} {
	// TODO: 为什么标准库给的例子还使用了old，而不是直接使用原数组
	//old := *h
	//n := len(*h)
	//x := old[n-1]
	//*h = (*h)[0 : n-1]
	//return x
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[:n-1]
	return x
}


// 那么如何避免代码冗余的情况下构建出最大堆呢? 使用嵌入
type MaxHeap struct {
	MinHeap
}

func (h MaxHeap)Less(i,j int) bool {
	return h.MinHeap[i] > h.MinHeap[j]
}

