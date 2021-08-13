//https://www.hackerrank.com/challenges/repeated-string/problem

func repeatedString(s string, n int64) int64 {
	count := int64(0)

	for _, val := range s {
		if val == 'a' {
			count++
		}
	}

	p := n / int64(len(s))
	count = count * p

	k := n % int64(len(s))
	if k != 0 {
		for i:=0;i<int(k); i++ {
			if s[i] == 'a' {
				count++
			}
		}
	}
	return count
}
