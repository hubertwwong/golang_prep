package main

import "fmt"

func printRows(coins [][]int) {
  // header
  for i := 0 ; i<len(coins[0]) ; i++ {
    fmt.Printf("%2d", i)
  }
  fmt.Println("")
  for i:=0 ; i<len(coins) ; i++ {
    fmt.Println(coins[i])
  }
}

func mergeSortI(s []int) []int {
    if len(s) <= 1 {
        return s
    }
    
    l := mergeSortI(s[:len(s)/2])
    r := mergeSortI(s[len(s)/2:])
    return merge2I(l, r)
}

func merge2I(l, r []int) []int {
    merged := make([]int, len(l) + len(r))
    
    li:=0
    ri:=0
    
    for i := 0 ; i<len(merged) ; i++ {
        if li == len(l) {
            merged[i] = r[ri]
            ri++
        } else if ri == len(r) {
            merged[i] = l[li]
            li++
        } else if l[li] < r[ri] {
            merged[i] = l[li]
            li++
        } else {
            merged[i] = r[ri]
            ri++
        }
    }
    
    return merged
}



func coinChange(coins []int, amount int) int {
  sortedCoins := mergeSortI(coins)
  //fmt.Println("> sorted coins")
  //fmt.Println(sortedCoins)
  if amount == 0 {
    return 0
  } else if amount < sortedCoins[0] {
    return -1
  }

  // using the last row of the memo table to store the total number of coins.
  memo := make([][]int, len(sortedCoins)+1)
  for i := 0 ; i < len(memo) ; i++ {
    // +1 since array are zero based.
    memo[i] = make([]int, amount+1)
  }

  // init. values you know of.
  // on each amount that is the coin value, you know it takes one coin.
  for i := 0 ; i<len(sortedCoins) ; i++ {
    // only init the column if it exsit.
    // it could be that the amount is less than certain coin values.
    if sortedCoins[i] < len(memo[i]) {
      memo[i][sortedCoins[i]] = 1                 // update the coin row.
      memo[len(sortedCoins)][sortedCoins[i]] = 1  // update the coin total.
    }
  }

  //fmt.Println("Start after coin init")
  //printRows(memo)

  // being a little lazy on the i.
  // just using the total to compute the value.
  totalCoinI := len(sortedCoins)
  for i := sortedCoins[0] ; i < amount+1 ; i++ {
    //fmt.Println("\n>==== i", i)
    // if total is not zero, the total coins was calcuated, so skip it.
    if memo[totalCoinI][i] != 0 {
      continue
    }

    // Check to see if you can add 1 coin from the previous solutions to come up with min coins for this memo.
    // note that you are starting with the largst coin and go down.
    // Note you are checking every coin... You have to.. Don't be greedy.
    coinIndex := -1
    prevBestCoinCount := -1
    for j := len(sortedCoins)-1 ; j >= 0 ; j-- {
      coinsPrevJ := i - sortedCoins[j]
      //fmt.Println(">==== j", j, "coinsPrevJ", coinsPrevJ, "i", i, "sortedCoins", sortedCoins[j])
      if coinsPrevJ <= 0 {
        //fmt.Println("> skip")
        continue
      }

      //fmt.Println("> prev memo found? > j", coinsPrevJ, "> tci", totalCoinI)
      if memo[totalCoinI][coinsPrevJ] != 0 {
        if coinIndex == -1 {
          //fmt.Println("> memo found")
          coinIndex = coinsPrevJ
          prevBestCoinCount = memo[totalCoinI][coinsPrevJ]
        } else if prevBestCoinCount > memo[totalCoinI][coinsPrevJ] {
          //fmt.Println("> better memo found")
          coinIndex = coinsPrevJ
          prevBestCoinCount = memo[totalCoinI][coinsPrevJ]
        }
      }
    }

    // check if a memoSolution is found
    // update the column if it found.
    if coinIndex != -1 {
      // you found a solution. copy values over. add 1 when the k index is j
      for k := 0 ; k < len(sortedCoins)+1 ; k++ {
        //fmt.Println("k", k, "i", i, "coinIndex", coinIndex)
        memo[k][i] = memo[k][coinIndex]
        if k + 1 == len(sortedCoins)+1 {
          memo[k][i]++
          break
        }
      }
      //fmt.Println("> memo updated on col", i)
      //printRows(memo)
    }
  }

  // grab bottom right value of the array. This cointains the final total.
  //fmt.Println("final")
  printRows(memo)
  result := memo[totalCoinI][amount]
  // if the result was zero, you didn't find a result.
  if result == 0 {
    return -1
  } else {
    return result
  }
}

func main() {
  var coins []int
  var amt int

  // expected 20
  coins = []int{186,419,83,408}
  amt = 6249

  // expected 2
  // coins = []int{1,2,3,4,5}
  // amt = 6

  // coins = []int{1,2}
  // amt = 5

  // coins = []int{10,15,20}
  // amt = 60

  // coins = []int{1}
  // amt = 0

  // coins = []int{3,7,11}
  // amt = 23

  fmt.Println(coinChange(coins, amt))
}

/*
Saw the pic on the youtube video.. on the tushar CS video. 
Not doing the DP correctly..
You needed a double array. not a single.

You are oin the right track.
you need to check if the value is betewwn the coin...

FUCK YES.... !!!!
This passed leetcode. 04/21/2017....

beat 66.67 percent of the solutions 49ms for 180 test cases...

You probably dind't need the merge sort...
...

*/