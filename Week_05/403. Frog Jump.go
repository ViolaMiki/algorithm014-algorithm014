package Week_05

func canCross(stones []int) bool {
	dpMap := make([][]bool, len(stones))
	for k, _ := range dpMap {
		dpMap[k] = make([]bool, len(stones) + 1)
	}

	for i, stone := range stones {
		if i == 0 {
			dpMap[0][0] = true
			continue
		}
		for j := 0; j < i; j++ {
			k := stone - stones[j]

			if k <= j + 1 {
				dpMap[i][k] = dpMap[j][k - 1] || dpMap[j][k] || dpMap[j][k + 1]
				if i == len(stones) - 1 && dpMap[i][k] == true {
					return true
				}
			}
		}
	}

	return false
}