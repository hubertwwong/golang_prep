package main

import (
  "fmt"
  "sort"
)

func hasDupe(vals []int) (bool) {
  if len(vals) <= 1 {
    return false
  }

  sort.Ints(vals)

  for cur, prev := 1, 0 ; cur <len(vals) ; cur, prev = cur+1, prev+1 {
    if vals[cur] == vals[prev] {
      return true
    }
  }

  return false
}

func main() {
  result := hasDupe([]int{2,3,4,5,6,7,2,3,4,5,6})
  fmt.Println(result)
}

/*
10 mins to paper this.
googled the go sort...

i don't think i can multiple increment...
- so you can't....
got the for statement wrong...
the sorted int..
- does not return.
- does an in place sort.
this part took 15 mins...

need to actually read the questions...


*/