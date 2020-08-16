package Week_01

func removeDuplicates(nums []int) int {
	p, q := 0, 1
	for q < len(nums) {
		if nums[p] != nums[q] {
			p ++
			if p != q {
				nums[p] = nums[q]
			}
		}
		q++
	}
	return p + 1
}