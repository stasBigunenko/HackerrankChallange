//https://www.hackerrank.com/challenges/marcs-cakewalk/problem

//additional import
//import "sort"
//import "math"

func marcsCakewalk(calorie []int32) int64 {
	var res float64 = 0
	sort.SliceStable(calorie, func(i, j int) bool {
		return calorie[i] < calorie[j]
	})
	j := len(calorie)-1
	for i :=0; i < len(calorie); i++ {
		res2 := (math.Pow(float64(2), float64(i)))*float64(calorie[j])
		res = res + res2
		j--
	}
	return int64(res)
}