//https://www.hackerrank.com/challenges/bon-appetit/problem

func bonAppetit(bill []int32, k int32, b int32) {
	total := int32(0);

	for _, val := range bill{
		total += int32(val);
	}

	total = (total-bill[k]) / 2;

	if total == b {
		fmt.Println("Bon Appetit");
	}else{
		fmt.Println(b - total)
	}
	return
}
