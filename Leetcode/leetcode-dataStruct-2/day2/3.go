package day2

type MyHashMap struct {
	Value [][]int
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	return MyHashMap{}
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
	for _, v := range this.Value {
		if v[0] == key {
			v[1] = value
			return
		}
	}
	this.Value = append(this.Value, []int{key, value})
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	for _, v := range this.Value {
		if v[0] == key {
			return v[1]
		}
	}
	return -1
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	for index, value := range this.Value {
		if value[0] == key {
			this.Value = append(this.Value[:index], this.Value[index+1:]...)
		}
	}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
