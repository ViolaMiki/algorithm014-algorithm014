package Week_01

import "fmt"

func getHint(secret string, guess string) string {
	bulls, cows := 0, 0
	bucket := map[byte]int {}
	for k, v := range secret {
		if byte(v) == guess[k] {
			bulls++
		}
		bucket[byte(v)]++
	}

	for _,v := range guess {
		if bucket[byte(v)] > 0 {
			cows++
			bucket[byte(v)]--
		}
	}
	cows = cows - bulls
	return fmt.Sprintf("%dA%dB", bulls, cows)
}
