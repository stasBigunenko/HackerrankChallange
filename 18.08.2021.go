//https://www.hackerrank.com/challenges/jumping-on-the-clouds/problem

func jumpingOnClouds(c []int32) int32 {
	var count int32 = -1

	for i:=0; i<len(c); i++{
		count++
		if (i+2 < len(c)) {
			if c[i+2] == int32(0) {
				i += 1
				continue
			} else if c[i+2] == c[i+1] {
				i += 1
				continue
			}
		}
	}
	return count
}