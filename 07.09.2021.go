//https://www.hackerrank.com/challenges/two-strings/problem

func twoStrings(s1 string, s2 string) string {
	str := [26]int{0}
	res := "NO"

	for i := 0; i < len(s1); i++ {
		str[s1[i]-'a']++
		}

	for i := 0; i < len(s2); i++ {
		if str[s2[i]-'a'] > 0 {
			res = "YES"
			break
		}
	}
	return res
}
