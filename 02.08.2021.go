//https://www.hackerrank.com/challenges/camelcase/problem

func camelcase(s string) int32 {
	var count int32 = 1

	for i := range s{
		if s[i] >= 'A' && s[i] <= 'Z' {
			count++
		}
	}
	return count
}