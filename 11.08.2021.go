//https://www.hackerrank.com/challenges/service-lane/problem

func serviceLane(width []int32, cases [][]int32) []int32 {
	var min int32 = 0
	var sl []int32
	for i:=0; i< len(cases); i++ {
		min = width[cases[i][0]]
		for j:=cases[i][0]; j<=cases[i][1]; j++ {
			if min > width[j] {
				min = width[j]
			}
		}
		sl = append(sl, min)
	}
	return sl
}