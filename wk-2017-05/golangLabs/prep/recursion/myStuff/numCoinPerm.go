package main

import "fmt"

func numCoinPerm(val int) (int) {
  fmt.Println(val)
  if val < 0 {
    return 0
  } else if val == 0 {
    return 1
  }

  p := numCoinPerm(val-1)
  n := numCoinPerm(val-5)
  d := numCoinPerm(val-10)
  q := numCoinPerm(val-25)
  return p + n + d + q
}

func main() {
  fmt.Println(numCoinPerm(30))
}