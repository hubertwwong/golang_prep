package main

import "fmt"

func maxSubArray(nums []int) int {
  if len(nums) == 0 {
    return 0
  } else if len(nums) == 1 {
    return nums[0]
  }

  // my guess is that you don't need either array.
  bestSoFar := make([]int, len(nums))
  startPos := make([]int, len(nums))
  
  // init col 0. basically if you only have one thing, thats the best solution so far.
  bestSoFar[0] = nums[0]
  startPos[0] = 0
  
  // dyniamically find the best contigious sum.
  // we start at 1 since we started at zero.
  //bestSoFarI := nums[0] // current best solution.
  
  for i := 0 ; i < len(nums) ; i++ {
    // fmt.Println("\n======= i", i, "nums[i]" ,nums[i])
    // fmt.Println(">b", bestSoFar)
    // fmt.Println(">s", startPos)
    // fmt.Println("nums[i]", nums[i])
    
    // compute sum...
    bestNewSum := nums[i]
    bestLen := 1
    curNewSum := 0
    curLen := 0
    for j := i ; j < len(nums) ; j++ {
      curNewSum += nums[j]
      curLen++
      if curNewSum > bestNewSum {
        bestNewSum = curNewSum
        bestLen = curLen
        //fmt.Println("> New best found. bestNewSum", bestNewSum, "bestLen", bestLen)
      }

      //fmt.Println("> j", j, "curNewSum", curNewSum, "curLen", curLen)
    }
    //fmt.Println("> Final best found. bestNewSum", bestNewSum, "bestLen", bestLen)

    // best sum check
    if i != 0 {
      if bestNewSum > bestSoFar[i-1] {
        //fmt.Println("> Found a better sum. bestNewSum", bestNewSum, "bestLen", bestLen)
        for j := i ; j < i+bestLen ; j++ {
          bestSoFar[j] = bestNewSum
          startPos[j] = i
        }
      } else {
        //fmt.Println("> No new solution found. Copying over.")
        bestSoFar[i] = bestSoFar[i-1]
        startPos[i] = startPos[i-1]
      }
    } else {
      //fmt.Println("> Found initial better sum. bestNewSum", bestNewSum, "bestLen", bestLen)
      for j := i ; j < i+bestLen ; j++ {
        bestSoFar[j] = bestNewSum
        startPos[j] = i
      }
    }

  }

  // return the best possible sum.
  return bestSoFar[len(bestSoFar)-1]
}

func main() {
  var nums []int
  
  // 6
  // nums = []int{-2,1,-3,4,-1,2,1,-5,4}
  
  // 9
  // nums = []int{1,3,5}

  // 1
  // nums = []int{-2,1}

  // 21. pick the last 3
  // nums = []int{8,-19,5,-4,20}

  // 110
  nums = []int{1,2,-1,1,2,3,-1,1,2,100}

  fmt.Println("> output >", maxSubArray(nums))
}

/*

You do need a best pos...
a running best pos..

Passes a 150/200 case on leet code.
I think the issue is that i'm being greedy.
And thats leading to wrong solutions.

This is wrong... But didi pass.
No dynamic programming involved...
Need to think about this.


*/