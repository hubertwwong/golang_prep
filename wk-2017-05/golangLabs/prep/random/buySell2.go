package main

import "fmt"

func maxProfit(prices []int) int {
  lenPrices := len(prices)
  if lenPrices <= 1 {
    return 0
  }

  profit := 0
  s := -1
  b := prices[0]
  pp := -1

  // note. using -1 to mean that value is not set.
  for _, cp := range prices {
    // one off to set the prev price.
    // this should only run once.
    if pp == -1 {
      pp = cp
      continue
    }

    if pp>cp {  // possible time to sell.
      // sell value not set. good to set the new buy point.
      if s == -1 {
        b = cp
      } else {
        // add profit.
        profit = profit + (s-b)
        
        // set a new buy and sell price.
        b = cp
        s = -1
      }
    } else if cp > s {
      // sell price is higher than current price. just set it.
      s = cp
    }

    // set new previous price.
    pp = cp
  }

  // cleanup....
  if s!= -1 {
    profit = profit + (s-b)
  }
  
  // return the profit.
  return profit
}

func main() {
  fmt.Println(maxProfit([]int{1,2,3}))
  fmt.Println(maxProfit([]int{3,2,1}))
  fmt.Println(maxProfit([]int{1}))
  fmt.Println(maxProfit([]int{5,3,1,10,1,10,1,3}))
}

/*

https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/
some paper fixes.
no error in the code....
when i typed it up.
it compiled an returned the corrrect value.


*/