package main

import "fmt"

// Naive impl... Probably should be using a hash map.
// and passing some pointner or using a struct.
func dictContains(s string, d []string) bool {
  fmt.Println("")
  for i := 0 ; i< len(d) ; i++ {
    if s == d[i] {
      fmt.Println("> dictContains > true >", s)
      fmt.Println("")
      return true
    }
  }

  fmt.Println("> dictContains > false >", s)
  fmt.Println("")
  return false
}

func wordBreak(s string, d []string) bool {
  fmt.Println("\n> wordBreak > start", s)
  if len(s) == 0 {
    return true
  }

  // wb stores the dynamic programming results...
  wb := make([]bool, len(s)+1)

  // try all prefixes from 1 to n
  for i := 1 ; i <= len(s) ; i++ {
    fmt.Println("> wordBreak > loop", s[0:i], "|", s[i:], "|", i)
    // if wb if false, then run the algo.
    // if wb is true, you figured out that that word is in the dict.
    // skip the call.
    if wb[i] == false && dictContains(s[0:i], d) {
      wb[i] = true
    }

    //fmt.Println("> wordBreak > contains", wb)
    fmt.Println("> wordBreak > i", i, "> wb[i]", wb[i])
    // feel like you are splitting up the two parts of the condition.
    if wb[i] == true {
      fmt.Println("> wordBreak > true", i)
      // return if you hit the last prefix.
      if i == len(s) {
        return true
      }

      // This is checking the rest of the string.
      // And the offset is i vs. 0
      // Why is there this second call? This i dont' get.
      for j := i+1 ; j <= len(s) ; j++ {
        fmt.Println("> wordBreak > part duex", i, j, s[i:j])
        if wb[j] == false && dictContains(s[i:j], d) {
          wb[j] = true
        }

        if j == len(s) && wb[j] == true {
          return true
        }
      }
    }
  }
  
  // You tried all of the prefix and it didn't work.
  return false
}

func main() {
  d := []string{"dog", "car", "cat"}
  s := "catdog"
  result := false

  //result := wordBreak(s, d)
  //fmt.Println(result)

  // s = "foo"
  // result = wordBreak(s, d)
  // fmt.Println(s, ">", result)

  s = "catdogcat"
  result = wordBreak(s, d)
  fmt.Println(s, ">", result)

  // s = "cat"
  // result = wordBreak(s, d)
  // fmt.Println(s, ">", result)

  // s = "cats"
  // result = wordBreak(s, d)
  // fmt.Println(s, ">", result)

  // s = "catdog"
  // result = wordBreak(s, d)
  // fmt.Println(s, ">", result)
}

/*

This is the non naive version... I think the indicies are wrong...
the slices are counting different from how they work in other languages. I think...

I think the code works at this point...
Need to understand it.
WB array stores the results to the dictionary check call.

Big thing to understand.... Memoization is 1 larger than the array.


Trying this out.
http://www.geeksforgeeks.org/dynamic-programming-set-32-word-break-problem/

*/