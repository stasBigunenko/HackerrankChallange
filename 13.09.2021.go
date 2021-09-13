//https://www.hackerrank.com/challenges/the-love-letter-mystery/problem

func theLoveLetterMystery(s string) int32 {
	count := 0

	i := 0
	j := len(s)-1
	for i<j  {
		res := int32(s[j]) - int32(s[i])
		if res < 0 {
			res *= (-1)
		}
		count += res
		i++
		j--
	}
	return count
}