package main

import "fmt"
import "math"
import "time"

func countPrimes (n int) int {
  if n<2 {
    return 0
  }
  // np is the number of primes.
  np := 0

  primes:=make([]int, 0)
  
  for i:=2 ; i<n ; i++ {
    // some optimizations
    
    //fmt.Println(i, n)
    np += isPrime(i, &primes)
  }
  return np
}

func isPrime(n int, primes *[]int) int {
  p := *primes
  nroot := math.Sqrt(float64(n))
  //fmt.Println(">------", n)
  for i := 0 ; i<len(p) ; i++ {
    //fmt.Println(">", n, i, p[i], n%p[i])
    if n%p[i] == 0 {
      //fmt.Println("returning")
      return 0
    }

    // break out of look if i > root.
    // optimization..
    if float64(i) > nroot {
      break
    }
  }
  //fmt.Println(">0", *primes)
  *primes = append(*primes, n)
  //fmt.Println(">1", *primes)
  
  return 1
}

func main() {
  start := time.Now()
  fmt.Println(countPrimes(1500000))
  elapsed := time.Since(start)
  fmt.Println(elapsed)
}

/*
bugs...
logic ones i think...
it only looping once. isPrime that is.
no syntax bugs. but a logic one...

isPrime bugs...
- should just loop thru the whole list of times..
- you should check against n not i

main loop bugs...
- not going to n.
vs. n-1

got prime definition wrong....
didn't read the definition...
should be less than not less than equal to...

and this was most of it...
optimization is the next thing...

*/