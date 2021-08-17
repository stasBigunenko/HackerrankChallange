//https://www.hackerrank.com/challenges/tutorial-intro/problem

func introTutorial(V int32, arr []int32) int32 {
	var res int32
	for i := 0; i<len(arr); i++ {
		if arr[i] == V {
		res = int32(i)
		break
		}
	}
	return res
}