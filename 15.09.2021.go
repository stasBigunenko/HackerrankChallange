//https://www.hackerrank.com/challenges/maximum-perimeter-triangle/problem

//additional import "sort"

func maximumPerimeterTriangle(sticks []int32) []int32 {
	res := []int32{-1}

	sort.Slice(sticks, func (i, j int) bool {
		return sticks[i] < sticks[j]
	})

	for i := len(sticks)-1; i >= 0; i-- {
		if i-2 >= 0 && sticks[i-2]+sticks[i-1] > sticks[i]{
		res = sticks[i-2:i+1]
		break
		}
	}

	return res
}