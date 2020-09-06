package Week_04

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	len, col := len(grid), len(grid[0])
	res := 0

	for x := 0; x < len; x++ {
		for y := 0; y < col; y++ {
			if grid[x][y] == '1' {
				res++
				dsf(x, y, grid)
			}
		}
	}
	return res
}

func dsf(x, y int, grid [][]byte) {
	if (x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) || grid[x][y] == '0') {
		return
	}

	grid[x][y] = '0'
	dsf(x, y - 1, grid)
	dsf(x, y + 1, grid)
	dsf(x - 1, y, grid)
	dsf(x + 1, y, grid)
}
