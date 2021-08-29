//https://www.hackerrank.com/challenges/alternating-characters/problem

func alternatingCharacters(s string) int32 {
	var count int32 = 0

	for i :=1; i < len(s); i++ {
		if int(s[i]) == int(s[i-1]) {
			count++
		}
	}
	return count
}
