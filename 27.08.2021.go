//https://www.hackerrank.com/challenges/pangrams/problem

func pangrams(s string) string {
	count:= [26]int{}
	var res string = "pangram"

	s = strings.ToLower(s)

	for _, val := range s {
		if val >= 'a' && val <= 'z' {
			count[val-'a'] = 1
		}
	}
	for i:=0; i < 26; i++ {
		if count[i] != 1 {
			res = "not pangram"
			break
		}
	}
	return res
}
