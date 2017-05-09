package main

import "fmt"

// Naive impl... Probably should be using a hash map.
// and passing some pointner or using a struct.
func dictContains(s string, d []string) bool {
 for i := 0 ; i< len(d) ; i++ {
    if s == d[i] {
      //fmt.Println("> dictContains > true >", s)
      return true
    }
  }
  //fmt.Println("> dictContains > false >", s)
  return false
}

func wordBreakNaive(s string, d []string) bool {
  //fmt.Println("\n> wordBreakNaive > start", s)
  if len(s) == 0 {
    return true
  }

  // try all prefixes from 1 to n
  // <= vs. < because of using slices.
  for i := 1 ; i <= len(s) ; i++ {
    //fmt.Println("> wordBreakNaive > loop", s[0:i], "|", s[i:], "|", i)
    if dictContains(s[0:i], d) && wordBreakNaive(s[i:], d) {
      return true
    }
  }
  
  // You tried all of the prefix and it didn't work.
  return false
}

func main() {
  d := []string{"dog", "car", "cat"}
  s := "catdog"
  result := false

  //result := wordBreakNaive(s, d)
  //fmt.Println(result)

  s = "foo"
  result = wordBreakNaive(s, d)
  fmt.Println(s, ">", result)

  s = "catdogcat"
  result = wordBreakNaive(s, d)
  fmt.Println(s, ">", result)

  s = "cats"
  result = wordBreakNaive(s, d)
  fmt.Println(s, ">", result)
}

/*

Trying this out.
http://www.geeksforgeeks.org/dynamic-programming-set-32-word-break-problem/

*/