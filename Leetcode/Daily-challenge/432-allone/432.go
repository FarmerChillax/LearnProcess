package allone

type AllOne struct {
	Value  map[string]int
	MaxKey string
	MinKey string
}

func Constructor() AllOne {
	allone := AllOne{
		Value: make(map[string]int),
	}
	return allone
}

func (this *AllOne) Inc(key string) {
	if _, exist := this.Value[key]; exist {
		this.Value[key]++
		if this.Value[key] > this.Value[this.MaxKey] {
			this.MaxKey = key
		}
		if this.Value[key] < this.Value[this.MinKey] {
			this.MinKey = key
		}
	} else {
		this.Value[key]++
		if this.MaxKey == "" {
			this.MaxKey = key
		}
		if this.MinKey == "" {
			this.MinKey = key
		}
		if this.Value[key] < this.Value[this.MinKey] {
			this.MinKey = key
		}
	}

}

func (this *AllOne) Dec(key string) {
	if this.Value[key] > 0 {
		this.Value[key]--
	} else {
		this.Value[key] = 0
	}

	var max, min int = 0, 999

	for k, value := range this.Value {
		if value > max {
			max = value
			this.MaxKey = k
		}
		if value < min && value != 0 {
			min = value
			this.MinKey = k
		}
	}

}

func (this *AllOne) GetMaxKey() string {
	return this.MaxKey
}

func (this *AllOne) GetMinKey() string {
	return this.MinKey
}

/**
 * Your AllOne object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Inc(key);
 * obj.Dec(key);
 * param_3 := obj.GetMaxKey();
 * param_4 := obj.GetMinKey();
 */
