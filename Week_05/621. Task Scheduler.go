package Week_05

func leastInterval(tasks []byte, n int) int {
	charArr := [26]int{}
	max := 0
	lastRow := 0
	for _, letter := range tasks {
		charArr[letter - 'A']++
	}

	for _, count := range charArr {
		if max < count {
			max = count
			lastRow = 1
		} else if max == count {
			lastRow++
		}
	}

	return Max((max - 1) * (n + 1) + lastRow, len(tasks))
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
