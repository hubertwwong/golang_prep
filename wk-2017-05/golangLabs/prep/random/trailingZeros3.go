package main

import "fmt"
import "math"
import "time"

// no memo. this is faster.
// func trailingZeroes(n int) int {
//     // curNum5
//     num5 := 1
//     numZeros := 0
//     for i := 5 ; i <= n ; i+=5 {
//         // assuming that counting 5 factors is good enough.
//         if i%5 == 0 {
//             curNumZeros := 1
//             curNum := i/num5 * 5:
            
//             // for curNum := i/5 ; curNum % 5 == 0 ; curNum /= 5 {
//             //     curNumZeros++
//             // }
//             numZeros += curNumZeros
//         }
//     }
//     return numZeros
// }

// no memo. this is faster. THIS IS THE CLEAN VERSION.
// func trailingZeroes(n int) int {
//     numZeros := 0

//     for i := 5 ; i <= n ; i+=5 {
//         // assuming that counting 5 factors is good enough.
//         if i%5 == 0 {
//             curNumZeros := 1
//             for curNum := i/5 ; curNum % 5 == 0 ; curNum /= 5 {
//                 curNumZeros++
//             }
//             numZeros += curNumZeros
//         }
//     }

//     return numZeros
//}

// compute powers of 5
func fivesPow(n int) map[int]int {
  powers := make(map[int]int)

  for i:=0 ; int(math.Pow(float64(5), float64(i)))<n ; i++ {
    powers[int(math.Pow(float64(5), float64(i)))] = i
  }

  return powers
}

// no memo. this is faster.
func trailingZeroes(n int) int {
    fives := fivesPow(n)
    numZeros := 0

    // time tracking...
    // var timeOverall time.Duration  
    // numIterations := 0

    for i := 5 ; i <= n ; i+=5 {
        //start := time.Now()

        // assuming that counting 5 factors is good enough.
        if i%5 == 0 {
            curNumZeros := 1
            for curNum := i/5 ; curNum % 5 == 0 ; curNum /= 5 {
                if val, ok := fives[curNum]; ok {
                  curNumZeros+=val
                  break
                } else {
                  curNumZeros++
                }
            }
            numZeros += curNumZeros
        }
        // elapsed := time.Since(start)
        // timeOverall += elapsed
        // numIterations++
    }

    //fmt.Println(">avg time", timeOverall/time.Duration(numIterations))

    return numZeros
}

func main() {
  start := time.Now()

  //fmt.Println(trailingZeroes(1808548329))
  // 452137076 is the answer.
  
  // fmt.Println(fivesPow(1808548329))
  // elapsed := time.Since(start)
  // fmt.Println(elapsed)

  fmt.Println(trailingZeroes(1808548329))
  elapsed := time.Since(start)
  fmt.Println(elapsed)

  //fmt.Println(trailingZeroes(30))
}