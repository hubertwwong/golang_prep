package main

import "fmt"

func climbStairs(n int) int {
  vals := make([]int, n)
  return climbStairsDP(n, vals)
}

func climbStairsDP(n int, vals []int) int {
  fmt.Println("")
  fmt.Println("> n", n, " > v >", vals)
  if n < 0 {
    return 0
  } else if n == 0 {
    vals[0] = 0
    return 0
  } else if n == 1 {
    vals[1] = 1
    return 1
  } else if n == 2 {
    vals[2] = 2
    return 2
  }

  result := 0
  if vals[n-1] != 0 {
    //fmt.Println(">", vals[n-1], ">", n-1)
    result = result + vals[n-1]
  } else {
    //fmt.Println(">n-1", n-1)
    n1 := climbStairsDP(n-1, vals)
    //fmt.Println(">n-1 > ret", n1)
    vals[n-1] = n1
    result = result + n1
  }

  if vals[n-2] != 0 {
    //fmt.Println(">", vals[n-2], ">", n-2)
    result = result + vals[n-2]
  } else {
    //fmt.Println(">n-2", n-2)
    n2 := climbStairsDP(n-2, vals)
    //fmt.Println(">n-2 > ret", n2)
    vals[n-2] = n2
    result = result + n2
  }
  //fmt.Println("")
  //fmt.Println("> ret", result, ">n", n)
  return result
}

func main() {
  fmt.Println(climbStairs(1))
}