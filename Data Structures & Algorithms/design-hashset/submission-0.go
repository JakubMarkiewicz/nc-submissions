type MyHashSet struct {
	value []int
}

func Constructor() MyHashSet {
	return MyHashSet{}
}

func (this *MyHashSet) Add(key int) {
	if !this.Contains(key) {
		this.value = append(this.value, key)
	}
}

func (this *MyHashSet) Remove(key int) {
	if this.Contains(key) {
		newValue := []int{}	

		for _, value := range this.value {
			if value != key {
				newValue = append(newValue, value)
			}
		}
		this.value = newValue
	}
    
}

func (this *MyHashSet) Contains(key int) bool {
	for _, value := range this.value {
		if value == key {
			return true
		}
	}

	return false
    
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
 