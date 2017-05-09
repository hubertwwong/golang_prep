package main

import "fmt"
//import "math"
import "time"

func trailingZeroes(n int) int {
  if n < 5 {
      return 0
  }
  
  twosA := make([]int, n+1)
  fivesA := make([]int, n+1)
  twos:=0
  fives:=0
  //hits:=0
  
  for i:=2 ; i<=n ; i++ {
    curI := i
    curTwos := 0
    curFives := 0

    // count twos..
    for j:=curI ; j>0 ; {
      if j%2 == 0 {
        if twosA[j/2] != 0 {
          curTwos = twosA[j/2] + 1
          //hits++
          break
        } else {
          curTwos++
          j/=2
        }
      } else {
          break
      }
    }

    // count fives.
    for j:=curI ; j>0 ; {
      if j%5 == 0 {
        if fivesA[j/5] != 0 {
          curFives = fivesA[j/5] + 1
          //hits++
          break
        } else {
          curFives++
          j=j/5
        }
      } else {
        break
      }
    }

    // final
    //fmt.Println(hits)
    twosA[i] = curTwos
    fivesA[i] = curFives
    twos += curTwos
    fives += curFives
  }
  
  //fmt.Println(twos, fives)
  // simple check to see if you have more twos than fives.
  // should be generally true.
  // number of fives should be the number of zeros that you have...
  if twos >= fives {
      return fives
  } else {
      return 0
  }
}

func main() {
  start := time.Now()
  //fmt.Println(trailingZeroes(1808548329))
  fmt.Println(trailingZeroes(99999999))
  elapsed := time.Since(start)
  fmt.Println(elapsed)
}

/*
optimizations
*/