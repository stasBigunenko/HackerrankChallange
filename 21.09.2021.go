//https://www.hackerrank.com/challenges/missing-numbers/problem

//additional import
//"sort"

func missingNumbers(arr []int32, brr []int32) []int32 {
	var r []int32
	var res []int32
	for i, val := range brr {
		r = append(r, brr[i])
		for j, val2 := range arr {
			if val == val2 {
				r[i] += 1
				arr[j] = -1
				break
			}
		}
	}
	for i, val := range brr {
		tmp := r[i] - val
		if tmp <0 {
			tmp *= -1
		}
		if tmp == 0 {
			res = append(res, brr[i])
		}
	}

	sort.SliceStable(res, func(i, j int) bool {
		return res[i] < res[j]
	})
	for i:=1; i<len(res); i++ {
		if res[i-1] == res[i] {
			res = append(res[:i], res[i+1:]...)
		}
	}
	return res
}