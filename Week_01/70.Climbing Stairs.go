package Week_01

func climbStairs(n int) int {
	if n < 1 {
		return 0
	}
	p, q, r := 0, 0, 1
	for i := 0; i < n; i++ {
		p = q
		q = r
		r = p + q
	}
	return r
}
