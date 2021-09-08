//https://www.hackerrank.com/challenges/anagram/problem

func anagram(s string) int32 {
	var res int32 = 0
	sub := [26]int32{}

	if len(s) % 2 != 0 {
		res = -1
		return res
	}

	for i:=0; i<len(s)/2; i++ {
		sub[s[i]-'a']++
		sub[s[len(s)/2+i]-'a']--
	}

	for _, val := range sub {
		if val > 0 {
			res += val
		}
	}
	return res
}