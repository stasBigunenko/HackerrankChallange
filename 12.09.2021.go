//https://www.hackerrank.com/challenges/string-construction/problem

func stringConstruction(s string) int32 {
	var cost int32 = 1

	for i :=1; i < len(s); i++ {
		for j:=i-1; j>=0; j-- {
			if s[j] == s[i] {
				cost--
				break
			}
		}
		cost++
	}
	return cost
}
