//https://www.hackerrank.com/challenges/chocolate-feast/problem

func chocolateFeast(n int32, c int32, m int32) int32 {
    	var sum int32 = 0
    	sum = n / c
    
      	for i:=sum; i>=m; i-=m {
       	 	sum++
        	i++
    	}
    	return sum
}
