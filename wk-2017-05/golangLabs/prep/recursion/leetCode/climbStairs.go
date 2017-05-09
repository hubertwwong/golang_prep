package main

import "fmt"

func climbStairs(n int) int {
  if n <= 0 {
    return 0
  } else if n == 1 {
    return 1
  } else if n == 2 {
    return 2
  }

  twos := climbStairs(n-2)
  ones := climbStairs(n-1)
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