//https://www.hackerrank.com/challenges/weighted-uniform-string/problem

func weightedUniformStrings(s string, queries []int32) []string {
	var v []int32
	var str[]string

	for _,val := range s {
		weight := int32(val - 'a') + 1
		v = append(v, weight)
	}
	for i:=0; i<len(v)-1; i++{
		if s[i] == s[i+1] {
			v[i+1] += v[i]
		}

	}
	for _, val := range queries {
		flag := 0
		for _, val2 := range v {
			if val == val2 {
				flag = 1
				str = append(str, "Yes")
				break
			}
		}
		if flag == 0 {
			str = append(str, "No")
		}
	}
	return str
}