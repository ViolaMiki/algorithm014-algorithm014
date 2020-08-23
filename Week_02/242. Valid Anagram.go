package Week_02

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	hash := map[byte]int{}
	for _, m := range s {
		hash[byte(m)]++
	}

	for _, m := range t {
		w, ok := hash[byte(m)]
		if !ok {
			return false
		}

		if w <= 0 {
			return false
		}

		hash[byte(m)]--
	}

	for _, m := range hash {
		if m > 0 {
			return false
		}
	}

	return true
}