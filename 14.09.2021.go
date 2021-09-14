//https://www.hackerrank.com/challenges/palindrome-index/problem

func palindromeIndex(s string) int32 {
	length := len(s) - 1
	res := int32(-1)

	for i:=0; i<length; i++ {
		if s[i] == s[length] {
			length--
		}else{
			if s[i] == s[length-1] && s[i+1] == s[length-2] {
				res = int32(length)
			}else {
				res = int32(i)
			}
			break
		}
	}
	return res
}
