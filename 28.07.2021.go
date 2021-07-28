//https://www.hackerrank.com/challenges/the-hurdle-race/problem

func hurdleRace(k int32, height []int32) int32 {
	doses := int32(0)

	for _, val := range height {
		needed := val - k

		if doses < needed {
			doses = needed
		}
	}
	return doses
}