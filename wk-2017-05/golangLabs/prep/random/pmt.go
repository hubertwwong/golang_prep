package main

import "fmt"

func parMatchTable(s string) []int {
  pmt := make([]int, len(s))
  //maxSubStrFound := 0

  for i:=1 ; i<len(s) ; i++ {
    fmt.Println(i)
    pStart := 0
    pEnd := 0
    sStart := i
    //sEnd := len(s)-1
    curMax := 0
    
    // find largest sub string at position i.
    for {
      // loop thru prefix and suffix at a certain size.
      for j, k := pStart, sStart ; ; {
        fmt.Println(">>>", j, k, "i", i)
        if s[j] == s[k] {
          curMax++
        } else {
          break
        }

        // increments
        j++
        k++

        // break at the of the string.
        if k+1 >= i {
          break
        }
      }

      if curMax >= pmt[i] {
        pmt[i] = curMax
      }

      // increase length of substring.
      pEnd++
      sStart--

      // you want to break 1 before the entire string.
      if pEnd + 1 >= i {
        break
      }
    }
  }

  return pmt
}

func main() {
  fmt.Println(parMatchTable("abab"))
}

/*

This is the partial match portion of the KNP algorithm...
so no i++ in array...
unused variable..
don't have a break condition....
- using some -1 index.
on the inner loop... no break condition...

*/