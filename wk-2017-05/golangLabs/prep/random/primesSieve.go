package main

import "fmt"
//import "math"
//import "time"

func countPrimes (n int) int {
  if n<2 {
    return 0
  }

  // default to all false.
  allNums:=make([]bool, n)
  
  // mark
  for m:=2 ; m<len(allNums) ; {
    // mark everything for a given multiple m
    for i:=m+m ; i<len(allNums) ; i+=m {
      allNums[i] = true
    }

    // find the next multiple.
    // note. this is using the 
    for m+=1 ; m<len(allNums) ; m++ {
      if allNums[m] == false {
        break
      }
    }
  }

  // count
  c := 0
  //fmt.Println(allNums)
  for i:=2 ; i<len(allNums) ; i++ {
    if !allNums[i] {
      c++
    }
  }
  return c
}


func main() {
  //start := time.Now()
  fmt.Println(countPrimes(1500000))
  //elapsed := time.Since(start)
  //fmt.Println(elapsed)
}

/*
trying the ancient algo.
the Sieve of Eratosthenes

from primes to prime Sieve
it went from 1.5seconds to 50sm
basically you have a massive space trade off to get to this...

*/