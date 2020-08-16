package Week_01

func maxArea(height []int) int {
	i, j, max := 0, len(height) - 1, 0
	l := len(height) - 1
	h := 0
	for i < j {
		left := height[i]
		right := height[j]
		if left < right {
			h = left
			i++
		} else {
			h = right
			j--
		}

		v := h * l
		if v > max {
			max = v
		}
		l--
	}

	return max
}