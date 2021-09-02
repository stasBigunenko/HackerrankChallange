//https://www.hackerrank.com/challenges/hackerrank-in-a-string/problem

func hackerrankInString(s string) string {
	res :="NO"
	var count int = 0
	var tmp int = -1
	n := "hackerrank"

	for i := range n {
		if tmp+1 >= len(s) {
			break
		}
		for j:=tmp+1; j<len(s); j++ {
			if n[i] == s[j] {
				count++
				tmp = j
				break
			}
		}
		if count ==10 {
			res = "YES"
		}
	}
	return res
}
