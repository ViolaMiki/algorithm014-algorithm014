package Week_01

func rotate(nums []int, k int)  {
	count, tmp := 0, 0
	for start := 0; count < len(nums); start++ {
		cNum := nums[start]
		oldIndex := start
		for true {
			newIndex := (oldIndex + k) % len(nums)
			tmp = cNum
			cNum = nums[newIndex]
			nums[newIndex] = tmp
			oldIndex = newIndex
			count++
			if oldIndex == start {
				break
			}
		}
	}
}