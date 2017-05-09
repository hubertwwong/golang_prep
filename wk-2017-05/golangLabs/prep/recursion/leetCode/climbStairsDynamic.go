package main

import "fmt"

// this is a good technique.
func climbStairs(n int) int {
  memo := make([]int, n)
  return climbStairsMemo(n, memo)
}

func climbStairsMemo(n int, memo []int) int {
  if n <= 0 {
    return 0
  } else if n == 1 {
    return 1
  } else if n == 2 {
    return 2
  }

  ones := 0
  twos := 0
  if memo[n-2] == 0 {
    twos = climbStairsMemo(n-2, memo)
    memo[n-2] = twos
  } else {
    twos = memo[n-2]
  }

  if memo[n-1] == 0 {
    ones = climbStairsMemo(n-1, memo)
    memo[n-1] = ones
  } else {
    ones = memo[n-1]
  }

  return twos+ones
}

func main() {
  fmt.Println(climbStairs(1))
  fmt.Println(climbStairs(2))
  fmt.Println(climbStairs(3))
  fmt.Println(climbStairs(4))
  fmt.Println(climbStairs(5))
  fmt.Println(climbStairs(100))
}

/*

https://leetcode.com/problems/climbing-stairs/

11:55am

*/