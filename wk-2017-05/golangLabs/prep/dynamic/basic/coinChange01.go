package main

import "fmt"

// Need to account for the -1 on the skips.
// a sub problem of this is to count the sub problems..
func sumValues(min, max, target int, vals []int) {
  res := make([][]int, 0)
  
  // outer loop for number of pointers.
  for i := 2 ; i < (max-min) ; i {
    // 0 is the min pointer
    ptrs := make([]int, i)
    
  }

  return res
}

func coinChange(coins []int, amount int) int {
  if len(coins) == 0 {
      return -1
  } else if amount < 0 {
      return -1
  } else if amount == 0 {
      return 0
  }
  
  memo := make([]int, amount+1)
  
  // initial everything to -1
  for i := 0 ; i < len(memo) ; i++ {
      memo[i] = -1
  }
  
  // init the coins to the memo array 
  for _, v := range coins {
      if v < len(memo) {
          memo[v] = 1
      }
  }
  
  // iterate thru each coin amount.
  for i := memo[0]+1 ; i < len(memo) ; i++ {
      //fmt.Println("\n> i >",i)
      // skip.
      if memo[i] == 1 {
          //fmt.Println("> Skipping")
          continue
      }
      
      // compute the possible values...

      // Compute min amount of coins for the i-th position
      curMinCoins := -1
      for j, k := coins[0], i ; j <= k ; {
          // you need a skip mechanism???
          // do you...
          if memo[j] == -1 {
              j++
              continue
          } else if memo[k] == -1 {
              k--
              continue
          }
          
          curAmount := j + k
          // fmt.Println(">", j, k)
          if curAmount == i && (curMinCoins == -1 || curMinCoins > memo[j] + memo[k]) {
              //fmt.Println("> cur amount == i")
              curMinCoins = memo[j] + memo[k]
          } else if curAmount > i {
              k--
          } else {
              j++
          }
      }
      
      // update the min coins.
      memo[i] = curMinCoins
      //fmt.Println(memo)
      
      //fmt.Println(">", i, memo)
  }
  
  return memo[len(memo)-1]
}

func main() {
  var coins []int
  var amt int

  // expected 20
  // coins = []int{186,419,83,408}
  // amt = 6249

  // expected 2
  // coins = []int{1,2,3,4,5}
  // amt = 6

  coins = []int{10,15,30,56}
  amt = 100

  fmt.Println(coinChange(coins, amt))
}

/*

The big thing with this problem right now is that you are not compute all of the previous solutions...
Only using 2 values on your scan back on the memo table. This technically can be N values

*/