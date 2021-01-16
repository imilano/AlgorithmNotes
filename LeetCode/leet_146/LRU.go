package leet_146

import "time"

// 一个Cell
type Cell struct {
	Key       int
	Value     int
	RecentUse int64
}

// 一个LRU cache队列，通过map中的key快速映射到Data中相应节点的index
type LRUCache struct {
	Cap  int
	Data []*Cell
	Map  map[int]int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		Cap:  capacity,
		Data: make([]*Cell, 0),
		Map:  make(map[int]int),
	}
}

//func (this *LRUCache) Get(key int) int {
//	value := -1
//
//	if v,ok := this.Map[key]; ok && v != -1 {
//		value = this.Data[v-1].Value
//		this.Data[v-1].RecentUse = time.Now().UnixNano()
//	}
//
//	return value
//}

func (this *LRUCache) Get(key int) int {
	value := -1

	if v, ok := this.Map[key]; ok && v != -1 { // Check -1 to avoid Data[-1]
		value = this.Data[v].Value
		this.Data[v].RecentUse = time.Now().UnixNano()
	}

	return value
}

func (this *LRUCache) Put(key int, value int) {
	// If already exist, modify value and time
	if index, ok := this.Map[key]; ok && index != -1 { // Check -1 to avoid Data[-1]
		this.Data[index].Value = value
		this.Data[index].RecentUse = time.Now().UnixNano()
	} else {
		// If the value  doesn't exist, then insert new node or replace with new one
		cell := &Cell{
			Key:       key,
			Value:     value,
			RecentUse: time.Now().UnixNano(),
		}

		// If  reach out its capacity, then replace the oldest one
		if this.Cap == len(this.Data) {
			min := time.Now().UnixNano()
			var idx int

			for i, v := range this.Data {
				if v.RecentUse < min {
					min = v.RecentUse
					idx = i
				}
			}

			this.Map[this.Data[idx].Key] = -1
			this.Data[idx] = cell
			this.Map[key] = idx
		} else {
			// It is not full, so we insert a node
			this.Data = append(this.Data, cell)
			this.Map[key] = len(this.Data) - 1
		}
	}

	//if v,ok := this.Map[key]; ok && v != -1 {
	//	index := this.Map[key] -1
	//	this.Data[index].Value = value
	//	this.Data[index].RecentUse = time.Now().UnixNano()
	//} else {
	//	// If not exist,ready to insert or replace a node
	//	cell := &Cell{
	//		Key: key,
	//		Value: value,
	//		RecentUse: time.Now().UnixNano(),
	//	}
	//	// If reach its capacity, then remove the oldest, which means to replace the least used one with new key,value and time
	//	if this.Cap == len(this.Data) {
	//		min := time.Now().UnixNano()
	//		var index int
	//
	//		// get the least-used-one index
	//		for k,v := range this.Data {
	//			if v.RecentUse < min {
	//				min = v.RecentUse
	//				index = k
	//			}
	//		}
	//
	//		//this.Data[index].RecentUse = time.Now().UnixNano()
	//		//this.Data[index].Value = value
	//		//this.Data[index].Key = key
	//		//this.Map[key] = index+1
	//		this.Map[this.Data[index].Key] = -1  // Set this to -1, means we remove this node ,so when get(key), we can return -1
	//		this.Data[index] = cell 		// So we replace the node with new one
	//		this.Map[key] = index + 1
	//	}  else {
	//		// If not reach its capacity, then just insert
	//		this.Data = append(this.Data,cell)
	//		this.Map[key] = len(this.Data)
	//	}
	//}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
