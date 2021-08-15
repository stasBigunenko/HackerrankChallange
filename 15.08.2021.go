//https://www.hackerrank.com/challenges/cut-the-sticks/problem

func cutTheSticks(arr []int32) []int32 {
	var sticks []int32

	for i := range arr {
		for j := i; j < len(arr); j++ {
			if arr[j] <= arr[i] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}

	sticks = append(sticks, int32(len(arr)))
	for i:=1; i<len(arr); i++ {
		if arr[i] != arr[i-1] {
			sticks = append(sticks, int32(len(arr)-i))
		}
	}

	return sticks
}