//https://www.hackerrank.com/challenges/making-anagrams/problem

func makingAnagrams(s1 string, s2 string) int32 {
	var res [26]int

	for i:=0; i<len(s1); i++ {
		res[s1[i]-'a'] += 1
	}
	for i:=0; i<len(s2); i++ {
		res[s2[i]-'a'] -= 1
	}

	sum := int32(0)
	for i:=0; i<26; i++ {
		if res[i] < 0 {
			res[i] *= -1
		}
		sum += int32(res[i])
	}

	return sum
}