package Week_07

type LRUCache struct {
	dict map[int]*LinkNode
	cap, count int
	head, tail *LinkNode
}

type LinkNode struct {
	k, v int
	pre, next *LinkNode
}


func Constructor(capacity int) LRUCache {
	cache := LRUCache{
		dict: map[int]*LinkNode{},
		cap: capacity,
		head: &LinkNode{0, 0, nil, nil},
		tail: &LinkNode{0, 0, nil, nil},
	}
	cache.head.next = cache.tail
	cache.tail.pre = cache.head
	return cache
}


func (this *LRUCache) Get(key int) int {
	if v, ok := this.dict[key]; ok {
		this.moveToHead(v)
		return v.v
	} else {
		return -1
	}
}


func (this *LRUCache) Put(key int, value int)  {
	if v, ok := this.dict[key]; ok {
		v.v = value
		this.moveToHead(v)
	} else {
		this.count++
		if (this.count > this.cap) {
			delKey := this.delTail()
			delete(this.dict, delKey)
			this.count--
		}
		v := &LinkNode{key, value, nil, nil}
		this.addNode(v)
		this.dict[key] = v
	}

}

func (this *LRUCache) moveToHead(v *LinkNode) {
	v.pre.next = v.next
	v.next.pre = v.pre

	v.next = this.head.next
	v.pre = this.head

	this.head.next.pre = v
	this.head.next = v
}

func (this *LRUCache) addNode(v *LinkNode) {
	v.next = this.head.next
	v.pre = this.head

	this.head.next.pre = v
	this.head.next = v
}

func (this *LRUCache) delTail() int {
	last := this.tail.pre
	last.pre.next = this.tail
	this.tail.pre = last.pre
	return last.k
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
