package main

import "fmt"

func fibDyn(val int, memo map[int]int) (int, map[int]int) {
	if val == 0 {
		memo[0] = 0
		return 0, memo
	} else if val == 1 {
		memo[1] = 1
		return 1, memo
	}

	fibMinus2 := -1
	fibMinus1 := -1

	if memo[val-2] != 0 {
		fibMinus2 = memo[val-2]
		fmt.Println("2", fibMinus2)
	} else {
		fibMinus2, memo = fibDyn(val-2, memo)
	}

	if memo[val-1] != 0 {
		fibMinus1 = memo[val-1]
		fmt.Println("1", fibMinus1)
	} else {
		fibMinus1, memo = fibDyn(val-1,memo)
	}

	result := fibMinus1 + fibMinus2
	memo[val] = result
	return result, memo
}

func main() {
	memo := make(map[int]int)
	fmt.Println(fibDyn(6, memo))
}

/*

you need to come up with a list of stuff to check
one is return types..
check if the signatures are correct.
assignment. check that you are assigning things..

*/