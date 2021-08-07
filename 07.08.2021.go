//https://www.hackerrank.com/challenges/chocolate-feast/problem

func chocolateFeast(n int32, c int32, m int32) int32 {
	var sum int32 = 0
	sum = n / c
	wrap := sum
	for i:=wrap; wrap>=m; i-=m {
		sum++
		wrap = wrap - m + 1
	}
	return sum
}