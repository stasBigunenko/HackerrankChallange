//https://www.hackerrank.com/challenges/the-birthday-bar/problem

func birthday(s []int32, d int32, m int32) int32 {
	var count int32 = 0
	var sub []int32
	var chock int32
	for i :=0; i <len(s); i++ {
		chock = 0
		if s[i] == d && m == 1 {
			count++
		}else {
			if s[i] < d && i+int(m) <= len(s) {
				sub = s[i:(int32(i)+m)]
				for j := range sub {
					chock+=sub[j]
				}
				if chock == d {
					count++
				}
			}
		}
	}
	return count
}
