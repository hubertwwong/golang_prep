package main

import "fmt"
import "time"

// no memo. this is faster.
// func trailingZeroes(n int) int {
//     numZeros := 0

//     for i := 5 ; i <= n ; i+=5 {
//       // assuming that counting 5 factors is good enough.
//       if i%5 == 0 {
//           curNumZeros := 1
//           for curNum := i/5 ; curNum % 5 == 0 ; curNum/=5 {
//             curNumZeros++
//           }
//           numZeros += curNumZeros
//       }
//     }

//     return numZeros
// }

func trailingZeroes(n int) int {
    numZeros := 0

    for i := 5 ; i <= n ; i+=5 {
        // assuming that counting 5 factors is good enough.
        //curNumZeros := 1
        for fac := 5 ; i%fac == 0 ; fac *= 5 {
            //curNumZeros++
						numZeros++
        }
        //numZeros += curNumZeros
    }

    return numZeros
}

// func loopTest(n int) float64 {
//   count := 0.0
//   for i :=0 ; i<n ; i+=5 {
//     count = float64(i) * 1.1
//   }
//   return count
// }

func main() {
  start := time.Now()

  fmt.Println(trailingZeroes(1808548329))
  //fmt.Println(loopTest(1808548329))
  
  elapsed := time.Since(start)
  fmt.Println(elapsed)
}