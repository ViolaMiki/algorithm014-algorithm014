package Week_01

type MyCircularDeque struct {
	capacity int
	head int
	tail int
	arr  []int
}


/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		arr: make([]int, k + 1),
		head: 0,
		tail: 0,
		capacity: k + 1,
	}
}


/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}

	this.head = (this.head + this.capacity - 1) % this.capacity
	this.arr[this.head] = value
	return true
}


/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}

	this.arr[this.tail] = value
	this.tail = (this.tail + 1) % this.capacity
	return true
}

/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}

	this.arr[this.head] = 0
	this.head = (this.head + 1) % this.capacity
	return true
}


/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}

	this.tail = (this.tail + this.capacity - 1) % this.capacity
	this.arr[this.tail] = 0
	return true
}


/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}
	return this.arr[this.head]
}


/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.arr[(this.tail - 1 + this.capacity) % this.capacity]
}


/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
	if this.head == this.tail {
		return true
	}
	return false
}


/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
	if (this.tail + 1) % this.capacity == this.head {
		return true
	}
	return false
}


/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
