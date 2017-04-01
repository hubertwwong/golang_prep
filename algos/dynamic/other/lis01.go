package main

import "fmt"

func lis(vals []int) (int) {
	valsLen := len(vals)
	if valsLen <= 1 {
		return valsLen
	}

	// LIS at n index.
	dp := make([]int, valsLen)

	// outer loop contains best position at dp.
	for i:=0 ; i<valsLen ; i++ {
		// 1 is always the min length.
		dp[i] = 1

		// inner loop cycles thru each item. to check the increasing sub sequese.
		fmt.Println("i", i)
		for j:=i-1 ; j>=0 ; j-- {
			// if the subsequence is larger.
			// fmt.Println(dp[j] + 1, dp[i], vals[j], vals[i])
			if dp[j] + 1 > dp[i] && vals[j] < vals[i] {
				dp[i] = dp[j]+1
			}
		}
	}
	fmt.Println(dp)

	// result. loop thru and figure out the best solution.
	result := 0
	for i:=0 ; i<valsLen ; i++ {
		if dp[i] > result {
			result = dp[i]
		}
	}
	return result
} 

func main() {
	fmt.Println(lis([]int{1,2,3,7,4,9}))
}


/*
largest inreasing sub sequence.
going to try a naive approach..
again..

10:49a
bugs...
assign should be check for new variables...
>=0 on countdown loops
<len on larger loops..

questions...
how to get dp[]......
- this is the memo..
min case is...
dp[0]

having issuesing thinking of the problem...
i get the general idea.

*/