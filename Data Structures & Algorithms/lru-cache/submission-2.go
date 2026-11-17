type ListNode struct {
	Key int
	Val int
	Next *ListNode
	Prev *ListNode
}

type LRUCache struct {
   capacity int
   data map[int]*ListNode
   lastUsed *ListNode
   recentlyUsed *ListNode
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		data: make(map[int]*ListNode),
		lastUsed: nil,
		recentlyUsed: nil,
	}
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.data[key]
	if !ok {
		return -1
	}

	this.shift(node)

	return node.Val
}

func (this *LRUCache) Put(key int, value int) {

	var node *ListNode
	if v, ok := this.data[key]; ok {
		node = v
		node.Val = value
	} else {
		node  = &ListNode{Key: key, Val: value}
	}
	
	if len(this.data) == 0 {
		this.recentlyUsed = node
		this.lastUsed = node
	} else {
		this.shift(node)
		if _, ok := this.data[key]; !ok {

		this.evictLRU()
		}
	}

	this.data[key] = node
}

func (this *LRUCache) evictLRU() {
	if len(this.data) < this.capacity {
		return
	}	

	delete(this.data, this.lastUsed.Key)
	this.lastUsed = this.lastUsed.Next
	this.lastUsed.Prev = nil
}

func (this *LRUCache) shift(node *ListNode) {
	if this.recentlyUsed == node {
		return
	}

	if this.lastUsed == node {
		this.lastUsed = this.lastUsed.Next
		this.lastUsed.Prev = nil
	}


	prev, next := node.Prev, node.Next

	if prev != nil {
		prev.Next = next
	}
	if next != nil {
		next.Prev = prev
	}

	recentlyUsed := this.recentlyUsed

	node.Prev = this.recentlyUsed
	node.Next = nil
	
	recentlyUsed.Next = node
	this.recentlyUsed = node

}