//https://www.hackerrank.com/challenges/drawing-book/problem

func pageCount(n int32, p int32) int32 {
	var min int32 = 0

	min = n/2 - p/2
	if min > p/2 {
		min = p/2
	}
	return min
}
