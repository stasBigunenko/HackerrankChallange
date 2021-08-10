//https://www.hackerrank.com/challenges/halloween-sale/problem

func howManyGames(p int32, d int32, m int32, s int32) int32 {
	sum := int32(0)
	games := int32(0)

	for {
		if sum + p <= s {
			sum += p
			games++
			p -= d
			if p < m {
				p = m
			}
		}else {
			break
		}
	}
	return games
}