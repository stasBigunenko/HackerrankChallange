//https://www.hackerrank.com/challenges/flatland-space-stations/problem

func flatlandSpaceStations(n int32, c []int32) int32 {
	max := int32(0)

	for i:=int32(0); i<n; i++{
		min := int32(100000)
		for j:=0; j<len(c); j++	{
			res := i - c[j]
			if res < 0 {
				res *= -1
			}
			if min > res {
				min = res
			}
		}
		if max < min {
			max = min
		}
	}
	return max
}