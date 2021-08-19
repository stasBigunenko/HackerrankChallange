//https://www.hackerrank.com/challenges/mars-exploration/problem

func marsExploration(s string) int32 {
	var count int32 = 0

	for i := range s {
		if i % 3 == 1 && string(s[i]) != "O" {
			count++
		}
		if (i % 3 == 0 || i % 3 == 2) && string(s[i]) != "S" {
			count++
		}
	}
	return count
}