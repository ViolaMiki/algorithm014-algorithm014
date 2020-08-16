package Week_01

func maxSlidingWindow(nums []int, k int) []int {
	deque := Constructor(3)
	res := make([]int, len(nums) - k + 1)
	for i := 0; i < k; i++ {
		for !deque.IsEmpty() && nums[deque.GetRear()] < nums[i] {
			deque.DeleteLast()
		}
		deque.InsertLast(i)
	}
	res[0] = nums[deque.GetFront()]
	for j := k; j < len(nums); j++ {
		if !deque.IsEmpty() && deque.GetFront() < j - k + 1 {
			deque.DeleteFront()
		}

		for !deque.IsEmpty() && nums[deque.GetRear()] < nums[j] {
			deque.DeleteLast()
		}
		deque.InsertLast(j)
		res[j - k + 1] = nums[deque.GetFront()]
	}
	return res
}
