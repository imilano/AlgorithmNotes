package Search


// 递归版本
func binary_search_recur(arr []int, low int, high int, key int) int {
	if low > high {
		return -1
	}

	mid := low + (high-low)/2 // 防止整数溢出
	
	if arr[mid] > key {
		return binary_search_recur(arr, low, mid-1,key)
	} else if arr[mid] < key {
		return binary_search_recur(arr, mid+1, high, key)
	}
	return mid;
}

// 迭代版本
func binary_search_itera(arr []int, low int, high int, key int) int {
	var mid int

	for low <= high {
		mid = low + (high-low)/2
		if arr[mid] == key {
			return mid
		}else if arr[mid] > key {
			high = mid -1
		} else if arr[mid] < key {
			low = mid +1
		}
	}
	return -1
 }
