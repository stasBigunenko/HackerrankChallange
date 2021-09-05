//https://www.hackerrank.com/challenges/gem-stones/problem

func gemstones(arr []string) int32 {
	r := [26]int32{0}
	res := int32(0)

	for i:=0; i<len(arr); i++ {
		for _, val := range arr[i] {
			if r[val - 'a'] == int32(i)  {
				r[val - 'a']++
				if r[val - 'a'] == int32(len(arr)) {
					res++
				}
			}
		}
	}
	return res
}