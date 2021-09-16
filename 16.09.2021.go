//https://www.hackerrank.com/challenges/game-of-stones-1/problem

func gameOfStones(n int32) string {
	res := "First"

	if n < 1 || n % 7 == 0 || n % 7 == 1 {
		res = "Second"
	}

	return res
}