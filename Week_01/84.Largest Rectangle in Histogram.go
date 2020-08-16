package Week_01

func largestRectangleArea(heights []int) int {
	stack := make([]int, len(heights) + 1)
	index := 0

	stack[index] = -1
	max := 0;
	i := 0

	for i < len(heights)  {
		if stack[index] == -1 || heights[stack[index]] <= heights[i] {
			index++
			stack[index] = i
			i++
		} else {
			currentI := stack[index]
			index--
			preI := stack[index]
			area := heights[currentI] * (i - preI - 1)
			max = Max(max, area)
		}
	}

	leftIndex := stack[index]
	for index != 0 {
		currentI := stack[index]
		index--
		preI := stack[index]
		area := heights[currentI] * (leftIndex - preI)
		max = Max(max, area)
	}
	return max
}

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

