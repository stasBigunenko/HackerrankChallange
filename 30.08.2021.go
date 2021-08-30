//https://www.hackerrank.com/challenges/mark-and-toys/problem

//additional import "sort"

func maximumToys(prices []int32, k int32) int32 {
	sort.Slice(prices, func(i, j int) bool { return prices[i] < prices[j] })
	count := int32(0)

	for i := range prices {
		if k >= prices[i] {
			k -= prices[i]
			count++
		}
	}
	return count
}
