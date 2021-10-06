//https://www.hackerrank.com/challenges/acm-icpc-team/problem

func acmTeam(topic []string) []int32 {
	var res []int32
	var str []string
	var str2 []string
	max := 0
	var teams int32
	for i:=0; i<len(topic)-1; i++ {
		str = strings.Split(topic[i], "")
		for k := i+1; k<len(topic); k++ {
			str2 = strings.Split(topic[k], "")
			count := 0
			for j := 0; j < len(str); j++ {
				if str[j] == "1" || str2[j] == "1" {
					count++
				}
			}
			if max < count {
				max = count
				teams = 1
			}else if count == max {
				teams++
			}
		}
	}
	res = []int32{int32(max), teams}
	return res
}