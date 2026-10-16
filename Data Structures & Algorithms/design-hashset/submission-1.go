type MyHashSet struct {
	value []bool
}

func Constructor() MyHashSet {
	return MyHashSet{value: make([]bool, 1000001)}
}

func (this *MyHashSet) Add(key int) {
		this.value[key] = true
}

func (this *MyHashSet) Remove(key int) {
	this.value[key] = false
}

func (this *MyHashSet) Contains(key int) bool {
	return  this.value[key]
    
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
 