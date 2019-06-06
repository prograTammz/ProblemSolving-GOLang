package main

func numJewelsInStones(J string, S string) int {
	var count int
	for _, chari := range J {
		for _, charj := range S {
			if chari == charj {
				count++
			}
		}
	}
	return count
}
