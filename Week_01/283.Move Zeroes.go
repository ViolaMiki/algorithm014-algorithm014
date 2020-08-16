package Week_01

func moveZeroes(nums []int)  {
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			j := i + 1
			for j < len(nums) && nums[j] == 0 {
				j++
			}

			if j == len(nums) {
				return
			}
			nums[i] = nums[j]
			nums[j] = 0
		}
	}
}
