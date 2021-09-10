//https://www.hackerrank.com/challenges/luck-balance/problem

//Additional import "sort"

func luckBalance(k int32, contests [][]int32) int32 {
	var res int32 = 0
	var count int32 = 0
	var ss []int32

	for i := range contests {
		res += contests[i][0]
		if contests[i][1] == 1 {
			count++
			ss = append(ss, contests[i][0])
		}
	}
	sort.Slice(ss, func (i, j int) bool {
		return ss[i] < ss[j]
	})
	for i := int32(0); i < (count-k); i++ {
		res -= ss[i]*2
	}
	return res
}