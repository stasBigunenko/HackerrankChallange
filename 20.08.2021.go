//https://www.hackerrank.com/challenges/permutation-equation/problem

func permutationEquation(p []int32) []int32 {
	var res []int32
	var numb int32

	for i := int32(1); i<=int32(len(p)); i++{
		for j := int32(0); j<int32(len(p)); j++ {
			if i == p[j] {
				numb = j + 1
				break
			}
		}
		for j := int32(0); j<int32(len(p)); j++ {
			if p[j] == numb {
				res = append(res, j+1)
				break
			}
		}
	}
	return res
}